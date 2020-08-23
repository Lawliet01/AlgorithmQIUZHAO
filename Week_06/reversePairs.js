
/**
 * @param {number[]} nums
 * @return {number}
 */
var reversePairs = function(nums) {
    return mergeSort(nums, 0, nums.length - 1);
};

function mergeSort(arr, start, end){
    if (start >= end) {
        return 0;
    }

    const mid = end + start >> 1
    
    let counter = mergeSort(arr, start, mid) + mergeSort(arr, mid + 1, end);


    for (let i = start, j = mid + 1; i <= mid; i ++ ) {
        while (j <= end && arr[i] > arr[j] * 2) j ++
        counter += j - (mid+1)
    }


    let i = start, j = mid + 1, k = 0

    let temp = []

    while(i <= mid && j <= end) {
        temp[k++] = arr[i] < arr[j] ? arr[i++]: arr[j++]
    }

    temp = temp.concat(i > mid ? arr.slice(j, end + 1) : arr.slice(i, mid +1))
    // console.log(temp);
    arr.splice(start, temp.length, ...temp)

    return counter
}

console.log(reversePairs([1, 3, 2, 3, 1]));
console.log(reversePairs([2, 4, 3, 5, 1]));
