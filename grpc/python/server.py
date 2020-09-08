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




