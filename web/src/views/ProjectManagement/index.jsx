import React, { Component } from 'react'
import { Switch, Route, Redirect } from 'react-router-dom'
import DocumentTitle from 'react-document-title'
import { PageHeader, Tabs, Button, Descriptions } from 'antd';
import PropositionGuide from '../../components/PropositionGuide'
import './index.less'

const { TabPane } = Tabs;
  
export default class index extends Component {

    render() {
        return (
            <DocumentTitle title="项目管理">
                <div className="project-management-page" data-component="project-management-page">
                    <PageHeader
                        className="site-page-header-responsive"
                        onBack={() => window.history.back()}
                        title="项目管理"
                        subTitle="2016年区域教学质量"
                        extra={[
                            <Button>编辑项目</Button>,
                            <Button>导出成员</Button>
                        ]}
                        footer={
                            <Tabs defaultActiveKey="proposition-guide" type="card" onChange={(e)=>{
                                console.log(e)
                            }}>
                                <TabPane tab="命题指南与测试" key="proposition-guide">
                                </TabPane>
                                <TabPane tab="6人访谈" key="6-interviews"></TabPane>
                                <TabPane tab="30人测试" key="30-person-test"></TabPane>
                                <TabPane tab="300人测试" key="300-person-test"></TabPane>
                                <TabPane tab="组卷" key="paper-formation"></TabPane>
                                <TabPane tab="最终试卷" key="final-paper"></TabPane>
                            </Tabs>
                        }
                        >
                        <Descriptions size="small" column={3}>
                            <Descriptions.Item label="创建时间" key="1">2016-6-20</Descriptions.Item>
                            <Descriptions.Item label="学科" key="2">语文</Descriptions.Item>
                            <Descriptions.Item label="学段" key="3">初中</Descriptions.Item>
                            <Descriptions.Item label="试卷" key="4">8</Descriptions.Item>
                            <Descriptions.Item label="试题" key="5">160</Descriptions.Item>
                        </Descriptions>
                    </PageHeader>
                    <div className="container">
                        <Switch>
                            <Redirect from="/home/project-management" to="/home/project-management/proposition-guide" exact></Redirect>
                            <Route path="/home/project-management/proposition-guide" component={PropositionGuide} exact></Route>
                        </Switch>
                    </div>
                </div>
            </DocumentTitle>
        )
    }
}
