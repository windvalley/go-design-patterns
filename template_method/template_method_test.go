package templatemethod

import (
	"fmt"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	templateA := NewTemplateStyleA()
	templateB := NewTemplateStyleB()

	homepage := NewHomepage(templateA)
	homepageA := homepage.Render("this is content")

	homepage = NewHomepage(templateB)
	homepageB := homepage.Render("this is content")

	fmt.Printf("Homepage that use templateStyleA: \n %s \n", homepageA)
	fmt.Printf("Homepage that use templateStyleB: \n %s \n", homepageB)
}
