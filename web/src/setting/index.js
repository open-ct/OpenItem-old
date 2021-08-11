/*
 * @Author: your name
 * @Date: 2021-07-31 17:08:57
 * @LastEditTime: 2021-07-31 17:08:58
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web-client\src\setting\index.js
 */
export let ServerUrl = "http://localhost:8080";
export function initServerUrl() {
    const hostname = window.location.hostname;
    if (hostname === 'localhost') {
        ServerUrl = `http://${hostname}:8080`;
    }
}