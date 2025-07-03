package main

import "fmt"

var version string // ビルド時にIdflagsフラグ経由でバージョンを埋め込むための変数

func main() {
	fmt.Printf("Exmaple %s\n", version)
}