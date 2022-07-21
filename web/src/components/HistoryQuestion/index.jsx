import React, { Component } from 'react'
import {Row,Col,Button} from 'antd'
import './index.less'

export default class index extends Component {
    render() {
        return (
            <div className="history-question-item" data-component="history-question-item">
                <Row gutter={[16,16]} className="queston-content">
                    <Col span="16" className="question-title">
                        <span>题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目......题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目.题目题目题目题目题目题目题目题目.(最多5行过多省略)题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目......题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目.题目题目题目题目题目题目题目题目.(最多5行过多省略)</span>
                    </Col>
                    <Col span="8" className="question-params">
                        <Row className="params-item">
                            <Col span="12" className="title">
                                <span>学科名称：</span>
                            </Col>
                            <Col span="12" className="value" >
                                <span>数学</span>
                            </Col>
                        </Row>
                        <Row className="params-item">
                            <Col span="12" className="title">
                                <span>能力纬度：</span>
                            </Col>
                            <Col span="12" className="value">
                                <span>参数1；参数2；参数3</span>
                            </Col>
                        </Row>
                        <Row className="params-item">
                            <Col span="12" className="title">
                                <span>内容纬度</span>
                            </Col>
                            <Col span="12" className="value">
                                <span>参数1；参数2；参数3</span>
                            </Col>
                        </Row>
                        <Row className="params-item">
                            <Col span="12" className="title">
                                <span>难度等级：</span>
                            </Col>
                            <Col span="12" className="value">
                                <span>参数1</span>
                            </Col>
                        </Row>
                    </Col>
                </Row>
                <Row gutter={[16,16]} className="question-footer">
                    <Col span="16" className="question-answer">
                        <span>答案：xxxx</span>
                    </Col>
                    <Col span="8">
                        <Button type="primary" size="small" style={{backgroundColor:"#56BFFF",borderColor:"#56BFFF"}}>使用</Button>
                    </Col>
                </Row>
            </div>
        )
    }
}
