/**
 * @param {string} s
 * @return {number}
 */
var firstUniqChar = function (s) {
    const map = new Map()
    for (let i = 0; i< s.length; i++) {
        map.set(s[i], map.has(s[i])? -1: i )
    }

    for (let res of map.entries()) {
        if (res[1] !== -1) {
            return res[1]
        }
    }

    return -1
};

console.log(firstUniqChar("loveleetcode"));
