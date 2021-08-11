## 接口文档演示

**[2021-8-7] 注: 除基本的注册登录外, 其他接口操作需要在请求Header中加入登录时用户放回的token (目前不需要加token也可以返回正常响应, 后续加入认证功能后需要token才能执行操作).**

### 用户操作接口

**用户注册**
- 请求方式: `POST`
  - 请求URL: `http://49.232.73.36:8081/user/register`(不需要token)
- 请求格式: `JSON`
```json
{
    "name": "robot-8",
    "nick_name": "robot-8",
    "password": "123456",
    "email": "robot8@robot.com",
    "phone": "13111111118",
    "gender": 1,
    "degree": "本科",
    "position": "教师",
    "employer": "none",
    "major": "英语"
}
```
- 响应格式: `JSON`
```json
{
    "code": 2000,
    "message": "注册失败",
    "data": {
        "user_id": 0,
        "name": "",
        "description": "用户已存在（手机号或邮箱已被注册）"
    }
}
// or
{
    "code": 1000,
    "message": "注册成功",
    "data": {
        "user_id": 8,
        "name": "robot-9",
        "description": "ok"
    }
}
```

**用户登录**
- 请求方式: `POST`
  - 请求URL: `http://49.232.73.36:8081/user/login`(不需要token)
- 请求格式: `JSON`
```json
{
    "email": "robot8@robot.com",
    "password":"123456"
}
```
- 响应格式: `JSON`
```json
{
    "code": 1000,
    "message": "登录成功",
    "data": {
        "user_id": 7,
        "user_info": {
            "name": "robot-8",
            "email": "robot8@robot.com",
            "phone": "13111111118",
            "gender": 1,
            "degree": "本科",
            "position": "教师",
            "employer": "none",
            "major": "英语"
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjcsImV4cCI6MTYyODQwMDcxMSwiaWF0IjoxNjI4NDAwMTExLCJpc3MiOiJyb2JvdDhAcm9ib3QuY29tIn0.waGOCWn8nzvurBYxxv-L-zJZkQnxYp1DgZcPXPfnsFA",
        "description": "ok"
    }
}
// or
{
    "code": 2000,
    "message": "登陆失败",
    "data": {
        "user_id": 7,
        "user_info": {
            "name": "",
            "email": "",
            "phone": "",
            "gender": 0,
            "degree": "",
            "position": "",
            "employer": "",
            "major": ""
        },
        "token": "",
        "description": "认证错误,请重试"
    }
}
```

**用户信息更改**
- 请求方式: `POST`
  - 请求URL: `http://49.232.73.36:8081/user/updateInfo` (需要token)
- 请求格式: `JSON`
```json
{
    "id": 5,
    "name": "robot-5",
    "nick_name": "robot-5",
    "email": "robot5@robot.com",
    "phone": "13111111115",
    "gender": 2,
    "degree": "本科",
    "position": "助理",
    "employer": "none",
    "major": "物理"
}
```
- 响应格式: `JSON`
```json
{
    "code": 1000,
    "message": "信息更新成功",
    "data": {
        "user_id": 5,
        "name": "robot-5",
        "description": "ok"
    }
}
// or
{
    "code": 2000,
    "message": "信息更新失败",
    "data": {
        "user_id": 5,
        "name": "robot-5",
        "description": "信息已被注册，请重试"
    }
}
```

**用户密码更改**
- 请求方式: `POST`
  - 请求URL: `http://49.232.73.36:8081/user/updatePwd` (需要token)
- 请求格式: `JSON`
```json
{
    "id": 6,
    "old_password":"abcdef",
    "new_password":"123456"
}
```
- 响应格式: `JSON`
```json
{
    "code": 2000,
    "message": "信息更新失败",
    "data": {
        "user_id": 6,
        "name": "robot-6",
        "description": "认证失败,请重试"
    }
}
// or
{
    "code": 1000,
    "message": "信息更新成功",
    "data": {
        "user_id": 6,
        "name": "robot-6",
        "description": "ok"
    }
}
```

**使用id获取用户信息**
- 请求方式: `GET`
  - 请求URL: `http://49.232.73.36:8081/user/getInfo?user_id=5`(需要token)
- 请求格式: `URL-Params`(见上方url参数)
- 响应格式: `JSON`
```json
{
    "code": 1000,
    "message": "获取信息成功",
    "data": {
        "user_id": 5,
        "user_info": {
            "name": "robot-5",
            "email": "robot5@robot.com",
            "phone": "13111111115",
            "gender": 2,
            "degree": "本科",
            "position": "助理",
            "employer": "none",
            "major": "物理"
        },
        "description": "ok"
    }
}
```

**用户退出登录** *调试中*
- 请求方式: `GET`
- 请求格式: `URL-Params`
```json
{}
```
- 响应格式: `JSON`
```json
{}
```

**删除用户**
- 请求方式: `GET`
  - 请求URL: `http://49.232.73.36:8081/user/delete?user_id=2`(需要token)
- 请求格式: `URL-Params`(见上方URL参数)
- 响应格式: `JSON`
```json
{
    "code": 2000,
    "message": "用户删除失败",
    "data": {
        "user_id": 2,
        "name": "",
        "description": "对应信息不存在"
    }
}
// or
{
    "code": 1000,
    "message": "用户删除成功",
    "data": {
        "user_id": 1,
        "name": "",
        "description": "ok"
    }
}
```



### 项目操作接口

### 文件操作接口

**文件上传**
- 请求方式: `POST`
  - 请求URL: `http://49.232.73.36:8081/file/upload`
- 请求格式: `form-data`
```
key: file_name
value: file
```
- 响应格式: `JSON`
```json
{
    "code": 1000,
    "message": "上传文件成功",
    "data": {
        "file_id": "70050161-6d4a-40b8-b2d5-49e0f7b9d497",
        "file_name": "横向多种条形图.png",
        "file_type": ".png",
        "file_path": "static/files/2021-8-7/70050161-6d4a-40b8-b2d5-49e0f7b9d497.png",
        "description": "ok"
    }
}
```

**文件下载**
- 请求方式: `GET`
  - 请求URL: `http://49.232.73.36:8081/file/download?file_id=18541826-bd1c-47ba-a7d8-a9ba7497a152`
- 请求格式: `URL-Params`(上面URL的参数是文件的uuid)
- 响应格式: `JSON` or `FILE`(若请求成功直接下载)
```json
// 请求失败情况
{
    "code": 2000,
    "message": "获取文件信息失败",
    "data": {
        "file_id": "18541826-bd1c-47ba-a7d8-a9ba7497a153",
        "file_name": "",
        "file_type": "",
        "file_tags": null,
        "file_path": "",
        "file_description": "",
        "file_owner": "",
        "description": "文件操作失败"
    }
}
```

**获取文件信息**
- 请求方式: `GET`
  - 请求URL: `http://49.232.73.36:8081/file/getInfo?file_id=aad492b1-9cfb-4866-946c-790204aadd4e`
- 请求格式: `URL-Params`(见上方请求URL)
- 响应格式: `JSON` 
```json
{
    "code": 1000,
    "message": "获取文件信息成功",
    "data": {
        "file_id": "aad492b1-9cfb-4866-946c-790204aadd4e",
        "file_name": "横向多种条形图.png",
        "file_type": ".png",
        "file_tags": [],
        "file_path": "static/files/2021-8-8/aad492b1-9cfb-4866-946c-790204aadd4e.png",
        "file_description": "",
        "file_owner": "",
        "description": "ok"
    }
}
```

**文件搜索** *完善中*