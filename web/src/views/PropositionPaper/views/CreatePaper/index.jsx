import React, { Component } from 'react'
import {PageHeader,Button,Spin,Descriptions,message,Row,Col,Input,Pagination,Tag,Select,Slider,Modal} from 'antd'
import request from '../../../../utils/request'
import UpLoadQuestionModal from '../../../../components/UpLoadQuestionModal'
import HistoryQuestion from '../../../../components/HistoryQuestion'
import BraftEditor from 'braft-editor'
import store from '../../../../store'
import 'braft-editor/dist/index.css'
import './index.less'

const { Option } = Select
const { Search } = Input

export default class index extends Component {

    state={
        initLoading:true,
        createTime:0,
        projectInfo:{},
        upLoadQuestionModalParams:{
            show:false,
            type:"update-paper"
        },
        questionList:[{
            state:"edit",               //表示编辑中
            subject:"",
            difficulty:1,
            ability:this.props.match.params.ability,
            content:this.props.match.params.content,
            body:BraftEditor.createEditorState(null)
        }],
        paperViewVisible:false,
        loadingState:false
    }

    componentDidMount(){
        let t =new Date()
        this.setState({
            createTime:`${t.getFullYear()}-${t.getMonth().toString().padStart(2,'0')}-${t.getDate().toString().padStart(2,'0')} ${t.getHours().toString().padStart(2,'0')}:${t.getMinutes().toString().padStart(2,'0')}:${t.getSeconds().toString().padStart(2,'0')}`
        })
        this.getProjectInfo()
    }

    upLoadQuestion=(info)=>{
        this.setState({
            loadingState:true
        })
        return request({
            url:"http://49.232.73.36:8082/qbank/question",
            method:"POST",
            data:{
                advanced_props:{
                    ctt_diff_1:info.difficulty,
                    ctt_diff_2:info.difficulty,
                    ctt_level:info.difficulty,
                    irt_level:info.difficulty
                },
                apply_record:{
                    grade_fits:this.state.projectInfo.basic_info.grade_range.join(','),
                    participant_count:0,
                    test_count:0,
                    test_region:[],
                    test_year:`${new Date().getFullYear()}`,
                },
                author:store.getState().userInfo.Id,
                basic_props:{
                    ability_dimension:info.ability,
                    description:"暂无",
                    details:info.body.toHTML(),
                    details_dimension:info.content,
                    encode:"",
                    keywords:[],
                    sub_ability_dimension:"",
                    sub_details_dimension:"",
                    subject:info.subject,
                    subject_requirements:""
                },
                extra_props:{
                    is_question_group:false,
                    is_scene:true,
                    material_length:0,
                    reading_material_topic:""
                },
                info:{
                    answer:"",
                    body:info.body.toHTML(),
                    solution:"无",
                    title:"无",
                    type:"选择题"
                },
                source_project:this.props.match.params.project,
                spec_props:{
                    article_type:"无",
                    length:"无",
                    topic:"无"
                }
            }
        })
    }

    getProjectInfo=()=>{
        this.setState({
            initLoading:true
        })
        request({
            url:`http://49.232.73.36:8081/review/proj/detailed/${this.props.match.params.project}`,
            method:"GET"
        }).then(res=>{
            this.setState({
                projectInfo:res.data.basic_info,
                initLoading:false
            })
        }).catch(err=>{
            this.setState({
                initLoading:false
            })
            this.props.history.push("/home/proposition-paper/home");
            message.error("编辑器加载失败！");
        })
    }

    render() {
        return (
            <div className="create-paper-page" data-component="create-paper-page">
                <PageHeader
                    ghost={false}
                    onBack={() => this.props.history.push("/home/proposition-paper/home")}
                    title="命题组卷"
                    subTitle="上传试题"
                    extra={[
                        <Button key="1" onClick={()=>{
                            this.setState({
                                upLoadQuestionModalParams:{
                                    show:true,
                                    type:"update-paper"
                                }
                            })
                        }}>编辑内容</Button>,
                    ]}
                    >
                    {
                        this.state.initLoading?(
                            <Spin spinning={true} tip="初始化中"></Spin>
                        ):(
                            <Descriptions size="small" column={3}>
                                <Descriptions.Item label="创建时间" key="createAt">{this.state.createTime}</Descriptions.Item>
                                <Descriptions.Item label="项目" key="peojects">{this.state.projectInfo.basic_info.name}</Descriptions.Item>
                                <Descriptions.Item label="学科" key="subjects">{
                                    this.state.projectInfo.basic_info.subjects.map((item,index)=>(
                                        <span>{item}{index===this.state.projectInfo.basic_info.subjects.length-1?"":"、"}</span>
                                    ))
                                }</Descriptions.Item>
                                
                                <Descriptions.Item label="内容纬度" key="content">{
                                    this.props.match.params.content.split(',').map((item,index)=>(
                                        <span>{item}{index===this.props.match.params.content.split(',').length-1?"":"、"}</span>
                                    ))
                                }</Descriptions.Item>
                                <Descriptions.Item label="能力纬度" key="ability">
                                    {
                                        this.props.match.params.ability.split(',').map((item,index)=>(
                                            <span>{item}{index===this.props.match.params.ability.split(',').length-1?"":"、"}</span>
                                        ))
                                    }
                                </Descriptions.Item>
                            </Descriptions>
                        )
                    }
                </PageHeader>
                <Spin spinning={this.state.loadingState} tip="上传中">
                    <Row className="container">
                        <Col span="8" className="left-box">
                            <div className="title-box">
                                <div className="title-value">
                                    <div className="ver-line"></div>
                                    <div className="value">
                                        <span>相关题库</span>
                                    </div>
                                </div>
                                <Search placeholder="input search text" style={{width:"2.64rem"}} size="middle" />
                            </div>
                            <HistoryQuestion/>
                            <HistoryQuestion/>
                            <HistoryQuestion/>
                            <Pagination defaultCurrent={1} total={50} className="page-spare"/>
                        </Col>
                        <Col span="16" className="right-box">
                            {
                                this.state.questionList.map((item,index)=>(
                                    <Row className="question-item" key={index+Math.random(100)} gutter={[8,8]}>
                                        <Col className="info-box" span="18">
                                            <Row className="main-box">
                                                {
                                                    item.state==='edit'?(
                                                        <BraftEditor
                                                            value={item.body}
                                                            onChange={(editorState)=>{
                                                                let questionList = [...this.state.questionList]
                                                                questionList[index].body = editorState
                                                                this.setState({ questionList })
                                                            }}
                                                            onSave={()=>{
                                                                console.log("保存题目")
                                                            }}
                                                        />
                                                    ):(
                                                        <div className="view-box" dangerouslySetInnerHTML={{__html:item.body.toHTML()}}>

                                                        </div>
                                                    )
                                                }
                                            </Row>
                                            <Row className="btn-line">
                                                <Button type="primary" size="small" style={{width:"1rem",marginRight:".2rem"}} onClick={()=>{
                                                    let questionList = [...this.state.questionList]
                                                    questionList[index].state = item.state==='edit'?"complete":"edit"
                                                    this.setState({
                                                        questionList
                                                    })
                                                }}>
                                                    {
                                                        item.state==='edit'?"保存":"编辑"
                                                    }
                                                </Button>
                                                <Button type="primary" danger size="small"  style={{width:"1rem"}} onClick={()=>{
                                                    let questionList = [...this.state.questionList]
                                                    questionList.splice(index,1);
                                                    this.setState({
                                                        questionList
                                                    })
                                                }}>
                                                    删除
                                                </Button>
                                            </Row>
                                        </Col>
                                        <Col className="params-box" span="6">
                                            <div className="title">
                                                <span>参数编辑</span>
                                            </div>
                                            <Row className="params-item">
                                                <Col className="label" span="8">
                                                    <span>学科</span>
                                                </Col>
                                                <Col className="value" span="16">
                                                    <Select 
                                                        placeholder="选择学科" 
                                                        value={item.subject} 
                                                        defaultValue={this.props.match.params.subject} 
                                                        onSelect={(e)=>{
                                                            let questionList = [...this.state.questionList]
                                                            questionList[index].subject = e
                                                            this.setState({
                                                                questionList
                                                            })
                                                        }}
                                                        size="small"
                                                    >
                                                        {
                                                            this.state.initLoading?"":this.state.projectInfo.basic_info.subjects.map((item,index)=>(
                                                                <Option value={item} key={index+Math.random(100)}>{item}</Option>
                                                            ))
                                                        }
                                                    </Select>
                                                </Col>
                                            </Row>
                                            <Row className="params-item" style={{marginTop:'.17rem'}}>
                                                <Col span="8" className="label">
                                                    <span>难度</span>
                                                </Col>
                                                <Col span="16" className="value">
                                                    <Slider marks={{1:1,2:2,3:3,4:4,5:5}} step={null} value={item.difficulty} defaultValue={1} max={5 } min={1} onChange={(e)=>{
                                                        let questionList = [...this.state.questionList]
                                                        questionList[index].difficulty = e
                                                        this.setState({
                                                            questionList
                                                        })
                                                    }}/>
                                                </Col>
                                            </Row>
                                            <Row className="params-item" style={{marginTop:'.17rem'}}>
                                                <Col span="8" className="label">
                                                    <span>能力纬度</span>
                                                </Col>
                                                <Col span="16" className="value">
                                                    <div className="tag-list">
                                                        {
                                                            item.ability.split(',').map(item=>(
                                                                <Tag>{item}</Tag>
                                                            ))
                                                        }
                                                    </div>
                                                </Col>
                                            </Row>
                                            <Row className="params-item" style={{marginTop:'.1rem'}}>
                                                <Col span="8" className="label">
                                                    <span>内容纬度</span>
                                                </Col>
                                                <Col span="16" className="value">
                                                    <div className="tag-list">
                                                        {
                                                            item.content.split(',').map(item=>(
                                                                <Tag>{item}</Tag>
                                                            ))
                                                        }
                                                    </div>
                                                </Col>
                                            </Row>
                                        </Col>
                                        <Row  gutter={[8,8]} className="oper-btn-line">
                                            <Col span="18" className="left">
                                                {
                                                    index===this.state.questionList.length-1?(
                                                        <Button type="primary" size="large" onClick={()=>{
                                                            let questionList = [...this.state.questionList]
                                                            questionList.push({
                                                                state:"edit",               //表示编辑中
                                                                subject:"",
                                                                difficulty:1,
                                                                ability:this.props.match.params.ability,
                                                                content:this.props.match.params.content,
                                                                body:BraftEditor.createEditorState(null)
                                                            })
                                                            this.setState({
                                                                questionList
                                                            })
                                                        }}>
                                                            添加一题    
                                                        </Button>
                                                    ):""
                                                }
                                            </Col>
                                            <Col span="6" className="right">
                                                <Button size="large" style={{backgroundColor:"#EEEEEE"}} onClick={()=>{
                                                    this.setState({
                                                        paperViewVisible:true
                                                    })
                                                }}>
                                                    预览试卷  
                                                </Button>
                                                <Button type="primary" size="large" style={{width:"1.2rem"}} onClick={async ()=>{
                                                    let questionIdList = []
                                                    for(let i=0;i<this.state.questionList.length;i++){
                                                        await this.upLoadQuestion(this.state.questionList[i]).then((res)=>{
                                                            questionIdList.push({
                                                                comment:"暂无",
                                                                question_id:res.data.data,
                                                                score:0
                                                            })
                                                        }).catch(err=>{
                                                            this.setState({
                                                                loadingState:false
                                                            })
                                                            message.error(err.message||"试题上传失败，请重试")
                                                        })
                                                    }
                                                    request({
                                                        url:"http://49.232.73.36:8082/qbank/testpaper",
                                                        method:"POST",
                                                        data:{
                                                            author:store.getState().userInfo.Id,
                                                            info:[{
                                                                description:"暂无",
                                                                question_list:questionIdList,
                                                                score:0,
                                                                title:"无"
                                                            }],
                                                            props:{
                                                                difficulty:"1",
                                                                grade_range:this.state.projectInfo.basic_info.grade_range,
                                                                subjects:this.state.projectInfo.basic_info.subjects,
                                                                time_limit:"0"
                                                            },
                                                            source_project:this.props.match.params.project,
                                                            title:"无"
                                                        }
                                                    }).then(res=>{
                                                        this.setState({
                                                            loadingState:false
                                                        })
                                                        this.props.history.push("/home/proposition-paper/home")
                                                        message.success("试题上传成功")
                                                    }).catch(err=>{
                                                        this.setState({
                                                            loadingState:false
                                                        })
                                                        message.error(err.message||"试题上传失败，请重试")
                                                    })
                                                }}>
                                                    保存    
                                                </Button>
                                            </Col>
                                        </Row>
                                    </Row>
                                ))
                            }
                        </Col>
                    
                    </Row>
                </Spin>
                <UpLoadQuestionModal
                    {...this.state.upLoadQuestionModalParams}
                    onClose={()=>{
                        this.setState({
                            upLoadQuestionModalParams:{
                                show:false,
                                type:"update-paper"
                            }
                        })
                    }}
                />
                <Modal width={1200} title="试卷预览" visible={this.state.paperViewVisible} onOk={()=>{
                    this.setState({
                        paperViewVisible:false
                    })
                }} onCancel={()=>{
                    this.setState({
                        paperViewVisible:false
                    })
                }}>
                    <div className="view-box">
                        {
                            this.state.questionList.map(item=>(
                                <div className="question-view-item" dangerouslySetInnerHTML={{__html:item.body.toHTML()}}></div>
                            ))
                        }
                    </div>
                </Modal>
            </div>
        )
    }
}
