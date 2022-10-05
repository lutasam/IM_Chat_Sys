# chat
聊天程序后台代码  the backend code of my chat program
## 整体架构  
```
--biz 业务逻辑都存放在这  
  --bo          封装request和response结构体
  --common      常量
  --handler     controller
  --service  
  --dal  
  --model       放数据结构
  --repository  数据库连接操作
  --util        工具类
  --middleware  中间件
--conf          配置
--main.go  
--router.go     路由
```
## 如何启动
main.go
