//封装的材料仓库组件
import React, { Component } from 'react'
import { Input, Pagination, Empty, Upload, message } from 'antd'
import ModulaCard from '../ModulaCard'
import UpLoadModal from '../UpLoadModal'
import { PlusOutlined, FileTextFilled, FileExcelFilled, FileZipFilled, FileMarkdownFilled, FilePptFilled } from '@ant-design/icons'
import request  from '../../utils/request'
import './index.less'

const { Search } = Input

export default class inedx extends Component {

    state = {
        fileType:[<FileTextFilled />,<FileExcelFilled/>,<FileZipFilled/>,<FileMarkdownFilled />,<FilePptFilled />],
        fileList:[],
        upLoadVisible:false,
        getFileLoading:false
    }

    downLoadFile = (file_id)=>{
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

    componentDidMount(){
        this.getFileList()
    }

    getFileList = ()=>{
        this.setState({
            getFileLoading:true
        })
        request({
            url:`/proj/detailed/${this.props.projectId}`,
            method:"GET"
        }).then(res=>{
            request({
                url:"/query/file",
                method:"POST",
                data:{
                    id_list:res.data.materials.files
                }
            }).then(res=>{
                this.setState({
                    fileList:Object.values(res.data),
                    getFileLoading:false
                })
            }).catch(err=>{
                this.setState({
                    getFileLoading:false
                })
            })
        }).catch(err=>{
            this.setState({
                getFileLoading:false
            })
        })
    }


    fileViewLoader = ()=>{
        if(this.props.role==="2"&&this.state.fileList.length===0){
            return (
                <div className="empty-state-box">
                    <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />
                </div>
            )
        }else if((this.props.role==="2"&&this.state.fileList.length!==0)||this.props.role==="3"||this.props.role==="4"){
            return this.state.fileList.map(item=>(
                <div className="file-item" key={item.Id} onClick={this.downLoadFile.bind(this,item.uuid)}>
                    <div className="icon">
                        <FileTextFilled />
                    </div>
                    <div className="name">
                        <span>{item.name}</span>
                    </div>
                </div>
            ))
        }
        return (
            <></>
        )
    }

    render() {
        return (
            <ModulaCard title="材料仓库" right={<Search placeholder="input search text" size="small" style={{ width: 200 }}/> }>
                <div className="material-warehouse-box" data-component="material-warehouse-box">
                    <div className="container">
                        {
                            this.props.role==="3"||(this.props.role==="4"&&this.props.stepName!=="测试框架与论证报告")?(
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
                                        <div className="file-load-btn" onClick={()=>{
                                            this.setState({
                                                upLoadVisible:true
                                            })
                                        }}>
                                            <PlusOutlined />
                                            <div style={{ marginTop: 8,width:'100%' }}>上传</div>
                                        </div>
                                    </Upload>
                                </div>
                            ):(
                                <></>
                            )
                        }
                        {this.fileViewLoader()}
                        <UpLoadModal
                            show={this.state.upLoadVisible}
                            projectId={this.props.projectId}
                            stepId = {this.props.stepId}
                            onClose={()=>{
                                this.setState({
                                    upLoadVisible:false
                                })
                            }}
                            onUpdate = {()=>{
                                this.getFileList()
                            }}
                        ></UpLoadModal>
                    </div>
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
