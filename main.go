package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/emadgh/go-persian-tools"
	"github.com/nbjahan/go-jalali/jalali"
)

type WpItem struct {
	Title    string
	Content  string
	Tags     []string
	Date     time.Time
	Comments []*WpComment
}

type WpComment struct {
	Name    string
	Comment string
	Date    time.Time
}

var (
	url          string   = "http://shabeasheghan.blogfa.com/" //
	page_query   string   = "?p="                              // dont change this for blogfa.com
	current_page int      = 1                                  // start page
	post_links   []string = make([]string, 0)                  // stored links, don't change it
	page_limit   int      = -1                                 // number of pages to go
	chunk        int      = 5                                  // chunk json file, default 5 post per jsonfile

	posts []*WpItem

	// maximum number of workers ( Parallel Downloading pages)
	maxWorker     = 8
	wg            sync.WaitGroup
	posts_channel chan string
)

func main() {
	posts_channel = make(chan string)
	for i := 0; i < maxWorker; i++ {
		go Worker()
	}
	// fetching posts links
	for {
		doc, err := goquery.NewDocument(url + page_query + strconv.Itoa(current_page))
		if err != nil {
			continue // try until get it
		}
		doc.Find(".posttitle").Each(func(i int, s *goquery.Selection) {
			href, ok := s.Attr("href")
			if !ok {
				return
			}
			href = strings.TrimSpace(href)
			if href == "" {
				return
			}
			log.Println("link:" + href)

			wg.Add(1)
			posts_channel <- href
		})

		current_page++
		if page_limit > 0 && current_page > page_limit {
			break
		}
	}

	wg.Wait()

	plen := len(posts)
	for i := 0; i < (plen / chunk); i++ {
		min := i * chunk
		max := (i + 1) * chunk
		if max > plen {
			max = plen
		}
		json, _ := json.Marshal(posts[min:max])
		err := ioutil.WriteFile("posts_"+strconv.Itoa(i)+".json", json, 0644)
		log.Println(err)
	}
}

func getPost(pl string) {
	for {
		doc, err := goquery.NewDocument(url + pl)
		if err != nil {
			continue
		}
		sel := doc.Find(".post")

		tags := make([]string, 0)
		sel.Find(".tagname").Each(func(i int, s *goquery.Selection) {
			tags = append(tags, strings.TrimSpace(s.Text()))
		})
		_date := Date(sel.Find(".postdate").Text())
		htmlContent, err := sel.Find(".postcontent").Html()
		wp := &WpItem{
			Title:    strings.TrimSpace(sel.Find(".posttitle").Text()),
			Content:  strings.TrimSpace(htmlContent),
			Tags:     tags,
			Date:     _date,
			Comments: getComments(pl[6:]),
		}

		// log.Println(wp)
		posts = append(posts, wp)
		break
	}
}
func getComments(id string) []*WpComment {
	log.Println("comment:" + id)

	comments := make([]*WpComment, 0)
	for {
		doc, err := goquery.NewDocument(url + "comments/?blogid=shabeasheghan&postid=" + id)
		if err != nil {
			continue
		}
		doc.Find(".box").Each(func(i int, s *goquery.Selection) {
			comment, _ := s.Find(".body").Html()
			_date := Date(s.Find(".date").Text())
			comments = append(comments, &WpComment{
				Name:    s.Find(".author").Text(),
				Date:    _date,
				Comment: comment,
			})
		})
		break
	}
	return comments
}

func Worker() {
	for {
		select {
		case link := <-posts_channel:

			getPost(link)
			wg.Done()
		}
	}
}

func Date(d string) time.Time {
	d = strings.TrimSpace(d)
	if d == "" {
		return time.Now()
	}

	months := []string{"فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور", "مهر", "آبان", "آذر", "دی", "بهمن", "اسفند"}
	d = persiantools.NumbersToEnglish(d)

	ts := strings.Split(d, " ")
	if len(ts) < 3 {
		return time.Now()
	}
	day_ := ts[1]
	month_ := ts[2]
	year_ := ""
	mlen := len(month_)
	if mlen-4 > 0 {
		year_ = month_[len(month_)-4:]
		if _, err := strconv.Atoi(year_); err != nil {
			year_ = ts[3]
		} else {
			month_ = month_[:len(month_)-4]
		}
	} else {
		year_ = ts[3]
	}

	month := 0
	for i, m := range months {
		if month_ == m {
			month = i
		}
	}
	year, _ := strconv.Atoi(year_)
	day, _ := strconv.Atoi(day_)

	hour, min := 0, 0
	if len(ts) == 5 {
		hm := strings.Split(ts[4], ":")
		hour, _ = strconv.Atoi(hm[0])
		min, _ = strconv.Atoi(hm[1])
	}

	return jalali.Jtog(year, month, day, hour, min)
}
