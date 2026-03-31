func getSum(a int, b int) int {
    carry := 0
    res := 0

    for i := 0; i < 32; i++ {
        aBit := (a >> i) & 1
        bBit := (b >> i) & 1
        curBit := aBit ^ bBit ^ carry
        if (aBit + bBit + carry) >= 2 {
            carry = 1
        } else {
            carry = 0
        }
        if curBit == 1 {
            res |= (1 << i)
        }
    }

    return int(int32(res))
}