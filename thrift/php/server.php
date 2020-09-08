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