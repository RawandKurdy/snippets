// Minimum Substring With All Chars - Interview Question -
// Asked by Facebook
// In Javascript
function minSubstringWithAllChars(s, t) {
    let size = s.length + 1 // length of the text
    for (let x = 0; x < size; x++) {
        for (let p = 0; p < size - x; p++) {
            // sub string candidate
            // also gets filtered
            // chars that are not in t get removed for now
            c = s.substring(p, p + x).split("").filter((el) => {
                if (t.indexOf(el) != -1) {
                    return el
                }
            })

            // if sub string contain all characters that are in t
            // then correct answer is found :D
            if (new Set(c).size == t.length) {
                return s.substring(p, p + x)
            }
        }
    }
    return ""
}

// Example
s = "adobecodebanc"
t = "abc"
console.log(minSubstringWithAllChars(s, t)) // "banc"