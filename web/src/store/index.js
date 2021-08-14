/*
 * @Author: your name
 * @Date: 2021-08-13 01:12:32
 * @LastEditTime: 2021-08-13 01:23:29
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web\src\store\user\index.js
 */
import { createStore } from 'redux';
import reducer from './reducer';

const store = createStore(reducer);

export default store;
