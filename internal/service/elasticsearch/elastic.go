package elasticsearch

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"
)

type ESClient struct {
	Client *elastic.Client
}

var esClient *ESClient

func InitES(ctx context.Context) {
	var (
		host     *gvar.Var
		port     *gvar.Var
		username *gvar.Var
		password *gvar.Var
		url      string
	)
	host, _ = g.Cfg().Get(ctx, "elastic.host")
	port, _ = g.Cfg().Get(ctx, "elastic.port")
	username, _ = g.Cfg().Get(ctx, "elastic.username")
	password, _ = g.Cfg().Get(ctx, "elastic.password")
	url = host.String() + ":" + port.String()
	var err error
	esClient = &ESClient{}
	esClient.Client, err = elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false), elastic.SetHealthcheck(false), elastic.SetBasicAuth(username.String(), password.String()))
	if err != nil {
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := esClient.Client.Ping(url).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	g.Log().Line().Infof(ctx, "Elasticsearch returned with code %d and version %s", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esVersion, err := esClient.Client.ElasticsearchVersion(url)
	if err != nil {
		// Handle error
		panic(err)
	}
	g.Log().Line().Infof(ctx, "Elasticsearch version %s", esVersion)
}

func Client() *ESClient {
	return esClient
}
