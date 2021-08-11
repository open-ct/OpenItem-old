import React, { Component } from 'react'
import { Space, Tag, Modal, Table } from 'antd'
import { FileTextTwoTone } from '@ant-design/icons'
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
    }

    recordListColums = [{
        title: '上传时间',
        dataIndex: 'date',
        key: 'date',
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
    }]

    render() {
        return (
            <Modal title="修改记录" width="9.51rem" visible={this.props.show} onOk={()=>{
                this.props.onComplete()
            }} onCancel={()=>{
                this.props.onCancel()
            }}>
                <div className="modify-record-box" data-component="modify-record-box">
                    <div className="basic">
                        <div className="id">
                            材料编号:openct15646541
                        </div>
                        <div className="date">
                            截止日期:2021-08-11
                        </div>
                    </div>
                    <Table 
                        dataSource={this.state.recordList} 
                        columns={this.recordListColums} size="small" 
                        pagination={false} 
                    />
                </div>
            </Modal>
        )
    }
}
