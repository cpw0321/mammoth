---
title: mammoth v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.5"

---

# mammoth

> v1.0.0

# 认证服务

## POST 用户注册

POST /user/register

> Body 请求参数

```json
{
  "user_name": "admin",
  "password": "123456"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» user_name|body|string| 是 | 用户名|用户名|
|» password|body|string| 是 | 密码|密码|

> 返回示例

> 成功

```json
{
  "code": 0,
  "message": "success",
  "data": ""
}
```

```json
{
  "code": 500,
  "message": "用户已存在",
  "data": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|返回状态码|
|» message|string|true|none|返回错误信息|
|» data|string|true|none|返回值|

## POST 用户登录

POST /user/login

> Body 请求参数

```json
{
  "user_name": "admin",
  "password": "123456"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» user_name|body|string| 是 | 用户名|用户名|
|» password|body|string| 是 | 密码|密码|

> 返回示例

> 成功

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NTc1MTE1MjAsImlhdCI6MTY1NzUxMDA4MCwibmJmIjoxNjU3NTEwMDgwfQ.HUc7kmltFBx8VJ1kKgDgOatY3RKQyvu9porMAO0K3Vo"
  }
}
```

```json
{
  "code": 500,
  "message": "key is of invalid type",
  "data": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|返回状态码|
|» message|string|true|none|返回错误信息|
|» data|object|true|none|返回值|
|»» token|string|true|none|用户token|

## GET 获取角色列表 

GET /role/list

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|name|query|string| 否 ||角色名|
|order_field|query|string| 否 ||排序字段|
|order_type|query|string| 否 ||排序,降序/升序|
|page|query|string| 否 ||页数|
|page_size|query|string| 否 ||页大小|
|Authorization|header|string| 否 ||none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|返回状态码|
|» message|string|true|none|返回错误信息|
|» data|object|true|none|返回值|
|»» total|integer|true|none|总数|
|»» page|integer|true|none|页数|
|»» pageSize|integer|true|none|页大小|
|»» list|[object]|true|none|角色列表|
|»»» id|integer|false|none|用户id|
|»»» name|string|false|none|用户名|

## GET 获取用户列表

GET /user/list

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|user_name|query|string| 否 ||用户名|
|order_field|query|string| 否 ||排序字段|
|order_type|query|string| 否 ||排序,降序/升序|
|page|query|string| 否 ||页数|
|page_size|query|string| 否 ||页大小|
|Authorization|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 1,
    "page": 0,
    "pageSize": 0,
    "list": [
      {
        "id": 1,
        "user_name": "admin"
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|返回状态码|
|» message|string|true|none|返回错误信息|
|» data|object|true|none|返回值|
|»» total|integer|true|none|总数|
|»» page|integer|true|none|页数|
|»» pageSize|integer|true|none|页大小|
|»» list|[object]|true|none|用户列表|
|»»» id|integer|false|none|用户id|
|»»» user_name|string|false|none|用户名|

## POST 分配用户角色

POST /user/role

> Body 请求参数

```json
{
  "user_id": 0,
  "role_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» user_id|body|integer| 是 | 用户id|用户id|
|» role_id|body|integer| 是 | 角色id|角色id|

> 返回示例

> 成功

```json
{
  "code": 0,
  "message": "success",
  "data": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|none|
|» message|string|true|none|none|
|» data|string|true|none|none|

## POST 新建角色

POST /role/create

> Body 请求参数

```json
{
  "name": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 是 | 角色名称|角色名称|

> 返回示例

> 成功

```json
{
  "code": 0,
  "message": "success",
  "data": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|none|
|» message|string|true|none|none|
|» data|string|true|none|none|

# 数据模型

