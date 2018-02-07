import vue from 'vue'
// import moment from 'moment'
import vueResource from 'vue-resource'
import vueRouter from 'vue-router'

vue.config.devtools = false
// vue.filter('cleanDate', value => moment(value).format('M/D/YY'))
vue.use(vueResource)
vue.use(vueRouter)
export default vue
