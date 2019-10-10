func mySqrt(x int) int {
    if x == 0 {
        return x
    }
    for i := 1; i <= x; i++{
        if i * i > x{
            return i-1 // Round down
        }else if i * i == x{
            return i // Perfect square
        }
    }
    return -1
}

// Can also do binary search

func mySqrt(x int) int {
    left := 1
    right := x

    for left <= right{
        mid := (left + right) / 2
        if mid * mid == x { // If its the square, we found our valu
            return mid
        }
        if mid * mid < x{ // If the square of current try is less than x, need to search the right
            left = mid + 1
        }
        if mid * mid > x { // If the square of curret try is greater than x, need to search the left
            right = mid - 1
        }
    }
    return left - 1
}
