package main

import(
	"fmt"
	"crypto/sha1"
	"os"
	"log"
)

func main(){
	var codes []string

	h := sha1.New()
	fo, err := os.Create("section2")
	if err != nil{
		log.Fatal(err)
		os.Exit(2)
	}

	for i := 0 ; i < 250 ; i++{
		h.Write([]byte(string(i)))
		codes = append(codes, fmt.Sprintf("MQ%x", h.Sum(nil) )[:10]);
	}

	
	for _,code := range codes{
		fmt.Println(code)
		fo.Write([]byte(code + "\n"))
	}

	fo.Close();

}
