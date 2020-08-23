/**
 * @param {string} J
 * @param {string} S
 * @return {number}
 */
var numJewelsInStones = function (J, S) {
    let res = 0;
    const hashTable = new Set(J)
    for (let s of S) {
        if (hashTable.has(s)) {
            res ++
        }
    }
    return res
};

console.log(numJewelsInStones("aA", "aAAbbbb"));
