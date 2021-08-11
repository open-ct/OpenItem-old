/*
 * @Author: your name
 * @Date: 2021-07-31 16:42:56
 * @LastEditTime: 2021-08-12 01:39:40
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
import { Button, Upload, } from 'antd'
import ModulaCard from '../ModulaCard'
import CalendarButton from '../CalendarButton'

import {FieldTimeOutlined,LinkOutlined} from '@ant-design/icons'
import './index.less'

export default class index extends Component {
    state = {
        identity:[{
            name:'expert',
            isEditer:true
        },{
            name:'assistant',
            isEditer:false
        },{
            name:'teacher',
            isEditer:true
        }]
    }
    //根据不同身份加载不同操作情况
    operationRender = ()=>{
        let operationDomList = [{
            identity:'expert',
            dom:()=>(
                <div className="choice-box">
                    <CalendarButton label="截止时间：2021年8月20日" icon={<FieldTimeOutlined />} onDateChange={(date)=>{
                        this.props.onDateChange(date)
                    }}/>
                    <Upload >
                        <Button size="small" icon={<LinkOutlined />} style={{marginLeft:'.1246rem'}}>{this.props.file.name}</Button>
                    </Upload>
                </div>
            )
        },{
            identity:'assistant',
            dom:()=>(
                <div className="choice-box">
                    <Button size="small" style={{marginRight:'.1246rem'}} icon={<FieldTimeOutlined />}>截止时间：2021年8月20日</Button>
                    <Button size="small" icon={<LinkOutlined />} onClick={this.props.onDownLoadFile}>{this.props.file.name}</Button>
                </div>
            )
        },{
            identity:'teacher',
            dom:()=>(
                <div className="choice-box">
                    <Button size="small" style={{marginRight:'.1246rem'}} icon={<FieldTimeOutlined />}>截止时间：2021年8月20日</Button>
                    <Button size="small" icon={<LinkOutlined />} onClick={this.props.onDownLoadFile}>{this.props.file.name}</Button>
                </div>
            )
        }]
        return operationDomList[operationDomList.findIndex(item=>item.identity===this.props.identity)].dom()
    }

    render() {
        return (
            <ModulaCard title="任务要求">
                <div className="task-requirements-box" data-component="task-requirements-box">
                    <div className="describe-box">
                        <div className="context">
                            <p>{this.props.describe}</p>
                        </div>
                        <div className="btn-box">
                            {this.state.identity[this.state.identity.findIndex(item=>item.name===this.props.identity&&item.isEditer)]?(<Button type="link">编辑</Button>):''}
                        </div>
                    </div>
                    {this.operationRender()}
                </div>
            </ModulaCard>
        )
    }
}
