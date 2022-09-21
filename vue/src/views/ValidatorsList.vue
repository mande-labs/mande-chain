<template>
  <div class="validators-list">
    <div class="title-wrapper">
      <div class="title">Validators</div>
    </div>
    <div v-if="!address || state.validators.length <= 0" class="empty">
      Validators list is empty
    </div>
    <div v-else-if="state.validators.length > 0" class="list">
        <div class="val-list">
          <div class="val-header">
            <div class="val-container row">
              <div class="col-lg-3 validator-name val-text"> Name </div>
              <div class="col-lg-5 validator-address val-text"> Address </div>
              <div class="col-lg-3 validator-cred val-text"> Credibility </div>
            </div>
          </div>
          <div v-for="(vdata) in state.validators" :key="`${index}`">
            <div class="val-row">
              <div class="val-container row">
                <div class="col-lg-3 validator-name val-text"> {{vdata.moniker}} </div>
                <div class="col-lg-5 validator-address val-text"> {{shortAddr(vdata.operatorAddress)}} </div>
                <div class="col-lg-3 validator-cred val-text"> {{vdata.tokens}} </div>
              </div>
            </div>
          </div>
        </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, watch } from 'vue'
import { useStore } from 'vuex'
import { useAddress } from '@starport/vue/src/composables'

export interface Validator {
  operatorAddress: string,
  status: string,
  tokens: string,
  moniker: string
}

export interface State {
  validators: Array<Validator>
}

export let initialState: State = {
  validators: []
}

export default defineComponent ({
  name: 'ValidatorsList',

  setup() {
    // store
    let $s = useStore()
    let { address } = useAddress({ $s })
    let state: State = reactive(initialState)

    let loadNewItems = async () => {
      let res = await $s.dispatch('cosmos.staking.v1beta1/QueryValidators', {
          subscribe: true
      })

      let validatorsList = computed(() => $s.getters['cosmos.staking.v1beta1/getValidators']())

      console.log('check below')
      console.log(validatorsList)

      var bufferValidators: Array<Validator> = []
      var val: Validator = {
        operatorAddress: '',
        status: '',
        tokens: '',
        moniker: ''
      }

      for (let i=0; i<validatorsList.value.validators.length; i++) {
        val.operatorAddress = validatorsList.value.validators[i].operator_address
        val.status = validatorsList.value.validators[i].status
        val.tokens = validatorsList.value.validators[i].tokens
        val.moniker = validatorsList.value.validators[i].description.moniker

        bufferValidators.push(val)
      }

      state.validators = bufferValidators
    }

    let shortAddr = (addr) => {
      return addr.substring(0, 10) + '...' + addr.slice(-4)
    }

    watch(
      () => address.value,
      async () => {
        await loadNewItems()
      }
    )

    return {
      state,
      address,
      shortAddr
    }
  }
})
</script>

<style scoped lang="scss">
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
  margin-bottom: 32px;
}
.empty {
  /* Body/M */

  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 150%;
  /* identical to box height, or 24px */

  /* light/muted */

  color: rgba(0, 0, 0, 0.667);
}
.val-list {
  display: flex;
  flex-direction: column;
}
.val-header {
  height: 4rem;
}
.val-row {
  background: #fff;
  position: relative;
  height: 7rem;

  &:after {
    content: '';
    display: none;
    width: calc(100% + 32px);
    height: 60px;
    background: rgba(0, 0, 0, 0.03);
    border-radius: 8px;
    position: absolute;
    top: -20px;
    left: -16px;
  }

  &:hover {
    &:after {
      display: block;
    }
  }

  &.header {
    font-size: 25px;
  }
}

.val-container {
  position: relative;
  display: flex;
  align-items: center;

  .val-text {
    font-family: Inter;
    font-style: normal;
    font-weight: 500;
    font-size: 13px;
    line-height: 153.8%;
    color: #000000;
  }
}

</style>
