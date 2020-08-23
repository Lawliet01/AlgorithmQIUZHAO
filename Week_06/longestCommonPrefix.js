/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
  if (strs.length === 0) {
      return ""
  }

  const firstStr = strs[0];
  let curr = 0;
  outer: while (curr <= firstStr.length) {
    for (let str of strs.slice(1)) {
        if (str[curr] !== firstStr[curr] || !str[curr]) {
          break outer;
        }
    }
    curr ++
  }

  return firstStr.slice(0,curr)
};

console.log(longestCommonPrefix(["flower","flow","flight"]))
console.log(longestCommonPrefix([]));
console.log(longestCommonPrefix(["sss"]));
console.log(longestCommonPrefix(["sss", "ss"]));