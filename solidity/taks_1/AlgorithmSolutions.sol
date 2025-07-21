// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract AlgorithmSolutions {

    // 反转字符串
    function reverseString(string memory str) public pure returns (string memory) {
        bytes memory b = bytes(str);
        bytes memory reversed = new bytes(b.length);
        
        for (uint256 i = 0; i < b.length; i++) {
            reversed[i] = b[b.length - 1 - i];
        }
        
        return string(reversed);
    }

    // 整数转罗马数字
    function integerToRoman(uint256 num) public pure returns (string memory) {

        string memory result = "";
        while (num > 0) {
            if (num >= 1000) {
                num -= 1000;
                result = string.concat(result, "M");
            } else if (num >= 900) {
                num -= 900;
                result = string.concat(result, "CM");
            } else if (num >= 500) {
                num -= 500;
                result = string.concat(result, "D");
            } else if (num >= 400) {
                num -= 400;
                result = string.concat(result, "CD");
            } else if (num >= 100) {
                num -= 100;
                result = string.concat(result, "C");
            } else if (num >= 90) {
                num -= 90;
                result = string.concat(result, "XC");
            } else if (num >= 50) {
                num -= 50;
                result = string.concat(result, "L");
            } else if (num >= 40) {
                num -= 40;
                result = string.concat(result, "XL");
            } else if (num >= 10) {
                num -= 10;
                result = string.concat(result, "X");
            } else if (num >= 9) {
                num -= 9;
                result = string.concat(result, "IX");
            } else if (num >= 5) {
                num -= 5;
                result = string.concat(result, "V");
            } else if (num >= 4) {
                num -= 4;
                result = string.concat(result, "IV");
            } else {
                num -= 1;
                result = string.concat(result, "I");
            }
        }
        return result;
    }

    // 罗马数字转数整数
    function romanToInteger(string memory str) public pure returns (uint256) {
        bytes memory s = bytes(str);
        uint256 result = 0;
        uint256 i = 0;

        while (i < s.length) {
            uint256 value1 = _romanCharToInt(s[i]);
            uint256 value2 = 0;
            if (i + 1 < s.length) {
                value2 = _romanCharToInt(s[i + 1]);
            }

            if (value2 > value1) {
                result += (value2 - value1);
                i += 2;
            } else {
                result += value1;
                i += 1;
            }
        }
        return result;
    }

    function _romanCharToInt(bytes1 c) internal pure returns (uint256) {
        if (c == "I") return 1;
        if (c == "V") return 5;
        if (c == "X") return 10;
        if (c == "L") return 50;
        if (c == "C") return 100;
        if (c == "D") return 500;
        if (c == "M") return 1000;
        revert("Invalid Roman numeral character");
    }

    //  合并两个有序数组
    function mergeTwoSortedArrays(uint256[] memory arr1, uint256[] memory arr2) public pure returns (uint256[] memory) {
        uint256[] memory result = new uint256[](arr1.length + arr2.length);
        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;

        while(i < arr1.length && j < arr2.length) {
            if (arr1[i] < arr2[j]) {
                result[k++] = arr1[i++];
            } else {
                result[k++] = arr2[j++];
            }
        }

        while(i < arr1.length) {
            result[k++] = arr1[i++];
        }

        while(j < arr2.length) {
            result[k++] = arr2[j++];
        }
        return result;
    }

    // 二分查找
    function binarySearch(uint256[] memory arr, uint256 target) public pure returns (int256) {
        int256 left = 0;
        int256 right = int256(arr.length - 1);
        while (left <= right) {
            int256 mid = (left + right) >> 1;
            if(arr[uint256(mid)] == target) {
                return mid;
            } else if (arr[uint256(mid)] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return -1;
    }
} 