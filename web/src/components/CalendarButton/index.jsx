/*
 * @Author: your name
 * @Date: 2021-07-31 16:42:56
 * @LastEditTime: 2021-08-12 01:22:48
 * @LastEditors: Please set LastEditors
 * @Description: 抽离出的日历按钮组件
 * @FilePath: \OpenItem\web-client\src\index.js
 */
import React, { Component } from 'react'
import { Button, Calendar, Select, Row, Col,  } from 'antd'
import './index.less'

/**
 * @description: 参数
 * @icon 按钮icon
 * @label 按钮名称
 */

export default class index extends Component {

    state = {
        calendarVisible:false
    }

    render() {
        return (
            <div className="calendar-button-box" data-component="calendar-button-box">
                <Button size="small" icon={this.props.icon} onClick={()=>{
                    this.setState({
                        calendarVisible:true
                    })
                }}>{this.props.label}</Button>
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
                        onSelect={(_date)=>{
                            let date = new Date(_date._d)
                            this.setState({
                                calendarVisible:false
                            })
                            this.props.onDateChange(`${date.getFullYear()}-${(date.getMonth()+1).toString().padStart(2,'0')}-${date.getDate().toString().padStart(2,'0')}`)
                        }}
                    />
                </div>
            </div>
        )
    }
}

