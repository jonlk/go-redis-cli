package main

import (
	"bufio"
	"fmt"
	"go-redis-cli/cache"
	"os"
	"strings"
)

func Listen() {

	reader := bufio.NewScanner(os.Stdin)
	listening := true

	client := cache.CreateClient()

	for listening {

		fmt.Print("Enter a cache operation - (g)et <key>, (s)et <key> <value>, (/q)uit: ")

		for reader.Scan() {

			input := strings.Split(reader.Text(), " ")
			inputLength := len(input)

			if inputLength == 1 && input[0] == "/q" {

				listening = false
				break

			} else if inputLength == 2 && input[0] == "g" {

				key := input[1]
				client.GetKeyValue(key)
				break

			} else if inputLength >= 3 && input[0] == "s" {

				key := input[1]
				remainingLength := inputLength - 2

				var s []string
				for i := 0; i < remainingLength; i++ {
					s = append(s, input[i+2])
				}

				value := strings.Join(s, " ")
				client.SetKeyValue(key, value)

				break

			} else {
				fmt.Println("Invalid input!")
				break

			}
		}
	}
}
