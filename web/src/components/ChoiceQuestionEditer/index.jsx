import React, { Component } from 'react'
import { Divider, Radio, Space, Row, Col,Form,Select,Tag,Input,Checkbox, message,Button } from "antd"
import { PictureFilled,PlusCircleOutlined,MinusCircleOutlined,UpCircleOutlined,DownCircleOutlined } from '@ant-design/icons';
import DifficultyBox from '../DifficultyBox'
import './index.less'

const { Option } = Select
const { TextArea } = Input


export default class index extends Component {

    state = {
        questionTitle:"这里是题目",
        questionOption:[{
            value:"测试",
            picture:"",
            isCorrect:false
        }],
        questionParams:{
            subject:"",
            difficulty:1,
        }
    }

    render() {
        return (
            <div className="choice-question-editer" data-component="choice-question-editer">
                <div className="question-title">
                    <span>{this.state.questionTitle}</span>
                </div>
                <Divider />
                <Radio.Group>
                    <Space direction="vertical">
                        {
                            this.state.questionOption.map((item,index)=>(
                                <Radio value={item.value} key={`${index}-${Math.random(100)}`}>{item.value}</Radio>
                            ))
                        }
                    </Space>
                </Radio.Group>
                <Row style={{marginTop:".3rem"}}  gutter={[8, 8]}>
                    <Col span={16}>
                        <div className="question-title-editer">
                            <div className="title">
                                <span className="title-value">题目编辑</span>
                                <PictureFilled className="icon"/>
                            </div>
                            <div className="title-editer-box">
                                <TextArea placeholder="题目输入区域" autoSize className="title-input" value={this.state.questionTitle} defaultValue={this.state.questionTitle} onChange={(e)=>{
                                    this.setState({
                                        questionTitle:e.target.value
                                    })
                                }}/>
                            </div>
                        </div>
                    </Col>
                    <Col span={8}>
                        <div className="question-params">
                            <div className="title">
                                <span>参数编辑</span>
                            </div>
                            <div className="params-list-box">
                                <Form
                                    labelCol={{ span: 7 }}
                                    wrapperCol={{ span: 17 }}
                                    labelAlign="left"
                                >
                                    <Form.Item
                                        label="学科"
                                        name="subjects"
                                        colon={false}
                                    >
                                        <Select 
                                            placeholder="选择学科" 
                                            value={this.state.questionParams.subject} 
                                            defaultValue={this.props.defaultSubjectValue} 
                                            onSelect={(e)=>{
                                                console.log("传出")
                                            }}
                                            size="small"
                                        >
                                            {
                                                this.props.subjectList.map((item,index)=>(
                                                    <Option value={item} key={index+Math.random(100)}>{item}</Option>
                                                ))
                                            }
                                        </Select>
                                    </Form.Item>
                                    <Form.Item
                                        label="难度"
                                        name="difficult"
                                        colon={false}
                                    >
                                        <DifficultyBox
                                            onChange={(e)=>{
                                                let questionParams = Object.assign(this.state.questionParams,{difficulty:e})
                                                this.setState({
                                                    questionParams
                                                })
                                            }}
                                        />
                                    </Form.Item>
                                    <Form.Item
                                        label="能力纬度"
                                        name="ability"
                                        colon={false}
                                    >
                                        {
                                            this.props.ability.map(item=>(
                                                <Tag>{item}</Tag>
                                            ))
                                        }
                                    </Form.Item>
                                    <Form.Item
                                        label="内容纬度"
                                        name="content"
                                        colon={false}
                                    >
                                        {
                                            this.props.content.map(item=>(
                                                <Tag>{item}</Tag>
                                            ))
                                        }
                                    </Form.Item>
                                </Form>
                            </div>
                        </div>
                    </Col>
                </Row>
                <div className="option-list-box">
                    <div className="title">
                        <Row className="title-line">
                            <Col span="14" className="title-item">
                                <span>选项文字</span>
                            </Col>
                            <Col span="4" className="title-item">
                                <span>图片</span>
                            </Col>
                            <Col span="3" className="title-item">
                                <span>正确答案</span>
                            </Col>
                            <Col span="3" className="title-item">
                                <span>上移下移</span>
                            </Col>
                        </Row>
                    </div>
                    {
                        this.state.questionOption.map((item,index)=>(
                            <Row className="option-oper-item" key={Math.random(100000)}>
                                <Col span="14" className="item-value">
                                    <Row gutter={[16, 16]} className="value-box">
                                        <Col span="19">
                                            <Input placeholder="请输入选项内容" value={item.value} onChange={(e)=>{
                                                let questionOption = [...this.state.questionOption]
                                                questionOption[index].value=e.target.value
                                                this.setState({
                                                    questionOption
                                                })
                                            }}/>
                                        </Col>
                                        <Col span="5" className="value-oper">
                                            <PlusCircleOutlined className="btn" onClick={()=>{
                                                let questionOption = [...this.state.questionOption]
                                                questionOption.push({
                                                    value:"",
                                                    picture:"",
                                                    isCorrect:false
                                                })
                                                this.setState({
                                                    questionOption
                                                })
                                            }}/>
                                            <MinusCircleOutlined className="btn" onClick={()=>{
                                                if(this.state.questionOption.length===1){
                                                    message.warning("至少保留一个选项！")
                                                }else{
                                                    let questionOption = [...this.state.questionOption]
                                                    questionOption.splice(index,1)
                                                    this.setState({
                                                        questionOption
                                                    })
                                                }
                                            }}/>
                                        </Col>
                                    </Row>
                                </Col>
                                <Col span="4" className="item-picture">
                                    <PictureFilled className="btn"/>
                                </Col>
                                <Col span="3" className="item-isCorrect">
                                    <Checkbox checked={item.isCorrect} onChange={(e)=>{
                                        let questionOption = [...this.state.questionOption]
                                        questionOption[index].isCorrect=e.target.checked
                                        this.setState({
                                            questionOption
                                        })
                                    }}/>
                                </Col>
                                <Col span="3" className="item-move">
                                    <UpCircleOutlined className="btn" onClick={()=>{
                                        if(index===0){
                                            message.error("当前为第一个选项，不可上移")
                                        }else{
                                            let questionOption = [...this.state.questionOption]
                                            let s = questionOption[index]
                                            questionOption[index] = questionOption[index-1]
                                            questionOption[index-1] = s
                                            this.setState({
                                                questionOption
                                            })
                                        }
                                    }}/>
                                    <DownCircleOutlined className="btn" onClick={()=>{
                                        if(index===this.state.questionOption.length-1){
                                            message.warning("当前为最后一个元素，不可下移")
                                        }else{
                                            let questionOption = [...this.state.questionOption]
                                            let s = questionOption[index]
                                            questionOption[index] = questionOption[index+1]
                                            questionOption[index+1] = s
                                            this.setState({
                                                questionOption
                                            })
                                        }
                                    }}/>
                                </Col>
                            </Row>
                        ))
                    }
                </div>
                <div className="question-complete-box">
                    <Button type="primary" block>完成编辑</Button>
                </div>
            </div>
        )
    }
}
