#!/usr/bin/python

from __future__ import print_function

import grpc

import rpc_pb2 
import rpc_pb2_grpc


def conn():
    conn = grpc.insecure_channel('127.0.0.1:50052')
    return rpc_pb2_grpc.StoreApiServiceStub(channel=conn)

def get_config(client):
    response = client.GetConfigStoreInfo(rpc_pb2.GetConfigStoreInfoRequest(store='ks3'))
    return response

if __name__ == '__main__':
    client = conn()

    # get config
    resConf = get_config(client)
    print(resConf)

