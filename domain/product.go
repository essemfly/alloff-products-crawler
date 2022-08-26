package domain

type Source struct {
	Code          string
	ExcelFilename string
	TextFilename  string
	Tabs          []string
	Brands        []string
}

type ProductOption struct {
	SizeInfo string
	SizeName string
	Quantity int
}

type WriteProductRow interface {
	ToProductTemplate() []string
}
type Product struct {
	Source            Source
	ProductURL        string
	Images            []string
	ImageFilenames    []string
	Brand             string
	ProductID         string
	ProductStyleisNow string
	Season            string
	Year              int
	Color             string
	MadeIn            string
	Material          string
	Name              string
	Description       string
	Category          string
	Quantity          int
	OriginalPrice     float64
	CurrencyType      string
	DiscountRate      int
	SizeOptions       []ProductOption
	FTA               bool
}
