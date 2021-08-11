/*
 * @Author: your name
 * @Date: 2021-07-31 16:42:56
 * @LastEditTime: 2021-08-10 15:02:04
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web-client\src\App.js
 */
import './App.less';
import React, { Component } from 'react'
import {Switch,Route,Redirect,BrowserRouter as Router} from 'react-router-dom'
import Login from './views/Login'
import Home from './views/Home'


export default class App extends Component {
  getRemSize = ()=>{
    let whdef = 100/1920
    let wW = window.innerWidth
    let rem = wW * whdef
    document.documentElement.style.fontSize = rem + 'px'
  }

  componentDidMount = ()=>{
    window.resize = ()=>{
      this.getRemSize()
    }
    this.getRemSize()
  }

  render() {
    return (
      <div id="App">
        <Router>
          <Switch>
            <Redirect from="/" to="/login" exact></Redirect>
            <Route path="/login" component={Login} exact></Route>
            <Route path="/home" component={Home}></Route>
          </Switch>
        </Router>
      </div>
    )
  }
}

