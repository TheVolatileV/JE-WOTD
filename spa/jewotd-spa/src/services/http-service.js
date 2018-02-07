import vue from './configured-vue'

const publicMethods = {
  getWords () {
    return vue.http.get('http://localhost:80/api/v1').then(resp => {
      return resp.body
    })
  }
}
export default publicMethods
