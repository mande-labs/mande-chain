<template>
  <div class="container">
    <div class="row row-sm-revers">
      <div class="col-lg-3">
        <SpTokenTransfer />
      </div>
      <div class="col-lg-3 col-lg-offset-1">
        <Voting />
      </div>
      <div class="col-md-3 col-lg-offset-1">
        <SpAssets />
      </div>
    </div>
    <div class="row row-sm-revers">
      <div class="col-lg-6">
        <SpTokenTransferList />
      </div>
      <div class="col-md-5 col-lg-offset-1">
        <VoteTransactionList />
      </div>
    </div>
  </div>
</template>

<script>
import { SpAssets, SpTokenTransfer, SpTokenTransferList } from '@starport/vue'
import Voting from './Voting.vue'
import VoteTransactionList from './VoteTransactionList.vue'
import { computed, defineComponent } from 'vue'
import { useStore } from 'vuex'

export default defineComponent ({
  name: 'Portfolio',

  components: { SpTokenTransfer, SpAssets, SpTokenTransferList, Voting, VoteTransactionList },

  setup() {
    // store
    let $s = useStore()

    // computed
    let address = computed(() => $s.getters['common/wallet/address'])

    return {
      address
    }
  }
})
</script>

<style scoped>
.row {
  display: flex;
  flex-wrap: wrap;
}
.col {
  flex-grow: 1;
  padding: 20px;
}
</style>
