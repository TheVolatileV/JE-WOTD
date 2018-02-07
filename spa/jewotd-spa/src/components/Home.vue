<template>
  <div class="main">
    <h1 class="main__title">今日の単語</h1>
    <div>
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
.main {
  background-color: gainsboro
  width: 75%
  margin: auto
  margin-top: 5%
  font-family: "Yu Gothic"

  &__title {
    width: 50%
    margin: auto
    text-align: center
    border-bottom: 3px solid gray
  }
  &__body {
    margin-left: 5%
    font-size: 20px
  }
}
</style>
