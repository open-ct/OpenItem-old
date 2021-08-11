import React, { Component } from 'react'
import { Row, Col, Pagination, Space, Tag, Upload, Button, Table, Modal, Radio, message, Form, Input, Calendar, Select, Statistic, Empty } from 'antd'
import DocumentTitle from 'react-document-title'
import ModulaCard from '../ModulaCard'
import { FileTextTwoTone, UploadOutlined, FieldTimeOutlined, LinkOutlined, InboxOutlined,PlusOutlined,FileExcelFilled,FileZipFilled,FileTextFilled,FileMarkdownFilled,FilePptFilled } from '@ant-design/icons';
import './index.less'

const { Search } = Input

export default class index extends Component {

    state = {
        expertTableData:[{
            key: '1',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料1',
            result:0,
            feedbackMaterial:'反馈材料1'
        },{
            key: '2',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料2',
            result:1,
            feedbackMaterial:'反馈材料2'
        },{
            key: '3',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料3',
            result:2,
            feedbackMaterial:'反馈材料3'
        },{
            key: '4',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料4',
            result:3,
            feedbackMaterial:''
        },{
            key: '5',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料1',
            result:0,
            feedbackMaterial:'反馈材料1'
        },{
            key: '6',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料2',
            result:1,
            feedbackMaterial:'反馈材料2'
        },{
            key: '7',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料3',
            result:2,
            feedbackMaterial:'反馈材料3'
        },{
            key: '8',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料4',
            result:3,
            feedbackMaterial:''
        },{
            key: '9',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料1',
            result:0,
            feedbackMaterial:'反馈材料1'
        },{
            key: '10',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料2',
            result:1,
            feedbackMaterial:'反馈材料2'
        },{
            key: '11',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料3',
            result:2,
            feedbackMaterial:'反馈材料3'
        },{
            key: '12',
            date:' 2021-8-10',
            user: 'stacker',
            reviewMaterials:'待评审材料4',
            result:3,
            feedbackMaterial:''
        }],                         //专家用户数据

        assistantTableData:[{
            key: '1',
            id:'openct-12345689',
            date: '2021-08-11',
            reviewMaterials:'待评审材料1',
            result:0,
            feedback:'反馈意见',
            isDelay:false,
            isModify:false
        },{
            key: '2',
            id:'openct-12345689',
            date: '2021-08-11',
            reviewMaterials:'待评审材料2',
            result:1,
            feedback:'',
            isDelay:true,
            isModify:false
        },{
            key: '3',
            id:'openct-12345689',
            date: '2021-08-11',
            reviewMaterials:'待评审材料1',
            result:2,
            feedback:'反馈意见',
            isDelay:false,
            isModify:false
        },{
            key: '4',
            id:'openct-12345689',
            date: '2021-08-11',
            reviewMaterials:'待评审材料2',
            result:3,
            feedback:'',
            isDelay:true,
            isModify:false
        },{
            key: '5',
            id:'openct-12345689',
            date: '2021-08-11',
            reviewMaterials:'待评审材料1',
            result:0,
            feedback:'反馈意见',
            isDelay:false,
            isModify:false
        },{
            key: '6',
            id:'openct-12345689',
            date: '2021-08-11',
            reviewMaterials:'待评审材料2',
            result:1,
            feedback:'反馈意见',
            isDelay:true,
            isModify:true
        },{
            key: '7',
            id:'openct-12345689',
            date: '2021-08-11',
            reviewMaterials:'待评审材料1',
            result:2,
            feedback:'反馈意见',
            isDelay:false,
            isModify:false
        },{
            key: '8',
            id:'openct-12345689',
            date: '2021-08-11',
            reviewMaterials:'待评审材料2',
            result:1,
            feedback:'反馈意见',
            isDelay:true,
            isModify:true
        },{
            key: '9',
            id:'openct-12345689',
            date: '2021-08-11',
            reviewMaterials:'待评审材料1',
            result:3,
            feedback:'',
            isDelay:false,
            isModify:false
        },{
            key: '10',
            id:'openct-12345689',
            date: '2021-08-11',
            reviewMaterials:'待评审材料2',
            result:0,
            feedback:'反馈意见',
            isDelay:true,
            isModify:false
        }],                         //助理用户数据

        teacherTableData:[{
            key:'1',
            id:'openct-12345689',
            date:'2021-08-11',
            user:'stacker',
            reviewMaterials:'待评审材料1',
            result:0,
            feedbackMaterial:'反馈材料1'
        },{
            key:'2',
            id:'openct-12345689',
            date:'2021-08-11',
            user:'stacker',
            reviewMaterials:'待评审材料1',
            result:1,
            feedbackMaterial:'反馈材料1'
        },{
            key:'3',
            id:'openct-12345689',
            date:'2021-08-11',
            user:'stacker',
            reviewMaterials:'待评审材料1',
            result:2,
            feedbackMaterial:'反馈材料1'
        },{
            key:'4',
            id:'openct-12345689',
            date:'2021-08-11',
            user:'stacker',
            reviewMaterials:'待评审材料1',
            result:3,
            feedbackMaterial:'反馈材料1'
        }],

        recordList:[{
            key:'1',
            date:'2021-08-11',
            result:0,
            reviewMaterials:'待评审材料1',
            feedback:'反馈意见',
        },{
            key:'2',
            date:'2021-08-11',
            result:1,
            reviewMaterials:'待评审材料1',
            feedback:'反馈意见',
        },{
            key:'3',
            date:'2021-08-11',
            result:2,
            reviewMaterials:'待评审材料1',
            feedback:'反馈意见',
        },{
            key:'4',
            date:'2021-08-11',
            result:3,
            reviewMaterials:'待评审材料1',
            feedback:'',
        }],                         //历史修改记录数据

        defaultFileList: [{
              uid: '1',
              name: 'xxx.png',
              status: 'done',
              response: 'Server Error 500', // custom error message to show
              url: 'http://www.baidu.com/xxx.png',
        },{
              uid: '2',
              name: 'yyy.png',
              status: 'done',
              url: 'http://www.baidu.com/yyy.png',
        },{
              uid: '3',
              name: 'zzz.png',
              status: 'error',
              response: 'Server Error 500', // custom error message to show
              url: 'http://www.baidu.com/zzz.png',
        }],

        identity:'expert',          //控制当前用户身份 expert专家/assistant助理/teacher教师
        reviewResultsVisible:false, //评审意见修改框控制
        modifyRecordVisible:false,
        upLoadMaterialVisible:false,
        calendarVisible:false,      //展示日历
    }

    expertTableColumns = [{
        title: '上传时间',
        dataIndex: 'date',
        key: 'date',
        align: 'center',
        width:150
    },{
        title: '上传用户',
        dataIndex: 'user',
        key: 'user',
        align: 'center',
        width:120
    },{
        title: '评审材料',
        dataIndex: 'review-materials',
        key: 'review-materials',
        align: 'center',
        width:220,
        render: (text, record) => (
            <Space size="middle">
              <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.reviewMaterials}</span><FileTextTwoTone />
            </Space>
        )
    },{
        title: '评审结果',
        dataIndex: 'result',
        key: 'result',
        align: 'center',
        width: 100,
        render: (text, record) => {
            let levelList = [{
                mode:'通过',
                color:'#87D068'
            },{
                mode:'再修改',
                color:'#2DB7F5'
            },{
                mode:'驳回',
                color:'#FF5500'
            },{
                mode:'审核意见',
                color:'default'
            }]
            return (
                <Space size="middle">
                    <Tag color={levelList[record.result].color} onClick={()=>{
                        this.setState({
                            reviewResultsVisible:true
                        })
                    }} style={{cursor:'pointer'}}>{levelList[record.result].mode}</Tag>
                </Space>
            )
        }
    },{
        title: '反馈批注材料',
        dataIndex: 'feedback-material',
        key: 'feedback-material',
        align: 'center',
        render: (text, record) => {
            if(record.feedbackMaterial===""){
                return (
                    <Space size="middle">
                        <Upload>
                            <Button icon={<UploadOutlined />} size="small">上传批注</Button>
                        </Upload>
                    </Space>
                )
            }
            return (
                <Space size="middle">
                  <span>{record.feedbackMaterial}</span><FileTextTwoTone />
                </Space>
            )
        }
    }]  //专家用户展示的Table

    assistantTableColumns = [{
        title: '材料编号',
        dataIndex: 'id',
        key: 'id',
    },{
        title: '上传时间',
        key: 'date',
        render: (text, record) => (
            <Space size="middle">
               <span>{record.date}</span>{record.isDelay?(<Tag color="error">延时提交</Tag>):''}
            </Space>
        )
    },{
        title: '评审材料',
        key: 'review-materials',
        width:220,
        render: (text, record) => (
            <Space size="middle">
              <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.reviewMaterials}</span><FileTextTwoTone />
            </Space>
        )
    },{
        title: '评审结果',
        key: 'result',
        render: (text, record) => {
            let levelList = [{
                mode:'通过',
                color:'#87D068'
            },{
                mode:'再修改',
                color:'#2DB7F5'
            },{
                mode:'驳回',
                color:'#FF5500'
            },{
                mode:'等待评审',
                color:'default'
            }]
            return (
                <Space size="middle">
                    <Tag color={levelList[record.result].color} onClick={()=>{
                        this.setState({
                            reviewResultsVisible:true
                        })
                    }} style={{cursor:'pointer'}}>{levelList[record.result].mode}</Tag>
                </Space>
            )
        }
    },{
        title: '反馈意见',
        key: 'feedback',
        render: (text, record) => {
            if(record.feedback===""){
                return (
                    <Space size="middle">
                        <span style={{cursor:'pointer'}}>等待评审...</span>
                    </Space>
                )
            }
            return (
                <Space size="middle">
                  <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.feedback}</span><FileTextTwoTone />
                </Space>
            )
        }
    },{
        title: '修改记录',
        key: 'isModify',
        align: 'center',
        width:80,
        render: (text, record) => (
            <Space size="middle">
               {record.isModify?(<Button type="link" onClick={()=>{
                    this.setState({
                        modifyRecordVisible:true
                    })
               }}>查看</Button>):(<span>无</span>)}
            </Space>
        )
    }]  //助理用户展示Table

    teacherTableColumns = [{
        title: '材料编号',
        dataIndex: 'id',
        key: 'id',
    },{
        title: '上传时间',
        dataIndex: 'date',
        key: 'date',
    },{
        title: '上传用户',
        dataIndex: 'user',
        key: 'user',
    },{
        title: '评审材料',
        key: 'review-materials',
        width:220,
        render: (text, record) => (
            <Space size="middle">
              <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.reviewMaterials}</span><FileTextTwoTone />
            </Space>
        )
    },{
        title: '评审结果',
        key: 'result',
        render: (text, record) => {
            let levelList = [{
                mode:'通过',
                color:'#87D068'
            },{
                mode:'再修改',
                color:'#2DB7F5'
            },{
                mode:'驳回',
                color:'#FF5500'
            },{
                mode:'等待评审',
                color:'default'
            }]
            return (
                <Space size="middle">
                    <Tag color={levelList[record.result].color}>{levelList[record.result].mode}</Tag>
                </Space>
            )
        }
    },{
        title: '反馈批注材料',
        dataIndex: 'feedback-material',
        key: 'feedback-material',
        align: 'center',
        render: (text, record) => {
            return (
                <Space size="middle">
                  <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.feedbackMaterial}</span><FileTextTwoTone />
                </Space>
            )
        }
    }]

    recordListColums = [{
        title: '上传时间',
        dataIndex: 'date',
        key: 'date',
    },{
        title: '评审材料',
        key: 'review-materials',
        width:220,
        render: (text, record) => (
            <Space size="middle">
              <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.reviewMaterials}</span><FileTextTwoTone />
            </Space>
        )
    },{
        title: '评审结果',
        key: 'result',
        render: (text, record) => {
            let levelList = [{
                mode:'通过',
                color:'#87D068'
            },{
                mode:'再修改',
                color:'#2DB7F5'
            },{
                mode:'驳回',
                color:'#FF5500'
            },{
                mode:'等待评审',
                color:'default'
            }]
            return (
                <Space size="middle">
                    <Tag color={levelList[record.result].color}>{levelList[record.result].mode}</Tag>
                </Space>
            )
        }
    },{
        title: '反馈意见',
        key: 'feedback',
        render: (text, record) => {
            if(record.feedback===""){
                return (
                    <Space size="middle">
                        <span style={{cursor:'pointer'}}>等待评审...</span>
                    </Space>
                )
            }
            return (
                <Space size="middle">
                  <span style={{cursor:'pointer'}} onClick={this.downLoadFile}>{record.feedback}</span><FileTextTwoTone />
                </Space>
            )
        }
    }]

    //根据身份不同加载右侧table
    loadMaterialReview = ()=>{
        return (
            <Table 
                dataSource={this.state[`${this.state.identity}TableData`]} 
                columns={this[`${this.state.identity}TableColumns`]} size="small" 
                pagination={false} 
                scroll={{ y: '5.8rem'}}
            />
        )
    }

    //根据身份不同加载不同任务要求交互
    loadTaskRequirements = ()=>{
        let editRender = ['expert','teacher']
        let choice = [{
            identity:'expert',
            dom:()=>(
                <div className="choice-box">
                    <div className="calendar-btn">
                        <Button size="small" icon={<FieldTimeOutlined />} onClick={()=>{
                            this.setState({
                                calendarVisible:true
                            })
                        }}>截止时间：2021年8月20日</Button>
                        <div className="calendar-box" style={{maxHeight:this.state.calendarVisible?'5rem':0}}>
                            <Calendar
                                fullscreen={false}
                                headerRender={({ value, type, onChange, onTypeChange }) => {
                                    const start = 0;
                                    const end = 12;
                                    const monthOptions = [];

                                    const current = value.clone();
                                    const localeData = value.localeData();
                                    const months = [];
                                    for (let i = 0; i < 12; i++) {
                                    current.month(i);
                                    months.push(localeData.monthsShort(current));
                                    }

                                    for (let index = start; index < end; index++) {
                                    monthOptions.push(
                                        <Select.Option className="month-item" key={`${index}`}>
                                        {months[index]}
                                        </Select.Option>,
                                    );
                                    }
                                    const month = value.month();

                                    const year = value.year();
                                    const options = [];
                                    for (let i = year - 10; i < year + 10; i += 1) {
                                    options.push(
                                        <Select.Option key={i} value={i} className="year-item">
                                        {i}
                                        </Select.Option>,
                                    );
                                    }
                                    return (
                                    <div style={{ padding: 8 }}>
                                        <Row gutter={8} justify="end">
                                        <Col>
                                            <Select
                                            size="small"
                                            dropdownMatchSelectWidth={false}
                                            className="my-year-select"
                                            onChange={newYear => {
                                                const now = value.clone().year(newYear);
                                                onChange(now);
                                            }}
                                            value={String(year)}
                                            >
                                            {options}
                                            </Select>
                                        </Col>
                                        <Col>
                                            <Select
                                            size="small"
                                            dropdownMatchSelectWidth={false}
                                            value={String(month)}
                                            onChange={selectedMonth => {
                                                const newValue = value.clone();
                                                newValue.month(parseInt(selectedMonth, 10));
                                                onChange(newValue);
                                            }}
                                            >
                                            {monthOptions}
                                            </Select>
                                        </Col>
                                        </Row>
                                    </div>
                                    );
                                }}
                                onSelect={()=>{
                                    this.setState({
                                        calendarVisible:false
                                    })
                                }}
                            />
                        </div>
                    </div>
                    <Upload >
                        <Button size="small" icon={<LinkOutlined />} style={{marginLeft:'.1246rem'}}>附件：xxxx</Button>
                    </Upload>
                </div>
            )
        },{
            identity:'assistant',
            dom:()=>(
                <div className="choice-box">
                    <Button size="small" style={{marginRight:'.1246rem'}} icon={<FieldTimeOutlined />}>截止时间：2021年8月20日</Button>
                    <Button size="small" icon={<LinkOutlined />} onClick={this.downLoadFile}>附件：xxxx</Button>
                </div>
            )
        },{
            identity:'teacher',
            dom:()=>(
                <div className="choice-box">
                    <Button size="small" style={{marginRight:'.1246rem'}} icon={<FieldTimeOutlined />}>截止时间：2021年8月20日</Button>
                    <Button size="small" icon={<LinkOutlined />} onClick={this.downLoadFile}>附件：xxxx</Button>
                </div>
            )
        }]
        return (
            <div className="task-requirements-box">
                <div className="describe-box">
                    <div className="context">
                        <p>任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明任务要求具体说明，详细内容可参考附件</p>
                    </div>
                    <div className="btn-box">
                        {editRender.includes(this.state.identity)?(<Button type="link">编辑</Button>):''}
                    </div>
                </div>
                {choice[choice.findIndex(item=>item.identity===this.state.identity)].dom()}
            </div>
        )
    }

    //根据身份不同加载不同材料仓库
    loadMaterialWarehouse = ()=>{
        let domList = [{
            identity:'expert',
            dom:()=>(
                <div className="expert-box">
                    <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />
                </div>
            )
        },{
            identity:'assistant',
            dom:()=>(
                <Upload
                    name="avatar"
                    listType="picture-card"
                    className="avatar-uploader"
                    showUploadList={false}
                    action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
                    openFileDialogOnClick={false}
                >
                    <div className="assistant-box" onClick={()=>{
                        this.setState({
                            upLoadMaterialVisible:true
                        })
                    }}>
                        <PlusOutlined />
                        <div style={{ marginTop: 8,width:'100%' }}>Upload</div>
                    </div>
                </Upload>
            )
        },{
            identity:'teacher',
            dom:()=>{
                let fileType = [<FileTextFilled />,<FileExcelFilled/>,<FileZipFilled/>,<FileMarkdownFilled />,<FilePptFilled />]
                let fileList = [{
                    id:'AK000123456',
                    fileType:0
                },{
                    id:'AK000123456',
                    fileType:1
                },{
                    id:'AK000123456',
                    fileType:2
                },{
                    id:'AK000123456',
                    fileType:3
                },{
                    id:'AK000123456',
                    fileType:4
                },{
                    id:'AK000123456',
                    fileType:0
                },{
                    id:'AK000123456',
                    fileType:1
                },{
                    id:'AK000123456',
                    fileType:2
                },{
                    id:'AK000123456',
                    fileType:3
                },{
                    id:'AK000123456',
                    fileType:4
                },{
                    id:'AK000123456',
                    fileType:4
                },{
                    id:'AK000123456',
                    fileType:4
                }]
                return (
                    <div className="teacher-box">
                        {fileList.map(item=>(
                            <div className="file-item">
                                {fileType[item.fileType]}
                                <div className="key">
                                    {item.id}
                                </div>
                            </div>
                        ))}        
                    </div>
                )
            }
        }]
        return (
            <div className="container">
                {domList[domList.findIndex(item=>item.identity===this.state.identity)].dom()}                               
            </div>
        )
    }



    downLoadFile = ()=>{
        message.success('下载成功');
    }

    render() {
        return (
            <DocumentTitle title="项目管理-命题指南与测试">
                <div className="proposition-guide-page" data-component="proposition-guide-page">
                    <Row style={{height:'7.48rem',marginBottom:'.22rem'}} gutter={24}>
                        <Col span={9}>
                            <Row style={{height:'1.9rem',marginBottom:'.22rem'}}>
                                <ModulaCard title="任务要求">
                                    {this.loadTaskRequirements()}
                                </ModulaCard>
                            </Row>
                            <Row style={{height:'1.9rem',marginBottom:'.12rem'}}>
                                <ModulaCard title="完成情况">
                                    <div className="situation-box">
                                        <div className="left-box">
                                            <div className="title">
                                                <span>自2021年4月20日截止今日</span>
                                            </div>
                                            <div className="value-list">
                                                <Row gutter={16}>
                                                    <Col span={12}>
                                                        <Statistic title="材料总计" value={93}/>
                                                    </Col>
                                                    <Col span={12}>
                                                        <Statistic title="通过率" value={3} suffix="%" />
                                                    </Col>
                                                </Row>
                                            </div>
                                        </div>
                                        <div className="right-box">
                                            <Row gutter={16} style={{width:'4.2rem'}}>
                                                <Col span={4}>
                                                    <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#87D068'}}></div>通过</div>} value={3}/>
                                                </Col>
                                                <Col span={4}>
                                                    <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#FF5500'}}></div>驳回</div>}  value={3} />
                                                </Col>
                                                <Col span={5}>
                                                    <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#FF5500'}}></div>未上传</div>}  value={3} />
                                                </Col>
                                                <Col span={5}>
                                                    <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#2DB7F5'}}></div>待审核</div>}  value={3}/>
                                                </Col>
                                                <Col span={5}>
                                                    <Statistic title={<div className="statistic-item-title"><div className="circle" style={{backgroundColor:'#2DB7F5'}}></div>再修改</div>}  value={3} />
                                                </Col>
                                            </Row>
                                        </div>
                                    </div>
                                </ModulaCard>
                            </Row>
                            <Row style={{height:'3.34rem'}}>
                                <ModulaCard title="材料仓库" right={<Search placeholder="input search text" size="small" style={{ width: 200 }} />}>
                                    <div className="warehouse-box">
                                        {this.loadMaterialWarehouse()}
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
                            </Row>
                        </Col>
                        <Col span={15}>
                            <ModulaCard title="材料审核">
                                <div className="examine-box">
                                    {this.loadMaterialReview()}
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
                        </Col>
                    </Row>
                    <Modal title="评审意见" visible={this.state.reviewResultsVisible} onOk={()=>{
                        this.setState({
                            reviewResultsVisible:false
                        })
                    }} onCancel={()=>{
                        this.setState({
                            reviewResultsVisible:false
                        })
                    }}>
                        <span>材料评审结果：</span>
                        <Radio.Group name="radiogroup" defaultValue={1}>
                            <Radio value={1}>通过</Radio>
                            <Radio value={2}>再修改</Radio>
                            <Radio value={3}>驳回</Radio>
                        </Radio.Group>
                    </Modal>
                    <Modal title="修改记录" width="9.51rem" visible={this.state.modifyRecordVisible} onOk={()=>{
                        this.setState({
                            modifyRecordVisible:false
                        })
                    }} onCancel={()=>{
                        this.setState({
                            modifyRecordVisible:false
                        })
                    }}>
                        <div className="assistant-record-box">
                            <div className="basic">
                                <div className="id">
                                    材料编号:openct15646541
                                </div>
                                <div className="date">
                                    截止日期:2021-08-11
                                </div>
                            </div>
                            <Table 
                                dataSource={this.state.recordList} 
                                columns={this.recordListColums} size="small" 
                                pagination={false} 
                            />
                        </div>
                    </Modal>
                    <Modal title="上传材料" width="6.26rem" visible={this.state.upLoadMaterialVisible}  closable={false} footer={[(
                        <Button key="upload" type="primary" onClick={()=>{
                            message.success('上传成功，系统编码为系统编码为AK08198000')
                            this.setState({
                                upLoadMaterialVisible:false
                            })
                        }}>
                            上传
                        </Button>
                    )]}>
                        <div className="upload-material-box">
                            <Form
                                name="material"
                                labelCol={{ span: 6 }}
                                wrapperCol={{ span: 18 }}
                            >
                                <Form.Item
                                    label="材料内容"
                                    name="category-materials-box"
                                    rules={[{ required: true}]}
                                >
                                    <Radio.Group size="small">
                                        <Radio value={1}>新材料</Radio>
                                        <Radio value={2}>再修改材料</Radio>
                                        <Radio value={3}>驳回修改材料</Radio>
                                    </Radio.Group>
                                </Form.Item>

                                <Form.Item
                                    label="材料编号"
                                    name="material-id"
                                    className="material__id"
                                    rules={[{ required: true}]}
                                    extra="新材料自动生成材料编号，不可修改"
                                >
                                    <Input placeholder="AK9800000123" disabled/>
                                </Form.Item>

                                <Form.Item 
                                    label="上传"
                                    rules={[{ required: true}]}
                                    labelCol={{span:4}}
                                    required={true}
                                >
                                    <Upload.Dragger 
                                        name="files" 
                                        action="/upload.do"
                                        defaultFileList={this.state.defaultFileList}
                                    >
                                        <p className="ant-upload-drag-icon">
                                        <InboxOutlined />
                                        </p>
                                        <p className="ant-upload-text">Click or drag file to this area to upload</p>
                                        <p className="ant-upload-hint" style={{fontSize:'.1rem'}}>Support for a single or bulk upload. Strictly prohibit from uploading company data or other band files</p>
                                    </Upload.Dragger>

                                </Form.Item>
                            </Form>
                        </div>
                    </Modal>
                </div>
            </DocumentTitle>
        )
    }
}
