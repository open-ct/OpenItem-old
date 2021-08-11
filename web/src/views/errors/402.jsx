import React, { Component } from 'react'
import DocumentTitle from 'react-document-title'
import Rockets from '../../asset/images/Rockets.png'
import Cloude from '../../asset/images/Cloud.png'
import Unhapply from '../../asset/images/Unhapply.png'
import './index.less'

export default class index extends Component {
    render() {
        return (
            <DocumentTitle title="402:拒绝访问">
                <div className="error_page__404" data-component="error_page__404">
                    <img src={Rockets} alt="火箭图" className="pattern__1 pattern" />
                    <img src={Cloude} alt="云" className="pattern__4 pattern"/>
                    <img src={Unhapply} alt="不高兴" className="pattern__5 pattern"/>
                    <div className="line__5 pattern line"></div>
                    <span className="title pattern">402</span>
                    <p className="describe pattern">
                        错误原因错误原因错误原因错误原因
                        错误原因错误原因错误原因错误原因
                        错误原因错误原因错误原因错误原因
                    </p>
                </div>
            </DocumentTitle>
            
        )
    }
}