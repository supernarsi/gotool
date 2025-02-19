package encrypt_date

func DateStringEncrypt(date string) string {
	if len(date) != 10 {
		return ""
	}

	lastChar := date[len(date)-1]
	lastCharNum := lastChar - '0'
	offsetNum := lastCharNum%3 + 1
	offsetChar := byte(offsetNum + '0')

	encDate := make([]byte, 0, len(date)+1)
	encDate = append(encDate, offsetChar)

	for i := 0; i < len(date); i++ {
		if posIsSep(i) {
			encDate = append(encDate, lastChar)
		} else {
			encDate = append(encDate, numToChar(date[i], offsetNum))
		}
	}

	return string(encDate)
}

func DateDecrypt(encryptStr string) string {
	if len(encryptStr) != 11 {
		return ""
	}

	offsetChar := encryptStr[0]
	if !isSingleDigit(offsetChar) {
		return ""
	}

	offset := offsetChar - '0'
	decDate := make([]byte, 0, len(encryptStr)-1)

	for i := 1; i < len(encryptStr); i++ {
		if posIsSep(i - 1) {
			decDate = append(decDate, '-')
		} else {
			if !isUpperLetter(encryptStr[i]) {
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

func isUpperLetter(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

func isSingleDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func posIsSep(pos int) bool {
	return pos == 4 || pos == 7
}
