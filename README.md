<!--
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-05-17 23:01:27
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 12:53:29
-->
# Gin+Vue实现前后端分离的秒杀商城  

## 微服务拆分  

[用户微服务](./zhiliao_user_srv/)
[商品微服务](./zhiliao_product_srv/)
[秒杀微服务](./zhiliao_seckill_srv/)  

## 日志級別设置  

```json
    "trace":logrus.TraceLevel,
    "debug":logrus.DebugLevel,
    "info":logrus.InfoLevel,
    "warn":logrus.WarnLevel,
    "error":logrus.ErrorLevel,
    "fatal":logrus.FatalLevel,
    "panic":logrus.PanicLevel,
```

## 运行  

`consul`  
```
$ consul agent -dev
```