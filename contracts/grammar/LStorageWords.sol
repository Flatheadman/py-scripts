// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract LStorageWords {
    /***************memory***************/
    // memory 关键字用于声明复杂类型的变量，这些变量的生命周期仅限于函数执行期间。这意味着它们在函数调用结束后会被销毁，不会永久存储在区块链上。
    // 当变量是复杂类型（数组、结构体、字符串、映射）时，必须使用 memory 关键字。这是因为复杂类型的数据通常太大，无法直接存储在调用栈中。
    function processData(uint[] memory data) public pure returns (uint sum) {
        for (uint i = 0; i < data.length; i++) {
            sum += data[i];
        }
    }
}