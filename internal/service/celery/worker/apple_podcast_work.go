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

func ParseApplePodcastAllCategories(url string) {
	var (
		ctx             = gctx.New()
		categoryUrlList []string
		respStr         string
	)

	g.Log().Line().Debug(ctx, "Start parse apple podcast category page")
	respStr = network.GetContent(ctx, url)
	if respStr == "" {
		g.Log().Line().Error(ctx, "Get apple podcast category page failed")
	}
	categoryUrlList = parseAllApplePodcastCategoryUrlList(ctx, respStr)
	for _, categoryLink := range categoryUrlList {
		jobs.AssignApplePodcastCategoryItemJob(ctx, categoryLink)
	}
}

func ParseApplePodcastCategoryItems(categoryUrl string) {
	var (
		ctx         = gctx.New()
		itemUrlList []string
		respStr     string
	)

	g.Log().Line().Debug(ctx, "Start parse apple podcast category item page")
	respStr = network.GetContent(ctx, categoryUrl)
	if respStr == "" {
		g.Log().Line().Error(ctx, "Get apple podcast category item failed")
	}

	itemUrlList = parseApplePodcastCategoryItemUrls(ctx, respStr)
	for _, itemUrl := range itemUrlList {
		jobs.AssignApplePodcastItemRSSJob(ctx, itemUrl)
	}
}

func GetApplePodcastItemRSS(itemUrl string) {
	var (
		ctx        = gctx.New()
		itemId     string
		rssContent string
	)

	itemId = parseApplePodcastItemId(ctx, itemUrl)
	rssContent = getApplePodcastItemRSSByLookupAPI(ctx, itemId)
	if isStringRSSXml(rssContent) {
		storeFeed(ctx, rssContent)
	}
}

func parseAllApplePodcastCategoryUrlList(ctx context.Context, htmlStr string) (categoryUrlList []string) {

	defer func(ctx context.Context) {
		if rec := recover(); rec != nil {
			g.Log().Line().Error(ctx, fmt.Sprintf("parse apple podcast all categories failed: %s", rec))
		}
	}(ctx)

	docs := soup.HTMLParse(htmlStr)
	navDoc := docs.Find("div", "class", "grid3-column")
	categoryUrlDocList := navDoc.FindAll("a")
	for _, categoryUrlDoc := range categoryUrlDocList {
		var categoryUrl string
		categoryUrl = categoryUrlDoc.Attrs()["href"]
		categoryUrlList = append(categoryUrlList, categoryUrl)
	}

	return
}

func parseApplePodcastCategoryItemUrls(ctx context.Context, htmlStr string) (itemUrlList []string) {
	defer func(ctx context.Context) {
		if rec := recover(); rec != nil {
			g.Log().Line().Error(ctx, fmt.Sprintf("parse apple podcast category item url list failed: %s", rec))
		}
	}(ctx)

	docs := soup.HTMLParse(htmlStr)
	itemsWrapDocs := docs.Find("div", "class", "grid3-column")
	itemATags := itemsWrapDocs.FindAll("a")
	for _, itemAtag := range itemATags {
		var itemUrl string
		itemUrl = itemAtag.Attrs()["href"]
		itemUrlList = append(itemUrlList, itemUrl)
	}

	return
}

func parseApplePodcastItemId(ctx context.Context, itemUrl string) (itemId string) {
	var (
		split []string
	)

	split = gstr.Split(itemUrl, "/id")
	if len(split) < 2 {
		g.Log().Line().Error(ctx, fmt.Sprintf("parse apple podcast item id by url (%s) failed", itemUrl))
	}
	itemId = split[1]
	return
}

func getApplePodcastItemRSSByLookupAPI(ctx context.Context, itemId string) (rss string) {
	var (
		apiUrl      string
		respJsonStr string
		rssLink     string
		respJson    *gjson.Json
	)

	apiUrl = fmt.Sprintf(consts.APPLE_PODCAST_ITUNE_LOOKUP_API, itemId)
	respJsonStr = network.GetContent(ctx, apiUrl)
	if respJsonStr == "" {
		g.Log().Line().Error(ctx, "get content from apple podcast itune lookup api failed with item id ", itemId)
	}
	respJson = gjson.New(respJsonStr)
	rssLink = respJson.Get("results.0.feedUrl").String()
	rss = network.GetContent(ctx, rssLink)
	if rss == "" {
		g.Log().Line().Error(ctx, fmt.Sprintf("get rss content by apple podcast itune api with item id (%s) link (%s) failed", itemId, rssLink))
	}

	return
}
