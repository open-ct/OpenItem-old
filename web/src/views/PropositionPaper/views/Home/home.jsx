import React, { Component } from 'react'
import { Button, Radio, Modal, Table, Pagination, Space,Spin, message, Card,Row,Col } from 'antd'
import UpLoadQuestionModal from '../../../../components/UpLoadQuestionModal'
import PropositionPaperIcon from '../../../../asset/images/proposition-paper-icon.png'
import { LockOutlined , EditOutlined, EllipsisOutlined } from '@ant-design/icons';
import request from '../../../../utils/request'
import store from '../../../../store';
import './index.less'

const { Meta } = Card

export default class index extends Component {

    state = {
        mode: 'testpaper',
        questionData:[],
        testpaperData:[],
        recordData:[{
            key:'1',
            date:'text',
            content:'/',
            param1:'/',
            param2:'/',
            param3:'/'

        },{
            key:'2',
            date:'text',
            content:'内容修改版本的文字内容......',
            param1:'内容纬度1；内容纬度2',
            param2:'/',
            param3:'/'

        },{
            key:'3',
            date:'text',
            content:'/',
            param1:'/',
            param2:'/',
            param3:'/'

        },{
            key:'4',
            date:'text',
            content:'/',
            param1:'/',
            param2:'/',
            param3:'/'

        },{
            key:'5',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'6',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'7',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'8',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'9',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'10',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'11',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'12',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'13',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        },{
            key:'14',
            date:'text',
            content:'text',
            param1:'text',
            param2:'text',
            param3:'text'

        }],

        upLoadQuestionModalParams:{
            show:false,
            type:"create"
        },

        questionRecord:{
            Id:"",
            viewIndex:0,
            recordList:[],
        },

        modifyRecordVisible:false,
        upLoadQuestionVisible:false,
        editQuestionVisible:false,
        createPaperVisible:false,
        loadingState:false,
        recordLoading:false,
        testpaperVisible:{
            show:false,
            id_list:[],
            loadingState:false,
            questionList:[]
        }
    }

    columns = [{
        title: '时间',
        key: 'date',
        width:110,
        render: (text, record) => (
            <span>{this.timeFilter(record.CreateAt)}</span>
        ),
    },{
        title: '内容',
        key: 'content',
        width:110,
        render: (text, record) => (
            <span>{record.basic_props.description}</span>
        ),
    },{
        title: '难度',
        key: 'param1',
        width:110,
        render: (text, record) => (
            <span>{record.advanced_props.irt_level}</span>
        ),
    },{
        title: '阶段',
        key: 'param2',
        width:110,
        render: (text, record) => (
            <span>{record.apply_record.grade_fits}</span>
        ),
    },{
        title: '学科',
        key: 'param3',
        width:110,
        render: (text, record) => (
            <span>{record.basic_props.subject}</span>
        ),
    },{
        title: '操作',
        key: 'operate',
        render: (text, record) => {
            return (
                <Button type="link" onClick={()=>{
                    this.setState({
                        questionRecord:Object.assign(this.state.questionRecord,{viewIndex:this.state.questionRecord.recordList.findIndex(item=>item.Id===record.Id)})
                    })
                }}>预览</Button>
            )
        },
    }];

    handleModeChange = e => {
        const mode = e.target.value;
        if(e.target.value==="questions"){
            this.getUserQuestionList()
        }else{
            this.getUserTestpaperList()
        }
        this.setState({ mode });
    };

    getQuestionRecord = (Id)=>{
        this.setState({
            recordLoading:true
        })
        request({
            url:`http://49.232.73.36:8082/qbank/question/trace/${Id}`,
            method:"GET"
        }).then(res=>{
            this.setState({
                questionRecord:Object.assign(this.state.questionRecord,{recordList:res.data.data,viewIndex:0}),
                recordLoading:false
            })
        }).catch(err=>{
            this.setState({
                recordLoading:false,
                modifyRecordVisible:false
            })
            message.error(err.message||"请求错误！")
        })
    }

    getQuestionListInfo = ()=>{
        this.setState({
            testpaperVisible:Object.assign(this.state.testpaperVisible,{loadingState:true})
        })
        request({
            url:"http://49.232.73.36:8082/qbank/query/t_question",
            method:"POST",
            data:{
                id_list:this.state.testpaperVisible.id_list
            }
        }).then(res=>{
            this.setState({
                testpaperVisible:Object.assign(this.state.testpaperVisible,{loadingState:false,questionList:Object.values(res.data.data)})
            })
        }).catch(err=>{
            this.setState({
                testpaperVisible:Object.assign(this.state.testpaperVisible,{loadingState:false,show:false})
            })
            message.error(err.message||"请求错误")
        })
    }

    loadQuestionData = ()=>{
        if(this.state.mode==="questions"){
            return this.state.questionData.map(item=>(
                // <div className="question-item" key={item.Id}>
                //     <div className="header">
                //         <span className="category">{item.basic_props.subject}</span>
                //         <span className="date">{this.timeFilter(item.CreateAt)}</span>
                //     </div>
                //     <div className="container">
                //         <div className="info">
                //             <span>{item.info.title}</span>
                //         </div>
                //         <div className="project">
                //             <span>{item.source_project}</span>
                //         </div>
                //     </div>
                //     <div className="footer">
                //         <div className="action-item b-right">
                //             <LockOutlined/>
                //         </div>
                //         <div className="action-item b-right" onClick={()=>{
                //             this.setState({
                //                 upLoadQuestionModalParams:{
                //                     type:"edit",
                //                     show:true
                //                 }
                //             })
                //         }}>
                //             <EditOutlined/>
                //         </div>
                //         <div className="action-item" onClick={()=>{
                //             this.setState({
                //                 modifyRecordVisible:true
                //             })
                //             this.getQuestionRecord(item.uuid)
                //         }}>
                //             <EllipsisOutlined/>
                //         </div>
                //     </div>
                // </div>
                <Card
                    style={{ width: 300,float:"left" }}
                    actions={[
                        <LockOutlined key="lock"/>,
                        <EditOutlined key="edit" onClick={()=>{
                            this.setState({
                                upLoadQuestionModalParams:{
                                    type:"update",
                                    show:true
                                }
                            })
                        }}/>,
                        <EllipsisOutlined key="ellipsis" onClick={()=>{
                            this.setState({
                                modifyRecordVisible:true
                            })
                            this.getQuestionRecord(item.uuid)
                        }}/>
                    ]}
                >
                    <Meta
                        title={(
                            <div className="header" style={{display:"flex",alignItems:"center",justifyContent:"space-between"}}>
                                <span>{item.basic_props.subject||"未知"}</span>
                                <span>{this.timeFilter(item.CreateAt)}</span>
                            </div>
                        )}
                        description={item.info.title}
                    />
                </Card>
            ))
        }else{
            return this.state.testpaperData.map(item=>(
                <Card
                    style={{ width: 300,float:"left" }}
                    actions={[
                        <LockOutlined key="lock"/>,
                        <EditOutlined key="edit" onClick={()=>{
                            this.setState({
                                upLoadQuestionModalParams:{
                                    type:"update-paper",
                                    show:true
                                }
                            })
                        }}/>,
                        <EllipsisOutlined key="ellipsis" onClick={()=>{
                            this.setState({
                                testpaperVisible:Object.assign(this.state.testpaperVisible,{id_list:item.info[0].question_list.map(item=>item.question_id),show:true})
                            })
                            this.getQuestionListInfo()
                        }}/>
                    ]}
                >
                    <Meta
                        title={(
                            <div className="header" style={{display:"flex",alignItems:"center",justifyContent:"space-between"}}>
                                <span>{item.props.subjects[0]||"未知"}</span>
                                <span>{this.timeFilter(item.CreateAt)}</span>
                            </div>
                        )}
                        description={item.info[0].description}
                    />
                </Card>
            ))
        }
        
    }

    timeFilter = (time)=>{
        let date = new Date(time);
        //${date.getHours().toString().padStart(2,'0')}:${date.getMinutes().toString().padStart(2,'0')}:${date.getSeconds().toString().padStart(2,'0')}
        return `${date.getFullYear()}-${date.getMonth().toString().padStart(2,'0')}-${date.getDate().toString().padStart(2,'0')}`
    }

    componentDidMount = ()=>{
       this.state.mode==="testpaper"?this.getUserTestpaperList():this.getUserQuestionList()
    }

    getUserQuestionList = ()=>{
        this.setState({
            loadingState:true
        })
        request({
            url:`http://49.232.73.36:8082/qbank/question/user_t/${store.getState().userInfo.Id}`,
            method:"GET",
        }).then(res=>{
            this.setState({
                questionData:res.data.data,
                loadingState:false
            })
        }).catch(err=>{
            this.setState({
                loadingState:false
            })
            message.error(err.message||"加载失败")
        })
    }

    getUserTestpaperList = ()=>{
        this.setState({
            loadingState:true
        })
        request({
            url:`http://49.232.73.36:8082/qbank/testpaper/user_t/${store.getState().userInfo.Id}`,
            method:"GET"
        }).then(res=>{
            console.log(res.data.data)
            this.setState({
                testpaperData:res.data.data,
                loadingState:false
            })
        }).catch(err=>{
            this.setState({
                loadingState:false
            })
            message.error(err.message||"加载失败")
        })
    }

    render() {
        return (
            <div className="proposition-paper-home-page" data-component="proposition-paper-home-page">
                <div className="header">
                    <div className="describe">
                        <div className="content">
                            <h1>实现高效命题组卷</h1>
                            <div className="context">
                                单题编辑上传，试卷调用，轻松快捷<br/>
                                xxxxxxxxxx示例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言例宣传语言
                            </div>
                        </div>
                        <div className="btn-box">
                            <Button type="primary" onClick={()=>{
                                this.setState({
                                    upLoadQuestionModalParams:{
                                        type:"create",
                                        show:true
                                    }
                                })
                            }}>上传试题</Button>
                            <Button type="primary" style={{marginLeft:'.18rem'}} onClick={()=>{
                                this.setState({
                                    upLoadQuestionModalParams:{
                                        type:"create-paper",
                                        show:true
                                    }
                                })
                            }}>创建试卷</Button>
                        </div>
                    </div>
                    <img src={PropositionPaperIcon} alt="图片" className="icon"></img>
                </div>
                <div className="container">
                    <div className="category-list">
                        <Radio.Group onChange={this.handleModeChange} value={this.state.mode} >
                            <Radio.Button value="questions">我的试题</Radio.Button>
                            <Radio.Button value="testpaper" style={{ marginLeft: 4 }}>我的试卷库</Radio.Button>
                        </Radio.Group>
                    </div>
                    <div className="main">
                        <Spin spinning={this.state.loadingState} tip="加载中">
                            {
                                this.state.loadingState?"":this.loadQuestionData()
                            }
                        </Spin>
                    </div>
                </div>
                <UpLoadQuestionModal
                    {...this.state.upLoadQuestionModalParams}
                    onClose={()=>{
                        let upLoadQuestionModalParams = Object.assign(this.state.upLoadQuestionModalParams,{show:false})
                        this.setState({
                            upLoadQuestionModalParams
                        })
                    }}
                />
                <Modal 
                    title="修改记录" 
                    visible={this.state.modifyRecordVisible} 
                    footer={null} 
                    onCancel={()=>{
                        this.setState({
                            modifyRecordVisible:false
                        })
                    }}  
                    width="16.12rem"
                >
                    <Spin spinning={this.state.recordLoading} tip="加载中">
                        <div className="question-modify-record-box">
                            <div className="preview-box">
                                <div className="header">
                                    <span>题目效果预览窗</span>
                                </div>
                                <div className="view-box" dangerouslySetInnerHTML={{__html:this.state.recordLoading?"加载中":(
                                        this.state.questionRecord.recordList[this.state.questionRecord.viewIndex]?this.state.questionRecord.recordList[this.state.questionRecord.viewIndex].info.body:"渲染失败"
                                    )}}>
                                </div>
                            </div>
                            <div className="record-box">
                                <Space align="center">
                                    修改历史版本
                                </Space>
                                <div className="container">
                                    <Table 
                                        dataSource={this.state.questionRecord.recordList} 
                                        columns={this.columns} 
                                        pagination={false}
                                        scroll = {{y:'calc(100% - .56rem)'}}
                                    />
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
                        </div>
                    </Spin>
                </Modal>
                <Modal 
                    title="试卷详情" 
                    visible={this.state.testpaperVisible.show} 
                    width="80%"
                    onOk={()=>{
                        this.setState({
                            testpaperVisible:Object.assign(this.state.testpaperVisible,{show:false})
                        })
                    }} 
                    onCancel={()=>{
                        this.setState({
                            testpaperVisible:Object.assign(this.state.testpaperVisible,{show:false})
                        })
                    }}
                >
                    <Spin spinning={this.state.testpaperVisible.loadingState} tip="加载中">
                        {
                            this.state.testpaperVisible.loadingState?"":(
                                <>
                                    {
                                        this.state.testpaperVisible.questionList.map((item,index)=>(
                                            <div className="paper-question-item" key={item.id}>
                                                <Row className="header">
                                                    <Col span="4">
                                                        <span>序号：<span style={{fontWeight:"bold",color:"red"}}>{index+1}</span></span>
                                                    </Col>
                                                    <Col span="6">
                                                        <span>
                                                            测试年份：<span style={{fontWeight:"bold",color:"green"}}>{item.apply_record.test_year}</span>
                                                        </span>
                                                    </Col>
                                                    <Col span="4">
                                                        <span>
                                                            试题难度：<span style={{fontWeight:"bold",color:"blue"}}>{item.advanced_props.irt_level}</span>
                                                        </span>
                                                    </Col>
                                                    <Col span="4">
                                                        <span>
                                                            试题类型：{item.info.type}
                                                        </span>
                                                    </Col>
                                                    <Col span="4">
                                                        <span>
                                                            题目答案：{item.info.answer||"无"}
                                                        </span>
                                                    </Col>
                                                </Row>
                                                <div className="body" dangerouslySetInnerHTML={{__html:item.info.body}}></div>
                                                <div className="footer">
                                                    <Button type="primary" style={{float:"right"}} onClick={()=>{
                                                        this.getQuestionRecord(item.uuid)
                                                        this.setState({
                                                            modifyRecordVisible:true
                                                        })
                                                    }}>查看历史版本</Button>
                                                </div>
                                                <br/>
                                                <br/>
                                            </div>
                                        ))
                                    }
                                </>
                            )
                        }
                    </Spin>
                </Modal>
            </div>
        )
    }
}
