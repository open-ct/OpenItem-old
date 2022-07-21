import React, { Component } from 'react'
import DocumentTitle from 'react-document-title'
import Rockets from '../../asset/images/Rockets.png'
import CryingFace from '../../asset/images/Crying_face.png'
import FluExpression from '../../asset/images/Flu_expression.png'
import './index.less'

export default class index extends Component {
    render() {
        return (
            <DocumentTitle title="404:页面丢失了">
                <div className="error_page__404" data-component="error_page__404">
                    <img src={Rockets} alt="火箭图" className="pattern__1 pattern" />
                    <img src={CryingFace} alt="哭脸图" className="pattern__2 pattern"/>
                    <img src={FluExpression} alt="尴尬图" className="pattern__3 pattern"/>
                    <div className="line__1 pattern line"></div>
                    <div className="line__2 pattern line"></div>
                    <div className="line__3 pattern line"></div>
                    <div className="line__4 pattern line"></div>
                    <span className="title pattern">404</span>
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