import React, { Component } from 'react'
import { Layout, Menu, Avatar, Button } from 'antd';
import DocumentTitle from 'react-document-title'
import {Switch, Route,Redirect} from 'react-router-dom'
import {  UserOutlined } from '@ant-design/icons';
import PendingTasks from '../PendingTasks'
import ProjectManagement from '../ProjectManagement'
import PropositionPaper from '../PropositionPaper'
import NotFound from '../errors/404'
import Forbid from '../errors/402'
import store from '../../store';
import './index.less'


const { Header , Content } = Layout;

export default class index extends Component {

    state = {
        menuCurrentSelected:"pending-tasks"
    }

    funcChange = (e)=>{
        this.setState({
            menuCurrentSelected:e.key
        })
        this.props.history.push(`/home/${e.key}`)
    }

    render() {
        console.log(store.getState())
        return (
            <DocumentTitle title="命题与题库系统">
                <Layout className="home-page" data-component="home-page">
                    <Header>
                        <div className="left-box">
                            <div className="logo">
                                <div className="circular"></div>
                            </div>
                            <span>Landing</span>
                        </div>
                        <div className="right-box">
                            <div className="header-menu">
                            <Menu onClick = {this.funcChange} defaultSelectedKeys={['pending-tasks']} mode="horizontal" selectedKeys={[this.state.menuCurrentSelected]}>
                                <Menu.Item key="pending-tasks">
                                    项目管理
                                </Menu.Item>
                                <Menu.Item key="personnel-management">
                                    人员管理
                                </Menu.Item>
                                <Menu.Item key="proposition-paper">
                                    命题组卷
                                </Menu.Item>
                                <Menu.Item key="system-settings">
                                    系统设置
                                </Menu.Item>
                            </Menu>
                            </div>
                            <div className="button-group">
                                <Button ghost size="small">操作1</Button>
                                <Button ghost size="small">工作台</Button>
                            </div>
                            <div className="user">
                                <Avatar icon={<UserOutlined />} size={24}/>
                                {/* <span>{store.getState().userInfo.name}</span> */}
                                <span>stacker</span>
                            </div>
                        </div>
                    </Header>
                    <Content>
                        <Switch>
                            <Redirect from="/home" to="/home/pending-tasks" exact></Redirect>
                            <Route path="/home/pending-tasks" component={PendingTasks} exact></Route>
                            <Route path="/home/project-management" component={ProjectManagement}></Route>
                            <Route path="/home/proposition-paper" component={PropositionPaper}></Route>
                            <Route path="/402" component={Forbid} exact></Route>
                            <Route component={NotFound}></Route>
                        </Switch>
                    </Content>
                </Layout>
            </DocumentTitle> 
        )
    }
}
