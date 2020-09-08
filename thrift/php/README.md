```

thrift php

0.需要的步骤:

	配置thrift的模块
	生成服务端和客户端的代码
	下载依赖
	定义服务端接口和注册
	客户端调用远程方法

cd /Users/liubing/go/src/staff_go/thrift/php

1.配置thrift的模块

cat user.thrift
# 命名空间，可以不写也可以写多个，按照使用服务端、客户端语言来写即可
namespace go thrift.user
namespace php thrift.user

# 定义一个用户结构
struct UserInfo {
    #序号:字段类型 字段名
    1:i64 id
    2:string username
    3:string password
    4:string email
}

# 定义一个用户服务
service User{
    # 定义一个GetUser方法（接收一个用户id，返回上面定义的用户信息）
    UserInfo GetUser(1:i64 id)
    # 定义一个GetName方法（接收一个用户id，返回用户名称）
    string GetName(1:i64 id)

   # 方法定义格式：
   # 返回的类型 方法名(序号:参数类型 参数名 ... )
   # bool Test(1:i32 id, 2:string name, 3:i32 age ... )
}

2.生成服务端和客户端的代码

thrift -out . --gen php:server user.thrift  

tree thrift

thrift
└── user
    ├── UserClient.php
    ├── UserIf.php
    ├── UserInfo.php
    ├── UserProcessor.php
    ├── User_GetName_args.php
    ├── User_GetName_result.php
    ├── User_GetUser_args.php
    └── User_GetUser_result.php

3.下载依赖

地址: https://github.com/apache/thrift  (clone 下来复制 lib/php/lib 到目录thrift, 不管在服务端还是客户端都需要用到.)

cd ~

git clone https://github.com/apache/thrift

cd /Users/liubing/go/src/staff_go/thrift/php && mkdir lib-php

cp -R ~/thrift/lib/php/* ./lib-php


然后需要修改/lib-php/里的lib目录名为Thrift，否则后续会一直提示Class 'Thrift\Transport\TSocket' not found

4.定义服务端接口和注册
cat server.php 
<?php

error_reporting(E_ALL);

$ROOT_DIR = realpath(dirname(__FILE__) . '/lib-php/');
$GEN_DIR = realpath(dirname(__FILE__)) . '/thrift/';
require_once $ROOT_DIR . '/Thrift/ClassLoader/ThriftClassLoader.php';

use Thrift\ClassLoader\ThriftClassLoader;

$loader = new ThriftClassLoader();
$loader->registerNamespace('Thrift',$ROOT_DIR);
$loader->registerNamespace('user', $GEN_DIR);
$loader->register();


use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TPhpStream;
use Thrift\Transport\TBufferedTransport;



//use Thrift\Exception\TException;
//use Thrift\Factory\TTransportFactory;
//use Thrift\Factory\TBinaryProtocolFactory;
//use Thrift\Server\TServerSocket;
//use Thrift\Server\TSimpleServer;

try {

    require_once $GEN_DIR. 'user/UserIf.php';
    require_once $GEN_DIR. 'user/UserInfo.php';
    require_once $GEN_DIR. 'user/User_GetName_args.php';
    require_once $GEN_DIR. 'user/User_GetName_result.php';
    require_once $GEN_DIR. 'user/User_GetUser_args.php';
    require_once $GEN_DIR. 'user/User_GetUser_result.php';

    //实现方法, 其实实际的内容就在这里处理.
    require_once realpath(dirname(__FILE__)).'/handler.php';
    require_once $GEN_DIR. 'user/UserProcessor.php';


    header('Content-Type', 'application/x-thrift');
    if (php_sapi_name() == 'cli') {
        echo "\r\n";
    }


//    $handler = new \Handler();
//    $processor = new \thrift\user\UserProcessor($handler);
//    $transportFactory = new TTransportFactory();
//    $protocolFactory = new TBinaryProtocolFactory(true, true);
//    //作为cli方式运行，监听端口，官方实现
//    $transport = new TServerSocket('127.0.0.1', 9090);
//    $transport->setAcceptTimeout(3000);
//    $server = new TSimpleServer($processor, $transport, $transportFactory, $transportFactory, $protocolFactory, $protocolFactory);
//    $server->serve();




    $handler = new Handler();
    $processor = new \thrift\user\UserProcessor($handler);

    $transport = new TBufferedTransport(new TPhpStream(TPhpStream::MODE_R | TPhpStream::MODE_W));
    $protocol = new TBinaryProtocol($transport, true, true);

    $transport->open();
    $processor->process($protocol, $protocol);
    $transport->close();




} catch (TException $tx) {
    print 'TException: '.$tx->getMessage()."\n";
}


上面定义的 Handler 需要自己定义这个类, 要实现GetName 和GetUser方法, 其实实际的内容就在这里处理.
cat handler.php 
<?php



class Handler implements thrift\user\UserIf{


    public function GetName($id){
        return "go server";
    }

    public function GetUser($id){

        $response = new thrift\user\UserInfo();
        $response->id = 111;
        $response->username = "liu";
        return $response;
    }
}



5.定义客户端

 cat client.php
<?php

error_reporting(E_ALL);

$ROOT_DIR = realpath(dirname(__FILE__) . '/lib-php/');
$GEN_DIR = realpath(dirname(__FILE__)) . '/thrift/';
require_once $ROOT_DIR . '/Thrift/ClassLoader/ThriftClassLoader.php';

use Thrift\ClassLoader\ThriftClassLoader;

$loader = new ThriftClassLoader();
$loader->registerNamespace('Thrift',$ROOT_DIR);
$loader->registerNamespace('user', $GEN_DIR);
$loader->register();

use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\THttpClient;
use Thrift\Transport\TBufferedTransport;
use Thrift\Exception\TException;

try {
    require_once $GEN_DIR. 'user/UserIf.php';
    require_once $GEN_DIR. 'user/UserInfo.php';
    require_once $GEN_DIR. 'user/User_GetName_args.php';
    require_once $GEN_DIR. 'user/User_GetName_result.php';
    require_once $GEN_DIR. 'user/User_GetUser_args.php';
    require_once $GEN_DIR. 'user/User_GetUser_result.php';
    require_once $GEN_DIR. 'user/UserClient.php';



    if (array_search('--http', $argv)) {
        $socket = new THttpClient('localhost', 8080, 'server.php');
    } else {
        $socket = new TSocket('localhost', 9090);
    }
    $transport = new TBufferedTransport($socket, 1024, 1024);
    $protocol = new TBinaryProtocol($transport);
    $client = new \thrift\user\UserClient($protocol);


    $transport->open();


    $s = $client->GetName(100);
    var_dump($s);

    $a = $client->GetUser(100);
    var_dump($a);
    echo PHP_EOL;
    var_dump($a->id);

    $transport->close();

} catch (TException $tx) {
  print 'TException: '.$tx->getMessage()."\n";
}

6.客户端调用远程方法

thrift实现的服务端不能自己起server服务独立运行，还需要借助php-fpm运行 这里我们直接使用php -S 0.0.0.0:8080启动httpserver，就不使用php-fpm演示了

php -S 0.0.0.0:8080

PHP 7.1.33 Development Server started at Fri Aug 21 18:17:26 2020
Listening on http://0.0.0.0:8080
Document root is /Users/liubing/go/src/staff_go/thrift/php
Press Ctrl-C to quit.

我们使用php客户端，注意需要加参数，调用http协议连接：
php client.php --http 


----------------------compose版本
可以更改为compose
cd /Users/liubing/go/src/staff_go/thrift/php

compose init  一直回车

cat composer.json
{
    "name": "liubing/php",
    "authors": [
        {
            "name": "liubing1",
            "email": "liubing1@conew.com"
        }
    ],
    "require": {},
    "autoload": {
        "classmap": ["thrift/","handler.php"]
    }
}

执行composer install更新自动加载，更新执行composer dump-autoload


cat server_compose.php 
cat server_compose.php 
<?php

error_reporting(E_ALL);
require 'vendor/autoload.php';
$ROOT_DIR = realpath(dirname(__FILE__) . '/lib-php/');
$GEN_DIR = realpath(dirname(__FILE__)) . '/thrift/';
require_once $ROOT_DIR . '/Thrift/ClassLoader/ThriftClassLoader.php';

use Thrift\ClassLoader\ThriftClassLoader;

$loader = new ThriftClassLoader();
$loader->registerNamespace('Thrift',$ROOT_DIR);
$loader->registerNamespace('user', $GEN_DIR);
$loader->register();


use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TPhpStream;
use Thrift\Transport\TBufferedTransport;



//use Thrift\Exception\TException;
//use Thrift\Factory\TTransportFactory;
//use Thrift\Factory\TBinaryProtocolFactory;
//use Thrift\Server\TServerSocket;
//use Thrift\Server\TSimpleServer;

try {



    header('Content-Type', 'application/x-thrift');
    if (php_sapi_name() == 'cli') {
        echo "\r\n";
    }


//    $handler = new \Handler();
//    $processor = new \thrift\user\UserProcessor($handler);
//    $transportFactory = new TTransportFactory();
//    $protocolFactory = new TBinaryProtocolFactory(true, true);

//    //作为cli方式运行，监听端口，官方实现
//    $transport = new TServerSocket('127.0.0.1', 9090);
//    $transport->setAcceptTimeout(3000);
//    $server = new TSimpleServer($processor, $transport, $transportFactory, $transportFactory, $protocolFactory, $protocolFactory);
//    $server->serve();




    $handler = new Handler();
    $processor = new \thrift\user\UserProcessor($handler);

    $transport = new TBufferedTransport(new TPhpStream(TPhpStream::MODE_R | TPhpStream::MODE_W));
    $protocol = new TBinaryProtocol($transport, true, true);

    $transport->open();
    $processor->process($protocol, $protocol);
    $transport->close();




} catch (TException $tx) {
    print 'TException: '.$tx->getMessage()."\n";
}

cat client_compose.php 
<?php

error_reporting(E_ALL);
require 'vendor/autoload.php';

$ROOT_DIR = realpath(dirname(__FILE__) . '/lib-php/');
$GEN_DIR = realpath(dirname(__FILE__)) . '/thrift/';
require_once $ROOT_DIR . '/Thrift/ClassLoader/ThriftClassLoader.php';

use Thrift\ClassLoader\ThriftClassLoader;

$loader = new ThriftClassLoader();
$loader->registerNamespace('Thrift',$ROOT_DIR);
$loader->registerNamespace('user', $GEN_DIR);
$loader->register();

use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\THttpClient;
use Thrift\Transport\TBufferedTransport;
use Thrift\Exception\TException;

try {

    if (array_search('--http', $argv)) {
        $socket = new THttpClient('localhost', 8080, 'server_compose.php');
    } else {
        $socket = new TSocket('localhost', 9090);
    }
    $transport = new TBufferedTransport($socket, 1024, 1024);
    $protocol = new TBinaryProtocol($transport);
    $client = new \thrift\user\UserClient($protocol);


    $transport->open();


    $s = $client->GetName(100);
    var_dump($s);

    $a = $client->GetUser(100);
    var_dump($a);
    echo PHP_EOL;
    var_dump($a->id);

    $transport->close();

} catch (TException $tx) {
  print 'TException: '.$tx->getMessage()."\n";
}
6.客户端调用远程方法

thrift实现的服务端不能自己起server服务独立运行，还需要借助php-fpm运行 这里我们直接使用php -S 0.0.0.0:8080启动httpserver，就不使用php-fpm演示了

php -S 0.0.0.0:8080

PHP 7.1.33 Development Server started at Fri Aug 21 18:17:26 2020
Listening on http://0.0.0.0:8080
Document root is /Users/liubing/go/src/staff_go/thrift/php
Press Ctrl-C to quit.

我们使用php客户端，注意需要加参数，调用http协议连接：
php client_compose.php --http 


```