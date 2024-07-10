package main

import (
	"bufio"
	"os"
	"toylsp/rpc"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(msg any) {

}
