import React, { Component } from 'react'
import { Layout } from 'antd';
import DocumentTitle from 'react-document-title'
import './index.less'
const { Header , Sider, Content } = Layout;

export default class index extends Component {
    render() {
        return (
            <DocumentTitle title="命题与题库系统">
                <Layout className="home-page" data-component="home-page">
                    <Header>Header</Header>
                    <Layout className="container">
                        <Sider>Sider</Sider>
                        <Content>Content</Content>
                    </Layout>
                </Layout>
            </DocumentTitle> 
        )
    }
}
