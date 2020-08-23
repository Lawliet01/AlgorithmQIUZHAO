/**
 * @param {character[]} s
 * @return {void} Do not return anything, modify s in-place instead.
 */
var reverseString = function (s) {
    if (s.length > 1) {
        let l = 0, r= s.length - 1
        while (l < r) {
            const store = s[l]
            s[l++] = s[r]
            s[r--] = store
        }
    }

    return s
};

console.log(reverseString(["h", "e", "l", "l", "o"]));
console.log(reverseString(["H", "a", "n", "n", "a", "h"]));

