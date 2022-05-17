package main

import (
	"fmt"
	m "math/rand"
	t "time"
)

// getLikesAndShares return two ints
func getLikesAndShares(postId int) (int, int) {
	var likesForPost, sharesForPost int
	switch postId {
	case 1:
		likesForPost = 5
		sharesForPost = 7
	case 2:
		likesForPost = 3
		sharesForPost = 11
	case 3:
		likesForPost = 22
		sharesForPost = 1
	case 4:
		likesForPost = 7
		sharesForPost = 9
	}
	fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)
	// return for likesForPost and sharesForPost
	return likesForPost, sharesForPost
}

func main() {
	var likes, shares int
	m.Seed(t.Now().UnixNano())
	rr := m.Intn(3)

	likes, shares = getLikesAndShares(rr)
	//fmt.Println("******************")
	//fmt.Println(rr)
	//fmt.Println("******************")

	/* if likes < 1 { //-----------------I need to fix it xD
	   fmt.Println("No cookies today.")
	   }*/

	if likes < 5 {
		fmt.Println("Well we have some likes.")
	}
	if likes >= 5 {
		fmt.Println("Woohoo! We got some likes.")
	}
	if shares > 10 {
		fmt.Println("We went viral!")
	}
}
