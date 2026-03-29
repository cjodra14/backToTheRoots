package begginer

func CanConstruct(ransomNote string, magazine string) bool {
	magazineMap := make([]int16, 26)
	for _, letter := range magazine {
		magazineMap[letter-'a']++
	}
	for _, letter := range ransomNote {
		if magazineMap[letter-'a'] == 0 {
			return false
		}

		magazineMap[letter-'a']--
	}

	return true
}
