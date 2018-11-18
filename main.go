package main

import (
	"cninfoDownloader/cninfoParser"
	"fmt"
)

func main() {

	var stockid,downloadPath string
	var num int

	fmt.Print("请输入股票号码(默认 002586):")
	if n,_ := fmt.Scanf("%s",&stockid); n == 0 {
		stockid = "002586"
	}

	fmt.Print("请输入要下载通知的页数(默认 500):")
	if n,_ := fmt.Scanf("%d",&num); n == 0 {
		num = 50
	}

	fmt.Printf("请输入要下载的文件路径(默认 D:/cninfo/，需要自己创建好，注意反斜杠方向):")
	if n,_ := fmt.Scanf("%s",&downloadPath); n == 0 {
		downloadPath = "D:/cninfo/"+stockid+"/"
	}

	cninfoParser.DownloadPdfs(stockid,num, downloadPath);
}
