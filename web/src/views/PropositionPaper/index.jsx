import React, { Component } from 'react'
import {Switch, Route,Redirect} from 'react-router-dom'
import CreatePaper from './views/CreatePaper'
import UploadQuestions from './views/UploadQuestions'
import EditerQuestion from './views/EditerQuestion'
import Home from './views/Home/home'

export default class index extends Component {
    render() {
        return (
            <div className="proposition-paper-page" data-component="proposition-paper-page" style={{width:'100%',height:'100%'}}>
                <Switch>
                    <Redirect from="/home/proposition-paper" to="/home/proposition-paper/home" exact></Redirect>
                    <Route path="/home/proposition-paper/home" component={Home} exact></Route>
                    <Route path="/home/proposition-paper/edit-question" component={EditerQuestion} exact></Route>
                    <Route path="/home/proposition-paper/create-paper/:project/:subject/:ability/:content/:type" component={CreatePaper} exact></Route>
                    <Route path="/home/proposition-paper/upload-questions/:project/:subject/:ability/:content/:type" component={UploadQuestions} exact></Route>
                </Switch>
            </div>
        )
    }
}
