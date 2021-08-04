/*
 * @Author: your name
 * @Date: 2021-07-31 16:42:56
 * @LastEditTime: 2021-08-04 00:48:13
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web-client\src\index.js
 */
import React from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter as Router} from 'react-router-dom'
import './index.less';
import 'antd/dist/antd.less';
import App from './App.jsx';

ReactDOM.render(
  // <React.StrictMode>
    <Router>
      <App />
    </Router>
  // </React.StrictMode>
  ,
  document.getElementById('root')
);

