```

thrift python

0.需要的步骤:

	配置thrift的模块
	生成服务端和客户端的代码
	下载依赖
	定义服务端接口和注册
	客户端调用远程方法

cd /Users/liubing/go/src/staff_go/thrift/python

1.配置thrift的模块

cat user.thrift
# 命名空间，可以不写也可以写多个，按照使用服务端、客户端语言来写即可
namespace go thrift.user
namespace php thrift.user

# 定义一个用户结构
struct UserInfo {
    #序号:字段类型 字段名
    1:i64 id
    2:string username
    3:string password
    4:string email
}

# 定义一个用户服务
service User{
    # 定义一个GetUser方法（接收一个用户id，返回上面定义的用户信息）
    UserInfo GetUser(1:i64 id)
    # 定义一个GetName方法（接收一个用户id，返回用户名称）
    string GetName(1:i64 id)

   # 方法定义格式：
   # 返回的类型 方法名(序号:参数类型 参数名 ... )
   # bool Test(1:i32 id, 2:string name, 3:i32 age ... )
}

2.生成服务端和客户端的代码

thrift -out . --gen py user.thrift  

tree user 
user
├── User-remote
├── User.py
├── __init__.py
├── constants.py
└── ttypes.py

3.下载依赖

pip install thrift


4.定义服务端接口和注册
cat server/server.py
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



5.定义客户端

 cat client/client.py 
#! /usr/bin/env python
# -*- coding: utf-8 -*-

import os
import sys
sys.path.append(os.path.abspath(os.path.join(os.path.dirname('__file__'), os.path.pardir)))

from user.User import Client
from user.ttypes import *
from user.constants import *

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol


__HOST = 'localhost'
__PORT = 9090

try:
    print (os.path.abspath(os.path.join(os.path.dirname('__file__'), os.path.pardir)))
    # Make socket
    transport = TSocket.TSocket(__HOST, __PORT)

    # Buffering is critical. Raw sockets are very slow
    transport = TTransport.TBufferedTransport(transport)

    # Wrap in a protocol
    protocol = TBinaryProtocol.TBinaryProtocol(transport)

    # Create a client to use the protocol encoder
    client = Client(protocol)

    # Connect!
    transport.open()



    msg = client.GetName(100)
    print (msg)
    msg = client.GetUser(100)
    print (msg)
    transport.close()

except Thrift.TException as tx:
    print ("%s" % (tx.message))



6.客户端调用远程方法

打开2个终端 分别
 python server/server.py
 
 python client/client.py
 


```