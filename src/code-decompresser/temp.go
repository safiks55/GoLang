package main

import "fmt"

func main() {
	alpha := []string{"a", "b", "c"}
	dec := []string{"010", "100", "101"}
	temp := []string{"010", "010", "010", "100", "101", "101", "101"}

	mp := make(map[string]string)
	{
		for i := 0; i < len(alpha); i++ {
			mp[dec[i]] = alpha[i]
		}
	}
	var t string
	message := ""

	for i := range temp {

		t = mp[temp[i]]
		message += t
	}
	fmt.Println(message)
}
