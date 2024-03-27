# [713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)

## Description
Given an array of integers `nums` and an integer `k`, return the number of contiguous subarrays where the product of all the elements in the subarray is strictly less than `k`.

**Example 1:**
```
Input: nums = [10,5,2,6], k = 100
Output: 8
Explanation: The 8 subarrays that have product less than 100 are:
[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.
```

**Example 2:**
```
Input: nums = [1,2,3], k = 0
Output: 0
```

**Constraints:**
- `1 <= nums.length <= 3 * 10^4`
- `1 <= nums[i] <= 1000`
- `0 <= k <= 10^6`

**solution approach**
每新增一個數字若乘積未超過 `k`，則會新增現在 subarray 的 length 組合，為包含自己以及逐一往前的 subarray 數量，ex:

`nums = [1, 2, 3, 4]`, `k = 100`
在所有 subarray 組合乘積皆不超過 k 的情況下：
- 在計算 1 時僅有 `[1]` 符合
- 在計算 2 時有 `[2]`, `[1, 2]` 符合
- 在計算 3 時有 `[3]`, `[2, 3]`, `[1, 2, 3]` 符合
- 在計算 4 時有 `[4]`, `[3, 4]`, `[2, 3, 4]`, `[1, 2, 3, 4]` 符合

依此類推
