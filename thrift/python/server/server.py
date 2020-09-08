#! /usr/bin/env python
# -*- coding: utf-8 -*-

import os
import sys
cur_path =os.path.abspath(os.path.join(os.path.dirname('__file__'), os.path.pardir))
sys.path.append(cur_path)




from user import User
from user.ttypes import UserInfo



from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer


import json

__HOST = 'localhost'
__PORT = 9090


class FormatDataHandler(object):

    def GetName(self,id):
        print ("receive GetName request: id %d\n", id)
        return "go server";

    def GetUser(self,id):
        print ("receive GetUser request: id %d\n", id)
        reslut =  UserInfo(100,"liu","xx")
        return reslut


if __name__ == '__main__':
    handler = FormatDataHandler()

    processor = User.Processor(handler)
    transport = TSocket.TServerSocket(__HOST, __PORT)
    # 传输方式，使用buffer
    tfactory = TTransport.TBufferedTransportFactory()
    # 传输的数据类型：二进制
    pfactory = TBinaryProtocol.TBinaryProtocolFactory()

    # 创建一个thrift 服务
    rpcServer = TServer.TSimpleServer(processor,transport, tfactory, pfactory)

    print('Starting the rpc server at', __HOST,':', __PORT)
    rpcServer.serve()
    print('done')
