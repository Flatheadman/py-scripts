// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract LFunction {

    /***************函数的返回值***************/
    function returnsInt() public pure returns (int) { // 返回值声明使用returns关键字，并使用括号包裹返回值的类型，只有一个返回值也要括号
        // 一个返回值可以直接返回，多个返回值需要使用元组的方式返回（见下面的函数）
        return (42); // 也可以用 return 42;
    }
    function returnsIntAndString() public pure returns (int, string memory) { 
        return (42, "Hello from Solidity!");
    }
    function returnsUse() public pure returns (int e, string memory c) { // 如果在函数名中声明了返回变量，那么在函数体中就不需要再声明，且会自动返回这些变量而无需显式return
        int a = returnsInt();
        int b;
        (b, c) = returnsIntAndString();
        // 错误写法：
        // (int b, c) = returnsIntAndString(); // 这种写法混合了类型声明和类型推断。 你声明了b的类型为int，但尝试让编译器推断c的类型。 Solidity不允许这种混合方式。 要么全部显式声明类型，要么全部省略类型（前提是变量已在之前声明）。

        e = a+b;
    }


}