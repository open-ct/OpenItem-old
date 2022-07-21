import React, { Component } from 'react'
import './index.less'

export default class index extends Component {
    render() {
        return (
            <div className="modula-card-box" data-component="modula-card-box">
                <div className="title">
                    <div className="icon"></div>
                    <span>{this.props.title}</span>
                    <div className="right">
                        {this.props.right}
                    </div>
                </div>
                {this.props.children}
            </div>
        )
    }
}
