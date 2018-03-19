import vue from './configured-vue'

const publicMethods = {
  getWords () {
    return vue.http.get('https://jewotd-api.herokuapp.com/api/v1').then(resp => {
      return resp.body
    })
  },
  getForcedWord () {
    return vue.http.get('https://jewotd-api.herokuapp.com/api/v1/force').then(resp => {
      return resp.body
    })
  }
}
export default publicMethods
