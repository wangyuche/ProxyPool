package proxynova

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type Proxynova struct {
}

func (this *Proxynova) Crawler() {
	c := colly.NewCollector()
	q, _ := queue.New(
		10,
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)
	c.OnHTML("#tbl_proxy_list > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_int int, e1 *colly.HTMLElement) {
			fmt.Print(e1.DOM.Find("td:nth-child(1)").Text())
		})
	})
	q.AddURL("https://www.proxynova.com/proxy-server-list")
	q.Run(c)
}

//#tbl_proxy_list > tbody > tr:nth-child(1) > td:nth-child(1)
//#tbl_proxy_list > tbody > tr:nth-child(2)
//#tbl_proxy_list > tbody > tr:nth-child(1) > td:nth-child(1) > script
