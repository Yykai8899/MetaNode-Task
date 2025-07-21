// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting{

    mapping(address => uint256) public votes;
    address[] public candidates;

    constructor() {}

    function vote(address candidate) public {
        if (votes[candidate] == 0) {
            candidates.push(candidate);
        }
        votes[candidate]++;
    }

    function getVotes(address candidate) public view returns (uint256) {
        return votes[candidate];
    }

    function resetVotes() public {
        for (uint i = 0; i < candidates.length; i++) {
            votes[candidates[i]] = 0;
        }
    }
}