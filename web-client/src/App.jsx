/*
 * @Author: your name
 * @Date: 2021-07-31 16:42:56
 * @LastEditTime: 2021-08-04 00:50:22
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web-client\src\App.js
 */
import './App.less';
import React, { Component } from 'react'
import {Switch,Route,Redirect} from 'react-router-dom'
import Login from './views/login'
import Home from './views/home'

export default class App extends Component {
  render() {
    return (
      <div id="App">
        <Switch>
          <Route path="/login" component={Login}></Route>
          <Route page="/home" component={Home}></Route>
          <Redirect to="/login"/>
        </Switch>
      </div>
    )
  }
}
