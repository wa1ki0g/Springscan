# Springscan

poc(持续更新中):
  Spring敏感端点路径扫描
  H2databaceRce
  SnakeYamlRce
  SpringCloudFunctionSpELRce
  SpringDataCommonsRce
  SpringFramework
  
# Spring敏感端点路径扫描:
  直接在go文件中定义了路径字典，爆破应该会快一点，网上能查到的全差不多都加进来了

<img width="1439" alt="image" src="https://user-images.githubusercontent.com/70200814/213861021-ad76377c-b28b-49b3-b316-e644b648301c.png">


# 用法

springscan version: 1.0.0
Usage:  [-uft] [-u url] [-f filename] [-t thread] [-h help]

Options:
  -f string
    	a target filename
  -h	this help
  -t int
    	thread Num (default 8)
  -u string
    	a target url(Please add http or https)
      

目前批量扫描这里还没写,只能扫单个

# 演示

<img width="1214" alt="image" src="https://user-images.githubusercontent.com/70200814/213861234-342a6e6e-d84a-427c-abb6-05c78b859669.png">

