package gotipboard

const (
	TileText             = "text"
	TilePieChart         = "pie_chart"
	TileLineChart        = "line_chart"
	TileCumulativeFlow   = "cumulative_flow"
	TileSimplePercentage = "simple_percentage"
	TileListing          = "listing"
	TileBarChart         = "bar_chart"
	TileFancyListing     = "fancy_listing"
	TileBigValue         = "big_value"
	TileJustValue        = "just_value"
	TileAdvancedPlot     = "advanced_plot"
	TileNormChart        = "norm_chart"
)

type dataTile struct {
	Title       string `json:"title,omitempty"`
	Subtitle    string `json:"subtitle,omitempty"`
	Description string `json:"description,omitempty"`
}

type DataText struct {
	Text string `json:"text,omitempty"`
}

type DataPieChart struct {
	dataTile
	PieData map[string]int `json:"pie_data"`
}

type DataLineChart struct {
	dataTile
	SeriesList []map[string]int `json:"series_list,omitempty"`
	/*
		[
			[["23.09", 8326], ["24.09", 260630], ["25.09", 240933], ["26.09", 229639], ["27.09", 190240], ["28.09", 125272], ["29.09", 3685]],
			[["23.09", 3685], ["24.09", 125272], ["25.09", 190240], ["26.09", 229639], ["27.09", 240933], ["28.09", 260630], ["29.09", 108326]]
		]
	*/
}

type itemCumulativeFlow struct {
	Label  string `json:"label,omitempty"`
	Series []int  `json:"series,omitempty"`
}

type DataCumulativeFlow struct {
	dataTile
	SeriesList []itemCumulativeFlow `json:"series_list,omitempty"`
}

type DataSimplePercentage struct {
	dataTile
	BigValue   string `json:"big_value,omitempty"`
	LeftLabel  string `json:"left_label,omitempty"`
	LeftValue  string `json:"left_value,omitempty"`
	RightLabel string `json:"right_label,omitempty"`
	RightValue string `json:"right_value,omitempty"`
}

type DataListing struct {
	Items []string `json:"items,omitempty"`
}

type DataBarChart struct {
	dataTile
	Ticks      []string `json:"ticks,omitempty"`
	SeriesList [][]int  `json:"series_list,omitempty"`
}

type itemFancyListing struct {
	Label       string `json:"label,omitempty"`
	Text        string `json:"text,omitempty"`
	Description string `json:"description,omitempty"`
}

type DataFancyListing struct {
	Data []itemFancyListing `json:"data,omitempty"`
}

type DataBigValue struct {
	dataTile
	BigValue        string `json:"big-value,omitempty"`
	UpperLeftLabel  string `json:"upper-left-label,omitempty"`
	UpperLeftValue  string `json:"upper-left-value,omitempty"`
	UpperRightLabel string `json:"upper-right-label,omitempty"`
	UpperRightValue string `json:"upper-right-value,omitempty"`
	LowerLeftLabel  string `json:"lower-left-label,omitempty"`
	LowerLeftValue  string `json:"lower-left-value,omitempty"`
	LowerRightLabel string `json:"lower-right-label,omitempty"`
	LowerRightValue string `json:"lower-right-value,omitempty"`
}

type DataJustValue struct {
	dataTile
	JustValue string `json:"just-value,omitempty"`
}

type DataAdvancedPlot struct {
	dataTile
	PlotData interface{} `json:"plotData,omitempty"`
}

type DataNormChart struct {
	dataTile
	PlotData []map[int]float64 `json:"plot_data,omitempty"`
}


// https://github.com/allegro/tipboard/issues/48

func (tile *dataTile) SetTitle(title string) {
	tile.Title = title
}

func (tile *dataTile) SetDescription(description string) {
	tile.Description = description
}

func (tile *DataBarChart) SetDescription(description string) {
	tile.Subtitle = description
}

func (tile *DataSimplePercentage) SetDescription(description string) {
	tile.Subtitle = description
}

func (tile *DataLineChart) SetTitle(title string) {
	tile.Subtitle = title
}

// get tile name

func (tile *DataText) Tile() string {
	return TileText
}

func (tile *DataPieChart) Tile() string {
	return TilePieChart
}

func (tile *DataLineChart) Tile() string {
	return TileLineChart
}

func (tile *DataCumulativeFlow) Tile() string {
	return TileCumulativeFlow
}

func (tile *DataSimplePercentage) Tile() string {
	return TileSimplePercentage
}

func (tile *DataListing) Tile() string {
	return TileListing
}

func (tile *DataBarChart) Tile() string {
	return TileBarChart
}

func (tile *DataFancyListing) Tile() string {
	return TileFancyListing
}

func (tile *DataBigValue) Tile() string {
	return TileBigValue
}

func (tile *DataJustValue) Tile() string {
	return TileJustValue
}

func (tile *DataAdvancedPlot) Tile() string {
	return TileAdvancedPlot
}

func (tile *DataNormChart) Tile() string {
	return TileNormChart
}
