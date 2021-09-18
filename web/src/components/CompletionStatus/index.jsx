//抽离出完成情况组件

import React, { Component } from 'react'
import { Col, Statistic,Row} from 'antd'
import ModulaCard from '../ModulaCard'
import request from '../../utils/request'
import './index.less'


export default class index extends Component {

    state = {
        status:{},
        loadingState:false
    }

    getStatusList = ()=>{
        this.setState({
            loadingState:true
        })
        request({
            url:`/proj/step/stat/${this.props.stepId}`,
            method:"GET"
        }).then(res=>{
            this.setState({
                status:res.data,
                loadingState:false
            })
        }).catch(err=>{
            this.setState({
                loadingState:false
            })
        })
    }

    componentDidMount(){
        this.getStatusList()
    }

    render() {
        return (
            <ModulaCard title="完成情况">
                {
                    this.state.loadingState?(
                        <></>
                    ):(
                        <div className="completion-status-box" data-component="completion-status-box">
                            <div className="left-box">
                                <div className="title">
                                    <span>{this.props.title}</span>
                                </div>
                                <div className="value-list">
                                    <Row gutter={16}>
                                        <Col span={12}>
                                            <Statistic title="材料总计" value={this.state.status.total}/>
                                        </Col>
                                        <Col span={12}>
                                            <Statistic title="通过率" value={this.state.status.pass_rate} suffix="%" />
                                        </Col>
                                    </Row>
                                </div>
                            </div>
                            <div className="right-box">
                                <Row gutter={16} style={{width:'4.2rem'}}>
                                    <Col span={4}>
                                        <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#87D068'}}></div>通过</div>} value={this.state.status.pass}/>
                                    </Col>
                                    <Col span={4}>
                                        <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#FF5500'}}></div>驳回</div>}  value={this.state.status.return} />
                                    </Col>
                                    <Col span={5}>
                                        <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#FF5500'}}></div>未上传</div>}  value={this.state.status.to_upload} />
                                    </Col>
                                    <Col span={5}>
                                        <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#2DB7F5'}}></div>待审核</div>}  value={this.state.status.to_audit}/>
                                    </Col>
                                    <Col span={5}>
                                        <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#2DB7F5'}}></div>再修改</div>}  value={this.state.status.to_correct} />
                                    </Col>
                                </Row>
                            </div>
                        </div>
                    )
                }
            </ModulaCard>
        )
    }
}
