import React, { Component } from 'react'
import { Button, Radio, Modal, Table, Pagination, Space, Form, Select } from 'antd'

import PropositionPaperIcon from '../../../../asset/images/proposition-paper-icon.png'
import { LockOutlined , EditOutlined, EllipsisOutlined } from '@ant-design/icons';
import './index.less'

const {Option} = Select

export default class index extends Component {

    state = {
        mode: 'questions',
        questionData:[],
        recordData:[{
            key:'1',
            date:'text',
            content:'/',
            param1:'/',
            param2:'/',
            param3:'/'

        },{
            key:'2',
            date:'text',
            content:'内容修改版本的文字内容......',
            param1:'内容纬度1；内容纬度2',
            param2:'/',
            param3:'/'

        },{
            key:'3',
            date:'text',
            content:'/',
            param1:'/',
            param2:'/',
            param3:'/'

        },{
            key:'4',
            date:'text',
            content:'/',
            param1:'/',
            param2:'/',
            param3:'/'

        },{
            key:'5',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'6',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'7',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'8',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'9',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'10',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'11',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'12',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'13',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'14',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        }],
        modifyRecordVisible:false,
        upLoadQuestionVisible:false,
        editQuestionVisible:false,
    }

    columns = [{
        title: '时间',
        dataIndex: 'date',
        key: 'date',
        width:110
    },{
        title: '内容',
        dataIndex: 'content',
        key: 'content',
        width:110
    },{
        title: '参数1',
        dataIndex: 'param1',
        key: 'param1',
        width:110
    },{
        title: '参数2',
        dataIndex: 'param2',
        key: 'param2',
        width:110
    },{
        title: '参数3',
        dataIndex: 'param3',
        key: 'param3',
        width:110
    },{
        title: '操作',
        key: 'operate',
        render: (text, record) => (
            <Button type="link">预览</Button>
        ),
    }];

    handleModeChange = e => {
        const mode = e.target.value;
        this.setState({ mode });
    };

    loadQuestionData = ()=>{
        if(this.state.mode==="questions"){
            return this.state.questionData.map((item,index)=>(
                <div className="question-item" key={index}>
                    <div className="header">
                        <span className="category">{item.category}</span>
                        <span className="date">{item.date}</span>
                    </div>
                    <div className="container">
                        <div className="info">
                            <span>{item.info}</span>
                        </div>
                        <div className="project">
                            <span>{item.projectName}</span>
                        </div>
                    </div>
                    <div className="footer">
                        <div className="action-item b-right">
                            <LockOutlined/>
                        </div>
                        <div className="action-item b-right" onClick={()=>{
                            this.setState({
                                editQuestionVisible:true
                            })
                        }}>
                            <EditOutlined/>
                        </div>
                        <div className="action-item" onClick={()=>{
                            this.setState({
                                modifyRecordVisible:true
                            })
                        }}>
                            <EllipsisOutlined/>
                        </div>
                    </div>
                </div>
            ))
        }
        return ""
    }

    componentDidMount = ()=>{
        let dataList = []
        for(let i=0;i<12;i++){
            dataList.push({
                category:"数学",
                info:"题目信息...........",
                projectName:"所属项目名称abbbbb",
                date:"2020/1/20"
            })
        }
        this.setState({
            questionData:dataList
        })
    }

    render() {
        return (
            <div className="proposition-paper-home-page" data-component="proposition-paper-home-page">
                <div className="header">
                    <div className="describe">
                        <div className="content">
                            <h1>实现高效命题组卷</h1>
                            <div className="context">
                                单题编辑上传，试卷调用，轻松快捷<br/>
                                xxxxxxxxxx示例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言
                            </div>
                        </div>
                        <div className="btn-box">
                            <Button type="primary" onClick={()=>{
                                this.setState({
                                    upLoadQuestionVisible:true
                                })
                            }}>上传试题</Button>
                            <Button type="primary" style={{marginLeft:'.18rem'}}>创建试卷</Button>
                        </div>
                    </div>
                    <img src={PropositionPaperIcon} alt="图片" className="icon"></img>
                </div>
                <div className="container">
                    <div className="category-list">
                        <Radio.Group onChange={this.handleModeChange} value={this.state.mode} >
                            <Radio.Button value="questions">我的试题</Radio.Button>
                            <Radio.Button value="warehouse" style={{ marginLeft: 4 }}>我的试卷库</Radio.Button>
                        </Radio.Group>
                    </div>
                    <div className="main">
                        {this.loadQuestionData()}
                    </div>
                </div>
                <Modal title="修改记录" visible={this.state.modifyRecordVisible} footer={null} onCancel={()=>{
                    this.setState({
                        modifyRecordVisible:false
                    })
                }}  width="16.12rem">
                    <div className="question-modify-record-box">
                        <div className="preview-box">
                            <div className="header">
                                <span>题目效果预览窗</span>
                            </div>
                        </div>
                        <div className="record-box">
                            <Space align="center">
                                修改历史版本
                            </Space>
                            <div className="container">
                                <Table 
                                    dataSource={this.state.recordData} 
                                    columns={this.columns} 
                                    pagination={false}
                                    scroll = {{y:'calc(100% - .56rem)'}}
                                />
                            </div>
                            <div className="footer">
                                <Pagination
                                    total={85}
                                    showTotal={total => `Total ${total} items`}
                                    defaultPageSize={20}
                                    defaultCurrent={1}
                                    size="small"
                                />
                            </div>
                        </div>
                    </div>
                </Modal>
                <Modal 
                    title="上传试题" 
                    cancelText="取消创建"
                    okText="下一步"
                    visible={this.state.upLoadQuestionVisible} 
                    onOk={()=>{
                        this.props.history.push("/home/proposition-paper/upload-questions")
                    }} 
                    onCancel={()=>{
                        this.setState({
                            upLoadQuestionVisible:false
                        })
                    }}
                >
                    <div className="upLoad-question-title">
                        <span>试题编号编号:absnahghj（自动生成）</span>
                    </div>
                    <Form 
                        labelCol={{ span: 5 }}
                        wrapperCol={{ span: 19 }}
                        labelAlign="left"
                    >
                        <Form.Item
                            name="project"
                            label="项目"
                            rules={[{ required: true, message: '请选择项目名称' }]}
                        >
                            <Select placeholder="选择项目名称">
                                <Option value="male">项目一</Option>
                                <Option value="female">项目二</Option>
                                <Option value="other">项目三</Option>
                            </Select>
                        </Form.Item> 
                        <Form.Item
                            name="subject"
                            label="学科"
                            rules={[{ required: true, message: '请选择学科' }]}
                        >
                            <Select placeholder="选择学科">
                                <Option value="male">学科一</Option>
                                <Option value="female">学科二</Option>
                                <Option value="other">学科三</Option>
                            </Select>
                        </Form.Item>
                        <Form.Item
                            name="content"
                            label="内容纬度"
                            rules={[{ required: true, message: '请选择内容纬度' }]}
                        >
                            <Select placeholder="选择内容纬度">
                                <Option value="male">内容纬度一</Option>
                                <Option value="female">内容纬度二</Option>
                                <Option value="other">内容纬度三</Option>
                            </Select>
                        </Form.Item>
                        <Form.Item
                            name="ability"
                            label="能力维度"
                            rules={[{ required: true, message: '请选择能力维度' }]}
                        >
                            <Select placeholder="选择能力维度">
                                <Option value="male">能力维度一</Option>
                                <Option value="female">能力维度二</Option>
                                <Option value="other">能力维度三</Option>
                            </Select>
                        </Form.Item>
                        <Form.Item
                            name="type"
                            label="题型"
                            rules={[{ required: true, message: '请选择题型' }]}
                        >
                            <Select placeholder="选择题型">
                                <Option value="male">题型一</Option>
                                <Option value="female">题型二</Option>
                                <Option value="other">题型三</Option>
                            </Select>
                        </Form.Item>
                    </Form>
                </Modal>
                <Modal 
                    title="编辑试题信息" 
                    cancelText="取消"
                    okText="保存"
                    visible={this.state.editQuestionVisible} 
                    onOk={()=>{
                        
                    }} 
                    onCancel={()=>{
                        this.setState({
                            editQuestionVisible:false
                        })
                    }}
                >
                    <div className="upLoad-question-title">
                        <span>试题编号编号:absnahghj（自动生成）</span>
                    </div>
                    <Form 
                        labelCol={{ span: 5 }}
                        wrapperCol={{ span: 19 }}
                        labelAlign="left"
                    >
                        <Form.Item
                            name="project"
                            label="项目"
                            rules={[{ required: true, message: '请选择项目名称' }]}
                        >
                            <Select placeholder="选择项目名称">
                                <Option value="male">项目一</Option>
                                <Option value="female">项目二</Option>
                                <Option value="other">项目三</Option>
                            </Select>
                        </Form.Item> 
                        <Form.Item
                            name="subject"
                            label="学科"
                            rules={[{ required: true, message: '请选择学科' }]}
                        >
                            <Select placeholder="选择学科">
                                <Option value="male">学科一</Option>
                                <Option value="female">学科二</Option>
                                <Option value="other">学科三</Option>
                            </Select>
                        </Form.Item>
                        <Form.Item
                            name="content"
                            label="内容纬度"
                            rules={[{ required: true, message: '请选择内容纬度' }]}
                        >
                            <Select placeholder="选择内容纬度">
                                <Option value="male">内容纬度一</Option>
                                <Option value="female">内容纬度二</Option>
                                <Option value="other">内容纬度三</Option>
                            </Select>
                        </Form.Item>
                        <Form.Item
                            name="ability"
                            label="能力维度"
                            rules={[{ required: true, message: '请选择能力维度' }]}
                        >
                            <Select placeholder="选择能力维度">
                                <Option value="male">能力维度一</Option>
                                <Option value="female">能力维度二</Option>
                                <Option value="other">能力维度三</Option>
                            </Select>
                        </Form.Item>
                        <Form.Item
                            name="type"
                            label="题型"
                            rules={[{ required: true, message: '请选择题型' }]}
                        >
                            <Select placeholder="选择题型">
                                <Option value="male">题型一</Option>
                                <Option value="female">题型二</Option>
                                <Option value="other">题型三</Option>
                            </Select>
                        </Form.Item>
                    </Form>
                </Modal>
            </div>
        )
    }
}
