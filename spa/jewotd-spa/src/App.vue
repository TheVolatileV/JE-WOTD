<template>
  <div id="app" class="app">
    <div class="register-email">
      <input type="email" class="register-email-input" placeholder="メールアドレスを入力" v-model="email" />
      <button class="main__title__btn" @click.prevent="registerEmail">登録</button>
    </div>
    <router-view/>
  </div>
</template>

<script>
import httpService from './services/http-service'

export default {
  name: 'App',
  data () {
    return {
      email: ''
    }
  },
  methods: {
    registerEmail () {
      httpService.registerEmail(this.email).then(resp => {
        console.log(resp)
      })
    }
  }
}
</script>

<style lang="stylus">
.background-image
  background-image: url(./assets/city-cropped.jpeg);
  background-size: cover;
  -webkit-filter: blur(5px);
  filter: blur(5px);
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;

.app
  position: fixed;
  left: 0;
  top: 0;
  padding-left: 15%;
  padding-top: 10%;
  width: 75%;

.register-email
  margin-left: 60%

.register-email-input
  border: 2px
  border-radius: 9px
  padding: 3px 0px 1px 15px
</style>
