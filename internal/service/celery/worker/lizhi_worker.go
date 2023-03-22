package worker

import (
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery/jobs"
	"guoshao-fm-crawler/internal/service/network"
	"strconv"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
)

func ParseLizhiAllCategories(url string) {
	var (
		ctx             = gctx.New()
		homePageRespStr string
		categoryUrlList []string
	)

	homePageRespStr = network.GetContent(ctx, consts.LIZHI_FM_ENTRY_URL)
	if homePageRespStr == "" {
		g.Log().Error(ctx, "Get Lizhi FM home page failed")
	}
	categoryUrlList = getLizhiCategory(homePageRespStr)
	for _, url := range categoryUrlList {
		jobs.AssignLizhiCateGoryParseJob(ctx, url)
	}
}

func ParseLizhiPodcastByCategoryPage(url string) {

	var (
		ctx                 = gctx.New()
		categoryPageRespStr string
		podcastIdList       []string
		totalRadioCount     int
		pageNum             int
		nextCategoryUrl     string
	)

	categoryPageRespStr = network.GetContent(ctx, consts.LIZHI_FM_ENTRY_URL)
	if categoryPageRespStr == "" {
		g.Log().Error(ctx, "Get Lizhi FM category page failed")
	}
	podcastIdList = getLizhiPodcastIdList(categoryPageRespStr)
	for _, podcastId := range podcastIdList {
		var (
			rssUrl string
		)
		rssUrl = genLizhiRSSUrl(podcastId)
		jobs.AssignLizhiPodcastXmlJob(ctx, rssUrl)
	}
	totalRadioCount = getLizhiCurrentCategoryPageRadioCount(url)
	if totalRadioCount < 24 {
		return
	}
	pageNum = getLizhiCategoryCurrentPageNum(url)
	nextCategoryUrl = genNextLizhiCategoryPageUrl(url, pageNum)

	g.Log().Info(ctx, podcastIdList, pageNum, nextCategoryUrl)
}

func ParseLizhiPodcastRSS(url string) {
	var (
		ctx        = gctx.New()
		podcastUrl string
		respStr    string
	)

	respStr = network.GetContent(ctx, podcastUrl)
	if isStringXml(respStr) {
		//The lizhi fm radio is RSS
		storeFeed(ctx, respStr)
	}
}

func getLizhiCategory(homePageRespStr string) (categoryList []string) {
	var (
		rootDocs soup.Root
	)

	rootDocs = soup.HTMLParse(homePageRespStr)
	allRadioTag := rootDocs.Find("div", "id", "allRadioTag")
	categoryTags := allRadioTag.FindAll("a")
	for _, categoryTag := range categoryTags {
		var (
			categoryLink string
		)
		categoryLink = categoryTag.Attrs()["href"]
		categoryLink = gstr.Replace(categoryLink, "//www", "www")
		categoryLink = categoryLink + "1.html"
		categoryList = append(categoryList, categoryLink)
	}
	return
}

func genNextLizhiCategoryPageUrl(url string, currentPage int) (nextUrl string) {
	var (
		strArray       []string
		nextPageNum    int
		nextPageNumStr string
	)

	strArray = gstr.Split(url, "/")
	nextUrl = gstr.Join(strArray[:len(strArray)-1], "/")
	nextPageNum = currentPage + 1
	nextPageNumStr = strconv.Itoa(nextPageNum)
	nextUrl = nextUrl + "/" + nextPageNumStr + ".html"

	return
}

func genLizhiRSSUrl(podcastId string) (url string) {

	url = fmt.Sprintf(consts.LIZHI_FM_RSS_URL, podcastId)

	return
}

func getLizhiCategoryCurrentPageNum(url string) (num int) {

	var (
		strArray    []string
		endStrArray []string
		endStr      string
	)

	strArray = gstr.Split(url, "/")
	endStr = strArray[len(strArray)-1]
	endStrArray = gstr.Split(endStr, ".")
	if len(endStrArray) != 0 {
		num, _ = strconv.Atoi(endStrArray[0])
	}

	return
}

func getLizhiPodcastIdList(htmlStr string) (idList []string) {
	var (
		rootDocs soup.Root
	)

	rootDocs = soup.HTMLParse(htmlStr)
	radioTagList := rootDocs.FindAll("p", "class", "radio-last-audio")
	for _, radioTag := range radioTagList {
		linkTag := radioTag.Find("a")
		podcastLink := linkTag.Attrs()["href"]
		podcastLink = gstr.Replace(podcastLink, "//www", "www")
		strArray := gstr.Split(podcastLink, "/")
		if len(strArray) == 0 {
			continue
		}
		podcastId := strArray[1]
		idList = append(idList, podcastId)
	}

	return
}

func getLizhiCurrentCategoryPageRadioCount(htmlStr string) (count int) {

	var (
		rootDocs soup.Root
	)

	rootDocs = soup.HTMLParse(htmlStr)
	radioTagList := rootDocs.FindAll("li", "class", "radio_list")
	count = len(radioTagList)
	return
}