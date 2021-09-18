import React, { Component} from 'react'
import DocumentTitle from 'react-document-title'
import { Button, Table, Tag, Space, Pagination, Menu, Layout, Dropdown, Input, Modal, Form, message,Spin } from 'antd';
import { DownOutlined, PlusCircleOutlined } from '@ant-design/icons';
import request from '../../utils/request'
import ChangeTags from '../../components/ChangeTags'
import store from '../../store'
import './index.less'

const { Header, Footer, Content } = Layout;
const { Search, TextArea } = Input;


export default class index extends Component {
    state = {
        data:[],
        tableHeight:0,
        loadingState:false,
        isCreateProjectVisible:false,
        createLoading:false,
        form:{
          current:1,
          pageSize:10,
          pageSizeOptions:[5,10,15,20],
          showSizeChanger:true,
          total:50
        },
        createForm:{
          name:"",
          grade_range:[],
          subjects:[],
          description:"",
          requirement:"",
          summary:"",
          target:"",
        }
    }

    createFormRef = React.createRef()

    columns = [
        {
          title: '项目名称',
          key: 'name',
          align: 'center',
          width: 210,
          render: record => (
            <Space size="middle">
              <Button type="link" onClick={this.seekProjectManagement.bind(this,record)}>项目名称</Button>
            </Space>
          )
        },
        {
          title: '负责人',
          dataIndex: 'operator',
          key: 'director',
          align: 'center',
          width: 140,
        },
        {
          title: '学科',
          key: 'subject',
          dataIndex: 'subject',
          align: 'center',
          width: 210,
          render: tags => (
            <>
              {/* {tags.map((tag,index)=> {
                let colorList = ['green','geekblue','red']
                return (
                  <Tag color={colorList[index]} key={tag}>
                    {tag.toUpperCase()}
                  </Tag>
                );
              })} */}
              <Tag color="red" key="1">
                学科标签
              </Tag>
            </>
          ),
        },
        {
            title: '学段',
            dataIndex: 'period',
            key: 'period',
            align: 'center',
            width: 80,
        },
        {
            title: '试卷',
            dataIndex: 'paper',
            key: 'paper',
            align: 'center',
            width: 142,
        },
        {
            title: '试题',
            dataIndex: 'questions',
            key: 'questions',
            align: 'center',
            width: 121,
        },
        {
            title: '创建时间',
            dataIndex: 'CreateAt',
            key: 'create-time',
            align: 'center',
            width: 342,
        },
        {
          title:"Title",
          key:'title',
          align: 'center',
          render: (text, record) => (
            <Space size="middle">
              <Button type="link">Invite {record.name}</Button>
              <Button type="link">Delete</Button>
            </Space>
          )
        }
    ]
    pendingTaskMenu = ()=>{
      return (
          <Menu>
              <Menu.Item key="1">任务1</Menu.Item>
              <Menu.Item key="2">任务2</Menu.Item>
              <Menu.Item disabled key="3">任务3</Menu.Item>
              <Menu.Item key="4">任务4</Menu.Item>
          </Menu>
      )
    }
    componentDidMount = ()=>{
      // let data = []
      // for(let i=0;i<5;i++){
      //   data.push({
      //     key: i,
      //     name: 'project'+i,
      //     director:'负责人'+i,
      //     subject: ['学科1', '学科2', '学科3'],
      //     period:'test',
      //     paper:'test',
      //     questions:"test",
      //     createTime:"2021-08-08 02:05:48",
      //   })
      // }
      // this.setState({
      //   data
      // })
      this.getProjectList()
    }
    seekProjectManagement = (state)=>{
      this.props.history.push(`/home/project-management/${state.project_id}/${state.role}`)
    }

    getProjectList = ()=>{
      this.setState({
        loadingState:true
      })
      request({ method:'GET', url:`/proj/user/${store.getState().userInfo.Id}`}).then(res=>{
        this.setState({
          data:res.data,
          loadingState:false
        })
      }).catch(err=>{
        this.setState({
          loadingState:false
        })
      })
    }



    render() {
        return (
            <DocumentTitle title="命题与题库系统-待处理任务">
                <Layout className="pending-tasks-page" data-component="pending-tasks-page">
                  <Header>
                    <span className="title">项目列表/待处理任务</span>
                    <div className="right-box">
                      <Search placeholder="input search text" style={{width:'2.64rem',height:'.32rem'}}/>
                      <Dropdown overlay={this.pendingTaskMenu()}>
                          <span>待处理任务<DownOutlined /></span>
                      </Dropdown>
                      <Button type="primary" icon={<PlusCircleOutlined />} onClick={()=>{
                        this.setState({
                          isCreateProjectVisible:true
                        })
                      }}>添加项目</Button>
                    </div>
                  </Header>
                  <Content>
                    <Table 
                      loading={this.state.loadingState}
                      columns={this.columns} 
                      dataSource={this.state.data} 
                      rowSelection={{}} 
                      size="small" 
                      rowKey="Id"
                      pagination={false}
                      scroll={{ y: 'calc(100vh - 2.2rem)'}}
                    />
                  </Content>
                  <Footer>
                    <Pagination
                      size="small"
                      total={85}
                      showSizeChanger
                      showQuickJumper
                    />
                    <Modal 
                      title="创建项目" 
                      visible={this.state.isCreateProjectVisible} 
                      cancelText="取消创建"
                      okText="创建项目"
                      closable={!this.state.createLoading}
                      confirmLoading={this.state.createLoading}
                      maskClosable={!this.state.createLoading}
                      keyboard={!this.state.createLoading}
                      onOk={()=>{
                        const form = this.createFormRef.current
                        form.validateFields().then(data=>{
                          this.setState({
                            createLoading:true
                          })
                          request({ method:'POST', url:"/proj/template",data}).then(res=>{
                            message.success("项目创建成功！");
                            this.setState({
                              createLoading:false,
                              isCreateProjectVisible:false
                            })
                            this.createFormRef.current.resetFields()
                            this.getProjectList()
                          }).catch(err=>{
                            this.setState({
                              createLoading:false
                            })
                            message.error(err.message||"未知错误")
                          })
                        }).catch(err=>{
                          message.warning("请按要求填写表单项！")
                        })
                      }} 
                      onCancel={()=>{
                        if(this.state.createLoading){
                          message.error("项目创建中，不可阻断！")
                        }else{
                          this.createFormRef.current.resetFields()
                          this.setState({
                            isCreateProjectVisible:false
                          })
                        }
                      }}
                    >
                      <Spin spinning={this.state.createLoading} tip="项目创建中，请等待！">
                        <Form labelCol={{span: 4 }}  wrapperCol={{span: 20}} ref={this.createFormRef} name="createForm"  initialValues={this.state.createForm}>
                          <Form.Item name="name" label="项目名称" rules={[{ required: true,message: '项目名称不能为空！' }]}>
                            <Input placeholder="请输入项目名称" />
                          </Form.Item>
                          <Form.Item name="grade_range" label="年级范围" rules={[{ required: true,message:"请至少创建一个年级" }]}>
                            <ChangeTags onChange={grade_range=>{
                              let createForm = Object.assign(this.state.createForm,{
                                grade_range
                              })
                              this.setState({
                                createForm
                              })
                            }}></ChangeTags>
                          </Form.Item>
                          <Form.Item name="subjects" label="涉及学科" rules={[{ required: true,message:"请至少创建一个学科！" }]}>
                            <ChangeTags onChange={subjects=>{
                              let createForm = Object.assign(this.state.createForm,{
                                subjects
                              })
                              this.setState({
                                createForm
                              })
                            }}></ChangeTags>
                          </Form.Item>
                          <Form.Item name="target" label="目标说明" rules={[{ required: true,message:"目标说明不能为空！" }]}>
                            <TextArea
                              placeholder="请输入项目目标说明，若无，请填写无"
                              autoSize={{ minRows: 3, maxRows: 5 }}
                            />
                          </Form.Item>
                          <Form.Item name="summary" label="项目摘要" rules={[{ required: true,message:"项目摘要不能为空！" }]}>
                            <TextArea
                              placeholder="请输入项目摘要，若无，请填写无"
                              autoSize={{ minRows: 3, maxRows: 5 }}
                            />
                          </Form.Item>
                          <Form.Item name="description" label="项目描述" rules={[{ required: true,message:"项目描述不能为空！" }]}>
                            <TextArea
                              placeholder="请输入项目描述，若无，请填写无"
                              autoSize={{ minRows: 3, maxRows: 5 }}
                            />
                          </Form.Item>
                          <Form.Item name="requirement" label="项目要求" rules={[{ required: true,message:"项目要求不能为空！" }]}>
                            <TextArea
                              placeholder="请输入项目要求，若无，请填写无"
                              autoSize={{ minRows: 3, maxRows: 5 }}
                            />
                          </Form.Item>
                        </Form>
                      </Spin>
                      
                    </Modal>
                  </Footer>
                </Layout>
                
            </DocumentTitle>
        )
    }
}
