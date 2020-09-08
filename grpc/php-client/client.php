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
// $userClient->UserDelete();