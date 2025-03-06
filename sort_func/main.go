package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{9, 6, 8, 3, 4, 5, 1, 2, 7}
	//升序
	sort.Ints(nums)
	fmt.Println(nums)
	//降序
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) //这里是用了reverse函数将less方法进行逆序
	fmt.Println(nums)
	//sort.string()对字符串进行排序
	//sort.Float64s()对浮点数进行排序
	users := []User{
		{Name: "a", Age: 10},
		{Name: "b", Age: 30},
		{Name: "c", Age: 20},
	}
	sort.Slice(users, func(i, j int) bool { //采用匿名函数对结构体字段age进行排序
		return users[i].Age < users[j].Age
	})
	fmt.Println(users)

	userbyage := Users(users) //users实现了那三种方法 直接调用sort排序
	sort.Sort(userbyage)
	fmt.Println(userbyage)

}
