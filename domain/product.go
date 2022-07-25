package domain

type ProductOption struct {
	SizeInfo string
	SizeName string
	Quantity int
}

type Product struct {
	ProductURL    string
	Images        []string
	Brand         string
	ProductID     string
	Season        string
	Year          int
	Color         string
	MadeIn        string
	Material      string
	Name          string
	Description   string
	Category      string
	Quantity      int
	OriginalPrice float64
	CurrencyType  string
	DiscountRate  int
	SizeOptions   []ProductOption
	FTA           bool
}

func (pd *Product) ToProductTemplate() [][]string {
	records := [][]string{}
	row := []string{
		"신상품",
		"50000805",
		pd.Name,
	}
	records = append(records, row)
	return records
}

func AddProduct() {

}

/*
상품상태 0
카테고리ID 1 - 네이커에서 찾아서
상품명 2
판매가 3
재고수량 (옵션재고?) 4
A/S안내내용 5
A/S전화번호 6
대표이미지 파일명 7
추가 이미지 파일명 8
상품 상세정보 9 (이미지파일)
판매자 상품코드 10 비필수
반매자 바코드 11 비필수
제조사 12 비필수
브랜드 13 비필수
제조일자 14 비필수
유효일자 15 비필수
부가세 16 (과세상품, 면세상품, 영세상품))
미성년자 17 - Y/N
구매평 노출여부 18 - Y/N
원산지 코드 19 - 네이버에서 찾아서
수입사 20 - 필수
복수원산지 여부 21 - N
배송방법 23 - (택배, 소포, 등기, 직접배송)
배송비 유형 24 - (무료, 조건부 무료, 유료, 수량별부과)
기본배송비 25
배송비 결제방식 26 - (착불, 선결제, 착불 또는 선결제)
조건부무료-상품판매가합계 27
수량별부과-수량 28
반품배송비 29
교환배송비 30
지역별 차등배송비 31
판매자 특이사항 33
즉시할인 값 34
즉시할인 단위 35
복수구매할인 조건 값 36
복수구매 할인조건 단위 37
복수구매할인 단위 38
상품구매시 포인트 지급 값 39
상품구매시 포인트 지급 단위 40
텍스트리뷰 작성시 지급 포인트 41
포토 리뷰 작성시 지급 포인트 42
한달사용 텍스트리뷰 포인트 43
한달사용 포토리뷰 포인트 44
스토어찜고객 리뷰작성시 지급 포인트 45
무이자 할부개월 46
사은품 47
옵션형태 48 (단독형/조합형/입력형)
옵션명 49 (사이즈)
옵션값 50 (S,M,L)
옵션가 51
옵션 재고수량 52
추가상품명 53
추가상품 값 54
추가 상품가 55
추가상품 재고수량 56
상품정보제공고시 품명 57
상품정보제공고시 모델명 58
상품정보제공고시 인증허가사항 59
상품정보제공고시 제조자 60
스토어찜회원 전용여부 61
*/
