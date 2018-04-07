package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//引数処理
	n := flag.Int("n", 10, "出力行数")
	flag.Parse()

	fileName := flag.Arg(0)
	//fmt.Printf("n=%v,fileName=%v\n", *n, fileName)

	//ファイル読み込み
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(fileName + "が開けませんでした")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for i := 0; i < *n; i++ {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			break
		}
	}
}
