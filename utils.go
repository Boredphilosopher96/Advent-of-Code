package main


func Abs(num1 int, num2 int) int {
    if num1 == num2 {
        return 0
    } else if num1 > num2 {
        return num1 - num2
    } else {
        return num2 - num1
    }
}

