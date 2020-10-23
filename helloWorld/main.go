package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	//level0()
	//level1()
	//level2()
	//level3()
	//level4()
	level5()
}
func level0(){
	fmt.Println("HelloWorld")
}
func level1(){
	fmt.Println("Input:")
	var a int
	fmt.Scanf("%d",&a)
	var b string
	fmt.Scanf("%s",&b)
	var c int
	fmt.Scanf("%d",&c)
	fmt.Println("Output:")
	switch b{
	case "+":fmt.Printf("%d",a+c)
	case "-":fmt.Printf("%d",a-c)
	case "*","x","X":fmt.Printf("%d",a*c)
	case "/":fmt.Printf("%d",a/c)
	}
}
func level2(){
	signin := map[string]string{"abc":"def","123":"456","素律":"2333"}
	var account string
	var password string
	fmt.Print("请输入账号：")
	fmt.Scanf("%s",&account)
	fmt.Print("请输入密码：")
	fmt.Scanf("%s",&password)
	epw,exit := signin[account]
	if(exit){
		if(password==epw){
			println("登录成功")
		}else{
			println("密码错误")
		}
	}else{
		println("此账号不存在")
	}
}
func level3(){
	var nums []int
	var num int
	fmt.Print("输入一串数字求最大值（输入0结束）：\n")
	fmt.Scanf("%d",&num)
	for i:=0;num!=0;i++{
		nums = append(nums, num)
		fmt.Scanf("%d",&num)
	}
	var temp int
	for m,_:= range nums{
		for n:=m+1;n<len(nums);n++{
			if(nums[m]<nums[n]){
				temp = nums[m]
				nums[m] = nums[n]
				nums[n] = temp
			}
		}
	}
	fmt.Printf("最大值为：%d\n",nums[0])
	fmt.Printf("数组从大到小排序为：%v",nums)
}
func level4(){
	r := gin.Default()
	r.GET("/",func(c *gin.Context){
		c.String(200,"HelloWorld-素律")
	})
	r.Run()
}
func level5(){
	fmt.Print("输入一串字符来判断是否为回文串：")
	var target string
	var judge bool
	fmt.Scanf("%s",&target)
	var length int = len(target)
	for i:=0;i<length/2;i++{
		if(target[i]!=target[length-i-1]){
			judge = true
			break
		}
	}
	if(judge){
		println("该字符串不是回文串")
	}else{
		println("该字符串是回文串")
	}
}

