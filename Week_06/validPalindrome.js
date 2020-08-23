/**
 * @param {string} s
 * @return {boolean}
 */
var validPalindrome = function (s) {
    return test(s, false)
};

function test(s, isDeleted) {
    let i = 0, j = s.length - 1

    while (i < j) {
        if (s[i] === s[j]) {
            i++;j--;continue
        }

        if (!isDeleted) {
            return test(s.slice(i, j), true) || test(s.slice(i + 1, j + 1), true)
        } 
        
        return false
        
    }

    return true
}

console.log(validPalindrome("aba"));
console.log(validPalindrome("abccda"));


