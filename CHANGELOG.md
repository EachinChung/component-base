## v0.6.0 (2022-07-30)

### Refactor

- 重构为范型

### Fix

- **utils**: GetInstanceID 前缀为空时, 不再带有 -
- **middleware**: 去除默认的 AllowOriginFunc
- **options**: 去除不知所谓的默认值

### Feat

- **utils**: 获取本机外网IP

## v0.5.0 (2022-07-27)

### Feat

- 更为合理的特殊字符检查
- **options**: 添加默认值
- instance id 最小设置为 6 位
- 随机验证码

## v0.4.0 (2022-07-26)

### Fix

- 去除 WithUnscoped 多余的参数

### Refactor

- **utils**: 包名不再和标准库冲突

### Feat

- **middleware**: JWT 将自动在上下文存储 eid

## v0.3.0 (2022-07-24)

### Feat

- **middleware**: 正确的 JWT 命名
- **options**: casbin 权限管理
- **middleware**: 去除与 cors 重复的 options
- **core**: response 有更多的 option
- postgres

## v0.2.1 (2022-07-18)

### Fix

- 修正错误的 log 版本

## v0.2.0 (2022-07-18)

### Feat

- db 选项模式
- jwt 中间件

## v0.1.0 (2022-07-16)

### Feat

- 一些基础公共组件
