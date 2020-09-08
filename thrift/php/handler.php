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