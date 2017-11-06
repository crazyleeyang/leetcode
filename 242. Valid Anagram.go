func isAnagram(s string, t string) bool {
    if len(s) != len(t) { return false}

    var ch26 [26]int
    for i := range s {
        ch26[s[i]-'a']++
        ch26[t[i]-'a']--
    }

    for _, v := range ch26{
        if v != 0 {return false}
    }
    return true 
}
