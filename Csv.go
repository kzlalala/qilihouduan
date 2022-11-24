/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-11-16 09:37:46
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-11-17 10:12:46
 * @FilePath: \Memo_frame\CSV\Csv.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */

package CSV

import (
	"encoding/csv"
	"fmt"
	"go/types"

	// "fmt"
	"log"
	"os"
)

func WriteCSV(C [][]string) {
	fmt.Println("请输入文件名记录，输入9退出")
	var nam string
	fmt.Scanln(&nam)
	e :=".csv"
	name = nam + e
	if name=="0"{
	return 0
	}
	var slice []string
	for i := 0; i < 2; i++ {
		var a string
		fmt.Scan(&a)
		slice = append(slice, a)
	}
	//创建csv文件
	f, err := os.OpenFile("Debt.csv", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	//异步管理
	defer f.Close()
	// 写入UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	//创建一个新的写入文件流
	w := csv.NewWriter(f)
    data:=[][]string{
		{"1", "张三", "2000.720000","1500.720000","-500.000000","y"},
		{"2", "李四","500.000000","300.000000","-200.000000","y"},
		{"3", "pig","500.000000","500.000000","0.000000","n"},
		{"4", "fzf404","50.000000","30.000000","-20.000000","y"},
		{"5", "寇卓","21599.000000","5600.000000","-20499.000000","y"},
	}
	//写入数据
	w.WriteAll(data)
	w.Flush()
}


func ReadCSV() string [][] {
	fmt.Printfln("输入文件名查询，输入9退出")
	var nam string
	fmt.Scanln(&nam)
	e :=".csv"
	name = nam + e
	if name=="0"{
	return 0
	}

	num := 0
	for i := 0; i < sz-1; i++ {
		if n[i+1] == "+" {
			num += p[i+1]
		} else {
			num -= p[i+1]
		}
	}


	fileName := "Debt.csv"
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}

	return content
}

func ShowCSV() {
	fileName := "Debt.csv"
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}

}
