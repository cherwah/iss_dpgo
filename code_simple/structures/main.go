package main

import "structures/robot"

func main() {
	robot := robot.Robot{Name: "Robbie", Battery: 10, Speed: 5}

	robot.Activate()
	for i := 0; i < 10; i++ {
		robot.Move()
	}
}
