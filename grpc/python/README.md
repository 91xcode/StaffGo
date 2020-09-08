```

Python
1. 安装grpc包
pip install grpcio
2. 安装protobuf
pip install protobuf
3. 安装grpc的protobuf编译工具
包含了protoc编译器和生成代码的插件

pip install grpcio-tools
4. 生成代码
cat gen-python.sh
#!/usr/bin/env bash

protoDir="."
outDir="./user"

/usr/local/Cellar/python3/3.7.4_1/bin/python3.7 -m grpc_tools.protoc -I ${protoDir}/ --python_out=${outDir} --grpc_python_out=${outDir} ${protoDir}/*proto

sh gen-python.sh

5. 定义服务端
touch server.py

# coding:utf-8

from concurrent import futures
import logging
import grpc
# 支持新的包
import sys
sys.path.append("user")

import user.user_pb2 as user_pb2
import user.user_pb2_grpc as user_pb2_grpc

import json
import time



_HOST = '127.0.0.1'  # todo
_PORT = '7788'  # todo


# 实现 proto 文件中定义的 xxxServicer
class UserServer(user_pb2_grpc.UserServicer):
    # 实现 proto 文件中定义的GRPC调用
    def UserIndex(self, request, context):
        print('Received  user index request: page=%s page_size=%s' % (request.page, request.page_size))
        return user_pb2.UserIndexResponse(
            data=[
                user_pb2.UserEntity(name="sam", age=23),
                user_pb2.UserEntity(name="Tim", age=2),
            ],
            err=0,
            msg="success"
        )

    def UserView(self, request, context):
        print('Received  user index request: uid=%s ' % (request.uid))
        return user_pb2.UserViewResponse(
            data=user_pb2.UserEntity(
                    name="sam",
                    age=23
            ),
            err=0,
            msg="success"
        )


    def UserPost(self, request, context):
        print('Received  user index request: name=%s password=%s age=%s' % (request.name,request.password ,request.age))
        return user_pb2.UserPostResponse(
            err=0,
            msg="success"
        )

    def UserDelete(self, request, context):
        print('Received  user index request: uid=%s' % (request.uid))
        return user_pb2.UserDeleteResponse(
            err=0,
            msg="success"
        )


def serve():
    # 启动GRPC服务,监听特定的端口
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))  # 多线程服务器
    user_pb2_grpc.add_UserServicer_to_server(UserServer(), server)  # 注册本地服务
    server.add_insecure_port("{0}:{1}".format(_HOST, _PORT))  # 监听端口
    server.start()  # 开始接收请求进行服务
    try:
        while True:
            time.sleep(60 * 60 * 24)  # one day in seconds
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    logging.basicConfig()
    serve()


6. 定义客户端
touch client.py

# coding:utf-8

import grpc
import sys
sys.path.append("user")

import user.user_pb2 as user_pb2
import user.user_pb2_grpc as user_pb2_grpc

_HOST = '127.0.0.1'  # todo
_PORT = '7788'  # todo

from google.protobuf.json_format import MessageToDict


def main():
    with grpc.insecure_channel("{0}:{1}".format(_HOST, _PORT)) as channel:
        client = user_pb2_grpc.UserStub(channel=channel)
        response = client.UserIndex(user_pb2.UserIndexRequest(page=4,page_size=12))
        # # print(response.msg)
        result = MessageToDict(response)
        print (result["data"])

        response = client.UserView(user_pb2.UserViewRequest(uid=4))
        print(response)

        response = client.UserPost(user_pb2.UserPostRequest(name="liu",password="hao",age=22))
        print(response)

        response = client.UserDelete(user_pb2.UserDeleteRequest(uid=99))
        print(response)


if __name__ == '__main__':
    main()

启动服务/请求服务  python的服务端
/usr/local/Cellar/python3/3.7.4_1/bin/python3.7 server.py
# 新建一个窗口     python的客户端
/usr/local/Cellar/python3/3.7.4_1/bin/python3.7 client.py  


问题

AttributeError: module 'google.protobuf.descriptor' has no attribute '_internal_create_key'
➜  python git:(master) ✗ protoc --version
libprotoc 3.12.3
➜  python git:(master) ✗ pip show protobuf
Name: protobuf
Version: 3.11.2

版本要一致
pip install python3-protobuf
pip install protobuf

AttributeError: module 'google.protobuf.descriptor_pool' has no attribute 'Default'
pip install protobuf --upgrade
```