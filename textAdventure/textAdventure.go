package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choices struct {
	cmd         string
	description string
	nextNode    *storyNode
	nextChoice  *choices
}

type storyNode struct {
	text    string
	choices *choices
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choices{cmd, description, nextNode, nil}
	if node.choices == nil {
		node.choices = choice
	} else {
		currentChoice := node.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	currentChoice := node.choices
	for currentChoice != nil {
		fmt.Println(currentChoice.cmd, ":", currentChoice.description)
		currentChoice = currentChoice.nextChoice
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	currentChoice := node.choices
	for currentChoice != nil {
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd) {
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Println("sorry, I didn't understand that.")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
		You are in large chamger, deep underground.You see three passages leading out. 
		A north passage leads into darkness.
		To the south, a passage appears to head to some kind of a room.
		The eastern passages appers flat and well traveled.
	`}
	darkRoom := storyNode{text: "It is pitch black. You cannot see a thing."}

	darkRoomLit := storyNode{text: "The dark passage is now lit by your lantern. You can continue north or head back south"}

	monster := storyNode{text: "While stumbling around in the darkness, you are eaten by a monster"}

	trap := storyNode{text: "You head down the well traveled path when suddenly a trap door opens and you fall into a pit"}

	treasure := storyNode{text: "You arrive at a small chamber, filled with treasure!"}

	npcRoom := storyNode{text: "Flame golem  appears in front of you. What do you do ?"}

	golemIce := storyNode{text: "Flame golem is vulnerable to ice. Your spell destroys him."}

	golemWind := storyNode{text: `
		Wind spells allow flame golem to create fire tornado.
		You have been killed
	`}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &npcRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try to go back south", &monster)
	darkRoom.addChoice("O", "Turn on lantern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	npcRoom.addChoice("A", "Use ice spells ", &golemIce)
	npcRoom.addChoice("B", "Use wind spells ", &golemWind)

	golemIce.addChoice("B", "Go back", &start)

	start.play()
	fmt.Println()
	fmt.Println("The End.")
}
