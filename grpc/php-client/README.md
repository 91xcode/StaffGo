```
用PHP作为客户端调用Go的服务端


安装 grpc_php_plugin 插件
grpc_php_plugin插件可以帮助我们自动生成client stub客户端(封装了grpc的服务接口)，方便我们直接引入调用，否则只生成服务/请求/响应的实体类，用起来不太方便

# 下载 grpc 的库到本地
cd ~ && git clone -b $(curl -L https://grpc.io/release) https://github.com/grpc/grpc
# 更新子模块依赖
cd grpc && git submodule update --init
# 这里我们只编译 php 的插件 如果要编译所有的 make && make install
make grpc_php_plugin
# 插件路径
ll ./bins/opt/grpc_php_plugin

# /Users/liubing/grpc/bins/opt/grpc_php_plugin






PHP GRPC 扩展及依赖安装
➜  grpc git:(master) php -v
PHP 7.1.33 (cli) (built: Dec 19 2019 11:01:14) ( NTS )
Copyright (c) 1997-2018 The PHP Group
Zend Engine v3.1.0, Copyright (c) 1998-2018 Zend Technologies
    with Zend OPcache v7.1.33, Copyright (c) 1999-2018, by Zend Technologies



 #安装扩展
pecl install grpc
pecl install protobuf



mkdir -p /Users/liubing/go/src/staff_go/grpc/php-client/user

mkdir -p /Users/liubing/go/src/staff_go/grpc/php-client/grpc/bins/opt/

cd /Users/liubing/go/src/staff_go/grpc/php-client/grpc/bins/opt/

mv /Users/liubing/grpc/bins/opt/grpc_php_plugin .



生成PHP客户端
PHP只能做C端，且需要安装grpc和protobuf扩展和库


 cat gen-php.sh
#!/usr/bin/env bash

protoDir="."
outDir="./user"

protoc --proto_path=${protoDir} \
  --php_out=${outDir} \
  --grpc_out=${outDir} \
  --plugin=protoc-gen-grpc=/Users/liubing/go/src/staff_go/grpc/php-client/grpc/bins/opt/grpc_php_plugin \
  ${protoDir}/*.proto


使用 composer 管理依赖加载

cd /Users/liubing/go/src/staff_go/grpc/php-client/
 
# 使用 composer 管理项目
composer init

# 安装 grpc/protobuf 的客户端库文件
composer require grpc/grpc
composer require google/protobuf


cat composer.json
{
    "name": "liubing/php-client",
    "authors": [
        {
            "name": "liubing1",
            "email": "liubing1@conew.com"
        }
    ],
    "require": {
        "grpc/grpc": "^1.30",
        "google/protobuf": "^3.13"
    },
    "autoload": {
        "psr-4": {
            "User\\": "./user/User/",
            "GPBMetadata\\": "./user/GPBMetadata/"
        }
    }
}

# 更新 composer 加载器
composer dump-autoload

PHP客户端代码实例
在安装完php的grpc扩展和依赖库后，我们就可以编写代码了

cat client.php
<?php
/**
 * @throws Exception
 */

require_once __DIR__ . '/vendor/autoload.php';

use User\UserClient;
use User\UserEntity;
use User\UserIndexRequest;
use User\UserIndexResponse;
use User\UserViewRequest;
use User\UserViewResponse;
use User\UserPostRequest;
use User\UserPostResponse;
use User\UserDeleteRequest;
use User\UserDeleteResponse;

$userClient = new UserClient('127.0.0.1:7788', [
    'credentials' => Grpc\ChannelCredentials::createInsecure()
]);

$userIndexRequest = new UserIndexRequest();
$userIndexRequest->setPage(4);
$userIndexRequest->setPageSize(12);

//$response = $userClient->UserIndex($userIndexRequest);

/* @var $userIndexResponse UserIndexResponse */
list($userIndexResponse, $statusObj) = $userClient->UserIndex($userIndexRequest)->wait();


if (0 != $statusObj->code) {
    throw new Exception($statusObj->details, $statusObj->code);
}

printf("index request end: err %d msg %s \n", $userIndexResponse->getErr(), $userIndexResponse->getMsg());

/* @var UserEntity[] */
$data = $userIndexResponse->getData();
foreach ($data as $row) {
    echo $row->getName(), " ", $row->getAge() . PHP_EOL;
}

//
// $userClient->UserView();
// $userClient->UserPost();
// $userClient->UserDelete()




启动服务/请求服务  go的服务端
go run server.go
# 新建一个窗口     php的客户端
php -f client.php  
```