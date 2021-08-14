import React, { Component} from 'react'
import DocumentTitle from 'react-document-title'
import { Button, Table, Tag, Space, Pagination, Menu, Layout, Dropdown, Input, Spin } from 'antd';
import { DownOutlined, PlusCircleOutlined } from '@ant-design/icons';
import request from '../../utils/request'
import store from '../../store'
import './index.less'

const { Header, Footer, Content } = Layout;
const { Search } = Input;

export default class index extends Component {
    state = {
        data:[],
        tableHeight:0,
        form:{
          current:1,
          pageSize:10,
          pageSizeOptions:[5,10,15,20],
          showSizeChanger:true,
          total:50
        },
        loadingState:{
          show:false,
          title:'加载中...'
        }
    }
    columns = [
        {
          title: '项目名称',
          key: 'name',
          align: 'center',
          width: 210,
          render: (text, record) => (
            <Space size="middle">
              <Button type="link" onClick={this.seekProjectManagement}>{record.name}</Button>
            </Space>
          )
        },
        {
          title: '负责人',
          dataIndex: 'director',
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
              {tags.map((tag,index)=> {
                let colorList = ['green','geekblue','red']
                return (
                  <Tag color={colorList[index]} key={tag}>
                    {tag.toUpperCase()}
                  </Tag>
                );
              })}
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
            dataIndex: 'createTime',
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
      let data = []
      for(let i=0;i<5;i++){
        data.push({
          key: i,
          name: 'project'+i,
          director:'负责人'+i,
          subject: ['学科1', '学科2', '学科3'],
          period:'test',
          paper:'test',
          questions:"test",
          createTime:"2021-08-08 02:05:48",
        })
      }
      this.setState({
        data
      })
    }
    seekProjectManagement = ()=>{
      this.props.history.push('/home/project-management')
    }

    getProjectList = ()=>{
      this.setState({
        loadingState:{
          show:true,
          title:'数据加载中...'
        }
      })
      request({ method:'GET', url:`/proj/userAssign?user_id=${store.getState().userInfo.user_id}`}).then(res=>{
        console.log('res',res)
        this.setState({
          loadingState:{
            show:false,
            title:''
          }
        })
      }).catch(err=>{
        console.log('err',err)
        this.setState({
          loadingState:{
            show:false,
            title:''
          }
        })
      })
    }

    // componentDidMount = ()=>{
    //   this.getProjectList()
    // }

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
                      {/* <Dropdown overlay={this.pendingTaskMenu()}>待处理任务<DownOutlined /></Dropdown> */}
                      <Button type="primary" icon={<PlusCircleOutlined />}>添加项目</Button>
                    </div>
                  </Header>
                  <Content>
                    <Spin spinning={this.state.loadingState.show} size="large" tip={this.state.loadingState.title}>
                      <Table 
                        columns={this.columns} 
                        dataSource={this.state.data} 
                        rowSelection={{}} 
                        size="small" 
                        pagination={false}
                        scroll={{ y: 'calc(100vh - 2.2rem)'}}
                      />
                    </Spin>
                  </Content>
                  <Footer>
                    <Pagination
                      total={85}
                      showSizeChanger
                      showQuickJumper
                    />
                  </Footer>
                </Layout>
            </DocumentTitle>
        )
    }
}
