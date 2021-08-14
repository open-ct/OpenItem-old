/*
 * @Author: your name
 * @Date: 2021-08-12 00:21:27
 * @LastEditTime: 2021-08-14 14:41:42
 * @LastEditors: Please set LastEditors
 * @Description: axios拦截器配置
 * @FilePath: \OpenItem\web\src\utils\request.js
 */

import axios from 'axios'
import {getCookie,checkCookie} from './cookies'
//import {baseURL} from '../conf'
const request = axios.create()


request.interceptors.request.use(
    config => {
      if (checkCookie('token')) {
        //Authorization Bearer 
        config.headers['Token'] = getCookie('token')
      }
      return config
    },
    error => {
      return Promise.reject(error)
    }
)

request.interceptors.response.use(
  response => {
    const code = response.data.code
    if(code===1000){
      const res = response.data.data
      return res
    }else if(code===2000){
      return Promise.reject(response.data)
    }else{
      return Promise.reject(response.data)
    }
  },
  error => {
    return Promise.reject(error.response.data)
  }
)

export default request
