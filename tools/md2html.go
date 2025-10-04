package tools

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Md2Html(inputFile, outputFile string) error {
	mdContent, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	htmlContent := convertMarkdownToHTML(string(mdContent))

	return os.WriteFile(outputFile, []byte(htmlContent), 0644)
}

func convertMarkdownToHTML(md string) string {
	lines := strings.Split(md, "\n")
	var buf bytes.Buffer

	buf.WriteString("<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>Markdown</title></head><body>\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			buf.WriteString("<br/>\n")
			continue
		}
		switch {
		case strings.HasPrefix(line, "### "):
			buf.WriteString(fmt.Sprintf("<h3>%s</h3>\n", escapeHTML(line[4:])))
		case strings.HasPrefix(line, "## "):
			buf.WriteString(fmt.Sprintf("<h2>%s</h2>\n", escapeHTML(line[3:])))
		case strings.HasPrefix(line, "# "):
			buf.WriteString(fmt.Sprintf("<h1>%s</h1>\n", escapeHTML(line[2:])))
		default:
			line = parseInlineMarkdown(line)
			buf.WriteString(fmt.Sprintf("<p>%s</p>\n", line))
		}
	}

	buf.WriteString("</body></html>")
	return buf.String()
}

func parseInlineMarkdown(s string) string {
	reBold := regexp.MustCompile(`\*\*(.*?)\*\*`)
	s = reBold.ReplaceAllString(s, "<strong>$1</strong>")

	reItalic := regexp.MustCompile(`\*(.*?)\*`)
	s = reItalic.ReplaceAllString(s, "<em>$1</em>")

	return escapeHTML(s)
}

func escapeHTML(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(s)
}
