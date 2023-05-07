package worker

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery/jobs"
	"guoshao-fm-crawler/internal/service/network"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
)

func ParseSpreakerAllCategoryList() {
	var (
		ctx             = gctx.New()
		categoryStrResp string
		categoryIdList  []string
	)
	g.Log().Line().Debug(ctx, "Start parse spreaker all category")
	categoryStrResp = getSpreakerCategory(ctx)
	if categoryStrResp == "" {
		g.Log().Line().Error(ctx, "Get Spreaker category html failed")
	}
	categoryIdList = parseSpreakerCategoryItemList(ctx, categoryStrResp)
	for _, categoryId := range categoryIdList {
		var url string
		url = consts.SPREAKER_BASE_URL + categoryId
		jobs.AssignSpreakerSingleCategoryJob(ctx, url)
	}
}

func ParseSpreakerSingleCategory(url string) {
	var (
		ctx      = gctx.New()
		respStr  string
		nextUrl  string
		showUrls []string
	)

	g.Log().Line().Debug(ctx, "Start parse spreaker single category")
	respStr = network.GetContent(ctx, url)
	nextUrl = getSpreakerCategoryNextUrl(respStr)
	showUrls = getSpreakerShowLinks(respStr)
	if nextUrl != "" {
		jobs.AssignSpreakerSingleCategoryJob(ctx, nextUrl)
	} else {
		g.Log().Line().Error(ctx, fmt.Sprintf("The spreaker next url is empty by url %s", url))
	}

	if len(showUrls) > 0 {
		for _, showUrl := range showUrls {
			jobs.AssignSpreakerShowRSSJob(ctx, showUrl)
		}
	} else {
		g.Log().Line().Error(ctx, fmt.Sprintf("The spreaker show url list is empty from %s", url))
	}

}

func ParseSpreakerShowRSS(url string) {
	var (
		ctx      = gctx.New()
		showResp string
		rssLink  string
		rssResp  string
	)

	g.Log().Line().Debug(ctx, "Start parse spreaker show RSS")
	showResp = network.GetContent(ctx, url)
	if showResp == "" {
		g.Log().Line().Error(ctx, fmt.Sprintf("Get Spreaker show response from %s is empty", url))
		return
	}
	rssLink = getSpreakerShowRSSLink(showResp)
	if rssLink == "" {
		g.Log().Line().Error(ctx, fmt.Sprintf("Get Spreaker show RSS link from %s is empty", url))
		return
	}

	rssResp = network.TryGetRSSContent(ctx, rssLink)
	if rssResp == "" {
		g.Log().Line().Error(ctx, fmt.Sprintf("Get Spreaker RSS feed content from %s is empty", rssLink))
		return
	}

	if isStringRSSXml(rssResp) {
		storeFeed(ctx, rssResp, rssLink)
	}

}

func getSpreakerCategory(ctx context.Context) (resp string) {

	resp = network.GetContent(ctx, consts.SPREAKER_CATEGORY_URL)
	return
}

func parseSpreakerCategoryItemList(ctx context.Context, categoryHtmlStr string) (categoryIdList []string) {

	var (
		rootDocs soup.Root
	)

	rootDocs = soup.HTMLParse(categoryHtmlStr)
	allCategoryATagList := rootDocs.FindAllStrict("a", "class", "button button--primary button--small")
	for _, categoryATagItem := range allCategoryATagList {
		var (
			categoryStr string
		)
		categoryStr = categoryATagItem.Attrs()["href"]
		categoryStr = gstr.Replace(categoryStr, "/", "")
		categoryIdList = append(categoryIdList, categoryStr)
	}

	return
}

func getSpreakerCategoryNextUrl(respStr string) (nextUrl string) {
	var (
		err      error
		respJson *gjson.Json
		rootDocs soup.Root
	)
	respJson, err = gjson.LoadJson(respStr)
	if err == nil {
		nextUrl = respJson.GetJson("response.next_url").Var().String()
	} else {
		rootDocs = soup.HTMLParse(respStr)
		lastIdDoc := rootDocs.Find("div", "id", "expl_pager")
		if lastIdDoc.Error == nil {
			nextUrl = lastIdDoc.Attrs()["data-pager-api-next-url"]
		}
	}

	return
}

func getSpreakerShowLinks(respStr string) (linkList []string) {

	var (
		err      error
		respJson *gjson.Json
		rootDocs soup.Root
	)

	respJson, err = gjson.LoadJson(respStr)
	if err == nil {
		itemList := respJson.GetJsons("response.items")
		for _, item := range itemList {
			linkId := item.Get("permalink").String()
			link := consts.SPREAKER_BASE_URL + "show/" + linkId
			linkList = append(linkList, link)
		}
	} else {
		rootDocs = soup.HTMLParse(respStr)
		linkDocs := rootDocs.FindAllStrict("div", "class", "tile")
		for _, linkDoc := range linkDocs {
			linkATag := linkDoc.Find("a")
			link := linkATag.Attrs()["href"]
			linkList = append(linkList, link)
		}
	}
	return
}

func getSpreakerShowRSSLink(resp string) (rssLink string) {
	var (
		rootDocs soup.Root
	)

	rootDocs = soup.HTMLParse(resp)
	rssTag := rootDocs.Find("a", "id", "show_episodes_feed")
	if rssTag.Error == nil {
		rssLink = rssTag.Attrs()["href"]
	}

	return
}
