/*
 * @Author: your name
 * @Date: 2021-08-13 00:58:37
 * @LastEditTime: 2021-08-14 00:17:55
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web\src\store\user\reduce.js
 */
import { SET_USERINFO } from './constants'
import { getCookie, setCookie } from '../utils/cookies'
import { getStorage, setStorage } from '../utils/storage'
 
const userState = {
   token:getCookie('token'),
   userInfo:getStorage('userInfo')
}

export default function userReduce(state=userState,action){
    switch(action.type){
      case SET_USERINFO:
        setCookie('token',action.data.token)
        setStorage('userInfo',action.data.user_info)
        return {...state, userInfo:getStorage('userInfo'),token:getCookie('token')}
      default:
        return state
    }
}