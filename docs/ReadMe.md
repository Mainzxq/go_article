
第一批API设计:
==========

### 用户管理设计

1. 创建用户*（注册)  

| URL | Method | SC  |
| :---: | :---: | :---: |
| /user | POST | 201, 400, 500 |

2. 用户登录： url:/user/:username Method: POST, SC:200, 400, 500

| URL | Method | SC  |
| :---: | :---:  | :---: |
|  /user/:username | POST  | 200, 400, 500 |

3. 获取用户的基本信息：url:/user/:username Method: GET, SC: 200, 400, 401(并没有验证),403(无权限), 500,

| URL | Method | SC  |
| :---: | :---:  | :---: |
|  /user/:username | GET  | 200, 400, 401(并没有验证),403(无权限), 500 |

4. 用户注销: url:/user/:username Method:DELETE, SC: 204, 400,401, 403,500

| URL | Method | SC  |
| :---: | :---:  | :---: |
|  /user/:username | DELETE  | 204, 400,401, 403,500 |

***
### 资源设计
1. 列出文章列表和所属资源 url: /user/:username/articles Method: GET, SC: 200, 400, 500
2. 查看详细内容  url: /user/:username/articles/:atcl_id Method:GET SC: 200, 400, 500
3. 删除资源详细内容 url: /user/:username/articles/:actl_id Method:DELETE SC: 204, 400, 401, 403, 500
###### 管理员、发表者拥有删除权限

### 评论设计
1. 显示评论 

| URL | Method | SC  |
| :---: | :---:  | :---: |
|  /articles/:atcl_id/comments | GET  | 200, 400, 500 |

- 这一点需要注意
  
2. 发表评论:

| URL | Method | SC  |
| :---: | :---:  | :---: |
|  /articles/:atcl_id/comments | POST  | 201, 400, 500  |

3. 删除评论 url: /articles/:atcl_id/comments/:comment_id Method: DELETE SC: 204, 400, 401, 403, 500
###### 管理员、发表者拥有删除权限