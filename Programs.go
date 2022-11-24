/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-11-16 09:16:04
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-11-17 11:20:37
 * @FilePath: \Memo_frame\Program\Programs.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */

package Programs

import (
	"demo/CSV"
	"fmt"
	// "strconv"
)

type Debt struct {
	Name   string
	Borrow float64
	Refund float64
	Total  float64
	Valid  string
	key    string
	loop   bool
}

func NewDebt() *Debt {
	return &Debt{
		key:    "",
		loop:   false,
		Name:   "",
		Borrow: 0,
		Refund: 0,
		Total:  0,
		Valid:  "y",
	}
}

func (this *Debt) exit() { 
	
}

func (this *Debt) ShowCSV() {
	CSV.ShowCSV()
}

func (this *Debt) Borrows() {
	WriterCsv := csv.NewWriter(File)
	err1 := WriterCsv.Write(slice)
	if err1 != nil {
		fmt.Println("写入文件失败")
}
}
func (this *Debt) Refunds() {
	
}

func (this *Debt) Totals() {

}

func (this *Debt) IsDebt() {
    
}

func (this *Debt) MainMenu() {
	for {
		fmt.Println("------------------ 债务备忘录软件-----------------")
		fmt.Println("                   1 借债")
		fmt.Println("                   2 还债/存入")
		fmt.Println("                   3 总额和(模拟后台查看记录)")
		fmt.Println("                   4 欠债查询")
		fmt.Println("                   5 总览(模拟后台查看记录)")
		fmt.Println("                   6.退出系统 ")
		fmt.Println("请选择(1-6):")
		fmt.Println("首次借款后，下次可还款存入")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.Borrows()
		case "2":
			this.Refunds()
		case "3":
			this.Totals()
		case "4":
			this.IsDebt()
		case "5":
			this.ShowCSV()
		case "6":
			this.exit()
		default:
			fmt.Println("您的输入不正确，请按提示输入正确的选项..")
		}
		if this.loop {
			this.Total = 0
			break
		}
	}
}
