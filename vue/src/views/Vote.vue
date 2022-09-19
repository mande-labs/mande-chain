<template>
  <div>
    <div class="sp-checkers__main sp-box sp-shadow sp-form-group">
        <form class="sp-checkers__main__form">
          <div class="sp-checkers__main__rcpt__header sp-box-header">
            Cast a vote
          </div>

          <input class="sp-input" placeholder="Vote receiver" v-model="receiver" />
          <br />
          <br />
          <input class="sp-input" placeholder="Vote count" v-model="count" />
          <br />
          <br />
          <input type="radio" id="cast" value="1" v-model="mode" />
          <label for="cast">Cast</label>
          <input type="radio" id="uncast" value="0" v-model="mode" />
          <label for="uncast">Uncast</label>
          <br />
          <br />
          <sp-button @click="submit">Cast</sp-button>
        </form>
    </div>
  </div>
</template>


<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: "Vote",
  data() {
    return {
      creator: "",
      receiver: "",
      count: 0,
      mode: 1
    };
  },
  computed: {

    currentAccount() {
      console.log('curr acc')
      console.log('deps loaded')
      if (this.loggedIn) {
        console.log('logged in')
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
      console.log(result);
    },
  },
};
</script>