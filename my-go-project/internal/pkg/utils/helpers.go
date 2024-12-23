package utils

// StringInSlice checks if a string exists in a slice of strings.
func StringInSlice(str string, slice []string) bool {
    for _, v := range slice {
        if v == str {
            return true
        }
    }
    return false
}

// ReverseString reverses the input string and returns it.
func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// LogMessage logs a message to the console.
func LogMessage(message string) {
    fmt.Println(message)
}