package coltorti

func CrawlImage(productURL string) []string {
	return nil
}

func FilterEmptyImageUrl(urls []string) []string {
	ret := []string{}
	for _, url := range urls {
		if url == "" {
			continue
		}
		ret = append(ret, url)
	}
	return ret
}
