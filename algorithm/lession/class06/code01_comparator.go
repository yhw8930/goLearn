package main

import (
	"fmt"
	"sort"
)

// Student 学生结构体
type Student struct {
	Name string
	ID   int
	Age  int
}

// ------------------------------------------------------------
// 比较器1：ID 升序，ID 相同则 Age 降序
// 返回负数：a 排前面
// 返回正数：b 排前面
// 返回0：顺序无关
// ------------------------------------------------------------
func IdShengAgeJiangOrder(a, b *Student) int {
	if a.ID != b.ID {
		return a.ID - b.ID
	}
	return b.Age - a.Age
}

// ------------------------------------------------------------
// 比较器2：ID 升序
// ------------------------------------------------------------
func IdAscendingComparator(a, b *Student) int {
	return a.ID - b.ID
}

// ------------------------------------------------------------
// 比较器3：ID 降序
// ------------------------------------------------------------
func IdDescendingComparator(a, b *Student) int {
	return b.ID - a.ID
}

func PrintStudents(students []*Student) {
	for _, s := range students {
		fmt.Printf("Name : %s, Id : %d, Age : %d\n", s.Name, s.ID, s.Age)
	}
}

func PrintArray(arr []int) {
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func main() {
	// 1. 整数数组降序排序
	arr := []int{5, 4, 3, 2, 7, 9, 1, 0}
	// 比较器：返回负数表示 a 排前面
	sort.Slice(arr, func(i, j int) bool {
		return arr[j] < arr[i] // 等价 Java arg1 - arg0
	})
	PrintArray(arr)
	fmt.Println("===========================")

	// 2. 学生数组排序
	student1 := &Student{Name: "A", ID: 4, Age: 40}
	student2 := &Student{Name: "B", ID: 4, Age: 21}
	student3 := &Student{Name: "C", ID: 3, Age: 12}
	student4 := &Student{Name: "D", ID: 3, Age: 62}
	student5 := &Student{Name: "E", ID: 3, Age: 42}

	students := []*Student{student1, student2, student3, student4, student5}

	// 第一条打印：ID 升序 + Age 降序
	fmt.Println("第一条打印")
	sort.Slice(students, func(i, j int) bool {
		return IdShengAgeJiangOrder(students[i], students[j]) < 0
	})
	for _, s := range students {
		fmt.Printf("%s,%d,%d\n", s.Name, s.ID, s.Age)
	}

	// 第二条打印：切片排序（同上逻辑）
	fmt.Println("第二条打印")
	studentList := []*Student{student1, student2, student3, student4, student5}
	sort.Slice(studentList, func(i, j int) bool {
		return IdShengAgeJiangOrder(studentList[i], studentList[j]) < 0
	})
	for _, s := range studentList {
		fmt.Printf("%s,%d,%d\n", s.Name, s.ID, s.Age)
	}

	// 第三条打印：TreeMap 有序结构（Go 使用 *sort.Slice + map 模拟）
	fmt.Println("第三条打印")
	student1 = &Student{Name: "A", ID: 4, Age: 40}
	student2 = &Student{Name: "B", ID: 4, Age: 21}
	student3 = &Student{Name: "C", ID: 4, Age: 12}
	student4 = &Student{Name: "D", ID: 4, Age: 62}
	student5 = &Student{Name: "E", ID: 4, Age: 42}

	// Go 模拟 TreeMap：map + 有序切片
	treeMap := make(map[*Student]string)
	treeMap[student1] = "我是学生1，我的名字叫A"
	treeMap[student2] = "我是学生2，我的名字叫B"
	treeMap[student3] = "我是学生3，我的名字叫C"
	treeMap[student4] = "我是学生4，我的名字叫D"
	treeMap[student5] = "我是学生5，我的名字叫E"

	// 按 ID 升序遍历
	var keys []*Student
	for k := range treeMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return IdAscendingComparator(keys[i], keys[j]) < 0
	})

	for _, s := range keys {
		fmt.Printf("%s,%d,%d\n", s.Name, s.ID, s.Age)
	}
}
