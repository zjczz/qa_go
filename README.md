# qa_go

企业选题
qa_go
  --api            对外api层
    --v1           api版本
  --auth           授权jwt
  --cache redis    配置相关  
  --conf           读取本地环境变量
  --middleware     常用跨域、jwt、日志等中间件
  --model          数据库实体以及组合使用的gorm操作
  --routes         路由汇总
  --serializer     将实体映射成不同的viewmodel，以及常用的响应信息 
  --service        服务层封装了model层的方法供api层调用
  --utils          常用工具类
  --example.env    环境变量配置