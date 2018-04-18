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
  },
  registerEmail (email) {
    var dataOut = {'email': email}
    return vue.http.post('https://jewotd-api.herokuapp.com/api/v1/register', JSON.stringify(dataOut)).then(resp => {
      return resp
    })
  }
}
export default publicMethods
