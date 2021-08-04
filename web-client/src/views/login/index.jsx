import React, { Component } from 'react'
import {Button,Form,Input,message} from 'antd'
import DocumentTitle from 'react-document-title'
import './index.less'

export default class index extends Component {
    state = {
        mode:"login",
        registerForm:{
            username:"",
            password:""
        }
    }
    
    loginFormRef = React.createRef()

    login = ()=>{
        const form = this.loginFormRef.current
        form.validateFields().then(val=>{
            console.log(val)
            this.props.history.push('/home')
        }).catch(err=>{
            message.warning('请正确输入账号和密码');
        })
    }

    modeChange = (mode)=>{
        this.setState({
            mode
        })
    }
    modePanelLoad(){
        if(this.state.mode==="login"){
            return (
                <div className="login-box">
                    <div className="login-btn__change login-btn__box">
                        <Button type="primary" shape="round" size="large" onClick={this.modeChange.bind(this,'register')}>注册</Button>
                    </div>
                    <div className="login-form">
                        <Form 
                            name="loginForm"
                            labelCol={{ span: 7 }}
                            wrapperCol={{ span: 17 }}
                            ref = {this.loginFormRef}
                            initialValues = {{username:"1811020024",password:"123456"}}
                        >
                            <Form.Item
                                label="用户名"
                                name="username"
                                colon={false}
                                rules={[{ required: true, message: '请输入您的账号！' },{ min: 6, max: 10, message: '账号长度在 6 到 10 个字符', trigger: 'blur' }]}
                            >
                                <Input/>
                            </Form.Item>
                            <Form.Item
                                label="密码"
                                name="password"
                                colon={false}
                                rules={[{ required: true, message: '请输入您的密码！' },{ min: 6, max: 10, message: '密码长度在 6 到 10 个字符', trigger: 'blur' }]}
                            >
                                <Input.Password />
                            </Form.Item>
                        </Form>
                        <div className="more-choice-box">
                            <a href="./forget">忘记密码</a>
                        </div>
                    </div>
                    <div className="login-btn__operation login-btn__box">
                        <div className="space-box"></div>
                        <Button type="primary" shape="round" size="large" onClick={this.login}>登陆</Button>
                    </div>
                </div>
            )
        }
        return (
            <div className="login-box">
                <div className="login-btn__change login-btn__box">
                    <Button type="primary" shape="round" size="large" onClick={this.modeChange.bind(this,'login')}>登陆</Button>
                </div>
                <div className="login-form">
                    <Form 
                        name="registerForm"
                        labelCol={{ span: 7 }}
                        wrapperCol={{ span: 17 }}
                        initialValues={{ remember: true }}
                    >
                        <Form.Item
                            label="用户名"
                            name="username"
                            colon={false}
                            rules={[{ required: true, message: 'Please input your username!' }]}
                        >
                            <Input />
                        </Form.Item>
                        <Form.Item
                            label="密码"
                            name="password"
                            colon={false}
                            rules={[{ required: true, message: 'Please input your username!' }]}
                        >
                            <Input.Password />
                        </Form.Item>
                    </Form>
                    <div className="more-choice-box">
                        <a href="./forget">忘记密码</a>
                    </div>
                </div>
                <div className="login-btn__operation login-btn__box">
                    <div className="space-box"></div>
                    <Button type="primary" shape="round" size="large">注册</Button>
                </div>
            </div>
        )
    }
    render() {
        return (
            <DocumentTitle title="登陆/注册">
                <div className="login-page" data-component="login-page">
                    {this.modePanelLoad()}
                    <div className="login-describe">
                        <h1>命题与题库系统</h1>
                        <p className="describe-item">系统介绍xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</p>
                        <p className="describe-item">系统介绍xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</p>
                        <p className="describe-item">系统介绍xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</p>
                    </div>
                </div>
            </DocumentTitle>
        )
    }
}
