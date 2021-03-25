package stockx

type StockXProductResponse struct {
	Product ProductClass `json:"Product"`
}

type ProductClass struct {
	ID                   string               `json:"id"`
	UUID                 string               `json:"uuid"`
	Brand                string               `json:"brand"`
	Colorway             string               `json:"colorway"`
	Condition            Condition            `json:"condition"`
	CountryOfManufacture CountryOfManufacture `json:"countryOfManufacture"`
	Gender               string               `json:"gender"`
	ContentGroup         ContentGroup         `json:"contentGroup"`
	MinimumBid           int64                `json:"minimumBid"`
	Name                 string               `json:"name"`
	PrimaryCategory      string               `json:"primaryCategory"`
	SecondaryCategory    string               `json:"secondaryCategory"`
	UsHtsCode            string               `json:"usHtsCode"`
	UsHtsDescription     string               `json:"usHtsDescription"`
	ProductCategory      ContentGroup         `json:"productCategory"`
	ReleaseDate          string               `json:"releaseDate"`
	RetailPrice          int64                `json:"retailPrice"`
	Shoe                 string               `json:"shoe"`
	ShortDescription     string               `json:"shortDescription"`
	StyleID              string               `json:"styleId"`
	TickerSymbol         string               `json:"tickerSymbol"`
	Title                string               `json:"title"`
	DataType             string               `json:"dataType"`
	URLKey               string               `json:"urlKey"`
	SizeLocale           string               `json:"sizeLocale"`
	SizeTitle            string               `json:"sizeTitle"`
	SizeDescriptor       string               `json:"sizeDescriptor"`
	SizeAllDescriptor    string               `json:"sizeAllDescriptor"`
	Description          string               `json:"description"`
	LithiumIonBattery    bool                 `json:"lithiumIonBattery"`
	Type                 bool                 `json:"type"`
	ALim                 int64                `json:"aLim"`
	Year                 int64                `json:"year"`
	ShippingGroup        ContentGroup         `json:"shippingGroup"`
	Traits               []Trait              `json:"traits"`
	Meta                 Meta                 `json:"meta"`
	PortfolioItems       []interface{}        `json:"PortfolioItems"`
	Shipping             Shipping             `json:"shipping"`
	EnhancedImage        EnhancedImage        `json:"enhancedImage"`
	Media                Media                `json:"media"`
	CharityCondition     int64                `json:"charityCondition"`
	Breadcrumbs          []Breadcrumb         `json:"breadcrumbs"`
	Market               Market               `json:"market"`
	Children             map[string]Child     `json:"children"`
}

type Breadcrumb struct {
	Level   int64  `json:"level"`
	Name    string `json:"name"`
	URL     URL    `json:"url"`
	IsBrand bool   `json:"isBrand"`
}

type Child struct {
	ID                   string               `json:"id"`
	UUID                 string               `json:"uuid"`
	Brand                Brand                `json:"brand"`
	Colorway             string               `json:"colorway"`
	Condition            Condition            `json:"condition"`
	CountryOfManufacture CountryOfManufacture `json:"countryOfManufacture"`
	Gender               string               `json:"gender"`
	ContentGroup         ContentGroup         `json:"contentGroup"`
	MinimumBid           int64                `json:"minimumBid"`
	Name                 string               `json:"name"`
	PrimaryCategory      Brand                `json:"primaryCategory"`
	SecondaryCategory    string               `json:"secondaryCategory"`
	UsHtsCode            string               `json:"usHtsCode"`
	UsHtsDescription     string               `json:"usHtsDescription"`
	ProductCategory      ContentGroup         `json:"productCategory"`
	ReleaseDate          string               `json:"releaseDate"`
	RetailPrice          int64                `json:"retailPrice"`
	Shoe                 string               `json:"shoe"`
	ShortDescription     string               `json:"shortDescription"`
	StyleID              string               `json:"styleId"`
	TickerSymbol         string               `json:"tickerSymbol"`
	Title                string               `json:"title"`
	DataType             string               `json:"dataType"`
	URLKey               string               `json:"urlKey"`
	SizeLocale           string               `json:"sizeLocale"`
	SizeTitle            string               `json:"sizeTitle"`
	SizeDescriptor       string               `json:"sizeDescriptor"`
	SizeAllDescriptor    string               `json:"sizeAllDescriptor"`
	Description          string               `json:"description"`
	LithiumIonBattery    bool                 `json:"lithiumIonBattery"`
	Type                 bool                 `json:"type"`
	ALim                 int64                `json:"aLim"`
	Year                 int64                `json:"year"`
	ShippingGroup        ContentGroup         `json:"shippingGroup"`
	Traits               []Trait              `json:"traits"`
	Meta                 Meta                 `json:"meta"`
	PortfolioItems       []interface{}        `json:"PortfolioItems"`
	Shipping             Shipping             `json:"shipping"`
	EnhancedImage        EnhancedImage        `json:"enhancedImage"`
	Media                Media                `json:"media"`
	CharityCondition     int64                `json:"charityCondition"`
	Breadcrumbs          []Breadcrumb         `json:"breadcrumbs"`
	Market               Market               `json:"market"`
	ParentID             string               `json:"parentId"`
	ParentUUID           string               `json:"parentUuid"`
	ShoeSize             string               `json:"shoeSize"`
	SizeSortOrder        int64                `json:"sizeSortOrder"`
	SkuVariantGroup      interface{}          `json:"skuVariantGroup"`
}

type EnhancedImage struct {
	ProductUUID string `json:"productUuid"`
	ImageKey    string `json:"imageKey"`
	ImageCount  int64  `json:"imageCount"`
}

type Market struct {
	ProductID                 int64   `json:"productId"`
	SkuUUID                   *string `json:"skuUuid"`
	ProductUUID               string  `json:"productUuid"`
	LowestAsk                 int64   `json:"lowestAsk"`
	LowestAskSize             *string `json:"lowestAskSize"`
	ParentLowestAsk           int64   `json:"parentLowestAsk"`
	NumberOfAsks              int64   `json:"numberOfAsks"`
	HasAsks                   *int64  `json:"hasAsks"`
	SalesThisPeriod           int64   `json:"salesThisPeriod"`
	SalesLastPeriod           int64   `json:"salesLastPeriod"`
	HighestBid                int64   `json:"highestBid"`
	HighestBidSize            *string `json:"highestBidSize"`
	NumberOfBids              int64   `json:"numberOfBids"`
	HasBids                   *int64  `json:"hasBids"`
	AnnualHigh                int64   `json:"annualHigh"`
	AnnualLow                 int64   `json:"annualLow"`
	DeadstockRangeLow         int64   `json:"deadstockRangeLow"`
	DeadstockRangeHigh        int64   `json:"deadstockRangeHigh"`
	Volatility                float64 `json:"volatility"`
	DeadstockSold             int64   `json:"deadstockSold"`
	PricePremium              float64 `json:"pricePremium"`
	AverageDeadstockPrice     int64   `json:"averageDeadstockPrice"`
	LastSale                  int64   `json:"lastSale"`
	LastSaleSize              *string `json:"lastSaleSize"`
	SalesLast72Hours          int64   `json:"salesLast72Hours"`
	ChangeValue               int64   `json:"changeValue"`
	ChangePercentage          float64 `json:"changePercentage"`
	AbsChangePercentage       float64 `json:"absChangePercentage"`
	TotalDollars              int64   `json:"totalDollars"`
	UpdatedAt                 int64   `json:"updatedAt"`
	LastLowestAskTime         *int64  `json:"lastLowestAskTime"`
	LastHighestBidTime        *int64  `json:"lastHighestBidTime"`
	LastSaleDate              *string `json:"lastSaleDate"`
	CreatedAt                 string  `json:"createdAt"`
	DeadstockSoldRank         int64   `json:"deadstockSoldRank"`
	PricePremiumRank          int64   `json:"pricePremiumRank"`
	AverageDeadstockPriceRank int64   `json:"averageDeadstockPriceRank"`
	Featured                  *int64  `json:"featured"`
	LowestAskFloat            float64 `json:"lowestAskFloat"`
	HighestBidFloat           float64 `json:"highestBidFloat"`
}

type Media struct {
	The360        []string      `json:"360"`
	ImageURL      string        `json:"imageUrl"`
	SmallImageURL string        `json:"smallImageUrl"`
	ThumbURL      string        `json:"thumbUrl"`
	Has360        bool          `json:"has360"`
	Gallery       []interface{} `json:"gallery"`
}

type Meta struct {
	Charity            bool `json:"charity"`
	Raffle             bool `json:"raffle"`
	MobileOnly         bool `json:"mobile_only"`
	Restock            bool `json:"restock"`
	Deleted            bool `json:"deleted"`
	Hidden             bool `json:"hidden"`
	LockBuying         bool `json:"lock_buying"`
	LockSelling        bool `json:"lock_selling"`
	Redirected         bool `json:"redirected"`
	BrowsePageFeatured bool `json:"browse_page_featured"`
}

type Shipping struct {
	TotalDaysToShip         int64 `json:"totalDaysToShip"`
	HasAdditionalDaysToShip bool  `json:"hasAdditionalDaysToShip"`
	DeliveryDaysLowerBound  int64 `json:"deliveryDaysLowerBound"`
	DeliveryDaysUpperBound  int64 `json:"deliveryDaysUpperBound"`
}

type Trait struct {
	Name       TraitName   `json:"name"`
	Value      interface{} `json:"value"`
	Filterable bool        `json:"filterable"`
	Visible    bool        `json:"visible"`
	Highlight  bool        `json:"highlight"`
	Format     *Format     `json:"format,omitempty"`
}

type Brand string

const (
	Adidas Brand = "adidas"
)

type URL string

type Condition string

const (
	New Condition = "New"
)

type ContentGroup string

const (
	ContentGroupSneakers ContentGroup = "sneakers"
)

type CountryOfManufacture string

type Format string

const (
	Currency Format = "currency"
	Date     Format = "date"
)

type TraitName string

const (
	Colorway    TraitName = "Colorway"
	ReleaseDate TraitName = "Release Date"
	RetailPrice TraitName = "Retail Price"
	Style       TraitName = "Style"
)

type Value struct {
	Integer *int64
	String  *string
}
