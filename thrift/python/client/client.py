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

