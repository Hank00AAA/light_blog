package markdown

import(
	"fmt"
	"github.com/gomarkdown/markdown"
)

// MarkDownData
type MarkDownData struct {
	Data []byte

}

// NewMarkDown
func NewMarkdown() *MarkDownData {
	return &MarkDownData{}
}

// MarkdownAddEmptyRow
func (m *MarkDownData)MarkdownAddEmptyRow() {
	m.Data = append(m.Data, []byte("  \n  ")...)
}

// MarkdownToHtml
func (m *MarkDownData)MarkdownToHtml() []byte {
	return markdown.ToHTML(m.Data, nil, nil)
}

// MarkdownHttp
func (m *MarkDownData)MarkdownHttp(name string, url string) {
	m.Data = append(m.Data, []byte(fmt.Sprintf("[%v](%v)", name, url))...)
}

// MarkdownTitle
func (m *MarkDownData)MarkdownAddTitle(level int, name string) {
	m.Data = append(m.Data, []byte(fmt.Sprintf("# %v", name))...)
	m.MarkdownAddEmptyRow()
}

