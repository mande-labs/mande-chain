<template>
  <div class="tx-row">
    <span class="copy-btn" @click="copyAddress(tx.receiver)">
      <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 115.77 122.88" style="enable-background:new 0 0 115.77 122.88;height:10px" xml:space="preserve">
        <g><path class="st0" d="M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02 v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02 c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1 c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7 h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01 c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65 v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01 c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02 v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z"/></g>
      </svg>
    </span>
    <div class="tx-container">
      <div class="tx-icon" :class="dirDescription">
        <svg
          width="16"
          height="16"
          viewBox="0 0 16 16"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M8 13V3"
            stroke="black"
            stroke-width="1.6"
            stroke-miterlimit="10"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M4 6L8 2L12 6"
            stroke="black"
            stroke-width="1.6"
            stroke-miterlimit="10"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </div>

      <div style="width: 16px; height: 100%" />

      <div class="tx-info">
        <div class="tx-direction">
          {{ dirDescription }}
        </div>
        <div class="tx-meta">
          {{ txDate }}
        </div>
      </div>
      <div class="tx-payload">
        <template v-if="tx.amount">
          <span
            :key="`${tx.amount}-${index}`"
            class="tx-amount"
            :class="dirDescription"
          >
            {{ amountSign + ' ' + Math.abs(tx.amount) / 10**6 }}
          </span>
        </template>
        <div class="tx-denom">
          {{ 'to ' + shortAddr }}
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from 'vue'

import { VoteTxForUI } from '../composables/useVoteTxs'
import { useToast } from "vue-toastification"

enum MODE_DESC {
  "1" = 'Casted',
  "0" = 'Uncasted'
}

enum AMOUNT_SIGN {
  positive = '+',
  negative = '-'
}

export default defineComponent ({
  name: 'VoteTransactionListItem',

  props: {
    tx: {
      type: Object as PropType<VoteTxForUI>,
      required: true
    }
  },

  setup(props) {
    // computed
    let dirDescription = computed<string>(() => MODE_DESC[props.tx.mode])
    let addr = computed<string>(() =>
      props.tx.receiver
    )
    let shortAddr = computed<string>(
      () => props.tx.receiver.substring(0, 10) + '...' + props.tx.receiver.slice(-4)
    )
    let toast = useToast();
    let txDate = computed<string>(() => {
      let date = new Date(props.tx.timestamp)
      return date.toLocaleString('en-US', {
        month: 'long',
        day: 'numeric',
        hour: 'numeric',
        minute: 'numeric'
      })
    })
    let amountSign = computed<string>(() => AMOUNT_SIGN[props.tx.op])

    let copyAddress = async (addr) => {
      const clipboardData =
        event.clipboardData ||
        window.clipboardData ||
        event.originalEvent?.clipboardData ||
        await navigator.clipboard;

      if (!clipboardData) {
        alert("Validator address - " + addr);
      }

      clipboardData.writeText(addr);
      toast.success("Copied!", {
        timeout: 1500
      });
    }

    return {
      //computed
      shortAddr,
      dirDescription,
      amountSign,
      txDate,
      copyAddress
    }
  }
})
</script>

<style scoped lang="scss">
.tx-row {
  background: #fff;
  margin-bottom: 28px;
  position: relative;

  &:after {
    content: '';
    display: none;
    width: calc(100% + 32px);
    height: 80px;
    background: rgba(0, 0, 0, 0.03);
    border-radius: 8px;
    position: absolute;
    top: -10px;
    left: -16px;
  }

  &:hover {
    &:after {
      display: block;
    }
  }
}

.tx-container {
  position: relative;
  display: flex;
  align-items: center;
}

.tx-meta {
  /* Body/S */

  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 13px;
  line-height: 153.8%;
  /* identical to box height, or 20px */

  /* light/muted */

  color: rgba(0, 0, 0, 0.667);
}
.tx-info {
  display: flex;
  flex-direction: column;
  flex: 1;
}
.tx-direction {
  /* Label/S */

  font-family: Inter;
  font-style: normal;
  font-weight: 500;
  font-size: 13px;
  line-height: 153.8%;
  /* identical to box height, or 20px */

  /* light/text */

  color: #000000;
}

.tx-payload {
  text-align: right;
  .tx-amount {
    &:not(:first-child) {
      margin-left: 16px;
    }
  }
}

.tx-amount {
  /* Label/S */

  font-family: Inter;
  font-style: normal;
  font-weight: 500;
  font-size: 13px;
  line-height: 153.8%;
  /* identical to box height, or 20px */

  text-align: right;

  /* light/text */

  color: #000000;
}

.tx-denom {
  /* Body/S */

  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 13px;
  line-height: 153.8%;
  /* identical to box height, or 20px */

  text-align: right;

  /* light/muted */

  color: rgba(0, 0, 0, 0.667);

  padding-right: 12px;
}

.tx-icon {
  width: 32px;
  height: 32px;
  left: 1px;
  top: 14px;

  /* light/fg */

  background: rgba(0, 0, 0, 0.03);
  border-radius: 8px;

  display: flex;
  align-items: center;
  justify-content: center;
}

.tx-icon.Uncasted {
  transform: rotate(180deg);
}

.tx-amount.Casted {
  color: #008223;
}
.copy-btn {
  z-index: 100;
  cursor: pointer;
  position: absolute;
  top: 58%;
  right: 0;
}
</style>
