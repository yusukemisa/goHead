package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	//引数処理
	n := flag.Int("n", 10, "出力行数")
	c := flag.Int64("c", 0, "出力バイト数")
	flag.Parse()
	// n c同時オプション指定禁止
	if flag.NFlag() != 1 {
		fmt.Fprintf(os.Stderr, "%v: can't combine line and byte counts\n", os.Args[0])
		flag.Usage()
		return
	}
	files := flag.Args()
	//fmt.Printf("len=%v,files=%v", len(files), files)

	//ファイル読み込み
	readFiles(files, *n, *c)
}

func readFiles(files []string, n int, c int64) {
	for i, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: No such file or directory\n", fileName)
			file.Close()
			continue
		}
		//複数ファイル指定時のヘッダー出力
		printHeader(i, len(files), fileName)

		readFile(bufio.NewReader(file), n, c)
		//forの全ループがスコープになるのでdeferで閉じるのよくないみたいなので手動で都度クローズ
		file.Close()
	}
}

func printHeader(i int, length int, fileName string) {
	if 0 < i {
		fmt.Println("")
	}
	if 1 < length {
		fmt.Printf("==> %v <==\n", fileName)
	}
}

func readFile(reader *bufio.Reader, n int, c int64) {
	if 0 < c {
		//バイト数指定出力
		io.CopyN(os.Stdout, reader, c)
	} else {
		//行数指定出力
		for i := 0; i < n; i++ {
			line, err := reader.ReadString('\n')
			fmt.Print(line)
			if err == io.EOF {
				break
			}
		}
	}
}
