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
              <div class="col-lg-6 validator-address val-text"> Address </div>
              <div class="col-lg-3 validator-cred val-text"> Credibility </div>
            </div>
          </div>
          <div v-for="(vdata) in state.validators" :key="`${index}`">
            <div class="val-row">
              <div class="val-container row">
                <div class="col-lg-3 validator-name val-text"> {{vdata.moniker}} </div>
                <div class="col-lg-6 validator-address val-text"> {{shortAddr(vdata.operatorAddress)}} 
                  <span class="copy-btn" @click="copyAddress(vdata.operatorAddress)">
                    <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 115.77 122.88" style="enable-background:new 0 0 115.77 122.88;height:10px" xml:space="preserve">
                      <g><path class="st0" d="M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02 v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02 c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1 c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7 h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01 c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65 v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01 c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02 v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z"/></g>
                    </svg>
                  </span>
                </div>
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
import { Bech32, toHex, fromHex } from '@cosmjs/encoding'
import { useToast } from "vue-toastification";

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
    let toast = useToast();

    let loadNewItems = async () => {
      let res = await $s.dispatch('cosmos.staking.v1beta1/QueryValidators', {
        options: { subscribe: true },
      })

      let validatorsList = computed(() => $s.getters['cosmos.staking.v1beta1/getValidators']())

      var bufferValidators: Array<Validator> = []
      var val: Validator = {
        operatorAddress: '',
        status: '',
        tokens: '',
        moniker: ''
      }

      let addrPrefix = $s.getters["common/env/addrPrefix"]

      for (let i=0; i<validatorsList.value.validators.length; i++) {
        val.operatorAddress = Bech32.encode(addrPrefix, fromHex(toHex(Bech32.decode(validatorsList.value.validators[i].operator_address).data)))
        val.status = validatorsList.value.validators[i].status
        val.tokens = validatorsList.value.validators[i].tokens / 10**6
        val.moniker = validatorsList.value.validators[i].description.moniker

        bufferValidators.push(val)
      }

      state.validators = bufferValidators
    }

    let shortAddr = (addr) => {
      return addr.substring(0, 10) + '...' + addr.slice(-4)
    }

    let copyAddress = (addr) => {
      const clipboardData =
        event.clipboardData ||
        window.clipboardData ||
        event.originalEvent?.clipboardData ||
        navigator.clipboard;

      clipboardData.writeText(addr);
      toast.success("Copied!", {
        timeout: 1500
      });
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
      shortAddr,
      copyAddress
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

  .validator-address {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;

    .copy-btn {
      z-index: 100;
      cursor: pointer;
    }
  }
}
</style>
