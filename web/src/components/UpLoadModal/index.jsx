//上传文件弹出框组件

import React, { Component } from 'react'
import { Modal, Form, Button, Radio, Upload, Input } from 'antd'
import { InboxOutlined } from '@ant-design/icons'
import './index.less'

export default class index extends Component {
    state = {
        defaultFileList: [{
                uid: '1',
                name: 'xxx.png',
                status: 'done',
                response: 'Server Error 500', // custom error message to show
                url: 'http://www.baidu.com/xxx.png',
        },{
                uid: '2',
                name: 'yyy.png',
                status: 'done',
                url: 'http://www.baidu.com/yyy.png',
        },{
                uid: '3',
                name: 'zzz.png',
                status: 'error',
                response: 'Server Error 500', // custom error message to show
                url: 'http://www.baidu.com/zzz.png',
        }],
    }
    render() {
        return (
            <Modal title="上传材料" width="6.26rem" visible={this.props.show}  closable={false} footer={[(
                <Button key="upload" type="primary" onClick={()=>{
                    this.props.onLoad()
                }}>
                    上传
                </Button>
            )]}>
                <div className="up-load-modal-box" data-component="up-load-modal-box">
                    <Form
                        name="material"
                        labelCol={{ span: 6 }}
                        wrapperCol={{ span: 18 }}
                    >
                        <Form.Item
                            label="材料内容"
                            name="category-materials-box"
                            rules={[{ required: true}]}
                        >
                            <Radio.Group size="small">
                                <Radio value={1}>新材料</Radio>
                                <Radio value={2}>再修改材料</Radio>
                                <Radio value={3}>驳回修改材料</Radio>
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
                                name="files" 
                                action="/upload.do"
                                defaultFileList={this.state.defaultFileList}
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
