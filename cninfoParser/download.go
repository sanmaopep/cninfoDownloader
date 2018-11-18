package cninfoParser

import (
	"cninfoDownloader/services"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// 分析网页：http://www.cninfo.com.cn/information/companyinfo_n.html?fulltext?szsme002586
const baseurl = "http://www.cninfo.com.cn"
const fulltext = "/cninfo-new/disclosure/szse/fulltext"
const step = 10

// https://mholt.github.io/json-to-go/
type announcement struct {
	ID                    interface{} `json:"id"`
	SecCode               string      `json:"secCode"`
	SecName               string      `json:"secName"`
	OrgID                 string      `json:"orgId"`
	AnnouncementID        string      `json:"announcementId"`
	AnnouncementTitle     string      `json:"announcementTitle"`
	AnnouncementTime      int64       `json:"announcementTime"`
	AdjunctURL            string      `json:"adjunctUrl"`
	AdjunctSize           int         `json:"adjunctSize"`
	AdjunctType           string      `json:"adjunctType"`
	StorageTime           int64       `json:"storageTime"`
	ColumnID              string      `json:"columnId"`
	PageColumn            interface{} `json:"pageColumn"`
	AnnouncementType      string      `json:"announcementType"`
	AssociateAnnouncement interface{} `json:"associateAnnouncement"`
	Important             bool        `json:"important"`
	BatchNum              interface{} `json:"batchNum"`
	AnnouncementContent   interface{} `json:"announcementContent"`
	AnnouncementTypeName  string      `json:"announcementTypeName"`
}

type downloadListAPIResp struct {
	ClassifiedAnnouncements [][]announcement `json:"classifiedAnnouncements"`
	TotalSecurities   int                    `json:"totalSecurities"`
	TotalAnnouncement int                    `json:"totalAnnouncement"`
	TotalRecordNum    int                    `json:"totalRecordNum"`
	Announcements     interface{}            `json:"announcements"`
	CategoryList      []struct {
		Value0 string `json:"value0"`
		Value1 string `json:"value1"`
	} `json:"categoryList"`
	HasMore    bool `json:"hasMore"`
	Totalpages int  `json:"totalpages"`
}
// 格式 "002586" 15
func getDownloadList(stockid string, num int) []announcement {
	fmt.Println("正在获取下载文件列表........")


	turnnum := num / step

	//var respList []classifiedAnnouncement
	var resp downloadListAPIResp
	var announcements []announcement

	for i := 0; i < turnnum; i++{
		result, err := services.PostFormData(
			baseurl + fulltext,
			url.Values{
				"stock": {stockid},
				"searchkey":{""},
				"category":{""},
				"tabName": {"latest"},
				"column":{"szse_sme"},
				"pageNum": {strconv.Itoa(i+1)},
				"pageSize": {strconv.Itoa(step)},
				"sortName":{""},
				"sortType":{""},
				"limit":{""},
				"seDate":{""},
			},
		)
		if err == nil {
			json.Unmarshal(result,&resp)
			for _, subArray := range resp.ClassifiedAnnouncements {
				for _, v := range subArray {
					announcements = append(announcements,v)
				}
			}

			fmt.Printf("get annoucement list %d / %d\n",(i+1)*step, num)
		}else {
			panic("get annoucement list fail")
		}
	}
	return announcements
}


func DownloadPdfs(stockid string, num int, downloadDirectory string){
	announcements := getDownloadList(stockid, num)

	if err := services.CreatePath(downloadDirectory); err != nil {
		panic(err)
	}

	fmt.Println("开始下载文件........")
	announcementsLen := len(announcements)

	for i, v := range announcements {
		url := baseurl + "/" + v.AdjunctURL
		filepath := downloadDirectory + "/" + v.AnnouncementTitle + ".pdf"
		fmt.Println("---------------------------")
		fmt.Printf("download pdf %d/%d\n", i+1 , announcementsLen)
		fmt.Printf("name: %s\n",filepath)
		fmt.Printf("url: %s\n",url)

		if err := services.DownloadFile(filepath,url); err != nil {
			panic(err)
		}
	}
}
