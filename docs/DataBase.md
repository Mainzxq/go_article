数据库设计
===========

### 表结构设计

#### 用户表 users
<br /> 

##### Table: users  
-  <span style="color: green">id:</span> `int auto increase primary key unique,not null`
-   <span style="color: green">login_name:</span> `varchar(64) not null unique,`
-  <span style="color: green">pwd:</span>  ` text`  
> now we knew about what we need to do.

<br />

##### Table: article_info
-  <span style="color: green">id</span> `VARCHAR(64), primary key, not null`
-  <span style="color: green">author_id</span> `UNSIGNED int`
-  <span style="color: green">name</span>  `text`
-  <span style="color: green">display_ctime</span> `text`
-  <span style="color: green">create_time</span> `DATETIME`
