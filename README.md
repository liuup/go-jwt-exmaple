# go-jwt-example

一个使用jwt对用户登录功能进行校验的demo  

主要功能：  
1. 路由分组，将公共路由和私有路由进行分离
2. jwt登录校验

不足：  
1. 登录接口未做限制，可能会被操作无限生成jwt token
