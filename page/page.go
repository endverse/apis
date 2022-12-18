package page

type Page struct {
	TotalRecord int64 `json:"totalRecord"`
	FirstPageNo int   `json:"firstPageNo"`
	NextPageNo  int   `json:"nextPageNo"`
	LastPageNo  int   `json:"lastPageNo"`
	PageNo      int   `json:"pageNo"`
	Limit       int   `json:"limit"`
	Offset      int   `json:"offset"`
	PrePageNo   int   `json:"prePageNo"`
	PageNumList []int `json:"pageNumList"`
}

const (
	DefalutLimit      = 20
	DefaultPageNoSize = 3
)

func NewPage(totalRecord int64, pageNo, limit int) *Page {
	page := Page{}

	page.setTotalRecord(totalRecord)
	page.setLimit(&limit)
	page.setLastPageNo(int(totalRecord), limit)
	page.setPageNo(pageNo)
	page.setFirstPageNo()
	page.setNextPageNo()
	page.setOffset()
	page.setPrePageNo()
	page.setPageNumList()

	return &page
}

func (p *Page) setTotalRecord(totalRecord int64) {
	p.TotalRecord = totalRecord
}

func (p *Page) setLimit(limit *int) {
	if *limit > 0 {
		p.Limit = *limit
	} else {
		p.Limit = DefalutLimit
		*limit = DefalutLimit
	}
}

func (p *Page) setLastPageNo(totalRecord, limit int) {
	if (totalRecord % limit) == 0 {
		if totalRecord == 0 {
			p.LastPageNo = 1
		} else {
			p.LastPageNo = totalRecord / limit
		}
	} else {
		p.LastPageNo = (totalRecord / limit) + 1
	}
}

func (p *Page) setPageNo(pageNo int) {
	pageNum := pageNo
	if pageNo > p.LastPageNo {
		pageNum = p.LastPageNo
	}

	if pageNum <= 0 {
		pageNum = 1
	}

	p.PageNo = pageNum
}

func (p *Page) setFirstPageNo() {
	if p.PageNo > 1 {
		p.FirstPageNo = p.PageNo - 1
	} else {
		p.FirstPageNo = 1
	}
}

func (p *Page) setNextPageNo() {
	if p.LastPageNo > p.PageNo {
		p.NextPageNo = p.PageNo + 1
	} else {
		p.NextPageNo = p.PageNo
	}
}

func (p *Page) setOffset() {
	if p.TotalRecord == 0 {
		p.Offset = 0
	} else {
		p.Offset = (p.PageNo - 1) * p.Limit
	}
}

func (p *Page) setPrePageNo() {
	var prePage = p.PageNo - 1
	if prePage <= 0 {
		prePage = 1
	}

	p.PrePageNo = prePage
}

func (p *Page) setPageNumList() {
	if p.hasPrePage() && p.hasNextPage() && p.PageNo > 1 {
		p.PageNumList = p.assignPageNumList(p.PageNo - DefaultPageNoSize/2)
	} else if !p.hasNextPage() {
		p.PageNumList = p.assignPageNumList(p.PageNo - (DefaultPageNoSize - 1))
	} else if !p.hasPrePage() {
		p.PageNumList = p.assignPageNumList(p.PageNo)
	} else {
		p.PageNumList = []int{}
	}
}

func (p *Page) assignPageNumList(start int) []int {
	pageNumList := []int{}
	var startPageNo = start
	var loopCount = p.LastPageNo
	if loopCount > DefaultPageNoSize {
		loopCount = DefaultPageNoSize
	}

	if startPageNo <= 0 {
		startPageNo = 1
	}

	for i := 0; i < loopCount; i++ {
		pageNumList = append(pageNumList, startPageNo)
		startPageNo += 1
	}

	return pageNumList
}

func (p *Page) hasNextPage() bool {
	return p.PageNo != p.LastPageNo
}

func (p *Page) hasPrePage() bool {
	return p.PageNo != 1
}

func (p *Page) GetTotalRecord() int64 {
	return p.TotalRecord
}

func (p *Page) GetFirstPageNo() int {
	return p.FirstPageNo
}

func (p *Page) GetNextPageNo() int {
	return p.NextPageNo
}

func (p *Page) GetLastPageNo() int {
	return p.LastPageNo
}

func (p *Page) GetPageNo() int {
	return p.PageNo
}

func (p *Page) GetLimit() int {
	return p.Limit
}

func (p *Page) GetOffset() int {
	return p.Offset
}

func (p *Page) GetPrePageNo() int {
	return p.PrePageNo
}

func (p *Page) GetPageNumList() []int {
	return p.PageNumList
}
