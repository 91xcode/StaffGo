<?php


$db = new PDO('mysql:host=localhost;dbname=demo', 'root', '');

try {

    foreach ($db->query('select * from user') as $row){

    print_r($row);

    }

    $db = null; //关闭数据库

} catch (PDOException $e) {

    echo $e->getMessage();

}