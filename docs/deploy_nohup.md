### nohup 开发测试环境部署方案

在项目开发、测试环境，我们需要的只是快速部署、测试项目功能，因此该方案是最简单、也是最适合开发调试环境的首选方案.


```code 

# 首先编译出可执行文件
# 将可执行文件 + config目录  + public 目录 + storage 目录 合计4项全部复制在服务器。
# 进入可执行文件目录执行
 nohup  可执行文件名  & 
 
 
 # 版本更新
 
 # 杀死旧进程
 kill  -9  （进程pid）
 # 重新启动进程
 nohup 可执行文件名  &  

```


