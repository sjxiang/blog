
# 创建测试用户。执行以下命令：
curl -XPOST -H"Content-Type: application/json" -d'{"username":"sjxiang1997","password":"admin12345","nickname":"sjxiang1997","email":"459444344164@qq.com","phone":"1818888xxxx"}' http://127.0.0.1:4500/v1/users


# 用户登录 blog。执行以下命令：
curl -s -XPOST -H"Content-Type: application/json" -d'{"username":"sjxiang1997","password":"admin12345"}' http://127.0.0.1:4500/login

# 修改用户 sjxiang1997 的密码。执行以下命令：
curl -XPUT -H"Content-Type: application/json" -d'{"old_password":"admin12345","new_password":"admin45678"}' http://127.0.0.1:4500/v1/users/sjxiang1997/change-password

# 使用新密码登录 blog。执行以下命令：
curl -s -XPOST -H"Content-Type: application/json" -d'{"username":"sjxiang1997","password":"admin45678"}' http://127.0.0.1:4500/login


# 健康检查
curl http://127.0.0.1:4500/healthz

curl -XPUT http://127.0.0.1:4500/v1/users/:sjxiang1997