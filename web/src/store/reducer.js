/*
 * @Author: your name
 * @Date: 2021-08-13 00:58:37
 * @LastEditTime: 2021-09-17 13:46:36
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web\src\store\user\reduce.js
 */
import { SET_PROJECTBASEINFO, SET_USERINFO } from './constants'
import { getCookie, setCookie } from '../utils/cookies'
import { getStorage, setStorage } from '../utils/storage'
 
const userState = {
   token:getCookie('token'),
   userInfo:getStorage('userInfo'),
   currentProjectBaseInfo:getStorage('projectBaseInfo')
}

export default function userReduce(state=userState,action){
    switch(action.type){
      case SET_USERINFO:
        setCookie('token',action.data.message)
        setStorage('userInfo',{...action.data.data.profile,Id:action.data.data.uuid})
        return {...state, userInfo:getStorage('userInfo'),token:getCookie('token')}
      case SET_PROJECTBASEINFO:
        setStorage('projectBaseInfo',action.data)
        return {...state, userInfo:getStorage('userInfo'),token:getCookie('token'),projectBaseInfo:getStorage('projectBaseInfo')}
      default:
        return state
    }
}