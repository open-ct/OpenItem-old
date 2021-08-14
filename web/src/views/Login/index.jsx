import React, { Component } from 'react'
import { Button, Form, Input, Spin } from 'antd'
import DocumentTitle from 'react-document-title'
import { InfoCircleOutlined } from '@ant-design/icons';
import request from '../../utils/request'
import { setuserinfo } from '../../store/action'
import store from '../../store'
import './index.less'

export default class index extends Component {
    state = {
        mode:"login",
        errorState:{
            show:true,
            title:'登陆错误'
        },
        loadingState:{
            show:false,
            title:'登陆中'
        },
        registerForm:{
            username:"",
            password:""
        }
    }
    
    formRef = React.createRef()

    login = ()=>{
        const form = this.formRef.current
        form.validateFields().then(data=>{
            this.setState({
                loadingState:{
                    show:false
                }
            })
            this.props.history.push('/home')
            request({ method:'POST', url:'/user/login', data }).then(res=>{
                store.dispatch(setuserinfo(res))
                this.setState({
                    loadingState:{
                        show:false
                    }
                })
                this.props.history.push('/home')
            }).catch(err=>{
                this.setState({
                    errorState:{
                        title:err.message,
                        show:true
                    },
                    loadingState:{
                        show:false
                    }
                })
            })
        }).catch(err=>{
            this.setState({
                errorState:{
                    title:"请正确输入账号和密码",
                    show:true
                },
                loadingState:{
                    show:false,
                    title:''
                }
            })
        })
    }


    register = ()=>{
        const form = this.formRef.current
        form.validateFields().then(data=>{
            request({ method:'POST', url:'/user/register', data }).then(res=>{
                console.log(res)
            }).catch(err=>{
                console.log(err)
                this.setState({
                    errorState:{
                        title:err.message,
                        show:true
                    },
                    loadingState:{
                        show:false
                    }
                })
            })
        }).catch(err=>{
            this.setState({
                errorState:{
                    title:"请正确输入账号和密码",
                    show:true
                },
                loadingState:{
                    show:false,
                    title:''
                }
            })
        })
    }

    render() {
        return (
            <DocumentTitle title="登陆/注册">
                <Spin tip={this.state.loadingState.title}  size="large" spinning={this.state.loadingState.show}>
                    <div className="login-page" data-component="login-page">
                        <div className="login-box">
                            <div className="login-btn__change login-btn__box">
                                <Button type="primary" shape="round" size="large" onClick={()=>{
                                    this.setState({
                                        mode:this.state.mode==="login"?"register":"login"
                                    })
                                }}>{this.state.mode==='login'?'注册':'登陆'}</Button>
                            </div>
                            <div className="state-box">
                                {this.state.errorState.show?(
                                    <div className="message">
                                        <InfoCircleOutlined />
                                        <span className="content">{this.state.errorState.title}</span>
                                    </div>
                                ):""}
                            </div>
                            <div className="login-form">
                                <Form 
                                    name="form"
                                    labelCol={{ span: 4 }}
                                    wrapperCol={{ span: 20 }}
                                    ref = {this.formRef}
                                    initialValues = {{email:"robot-0@pqbs.com",password:"123456"}}
                                >
                                    <Form.Item
                                        label="用户名"
                                        name="email"
                                        colon={false}
                                        rules={[{ required: true, message: '请输入您的邮箱！' },{ min: 6, max: 20, message: '邮箱长度在 6 到 20 个字符', trigger: 'blur' }]}
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
                                <Button type="primary" shape="round" size="large" onClick={()=>{
                                    this.setState({
                                        errorState:{show:false,title:''},
                                        loadingState:{title:this.state.mode==="login"?'登陆中':'注册中',show:true}
                                    })
                                    this.state.mode==="login"?this.login():this.register()
                                }}>{this.state.mode==='login'?'登陆':'注册'}</Button>
                            </div>
                        </div>
                        <div className="login-describe">
                            <h1>命题与题库系统</h1>
                            <p className="describe-item">系统介绍xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</p>
                            <p className="describe-item">系统介绍xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</p>
                            <p className="describe-item">系统介绍xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</p>
                        </div>
                    </div>
                </Spin>
            </DocumentTitle>
        )
    }
}
