package main

import "os"

type ClassCSS struct {
	name string	// class名
	num int		// class数
}

func main() {

}

func parser(filename string)error  {
	file, err := os.Open(filename)

	if err != nil{
		return err
	}


}

func convertClass(){

}