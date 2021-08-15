package utils

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"proj-review/log"
)

func readAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func markdownToHTML(md string) string {
	myHTMLFlags := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	myExtensions := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS |
		blackfriday.EXTENSION_HARD_LINE_BREAK

	renderer := blackfriday.HtmlRenderer(myHTMLFlags, "", "")
	bytes := blackfriday.MarkdownOptions([]byte(md), renderer, blackfriday.Options{
		Extensions: myExtensions,
	})
	theHTML := string(bytes)
	return bluemonday.UGCPolicy().Sanitize(theHTML)
}

func GenApiDoc() {
	path := "views/api.md"
	res, err := readAll(path)
	if err != nil {
		log.Logger.Warn("Build API docs error: " + err.Error())
		return
	}

	result := markdownToHTML(string(res))
	// 写入到文件
	dstFile, err := os.Create("views/api.html")
	if err != nil {
		log.Logger.Warn("Build API docs error: " + err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(result + "\n")
}
