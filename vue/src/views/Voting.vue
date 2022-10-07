<template>
  <section>
    <div class="tx">

      <!-- feedbacks -->
      <div v-if="isTxOngoing" class="feedback">
        <div class="loading-spinner">
          <SpSpinner size="46"></SpSpinner>
        </div>
        <div style="width: 100%; height: 24px" />

        <div class="tx-ongoing-title">Opening Keplr</div>

        <div style="width: 100%; height: 8px" />

        <div class="tx-ongoing-subtitle">Sign transaction...</div>
      </div>

      <div v-else-if="isTxSuccess" class="feedback">
        <div class="check-icon">
          <svg
            width="64"
            height="63"
            viewBox="0 0 64 63"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <circle
              cx="32"
              cy="31.5"
              r="29.5"
              stroke="#00CF30"
              stroke-width="4"
              stroke-linecap="round"
            />
            <path
              d="M19 30.1362L28.6557 40L45 23"
              stroke="#00CF30"
              stroke-width="4"
            />
          </svg>
        </div>

        <div style="width: 100%; height: 24px" />

        <div class="tx-feedback-title">Votes submitted</div>

        <div style="width: 100%; height: 40px" />

        <div style="width: 100%">
          <SpButton style="width: 100%" @click="resetTx">Done</SpButton>
        </div>
      </div>

      <div v-else-if="isTxError" class="feedback">
        <div class="warning-icon">
          <svg
            width="58"
            height="54"
            viewBox="0 0 58 54"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M29 44.5625C29.7249 44.5625 30.3125 43.9749 30.3125 43.25C30.3125 42.5251 29.7249 41.9375 29 41.9375C28.2751 41.9375 27.6875 42.5251 27.6875 43.25C27.6875 43.9749 28.2751 44.5625 29 44.5625Z"
              fill="#FE475F"
            />
            <path
              d="M1.4375 52.4375L29 1.25L56.5625 52.4375H1.4375Z"
              stroke="#FE475F"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <path
              d="M29 19.625V34.0625"
              stroke="#FE475F"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <path
              d="M29 44.5625C29.7249 44.5625 30.3125 43.9749 30.3125 43.25C30.3125 42.5251 29.7249 41.9375 29 41.9375C28.2751 41.9375 27.6875 42.5251 27.6875 43.25C27.6875 43.9749 28.2751 44.5625 29 44.5625Z"
              stroke="#FE475F"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
        </div>

        <div style="width: 100%; height: 24px" />

        <div class="tx-feedback-title">Something went wrong</div>

        <div style="width: 100%; height: 16px" />

        <div style="width: 100%; height: 50px" />

        <div style="width: 100%">
          <SpButton style="width: 100%" @click="sendTx">Try again</SpButton>

          <div style="width: 100%; height: 8px" />

          <SpButton style="width: 100%" type="secondary" @click="resetTx"
            >Cancel</SpButton
          >
        </div>
      </div>

      <div v-else>
        <div class="title-wrapper">
          <h3 class="title" style="padding-bottom: 30px">Voting</h3>
        </div>
        <form>
          <div class="input-label">
            Cast vote to
          </div>
          <div class="input-wrapper">
            <input
              v-model="state.tx.receiver"
              class="input"
              :class="{
                error: state.tx.receiver.length > 0 && !validReceiver
              }"
              placeholder="Recipient address"
              :disabled="!hasAnyBalance"
            />
            <div
              v-if="state.tx.receiver.length > 0 && !validReceiver"
              class="error-message"
            >
              Invalid address
            </div>
          </div>
          <br /><br />
          <div class="input-label">
            Vote count
          </div>
          <div class="input-wrapper">
            <input
              v-model="state.tx.count"
              class="input"
              :class="{
                error: state.tx.count.length > 0 && !validVoteCount
              }"
              placeholder="100"
              :disabled="!hasAnyBalance"
            />
            <div
              v-if="state.tx.count.length > 0 && !validVoteCount"
              class="error-message"
            >
              Invalid count
            </div>
          </div>
          <br /><br />
          <div class="input-wrapper">
            <input type="radio" id="cast" value="1" v-model="state.tx.mode" :disabled="!hasAnyBalance"/>&nbsp;
            <label class="input-label" for="cast">Cast</label>
            &nbsp;&nbsp;&nbsp;
            <input type="radio" id="uncast" value="0" v-model="state.tx.mode" :disabled="!hasAnyBalance"/>&nbsp;
            <label class="input-label" for="uncast">Uncast</label>
          </div>
          <br /><br />
          <sp-button :disabled="!ableToTx" @click="sendTx">Submit</sp-button>
        </form>
      </div>
    </div>
  </section>
</template>


<script lang="ts">
import { Bech32 } from '@cosmjs/encoding'
import { computed, defineComponent, reactive, watch } from 'vue'
import { useStore } from 'vuex'
import { useAddress, useAssets } from '@starport/vue/src/composables'
import SpSpinner from '@starport/vue/src/components/SpSpinner'

// types
export interface TxData {
  creator: string,
  receiver: string,
  count: string,
  mode: int
}

export enum UI_STATE {
  'TX_SIGNING' = 300,
  'TX_SUCCESS' = 301,
  'TX_ERROR' = 302
}

export interface State {
  tx: TxData
  currentUIState: UI_STATE
}

export let initialState: State = {
  tx: {
    creator: '',
    receiver: '',
    count: '',
    mode: 1
  },
  currentUIState: UI_STATE.SEND
}

export default defineComponent ({
  name: "Voting",

  components: { SpSpinner },

  setup() {
    // store
    let $s = useStore()

    // state
    let state: State = reactive(initialState)

    // composables
    let { address } = useAddress({ $s })
    let { balances } = useAssets({ $s })

    // actions
    let sendMsgCreateVote = (opts: any) =>
      $s.dispatch('mandechain.voting/sendMsgCreateVote', opts)

    // methods
    let resetTx = (): void => {
      state.tx.creator = ''
      state.tx.receiver = ''
      state.tx.count = ''
      state.tx.mode = 1

      state.currentUIState = UI_STATE.SEND
    }
    let sendTx = async (): Promise<void> => {
      state.currentUIState = UI_STATE.TX_SIGNING

      let payload: any = {
        creator: address.value,
        receiver: state.tx.receiver,
        count: parseFloat(state.tx.count) * 10**6,
        mode: state.tx.mode
      }

      let send

      try {
        send = () =>
          sendMsgCreateVote({
            value: payload,
            fee: [],
            memo: ''
          })

        const txResult = await send()

        if (txResult.code) {
          throw new Error()
        }
        state.currentUIState = UI_STATE.TX_SUCCESS
      } catch (e) {
        console.error(e)
        state.currentUIState = UI_STATE.TX_ERROR
      }
    }
    let isTxOngoing = computed<boolean>(() => {
      return state.currentUIState === UI_STATE.TX_SIGNING
    })
    let isTxSuccess = computed<boolean>(() => {
      return state.currentUIState === UI_STATE.TX_SUCCESS
    })
    let isTxError = computed<boolean>(() => {
      return state.currentUIState === UI_STATE.TX_ERROR
    })
    let validReceiver = computed<boolean>(() => {
      let valid: boolean

      try {
        valid = !!Bech32.decode(state.tx.receiver)
      } catch {
        valid = false
      }
      return valid
    })
    let validVoteCount = computed<boolean>(
      () =>
        !isNaN(state.tx.count) && state.tx.count != 0 && state.tx.count != ''
    )
    let ableToTx = computed<boolean>(
      () =>
          validVoteCount.value &&
          validReceiver.value &&
          !!address.value
    )
    let parseAmount = (amount: string): number => {
      return amount == '' ? 0 : parseInt(amount)
    }
    let hasAnyBalance = computed<boolean>(
      () =>
        balances.value.assets.length > 0 &&
        balances.value.assets.some((x) => parseAmount(x.amount.amount) > 0)
    )

    //watch
    watch(
      () => address.value,
      async () => {
        resetTx()
      }
    )

    return {
      state,
      address,
      resetTx,
      sendTx,
      isTxOngoing,
      isTxSuccess,
      isTxError,
      ableToTx,
      validVoteCount,
      validReceiver,
      hasAnyBalance
    }
  }
});
</script>

<style scoped lang="scss">
section {
  padding-bottom: 132px;
}
.title-wrapper {
  display: flex;
  flex-wrap: wrap;
  width: 100%;
}
.title-wrapper .title {
  font-family: Inter, serif;
  font-style: normal;
  font-weight: 600;
  font-size: 28px;
  line-height: 127%;
  letter-spacing: -0.02em;
  font-feature-settings: "zero";
  color: #000000;
  margin-top: 0;
}
.input-label {
  font-family: Inter;
  font-style: normal;
  font-weight: 400;
  font-size: 13px;
  line-height: 153.8%;
  color: rgba(0, 0, 0, 0.667);
}
.input {
  margin-top: 4px;
  padding: 12px 16px;
  height: 48px;
  background-color: rgba(0, 0, 0, 0.03);
  border-radius: 10px;
  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 130%;
  color: #000000;
  width: 100%;
  outline: 0;
  transition: background-color 0.2s cubic-bezier(0.645, 0.045, 0.355, 1);
  display: block;

  &:not([disabled]) {
    &:hover {
      background: rgba(0, 0, 0, 0.07);
    }
  }

  &:focus {
    background: rgba(0, 0, 0, 0.07);
    color: #000;
  }

  &.error {
    box-shadow: 0 0 0 1px rgba(254, 71, 95, 1);
  }
}
.feedback {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.tx-feedback-title {
  font-family: Inter;
  font-style: normal;
  font-weight: bold;
  font-size: 21px;
  line-height: 152%;
  /* identical to box height, or 32px */

  text-align: center;
  letter-spacing: -0.017em;

  /* light/text */

  color: #000000;
}
.tx-feedback-subtitle.amount {
  text-transform: uppercase;
}
.tx-feedback-subtitle {
  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 150%;
  /* identical to box height, or 24px */

  text-align: center;

  /* light/muted */

  color: rgba(0, 0, 0, 0.667);
}
.tx-ongoing-title {
  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 150%;
  /* identical to box height, or 24px */

  text-align: center;

  /* light/muted */

  color: rgba(0, 0, 0, 0.667);
}

.tx-ongoing-subtitle {
  font-family: Inter;
  font-style: normal;
  font-weight: bold;
  font-size: 21px;
  line-height: 152%;
  /* identical to box height, or 32px */

  text-align: center;
  letter-spacing: -0.017em;

  /* light/text */

  color: #000000;
}
.error-message {
  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 13px;
  line-height: 153.8%;
  color: #d80228;
  margin-top: 5px;
}
.input-wrapper {
  display: block;
}
</style>
