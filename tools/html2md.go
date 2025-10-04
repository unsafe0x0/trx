package tools

import (
	"os"
	"regexp"
	"strings"
)

func Html2Md(inputFile, outputFile string) error {
	htmlContent, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	mdContent := convertHTMLToMarkdown(string(htmlContent))

	return os.WriteFile(outputFile, []byte(mdContent), 0644)
}

func convertHTMLToMarkdown(html string) string {
	html = regexp.MustCompile(`(?i)<(html|head|body)[^>]*>`).ReplaceAllString(html, "")
	html = regexp.MustCompile(`(?i)</(html|head|body)>`).ReplaceAllString(html, "")
	html = regexp.MustCompile(`(?i)<h1>(.*?)</h1>`).ReplaceAllString(html, "# $1\n")
	html = regexp.MustCompile(`(?i)<h2>(.*?)</h2>`).ReplaceAllString(html, "## $1\n")
	html = regexp.MustCompile(`(?i)<h3>(.*?)</h3>`).ReplaceAllString(html, "### $1\n")
	html = regexp.MustCompile(`(?i)<p>(.*?)</p>`).ReplaceAllString(html, "$1\n")
	html = regexp.MustCompile(`(?i)<br\s*/?>`).ReplaceAllString(html, "\n")
	html = regexp.MustCompile(`(?i)<(strong|b)>(.*?)</\\1>`).ReplaceAllString(html, "**$2**")
	html = regexp.MustCompile(`(?i)<(em|i)>(.*?)</\\1>`).ReplaceAllString(html, "*$2*")
	html = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "&amp;", "&")
	html = strings.ReplaceAll(html, "&lt;", "<")
	html = strings.ReplaceAll(html, "&gt;", ">")
	html = strings.ReplaceAll(html, "&quot;", "\"")
	html = strings.ReplaceAll(html, "&#39;", "'")
	html = regexp.MustCompile(`\n+`).ReplaceAllString(html, "\n")

	return strings.TrimSpace(html)
}
