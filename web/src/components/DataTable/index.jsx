import React, { Component } from 'react'
import { Table, Pagination, message, Space,Tag, Upload,Button,Modal,Radio } from 'antd'
import ModulaCard from '../ModulaCard'
import ModifyRecordModal from '../ModifyRecordModal'
import request from '../../utils/request'
import { FileTextTwoTone, UploadOutlined } from '@ant-design/icons';
import './index.less'

export default class index extends Component {

    state = {
        data:[],
        loadingState:false,
        reviewResultsVisible:false,
        statusChangeLoading:false,
        modifyRecordVisible:false,
        selectedSubmitId:"",
        statusChangeParams:{
            value:1,
            submitId:""
        }
    }

    modifyRecordRef = React.createRef()

    colums=[[{
            title: '上传时间',
            dataIndex: 'CreateAt',
            key: 'CreateAt',
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
            key: 'contents',
            align: 'center',
            width:220,
            render: (text, record) => (
                <Space size="middle">
                    {
                        record.contents?(
                            <span style={{cursor:'pointer'}} onClick={this.downLoadFile.bind(this,record.contents[0].item_id)}>{record.contents[0].comment} <FileTextTwoTone /></span>
                        ):"无"
                    }
                </Space>
            )
        },{
            title: '评审结果',
            key: 'status',
            align: 'center',
            width: 100,
            render: (text, record) => {
                let levelList = [{
                    mode:'等待审核',
                    color:'default'
                },{
                    mode:'通过',
                    color:'#87D068'
                },{
                    mode:'再修改',
                    color:'#2DB7F5'
                },{
                    mode:'驳回',
                    color:'#FF5500'
                }]
                return (
                    <Space size="middle">
                        <Tag color={levelList[record.status].color} onClick={()=>{
                            this.setState({
                                reviewResultsVisible:true,
                                statusChangeParams:{
                                    submitId:record.uuid,
                                    value:record.status
                                }
                            })
                        }} style={{cursor:'pointer'}}>{levelList[record.status].mode}</Tag>
                    </Space>
                )
            }
        },{
            title: '反馈批注材料',
            dataIndex: 'feedback-material',
            key: 'feedback-material',
            align: 'center',
            render: (text, record) => {
                if(!record.feedbackMaterial){
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
        }],[{
            title: '材料编号',
            dataIndex: 'uuid',
            key: 'id',
            align: 'center',
        },{
            title: '上传时间',
            key: 'date',
            align: 'center',
            render: (text, record) => (
                <Space size="middle">
                   <span>{record.CreateAt}</span>{record.isDelay?(<Tag color="error">延时提交</Tag>):''}
                </Space>
            )
        },{
            title: '评审材料',
            key: 'review-materials',
            width:220,
            align: 'center',
            render: (text, record) => (
                <Space size="middle">
                    {
                        record.contents?(
                            <span style={{cursor:'pointer'}} onClick={this.downLoadFile.bind(this,record.contents[0].item_id)}>{record.contents[0].comment} <FileTextTwoTone /></span>
                        ):"无"
                    }
                </Space>
            )
        },{
            title: '评审结果',
            key: 'result',
            align: 'center',
            render: (text, record) => {
                let levelList = [{
                    mode:'等待评审',
                    color:'default'
                },{
                    mode:'通过',
                    color:'#87D068'
                },{
                    mode:'再修改',
                    color:'#2DB7F5'
                },{
                    mode:'驳回',
                    color:'#FF5500'
                }]
                return (
                    <Space size="middle">
                        <Tag color={levelList[record.status].color}  style={{cursor:'pointer'}}>{levelList[record.status].mode}</Tag>
                    </Space>
                )
            }
        },{
            title: '反馈意见',
            key: 'feedback',
            align: 'center',
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
                      {/* <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.feedback}</span><FileTextTwoTone /> */}
                      <span>无</span>
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
                   {record.contents?(<Button type="link" onClick={()=>{
                        this.setState({
                            selectedSubmitId:record.uuid,
                            modifyRecordVisible:true
                        })
                        this.modifyRecordRef.current.getRecordList()
                   }}>查看</Button>):(<span>无</span>)}
                </Space>
            )
        }],[{
            title: '上传时间',
            dataIndex: 'CreateAt',
            key: 'CreateAt',
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
            key: 'contents',
            align: 'center',
            width:220,
            render: (text, record) => (
                <Space size="middle">
                    {
                        record.contents?(
                            <span style={{cursor:'pointer'}} onClick={this.downLoadFile.bind(this,record.contents[0].item_id)}>{record.contents[0].comment} <FileTextTwoTone /></span>
                        ):"无"
                    }
                </Space>
            )
        },{
            title: '评审结果',
            key: 'status',
            align: 'center',
            width: 100,
            render: (text, record) => {
                let levelList = [{
                    mode:'等待审核',
                    color:'default'
                },{
                    mode:'通过',
                    color:'#87D068'
                },{
                    mode:'再修改',
                    color:'#2DB7F5'
                },{
                    mode:'驳回',
                    color:'#FF5500'
                }]
                return (
                    <Space size="middle">
                        <Tag color={levelList[record.status].color}  style={{cursor:'pointer'}}>{levelList[record.status].mode}</Tag>
                    </Space>
                )
            }
        },{
            title: '反馈批注材料',
            dataIndex: 'feedback-material',
            key: 'feedback-material',
            align: 'center',
            render: (text, record) => {
                return (
                    <Space size="middle">
                      <span>无</span>
                    </Space>
                )
            }
        }]
    ]

    downLoadFile(file_id){
        message.info(`开始下载文件：${file_id}！`);
        request({
            url:`/file/${file_id}`,
            method: 'GET',
            responseType:'blob'
        }).then(res=>{
            console.log(res)
            const filename = res.headers["content-disposition"];
            const blob = new Blob([res.data]);
            var downloadElement = document.createElement("a");
            var href = window.URL.createObjectURL(blob);
            downloadElement.href = href;
            downloadElement.download = decodeURIComponent(filename.split("filename*=")[1].replace("utf-8''",""));
            document.body.appendChild(downloadElement);
            downloadElement.click();
            document.body.removeChild(downloadElement);
            window.URL.revokeObjectURL(href); 
            message.success("文件下载成功！");
        }).catch(err=>{
            message.error("文件下载失败！");
        })
    }

    loadClumsIndex = ()=>{
        if(this.props.role==="2"){
            return 0
        }else if(this.props.role==="3"){
            return 1
        }else if(this.props.role==="4"&&this.props.stepName==="测试框架与论证报告"){
            return 2
        }else if(this.props.role==="4"&&this.props.stepName!=="测试框架与论证报告"){
            return 1
        }else{
            return 0
        }
    }

    componentDidMount(){
        this.getDataList()
    }

    getDataList = ()=>{
        this.setState({
            loadingState:true
        })
        request({ method:'GET', url:`/proj/submits/${this.props.stepId}`}).then(res=>{
            this.setState({
                data:res.data,
                loadingState:false
            })
        }).catch(err=>{
            message.error(err.message||"审查材料加载失败！")
            this.setState({
                loadingState:false
            })
        })
    }

    render() {
        return (
            <ModulaCard title={this.props.title}>
                <div className="data-table-box" data-component="data-table-box">
                    <Table 
                        dataSource={this.state.data} 
                        columns={this.colums[this.loadClumsIndex()]} 
                        size="small" 
                        rowKey="Id"
                        pagination={false} 
                        scroll={{ y: '5.8rem'}}
                        loading={this.state.loadingState}
                    />
                    <div className="footer">
                        <Pagination
                            total={85}
                            showTotal={total => `Total ${total} items`}
                            defaultPageSize={20}
                            defaultCurrent={1}
                            size="small"
                        />
                    </div>
                </div>
                <Modal title="评审意见" visible={this.state.reviewResultsVisible}
                    cancelText="关闭" 
                    okText="审核"
                    confirmLoading={this.state.statusChangeLoading}
                    closable={!this.state.statusChangeLoading}
                    maskClosable={!this.state.statusChangeLoading}
                    keyboard={!this.state.statusChangeLoading}
                    onOk={()=>{
                        this.setState({
                            statusChangeLoading:true
                        })
                        request({
                            url:`/proj/submit`,
                            method: 'PUT',
                            data:{
                                new_status:this.state.statusChangeParams.value,
                                submit_id:this.state.statusChangeParams.submitId
                            }
                        }).then(res=>{
                            this.setState({
                                statusChangeLoading:false,
                                reviewResultsVisible:false
                            })
                            this.getDataList()
                            message.success("审核成功")
                            
                        }).catch(err=>{
                            this.setState({
                                statusChangeLoading:false
                            })
                            message.error(err.message||"审核失败")
                        })
                    }} 
                    onCancel={()=>{
                        if(this.state.statusChangeLoading){
                            message.warning("请等待")
                        }else{
                            this.setState({
                                reviewResultsVisible:false
                            })
                        }
                    }}
                >
                    <span>材料评审结果：</span>
                    <Radio.Group name="radiogroup" value={this.state.statusChangeParams.value} onChange={(e)=>{
                        let statusChangeParams = Object.assign(this.state.statusChangeParams,{
                            value:e.target.value
                        })
                        this.setState({
                            statusChangeParams
                        })
                    }}>
                        <Radio value={1}>通过</Radio>
                        <Radio value={2}>再修改</Radio>
                        <Radio value={3}>驳回</Radio>
                    </Radio.Group>
                </Modal>
                <ModifyRecordModal 
                    show={this.state.modifyRecordVisible}
                    submitId={this.state.selectedSubmitId} 
                    ref={this.modifyRecordRef}
                    onCancel={()=>{
                        this.setState({
                            modifyRecordVisible:false
                        })
                    }} 
                    onComplete={()=>{
                        this.setState({
                            modifyRecordVisible:false
                        })
                    }} 
                />
            </ModulaCard>
        )
    }
}
