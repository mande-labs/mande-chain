import axios, { AxiosResponse } from 'axios'
import { EventEmitter } from 'events'
import { computed, ComputedRef, Ref, ref, watch } from 'vue'
import { Store } from 'vuex'

import { Amount } from '@/utils/interfaces'

import useAPIPagination, {
  merge,
  Pager,
  Response as APIPagination
} from './useAPIPagination'

type Response = {
  filterSupportedTypes: (tx: object) => boolean
  normalize: (tx: object) => VoteTxForUI
  pager: ComputedRef<Pager>
  newTxs: Ref<number>
}

type Params = {
  $s: Store<any>
  opts: {
    order: 'asc' | 'desc'
    realTime: boolean
  }
}

type TxMode = "0" | "1"
type TxOperation = 'positive' | 'negative'

export type VoteTxForUI = {
  mode: TxMode
  op: TxOperation
  sender: string
  receiver: string
  amount: Amount[]
  hash: string
  type: string
  timestamp: string
  height: number
}

export default async function ({
  $s,
  opts: { order, realTime }
}: Params): Promise<Response> {
  // methods
  let normalizeAPIResponse = (resp: AxiosResponse) => {
    let { txs, tx_responses, pagination } = resp.data

    let merged = txs.map((tx, i) => {
      return { ...tx, ...tx_responses[i] }
    })

    return {
      data: merged,
      total: Number(pagination.total)
    }
  }
  let filterSupportedTypes = (tx: any) => {
    let isVoteCreated = (tx.body.messages[0]['@type'] as string).includes(
      'mandechain.voting.MsgCreateVote'
    )

    return isVoteCreated
  }
  let normalize = (tx: any): VoteTxForUI => {
    let findOutOp = (tx: VoteTxForUI): TxOperation => {
      let op: TxOperation = 'positive'
      if (tx.amount < 0) {
        op = 'negative'
      }

      return op
    }

    let normalized: any = {}

    normalized.sender = tx.body.messages[0].creator
    normalized.receiver = tx.body.messages[0].receiver
    normalized.amount = tx.body.messages[0].count
    normalized.height = Number(tx.height)

    normalized.type = tx.body.messages[0]['@type']
    normalized.timestamp = tx.timestamp
    normalized.hash = tx.txhash
    normalized.mode = tx.body.messages[0].mode
    normalized.op = findOutOp(normalized)

    return normalized as VoteTxForUI
  }
  let fetchTxs = async (offset: number, event: string, orderParam: number) =>
    axios.get(
      `${API_COSMOS.value}` +
        `/cosmos/tx/v1beta1/txs` +
        `?events=${event}` +
        `&pagination.offset=${offset}` +
        `&order_by=${orderParam}`
    )

  // store
  let address = computed<string>(() => $s.getters['common/wallet/address'])
  let client = computed<EventEmitter>(() => $s.getters['common/env/client'])
  let API_COSMOS = computed<string>(() => $s.getters['common/env/apiCosmos'])

  // computed
  let CASTED_EVENT = computed<string>(
    () => `vote_casted.sender%3D%27${address.value}%27`
  )
  let UNCASTED_EVENT = computed<string>(
    () => `vote_uncasted.sender%3D%27${address.value}%27`
  )

  // state
  let orderParam = order === 'asc' ? 1 : 2
  let newTxs = ref(0)

  // composables
  let castedAPIPagination: APIPagination = await useAPIPagination({
    opts: {},
    getters: {
      fetchList: async ({ offset }) =>
        normalizeAPIResponse(
          await fetchTxs(offset, CASTED_EVENT.value, orderParam)
        )
    }
  })
  let uncastedAPIPagination: APIPagination = await useAPIPagination({
    opts: {},
    getters: {
      fetchList: async ({ offset }) =>
        normalizeAPIResponse(
          await fetchTxs(offset, UNCASTED_EVENT.value, orderParam)
        )
    }
  })

  await castedAPIPagination.pager.load()
  await uncastedAPIPagination.pager.load()

  // computed
  let castedAndUncastedPager: ComputedRef<Pager> = computed(() =>
    merge(castedAPIPagination.pager, uncastedAPIPagination.pager)
  )

  //watch
  watch(
    () => address.value,
    async () => {
      await castedAndUncastedPager.value.load()
    }
  )

  if (realTime) {
    client.value.on('newblock', async () => {
      // there's got bet a better way to diff latest vs. current while sparing this wasted round-trip
      let casted = await fetchTxs(0, CASTED_EVENT.value, orderParam)
      let uncasted = await fetchTxs(0, UNCASTED_EVENT.value, orderParam)

      let currentTotal = castedAndUncastedPager.value.total.value
      let latestTotal = 
        Number(casted.data.pagination.total) + Number(uncasted.data.pagination.total)
      let diff = latestTotal - currentTotal

      newTxs.value = diff
    })
  }

  return {
    newTxs,
    pager: castedAndUncastedPager,
    normalize,
    filterSupportedTypes
  }
}
