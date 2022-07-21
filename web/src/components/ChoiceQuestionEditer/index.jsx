import React, { Component } from 'react'
import {withRouter} from 'react-router-dom'
import {  Row, Col,Select,Tag,Button,Spin,Slider,message } from "antd"
import BraftEditor from 'braft-editor'
import request from '../../utils/request'
import store from '../../store'
import 'braft-editor/dist/index.css'
import './index.less'

const { Option } = Select


class index extends Component {

    state = {
        editorState: BraftEditor.createEditorState(null),
        loadingState:false,
        questionParams:{
            subject:"",
            difficulty:1,
            answer:""
        }
    }

    upLoadQuestion = ()=>{
        this.setState({
            loadingState:true
        })
        request({
            url:"http://49.232.73.36:8082/qbank/question",
            method:"POST",
            data:{
                advanced_props:{
                    ctt_diff_1:this.state.questionParams.difficulty,
                    ctt_diff_2:this.state.questionParams.difficulty,
                    ctt_level:this.state.questionParams.difficulty,
                    irt_level:this.state.questionParams.difficulty
                },
                apply_record:{
                    grade_fits:this.props.grade_range.join(','),
                    participant_count:0,
                    test_count:0,
                    test_region:[],
                    test_year:`${new Date().getFullYear()}`,
                },
                author:store.getState().userInfo.Id,
                basic_props:{
                    ability_dimension:this.props.ability.join(','),
                    description:"暂无",
                    details:this.state.editorState.toHTML(),
                    details_dimension:this.props.content.join(','),
                    encode:"",
                    keywords:[],
                    sub_ability_dimension:"",
                    sub_details_dimension:"",
                    subject:this.state.questionParams.subject,
                    subject_requirements:""
                },
                extra_props:{
                    is_question_group:false,
                    is_scene:true,
                    material_length:0,
                    reading_material_topic:""
                },
                info:{
                    answer:this.state.questionParams.answer,
                    body:this.state.editorState.toHTML(),
                    solution:"无",
                    title:"无",
                    type:"选择题"
                },
                source_project:this.props.projectId,
                spec_props:{
                    article_type:"无",
                    length:"无",
                    topic:"无"
                }
            }
        }).then(res=>{
            this.setState({
                loadingState:false
            })
            this.props.history.push("/home/proposition-paper/home")
            message.success("上传成功")
        }).catch(err=>{
            this.setState({
                loadingState:false
            })
            message.error(err.message||"请求错误")
        })
    }


    componentDidMount(){
        this.setState({
            questionParams:Object.assign(this.state.questionParams,{subject:this.props.defaultSubjectValue})
        })
    }

    render() {
        return (
            <div className="choice-question-editer" data-component="choice-question-editer" id="choice-question-edit-box">
                <Spin spinning={this.state.loadingState} tip="上传试题中">
                    <Row>
                        <BraftEditor
                            value={this.state.editorState}
                            onChange={(editorState)=>{
                                this.setState({ editorState })
                            }}
                            onSave={()=>{
                                console.log("保存题目")
                            }}
                        />
                    </Row>
                    <Row className="question-params">
                        <div className="title">
                            <span>参数编辑</span>
                        </div>
                        <Row className="param-item" style={{marginTop:'.17rem'}}>
                            <Col span="4" className="label">
                                <span>学科</span>
                            </Col>
                            <Col span="20" className="value">
                                <Select 
                                    placeholder="选择学科" 
                                    value={this.state.questionParams.subject} 
                                    defaultValue={this.props.defaultSubjectValue} 
                                    onSelect={(e)=>{
                                        let questionParams = Object.assign(this.state.questionParams,{subject:e})
                                        this.setState({
                                            questionParams
                                        })
                                    }}
                                    size="small"
                                >
                                    {
                                        this.props.subjectList.map((item,index)=>(
                                            <Option value={item} key={index+Math.random(100)}>{item}</Option>
                                        ))
                                    }
                                </Select>
                            </Col>
                        </Row>
                        <Row className="param-item" style={{marginTop:'.17rem'}}>
                            <Col span="4" className="label">
                                <span>难度</span>
                            </Col>
                            <Col span="20" className="value">
                                <Slider marks={{1:1,2:2,3:3,4:4,5:5}} step={null} defaultValue={1} max={5 } min={1} onChange={(e)=>{
                                    let questionParams = Object.assign(this.state.questionParams,{difficulty:e});
                                    this.setState({
                                        questionParams
                                    })
                                }}/>
                            </Col>
                        </Row>
                        <Row className="param-item" style={{marginTop:'.3rem'}}>
                            <Col span="4" className="label">
                                <span>能力纬度</span>
                            </Col>
                            <Col span="20" className="value">
                                <div className="tag-list">
                                    {
                                        this.props.ability.map(item=>(
                                            <Tag>{item}</Tag>
                                        ))
                                    }
                                </div>
                            </Col>
                        </Row>
                        <Row className="param-item" style={{marginTop:'.17rem'}}>
                            <Col span="4" className="label">
                                <span>内容纬度</span>
                            </Col>
                            <Col span="20" className="value">
                                <div className="tag-list">
                                    {
                                        this.props.content.map(item=>(
                                            <Tag>{item}</Tag>
                                        ))
                                    }
                                </div>
                            </Col>
                        </Row>
                    </Row>
                    <div className="question-complete-box">
                        <Button type="primary" block onClick={()=>{
                            this.upLoadQuestion()
                        }}>完成编辑</Button>
                    </div>
                </Spin>
            </div>
        )
    }
}

export default withRouter(index)
