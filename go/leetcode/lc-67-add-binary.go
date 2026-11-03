func addBinary(a string, b string) string {
    answer := ""
    rem := 0
    i, j := len(a)-1, len(b)-1

    for ; i >= 0 || j >= 0 || rem != 0; {
        x, y := 0, 0
        ch := ""
        if i >= 0 {
            x = byteToInt(a[i])
        }
        if j >= 0 {
            y = byteToInt(b[j])
        }

        sum := x + y + rem

    	ch, rem = validateSum(sum)

        answer = ch + answer

        i--
        j--
    }

    return answer
}

func byteToInt(a byte) int{
    if (a == 48) {
        return 0
    }
    return 1
}

func validateSum(sum int) (string, int) {
    if sum == 3 {
        return  "1", 1
    } else if sum == 2 {
        return "0", 1
    } else if sum == 1 {
        return "1", 0
    }
    return "0", 0
}