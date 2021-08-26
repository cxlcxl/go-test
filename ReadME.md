## GO 入门学习
 使用了 `GinSkeleton` 为基础开发，所有配置介绍可以在 [ginskeleton](https://gitee.com/daitougege/gin-skeleton-admin-backend) 查询

### 快速上手 
- 1.go语言环境配置
```code  
// 1.安装的go语言版本必须>=1.15 .

// 2.配置go包的代理，打开你的 goland 终端并执行以下命令（windwos系统）
    // 其他操作系统自行参见：https://goproxy.cn  
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct

// 3.下载本项目依赖库  
    使用 goland(>=2019.3版本) 打开本项目，打开 goland 底部的 Terminal ,执行  go mod tidy 下载本项目依赖库  
```