package main

import "fmt"

func main() {
	//level1()
	level2()
}

func level1() {
	for y := 1.5; y > -1.5; y -= 0.1 {
		for x := -1.5; x < 1.5; x += 0.05 {
			a := x*x + y*y - 1
			b := a*a*a - x*x*y*y*y
			if b <= 0.0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func level2() {
	var q rune
	for true {
		fmt.Print("计算机系统，请输入数字：")
		calculate()
		fmt.Print("输入q退出,输入任意键继续计算。")
		fmt.Scanf("%c\n", &q)
		if q == 'q' {
			print("欢迎下次光临")
			break
		}
	}
}
func calculate() {
	var (
		a   float64
		b   float64
		res float64
		c   rune
	)
	fmt.Scanf("%f%c%f", &a, &c, &b)
	switch c {
	case '+':
		res = a + b
		fmt.Printf("The result is : %f\n", res)
	case '-':
		res = a - b
		fmt.Printf("The result is : %f\n", res)
	case '*':
		res = a * b
		fmt.Printf("The result is : %f\n", res)
	case '/':
		if b == 0 {
			fmt.Print("除数不能为0")
		} else {
			res = a / b
			fmt.Printf("The result is : %f\n", res)
		}
	default:
		fmt.Print("输入格式错误")
	}
}
