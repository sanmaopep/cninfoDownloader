package main

import (
	"cninfoDownloader/cninfoParser"
)

func main() {

	//var stockid string
	//var num int
	//fmt.Print("请输入股票号码(如 002586):")
	//fmt.Scanf("%s",&stockid)
	//fmt.Print("请输入要下载的文件数量(如200):")
	//fmt.Scanf("%s",&num)

	cninfoParser.GetDownloadList("002586",100);
}
