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