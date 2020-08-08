# sum_for_1d_array

## introduction 

This is the practice for solving problem on leetcode

[leetcode problem of running sum of 1d array](https://leetcode.com/problems/running-sum-of-1d-array/)

## problem description

Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.

## my solutions design

according to the problem description

I would like to the property

runningSum[i] =  runningSum[i-1] +nums[i]

thus, I would use a accumulater  to calculate the  element of each runningSum[i] when loop nums

## my self record

[leetcode 30days chanllenge first day](https://hackmd.io/clQdX_d1QB2NgqOnHOeDqQ?view#golang-%E9%90%B5%E4%BA%BA%E8%B3%BD-%E8%87%AA%E6%88%91%E6%8C%91%E6%88%B0%E8%B3%BD-leetcode-30-%E5%A4%A9-%E7%AC%AC%E4%B8%80%E5%A4%A9)

## reference article
[golang test](https://ithelp.ithome.com.tw/articles/10204692)