import axios from 'axios'

const instance = axios.create({
    baseURL:"http://192.168.0.105:8081"  // 一定要加http://
})


export default instance