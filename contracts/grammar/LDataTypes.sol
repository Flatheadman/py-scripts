// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;


contract LDataTypes {
    /***************struct结构体***************/
    // 注意：solidty的结构体是纯粹的数据存储类型，不能像golang那样定义方法
    // 创建一个结构体类型
    struct Person {
        string name;
        uint age;
        bool isActive;
    }
    // 声明一个上述结构体类型的变量
    Person public defaultPerson; 

    constructor() {
        // 推荐这种用字面量的方式初始化结构体，最节省gas费
        defaultPerson = Person("Default Person", 0, false);
        // 其他的初始化方式：
        // defaultPerson.name = "Default Person"; ... // 使用点语法
        // defaultPerson = Person({name: "Default Person", age: 0, isActive: false}); // 使用键值对
    }




}