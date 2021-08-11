//封装的材料仓库组件
import React, { Component } from 'react'
import { Input, Pagination, Empty, Upload, message } from 'antd'
import ModulaCard from '../ModulaCard'
import UpLoadModal from '../UpLoadModal'
import { PlusOutlined, FileTextFilled, FileExcelFilled, FileZipFilled, FileMarkdownFilled, FilePptFilled } from '@ant-design/icons'
import './index.less'

const { Search } = Input

export default class inedx extends Component {

    state = {
        fileType:[<FileTextFilled />,<FileExcelFilled/>,<FileZipFilled/>,<FileMarkdownFilled />,<FilePptFilled />],
        upLoadVisible:false
    }

    downLoadFile = ()=>{
        if(this.props.mode.includes('downLoad')){
            message.success('下载成功')
        }else{
            message.error('下载失败：您暂无权限下载')
        }
    }

    loadContainer = ()=>{
        return (
            <div className="container">
                {this.props.mode.includes('upLoadOnly')?(
                    <div className="upload-only-box" onClick={()=>{
                        this.setState({
                            upLoadVisible:true
                        })
                    }}>
                        <Upload
                            name="avatar"
                            listType="picture-card"
                            className="avatar-uploader"
                            showUploadList={false}
                            action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
                            openFileDialogOnClick={false}
                        >
                            <div className="btn-box" onClick={()=>{
                                this.setState({
                                    upLoadMaterialVisible:true
                                })
                            }}>
                                <PlusOutlined />
                                <div style={{ marginTop: 8,width:'100%' }}>Upload</div>
                            </div>
                        </Upload>
                    </div>
                ):(
                    <div className="data-show-box">
                        {this.props.fileList.length===0&&!this.props.mode.includes('upLoad')?(
                            <div className="empty-state-box">
                                <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />
                            </div>
                        ):''}
                        {this.props.mode.includes('upLoad')?(
                            <div className="upload-download-box" onClick={()=>{
                                this.setState({
                                    upLoadVisible:true
                                })
                            }}>
                                <Upload
                                    name="avatar"
                                    listType="picture-card"
                                    showUploadList={false}
                                    action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
                                    openFileDialogOnClick={false}
                                >
                                    <div className="file-load-btn">
                                        <PlusOutlined />
                                        <div style={{ marginTop: 8,width:'100%' }}>Upload</div>
                                    </div>
                                </Upload>
                            </div>
                        ):''}
                        {this.props.fileList.map(item=>(
                            <div className="file-item" onClick={this.downLoadFile}>
                                {this.state.fileType[item.fileType]}
                                <div className="key">
                                    {item.id}
                                </div>
                            </div>
                        ))}
                    </div>
                )}
                <UpLoadModal show={this.state.upLoadVisible} onLoad={()=>{
                    this.setState({
                        upLoadVisible:false
                    })
                    message.success('上传成功')
                }} />
            </div>
        )
    }

    render() {
        return (
            <ModulaCard title="材料仓库" right={<Search placeholder="input search text" size="small" style={{ width: 200 }}/> }>
                <div className="material-warehouse-box" data-component="material-warehouse-box">
                    {this.loadContainer()}
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
            </ModulaCard>
        )
    }
}
