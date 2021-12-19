# 规范与建议
## Commit 规范

### 标准格式
```
Action(modname/package): (add|update|..all action you done)
```

### example

```
update(wechatPong): add template message test 
```

## 接口标准

1. URL中字母全部小写
2. 如果有单词拼接，使用中划线‘-’，不使用下划线‘_’ (在搜索引擎中，把中划线当做空格处理，而下划线是被忽略的。使用中划线是对搜索引擎友好的写法)
3. 资源必须采用资源名词的复数形式
4. 层级尽量不超过三层
5. 参数不允许超过3个
6. 如果是内容资源的URL，不允许以参数方式显示
