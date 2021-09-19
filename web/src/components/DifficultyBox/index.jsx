import React, { Component } from 'react'
import './index.less'


export default class index extends Component {

    state={
        value:1
    }

    render() {
        return (
            <div className="difficulty-box" data-component="difficulty-box">
                <div className="difficulty-process-box">
                    {
                        [1,2,3,4,5].map(item=>(
                            <div className="value-item" key={Math.random(100)} style={{backgroundColor: item<=this.state.value?"#91D5FF":"#F5F5F5"}} onClick={()=>{
                                this.setState({
                                    value:item
                                })
                                this.props.onChange(item)
                            }}>
                                <div className="value">{item}</div>
                                <div className="center-cir"></div>
                            </div>
                        ))
                    }
                    <div className="difficulty-process-value" style={{width:`${25*(this.state.value-1)}%`}}></div>
                    <div className="selected-cir" style={{left:`${25*(this.state.value-1)}%`}}>
                        <div className="circle-1">
                            <div className="circle-2">
                                <div className="circle-3"></div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}
