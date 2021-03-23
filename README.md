# PEF (Performance Testing Framework)

## 0. 简介
PEF是一个性能测试框架，确切地说，它是一个[locust](https://locust.io)的slave进程，用于实际的压测执行。它使用[golang](https://golang.org)编写，基于[boomer](https://github.com/myzhan/boomer)与locust master交互。
locust master节点用于监控实时压测结果
## 1. 编译、安装

以下步骤针对Linux和Mac用户，Windows用户请使用go build自行编译，pef的目标文件只有一个二进制文件，在command line中运行编译得到的文件即可。

### 1.1. 编译
```sh
# cd to root directory of pef
# export GOOS=linux # windows或者macos，交叉编译的时候需要用到，如果不指定，则会用当前运行的OS类型
make  # compile and generate "./pef" binary
make package  # generate rpm package，需要安装fpm工具
```

### 1.2. 安装
在所有施压机上都需要安装pef的rpm包
```sh
# 统一测试环境中，如果存在yum.dx.corp的yum源，直接用yum安装
yum install -y pef
```

或
```sh
# 拷贝pef-xxx.rpm到环境中
yum install ./pef-*.rpm
```

安装完成后，可以在/usr/local/bin目录下找到pef文件。

### 1.3. 安装locust

要求本地已安装有docker环境

```sh
# 本地可以访问harbor
docker pull harbor.dx-corp.top/aladdin/locust
```

或
```sh
# 导入包含locust的tar.gz文件
docker load < locust.tar.gz
```

## 2.使用

```sh
pef [general|engine|captcha|constid|indicator|constid-dubbo|dns] --help
```

### 2.1. 启动locust
只需要在一台施压机上启动locust即可，locust master本身不执行压测，只有实时监控效果
```sh
# 启动locust master
docker run --restart always --name locust --network host -d harbor.dx-corp.top/aladdin/locust
```

### 2.2. 启动pef分布式压测
##### 在所有施压机上都需要启动pef，每个施压机作为一个slave节点用于实际的压测执行（机器资源充足条件下推荐）
##### 或者在同一台施压机启动多个pef进程，每个进程作为一个slave节点用于实际的压测执行，进程数<=CPU总数（机器资源不充足条件下推荐）

#### 2.2.1. engine的压测

```sh
pef engine -master-host 10.1.2.14 -host 10.1.2.12:7776 -app-key f3204b7c2e89ca19f757576965ffb066 -app-secret 9f044ec59fca9bd9e2d812d45f4bac26 -event-codes test_code1,test_code2 -field ip -field phone_number -field user_id
```

* -app-key: 被测应用的app key
* -app-secret: 被测应用的app secret
* -event-codes: 被测事件的code,多个code用英文逗号分隔,多个事件必须属于同一个产品
* -field: 随机发送的字段名，上述例子中随机发送的请求包括3个字段ip、phone_number和user_id
* -host: 被测程序的IP:PORT
* -master-host: locust master的IP地址


从文件读取压测数据

```sh
pef engine -master-host 10.1.2.14 -host 10.1.2.12:7776 -app-key f3204b7c2e89ca19f757576965ffb066 -app-secret 9f044ec59fca9bd9e2d812d45f4bac26 -event-codes test_code1,test_code2 -data-source test_data.csv
```

* -data-source: 测试数据源，random或者某个具体的文件路径
* -run-once: 是否只执行一次测试数据，仅当数据源为文件的时候有效，默认将循环测试文件中的数据

#### 2.2.2. captcha的压测

```sh
pef captcha -ak b09ca8827bac117cf192b378b7164029 -c 5c2498d5ZbNZJ1vTzfyZ32YziDFliQaqEbwBh8o1 -iface all -host 10.1.2.7:9091 -master-host 10.1.2.14
```

* -ak: 被测应用的app key
* -c: 有效的constid，能够通过验证码验证
* -iface: 被测的接口，a|v1|tokenVerify|all，选择all的时候会在请求a接口后，拿到a接口的返回值接着请求v1接口,v1接口返回状态码为200时，请求tokenVerify接口
* -host: 被测程序的IP:PORT
* -master-host: locust master的IP地址

#### 2.2.3. constid的压测

```sh
pef constid -iface all -app-key 90b3a12f68187787bc8d52cf7e7b6366 -app-secret f5e164094e1a9c0531da4aff90fdc46b -host 10.0.0.203:8090 -type random -master-host 10.1.2.14
```

* -iface: 测试接口，token、verify和all，默认为token，all表示先请求token接口，在请求getTokenInfo接口
* -app-key: 被测应用的app key
* -app-secret: 被测应用的app secret
* -type: 客户端类型，web、ios、android或者random
* -host: 被测程序的IP:PORT
* -master-host: locust master的IP地址

#### 2.2.4. constid-dubbo的压测

```sh
pef constid-dubbo -host 10.0.0.203:36502 -master-host 10.1.2.14
```

该命令将会生成随机的token，通过dubbo请求constid服务的getDeviceInfo接口。

* -host: 被测程序的dubbo服务的IP:PORT
* -master-host: locust master的IP地址

#### 2.2.5. indicator的压测

```sh
pef indicator -host 10.0.0.203:31000 -iface all -dc data_source_1 -ic ic_1 -ic ic_2 -field ip -field phone_number -master-host 10.1.2.14
```

该命令会调用indicator服务的dubbo接口，发送processMessage和mget消息。

* -iface: 测试接口，processMessage、mget和all，默认为all，all表示先请求mget，再请求processMessage
* -dc: 被测数据源code
* -ic: 被测指标code
* -field: 被测字段，这里指定了ip和phone_number两个字段
* -host: 被测程序的IP:PORT
* -master-host: locust master的IP地址

#### 2.2.6. DNS的压测

```sh
pef dns -type A -protocol udp -record test.dx.corp -host 10.1.2.14:53 -master-host 10.1.2.14
```

* -type: DNS类型，默认为A
* -protocol: 协议类型，默认为udp
* -record: 被测DNS记录
* -host: 被测程序的IP:PORT
* -master-host: locust master的IP地址

### 2.2.7. 通用接口的压测

```sh
pef general -host 127.0.0.1:8080 -api /api/test -master-host 10.1.2.39 -data-source test.data
```

* -api: 被测程序的接口的url地址
* -host: 被测程序的IP:PORT
* -master-host: locust master的IP地址
* -data-source: 测试数据源，random或者某个具体的文件路径
* -run-once: 是否只执行一次测试数据，仅当数据源为文件的时候有效，默认将循环测试文件中的数据

### 2.3. 开始压测

打开浏览器，输入地址http://10.1.2.14:8089 ，把10.1.2.14改成locustmaster对应的IP地址，查看slave连接情况，点击New test，设置用户数，开始测试。

注意观察QPS、RT、错误率等数据，同时注意监控相关服务器的基础性能和连接数等数据。

## 3. 资源监控

### 3.1 falcon-plus工具的安装
falcon-plus的工具由golang编写，以rpm形式发布，该版本在官方开源的open-falcon基础上修改而成。

#### 3.1.1 falcon-graph、falcon-api、falcon-transfer的安装

选择一台服务器作为监控的中心服务器，我们会在上面安装falcon-graph、falcon-api、falcon-transfer，将对应的rpm包拷贝到该服务器上。

将两个sql文件拷贝到该服务器上，在对应的mysql上创建对应的数据库和表
```sh
mysql -h [mysql数据库ip地址] -u root -p<密码> -e "source 4_graph-db-schema.sql"
mysql -h [mysql数据库ip地址] -u root -p<密码> -e "source 5_alarms-db-schema.sql"
```

安装falcon-graph，falcon-transfer，falcon-api
```sh
yum install ./falcon-*.rpm
```

修改/data/falcon-graph/cfg.json和/data/falcon-api/cfg.json，修改为正确的数据库连接地址。


启动服务
```sh
systemctl enable falcon-graph falcon-transfer falcon-api
systemctl start falcon-graph falcon-transfer falcon-api
```

#### 3.1.2 falcon-agent的安装

所有的服务器都需要安装falcon-agent，将falcon-agent对应的rpm包拷贝到对应服务器上
```sh
yum install ./falcon-agent*.rpm
# 修改/data/falcon-agent/cfg.json，配置正确的heartbeat和transfer服务地址，指向3.1.1安装的服务器地址
systemctl enable falcon-agent
systemctl start falcon-agent
```

### 3.2 grafana安装、配置

找一台服务器，可以跟falcon中心服务器位于同一台服务器上，安装grafana。
将grafana的rpm包拷贝到该服务器上，通过yum安装。

```sh
yum install ./grafana-*.rpm
```

安装openfalcon-datasource插件
```sh
cd /var/lib/grafana/plugins
git clone https://github.com/feiyuw/grafana-openfalcon-datasource
```

启动grafana：
```sh
systemctl enable grafana-server
systemctl start grafana-server
```

打开http://[grafana ip]:3000，使用admin/admin123登陆，添加一个openfalcon的datasource，地址指向：http://127.0.0.1:8080/api/v1/grafana

添加dashboard，即可查看监控图表安装。



**注意：**
内部或者客户处进行引擎的性能测试时，es的读写很容易成为一个性能瓶颈，但在真实运行情况中，es会分配更多的资源。

所以在性能测试的时候，遇到瓶颈可尝试将es关闭，关闭入口为进入ctu-engine服务的配置文件中：更改或添加switch.ctuservice.consolelog=false
