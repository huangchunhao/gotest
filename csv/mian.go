package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	//写入数据
	filename := "./csv/test.csv"
	columns := [][]string{{"姓名", "电话", "公司", "职位", "加入时间"}, {"1", "2", "刘犇,刘犇,刘犇", "4", "5"}}
	ExportCsv(filename,columns)

//以遍历的方式读取数据
	file, err := os.Open("./csv/test.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(record) // record has the type []string
	}

//以ALL方式读取数据
	file2, err := os.Open("./csv/test.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file2.Close()
	reader2 := csv.NewReader(file2)
	content,_ :=reader2.ReadAll()
	fmt.Println(content)

	fmt.Println(content[0])
	fmt.Println(content[0][0])
	fmt.Println(content[1])
	fmt.Println(content[1][0])

}

func ExportCsv(filePath string, data [][]string) {
	fp, err := os.Create(filePath) // 创建文件句柄
	if err != nil {
		log.Fatalf("创建文件["+filePath+"]句柄失败,%v", err)
		return
	}
	defer fp.Close()
	fp.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(fp) //创建一个新的写入文件流
	w.WriteAll(data)
	w.Flush()
}
