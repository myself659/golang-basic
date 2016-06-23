package main

import (
	"fmt"
	"math"
	"strings"
)

func IsPrime(num int64) bool {
	if num < 2 {
		return false
	}

	if num == 2 {
		return true
	}
	var i int64 = 2
	for ; i <= int64(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func GetPrimeSet(min int64, max int64, set *[]int64) {
	var i int64 = min
	for ; i <= max; i++ {
		if IsPrime(i) {
			*set = append(*set, i)
		}
	}
}

func GetPrimeNums(num int64, set []int64, nums *[]int64) {
	setlen := len(set)
	for i := 0; num != 0 && i < setlen; {
		if num%set[i] == 0 {
			*nums = append(*nums, set[i])
			num = num / set[i]
		} else {
			i++
		}
	}

}

func IsCodeNum(num int64, set []int64) bool {
	for i := 0; i < len(set); {
		if num%set[i] == 0 {
			num = num / set[i]

		} else {
			i++
		}
	}

	if num != 1 {

		return false
	}

	return true
}

func Decode(table map[int]string, format string, code int64, set []int64) string {
	var ret string
	var nums []int64
	GetPrimeNums(code, set, &nums)

	fmt.Println("decode nums for ", code, ":", nums)

	offset := 0
	j := 0
	flen := len(format)
	for ; offset < flen && j < len(nums); j++ {
		index := strings.Index(format[offset:], "%s")
		if index == -1 {
			ret += format[offset:flen]
			break
		}
		if index > 0 {
			ret += format[offset : offset+index]
		}
		offset += index
		offset += len("%s")
		val, ok := table[int(nums[j])]
		if ok {
			ret += val
		}
	}

	return ret
}

func Aliyuncdn(table map[string]int, formats [4]string, codes [3]int64) []string {
	ret := make([]string, 4)

	tablelen := len(table)
	dst := make(map[int]string, tablelen)
	for key, val := range table {
		dst[val] = key
	}

	var set []int64

	GetPrimeSet(1, 256, &set)

	fmt.Println("Prime set:", set)

	temp := Decode(dst, formats[0], codes[0], set)

	ret[0] = temp

	temp = Decode(dst, formats[1], codes[1], set)

	ret[1] = temp

	temp = Decode(dst, formats[2], codes[2], set)

	ret[2] = temp
	var forthcode int64 = 31945956
	for ; forthcode < 31945975; forthcode++ {
		if IsCodeNum(forthcode, set) {
			temp = Decode(dst, formats[3], forthcode, set)

			ret[3] = temp
			break
		}

	}

	return ret
}

func main() {

	table := map[string]int{"a": 3, "b": 43, "c": 7, "d": 31, "e": 47, "f": 53, "g": 59,
		"h": 61, "i": 19, "j": 11, "k": 67, "l": 5, "m": 41, "n": 37, "o": 17,
		"p": 71, "q": 73, "r": 79, "s": 83, "t": 13, "u": 29, "v": 89, "w": 2, "x": 97, "y": 23, "z": 101}

	var formats [4]string = [4]string{"%se%s%s%s%se", "%s%s", "%s%s%s%s", "%s%si%s%snc%sn"}
	var codes [3]int64 = [3]int64{5025370, 22763, 13540483}

	fmt.Println(Aliyuncdn(table, formats, codes))
}
