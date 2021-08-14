/*
 * @Author: your name
 * @Date: 2021-08-12 00:25:43
 * @LastEditTime: 2021-08-12 23:18:32
 * @LastEditors: Please set LastEditors
 * @Description: 操作Cookie
 * @FilePath: \OpenItem\web\src\utils\cookies.js
 */






/**
 * @description: 存值
 * @param {*} name
 * @param {*} value
 * @return {*}
 */
function setCookie(name,value){
    var d = new Date();
    d.setTime(d.getTime()+(60*60*1000)); //保存一个小时
    var expires = "expires="+d.toGMTString();
    document.cookie = name+"="+value+"; "+expires;
}


/**
 * @description: 获取指定Cookie的值
 * @param {*} name
 * @return {*}
 */
function getCookie(name){
    name = name + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i].trim();
        if (c.indexOf(name)===0) { return c.substring(name.length,c.length); }
    }
    return null;
}



/**
 * @description: 检查指定Cookie是否存在
 * @param {*} name Cookie名称
 * @return {*}
 */
function checkCookie(name){
    if(getCookie(name)){
      return true
    }else{
      return false
    }
}


/**
 * @description: 删除指定Cookie
 * @param {*} name Cookie名称
 * @return {*}
 */
function delCookie(name){
    document.cookie = name+"=;expires="+(new Date(0)).toGMTString();
}

module.exports = {
    setCookie,
    getCookie,
    checkCookie,
    delCookie
}