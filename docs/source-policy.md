# images unauthorized source Specification
- sourceDefaultPolicy 将定义你默认是否允许所有来源镜像
```
    当 sourceAllowPolicy/sourceRejectPolicy 存在配置
    将互相制约
```
- sourceAllowPolicy 授信的来源
- sourceRejectPolicy 拒绝的来源

- sourceAllowPolicy/sourceRejectPolicy 同时存在相同来源，拒绝权限更高
