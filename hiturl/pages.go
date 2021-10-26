package hiturl

type PageLink struct {
	number int
	url    string
}

func NewPageLink(number int, url string) *PageLink {
	return &PageLink{
		number: number,
		url:    url,
	}
}
