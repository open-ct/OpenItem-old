import React, { Component} from 'react'
import DocumentTitle from 'react-document-title'
import { Button, Table, Tag, Space, Pagination, Menu, Dropdown,Input } from 'antd';
import { PlusCircleOutlined,DownOutlined } from '@ant-design/icons';
import './index.less'

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
        }
    }
    columns = [
        {
          title: '项目名称',
          dataIndex: 'name',
          key: 'name',
          align: 'center',
          width: 150,
        },
        {
          title: '负责人',
          dataIndex: 'director',
          key: 'director',
          align: 'center',
          width: 150,
        },
        {
          title: '学科',
          key: 'subject',
          dataIndex: 'subject',
          align: 'center',
          width: 250,
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
            width: 80,
        },
        {
            title: '试题',
            dataIndex: 'questions',
            key: 'questions',
            align: 'center',
            width: 80,
        },
        {
            title: '创建时间',
            dataIndex: 'createTime',
            key: 'create-time',
            align: 'center',
            width: 200,
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
              <Menu.Item key="1">
                <a target="_blank" rel="noopener noreferrer" href="https://www.antgroup.com">
                  1st menu item
                </a>
              </Menu.Item>
              <Menu.Item icon={<DownOutlined />} disabled key="2">
                <a target="_blank" rel="noopener noreferrer" href="https://www.aliyun.com">
                  2nd menu item (disabled)
                </a>
              </Menu.Item>
              <Menu.Item disabled key="3">
                <a target="_blank" rel="noopener noreferrer" href="https://www.luohanacademy.com">
                  3rd menu item (disabled)
                </a>
              </Menu.Item>
              <Menu.Item danger key="4">a danger item</Menu.Item>
          </Menu>
      )
    }
    componentDidMount = ()=>{
      let data = []
      for(let i=0;i<20;i++){
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
    render() {
        return (
            <DocumentTitle title="命题与题库系统-待处理任务">
                <div className="pending-tasks-page" data-component="pending-tasks-page">
                    <div className="header">
                        <span className="title">项目列表/待处理任务</span>
                        <div className="right-box">
                            <Input.Search placeholder="input search text" style={{ width: 200 }} />
                            <Dropdown overlay={this.pendingTaskMenu()}>
                              <Button type="link">
                                待处理任务 <DownOutlined />
                              </Button>
                            </Dropdown>
                            <Button type="primary" size="small" icon={<PlusCircleOutlined />}>添加项目</Button>
                        </div>
                    </div>
                    <div className="container">
                        <Table 
                            columns={this.columns} 
                            dataSource={this.state.data} 
                            rowSelection={{}} 
                            size="small" 
                            pagination={false}
                            scroll={{ y: 'calc(100vh - 224px)'}}
                        />
                    </div>
                    <div className="footer">
                        <Pagination
                          total={85}
                          showSizeChanger
                          showQuickJumper
                        />
                    </div>
                </div>
            </DocumentTitle>
        )
    }
}
