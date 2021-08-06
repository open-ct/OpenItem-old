/*
 * @Author: your name
 * @Date: 2021-07-31 16:42:56
 * @LastEditTime: 2021-08-05 01:29:29
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web-client\src\App.js
 */
import './App.less';
import React, { Component } from 'react'
import {Switch,Route} from 'react-router-dom'
import Login from './views/Login'
import Home from './views/Home'

export default class App extends Component {
  render() {
    return (
      <div id="App">
        <Switch>
          <Route path="/" component={Login} exact />
          <Route path="/login" component={Login}></Route>
          <Route page="/home" component={Home}></Route>
        </Switch>
      </div>
    )
  }
}
