/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function (s) {
  s = s.trim()
  for (let i = s.length - 1; i >= 0; i--) {
    if (/\s/.test(s[i])) {
      return s.length - i - 1;
    }
  }
  return s.length;
};

console.log(lengthOfLastWord("Hello World"));
console.log(lengthOfLastWord(" "));
console.log(lengthOfLastWord("a "));
