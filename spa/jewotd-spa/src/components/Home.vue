<template>
  <div class="main">
    <h1 class="main__title">今日の単語</h1>
    <div class="main__container">
      <table class="main__body">
        <tbody>
          <tr>
            <td>日本語：</td>
            <td>{{japanese}}</td>
          </tr>
          <tr>
            <td>英語：</td>
            <td>{{english}}</td>
          </tr>
          <tr>
            <td>品詞：</td>
            <td>{{pos}}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import httpService from '../services/http-service'

export default {
  name: 'HelloWorld',
  data () {
    return {
      japanese: '行く',
      english: 'to go',
      pos: '動詞'
    }
  },
  mounted () {
    httpService.getWords().then(resp => {
      console.log(resp)
      this.japanese = resp.japanese
      this.english = resp.english
      this.pos = resp.partOfSpeech
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style  lang="stylus">
@charset "utf-8"

.main
  background-color: rgba(255,255,255,.65)
  width: 70%
  margin: 0 auto
  margin-top: 5%
  font-family: "Yu Gothic"
  border-radius: 30px

  &__title
    padding-top: 20px
    width: 50%
    margin: auto
    text-align: center
    border-bottom: 3px solid #808080
    font-size: 45px
    font-weight: lighter

  &__body
    margin: 0 auto
    font-size: 20px

  &__container
    padding-top: 20px
    padding-bottom: 30px
</style>
