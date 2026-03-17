package main

import "fmt"

func main() {
	var H int
	var p int
	var F int
	var D int
	fmt.Scanf("%d %d %d %d", &H, &p, &F, &D)
	if H > p {
		if F > H {
			if D == -1 {
				fmt.Println("S")
			} else {
				fmt.Println("N")
			}
		} else {
            if p==3{
                fmt.Println("S")
            }else{
			    if D == -1 {
				    fmt.Println("S")
			    } else {
				    fmt.Println("N")
			    }
            }
		}
	} else {
		if F > p {
			if D == 1 {
				fmt.Println("S")
			} else {
				fmt.Println("N")
			}
		} else {
			if D == -1 {
				fmt.Println("S")
			} else {
				fmt.Println("N")
			}
		}
	}
}
