import React, { Component } from 'react'
import { Switch, Route, Redirect } from 'react-router-dom'
import DocumentTitle from 'react-document-title'
import { PageHeader, Tabs, Button, Descriptions,Spin,message } from 'antd';
import Step from '../../components/Step'
import BuildTeam from '../../components/BuildTeam'
import NotFound from '../errors/404'
import request from '../../utils/request'

import './index.less'

const { TabPane } = Tabs;
  
export default class index extends Component {
    
    state = {
        projectBaseInfo:{},
        loadingState:true
    }

    componentDidMount(){
        this.getProjectBaseInfo()
    }

    getProjectBaseInfo(){
        this.setState({
            loadingState:true
        })
        request({ method:'GET', url:`http://49.232.73.36:8081/review/proj/detailed/${this.props.match.params.project_id}`}).then(res=>{
            this.setState({
                projectBaseInfo:res.data,
                loadingState:false
            })
        }).catch(err=>{
            this.props.history.push('/home')
            message.error(err.message||"项目信息加载失败，请重试！");
            this.setState({
                loadingState:false
            })
        })
    }

    dateFilter(time){
        let date = new Date(time)
        return `${date.getFullYear()}-${date.getMonth().toString().padStart(2,'0')}-${date.getDate().toString().padStart(2,'0')}`
    }

    tabCruuent=()=>{
        let path_list = this.props.location.pathname.split('/')
        return `${path_list[path_list.length-1]}_${path_list[path_list.length-2]}`
    }

    render() {
        return (
            <DocumentTitle title="项目管理">
                <div className="project-management-page" data-component="project-management-page" key="project-management-page">
                    <PageHeader
                        ghost={false}
                        onBack={() => this.props.history.push("/home")}
                        title="项目管理"
                        subTitle={this.state.loadingState?"加载中":this.state.projectBaseInfo.basic_info.basic_info.name}
                        extra={[
                            <Button key="2">编辑项目</Button>,
                            <Button key="1">导出成员</Button>,
                        ]}
                        footer={
                            this.state.loadingState?(
                                <Spin spinning={this.state.loadingState} tip="加载中"/>
                            ):(
                                <Tabs defaultActiveKey={`${this.state.projectBaseInfo.steps[0].uuid}_${this.state.projectBaseInfo.steps[0].name}`} type="card" activeKey={this.tabCruuent()} onChange={(e)=>{
                                    this.props.history.push(`/home/project-management/${this.props.match.params.project_id}/${this.props.match.params.role}/${e.split('_')[1]}/${e.split('_')[0]}`)
                                }}>
                                    {
                                        this.state.projectBaseInfo.steps.map(item=>(
                                            <TabPane key={`${item.uuid}_${item.name}`} tab={item.name}></TabPane>
                                        ))
                                    }
                                </Tabs>
                            )
                        }
                        >
                            {
                                this.state.loadingState?(
                                    <Spin spinning={this.state.loadingState} tip="加载中..." />
                                ):(
                                    <Descriptions size="small" column={3} style={{width:'auto'}}>
                                        <Descriptions.Item label="创建时间">{this.dateFilter(this.state.projectBaseInfo.basic_info.CreateAt)}</Descriptions.Item>
                                        <Descriptions.Item label="学科">
                                            {
                                                this.state.projectBaseInfo.basic_info.basic_info.subjects.map((item,index)=>(
                                                    <span>{`${item}${index===this.state.projectBaseInfo.basic_info.basic_info.subjects.length-1?'':'、'}`}</span>
                                                ))
                                            }
                                        </Descriptions.Item>
                                        <Descriptions.Item label="学段">
                                            {
                                                this.state.projectBaseInfo.basic_info.basic_info.grade_range.map((item,index)=>(
                                                    <span>{`${item}${index===this.state.projectBaseInfo.basic_info.basic_info.grade_range.length-1?'':'、'}`}</span>
                                                ))
                                            }
                                        </Descriptions.Item>
                                        <Descriptions.Item label="试卷">0</Descriptions.Item>
                                        <Descriptions.Item label="试题">0</Descriptions.Item>
                                    </Descriptions>
                                )
                            }
                    </PageHeader>
                    <div className="page-content-box">
                        {
                            this.state.loadingState?(<></>):(
                                <Switch>
                                    <Redirect from={`/home/project-management/${this.props.match.params.project_id}/${this.props.match.params.role}`} to={`/home/project-management/${this.props.match.params.project_id}/${this.props.match.params.role}/${this.state.projectBaseInfo.steps[0].name}/${this.state.projectBaseInfo.steps[0].uuid}`} exact></Redirect>
                                    {
                                        this.state.projectBaseInfo.steps.map(item=>(
                                            <Route path={`/home/project-management/:project_id/:role/${item.name}/:step_id`} component={item.name==="组建团队"?BuildTeam:Step} exact key={item.Id}></Route>
                                        ))
                                    }
                                    <Route component={NotFound} key="404"></Route>
                                </Switch>
                            )
                        }
                    </div>
                </div>
            </DocumentTitle>
        )
    }
}
