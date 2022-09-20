<template>
  <section>
      <div class="title-wrapper">
        <h3 class="title" style="padding-bottom: 22px">Voting</h3>
      </div>
      <form>
        <div class="input-label">
          Cast vote to
        </div>

        <div class="input-wrapper">
          <input
            v-model="receiver"
            class="input"
            placeholder="Recipient address"
          />
        </div>
        <br />
        <br />
        <div class="input-label">
          Vote amount
        </div>
        <div class="input-wrapper">
          <input
            v-model="count"
            class="input"
            placeholder="100"
          />
        </div>
        <br />
        <br />
        <input type="radio" id="cast" value="1" v-model="mode" />&nbsp;
        <label class="input-label" for="cast">Cast</label>
        &nbsp;&nbsp;&nbsp;
        <input type="radio" id="uncast" value="0" v-model="mode" />&nbsp;
        <label class="input-label" for="uncast">Uncast</label>
        <br />
        <br />
        <sp-button @click="submit">Submit</sp-button>
      </form>
  </section>
</template>


<script>
import { computed, defineComponent } from 'vue'
import { useStore } from 'vuex'

export default defineComponent ({
  name: "Voting",
  data() {
    return {
      creator: "",
      receiver: "",
      count: "",
      mode: 1
    };
  },
  computed: {
    currentAccount() {
      if (this.loggedIn) {
        return this.$store.getters['common/wallet/address']
      } else {
        return null
      }
    },
    loggedIn() {
      return this.$store.getters['common/wallet/loggedIn']
    }
  },
  methods: {
    async submit() {
      const value = {
        creator: this.currentAccount,
        receiver: this.receiver,
        count: this.count,
        mode: this.mode
      };
      const result = await this.$store.dispatch("mandechain.voting/sendMsgCreateVote", {
        value,
        fee: [],
        memo: '',
      });
      console.log(result.rawLog);
    },
  },
});
</script>

<style scoped>
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
}
.input:not([disabled]):hover {
    background: rgba(0, 0, 0, 0.07);
}
</style>
