import React, { Component } from 'react'
import { Layout, Menu, Avatar, Button } from 'antd';
import DocumentTitle from 'react-document-title'
import {Switch, Route,Redirect} from 'react-router-dom'
import {  UserOutlined } from '@ant-design/icons';
import PendingTasks from '../PendingTasks'
import ProjectManagement from '../ProjectManagement'
import './index.less'
const { Header , Content } = Layout;

export default class index extends Component {

    funcChange = (e)=>{
        this.props.history.push(`/home/${e.key}`)
    }

    render() {
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
                            <Menu onClick = {this.funcChange} defaultSelectedKeys={['pending-tasks']} mode="horizontal">
                                <Menu.Item key="pending-tasks">
                                    项目管理
                                </Menu.Item>
                                <Menu.Item key="2">
                                    人员管理
                                </Menu.Item>
                                <Menu.Item key="3">
                                    系统管理
                                </Menu.Item>
                            </Menu>
                            </div>
                            <div className="tab">
                                <span>Tab</span>
                            </div>
                            <div className="button-group">
                                <Button ghost size="small">操作1</Button>
                                <Button ghost size="small">操作2</Button>
                            </div>
                            <div className="user">
                                <Avatar icon={<UserOutlined />} size={24}/>
                                <span>YingMing</span>
                            </div>
                        </div>
                    </Header>
                    <Content>
                        <Switch>
                            <Redirect from="/home" to="/home/pending-tasks" exact></Redirect>
                            <Route path="/home/pending-tasks" component={PendingTasks} exact></Route>
                            <Route path="/home/project-management" component={ProjectManagement}></Route>
                        </Switch>
                    </Content>
                </Layout>
            </DocumentTitle> 
        )
    }
}
