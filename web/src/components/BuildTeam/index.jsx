import React, { Component } from 'react'
import { Button,message,Table,Tag,Space,Modal,Select,Input,Form } from 'antd'
import {withRouter} from 'react-router-dom'
import ModalCard from '../ModulaCard'
import request from '../../utils/request'
import store from '../../store'
import './index.less'

import { WarningTwoTone } from '@ant-design/icons'

const { Search,TextArea } = Input
const { confirm } = Modal
const { Option } = Select

class index extends Component {

    state={
        loadingState:false,
        roleChangeForm:{
            show:false,
            assignment_id:"",
            new_role:"",
            updateLoading:false
        },
        addMemberForm:{
            show:false,
            loadingState:false,
            role:1,
            user_id:""
        },
        emailForm:{
            show:false,
            loadingState:false,
            message:"",
            send_time:"",
            sender:store.getState().userInfo.name,
            destination:[],
            subject:""
        },
        memberList:[]
    }

    
    columns = [{
        title: '姓名',
        align:"center",
        render: (text,record) =>(
            <span>{record.info.name}</span>
        )
    },{
        title: '年龄',
        align:"center",
        render: (text,record) =>(
            <span>{record.info.age}</span>
        )
    },{
        title: '性别',
        align:"center",
        render: (text,record) =>(
            <span>{record.info.gender?"男":"女"}</span>
        )
    },{
        title: '专业',
        align:"center",
        render: (text,record) =>(
            <Tag color="processing">{record.info.degree}</Tag>
        )
    },{
        title: '职位',
        align:"center",
        render: (text,record) =>(
            <span>{record.info.position}</span>
        )
    },{
        title: '邮箱',
        align:"center",
        render: (text,record) =>(
            <Button type="link" onClick={()=>{
                this.setState({
                    emailForm:Object.assign(this.state.emailForm,{destination:[record.info.email],send_time:new Date().getTime().toString(),show:true})
                })
            }}>
                {record.info.email}
            </Button>
        )
    },{
        title: '组织',
        align:"center",
        render: (text,record) =>(
            <span>{record.info.organization}</span>
        )
    },{
        title: '位置',
        align:"center",
        render: (text,record) =>(
            <span>{record.info.location}</span>
        )
    },{
        title: '电话',
        align:"center",
        render: (text,record) =>(
            <span>{record.info.phone}</span>
        )
    },{
        title:"项目角色",
        align:"center",
        render:(text,record)=>{
            let roleList=["管理员","专家","学科助理","教师","外审专家"]
            return (
                <Tag color="green">{roleList[record.role-1]}</Tag>
            )
        }
    },{
        title: '状态',
        align:"center",
        render: (text,record) =>(
            <>
                {
                    record.info.is_confirmed?(
                        <Tag color="#87d068">已确认 √</Tag>
                    ):(
                        <Tag color="#f50">未确认 ×</Tag>
                    )
                }
            </>
        )
    },{
        title:"管理",
        align:"center",
        render: (text, record) => (
            <Space size="middle">
              <Button type="link" onClick={()=>{
                  this.setState({
                      roleChangeForm:Object.assign(this.state.roleChangeForm,{
                        assignment_id:record.Id,
                        new_role:record.role,
                        show:true
                      })
                  })
              }}>更改角色</Button>
              <Button type="link" danger onClick={()=>{
                  confirm({
                    icon: <WarningTwoTone />,
                    content: "该删除操作不可逆，是否继续？",
                    okText:"确认移除",
                    cancelText:"取消移除",
                    onOk:()=>{
                      this.setState({
                          loadingState:true
                      })
                      request({
                          url:`http://49.232.73.36:8081/review/proj/assign/${record.uuid}`,
                          method:"DELETE"
                      }).then(res=>{
                        this.setState({
                            loadingState:false
                        })
                        message.success("删除成功！");
                        this.getProjectMember()
                      }).catch(err=>{
                          message.error(err.message||"请求错误！");
                          this.setState({
                            loadingState:false
                            })
                      })
                    },
                    onCancel() {
                      message.info("已取消移除");
                    },
                  });
              }}>移除人员</Button>
            </Space>
        )
    }]

    emailFormRef = React.createRef()

    componentDidMount(){
        this.getProjectMember()
    }

    getProjectMember=()=>{
        this.setState({
            loadingState:true
        })
        let memberList = []
        request({
            url:`http://49.232.73.36:8081/review/proj/assign/${this.props.match.params.project_id}`,
            method:"GET"
        }).then(res=>{
            memberList = [...res.data.admins.map(item=>{
                item.roleName = "admin"
                return item;
            }),...res.data.assistants.map(item=>{
                item.roleName = "assistant"
                return item;
            }),...res.data.experts.map(item=>{
                item.roleName = "expert"
                return item;
            }),...res.data.out_experts.map(item=>{
                item.roleName = "out_expert"
                return item;
            }),...res.data.teachers.map(item=>{
                item.roleName = "teacher"
                return item;
            })]
            let id_list = memberList.map(item=>item.user_id)
            request({
                url:"http://49.232.73.36:8081/review/query/user",
                method:"POST",
                data:{id_list}
            }).then(res=>{
                let userInfo_list = Object.values(res.data);
                memberList = memberList.map((item,index)=>{
                    item.info = userInfo_list[index].profile;
                    return item;
                })
                // let memberData = Object.assign(this.state.memberData,{
                //     admins:memberList.filter(item=>item.roleName==="admin"),
                //     assistants:memberList.filter(item=>item.roleName==="assistant"),
                //     experts:memberList.filter(item=>item.roleName==="expert"),
                //     out_experts:memberList.filter(item=>item.roleName==="out_expert"),
                //     teachers:memberList.filter(item=>item.roleName==="teacher"),
                // })
                this.setState({
                    memberList,
                    loadingState:false
                })
            }).catch(err=>{
                this.setState({
                    loadingState:false
                })
                message.error(err.message||"请求错误！")
            })
        }).catch(err=>{
            this.setState({
                loadingState:false
            })
            message.error(err.message||"请求错误！")
        })
    }

    render() {
        return (
            <div className="build-team-page" data-component="build-team-page">
                <ModalCard 
                    title="项目成员"
                    right={(
                        <Button type="primary" size="small" onClick={()=>{
                            this.setState({
                                addMemberForm:Object.assign(this.state.addMemberForm,{show:true})
                            })
                        }}>添加成员</Button>
                    )}
                >
                    <div className="member-list">
                        <Table 
                            key="admins"
                            columns={this.columns}
                            rowKey="Id"
                            pagination={false}
                            dataSource={this.state.memberList}
                            size="small"
                            loading={this.state.loadingState}
                        />
                    </div>
                </ModalCard>
                <Modal 
                    title="角色分配修改" 
                    visible={this.state.roleChangeForm.show} 
                    okText="确认修改"
                    cancelText="取消修改"
                    closable={!this.state.roleChangeForm.updateLoading}
                    keyboard={!this.state.roleChangeForm.updateLoading}
                    maskClosable={!this.state.roleChangeForm.updateLoading}
                    confirmLoading={this.state.roleChangeForm.updateLoading}
                    onOk={()=>{
                        this.setState({
                            roleChangeForm:Object.assign(this.state.roleChangeForm,{updateLoading:true})
                        })
                        request({
                            url:"http://49.232.73.36:8081/review/proj/assign",
                            method:"PATCH",
                            data:{
                                assignment_id: this.state.roleChangeForm.assignment_id,
                                new_role: this.state.roleChangeForm.new_role,
                            }
                        }).then(res=>{
                            console.log(res)
                            this.setState({
                                roleChangeForm:Object.assign(this.state.roleChangeForm,{updateLoading:false,show:false})
                            })
                            this.getProjectMember()
                            message.success("修改成功！");
                        }).catch(err=>{
                            message.error(err.message||"请求错误！");
                            this.setState({
                                roleChangeForm:Object.assign(this.state.roleChangeForm,{updateLoading:false})
                            })
                        })
                    }} 
                    onCancel={()=>{
                        if(this.state.roleChangeForm.updateLoading){
                            message.error("修改中，操作不可中断！")
                        }else{
                            this.setState({
                                roleChangeForm:Object.assign(this.state.roleChangeForm,{show:false})
                            })
                        }
                    }}
                >
                    <label style={{lineHeight:".6rem"}}>新的角色：</label>
                    <Select value={this.state.roleChangeForm.new_role}  style={{width:"100%"}} onChange={(e)=>{
                        this.setState({
                            roleChangeForm:Object.assign(this.state.roleChangeForm,{new_role:e})
                        })
                    }}>
                        <Option value={1}>管理员</Option>
                        <Option value={2}>专家</Option>
                        <Option value={3}>学科助理</Option>
                        <Option value={4}>教师</Option>
                        <Option value={5}>外审人员</Option>
                    </Select>
                </Modal>
                <Modal 
                    title="添加成员" 
                    visible={this.state.addMemberForm.show} 
                    okText="确认添加"
                    cancelText="取消添加"
                    closable={!this.state.addMemberForm.loadingState}
                    keyboard={!this.state.addMemberForm.loadingState}
                    maskClosable={!this.state.addMemberForm.loadingState}
                    confirmLoading={this.state.addMemberForm.loadingState}
                    onOk={()=>{

                    }} 
                    onCancel={()=>{
                        if(this.state.addMemberForm.loadingState){
                            message.error("添加中，操作不可中断！")
                        }else{
                            this.setState({
                                addMemberForm:Object.assign(this.state.addMemberForm,{show:false})
                            })
                        }
                    }}
                >
                    <label style={{lineHeight:".6rem"}}>查找用户：</label>
                    <Search placeholder="请输入被添加用户账号" onSearch={()=>{

                    }} enterButton />
                    {
                        this.state.addMemberForm.user_id===""?"":(
                            <>
                                <label style={{lineHeight:".6rem"}}>角色分配：</label>
                                <Select value={this.state.addMemberForm.role}  style={{width:"100%"}} onChange={(e)=>{
                                    this.setState({
                                        addMemberForm:Object.assign(this.state.addMemberForm,{role:e})
                                    })
                                }}>
                                    <Option value={1}>管理员</Option>
                                    <Option value={2}>专家</Option>
                                    <Option value={3}>学科助理</Option>
                                    <Option value={4}>教师</Option>
                                    <Option value={5}>外审人员</Option>
                                </Select>
                            </>
                        )
                    }
                </Modal>
                <Modal
                   title="发送邮件" 
                   visible={this.state.emailForm.show} 
                   okText="确认发送"
                   cancelText="取消发送"
                   closable={!this.state.emailForm.loadingState}
                   keyboard={!this.state.emailForm.loadingState}
                   maskClosable={!this.state.emailForm.loadingState}
                   confirmLoading={this.state.emailForm.loadingState}
                   onOk={()=>{
                        this.emailFormRef.current.validateFields().then(formData=>{
                            let data = {
                                body:{
                                    message:formData.message,
                                    send_time:this.state.emailForm.send_time,
                                    sender:this.state.emailForm.sender
                                },
                                destination:this.state.emailForm.destination,
                                subject:formData.subject
                            }
                            this.setState({
                                emailForm:Object.assign(this.state.emailForm,{loadingState:true})
                            })
                            request({
                                url:"http://49.232.73.36:8081/review/noticer/email",
                                method:"POST",
                                data
                            }).then(res=>{
                                this.setState({
                                    emailForm:Object.assign(this.state.emailForm,{loadingState:false,show:false})
                                })
                                this.emailFormRef.current.resetFields()
                                message.success("发送成功");
                            }).catch(err=>{
                                this.setState({
                                    emailForm:Object.assign(this.state.emailForm,{loadingState:false})
                                })
                                message.error(err.message||"请求错误")
                            })
                        }).catch(err=>{
                            message.warning("请正确填写邮件内容")
                        })
                   }} 
                   onCancel={()=>{
                       if(this.state.emailForm.loadingState){
                           message.error("发送中，操作不可中断！")
                       }else{
                            this.emailFormRef.current.resetFields()
                           this.setState({
                            emailForm:Object.assign(this.state.emailForm,{show:false})
                           })
                       }
                   }} 
                >
                    <Form
                        name="emailForm"
                        labelCol={{ span: 4 }}
                        wrapperCol={{ span: 20 }}
                        initialValues={this.state.emailForm}
                        autoComplete="off"
                        ref = {this.emailFormRef}
                        >
                            <Form.Item
                                label="邮件主题"
                                name="subject"
                                rules={[{ required: true, message: '请输入邮件主题' }]}
                            >
                                <Input placeholder="请输入邮件主题"/>
                            </Form.Item>

                            <Form.Item
                                label="邮件内容"
                                name="message"
                                rules={[{ required: true, message: '请输入邮件内容' }]}
                            >
                                <TextArea placeholder="请输入邮件内容" autoSize={{ minRows: 2, maxRows: 6 }}/>
                            </Form.Item>
                        </Form>
                </Modal>
            </div>
        )
    }
}

export default withRouter(index)
