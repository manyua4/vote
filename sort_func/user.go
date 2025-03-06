package main

type User struct {
	Name string
	Age  int
}
type Users []User

func (users Users) Len() int {
	return len(users) //Len方法返回users中的元素个数
}
func (users Users) Less(i, j int) bool {
	return users[i].Age < users[j].Age //Less方法报告索引i的元素是否比索引j的元素小

}
func (users Users) Swap(i, j int) {
	users[i], users[j] = users[j], users[i] //Swap方法交换索引i和j的两个元素

}
