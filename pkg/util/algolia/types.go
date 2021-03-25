package algolia

type AlgolaQueryResponse struct {
	Hits                  []Hit       `json:"hits"`
	NbHits                int64       `json:"nbHits"`
	Page                  int64       `json:"page"`
	NbPages               int64       `json:"nbPages"`
	HitsPerPage           int64       `json:"hitsPerPage"`
	Facets                Facets      `json:"facets"`
	FacetsStats           FacetsStats `json:"facets_stats"`
	ExhaustiveFacetsCount bool        `json:"exhaustiveFacetsCount"`
	ExhaustiveNbHits      bool        `json:"exhaustiveNbHits"`
	Query                 string      `json:"query"`
	Params                string      `json:"params"`
	ProcessingTimeMS      int64       `json:"processingTimeMS"`
}

type Facets struct {
	Brand            map[string]int64     `json:"brand"`
	Gender           Gender               `json:"gender"`
	Categories       map[string]int64     `json:"categories"`
	LowestAsk        map[string]int64     `json:"lowest_ask"`
	QualityBid       map[string]int64     `json:"quality_bid"`
	LockSelling      LockSelling          `json:"lock_selling"`
	MediaImageURL    map[string]int64     `json:"media.imageUrl"`
	BuyingCountries  map[string]int64     `json:"buying_countries"`
	ProductCategory  ProductCategoryClass `json:"product_category"`
	SellingCountries map[string]int64     `json:"selling_countries"`
}

type Gender struct {
	Men       int64 `json:"men"`
	Unisex    int64 `json:"unisex"`
	Infant    int64 `json:"infant"`
	Preschool int64 `json:"preschool"`
	Women     int64 `json:"women"`
}

type LockSelling struct {
	False int64 `json:"false"`
}

type ProductCategoryClass struct {
	Sneakers     int64 `json:"sneakers"`
	Streetwear   int64 `json:"streetwear"`
	Collectibles int64 `json:"collectibles"`
	Handbags     int64 `json:"handbags"`
}

type FacetsStats struct {
	LowestAsk  LowestAsk `json:"lowest_ask"`
	QualityBid LowestAsk `json:"quality_bid"`
}

type LowestAsk struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
	Avg int64 `json:"avg"`
	Sum int64 `json:"sum"`
}

type Hit struct {
	ID               string              `json:"id,omitempty"`
	UUID             string              `json:"uuid,omitempty"`
	Name             string              `json:"name"`
	Brand            string              `json:"brand"`
	ThumbnailURL     string              `json:"thumbnail_url"`
	Media            Media               `json:"media"`
	URL              string              `json:"url"`
	ReleaseDate      string              `json:"release_date"`
	Categories       []string            `json:"categories"`
	ProductCategory  ProductCategoryEnum `json:"product_category"`
	TickerSymbol     string              `json:"ticker_symbol"`
	StyleID          string              `json:"style_id"`
	Make             string              `json:"make"`
	Model            string              `json:"model"`
	ShortDescription string              `json:"short_description"`
	Gender           string              `json:"gender"`
	Colorway         *string             `json:"colorway"`
	Price            int64               `json:"price"`
	Description      string              `json:"description"`
	HighestBid       int64               `json:"highest_bid"`
	TotalDollars     int64               `json:"total_dollars"`
	LowestAsk        int64               `json:"lowest_ask"`
	LastSale         int64               `json:"last_sale"`
	SalesLast72      int64               `json:"sales_last_72"`
	DeadstockSold    int64               `json:"deadstock_sold"`
	QualityBid       int64               `json:"quality_bid"`
	Active           int64               `json:"active"`
	NewRelease       int64               `json:"new_release"`
	Featured         int64               `json:"featured"`
	LockSelling      bool                `json:"lock_selling"`
	SellingCountries []string            `json:"selling_countries"`
	BuyingCountries  []string            `json:"buying_countries"`
	Traits           []Trait             `json:"Traits,omitempty"`
	SearchableTraits SearchableTraits    `json:"searchable_traits"`
	ObjectID         string              `json:"objectID"`
	HighlightResult  HighlightResult     `json:"_highlightResult"`
	HitTraits        []Trait             `json:"traits,omitempty"`
}

type HighlightResult struct {
	Name             Colorway            `json:"name"`
	URL              Colorway            `json:"url"`
	Categories       []Colorway          `json:"categories"`
	ProductCategory  Colorway            `json:"product_category"`
	TickerSymbol     Colorway            `json:"ticker_symbol"`
	StyleID          *Colorway           `json:"style_id,omitempty"`
	ShortDescription Colorway            `json:"short_description"`
	Gender           Colorway            `json:"gender"`
	Colorway         *Colorway           `json:"colorway,omitempty"`
	SearchableTraits map[string]Colorway `json:"searchable_traits"`
}

type Colorway struct {
	Value            string     `json:"value"`
	MatchLevel       MatchLevel `json:"matchLevel"`
	MatchedWords     []string   `json:"matchedWords"`
	FullyHighlighted *bool      `json:"fullyHighlighted,omitempty"`
}

type Trait struct {
	Name       string      `json:"name"`
	Value      interface{} `json:"value"`
	Filterable bool        `json:"filterable"`
	Visible    bool        `json:"visible"`
	Highlight  bool        `json:"highlight"`
	Format     *Format     `json:"format,omitempty"`
}

type Media struct {
	ImageURL      string        `json:"imageUrl"`
	SmallImageURL string        `json:"smallImageUrl"`
	ThumbURL      string        `json:"thumbUrl"`
	Gallery       []interface{} `json:"gallery"`
	Hidden        interface{}   `json:"hidden"`
}

type SearchableTraits struct {
	Colorway         *string `json:"Colorway,omitempty"`
	ReleaseDate      *string `json:"Release Date,omitempty"`
	RetailPrice      *int64  `json:"Retail Price,omitempty"`
	Style            *string `json:"Style,omitempty"`
	Season           *string `json:"Season,omitempty"`
	Color            *string `json:"Color,omitempty"`
	Retail           *int64  `json:"Retail,omitempty"`
	Size             *string `json:"size,omitempty"`
	Dimensions       *string `json:"Dimensions,omitempty"`
	Hardware         *string `json:"Hardware,omitempty"`
	Material         *string `json:"Material,omitempty"`
	BaseModel        *string `json:"base_model,omitempty"`
	BaseModelVariety *string `json:"base_model_variety,omitempty"`
	ColorGeneric     *string `json:"color_generic,omitempty"`
	Depth            *string `json:"depth,omitempty"`
	Height           *string `json:"height,omitempty"`
	InteriorColor    *string `json:"interior_color,omitempty"`
	Pattern          *string `json:"pattern,omitempty"`
	Serial           *string `json:"serial,omitempty"`
	Styles           *string `json:"styles,omitempty"`
	Width            *string `json:"width,omitempty"`
	Gender           *string `json:"gender,omitempty"`
}

type MatchLevel string

const (
	Full MatchLevel = "full"
	None MatchLevel = "none"
)

type Format string

const (
	Currency Format = "currency"
	Date     Format = "date"
)

type ProductCategoryEnum string

const (
	Collectibles ProductCategoryEnum = "collectibles"
	Handbags     ProductCategoryEnum = "handbags"
	Sneakers     ProductCategoryEnum = "sneakers"
	Streetwear   ProductCategoryEnum = "streetwear"
)

type Value struct {
	Integer *int64
	String  *string
}
