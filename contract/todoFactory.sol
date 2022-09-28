// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.9.0;

import "./todo.sol";

contract TodoFactory is Todo {

Todo[] public todoArray;

function createTodoContract() public {
Todo todo = new Todo();
todoArray.push(todo);
}

function tFAdd(uint256 _todoIndex, string memory content) public {
return Todo(address(todoArray[_todoIndex])).add(content);
}

function tFGet(uint256 _todoIndex, uint256 _whichTodo) public view returns (string memory){
return Todo(address(todoArray[_todoIndex])).get(_whichTodo).content;
}


}