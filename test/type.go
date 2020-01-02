package main

import "fmt"

/*//预先编译好要提取信息的正则表达式
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)

func main() {
	var str = `<div class="m-btn purple" data-v-8b1eac0c>44岁</div>`
	ParseProfile([]byte(str))
}
func ParseProfile(contents []byte) {
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		panic(err)
	}
	fmt.Println(age)
}

//封装正则匹配函数,提取字符串
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}*/
func main() {
	strings := []string{
		"asd",
		"zxc",
		"zxc",
		"qwe",
	}
	m := make(map[string]int)

	var i = 0
	for _, v := range strings {
		if isDuplicate(v, m) {
			i++
		}

	}
	fmt.Println(i)

	m2 := make(map[string]bool)
	fmt.Println(m2["qq"])
}

func isDuplicate(url string, m map[string]int) bool {
	if m[url] > 0 {
		fmt.Println(url, m[url])
		return false
	} else {
		m[url]++
	}
	return true
}
