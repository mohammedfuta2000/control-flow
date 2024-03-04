package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("Begin initializzation")
}

func main() {

	if z := rand.Intn(10); z <= 10 {
		fmt.Println("the result is less than the source")
	}

	myVals := map[string]string{"man": "woman", "black": "white"}
	if val,ok:=myVals["man"]; ok{
		fmt.Println(val)
	}

	// switch{
		// case x < 42:
			// 	fmt.Println("less than 42")
			// case x>42:
				// 	fmt.Println("greater than 42")
				// default:
					// 	fmt.Println("this is the default")
					// }
					
	x:= 24
	switch x{
	case 24:
		fmt.Println("this is 24")
		fallthrough
	case 30:
		fmt.Println("this is 30")
		fallthrough
	case 40:
		fmt.Println("this is 40")
		
	default:
		fmt.Println("this is default")
	}


	// SELECT STATEMEMNT
	ch1:= make(chan int)
	ch2:= make(chan int)

	d1:=time.Duration(rand.Int63n(250))
	d2:=time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1*time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2*time.Millisecond)
		ch2 <- 76
	}()
	// HOW DO I CATCH ALL
	select{
	case v1 := <-ch1:
		fmt.Println("value from channel 1:",v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2:",v2)
	}

	for i:=0; i<10;i++{
		if i%2!=0 {
			continue
		}
		if i==6 {
			break
		}
		fmt.Println(i)
	}

	for _,v:= range myVals{
		fmt.Println(v)
	}
}
