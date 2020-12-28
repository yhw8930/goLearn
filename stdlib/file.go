package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//readAll()
	//readByBuffer()
	//writeAll()
	//writeByBuffer()
	writeByIoutil()
}

func readFromFile() {
	file, err := os.Open("./stdlib/test.text")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer file.Close()
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d个字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}

func readAll() {
	file, err := os.Open("./stdlib/test.text")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer file.Close()
	for {
		var tmp = make([]byte, 128)
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Printf("读取了%d个字节数据\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

func readByBuffer() {
	file, err := os.Open("./stdlib/test.text")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

func readByIoutil() {
	content, err := ioutil.ReadFile("./stdlib/test.text")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func writeAll() {
	file, err := os.OpenFile("./stdlib/test.text", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello gogo\n"
	file.Write([]byte(str))        //写入字节切片数据
	file.WriteString("hello go\n") //直接写入字符串数据
}

func writeByBuffer() {
	file, err := os.OpenFile("./stdlib/test2.text", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

func writeByIoutil() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./stdlib/test2.text", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

// CopyFile 拷贝文件函数
func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}
