import React, { Component } from 'react'
import { Layout, Menu } from 'antd';
import DocumentTitle from 'react-document-title'
import {Switch,Route} from 'react-router-dom'
import { AppstoreOutlined, MailOutlined, SettingOutlined } from '@ant-design/icons';
import PendingTasks from '../PendingTasks'
import './index.less'
const { Header , Sider, Content } = Layout;
const {SubMenu} = Menu

export default class index extends Component {

    funcChange = (e)=>{
        this.props.history.push(`/home/${e.key}`)
    }

    render() {
        return (
            <DocumentTitle title="命题与题库系统">
                <Layout className="home-page" data-component="home-page">
                    <Header>
                        <div className="header-box">
                            <span className="header-title">命题与题库系统</span>
                        </div>
                    </Header>
                    <Layout className="container">
                        <Sider>
                            <Menu
                                style={{ width: 200,height:'100%' }}
                                defaultSelectedKeys={['pending-tasks']}
                                defaultOpenKeys={['sub1']}
                                mode="inline"
                                onClick = {this.funcChange}
                            >
                                <SubMenu key="sub1" title="项目管理" icon={<MailOutlined />}>
                                    <Menu.Item key="pending-tasks">Option 1</Menu.Item>
                                    <Menu.Item key="2">Option 2</Menu.Item>
                                    <Menu.Item key="3">Option 3</Menu.Item>
                                    <Menu.Item key="4">Option 4</Menu.Item>
                                </SubMenu>
                                <SubMenu key="sub2"  title="人员管理" icon={<AppstoreOutlined />}>
                                    <Menu.Item key="5">Option 5</Menu.Item>
                                    <Menu.Item key="6">Option 6</Menu.Item>
                                </SubMenu>
                                <SubMenu key="sub4" title="系统管理" icon={<SettingOutlined />}>
                                    <Menu.Item key="9">Option 9</Menu.Item>
                                    <Menu.Item key="10">Option 10</Menu.Item>
                                    <Menu.Item key="11">Option 11</Menu.Item>
                                    <Menu.Item key="12">Option 12</Menu.Item>
                                </SubMenu>
                            </Menu>
                        </Sider>
                        <Content>
                            <Switch>
                                <Route path="/home" component={PendingTasks} exact />
                                <Route path="/home/pending-tasks" component={PendingTasks}></Route>
                            </Switch>
                        </Content>
                    </Layout>
                </Layout>
            </DocumentTitle> 
        )
    }
}
