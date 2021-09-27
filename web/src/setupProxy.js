/*
 * @Author: your name
 * @Date: 2021-09-19 00:17:31
 * @LastEditTime: 2021-09-20 02:44:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \OpenItem\web\src\setupProxy.js
 */
const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function (app) {
    // /api 表示代理路径
    // target 表示目标服务器的地址
    app.use(createProxyMiddleware('/user', {target: "http://49.232.73.36:8081/review"}));
    app.use(createProxyMiddleware('/proj', {target: "http://49.232.73.36:8081/review"}));
    app.use(createProxyMiddleware('/file', {target: "http://49.232.73.36:8081/review"}));
    app.use(createProxyMiddleware('/query/assign', {target: "http://49.232.73.36:8081/review"}));
    app.use(createProxyMiddleware('/query/audit', {target: "http://49.232.73.36:8081/review"}));
    app.use(createProxyMiddleware('/query/file', {target: "http://49.232.73.36:8081/review"}));
    app.use(createProxyMiddleware('/query/proj', {target: "http://49.232.73.36:8081/review"}));
    app.use(createProxyMiddleware('/query/step', {target: "http://49.232.73.36:8081/review"}));
    app.use(createProxyMiddleware('/query/submit', {target: "http://49.232.73.36:8081/review"}));
    app.use(createProxyMiddleware('/query/user', {target: "http://49.232.73.36:8081/review"}));
    app.use(createProxyMiddleware('/question', {target: "http://49.232.73.36:8082/qbank"}));
    app.use(createProxyMiddleware('/testpaper', {target: "http://49.232.73.36:8082/qbank"}));
    app.use(createProxyMiddleware('/query/f_question', {target: "http://49.232.73.36:8082/qbank"}));
    app.use(createProxyMiddleware('/query/f_testpaper', {target: "http://49.232.73.36:8082/qbank"}));
    app.use(createProxyMiddleware('/query/t_question', {target: "http://49.232.73.36:8082/qbank"}));
    app.use(createProxyMiddleware('/query/t_testpaper', {target: "http://49.232.73.36:8082/qbank"}));
}