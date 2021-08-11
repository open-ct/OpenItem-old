//抽离出完成情况组件

import React, { Component } from 'react'
import { Col, Statistic,Row} from 'antd'
import ModulaCard from '../ModulaCard'
import './index.less'

export default class index extends Component {
    render() {
        return (
            <ModulaCard title="完成情况">
                <div className="completion-status-box" data-component="completion-status-box">
                    <div className="left-box">
                        <div className="title">
                            <span>{this.props.title}</span>
                        </div>
                        <div className="value-list">
                            <Row gutter={16}>
                                <Col span={12}>
                                    <Statistic title="材料总计" value={this.props.data.totalMaterial}/>
                                </Col>
                                <Col span={12}>
                                    <Statistic title="通过率" value={this.props.data.passingRate} suffix="%" />
                                </Col>
                            </Row>
                        </div>
                    </div>
                    <div className="right-box">
                        <Row gutter={16} style={{width:'4.2rem'}}>
                            <Col span={4}>
                                <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#87D068'}}></div>通过</div>} value={this.props.data.adopt}/>
                            </Col>
                            <Col span={4}>
                                <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#FF5500'}}></div>驳回</div>}  value={this.props.data.reject} />
                            </Col>
                            <Col span={5}>
                                <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#FF5500'}}></div>未上传</div>}  value={this.props.data.notUploaded} />
                            </Col>
                            <Col span={5}>
                                <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#2DB7F5'}}></div>待审核</div>}  value={this.props.data.reviewed}/>
                            </Col>
                            <Col span={5}>
                                <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#2DB7F5'}}></div>再修改</div>}  value={this.props.data.modify} />
                            </Col>
                        </Row>
                    </div>
                </div>
            </ModulaCard>
            
        )
    }
}
