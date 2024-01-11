package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	vmessURL  = "https://aliilapro.github.io/v2rayNG-Config/server.txt"
	ssURL     = "https://raw.githubusercontent.com/barry-far/V2ray-Configs/main/Splitted-By-Protocol/ss.txt"
	ssrURL    = "https://raw.githubusercontent.com/barry-far/V2ray-Configs/main/Splitted-By-Protocol/ssr.txt"
	trojanURL = "https://raw.githubusercontent.com/barry-far/V2ray-Configs/main/Splitted-By-Protocol/trojan.txt"
	vlessURL  = "https://raw.githubusercontent.com/barry-far/V2ray-Configs/main/Splitted-By-Protocol/vless.txt"
	vmssURL   = "https://raw.githubusercontent.com/barry-far/V2ray-Configs/main/Splitted-By-Protocol/vmess.txt"
)

func main() {
	printHeader()

prog:
	fmt.Println("Choose a proxy type:")
	fmt.Println("1) MIX of vmess and vless #priv8") // priv8_ripo
	fmt.Println("2) ss")
	fmt.Println("3) ssr")
	fmt.Println("4) trojan")
	fmt.Println("5) vless")
	fmt.Println("6) vmess")

	var choice int
	fmt.Print("Enter your choice [1-6]: ")
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > 6 {
		log.Fatal("Invalid choice")
		goto prog
	}

	var (
		url      string
		fileName string
	)

	switch choice {
	case 1:
		url = vmessURL
	case 2:
		url = ssURL
	case 3:
		url = ssrURL
	case 4:
		url = trojanURL
	case 5:
		url = vlessURL
	case 6:
		url = vmessURL
	}

	fileName = generateFileName()

	err = saveFile(url, fileName)
	if err != nil {
		log.Fatal("Failed to save file:", err)
	}

	fmt.Println("File saved successfully as", fileName)
	fmt.Println("Press Enter to exit")
	_, err = fmt.Scanln()
	if err != nil {
		return
	}
}

func generateFileName() string {
	now := time.Now()
	return fmt.Sprintf("Server_Choice_%02d-%02d-%02d.txt", now.Hour(), now.Minute(), now.Second())
}

func saveFile(url, fileName string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func printHeader() {
	header := `
	%sVV    VV%s  %sEEEEEEE%s   %sXXX    XXX%s  %sSSSSSS%s   %sXX    XX%s
	%sVV    VV%s  %sE%s           %sXX  XX%s   %sSSSS%s      %sXX   XX%s
	%sVV    VV%s  %sEEEEEEE%s     %sXXXXX%s    %sSSS%s        %sXXXXX%s
	 %sVV  VV%s   %sEE%s           %sXXXX%s      %sSSS%s      %sXXXX%s
	  %sVVVV%s    %sEE%s           %sXX  XX%s    %sSSSS%s    %sXX   XX%s
	   %sVV%s     %sEEEEEEE%s    %sXXX    XXX%s %sSSSSS%s   %sXXX   XXX%s
`
	// Color codes
	red := "\033[1;31m"
	green := "\033[1;32m"
	yellow := "\033[1;33m"
	blue := "\033[1;34m"
	purple := "\033[1;35m"
	reset := "\033[0m"

	colorfulHeader := fmt.Sprintf(header, red, reset, green, reset, yellow, reset, blue, reset, purple, reset,
		red, reset, green, reset, yellow, reset, blue, reset, purple, reset,
		red, reset, green, reset, yellow, reset, blue, reset, purple, reset,
		red, reset, green, reset, yellow, reset, blue, reset, purple, reset,
		red, reset, green, reset, yellow, reset, blue, reset, purple, reset,
		red, reset, green, reset, yellow, reset, blue, reset, purple, reset)

	fmt.Print(colorfulHeader)
}
