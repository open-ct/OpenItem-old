import React, { Component } from 'react'
import { Row, Col } from 'antd'
import TaskRequirements from '../TaskRequirements'
import CompletionStatus from '../CompletionStatus'
import MaterialWarehouse from '../MaterialWarehouse'
import DataTable from '../DataTable'
import './index.less'

export default class index extends Component {
    
    render() {
        return (
            <div className="step-page" data-component="step-page">
                <Row style={{height:'7.48rem',marginBottom:'.22rem'}} gutter={24}>
                    <Col span={9}>
                        <Row style={{height:'1.9rem',marginBottom:'.22rem'}}>
                            <TaskRequirements
                                role = {this.props.match.params.role}
                                stepId = {this.props.match.params.step_id}
                                projectId = {this.props.match.params.project_id}
                            ></TaskRequirements>
                        </Row>
                        <Row style={{height:'1.9rem',marginBottom:'.12rem'}}>
                            <CompletionStatus
                                title="自项目创建截止今日情况"
                                stepId = {this.props.match.params.step_id}
                            />
                        </Row>
                        <Row style={{height:'3.34rem'}}>
                            <MaterialWarehouse
                                role = {this.props.match.params.role}
                                projectId = {this.props.match.params.project_id}
                                stepId = {this.props.match.params.step_id}
                                stepName = {this.props.location.pathname.split('/')[this.props.location.pathname.split('/').length-2]}
                            > 
                            </MaterialWarehouse>
                        </Row>
                    </Col>
                    <Col span={15}>
                        <DataTable
                            title="材料评审"
                            stepName = {this.props.location.pathname.split('/')[this.props.location.pathname.split('/').length-2]}
                            role = {this.props.match.params.role}
                            stepId={this.props.match.params.step_id}
                        >
                        </DataTable>
                    </Col>
                </Row>
            </div>
        )
    }
}
