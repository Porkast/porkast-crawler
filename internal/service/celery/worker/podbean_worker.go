package worker

import (
	"context"
	"fmt"
	"porkast-crawler/internal/consts"
	"porkast-crawler/internal/service/celery/jobs"
	"porkast-crawler/internal/service/network"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func ParsePodbeanAllcategoryList(url string) {
	var (
		ctx          = gctx.New()
		htmlResp     string
		categoryList []string
	)

	g.Log().Line().Debug(ctx, "Start podbean all category list job")
	htmlResp = getPodbeanCategoryHtml(ctx, url)
	categoryList = getPodbeanCategoryUrlsFromHtml(ctx, htmlResp)
	for _, category := range categoryList {
		jobs.AssignPodbeanCategoryPopularListJob(ctx, category, 0)
	}
}

func ParsePodbeancategoryPopularShow(category string, page int) {
	var (
		ctx                 = gctx.New()
		categoryPopularHtml string
		rssLinkList         []string
		currentCategoryUrl  string
		nextPage            int
	)
	g.Log().Line().Debug(ctx, fmt.Sprintf("Start parse podbean category %s popular show, page %d", category, page))
	currentCategoryUrl = fmt.Sprintf(consts.PODBEAN_CATEGORY_POPULAR_API_URL, page, category)
	nextPage = page + 1
	categoryPopularHtml = getPodbeanCategoryPopularShowHtml(ctx, currentCategoryUrl)
	rssLinkList = getPodbeanCategoryPopularShowRSSList(ctx, categoryPopularHtml)
	if len(rssLinkList) != 0 {
		jobs.AssignPodbeanCategoryPopularListJob(ctx, category, nextPage)
		for _, rssLink := range rssLinkList {
			jobs.AssignPodbeanRSSJob(ctx, rssLink)
		}
	}
}

func ParsePodbeanShowRSS(rssLink string) {

	var (
		ctx     = gctx.New()
		rssResp string
	)

	g.Log().Line().Debug(ctx, "Start parse podbean show RSS")
	rssResp = network.TryGetRSSContent(ctx, rssLink)
	if rssResp == "" {
		g.Log().Line().Error(ctx, fmt.Sprintf("Get Podbean RSS feed content from %s is empty", rssLink))
		return
	}

	if isStringRSSXml(rssResp) {
		storeFeed(ctx, rssResp, rssLink, "ParsePodbeanShowRSS")
	}

}

func getPodbeanCategoryHtml(ctx context.Context, url string) (htmlResp string) {
	htmlResp = network.GetContent(ctx, url)
	return
}

func getPodbeanCategoryUrlsFromHtml(ctx context.Context, categoryHtmlStr string) (categories []string) {

	var (
		rootDocs soup.Root
	)
	rootDocs = soup.HTMLParse(categoryHtmlStr)
	categoryWrapTag := rootDocs.Find("div", "class", "rightside-panel")
	categoryWrapUl := categoryWrapTag.Find("ul")
	categoryATagList := categoryWrapUl.FindAll("a")
	for _, categoryATag := range categoryATagList {
		var category string
		category = categoryATag.Text()
		if gstr.Contains(category, "&") || gstr.Contains(category, " ") {
			categoryItemSplit := gstr.SplitAndTrim(category, " ")
			category = gstr.Join(categoryItemSplit, "+")
			category = gstr.Replace(category, "&", "%26")
		}
		categories = append(categories, category)
	}

	return
}

func getPodbeanCategoryPopularShowHtml(ctx context.Context, url string) (htmlStr string) {
	var (
		err         error
		respJsonStr string
		respJson    *gjson.Json
	)

	respJsonStr = network.GetContent(ctx, url)
	respJson, err = gjson.LoadJson(gconv.Bytes(respJsonStr))
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Parse the podbead category popular response failed with url : %s\nerror:%s\n", url, err))
	}

	htmlStr = respJson.Get("data.popuparData").String()

	return
}

func getPodbeanCategoryPopularShowRSSList(ctx context.Context, htmlStr string) (rssLinkList []string) {

	var (
		rootDocs soup.Root
		showUrl  string
	)
	htmlStr = gstr.Replace(htmlStr, "\\r", "")
	htmlStr = gstr.Replace(htmlStr, "\\n", "")
	htmlStr = gstr.Replace(htmlStr, "\\t", "")
	htmlStr = gstr.Replace(htmlStr, "\\\"", "\"")
	rootDocs = soup.HTMLParse(htmlStr)
	categoryATagList := rootDocs.FindAll("a", "class", "pro")
	for _, categoryATag := range categoryATagList {
		var rssLink string
		showUrl = categoryATag.Attrs()["href"]
		if gstr.Contains(showUrl, "http://podcast") && gstr.Contains(showUrl, "org/?source=pb") {
			rssLink = gstr.Replace(showUrl, "?source=pb", "feed.xml")
			rssLinkList = append(rssLinkList, rssLink)
		} else if gstr.Contains(showUrl, ".podbean.com/?source=pb") {
			showUrl = gstr.Replace(showUrl, "https://", "")
			showUrl = gstr.Replace(showUrl, ".podbean.com/?source=pb", "")
			rssLink = fmt.Sprintf(consts.PODBEAN_RSS_URL, showUrl)
			rssLinkList = append(rssLinkList, rssLink)
		}
	}
	return
}
