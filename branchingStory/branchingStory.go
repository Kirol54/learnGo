package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func (node *storyNode) printStory(depth int) {
	for i := 0; i < depth*2; i++ {
		fmt.Print("  ")
	}
	fmt.Print(node.text)
	fmt.Println()
	if node.yesPath != nil {
		node.yesPath.printStory(depth + 1)
	}
	if node.noPath != nil {
		node.noPath.printStory(depth + 1)
	}
}

func (node *storyNode) play() {
	fmt.Println(node.text)
	if node.yesPath != nil && node.noPath != nil {
		scanner := bufio.NewScanner(os.Stdin)

		for {
			scanner.Scan()
			answer := scanner.Text()
			if answer == "yes" {
				node.yesPath.play()
				break
			} else if answer == "no" {
				node.noPath.play()
				break
			} else {
				fmt.Println("this answer was not an option, please answer yes or no")
			}
		}
	}
}

func main() {

	root := storyNode{"you are at the entrance to a cave. Do you want to enter?", nil, nil}
	winning := storyNode{"you have won !", nil, nil}
	losing := storyNode{"you have lost", nil, nil}
	root.yesPath = &losing
	root.noPath = &winning

	root.printStory(0)
}
