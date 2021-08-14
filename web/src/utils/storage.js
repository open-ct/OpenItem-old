/*
 * @Author: your name
 * @Date: 2021-08-14 00:09:49
 * @LastEditTime: 2021-08-14 04:22:15
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web\src\utils\storage.js
 */

function getStorage(name) {
    return localStorage.getItem(name)===null?null:JSON.parse(localStorage.getItem(name))
}

function setStorage(name,value) {
    localStorage.setItem(name,JSON.stringify(value))
}

function checkStorage(name) {
    return localStorage.getItem(name)!==null
}

function delStorage(name) {
    localStorage.removeItem(name)
}

module.exports = {
    getStorage,
    setStorage,
    checkStorage,
    delStorage
}