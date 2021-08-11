import React, { Component } from 'react'
import { Table, Pagination } from 'antd'
import ModulaCard from '../ModulaCard'
import './index.less'

export default class index extends Component {
    constructor(props){
        super(props)
        console.log(props)
    }
    render() {
        return (
            <ModulaCard title={this.props.title}>
                <div className="data-table-box" data-component="data-table-box">
                    <Table 
                        dataSource={this.props.data} 
                        columns={this.props.colums} 
                        size="small" 
                        pagination={false} 
                        scroll={{ y: '5.8rem'}}
                    />
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
        )
    }
}
