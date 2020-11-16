## Counter-Strike Online 2 Server 

[![Build status](https://ci.appveyor.com/api/projects/status/a4pj1il9li5s08k5?svg=true)](https://ci.appveyor.com/project/KouKouChan/cso2-server)
[![](https://img.shields.io/badge/license-MIT-green)](./LICENSE)
[![](https://img.shields.io/badge/version-v0.3.15-blue)](https://github.com/KouKouChan/CSO2-Server/releases)

[English](./README.en.md) | [CodePage](./CodePage.md)

*声明：Counter-Strike Online 2 归 NEXON 所有 ，本程序仅用于学习之用*

### 一、介绍

CSOL2 服务器 v0.3.15

数据库：SQLite    ，可使用相应工具打开.db文件

用于 **2017年国服 Counter-Strike Online 2** 和 **2018年韩服 Counter-Strike Online 2**  

目前客户端请使用L-Leite的启动器,**韩服端竞技模式有问题的请下载最新的汉化包**。

基于L-Leite的[cso2-master-server](https://github.com/L-Leite/cso2-master-server)。

如果大家有什么建议或问题，欢迎提出。

欢迎大家帮忙本地化文件，具体见下方第七部分。

### 二、项目计划

    1.先实现基本的游戏游玩功能和联机功能 √
    2.重构代码 √
    3.主要实现偏单机的功能 ...(进行中)

### 三、基本已完成的功能

    登录、频道、房间、仓库、UDP、角色战绩(游戏结果界面)、数据库、个人信息、聊天、命令行和数据库、新手教程、开箱子

### 四、正在编写的功能

​    

### 五、已知问题

    1.房主离开后，其余玩家会卡住直到炸出房间
    2.服务端目前还未适配2017国服端的部分数据包，所以结算界面数据显示存在错误

### 六、大概已修复的问题

    1.主机开始游戏后，其他玩家不能加入，显示超时。需要和主机一起开始游戏才能加入。
    2.房间列表显示的房间信息及状态不准确，待刷新
    3.玩家仓库数据不准确
    4.房间列表的房主名字显示中文乱码
    5.每局结束显示的角色战绩可能存在一些问题
    6.房间无法加密码
    7.可能存在多协程共享变量不安全的问题
    8.由于房间用户与主管理器的用户重复，可能造成性能浪费
    9.房间ID和房间NUM在多频道下可能冲突（虽然目前是单频道）
    10.竞技模式下默认添加电脑

### 七、部分数据本地化方法

```
1.打开server.conf
2.修改LocaleFile选项，将其改为你的语种的文件名，比如 zh-cn.ini
3.进入 CSO2-Server\locales\ 目录
4.创建相同文件名文件，比如 zh-cn.ini
5.根据 zh-cn.ini 中的内容相应修改
```

### 八、客户端下载

  [点击2018韩服端下载](https://pan.baidu.com/s/1NGHisLeTB1nXH4zCtR6FSA) 提取码：5vca  

  [韩服端汉化包GoogleDrive](https://drive.google.com/file/d/1aaoKSBrAKgO30w-BCf1VJG6n6PUiS-88/view?usp=sharing)

  [点击2017年国服端下载](https://pan.baidu.com/s/1tTtks0fwROk0WUueC2gnOQ)  提取码：o9hd

  [单独启动器下载，如果你已有客户端](https://pan.baidu.com/s/1QGyRmjw24eJ5ycrFjorv_g)  提取码：amys

### 九、使用方法

1.需要有CSOL2客户端，同时使用第三方启动器

2.进入本项目的release页面下载最新版本的程序（ https://github.com/KouKouChan/CSO2-Server/releases ）

3 .(国服端请跳过该步骤) 建立bat文件，和游戏的bin目录同级，里面写入：

```shell
START ./bin/launcher.exe -masterip IP地址 -enablecustom -username 用户名 -password 密码
```

4.IP地址指的是你的服务端IP，如果是本地那么就填127.0.0.1（仅单人情况下），如果你要连接局域网别人的服务端那么就填别人的IP地址，如果你安装了汉化包，也可以再加上以下语句：

```shell
-lang schinese
```

5.先运行本项目的exe文件启动服务器，然后打开bat文件启动客户端即可**（国服端可能启动稍慢）**

- 从网盘里面下载得到的start-cso2.bat文件需要修改里面的IP地址和用户名！
- 如果你需要注册，请修改server.conf文件，将EnableRegister值修改为1，然后你可以使用浏览器打开 localhost:1314 来注册，默认注册端口为1314。
- 如果你想开启邮箱注册，那么你需要一个邮箱账号并且申请到了密钥，将密钥填入配置文件，同时开启EnableMail。
- 如果你是和别人联机玩，那么即使你的电脑运行着服务端也**不能**在bat文件里填127.0.0.1，不然对方无法通过你的ip连接你。

### 十、Console使用方法

CSO2-Server自带管理员功能，可通过命令行参数打开console功能管理服务器，前提需要服务器已经在运行。

1.运行服务器。

2.使用local-console.bat连接本地服务器或者使用如下命令连接服务器：

```
CSO2-Server.exe -console -ip YOURIP -port YOURPORT -username GMNAME -password PASSWORD
```

默认参数如下:

```
Usage of CSO2-Server.exe:
  -ip string
        主机名，默认为localhost (default "localhost")
  -password string
        密码，默认为cso2server123 (default "cso2server123")
  -port string
        端口号，默认为1315 (default "1315")
  -username string
        账号，默认为admin (default "admin")
```

3.连接成功后可以使用命令管理服务器了，你可以踢出玩家，或者给予玩家物品等。

### 十一、自定义文件方法

1.下载CSOL2解包工具，[点击这里下载](https://pan.baidu.com/s/14q1SoIdHwp1casMWG2OS-w) 提取码：41bs

2.解压后，打开工具，点击左上角File选项，点击Open folder，选中csol2的data文件夹即可

3.解压你需要的文件，并且将解压后的文件按你的想法进行修改

4.将文件放入csol2目录的custom文件夹下，打开游戏

### 十二、Docker下使用方法

1.首先你需要拥有Docker,请下载并安装Docker,同时配置好Docker,比如Docker源

2.输入以下命令拉取最新版的服务端:

```shell
docker pull koukouchan/cso2-server:latest
```

3.运行服务端

```shell
docker run -p 30001:30001 -p 30002:30002 -p 1314:1314 -p 1315:1315 koukouchan/cso2-server:latest
```

4.接下来打开客户端，连接服务器

### 十三、编译环境

*Go 1.15.3*

当你要架设局域网或外网时，请打开防火墙的端口。30001-TCP类型端口、30002-UDP类型端口

貌似建立互联网服务器需要双方玩家都能内网穿透，实测局域网能够连接，互联网无法房间内加入主机，可能需要架设虚拟局域网。

### 十四、编译方法

```shell
1. 在shell中执行 go get github.com/KouKouChan/CSO2-Server
2. 进入目录
3. 执行命令 go build
4. 运行生成的可执行文件即可
```

### 十五、Docker下编译方法

1.首先你需要拥有Docker,请下载并安装Docker和Git,同时配置好Docker,比如Docker源,使用如下命令安装git:

```
yum install git     #centos
或
apt-get install git #ubuntu
```

2.在终端下输入以下命令:

```shell
git clone https://github.com/KouKouChan/CSO2-Server
cd CSO2-Server
docker build -t cso2-server .
```

3.在第2步后，如果运行正常，会显示所有步骤都运行完毕。接下来是运行服务端，为了能够让游戏和Docker容器里面的服务端相连，你需要打开相应的端口映射，使用以下命令运行：

```shell
docker run -p 30001:30001 -p 30002:30002 -p 1314:1314 -p 1315:1315 cso2-server
```

4.接下来打开客户端，连接服务器

5.建议关闭docker时将容器数据导出，否则将丢失玩家数据！

### 十六、图片

![Image](./photos/main.png)

![Image](./photos/intro.png)

![Image](./photos/channel.png)

![Image](./photos/ingame.jpg)

![Image](./photos/result.jpg)
