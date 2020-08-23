/**
 * @param {string} s
 * @param {number} k
 * @return {string}
 */
var reverseStr = function (s, k) {
    s = s.split("")
    const len = s.length
    let curr = 0;

    while (curr < len) {
        const end = len > curr + k ? curr + k - 1: len - 1
        let i = curr , j = end
        while (i < j) {
            const store = s[i]
            s[i++] = s[j]
            s[j--] = store
        }
        curr += 2 * k
    }

    return s.join("")
};

console.log(reverseStr("abcdefg", 2));
console.log(reverseStr("", 2));
console.log(reverseStr("abc", 2));

