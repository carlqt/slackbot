package tweetbot

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

func FormatTweet(tweet string) string {
	thecodingloveLink := extractExternalURL(tweet)
	gifLink := extractGifLink(thecodingloveLink) //visitLink(link)
	// format tweet text to remove the urls
	formattedTweet := removeURL(tweet)
	return fmt.Sprintf(formattedTweet + "\n" + gifLink)
}

func extractExternalURL(text string) string {
	re := regexp.MustCompile(`(http|ftp|https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
	return re.FindStringSubmatch(text)[0]
}

func extractGifLink(externalLink string) (imgSrc string) {
	// visit the link to scrape
	client := &http.Client{}
	req, _ := http.NewRequest("GET", externalLink, nil)
	log.Println("Connecting to " + externalLink)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

L:
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			break L
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "img"
			if isAnchor {
				imgSrc = getSrc(t)
				break L
			}
		}
	}

	return imgSrc
}

func getSrc(t html.Token) (src string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, img := range t.Attr {
		if img.Key == "src" {
			src = img.Val
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}

func removeURL(tweet string) string {
	re := regexp.MustCompile(`(http|ftp|https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
	newStr := re.ReplaceAllString(tweet, "")
	return newStr
}
