/*
 * @Author: stacker
 * @Date: 2021-07-31 16:46:43
 * @LastEditTime: 2021-07-31 16:46:43
 * @LastEditors: stacker
 * @Description: less相关配置
 * @FilePath: \OpenItem\web-client\config-overrides.js
 */
const { override, addLessLoader } = require("customize-cra");
module.exports = override(
  addLessLoader({
    javascriptEnabled: true,
  })
);