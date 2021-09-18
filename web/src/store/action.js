/*
 * @Author: your name
 * @Date: 2021-08-13 01:12:25
 * @LastEditTime: 2021-09-17 13:47:42
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web\src\store\user\action.js
 */

import { SET_USERINFO,SET_PROJECTBASEINFO } from './constants'

export const setuserinfo = (data)=>({type:SET_USERINFO,data})
export const setprojectbaseinfo = (data)=>({type:SET_PROJECTBASEINFO,data})
