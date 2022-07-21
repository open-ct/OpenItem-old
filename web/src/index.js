/*
 * @Author: your name
 * @Date: 2021-07-31 16:42:56
 * @LastEditTime: 2021-08-13 01:38:48
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web-client\src\index.js
 */
import React from 'react';
import ReactDOM from 'react-dom';
import './index.less';
import 'antd/dist/antd.less';
import App from './App.jsx';
import store from './store'

ReactDOM.render( <App/>, document.getElementById('root') );

store.subscribe(()=>{
	ReactDOM.render(<App/>,document.getElementById('root'))
})