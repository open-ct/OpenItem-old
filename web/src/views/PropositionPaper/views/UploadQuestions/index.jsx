import React, { Component } from 'react'
import {PageHeader,Descriptions,Button,Layout,Menu,Radio,Space,Input,Form,Select,Checkbox,Modal} from 'antd'
import {PictureOutlined,UpCircleOutlined,DownCircleOutlined,PlusCircleOutlined,MinusCircleOutlined} from '@ant-design/icons'

import './index.less'

const { Search } = Input;
const { Option } = Select
const { Sider, Content,Footer } = Layout;

export default class index extends Component {

    state = {
        difficultyValue:4,
        editQuestionVisible:false
    }

    render() {
        return (
            <div className="upLoad-question-page" data-component="upLoad-question-page">
                <PageHeader
                    ghost={false}
                    onBack={() => this.props.history.push("home/proposition-paper/home")}
                    title="命题组卷"
                    subTitle="上传试题"
                    extra={[
                        <Button key="1" onClick={()=>{
                            this.setState({
                                editQuestionVisible:true
                            })
                        }}>编辑内容</Button>,
                    ]}
                    >
                    <Descriptions size="small" column={3}>
                        <Descriptions.Item label="创建时间">2017-01-10</Descriptions.Item>
                        <Descriptions.Item label="项目">项目名称</Descriptions.Item>
                        <Descriptions.Item label="学科">数学</Descriptions.Item>
                        <Descriptions.Item label="内容纬度">纬度1；纬度2</Descriptions.Item>
                        <Descriptions.Item label="能力纬度">纬度1；纬度2</Descriptions.Item>
                    </Descriptions>
                </PageHeader>
                <div className="main">
                    <Layout className="container">
                        <Sider theme="light" width="2.4rem" style={{backgroundColor:"#FAFAFA"}}>
                            <Menu
                                style={{ width: "2.4rem" }}
                                defaultSelectedKeys={['1']}
                                defaultOpenKeys={['1']}
                                mode="vertical"
                                theme="light"
                                style={{backgroundColor:"#FAFAFA"}}
                            >
                                <Menu.Item key="1">选择题</Menu.Item>
                                <Menu.Item key="2">填空题</Menu.Item>
                                <Menu.Item key="3">阅读题</Menu.Item>
                                <Menu.Item key="4">题型4</Menu.Item>
                                <Menu.Item key="5">题型5</Menu.Item>
                                <Menu.Item key="6">题型6</Menu.Item>
                                <Menu.Item key="7">题型7</Menu.Item>
                                <Menu.Item key="8">题型8</Menu.Item>
                                <Menu.Item key="9">题型9</Menu.Item>
                            </Menu>
                        </Sider>
                        <Content style={{backgroundColor:"white"}} className="content">
                            <div className="title">
                                <p>1.题目</p>
                            </div>
                            <div className="line"></div>
                            <Radio.Group value={1}>
                                <Space direction="vertical">
                                    <Radio value={1}>Option A</Radio>
                                    <Radio value={2}>Option B</Radio>
                                </Space>
                            </Radio.Group>
                            <Layout className="edit-box" style={{backgroundColor:"white"}}>
                                <Layout className="edit-core-box" style={{backgroundColor:"white"}}>
                                    <Sider theme="light" width="5.42rem" style={{marginRight:".06rem"}}>
                                        <div className="header-title">
                                            <span className="title">题目编辑</span>
                                            <PictureOutlined/>
                                        </div>
                                        <div className="subject-edit-box"></div>
                                    </Sider>
                                    <Content style={{backgroundColor:"white"}}>
                                        <div className="header-title">
                                            <span className="title">参数编辑</span>
                                            <PictureOutlined/>
                                        </div>
                                        <div className="params-edit-box">
                                            <Form
                                                name="basic"
                                                labelCol={{ span: 6 }}
                                                wrapperCol={{ span: 18 }}
                                                labelAlign="left"
                                            >
                                                <Form.Item
                                                    label="学科"
                                                    name="subjects"
                                                    colon={false}
                                                >
                                                    <Select placeholder="选择学科">
                                                        <Option value="male">学科一</Option>
                                                        <Option value="female">学科二</Option>
                                                        <Option value="other">学科三</Option>
                                                    </Select>
                                                </Form.Item>
                                                <Form.Item
                                                    label="难度"
                                                    name="difficulty"
                                                    colon={false}
                                                >
                                                    <div className="difficulty">
                                                        <div className="content-box">
                                                            <div className="value-box" style={{width:`${(this.state.difficultyValue-1)*25}%`}}></div>
                                                            {
                                                                [1,2,3,4,5].map(item=>(
                                                                    <div className={`value-item ${item===this.state.difficultyValue?"select-value":""} ${item>this.state.difficultyValue?"big-value":""}`} onClick={()=>{
                                                                        this.setState({
                                                                            difficultyValue:item
                                                                        })
                                                                    }}>
                                                                        <div className="value">{item}</div>
                                                                    </div>
                                                                ))
                                                            }
                                                        </div>
                                                    </div>
                                                </Form.Item>
                                                <Form.Item
                                                    label="能力纬度"
                                                    name="ability"
                                                    colon={false}
                                                >
                                                    <div className="params-item">参数1</div>
                                                    <div className="params-item">参数2</div>
                                                    <div className="params-item">参数3</div>
                                                    <div className="params-item">参数4</div>
                                                    <div className="params-item">参数5</div>
                                                </Form.Item>
                                                <Form.Item
                                                    label="内容纬度"
                                                    name="content"
                                                    colon={false}
                                                >
                                                    <div className="params-item">参数1</div>
                                                    <div className="params-item">参数2</div>
                                                    <div className="params-item">参数3</div>
                                                    <div className="params-item">参数4</div>
                                                    <div className="params-item">参数5</div>
                                                </Form.Item>
                                            </Form>
                                        </div>
                                    </Content>
                                </Layout>
                                <Footer style={{backgroundColor:"white",minHeight:"1rem"}} className="option-box">
                                    <div className="header-title">
                                        <span className="title">题目编辑</span>
                                        <div className="right-box">
                                            <span>图片</span>
                                            <span>正确答案</span>
                                            <span>上移下移</span>
                                        </div>
                                    </div>
                                    <div className="option-item">
                                        <div className="value-box">
                                            <Input placeholder="选项1" />
                                            <div className="oper-box">
                                                <PlusCircleOutlined className="oper-icon"/>
                                                <MinusCircleOutlined className="oper-icon"/>
                                            </div>
                                        </div>
                                        <div className="oper-box">
                                            <PictureOutlined className="oper-icon"/>
                                            <Checkbox/>
                                            <div className="move">
                                                <UpCircleOutlined className="oper-icon" style={{marginRight:".05rem"}}/>
                                                <DownCircleOutlined className="oper-icon"/>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="option-item">
                                        <div className="value-box">
                                            <Input placeholder="选项1" />
                                            <div className="oper-box">
                                                <PlusCircleOutlined className="oper-icon"/>
                                                <MinusCircleOutlined className="oper-icon"/>
                                            </div>
                                        </div>
                                        <div className="oper-box">
                                            <PictureOutlined className="oper-icon"/>
                                            <Checkbox/>
                                            <div className="move">
                                                <UpCircleOutlined className="oper-icon" style={{marginRight:".05rem"}}/>
                                                <DownCircleOutlined className="oper-icon"/>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="btn-box">
                                        <Button type="primary"  block>完成编辑</Button>
                                    </div>
                                </Footer>
                            </Layout>
                        </Content>
                        <Sider theme="light" width="7.47rem" className="question-box">
                            <div className="question-content-box">
                                <div className="title">相关题目</div>
                                <div className="filter-box">
                                    <span>筛选</span>
                                    <Search placeholder="input search text" style={{ width: 200 }} size="small"/>
                                </div>
                                <div className="question-item" style={{marginTop:".38rem"}}>
                                    <div className="content">
                                        <div className="left">
                                            <span>题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目......题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目.题目题目题目题目题目题目题目题目.(最多5行过多省略)</span>
                                        </div>
                                        <div className="right">
                                            <div className="info-line">
                                                <span>学科名称：数学</span>
                                            </div>
                                            <div className="info-line">
                                                <span>能力纬度：参数1；参数2；参数3...</span>
                                            </div>
                                            <div className="info-line">
                                                <span>内容纬度：参数1；参数2；参数3...</span>
                                            </div>
                                            <div className="info-line">
                                                <span>难度等级：参数1</span>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="footer">
                                        <div className="left">
                                            <span>答案：xxxxx</span>
                                        </div>
                                        <div className="right">
                                            <Button type="primary" size="small">使用</Button>
                                        </div>
                                    </div>
                                </div>
                                <div className="question-item" style={{marginTop:".38rem"}}>
                                    <div className="content">
                                        <div className="left">
                                            <span>题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目......题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目.题目题目题目题目题目题目题目题目.(最多5行过多省略)</span>
                                        </div>
                                        <div className="right">
                                            <div className="info-line">
                                                <span>学科名称：数学</span>
                                            </div>
                                            <div className="info-line">
                                                <span>能力纬度：参数1；参数2；参数3...</span>
                                            </div>
                                            <div className="info-line">
                                                <span>内容纬度：参数1；参数2；参数3...</span>
                                            </div>
                                            <div className="info-line">
                                                <span>难度等级：参数1</span>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="footer">
                                        <div className="left">
                                            <span>答案：xxxxx</span>
                                        </div>
                                        <div className="right">
                                            <Button type="primary" size="small">使用</Button>
                                        </div>
                                    </div>
                                </div>
                                <div className="question-item" style={{marginTop:".38rem"}}>
                                    <div className="content">
                                        <div className="left">
                                            <span>题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目......题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目题目.题目题目题目题目题目题目题目题目.(最多5行过多省略)</span>
                                        </div>
                                        <div className="right">
                                            <div className="info-line">
                                                <span>学科名称：数学</span>
                                            </div>
                                            <div className="info-line">
                                                <span>能力纬度：参数1；参数2；参数3...</span>
                                            </div>
                                            <div className="info-line">
                                                <span>内容纬度：参数1；参数2；参数3...</span>
                                            </div>
                                            <div className="info-line">
                                                <span>难度等级：参数1</span>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="footer">
                                        <div className="left">
                                            <span>答案：xxxxx</span>
                                        </div>
                                        <div className="right">
                                            <Button type="primary" size="small">使用</Button>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </Sider>
                    </Layout>
                </div>
                <Modal 
                    title="编辑试题信息" 
                    cancelText="取消"
                    okText="保存"
                    visible={this.state.editQuestionVisible} 
                    onOk={()=>{
                        
                    }} 
                    onCancel={()=>{
                        this.setState({
                            editQuestionVisible:false
                        })
                    }}
                >
                    <div className="upLoad-question-title">
                        <span>试题编号编号:absnahghj（自动生成）</span>
                    </div>
                    <Form 
                        labelCol={{ span: 5 }}
                        wrapperCol={{ span: 19 }}
                        labelAlign="left"
                    >
                        <Form.Item
                            name="project"
                            label="项目"
                            rules={[{ required: true, message: '请选择项目名称' }]}
                        >
                            <Select placeholder="选择项目名称">
                                <Option value="male">项目一</Option>
                                <Option value="female">项目二</Option>
                                <Option value="other">项目三</Option>
                            </Select>
                        </Form.Item> 
                        <Form.Item
                            name="subject"
                            label="学科"
                            rules={[{ required: true, message: '请选择学科' }]}
                        >
                            <Select placeholder="选择学科">
                                <Option value="male">学科一</Option>
                                <Option value="female">学科二</Option>
                                <Option value="other">学科三</Option>
                            </Select>
                        </Form.Item>
                        <Form.Item
                            name="content"
                            label="内容纬度"
                            rules={[{ required: true, message: '请选择内容纬度' }]}
                        >
                            <Select placeholder="选择内容纬度">
                                <Option value="male">内容纬度一</Option>
                                <Option value="female">内容纬度二</Option>
                                <Option value="other">内容纬度三</Option>
                            </Select>
                        </Form.Item>
                        <Form.Item
                            name="ability"
                            label="能力维度"
                            rules={[{ required: true, message: '请选择能力维度' }]}
                        >
                            <Select placeholder="选择能力维度">
                                <Option value="male">能力维度一</Option>
                                <Option value="female">能力维度二</Option>
                                <Option value="other">能力维度三</Option>
                            </Select>
                        </Form.Item>
                        <Form.Item
                            name="type"
                            label="题型"
                            rules={[{ required: true, message: '请选择题型' }]}
                        >
                            <Select placeholder="选择题型">
                                <Option value="male">题型一</Option>
                                <Option value="female">题型二</Option>
                                <Option value="other">题型三</Option>
                            </Select>
                        </Form.Item>
                    </Form>
                </Modal>
            </div>
        )
    }
}
