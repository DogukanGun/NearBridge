package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"time"
)

/*
PROCESS LIST
1. Github repo https://github.com/{protocol_name}
-> is request 200
=> score += 0.5 jump 2
-> is not request 200
=> jump 3
2. Github people https://github.com/orgs/{protocol_name}/people
-> is request 200
=> score += 0.5
==> next
3. Get Github repos find first tr that contains protocol name https://github.com/orgs/{protocol_name}/repositories
-> a project is found go https://github.com/{protocol_name}/{project_name}
--> Get last commit data and calculate last days
*/
func main() {
	scrapeDataFromETHScan()
}

func CalculateTransparencyScore(protocolName string) (score float64) {
	tempScore := 0.0
	html := scrapeDataFromGithub("https://github.com/"+protocolName, "h1")
	if html != nil {
		tempScore += 0.5
	}
	html = scrapeDataFromGithub("https://github.com/orgs/"+protocolName+"/people", "img.avatar")
	if html != nil && html.Attr("src") != "" {
		tempScore += 0.5
	}
	return tempScore
}

func scrapeDataFromETHScan() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"
	c.Cookies("ASP.NET_SessionId=sxfoww5f45de45j4hfe0wt2i; __cflb=0H28vPcoRrcznZcNZSuFrvaNdHwh857ud7yZhKihFzL; _gid=GA1.2.464227293.1696408011; etherscan_offset_datetime=+3; __stripe_mid=141eaa72-9c5a-493b-8795-89d8d0425910b75551; __cuid=d0de949cde024858bc96967bd4cd2ce0; amp_fef1e8=ce49c1c4-987f-4331-9986-32ee55082afdR...1hbton0n2.1hbtoui8l.15.2.17; etherscan_cookieconsent=True; cf_chl_2=30040249d3d8f14; cf_clearance=FrH1L4sQtZedgfNVnNS61ftDa92HISZ3yYmj56tGVNU-1696442579-0-1-4325b6b2.3e22381d.dd534ed7-150.2.1696441135; __stripe_sid=5c44808f-5362-46b0-bb9d-1505c834b366c990b4; _ga=GA1.2.186259805.1694865783; _ga_T1JC9RNQXV=GS1.1.1696441134.8.1.1696442613.4.0.0")
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("a.me-1", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		childNumber := 0
		e.ForEach("td", func(_ int, child *colly.HTMLElement) {
			// Increment the child number
			childNumber++

			// Check if it's the third child
			if childNumber == 3 {
				// This is the third child, you can access it here
				fmt.Println(child.Text)
			}
		})

	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})
	err := c.Limit(&colly.LimitRule{
		DomainGlob: "*etherscan.io*",
		Delay:      2 * time.Second, // Add a 2-second delay between requests.
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.Visit("https://etherscan.io/accounts/label/BINANCE")
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Wait()
}

func scrapeDataFromGithub(link string, htmlQuerySelector string) (res *colly.HTMLElement) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML(htmlQuerySelector, func(e *colly.HTMLElement) {
		res = e
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	err := c.Visit(link)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Wait()
	return
}
