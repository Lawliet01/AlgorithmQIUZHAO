/**
 * @param {string} str
 * @return {number}
 */

/**
* 换逻辑写，不用hashmap
* 用charCode来判断数字
* 也不用正则
* 练基本功
*/
var myAtoi = function (str) {
  const regExp = /^\s*(-|\+)?([0-9]*)/;
  const match = regExp.exec(str);
  let res = 0;
  if (match[2]) {
    const num = new Map([
      ["0", 0],
      ["1", 1],
      ["2", 2],
      ["3", 3],
      ["4", 4],
      ["5", 5],
      ["6", 6],
      ["7", 7],
      ["8", 8],
      ["9", 9],
    ]);

	for (let i of match[2]) {
		res = res * 10 + num.get(i)
	}

	if (match[1] === "-") {
		res = -res
	}

	return Math.max(Math.min(res, 2**31 - 1), -(2**31))

  }
  return res
};

console.log(myAtoi("42"));
console.log(myAtoi("   -42"));
console.log(myAtoi("4193 with words"));
console.log(myAtoi("words and 987"));
console.log(myAtoi("-91283472332"));
console.log(myAtoi("+1"));
