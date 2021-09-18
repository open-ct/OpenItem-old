/*
 * @Author: your name
 * @Date: 2021-07-31 16:42:56
 * @LastEditTime: 2021-09-18 18:17:53
 * @LastEditors: Please set LastEditors
 * @Description: 抽离出的任务要求组件
 * @FilePath: \OpenItem\web-client\src\index.js
 */


/**
 * @description: 组件参数
 * @identity 用户身份
 * @describe 要求描述
 * @onDateChange 日期改变回调
 */
import React, { Component } from 'react'
import { Button, Upload, message } from 'antd'
import ModulaCard from '../ModulaCard'
import CalendarButton from '../CalendarButton'
import {FieldTimeOutlined,LinkOutlined} from '@ant-design/icons'
import request from '../../utils/request'
import './index.less'

export default class index extends Component {

    state = {
        stepInfo:{},
        loadingState:false,
        upLoadState:false
    }

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
            request({
                url:'/proj/step/attachment',
                method:"POST",
                data:{
                    files_ids:[res.data.uuid],
                    step_id:this.props.stepId
                }
            }).then(res=>{
                console.log(res)
                this.setState({
                    upLoadState:false
                })
                this.getStepInfo()
                message.success("文件上传成功")
            }).catch(err=>{
                this.setState({
                    upLoadState:false
                })
                message.error("文件上传失败")
            })
        }).catch(err=>{
            this.setState({
                upLoadState:false
            })
            message.error("文件上传失败")
        })
    }

    downLoadFile = ()=>{
        if(this.state.stepInfo.attachments.length===0){
            message.warning("暂无可下载附件")
        }else{
            message.info(`开始下载文件：${this.state.stepInfo.attachments[0]}！`);
            request({
                url:`/file/${this.state.stepInfo.attachments[0]}`,
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
    }

    //根据不同身份加载不同操作情况
    operationRender = ()=>{
        let operationDomList = [()=>(
            <div className="choice-box">
                <CalendarButton label="截止时间：2021年8月20日" icon={<FieldTimeOutlined />} onDateChange={(date)=>{
                    message.error("改操作暂不可用")
                }}/>
                <Upload
                    name="filename" 
                    showUploadList={false}
                    beforeUpload = {()=>{
                        if(this.state.upLoadState){
                            message.error("当前存在上传中文件，请勿频繁上传！");
                            return false
                        }
                        return true
                    }} 
                    customRequest={this.upLoadFile.bind(this)}
                >
                    <Button size="small" icon={<LinkOutlined />} style={{marginLeft:'.1246rem'}}>{
                        this.state.stepInfo.attachments?"已上传":"暂无文件"
                    }</Button>
                </Upload>
            </div>
        ),()=>(
            <div className="choice-box">
                <Button size="small" style={{marginRight:'.1246rem'}} icon={<FieldTimeOutlined />}>截止时间：2021年8月20日</Button>
                <Button size="small" icon={<LinkOutlined />} onClick={this.downLoadFile}>{
                    this.state.stepInfo.attachments?"已上传":"暂无文件"
                }</Button>
            </div>
        ),()=>(
            <div className="choice-box">
                <Button size="small" style={{marginRight:'.1246rem'}} icon={<FieldTimeOutlined />}>截止时间：2021年8月20日</Button>
                <Button size="small" icon={<LinkOutlined />} onClick={this.downLoadFile}>{
                    this.state.stepInfo.attachments?"已上传":"暂无文件"
                }</Button>
            </div>
        )]
        return operationDomList[this.props.role-2]()
    }

    componentDidMount(){
        this.getStepInfo()
    }

    getStepInfo = ()=>{
        this.setState({
            loadingState:true,
        })
        request({
            url:`/proj/step/${this.props.stepId}`,
            method:'GET'
        }).then(res=>{
            this.setState({
                stepInfo:res.data,
                loadingState:false,
            })
        }).catch(err=>{
            this.setState({
                loadingState:false,
            })
        })
    }

    render() {
        return (
            <ModulaCard title="任务要求">
                {
                    this.state.loadingState?(
                        <></>
                    ):(
                        <div className="task-requirements-box" data-component="task-requirements-box">
                            <div className="describe-box">
                                <div className="context">
                                    <p>{this.state.stepInfo.requirement||"无"}</p>
                                </div>
                                <div className="btn-box">
                                    {this.props.role===2?(<Button type="link">编辑</Button>):''}
                                </div>
                            </div>
                            {this.operationRender()}
                        </div>
                    )
                }
            </ModulaCard>
        )
    }
}
