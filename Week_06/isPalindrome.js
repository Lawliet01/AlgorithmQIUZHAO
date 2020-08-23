/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function (s) {
    s = s.toLowerCase()
    let i = 0, j = s.length - 1
    while (i < j) {
        while (! /[\da-z]/.test(s[i]) && i < j) {
            i++
        }

        while (!/[\da-z]/.test(s[j]) && i < j) {
            j--
        }

        if (s[i++] !== s[j--]){
            return false
        }   
    }

    return true
};

console.log(isPalindrome("A man, a plan, a canal: Panama"));
console.log(isPalindrome("race a car"));
console.log(isPalindrome("ab_a"));


