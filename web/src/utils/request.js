/*
 * @Author: your name
 * @Date: 2021-08-12 00:21:27
 * @LastEditTime: 2021-08-12 00:26:38
 * @LastEditors: Please set LastEditors
 * @Description: axios拦截器配置
 * @FilePath: \OpenItem\web\src\utils\request.js
 */

import axios from 'axios'
import {baseURL} from '../conf'
import {getCookie} from './cookies'


const request = axios.create({baseURL})


request.interceptors.request.use(
    config => {
      if (checkToken()) {
        config.headers['Authorization'] = `Bearer ${getCookie('jwt') }`
      }
      return config
    },
    error => {
      return Promise.reject(error)
    }
)

request.interceptors.response.use(
  response => {
    //const code = response.data.code
    const res = response.data
    return res
  },
  error => {
    return Promise.reject(error.response.data)
  }
)

export default request
