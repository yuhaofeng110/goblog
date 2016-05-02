package manage

type Leftbar struct {
	ID    string // 内部ID
	Title string // 说明
	Extra string // 链接
	Text  string // 显示名称
}

type View struct {
}

type BaseData struct {
	DayTrend     [2][145]int    // 日点击趋势
	KeywordTop10 map[string]int // 关键词top10
	Page         map[string]int // 页面排行
	Engine       map[string]int // 搜索引擎top5
	Source       map[string]int // 全部来源
	Links        map[string]int // 外部链接top5
	Area         map[string]int // 访客地域
}
