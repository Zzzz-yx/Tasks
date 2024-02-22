package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("\t欢迎来到猜拳游戏！")
	rand.Seed(time.Now().UnixNano()) 
	var choice, myn string
	var computerchoice string      
	var mychoice string            
	var mycount, computercount int 
	var count int                  
	FistMap := map[int]string{0: "石头", 1: "剪刀", 2: "布"}
	fmt.Println("-------------------------------------")
	fmt.Print("游戏规则：三局两胜\n")
	fmt.Print("注意：1为继续玩     0为退出\n")
	fmt.Println("-------------------------------------")
	for {
		fmt.Printf("是否继续玩(1.继续玩,0.退出游戏):")
		fmt.Scan(&choice)
		choiceint, err := strconv.Atoi(choice)
		for err != nil || (choiceint < 0 || choiceint > 1) {
			fmt.Printf("你的输入不合法！请输入(0-1)之间的数：")
			fmt.Scan(&choice)
			choiceint, err = strconv.Atoi(choice)
		}
		if choiceint == 1 {
			for i := 0; i < 3; i++ {
				n := rand.Intn(3) 
				fmt.Printf("请输入你的出招(0为石头,1为剪刀,2为布):")
				fmt.Scan(&myn)
				mynint, errint := strconv.Atoi(myn)
				for errint != nil || (mynint > 2 || mynint < 0) {
					fmt.Printf("你的出招有误，重新输入(0为石头,1为剪刀,2为布):")
					fmt.Scan(&myn)
					mynint, errint = strconv.Atoi(myn)
				}
				computerchoice = FistMap[n]
				mychoice = FistMap[mynint]
				if (computerchoice == "石头" && mychoice == "布") || (computerchoice == "剪刀" && mychoice == "石头") || (computerchoice == "布" && mychoice == "剪刀") {
					fmt.Printf("机器出招（%s） --  我出招（%s）\n", computerchoice, mychoice)
					mycount += 1
				} else if (computerchoice == "石头" && mychoice == "剪刀") || (computerchoice == "剪刀" && mychoice == "布") || (computerchoice == "布" && mychoice == "石头") {
					fmt.Printf("机器出招（%s） --我出招（%s）\n", computerchoice, mychoice)
					computercount += 1
				} else {
					fmt.Printf("机器出招(%s) --我出招(%s)\n", computerchoice, mychoice)
				}
			}
			count += 1
			if mycount > computercount {
				fmt.Println("----最终还是我赢了呀，啊哈哈哈！----")
			} else if mycount < computercount {
				fmt.Println("----最终我还是输了，呜呜呜呜！----")
			} else {
				fmt.Println("----最终达成了个平手，下次再战----")
			}
		}
		if choiceint == 0 {
			fmt.Printf("拜拜了,本次你一共玩了%d次", count)
			break
		}
	}
}

