package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
)


const (
	html      = "html"
	tagCloser = "/"
	starter   = "<"
	closer    = ">"
	CSSclass  = "class"
)


var validTags = map[string]bool{
	"head": true, "body": true, "footer": true, "title": true, "div": true, "section": true, "p": true,
}


type Section struct {
	Tag     string
	Content string
}


type HTMLbuilder interface {
	SetSection(tag string, content string, css ...string) (template.HTML, error)
	Build(sections []Section)
	SaveToFile(name string) error
}


type PageBuilder struct {
	content template.HTML
}


func NewPageBuilder() *PageBuilder {
	return &PageBuilder{content: ""}
}


func (p *PageBuilder) SetSection(tag string, content string, css ...string) (template.HTML, error) {
	if !isValidSectionTag(tag) {
		return "", fmt.Errorf("not a valid tag")
	}


	tagContent := tag
	if len(css) > 0 {
		tagContent += ` class="` + strings.Join(css, " ") + `"`
	}


	tagArr := []string{starter, tagContent, closer, content, starter, tagCloser, tag, closer}
	return template.HTML(strings.Join(tagArr, "")), nil
}


func isValidSectionTag(tag string) bool {
	return validTags[tag]
}


func (p *PageBuilder) Build(sections []Section) {
	var page template.HTML


	startHTML := starter + html + closer
	page += template.HTML(startHTML)


	for _, section := range sections {
		sectionHTML, err := p.SetSection(section.Tag, section.Content)
		if err != nil {
			fmt.Printf("tag %s is not in the list of section tags, skipped\n", section.Tag)
			continue
		}
		page += sectionHTML
	}


	closeHTML := starter + tagCloser + html + closer
	page += template.HTML(closeHTML)


	p.content = page
}


func (p *PageBuilder) SaveToFile(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("error occurred while creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(string(p.content))
	if err != nil {
		return fmt.Errorf("error occurred while writing to file: %v", err)
	}
	return nil
}


func ParseJson(path string) ([]Section, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	var result []Section
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	return result, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <command> <path-to-json>")
		os.Exit(1)
	}

	
	builder := NewPageBuilder()


	configContent, err := ParseJson(os.Args[1])
	if err != nil {
		fmt.Println("Can't parse JSON:", err)
		os.Exit(1)
	}

	builder.Build(configContent)


	if err := builder.SaveToFile("test.html"); err != nil {
		fmt.Println("Error saving file:", err)
		os.Exit(1)
	}

	fmt.Println("The program has finished generating the HTML page.")
}
