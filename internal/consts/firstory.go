package consts

const FIRSTORY_ENTRY_JOB = "FIRSTORY_ENTRY_JOB"

const FIRSTORY_CATEGORY_LIST_JOB = "FIRSTORY_CATEGORY_JOB"


const FIRSTORY_CATEGORY_SHOW_RSS_JOB = "FIRSTORY_CATEGORY_SHOW_RSS_JOB"

const FIRSTORY_CATEGORY_BASE_URL = "https://open.firstory.me/browse/category/"

const FIRSTORY_GRAPHQL_BASE_URL = "https://prod-api.firstory.me/hosting/v1/graphql"

const FIRSOTRY_RSS_BASE_URL = "https://open.firstory.me/user/"

const FIRSTORY_GRAPHQL_SHOW_QUERY_JSON = `{"operationName": "SearchShow","variables": {"categoryId": "%s","first": 20,"skip": %d},"query": "query SearchShow($showIds: [ID!], $categoryId: ID, $queryString: String, $first: Int, $skip: Int) {playerShowFind(showIds: $showIds, categoryId: $categoryId, queryString: $queryString, take: $first, skip: $skip) {id name avatar urlSlug __typename }}"}`

const FIRSTORY_CATEGORY_GRAPHQL_QUERY_JSON = `{"operationName":"GetCategories","variables":{},"query":"query GetCategories {\n  playerCategoryFind(subCategory: false) {\n    ...CategoryFragment\n    __typename\n  }\n}\n\nfragment CategoryFragment on Category {\n  id\n  nameEn\n  nameZh\n  parentCategory {\n    id\n    nameEn\n    nameZh\n    __typename\n  }\n  __typename\n}\n"}`

const FIRSOTRY_SHOW_INFO_QUERY_JSON = `{"operationName":"GetShowInfo","variables":{"showId":"%s"},"query":"query GetShowInfo($showId: String!) {\n  playerShowFindOneByUrlSlug(urlSlug: $showId) {\n    ...ShowInfoFragment\n    episodeCount\n    __typename\n  }\n}\n\nfragment ShowInfoFragment on Show {\n  id\n  name\n  avatar\n  intro\n  isCreator\n  author\n  urlSlug\n  language\n  explicit\n  categories {\n    id\n    nameEn\n    nameZh\n    __typename\n  }\n  import {\n    status\n    originRssUrl\n    __typename\n  }\n  distributions {\n    status\n    platformType\n    platformId\n    platformUrl\n    __typename\n  }\n  externalLinks {\n    title\n    type\n    url\n    __typename\n  }\n  websiteSetting {\n    ...WebsiteSettingFragment\n    __typename\n  }\n  __typename\n}\n\nfragment WebsiteSettingFragment on WebsiteSetting {\n  active\n  gaTrackingId\n  fbPixelId\n  themeHexFirst\n  themeHexSecond\n  themeHexThird\n  customDomain\n  flinkShowDonate\n  flinkShowExternal\n  flinkShowRssFeed\n  flinkShowPlatforms\n  flinkShowComment\n  flinkShowVoiceMail\n  flinkShowDownloadAudioFile\n  playerShowDonate\n  playerShowRssFeed\n  playerHideLogo\n  playerShowShownote\n  playerShowComment\n  playerShowVoicemail\n  playerShowDownloadAudioFile\n  __typename\n}\n"}`
