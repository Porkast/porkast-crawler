package worker

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery/jobs"
	"guoshao-fm-crawler/internal/service/network"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func ParseFistoryAllCategoryList() {
	var (
		ctx             = gctx.New()
		categoryStrResp string
		categoryIdList  []string
	)

	categoryIdList = parseFirstoryAllCategoryList(categoryStrResp)
	for _, categoryId := range categoryIdList {
		jobs.AssignFirstoryCategoryJob(ctx, categoryId, 0)
	}
}

func ParseFirstoryCategoryItemList(categoryId string, skip int) {
	var (
		ctx        = gctx.New()
		respStr    string
		showIdList []string
	)

	respStr = getFirstoryCategoryShowsJsonStr(ctx, categoryId, skip)
	showIdList = parseFirstoryShowList(respStr)
	for _, showId := range showIdList {
		jobs.AssignFirstoryShowRSSJob(ctx, showId)
	}

}

func ParseFirstoryShowRSS(categoryId string) {
	var (
		ctx     = gctx.New()
		respStr string
		rssLink string
	)

	respStr = getFirstoryCategoryShowInfoJsonStr(ctx, categoryId)
	rssLink = getFirsotryShowRSSLink(ctx, respStr)
	respStr = network.TryGetRSSContent(ctx, rssLink)
	if isStringRSSXml(respStr) {
		storeFeed(ctx, respStr)
	}

}

func getFirstoryAllCategoryList(ctx context.Context) (resp string) {
	resp = network.PostContentByMobile(ctx, consts.FIRSTORY_GRAPHQL_BASE_URL, consts.FIRSTORY_CATEGORY_GRAPHQL_QUERY_JSON)
	return
}

func parseFirstoryAllCategoryList(jsonStr string) (categoryIdList []string) {
	var (
		categoryJsonResp *gjson.Json
		categoryJsonList []*gjson.Json
	)

	categoryJsonResp = gjson.New(jsonStr)
	categoryJsonList = categoryJsonResp.GetJsons("data.playerCategoryFind")
	for _, categoryJson := range categoryJsonList {
		var categoryId string
		categoryId = categoryJson.Get("id").String()
		categoryIdList = append(categoryIdList, categoryId)
	}
	return
}

func getFirstoryCategoryShowsJsonStr(ctx context.Context, categoryId string, skip int) (resp string) {
	var (
		queryStr string
	)
	queryStr = fmt.Sprintf(consts.FIRSTORY_GRAPHQL_SHOW_QUERY_JSON, categoryId, skip)
	resp = network.PostContentByMobile(ctx, consts.FIRSTORY_GRAPHQL_BASE_URL, queryStr)
	return
}

func parseFirstoryShowList(jsonStr string) (showIdList []string) {
	var (
		showRespJson *gjson.Json
		showJsonList []*gjson.Json
	)
	showRespJson = gjson.New(jsonStr)
	showJsonList = showRespJson.GetJsons("data.playerShowFind")
	for _, showJson := range showJsonList {
		var (
			showIdStr string
		)
		showIdStr = showJson.Get("id").String()
		showIdList = append(showIdList, showIdStr)
	}

	return
}

func getFirstoryCategoryShowInfoJsonStr(ctx context.Context, categoryId string) (resp string) {
	var (
		queryStr string
	)
	queryStr = fmt.Sprintf(consts.FIRSOTRY_SHOW_INFO_QUERY_JSON, categoryId)
	resp = network.PostContentByMobile(ctx, consts.FIRSTORY_GRAPHQL_BASE_URL, queryStr)
	return
}

func getFirsotryShowRSSLink(ctx context.Context, jsonStr string) (rssLink string) {
	var (
		respJson  *gjson.Json
		showIdStr string
	)
	respJson = gjson.New(jsonStr)
	if respJson == nil || respJson.IsNil() {
		g.Log().Line().Error(ctx, "Parse firstory show response json failed")
		return
	}

	rssLink = respJson.Get("data.playerShowFindOneByUrlSlug.import.originRssUrl").String()
	if rssLink == "" {
		showIdStr = respJson.Get("data.playerShowFindOneByUrlSlug.id").String()
		rssLink = consts.FIRSOTRY_RSS_BASE_URL + showIdStr
	}

	return
}
