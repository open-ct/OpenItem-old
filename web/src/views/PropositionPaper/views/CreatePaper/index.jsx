import React, { Component } from 'react'
import {PageHeader,Button,Spin,Descriptions,message,Row,Col,Input,Pagination,Form,Tag,Select} from 'antd'
import request from '../../../../utils/request'
import UpLoadQuestionModal from '../../../../components/UpLoadQuestionModal'
import HistoryQuestion from '../../../../components/HistoryQuestion'
import DifficultyBox from '../../../../components/DifficultyBox'
import './index.less'

//const { Option } = Select
const { Search } = Input

export default class index extends Component {

    state={
        initLoading:true,
        createTime:0,
        projectInfo:{},
        upLoadQuestionModalParams:{
            show:false,
            type:"update-paper"
        }
    }

    componentDidMount(){
        let t =new Date()
        this.setState({
            createTime:`${t.getFullYear()}-${t.getMonth().toString().padStart(2,'0')}-${t.getDate().toString().padStart(2,'0')} ${t.getHours().toString().padStart(2,'0')}:${t.getMinutes().toString().padStart(2,'0')}:${t.getSeconds().toString().padStart(2,'0')}`
        })
        this.getProjectInfo()
    }

    getProjectInfo=()=>{
        this.setState({
            initLoading:true
        })
        request({
            url:`/proj/detailed/${this.props.match.params.project}`,
            method:"GET"
        }).then(res=>{
            this.setState({
                projectInfo:res.data.basic_info,
                initLoading:false
            })
        }).catch(err=>{
            this.setState({
                initLoading:false
            })
            this.props.history.push("/home/proposition-paper/home");
            message.error("编辑器加载失败！");
        })
    }

    render() {
        return (
            <div className="create-paper-page" data-component="create-paper-page">
                <PageHeader
                    ghost={false}
                    onBack={() => this.props.history.push("/home/proposition-paper/home")}
                    title="命题组卷"
                    subTitle="上传试题"
                    extra={[
                        <Button key="1" onClick={()=>{
                            this.setState({
                                upLoadQuestionModalParams:{
                                    show:true,
                                    type:"update-paper"
                                }
                            })
                        }}>编辑内容</Button>,
                    ]}
                    >
                    {
                        this.state.initLoading?(
                            <Spin spinning={true} tip="初始化中"></Spin>
                        ):(
                            <Descriptions size="small" column={3}>
                                <Descriptions.Item label="创建时间" key="createAt">{this.state.createTime}</Descriptions.Item>
                                <Descriptions.Item label="项目" key="peojects">{this.state.projectInfo.basic_info.name}</Descriptions.Item>
                                <Descriptions.Item label="学科" key="subjects">{
                                    this.state.projectInfo.basic_info.subjects.map((item,index)=>(
                                        <span>{item}{index===this.state.projectInfo.basic_info.subjects.length-1?"":"、"}</span>
                                    ))
                                }</Descriptions.Item>
                                
                                <Descriptions.Item label="内容纬度" key="content">{
                                    this.props.match.params.content.split(',').map((item,index)=>(
                                        <span>{item}{index===this.props.match.params.content.split(',').length-1?"":"、"}</span>
                                    ))
                                }</Descriptions.Item>
                                <Descriptions.Item label="能力纬度" key="ability">
                                    {
                                        this.props.match.params.ability.split(',').map((item,index)=>(
                                            <span>{item}{index===this.props.match.params.ability.split(',').length-1?"":"、"}</span>
                                        ))
                                    }
                                </Descriptions.Item>
                            </Descriptions>
                        )
                    }
                </PageHeader>
                
                <Row className="container">
                    <Col span="8" className="left-box">
                        <div className="title-box">
                            <div className="title-value">
                                <div className="ver-line"></div>
                                <div className="value">
                                    <span>相关题库</span>
                                </div>
                            </div>
                            <Search placeholder="input search text" style={{width:"2.64rem"}} size="middle" />
                        </div>
                        <HistoryQuestion/>
                        <HistoryQuestion/>
                        <HistoryQuestion/>
                        <Pagination defaultCurrent={1} total={50} className="page-spare"/>
                    </Col>
                    <Col span="16" className="right-box">
                        <div className="right-content">
                            <Row className="top-box">
                                <Col span="18">
                                </Col>
                                <Col span="6" className="question-params">
                                    <div className="title">
                                        <span>参数编辑</span>
                                    </div>
                                    <div className="params-box">
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
                                                    onSelect={(e)=>{
                                                        console.log("传出")
                                                    }}
                                                    size="small"
                                                >
                                                    
                                                </Select>
                                            </Form.Item>
                                            <Form.Item
                                                label="难度"
                                                name="difficult"
                                                colon={false}
                                            >
                                                <DifficultyBox
                                                    onChange={(e)=>{
                                                        
                                                    }}
                                                />
                                            </Form.Item>
                                            <Form.Item
                                                label="能力纬度"
                                                name="ability"
                                                colon={false}
                                            >
                                                <Tag>参数1</Tag>
                                                <Tag>参数2</Tag>
                                            </Form.Item>
                                            <Form.Item
                                                label="内容纬度"
                                                name="content"
                                                colon={false}
                                            >
                                                <Tag>参数1</Tag>
                                                <Tag>参数2</Tag>
                                            </Form.Item>
                                        </Form>
                                    </div>
                                </Col>
                                <div className="btn-box">
                                    <Button type="primary" size="small" style={{marginRight:".12rem",width:".73rem"}}>编辑</Button>
                                    <Button type="primary" danger size="small" style={{width:".73rem"}}>删除</Button>
                                </div>
                            </Row>
                            <Row className="btn-line">
                                <Col span="17">
                                    <Button type="primary">添加一题</Button>
                                </Col>
                                <Col span="7">
                                    <Button type="info" style={{marginRight:".05rem"}}>预览试卷</Button>
                                    <Button type="primary">保存</Button>
                                </Col>
                            </Row>
                        </div>
                    </Col>
                   
                </Row>

                <UpLoadQuestionModal
                    {...this.state.upLoadQuestionModalParams}
                    onClose={()=>{
                        this.setState({
                            upLoadQuestionModalParams:{
                                show:false,
                                type:"update-paper"
                            }
                        })
                    }}
                />
            </div>
        )
    }
}
