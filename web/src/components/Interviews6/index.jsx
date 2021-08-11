import React, { Component } from 'react'
import DocumentTitle from 'react-document-title'
import { Row, Col, message, Space, Tag, Button, Upload, Modal, Radio } from 'antd'
import CompletionStatus from '../CompletionStatus'
import TaskRequirements from '../TaskRequirements'
import MaterialWarehouse from '../MaterialWarehouse'
import DataTable from '../DataTable'
import ModifyRecordModal from '../ModifyRecordModal'
import { FileTextTwoTone, UploadOutlined } from '@ant-design/icons';
import './index.less'

export default class index extends Component {

    state = {
        materialWarehouseMode:[{
            identity:'expert',
            mode:[],
            data:[{
                id:'AK000123456',
                fileType:0
            },{
                id:'AK000123456',
                fileType:1
            },{
                id:'AK000123456',
                fileType:2
            },{
                id:'AK000123456',
                fileType:3
            },{
                id:'AK000123456',
                fileType:4
            },{
                id:'AK000123456',
                fileType:0
            },{
                id:'AK000123456',
                fileType:1
            },{
                id:'AK000123456',
                fileType:2
            },{
                id:'AK000123456',
                fileType:3
            },{
                id:'AK000123456',
                fileType:4
            },{
                id:'AK000123456',
                fileType:4
            }]
        },{
            identity:'assistant',
            mode:['upLoad','downLoad'],
            data:[{
                id:'AK000123456',
                fileType:0
            },{
                id:'AK000123456',
                fileType:1
            },{
                id:'AK000123456',
                fileType:2
            },{
                id:'AK000123456',
                fileType:3
            },{
                id:'AK000123456',
                fileType:4
            },{
                id:'AK000123456',
                fileType:0
            },{
                id:'AK000123456',
                fileType:1
            },{
                id:'AK000123456',
                fileType:2
            },{
                id:'AK000123456',
                fileType:3
            },{
                id:'AK000123456',
                fileType:4
            },{
                id:'AK000123456',
                fileType:4
            }]
        },{
            identity:'teacher',
            mode:['upLoad','downLoad'],
            data:[{
                id:'AK000123456',
                fileType:0
            },{
                id:'AK000123456',
                fileType:1
            },{
                id:'AK000123456',
                fileType:2
            },{
                id:'AK000123456',
                fileType:3
            },{
                id:'AK000123456',
                fileType:4
            },{
                id:'AK000123456',
                fileType:0
            },{
                id:'AK000123456',
                fileType:1
            },{
                id:'AK000123456',
                fileType:2
            },{
                id:'AK000123456',
                fileType:3
            },{
                id:'AK000123456',
                fileType:4
            },{
                id:'AK000123456',
                fileType:4
            }]
        }],
        dataCategory:[{
            identity:'expert',
            data:[{
                key: '1',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料1',
                result:0,
                feedbackMaterial:'反馈材料1'
            },{
                key: '2',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料2',
                result:1,
                feedbackMaterial:'反馈材料2'
            },{
                key: '3',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料3',
                result:2,
                feedbackMaterial:'反馈材料3'
            },{
                key: '4',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料4',
                result:3,
                feedbackMaterial:''
            },{
                key: '5',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料1',
                result:0,
                feedbackMaterial:'反馈材料1'
            },{
                key: '6',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料2',
                result:1,
                feedbackMaterial:'反馈材料2'
            },{
                key: '7',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料3',
                result:2,
                feedbackMaterial:'反馈材料3'
            },{
                key: '8',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料4',
                result:3,
                feedbackMaterial:''
            },{
                key: '9',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料1',
                result:0,
                feedbackMaterial:'反馈材料1'
            },{
                key: '10',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料2',
                result:1,
                feedbackMaterial:'反馈材料2'
            },{
                key: '11',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料3',
                result:2,
                feedbackMaterial:'反馈材料3'
            },{
                key: '12',
                date:' 2021-8-10',
                user: 'stacker',
                reviewMaterials:'待评审材料4',
                result:3,
                feedbackMaterial:''
            }],
            colums:[{
                title: '上传时间',
                dataIndex: 'date',
                key: 'date',
                align: 'center',
                width:150
            },{
                title: '上传用户',
                dataIndex: 'user',
                key: 'user',
                align: 'center',
                width:120
            },{
                title: '评审材料',
                dataIndex: 'review-materials',
                key: 'review-materials',
                align: 'center',
                width:220,
                render: (text, record) => (
                    <Space size="middle">
                      <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.reviewMaterials}</span><FileTextTwoTone />
                    </Space>
                )
            },{
                title: '评审结果',
                dataIndex: 'result',
                key: 'result',
                align: 'center',
                width: 100,
                render: (text, record) => {
                    let levelList = [{
                        mode:'通过',
                        color:'#87D068'
                    },{
                        mode:'再修改',
                        color:'#2DB7F5'
                    },{
                        mode:'驳回',
                        color:'#FF5500'
                    },{
                        mode:'审核意见',
                        color:'default'
                    }]
                    return (
                        <Space size="middle">
                            <Tag color={levelList[record.result].color} onClick={()=>{
                                this.setState({
                                    reviewResultsVisible:true
                                })
                            }} style={{cursor:'pointer'}}>{levelList[record.result].mode}</Tag>
                        </Space>
                    )
                }
            },{
                title: '反馈批注材料',
                dataIndex: 'feedback-material',
                key: 'feedback-material',
                align: 'center',
                render: (text, record) => {
                    if(record.feedbackMaterial===""){
                        return (
                            <Space size="middle">
                                <Upload>
                                    <Button icon={<UploadOutlined />} size="small">上传批注</Button>
                                </Upload>
                            </Space>
                        )
                    }
                    return (
                        <Space size="middle">
                          <span>{record.feedbackMaterial}</span><FileTextTwoTone />
                        </Space>
                    )
                }
            }]
        },{
            identity:'assistant',
            data:[{
                key: '1',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料1',
                result:0,
                feedback:'反馈意见',
                isDelay:false,
                isModify:false
            },{
                key: '2',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料2',
                result:1,
                feedback:'',
                isDelay:true,
                isModify:false
            },{
                key: '3',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料1',
                result:2,
                feedback:'反馈意见',
                isDelay:false,
                isModify:false
            },{
                key: '4',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料2',
                result:3,
                feedback:'',
                isDelay:true,
                isModify:false
            },{
                key: '5',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料1',
                result:0,
                feedback:'反馈意见',
                isDelay:false,
                isModify:false
            },{
                key: '6',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料2',
                result:1,
                feedback:'反馈意见',
                isDelay:true,
                isModify:true
            },{
                key: '7',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料1',
                result:2,
                feedback:'反馈意见',
                isDelay:false,
                isModify:false
            },{
                key: '8',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料2',
                result:1,
                feedback:'反馈意见',
                isDelay:true,
                isModify:true
            },{
                key: '9',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料1',
                result:3,
                feedback:'',
                isDelay:false,
                isModify:false
            },{
                key: '10',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料2',
                result:0,
                feedback:'反馈意见',
                isDelay:true,
                isModify:false
            }],
            colums:[{
                title: '材料编号',
                dataIndex: 'id',
                key: 'id',
            },{
                title: '上传时间',
                key: 'date',
                render: (text, record) => (
                    <Space size="middle">
                       <span>{record.date}</span>{record.isDelay?(<Tag color="error">延时提交</Tag>):''}
                    </Space>
                )
            },{
                title: '评审材料',
                key: 'review-materials',
                width:220,
                render: (text, record) => (
                    <Space size="middle">
                      <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.reviewMaterials}</span><FileTextTwoTone />
                    </Space>
                )
            },{
                title: '评审结果',
                key: 'result',
                render: (text, record) => {
                    let levelList = [{
                        mode:'通过',
                        color:'#87D068'
                    },{
                        mode:'再修改',
                        color:'#2DB7F5'
                    },{
                        mode:'驳回',
                        color:'#FF5500'
                    },{
                        mode:'等待评审',
                        color:'default'
                    }]
                    return (
                        <Space size="middle">
                            <Tag color={levelList[record.result].color}>{levelList[record.result].mode}</Tag>
                        </Space>
                    )
                }
            },{
                title: '反馈意见',
                key: 'feedback',
                render: (text, record) => {
                    if(record.feedback===""){
                        return (
                            <Space size="middle">
                                <span style={{cursor:'pointer'}}>等待评审...</span>
                            </Space>
                        )
                    }
                    return (
                        <Space size="middle">
                          <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.feedback}</span><FileTextTwoTone />
                        </Space>
                    )
                }
            },{
                title: '修改记录',
                key: 'isModify',
                align: 'center',
                width:80,
                render: (text, record) => (
                    <Space size="middle">
                       {record.isModify?(<Button type="link" onClick={()=>{
                            this.setState({
                                modifyRecordVisible:true
                            })
                       }}>查看</Button>):(<span>无</span>)}
                    </Space>
                )
            }]
        },{
            identity:'teacher',
            data:[{
                key: '1',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料1',
                result:0,
                feedback:'反馈意见',
                isDelay:false,
                isModify:false
            },{
                key: '2',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料2',
                result:1,
                feedback:'',
                isDelay:true,
                isModify:false
            },{
                key: '3',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料1',
                result:2,
                feedback:'反馈意见',
                isDelay:false,
                isModify:false
            },{
                key: '4',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料2',
                result:3,
                feedback:'',
                isDelay:true,
                isModify:false
            },{
                key: '5',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料1',
                result:0,
                feedback:'反馈意见',
                isDelay:false,
                isModify:false
            },{
                key: '6',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料2',
                result:1,
                feedback:'反馈意见',
                isDelay:true,
                isModify:true
            },{
                key: '7',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料1',
                result:2,
                feedback:'反馈意见',
                isDelay:false,
                isModify:false
            },{
                key: '8',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料2',
                result:1,
                feedback:'反馈意见',
                isDelay:true,
                isModify:true
            },{
                key: '9',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料1',
                result:3,
                feedback:'',
                isDelay:false,
                isModify:false
            },{
                key: '10',
                id:'openct-12345689',
                date: '2021-08-11',
                reviewMaterials:'待评审材料2',
                result:0,
                feedback:'反馈意见',
                isDelay:true,
                isModify:false
            }],
            colums:[{
                title: '材料编号',
                dataIndex: 'id',
                key: 'id',
            },{
                title: '上传时间',
                key: 'date',
                render: (text, record) => (
                    <Space size="middle">
                       <span>{record.date}</span>{record.isDelay?(<Tag color="error">延时提交</Tag>):''}
                    </Space>
                )
            },{
                title: '评审材料',
                key: 'review-materials',
                width:220,
                render: (text, record) => (
                    <Space size="middle">
                      <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.reviewMaterials}</span><FileTextTwoTone />
                    </Space>
                )
            },{
                title: '评审结果',
                key: 'result',
                render: (text, record) => {
                    let levelList = [{
                        mode:'通过',
                        color:'#87D068'
                    },{
                        mode:'再修改',
                        color:'#2DB7F5'
                    },{
                        mode:'驳回',
                        color:'#FF5500'
                    },{
                        mode:'等待评审',
                        color:'default'
                    }]
                    return (
                        <Space size="middle">
                            <Tag color={levelList[record.result].color}>{levelList[record.result].mode}</Tag>
                        </Space>
                    )
                }
            },{
                title: '反馈意见',
                key: 'feedback',
                render: (text, record) => {
                    if(record.feedback===""){
                        return (
                            <Space size="middle">
                                <span style={{cursor:'pointer'}}>等待评审...</span>
                            </Space>
                        )
                    }
                    return (
                        <Space size="middle">
                          <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.feedback}</span><FileTextTwoTone />
                        </Space>
                    )
                }
            },{
                title: '修改记录',
                key: 'isModify',
                align: 'center',
                width:80,
                render: (text, record) => (
                    <Space size="middle">
                       {record.isModify?(<Button type="link" onClick={()=>{
                            this.setState({
                                modifyRecordVisible:true
                            })
                       }}>查看</Button>):(<span>无</span>)}
                    </Space>
                )
            }]
        }],
        identity:'teacher',
        reviewResultsVisible:false,
        modifyRecordVisible:false
    }

    downLoadFile = ()=>{
        message.success('下载成功')
    }

    render() {
        return (
            <DocumentTitle title="项目管理-6人访谈">
                <div className="interviews-6-page" data-component="interviews-6-page">
                    <Row style={{height:'7.48rem',marginBottom:'.22rem'}} gutter={24}>
                        <Col span={9}>
                            <Row style={{height:'1.9rem',marginBottom:'.22rem'}}>
                                <TaskRequirements 
                                    identity="expert" 
                                    describe="任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明，详细内容可参考附件"
                                    file={{
                                        name:'附件:xxxx'
                                    }}
                                    onDownLoadFile={()=>{
                                        message.success('下载成功')
                                    }}
                                    onDateChange = {(date)=>{
                                        console.log(date)
                                    }}
                                />
                                    
                            </Row>
                            <Row style={{height:'1.9rem',marginBottom:'.12rem'}}>
                                <CompletionStatus
                                    title="自2021年4月20日截止今日"
                                    data={{
                                        totalMaterial:93,
                                        passingRate:3,
                                        adopt:3,
                                        reject:3,
                                        notUploaded:3,
                                        reviewed:3,
                                        modify:3
                                    }}
                                />
                            </Row>
                            <Row style={{height:'3.34rem'}}>
                                <MaterialWarehouse
                                    mode={this.state.materialWarehouseMode[this.state.materialWarehouseMode.findIndex(item=>item.identity===this.state.identity)].mode}
                                    fileList={this.state.materialWarehouseMode[this.state.materialWarehouseMode.findIndex(item=>item.identity===this.state.identity)].data}
                                />
                            </Row>
                        </Col>
                        <Col span={15}>
                            <DataTable 
                                title="材料审核"
                                data={this.state.dataCategory[this.state.dataCategory.findIndex(item=>item.identity===this.state.identity)].data}
                                colums={this.state.dataCategory[this.state.dataCategory.findIndex(item=>item.identity===this.state.identity)].colums}
                            />
                        </Col>
                    </Row>
                    <ModifyRecordModal show={this.state.modifyRecordVisible} onCancel={()=>{
                        this.setState({
                            modifyRecordVisible:false
                        })
                    }} onComplete={()=>{
                        this.setState({
                            modifyRecordVisible:false
                        })
                    }} />
                    <Modal title="评审意见" visible={this.state.reviewResultsVisible} onOk={()=>{
                        this.setState({
                            reviewResultsVisible:false
                        })
                    }} onCancel={()=>{
                        this.setState({
                            reviewResultsVisible:false
                        })
                    }}>
                        <span>材料评审结果：</span>
                        <Radio.Group name="radiogroup" defaultValue={1}>
                            <Radio value={1}>通过</Radio>
                            <Radio value={2}>再修改</Radio>
                            <Radio value={3}>驳回</Radio>
                        </Radio.Group>
                    </Modal>
                </div>
            </DocumentTitle>
        )
    }
}
