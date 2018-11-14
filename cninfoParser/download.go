package cninfoParser

import (
	"cninfoDownloader/services"
	"fmt"
	"net/url"
)

const baseurl = "http://www.cninfo.com.cn"
const fulltext = "/cninfo-new/disclosure/szse/fulltext"
const step = 15

// "002586" 15
func GetDownloadList(stockid string, num int){
	turnnum := num / step

	for i := 0; i < turnnum; i++{
		result, err := services.PostFormData(
			baseurl + fulltext,
			url.Values{
				"stock": {stockid},
				"searchkey":{""},
				"category":{""},
				"tabName": {"latest"},
				"column":{"szse_sme"},
				"pageNum": {string(i+1)},
				"pageSize": {string(step)},
				"sortName":{""},
				"sortType":{""},
				"limit":{""},
				"seDate":{""},
			},
		)
		if err != nil {
			fmt.Println(result)
		}
	}




}
