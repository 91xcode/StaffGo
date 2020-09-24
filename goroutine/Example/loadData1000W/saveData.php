<?php

// 制定允许其他域名访问
header("Access-Control-Allow-Origin:*");
// 响应类型
header('Access-Control-Allow-Methods:POST');
// 响应头设置
header('Access-Control-Allow-Headers:x-requested-with, content-type');

// 连接数据库

header('content-type:text/html;charset=utf-8');

define('DB_HOST','127.0.0.1');

define('DB_USER','root');

define('DB_PASS','');

define('DB_NAME','load_data_test');

define('DB_PORT',3306);

define('DB_CHAR','utf8');

define('APPNAME','');

 

$sqli = new mysqli( DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT);

$sqli->query( "SET NAMES ".DB_CHAR );

 

//batch_data_to_mysql($sqli);

batch_data_to_file();


function batch_data_to_mysql($sqli){
    /**

     * 批量添加 方法3

     * 使用优化SQL语句，将SQL语句拼接使用 insert into table() values(),(),()然后一次性添加；

     */
    ini_set("max_execution_time", "200000");
    
    $arr = $_POST['dat'];
    

    if(empty($arr)){
        exit("数据为空!");
    }
    //将字符串转化成数组  
    $json=json_decode($arr);


    $time_s = date("H:i:s",time())."<br/>";

    $sql = "INSERT INTO test(`name`,`card_id`,`phone`,`bank_id`,`email`,`address`,`create_at`) VALUES ";

    foreach($json as $v){
        $sql .= "( '".$v->name."','".$v->card_id."','".$v->phone."','".$v->bank_id."','".$v->email."','".$v->address."','".$v->create_at."'),";
    }

    $sql = substr( $sql,0, strlen($sql)-1 );

//     echo $sql;die();

    $sqli->query( $sql );

    $time_e = date("H:i:s",time())."<br/>";

    response(1,"",array('time'=>$time_s."--".$time_e));  
}




function batch_data_to_file(){
    $arr = $_POST['dat'];
    
    if(empty($arr)){
        exit("数据为空!");
    }
    
        //将字符串转化成数组  
    $json=json_decode($arr);
    
    $time_s = date("H:i:s",time())."<br/>";
    
    $txt = '';
    
    foreach($json as $v){
        $txt .= $v->name.','.$v->card_id.','.$v->phone.','.$v->bank_id.','.$v->email.','.$v->address.','.$v->create_at.PHP_EOL;
    }
    
    file_put_contents("/tmp/test.txt", $txt, FILE_APPEND);
    
    $time_e = date("H:i:s",time())."<br/>";
    
    response(1,"",array('time'=>$time_s."--".$time_e));  
}

function response($code = 1, $msg = "", $data = array()) {
        $result = array(
            'code' => $code,
            'msg' => $msg,
        );
        if (!empty($data)) {
            $result['data'] = $data;
        }
        if (!headers_sent()) {
            header('Content-Type: application/json; charset=UTF-8');
        }
        echo json_encode($result);
        exit;
    }
