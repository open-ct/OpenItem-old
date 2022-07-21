import React, { Component } from 'react'
import {Modal,Form,Select,message,Tag} from 'antd'
import request from '../../utils/request'
import store from '../../store'
import { withRouter } from 'react-router-dom';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'

const {Option} = Select
const {confirm} = Modal

class index extends Component {

    state = {
        form:{
            project:null,
            subject:null,
            content:[],
            ability:[],
            type:null
        },
        projectLoading:false,
        projectList:[]
    }

    contentOption = [{
        value:"数与代数"
    },{
        value:"历史认知"
    },{
        value:"积累与运用"
    }]

    abilityOption = [{
        value:"听力"
    },{
        value:"了解"
    },{
        value:"认识与知识"
    },{
        value:"古诗文积累"
    }]

    formRef = React.createRef()

    titleList=[{
        type:"create",
        value:"上传试题信息"
    },{
        type:"update",
        value:"编辑试题信息"
    },{
        type:"edit",
        value:"编辑试题信息"
    },{
        type:"create-paper",
        value:"上传试卷信息"
    },{
        type:"update-paper",
        value:"更新试卷信息"
    }]

    getProjectList = ()=>{
        this.setState({
            projectLoading:true
        })
        request({
            url:`http://49.232.73.36:8081/review/proj/user/${store.getState().userInfo.Id}`,
            method:"GET"
        }).then(res=>{
            let id_list = []
            for(let i=0;i<res.data.length;i++){
                id_list.push(res.data[i].project_id)
            }
            request({
                url:"http://49.232.73.36:8081/review/query/proj",
                method:"POST",
                data:{id_list}
            }).then(res=>{
                this.setState({
                    projectList:Object.values(res.data),
                    projectLoading:false
                })
            }).catch(err=>{
                message.error(err.message||"加载失败，请重试")
                this.setState({
                    projectLoading:false
                })
            })
        }).catch(err=>{
            message.error(err.message||"加载失败，请重试")
            this.setState({
                projectLoading:false
            })
        })
    }

    tagRender = (props)=>{
        let colorList = ['gold','lime','green','cyan']
        const { label, closable, onClose } = props;
        const onPreventMouseDown = event => {
          event.preventDefault();
          event.stopPropagation();
        };
        return (
          <Tag
            color={colorList[Math.floor(Math.random() * (3))]}
            onMouseDown={onPreventMouseDown}
            closable={closable}
            onClose={onClose}
            style={{ marginRight: 3 }}
          >
            {label}
          </Tag>
        );
    }


    render() {
        return (
            <Modal 
                title={this.titleList[this.titleList.findIndex(item=>item.type===this.props.type)].value}
                cancelText="取消"
                okText={this.props.type==="create"||this.props.type==="create-paper"?"下一步":"保存"}
                visible={this.props.show} 
                onOk={()=>{
                    this.formRef.current.validateFields().then(data=>{
                        if(this.props.type==="create"){
                            this.props.history.push(`/home/proposition-paper/upload-questions/${data.project}/${data.subject}/${data.ability}/${data.content}/${data.type}`)
                        }else if(this.props.type==="update"){
                            let that = this
                            confirm({
                                icon: <ExclamationCircleOutlined />,
                                content: "此操作将覆盖当前撰写的试题，是否继续？",
                                onOk() {
                                    that.props.onClose()
                                    that.props.history.push(`/home/proposition-paper/upload-questions/${data.project}/${data.subject}/${data.ability}/${data.content}/${data.type}`)
                                },
                                onCancel() {
                                    that.props.onClose()
                                },
                            });
                        }else if(this.props.type==="edit"){
                            console.log("开始保存")
                            this.props.onClose()
                        }else if(this.props.type==="create-paper"){
                            this.props.history.push(`/home/proposition-paper/create-paper/${data.project}/${data.subject}/${data.ability}/${data.content}/${data.type}`)
                        }else if(this.props.type==="update-paper"){
                            let that = this
                            confirm({
                                icon: <ExclamationCircleOutlined />,
                                content: "此操作将覆盖当前撰写的试卷，是否继续？",
                                onOk() {
                                    that.props.onClose()
                                    that.props.history.push(`/home/proposition-paper/create-paper/${data.project}/${data.subject}/${data.ability}/${data.content}/${data.type}`)
                                },
                                onCancel() {
                                    that.props.onClose()
                                },
                            });
                        }else{
                            message.error("组件参数异常，失败！")
                            this.props.onClose()
                        }
                    }).catch(err=>{    
                        message.warning("请按要求选择选项")
                    })
                    
                }} 
                onCancel={()=>{
                    this.formRef.current.resetFields()
                    this.props.onClose()
                }}
            >
                <div className="upLoad-question-title">
                    <span>试题编号编号:absnahghj（自动生成）</span>
                </div>
                <Form 
                    labelCol={{ span: 5 }}
                    wrapperCol={{ span: 19 }}
                    labelAlign="left"
                    ref={this.formRef}
                    initialValues={this.state.form}
                >
                    <Form.Item
                        name="project"
                        label="项目"
                        rules={[{ required: true, message: '请选择项目名称' }]}
                    >
                        <Select placeholder="选择项目名称" loading={this.state.projectLoading} onFocus={this.getProjectList} onSelect={(e)=>{
                            let from  = Object.assign(this.state.form,{project:e})
                            this.setState({
                                from
                            })
                        }}>
                            {
                                this.state.projectLoading?(
                                    <></>
                                ):this.state.projectList.map(item=>(
                                    <Option key={item.Id} value={item.uuid}>{item.basic_info.name}</Option>
                                ))
                            }
                        </Select>
                    </Form.Item> 
                    <Form.Item
                        name="subject"
                        label="学科"
                        rules={[{ required: true, message: '请选择学科' }]}
                    >
                        <Select placeholder="选择学科" loading={this.state.projectLoading} onFocus={()=>{
                            if(!this.state.form.project){
                                message.warning("请先选择项目")
                            }
                        }} onSelect={(e)=>{
                            let form = Object.assign(this.state.form,{subject:e})
                            this.setState({
                                form
                            })
                        }}>
                            {
                                !this.state.form.project?(
                                    <></>
                                ):this.state.projectList[this.state.projectList.findIndex(item=>item.uuid===this.state.form.project)].basic_info.subjects.map((item,index)=>(
                                    <Option key={index} value={item}>{item}</Option>
                                ))
                            }
                        </Select>
                    </Form.Item>
                    <Form.Item
                        name="content"
                        label="内容纬度"
                        rules={[{ required: true, message: '请选择内容纬度' }]}
                    >
                        <Select
                            mode="multiple"
                            showArrow
                            tagRender={this.tagRender}
                            style={{ width: '100%' }}
                            options={this.contentOption}
                            onChange={(e)=>{
                                let form = Object.assign(this.state.form,{content:e})
                                this.setState({
                                    form
                                })
                            }}
                        />
                    </Form.Item>
                    <Form.Item
                        name="ability"
                        label="能力维度"
                        rules={[{ required: true, message: '请选择能力维度' }]}
                    >
                        <Select
                            mode="multiple"
                            showArrow
                            tagRender={this.tagRender}
                            style={{ width: '100%' }}
                            options={this.abilityOption}
                            onChange={(e)=>{
                                let form = Object.assign(this.state.form,{ability:e})
                                this.setState({
                                    form
                                })
                            }}
                        />
                    </Form.Item>
                    <Form.Item
                        name="type"
                        label="题型"
                        rules={[{ required: true, message: '请选择题型' }]}
                    >
                        <Select placeholder="选择题型" onSelect={(e)=>{
                            let form = Object.assign(this.state.form,{type:e})
                            this.setState({
                                form
                            })
                        }}>
                            <Option value="1">选择题</Option>
                            <Option value="2">填空题</Option>
                            <Option value="3">简答题</Option>
                        </Select>
                    </Form.Item>
                </Form>
            </Modal>
        )
    }
}

export default withRouter(index)