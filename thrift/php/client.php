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
        $socket = new THttpClient('localhost', 7788, 'server.php');
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