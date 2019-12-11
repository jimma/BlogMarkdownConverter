package blogindex

import (
	"bytes"
	"io"
	"strings"
	"sync"

	"github.com/antchfx/htmlquery"
	"github.com/lunny/html2md"
	"golang.org/x/net/html"
)

type BlogIndex struct {
	URL     string
	XPath   string
	Entries []string
}

type Blog struct {
	URL        string
	MDFileName string

	Title   string
	Author  string
	Date    string
	Content string
}

const (
	BlogIndexPath  string = "/html/body/div/section/div/div/div/div/section/header/div/h1/a"
	BlogTitleXPath string = "/html/body/div/section/div/div/div/div/div/header/div/h1/a"

	BlogAuthorXPath  string = "/html/body/div/section/div/div/div/div/div/header/div/span/a[1]"
	BlogDataXPath    string = "/html/body/div/section/div/div/div/div/div/header/div/span/text()[3]"
	BlogContentXpath string = "/html/body/div/section/div/div/div/div/div/section"
)

func (blogIndex *BlogIndex) Parse() ([]string, error) {
	doc, err := htmlquery.LoadURL(blogIndex.URL)

	if err != nil {
		return nil, err
	}

	nodes, err := htmlquery.QueryAll(doc, blogIndex.XPath)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, node := range nodes {
		res = append(res, node.Attr[0].Val)

	}
	blogIndex.Entries = res
	return res, nil
}

func (blogIndex *BlogIndex) GetBlogEntry() []Blog {
	var blogs []Blog
	var wg sync.WaitGroup
	for _, url := range blogIndex.Entries {
		res := convetToFileName(url)
		blog := Blog{
			URL:        url,
			MDFileName: res,
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			blog.GetBlogContet()
			blogs = append(blogs, blog)
		}()

	}
	wg.Wait()
	return blogs
}

func (blog *Blog) GetBlogContet() error {

	doc, err := htmlquery.LoadURL(blog.URL)

	if err != nil {
		return err
	}
	titleNodes, err := htmlquery.QueryAll(doc, BlogTitleXPath)
	if err != nil {
		return err
	}
	blog.Title = titleNodes[0].FirstChild.Data

	dataNode, err := htmlquery.QueryAll(doc, BlogDataXPath)
	if err != nil {
		return err
	}
	blog.Date = formateData(dataNode[0].Data)

	authorNode, err := htmlquery.QueryAll(doc, BlogAuthorXPath)
	if err != nil {
		return err
	}
	blog.Author = authorNode[0].FirstChild.Data
	contentNode, err := htmlquery.QueryAll(doc, BlogContentXpath)
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	buf.WriteString("---\n")
	buf.WriteString("layout:     post\n")
	buf.WriteString("title:       " + "\"" + blog.Title + "\"\n")
	buf.WriteString("subtitle:   \"\"\n")
	buf.WriteString("date:       " + blog.Date + "\n")
	buf.WriteString("author:     " + blog.Author + "\n")
	buf.WriteString("---\n")
	md := html2md.Convert(printNode(contentNode[0]))
	blog.Content = buf.String() + md
	return nil
}

func convetToFileName(url string) string {
	a := []rune(url)
	length := len(url)
	res := string(a[52:length])
	res = strings.ReplaceAll(res, "/", "-")
	return res
}
func formateData(data string) string {
	data = strings.ReplaceAll(data, "\n", "")
	a := []rune(data)
	res := string(a[4:27])
	return res
}

func printNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()

}
