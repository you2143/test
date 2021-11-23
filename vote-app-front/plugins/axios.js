import axios from 'axios'

export default ({ $axios }, inject) => {
    const _axios = axios.create()

    inject('_axios', _axios)
}