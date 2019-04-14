package main

import (
	"fmt" 
	"bufio"
	"os"
	"net/http"
	"io/ioutil"
)

func main() {  
fmt.Println("╔═╗┬ ┬┌─┐┬  ┬    ╔═╗┬┌┐┌┌┬┐┌─┐┬─┐")
fmt.Println("╚═╗├─┤├┤ │  │    ╠╣ ││││ ││├┤ ├┬┘")
fmt.Println("╚═╝┴ ┴└─┘┴─┘┴─┘  ╚  ┴┘└┘─┴┘└─┘┴└─\n")
fmt.Println("[=================]")
fmt.Println("Coder : Fajar Firdaus")
fmt.Println("Fb : Fajar Firdaus")
fmt.Println("IG : fajar_firdaus_7")
fmt.Println("YT : iTech7732")
fmt.Println("[=================]")


var input string
fmt.Print("Input : ")
fmt.Scan(&input)
	file, err := os.Open("wordlist.txt")
	if err != nil{
		fmt.Println("[!] 404 File not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		var scan = scanner.Text()
		resp, error := http.Get(input + scan)
		if error != nil{
			fmt.Println(input + scan + " [!] 404 Shell Not Found")
		}else{
			ioutil.ReadAll(resp.Body)
			fmt.Println(input + scan + " [+] Shell Found")
		}
	}
	if err := scanner.Err(); err != nil {
        fmt.Println("error")
    }
}
