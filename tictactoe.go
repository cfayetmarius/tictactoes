package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
	"strings"
)


type Case struct{
	content string
}

func (c *Case) changecontent(ncont string) {
	c.content = ncont
}
var caselist [9]Case

type Player struct {
	name string
	sign string
}

func (p Player) play() 	{
	scanner := bufio.NewReader(os.Stdin)
	printboard(caselist)
	fmt.Println("")
	fmt.Println(p.name, "Your turn !")
	fmt.Println("Choose a number to play between 123456789")
	text, _ := scanner.ReadString('\n')
	trans, err := strconv.Atoi(strings.Replace(text, "\r\n", "",  1))
	if err != nil {
		log.Fatal(err)
	}
	trans = trans - 1
	switch caselist[trans].content {
		case "#" :
			caselist[trans].changecontent(p.sign)
		default : 
			fmt.Println("Play something else")
			p.play()
	}
}



func printboard(cselist [9]Case) {
	fmt.Printf(cselist[0].content + "|" + cselist[1].content + "|" + cselist[2].content + "\n" + "-----" + "\n" + cselist[3].content + "|" + cselist[4].content + "|" + cselist[5].content + "\n" + "-----" + "\n" + cselist[6].content + "|" + cselist[7].content + "|" + cselist[8].content)
}

func initcases() {
	for i:=0; i!=9; i++  {
		caselist[i] = Case{"#"}
	}
}

func getstate() string{
	if caselist[0].content == caselist[1].content && caselist[1].content == caselist[2].content && caselist[0].content != "#" {
		return(caselist[0].content)
	}
	if caselist[3].content == caselist[4].content && caselist[4].content == caselist[5].content && caselist[3].content != "#" {
		return(caselist[3].content)
	}
	if caselist[6].content == caselist[7].content && caselist[7].content == caselist[8].content && caselist[6].content != "#" {
		return(caselist[6].content)
	}
	if caselist[0].content == caselist[3].content && caselist[3].content == caselist[6].content && caselist[0].content != "#" {
		return(caselist[0].content)
	}
	if caselist[1].content == caselist[4].content && caselist[4].content == caselist[7].content && caselist[1].content != "#" {
		return(caselist[1].content)
	}
	if caselist[2].content == caselist[5].content && caselist[5].content == caselist[8].content && caselist[2].content != "#" {
		return(caselist[2].content)
	}
	if caselist[0].content == caselist[4].content && caselist[4].content == caselist[8].content && caselist[0].content != "#" {
		return(caselist[0].content)
	}
	if caselist[2].content == caselist[4].content && caselist[4].content == caselist[6].content && caselist[2].content != "#" {
		return(caselist[2].content)
	}
	for i:=0; i !=9; i++ {
		if caselist[i].content == "#" {
			return("NOTEND")
		}
	}
	return("DRAW")
}


func main() {
	initcases()
	Jhon := Player{"Jhon", "X"}
	Jack := Player{"Jack", "O"}
	for true {
		if getstate() == "NOTEND" {
			Jhon.play()
			for i:=0; i<100; i++ {
				fmt.Println("\n")
			}
		}
		if getstate() == "NOTEND" {
			Jack.play()			
			for i:=0; i<100; i++ {
				fmt.Println("\n")
			}
		}
		if getstate() != "NOTEND" {
			break
		}
	}
	switch getstate() {
	case "DRAW" :
		fmt.Println("This is a draw !")
	case "O" :
		fmt.Println("Jack won the game !")
	case "X" :
		fmt.Println("Jhon won the game !")
	}
}