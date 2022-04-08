package web

import (
	"github.com/see792/gotool/util/struct2"
	"strings"
)

type Param map[string]interface{}

type Header map[string]interface{}

//http://
func GetUrlFileName(url string) string {
	indexName := "index.html"

	spiltIndex := strings.LastIndex(url, "/")

	if spiltIndex > 7 {

		endName := url[spiltIndex:]

		if endName == "/" {
			return indexName
		}

		return endName[1:]

	} else {
		return indexName
	}

}

func GetAbsUrl(url string, host string) string {

	if strings.Index(url, "http://") == 0 || strings.Index(url, "https://") == 0 {

		return url

	}

	if strings.LastIndex(host, "/") == len(host)-1 {
		return host + "" + url
	}
	return host + "/" + url

}
func GetUrlPath(url string) string {

	urlHost := GetUrlHostHttp(url)

	return url[len(urlHost)+1:]

}
func GetUrlHost(url string) string {

	headLen := 7
	spiltIndex := strings.Index(url, "http://")

	if spiltIndex < 0 {
		spiltIndex = strings.Index(url, "https://")
		headLen = 8
		if spiltIndex < 0 {
			return url
		}
	}

	endIndex := strings.Index(url[headLen:], "/")

	if endIndex < 0 {

		return url[headLen:]

	}
	return url[headLen : headLen+endIndex]

}
func GetUrlHostHttp(url string) string {

	headLen := 7
	spiltIndex := strings.Index(url, "http://")

	if spiltIndex < 0 {
		spiltIndex = strings.Index(url, "https://")
		headLen = 8
		if spiltIndex < 0 {

			return "http://" + url
		}
	}

	endIndex := strings.Index(url[headLen:], "/")

	if endIndex < 0 {

		return url

	}
	return url[:headLen+endIndex]

}

func ParseParam(param map[string]interface{}) string {
	if len(param)==0 {
		return ""
	}
	var connectParam = "&"
	var connectStr = ""
	for k:=range param {
		v:= struct2.InterfaceStrval(param[k])
		connectStr +=k+"="+v+connectParam
	}
	if len(connectStr)>0 {
		connectStr = connectStr[:len(connectStr)-1]
	}

	return connectStr


}
