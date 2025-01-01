# 定义数组变量
arrays = {
    "solidity": [
        {"key": "语法-函数", "value": ["contracts/grammar/LFunction.sol"]},
        {"key": "语法-数据类型", "value": ["contracts/grammar/LDataTypes.sol"]},
        {"key": "语法-存储位置关键字", "value": ["contracts/grammar/LStorageWords.sol"]},
    ],
    "golang": [
        {"key": "应用-grpc服务的创建", "value": [
            "proto文件的创建与编译: golang/proto/hello.proto", 
            "实现并运行服务端: golang/server/server.go",
            "实现并运行客户端: golang/client/client.go",
        ]},
        {"key": "应用-rpc连接池", "value": [
            "库文件: golang/multiclient/multiclient.go", 
            "测试文件: golang/multiclient/multiclient_test.go",
        ]},
    ],
}