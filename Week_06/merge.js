/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
var merge = function (intervals) {
    intervals.sort((a,b) => a[0] < b[0] ? -1: 1)
    
    const merge = []

    for (let interval of intervals) {
        if (merge.length === 0 || merge[merge.length - 1][1] < interval[0]) {
            merge.push(interval)
        } else {
            merge[merge.length - 1][1] = Math.max(merge[merge.length - 1][1] , interval[1])
        }
    }

    return merge
};

console.log(
  merge([
    [1, 3],
    [2, 6],
    [8, 10],
    [15, 18],
  ])
);

