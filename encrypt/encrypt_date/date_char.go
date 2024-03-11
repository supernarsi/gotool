package encrypt_date

func DateStringEncrypt(date string) string {
	if len(date) != 10 {
		return ""
	}
	var sigChar byte
	lastCharAsc := date[len(date)-1]
	lastCharNum := lastCharAsc - '0'
	offsetNum := lastCharNum%3 + 1
	var offsetAsc = offsetNum + '0'
	encDate := make([]byte, 1, len(date))
	encDate[0] = offsetAsc
	for i := 0; i < len(date); i++ {
		if !posIsSep(i) {
			sigChar = numToChar(date[i], offsetNum)
		} else {
			sigChar = lastCharAsc
		}
		encDate = append(encDate, sigChar)
	}
	return string(encDate)
}

func DateDecrypt(encryptStr string) string {
	if len(encryptStr) != 11 {
		return ""
	}
	offset := encryptStr[0]
	if !offsetCharIsSingleNum(offset) {
		return ""
	}
	offset -= '0'
	var decDate []byte
	for i := 1; i < len(encryptStr); i++ {
		if posIsSep(i - 1) {
			decDate = append(decDate, '-')
		} else {
			if encryptStr[i] > 'Z' || !charIsUpperLetter(encryptStr[i]) {
				return ""
			}
			decDate = append(decDate, charToNum(encryptStr[i], offset))
		}
	}
	return string(decDate)
}

func numToChar(numChar byte, offset byte) byte {
	return numChar - '0' + 'A' + offset
}

func charToNum(letterChar byte, offset byte) byte {
	return letterChar - 'A' + '0' - offset
}

func charIsUpperLetter(letter byte) bool {
	return letter >= 'A' && letter <= 'Z'
}

func posIsSep(pos int) bool {
	return pos == 4 || pos == 7
}

func offsetCharIsSingleNum(offset byte) bool {
	return offset >= '0' && offset <= '9'
}
