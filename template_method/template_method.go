package templatemethod

import "fmt"

// WebsiteTemplate ...
type WebsiteTemplate interface {
	Header() string
	Footer() string
}

// Homepage ...
type Homepage struct {
	WebsiteTemplate
}

// Render ...
func (h *Homepage) Render(content string) string {
	return fmt.Sprintf("%s \n %s \n %s \n", h.Header(), content, h.Footer())
}

// NewHomepage ...
func NewHomepage(template WebsiteTemplate) *Homepage {
	return &Homepage{template}
}

type templateStyleA struct{}

// NewTemplateStyleA ...
func NewTemplateStyleA() WebsiteTemplate {
	return &templateStyleA{}
}

func (*templateStyleA) Header() string {
	return "Welcome to my blog"
}

func (*templateStyleA) Footer() string {
	return "Copyright@2009-2020 Windvalley"
}

type templateStyleB struct{}

// NewTemplateStyleB ...
func NewTemplateStyleB() WebsiteTemplate {
	return &templateStyleB{}
}

func (*templateStyleB) Header() string {
	return "Cogito ergo sum"
}

func (*templateStyleB) Footer() string {
	return "2009-2020@Windvalley"
}
