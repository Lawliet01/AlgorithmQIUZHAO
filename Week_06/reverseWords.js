/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function (s) {
    const match = /\S+/g
    const store = []
    let curr = 0, end = 0
    s = s + " "
    for (let i of s) {
        const isSpace = /\s/.test(i)
        if (!isSpace) {
            end++
            continue
        }

        if (end > curr) {
            store.push(s.slice(curr, end))
        } 
        curr = end = end + 1
    }

    for (let i = 0, j = store.length - 1; i < j; i++, j--) {
        const temp = store[i]
        store[i] = store[j]
        store[j] = temp
    }

    return store.join(" ")
};

console.log(reverseWords("  hello world!  "));
console.log(reverseWords("the sky is blue"));

