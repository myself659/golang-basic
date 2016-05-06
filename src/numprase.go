package main

import "strconv"
import "fmt"

func main() {

	f, _ := strconv.ParseFloat("1.234", 3)
	fmt.Println(f)

	i, _ := strconv.ParseInt("20", 8, 64) //按8进制解析
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
