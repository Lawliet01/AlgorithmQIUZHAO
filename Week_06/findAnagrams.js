/**
 * @param {string} s
 * @param {string} p
 * @return {number[]}
 */
var findAnagrams = function (s, p) {

    const sLen = s.length;
    const pLen = p.length;

    if (sLen < pLen) {
        return []
    }

    const map = new Map()
    for (i = 0; i < pLen; i++) {
        if (i < pLen - 1) {
            map.set(s[i], map.has(s[i]) ? map.get(s[i]) + 1 : 1)
        }
        map.set(p[i], map.has(p[i]) ? map.get(p[i]) - 1: -1);
    }

    const res = []
    outer : for (i = 0; i <= sLen - pLen ; i++) {
        if (i !== 0) {
            map.set(s[i-1], map.get(s[i-1]) - 1)
        }
        map.set(s[i + pLen - 1], map.has(s[i+pLen - 1]) ? map.get(s[i+pLen - 1]) + 1 : 1)

        for (let v of map.values()) {
            if (v !== 0) {
                continue outer
            }
        }

        res.push(i)
    }
    return res
};


console.log(findAnagrams("cbaebabacd", "abc"));
console.log(findAnagrams("abab", "ab"));
