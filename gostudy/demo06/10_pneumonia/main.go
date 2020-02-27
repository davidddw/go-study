package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	runewidth "github.com/mattn/go-runewidth"
)

const (
	url = "https://ncov.dxy.cn/ncovh5/view/pneumonia"
)

type Provinces []Province

type Province struct {
	ProvinceName          string `json:"provinceName"`
	ProvinceShortName     string `json:"provinceShortName"`
	CurrentConfirmedCount int    `json:"currentConfirmedCount"`
	ConfirmedCount        int    `json:"confirmedCount"`
	SuspectedCount        int    `json:"suspectedCount"`
	CuredCount            int    `json:"curedCount"`
	DeadCount             int    `json:"deadCount"`
	Comment               string `json:"comment"`
	LocationId            int    `json:"locationId"`
	Citys                 []City `json:"cities"`
}

type City struct {
	CityName              string `json:"cityName"`
	CurrentConfirmedCount int    `json:"currentConfirmedCount"`
	ConfirmedCount        int    `json:"confirmedCount"`
	SuspectedCount        int    `json:"suspectedCount"`
	CuredCount            int    `json:"curedCount"`
	DeadCount             int    `json:"deadCount"`
	LocationId            int    `json:"locationId"`
}

func formatStr(s string, width int) string {
	return runewidth.FillRight(s, width)
}

const (
	provCol1       = 20
	provCol2       = 10
	cityCol1       = 18
	cityCol2       = 10
	timeFormatFile = "20060102"
)

func (ps *Provinces) getPrintData() string {
	var str []string
	title := []string{
		formatStr("地区", provCol1),
		formatStr("现存确诊", provCol2),
		formatStr("累计确诊", provCol2),
		formatStr("死亡", provCol2),
		formatStr("治愈", provCol2),
	}
	str = append(str, strings.Join(title, ""))
	str = append(str, strings.Repeat("=", 60))
	for _, p := range *ps {
		s1 := formatStr(p.ProvinceName, provCol1) +
			formatStr(fmt.Sprintf("%d", p.CurrentConfirmedCount), provCol2) +
			formatStr(fmt.Sprintf("%d", p.ConfirmedCount), provCol2) +
			formatStr(fmt.Sprintf("%d", p.DeadCount), provCol2) +
			formatStr(fmt.Sprintf("%d", p.CuredCount), provCol2)
		str = append(str, s1)
		for _, c := range p.Citys {
			s2 := "|-" + formatStr(c.CityName, cityCol1) +
				formatStr(fmt.Sprintf("%d", c.CurrentConfirmedCount), cityCol2) +
				formatStr(fmt.Sprintf("%d", c.ConfirmedCount), cityCol2) +
				formatStr(fmt.Sprintf("%d", c.DeadCount), cityCol2) +
				formatStr(fmt.Sprintf("%d", c.CuredCount), cityCol2)
			str = append(str, s2)
		}
	}
	return strings.Join(str, "\n")
}

func main() {
	dataStr := scrapeHtml(url)
	provinces := jsonToData(&dataStr)
	formatData := provinces.getPrintData()
	fmt.Println(formatData)
	filename := fmt.Sprintf("新冠实时疫情(%s).txt", time.Now().Format(timeFormatFile))
	ioutil.WriteFile(filename, []byte(formatData), 0644)
}

func jsonToData(s *string) Provinces {
	var p Provinces
	err := json.Unmarshal([]byte(*s), &p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func scrapeHtml(url string) string {
	var ret string
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status coder error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("#getAreaStat").Each(func(i int, s *goquery.Selection) {
		rawStr := s.Text()
		if len(rawStr) > 0 {
			r, _ := regexp.Compile("\\[(.*)\\]")
			ret = r.FindString(rawStr)
			ret = strings.Replace(ret, `\r\n`, "", -1)
		}
	})
	return ret
}
