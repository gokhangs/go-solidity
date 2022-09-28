//SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract Todo{
    Task[] tasks; //mapping
    address public owner;

    struct Task{
        string content;
        bool status;
    }

    modifier isOwner(){
        require(owner==msg.sender);
        _;
    }

    constructor(){
        owner = msg.sender;
    }

    //memory means we pass "content" as copy
    function add(string memory _content) public isOwner{
        tasks.push(Task(_content, false));
    }

    function get(uint _id) public view returns (Task memory){
        return tasks[_id];
    }

    function list() public view returns (Task[] memory) {
        return tasks;
    }

    function toggle(uint _id) public isOwner {
        tasks[_id].status = !tasks[_id].status;
    }

    function update(uint _id, string memory _content, bool status) public isOwner{ //} public returns (Task memory){
        tasks[_id].status = status;
        tasks[_id].content = _content;
        //return tasks[_id];
    }

    function remove(uint _id) public isOwner{

        //delete keyword sets value to default value
        // delete tasks[_id];
        for(uint i = _id; i < tasks.length -1; i++){
            tasks[_id] = tasks[_id +1];
        }
        tasks.pop();
    }


}