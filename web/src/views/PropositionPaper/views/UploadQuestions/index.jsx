import React, { Component } from 'react'
import {PageHeader,Descriptions,Button,Layout,Menu,Input, message,Spin,Pagination} from 'antd'
import ChoiceQuestionEditer from '../../../../components/ChoiceQuestionEditer'
import HistoryQuestion from '../../../../components/HistoryQuestion'
import UpLoadQuestionModal from '../../../../components/UpLoadQuestionModal'
import request from '../../../../utils/request'

import './index.less'

const { Search } = Input;
const { Sider, Content } = Layout;

export default class index extends Component {

    state = {
        difficultyValue:4,
        createTime:0,
        questionTitle:"",
        questionOption:[{
            value:"",
            picture:"",
            isCorrect:false,
        }],
        projectInfo:{},
        initLoading:true,

        upLoadQuestionModalParams:{
            show:false,
            type:"update"
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
            <div className="upLoad-question-page" data-component="upLoad-question-page">
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
                                    type:"update"
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
                <div className="main">
                    <Layout className="container">
                        <Sider theme="light" width="2.4rem" style={{backgroundColor:"#FAFAFA"}}>
                            <Menu
                                style={{ width: "2.4rem" }}
                                defaultSelectedKeys={['1']}
                                defaultOpenKeys={['1']}
                                mode="vertical"
                                theme="light"
                            >
                                <Menu.Item key="1">选择题</Menu.Item>
                                <Menu.Item key="2">填空题</Menu.Item>
                                <Menu.Item key="3">阅读题</Menu.Item>
                            </Menu>
                        </Sider>
                        <Content style={{backgroundColor:"white"}} className="content">
                            <ChoiceQuestionEditer
                                defaultSubjectValue={this.props.match.params.subject}
                                subjectList={this.state.initLoading?[]:this.state.projectInfo.basic_info.subjects}
                                ability={this.props.match.params.ability.split(',')}
                                content={this.props.match.params.content.split(',')}
                            />
                        </Content>
                        <Sider theme="light" width="7.47rem" className="question-box">
                            <div className="question-content-box">
                                <div className="title">相关题目</div>
                                <div className="filter-box">
                                    <span>筛选</span>
                                    <Search placeholder="input search text" style={{ width: 200 }} size="small"/>
                                </div>
                                <HistoryQuestion

                                />
                                <HistoryQuestion
                                    
                                />
                                <HistoryQuestion
                                    
                                />
                                <Pagination defaultCurrent={1} total={50} className="page-spare"/>
                            </div>
                        </Sider>
                    </Layout>
                </div>
                <UpLoadQuestionModal
                    {...this.state.upLoadQuestionModalParams}
                    onClose={()=>{
                        let upLoadQuestionModalParams = Object.assign(this.state.upLoadQuestionModalParams,{show:false})
                        this.setState({
                            upLoadQuestionModalParams
                        })
                    }}
                />
            </div>
        )
    }
}
