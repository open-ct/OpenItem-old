//上传文件弹出框组件

import React, { Component } from 'react'
import { Modal, Form, Radio, Upload, Input, message } from 'antd'
import { InboxOutlined } from '@ant-design/icons'
import {getCookie} from '../../utils/cookies'
import request from '../../utils/request'
import './index.less'

export default class index extends Component {
    state = {
        upLoadFileList: [],
        upLoadState:false,//用于控制文件上传
        createLoading:false,
        submitForm:{
            type:1
        }
    }

    upLoadFormRef = React.createRef()

    upLoadFile(info){
        this.setState({
            upLoadState:true
        })
        message.info("开始上传文件："+info.file.name)
        const formData = new FormData()
        formData.append('file', info.file)
        formData.append('source_project',this.props.projectId)
        formData.append('tags', "无")
        formData.append('description', "步骤附件")
        request({
            url:'/file',
            method:"POST",
            data:formData,
            headers:{
                "Content-Type":"multipart/form-data;"
            }
        }).then(res=>{
            console.log(res)
            let upLoadFileList = Object.assign(this.state.upLoadFileList,{})
            upLoadFileList[upLoadFileList.length-1].status = "done"
            upLoadFileList[upLoadFileList.length-1].id = res.data.uuid
            this.setState({
                upLoadFileList,
                upLoadState:false
            })
            message.success("文件上传成功")
        }).catch(err=>{
            let upLoadFileList = Object.assign(this.state.upLoadFileList,{})
            upLoadFileList[upLoadFileList.length-1].status = "error"
            this.setState({
                upLoadFileList,
                upLoadState:false
            })
            message.error("文件上传失败")
        })
    }

    createSubmit = ()=>{
        if(this.state.upLoadFileList.length===0){
            message.warning("请至少上传一个文件")
        }else{
            this.setState({
                createLoading:true
            })
            request({
                url:"/proj/submit",
                method:"POST",
                data:{
                    description:"暂无",
                    step_id:this.props.stepId,
                    title:"一个新的提交"
                }
            }).then(res=>{
                request({
                    url:"/proj/submit/content",
                    method:"POST",
                    data:{
                        comment:"一个新的文件",
                        item_id:this.state.upLoadFileList[0].id,
                        submit_id:res.data.uuid,
                        type:this.state.submitForm.type,
                        version:"1.0.0"
                    }
                }).then(res=>{
                    this.upLoadFormRef.current.resetFields()
                    this.setState({
                        createLoading:false,
                        upLoadFileList:[]
                    })
                    this.props.onClose()
                    message.success("创建成功")
                    this.props.onUpdate()
                }).catch(err=>{
                    this.setState({
                        createLoading:false
                    })
                    message.error(err.message||"创建失败")
                })
            }).catch(err=>{
                this.setState({
                    createLoading:false
                })
                message.error(err.message||"创建失败")
            })
        }
    }

    render() {
        return (
            <Modal title="上传材料" 
                width="6.26rem" 
                okText="确认创建"
                cancelText="放弃创建"
                visible={this.props.show} 
                closable={!this.state.createLoading}
                keyboard={!this.state.createLoading}
                confirmLoading={this.state.createLoading}
                onCancel={()=>{
                    if(this.state.createLoading){
                        message.info("请等待")
                    }else{
                        this.upLoadFormRef.current.resetFields()
                        this.setState({
                            upLoadFileList:[]
                        })
                        this.props.onClose()
                    }
                }}
                onOk={this.createSubmit}
            >
                <div className="up-load-modal-box" data-component="up-load-modal-box">
                    <Form
                        name="material"
                        labelCol={{ span: 6 }}
                        wrapperCol={{ span: 18 }}
                        initialValues = {this.state.submitForm}
                        ref={this.upLoadFormRef}
                    >
                        <Form.Item
                            label="材料内容"
                            name="type"
                            rules={[{ required: true}]}
                        >
                            <Radio.Group size="small"
                                onChange={(e)=>{
                                    this.setState({
                                        upLoadFormData:Object.assign({}, this.state.submitForm, {type:e.target.value})
                                    })
                                }}
                            >
                                <Radio value={1} key="option_1">新材料</Radio>
                                <Radio value={2} key="option_2">再修改材料</Radio>
                                <Radio value={3} key="option_3">驳回修改材料</Radio>
                            </Radio.Group>
                        </Form.Item>

                        <Form.Item
                            label="材料编号"
                            name="material-id"
                            className="material__id"
                            rules={[{ required: true}]}
                            extra="新材料自动生成材料编号，不可修改"
                        >
                            <Input placeholder="AK9800000123" disabled/>
                        </Form.Item>

                        <Form.Item 
                            label="上传"
                            rules={[{ required: true}]}
                            labelCol={{span:4}}
                            required={true}
                        >
                            <Upload.Dragger 
                                name="filename"                                  
                                fileList={this.state.upLoadFileList}
                                headers={{
                                    "Token":getCookie('token')
                                }}
                                beforeUpload={(e)=>{
                                    if(this.state.upLoadState||this.state.createLoading){
                                        message.warning("当前存在上传中文件，请等待")
                                        return false
                                    }else{
                                        let upLoadFileList = [...this.state.upLoadFileList]
                                        upLoadFileList.push({
                                            name:e.name,
                                            status:'uploading'
                                        })
                                        this.setState({
                                            upLoadFileList
                                        })
                                        return true
                                    }
                                }}
                                customRequest={this.upLoadFile.bind(this)}
                                onRemove={(e)=>{
                                    //uid
                                    let upLoadFileList = Object.assign(this.state.upLoadFileList,{})
                                    upLoadFileList = upLoadFileList.splice(upLoadFileList.findIndex(item=>item.uid===e.uid),1)
                                    this.setState({
                                        upLoadFileList
                                    })
                                    console.log(this.state.upLoadFileList)
                                }}
                            >
                                <p className="ant-upload-drag-icon">
                                <InboxOutlined />
                                </p>
                                <p className="ant-upload-text">Click or drag file to this area to upload</p>
                                <p className="ant-upload-hint" style={{fontSize:'.1rem'}}>Support for a single or bulk upload. Strictly prohibit from uploading company data or other band files</p>
                            </Upload.Dragger>

                        </Form.Item>
                    </Form>
                </div>
            </Modal>
        )
    }
}
