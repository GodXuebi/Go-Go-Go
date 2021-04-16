package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	/*如果 jobs 已经关闭了，并且通道中所有的值都已经接收完毕，
	那么 more 的值将是 false。如果jobs不为空，即使close，返回的more仍为true*/
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	/*关闭 一个通道意味着不能再向这个通道发送值了。
	这个特性可以用来给这个通道的接收方传达工作已经完成的信息。*/
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent jobs", j)
	}
	close(jobs)

	fmt.Println("sent all jobs")

	<-done
}
