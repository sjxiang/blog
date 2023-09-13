
# Blog，博客


# Feature

```text

用户管理
    用户注册
    用户登录
    获取用户列表（管理员权限）
    获取用户详情
    更新用户信息
    修改用户密码
    注销用户
    

博客管理
    创建博客
    获取博客列表
    获取博客详情
    更新博客内容
    删除博客
    批量删除博客
```

db2struct --gorm --no-json -H 172.21.0.3 -d blog -t user --package model --struct UserM -u root -p '123456' --target=user.go
