import React, { Component } from 'react'
import { Space,  Modal, Table, message,Spin } from 'antd'
import { FileTextTwoTone } from '@ant-design/icons'
import request from '../../utils/request'
import './index.less'

export default class index extends Component {

    state = {
        recordList:[{
            key:'1',
            date:'2021-08-11',
            result:0,
            reviewMaterials:'待评审材料1',
            feedback:'反馈意见',
        },{
            key:'2',
            date:'2021-08-11',
            result:1,
            reviewMaterials:'待评审材料1',
            feedback:'反馈意见',
        },{
            key:'3',
            date:'2021-08-11',
            result:2,
            reviewMaterials:'待评审材料1',
            feedback:'反馈意见',
        },{
            key:'4',
            date:'2021-08-11',
            result:3,
            reviewMaterials:'待评审材料1',
            feedback:'',
        }],
        submitInfo:{},
        loadingState:false 
    }

    getRecordList(){
        this.setState({
            loadingState:true
        })
        request({
            url:`/proj/submit/${this.props.submitId}`,
            method:"GET"
        }).then(res=>{
            this.setState({
                submitInfo:res.data,
                loadingState:false
            })
        }).catch(err=>{
            this.setState({
                loadingState:false
            })
            message.error("加载错误")
            this.props.onComplete()
        })
    }

    recordListColums = [{
        title: 'ID',
        dataIndex: 'uuid',
        align: 'center',
        key: 'id',
    },{
        title: '评论',
        dataIndex: 'comment',
        align: 'center',
        key: 'comment',
    },{
        title: '评审材料',
        key: 'review-materials',
        align: 'center',
        width:220,
        render: (text, record) => (
            <Space size="middle">
              <span style={{cursor:'pointer'}} onClick={this.downLoadFile.bind(this,record.uuid)}>点击下载</span><FileTextTwoTone />
            </Space>
        )
    },{
        title: '版本',
        dataIndex: 'version',
        align: 'center',
        key: 'version',
    }]

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


    render() {
        return (
            <Modal title="修改记录" 
                width="9.51rem" 
                visible={this.props.show}
                cancelText="关闭" 
                okText="确认"
                confirmLoading={this.state.loadingState}
                closable={!this.state.loadingState}
                maskClosable={!this.state.loadingState}
                keyboard={!this.state.loadingState} 
                onOk={()=>{
                    this.props.onComplete()
                }} 
                onCancel={()=>{
                    this.props.onCancel()
                }}
            >
                {
                    this.state.loadingState?(
                        <Spin spinning={this.state.loadingState} tip="加载中"></Spin>
                    ):(
                        <div className="modify-record-box" data-component="modify-record-box">
                            <div className="basic">
                                <div className="id">
                                    材料编号:{this.state.submitInfo.uuid}
                                </div>
                                <div className="date">
                                    截止日期:未知
                                </div>
                            </div>
                            <Table 
                                dataSource={this.state.submitInfo.contents||[]} 
                                columns={this.recordListColums} size="small" 
                                pagination={false} 
                            />
                        </div>
                    )
                }   
            </Modal>
        )
    }
}
