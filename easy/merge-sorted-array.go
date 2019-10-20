func merge(nums1 []int, m int, nums2 []int, n int) []int {
    var results []int

    var i, j int
    for i < m && j < n{
      var min int
      if nums1[i] > nums2[j]{
        min = nums2[j]
         j++
      }else{
        min = nums1[i]
          i++
      }
      results = append(results, min)
    }

    if i < m{
        // append rest of m
        for k := i; k < m; k++{
            results = append(results, nums1[k])
        }
    }
    if j < n{
        // append rest of n
         for k := j; k < n; k++{
            results = append(results, nums2[k])
         }
    }
    return results
}


// Notice we need to modify nums1, not return a new list.. oops
func merge(nums1 []int, m int, nums2 []int, n int)  {
    count := m + n - 1 // result index
    m-- // ponter at end of nums1
    n-- // pointer at end of nums2
    for m >= 0 && n >= 0{ // in bounds in both arrays
        if nums1[m] > nums2[n]{
            nums1[count] = nums1[m]
            count--
            m--
        }else{
            nums1[count] = nums2[n]
            count--
            n--
        }
    }

    for n >= 0{ // N must have the most due to it being the bigger array starting with zeros
        nums1[count] = nums2[n]
        count--
        n--
    }
}
