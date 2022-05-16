package main

import (
	"bytes"
	"confessions-encoder/protos"
	_ "embed"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

//go:embed src.pro
var templateBytes []byte

//go:embed Arial.ttf
var arialBytes []byte

func main() {

	inputs := getLDs("./scrape")

	arial, err := truetype.Parse(arialBytes)
	if err != nil {
		log.Fatalln(err)
	}

	face := truetype.NewFace(arial, &truetype.Options{
		Size: 90,
		DPI:  84.73,
	})

	for _, item := range inputs {
		formatSlides(face, item)
	}

	fmt.Println("done")
}

func formatSlides(face font.Face, input *Input) {
	template := &protos.Presentation{}
	err := proto.Unmarshal(templateBytes, template)
	if err != nil {
		log.Fatalln(err)
	}

	blank := template.Cues[0]
	text := template.Cues[2]
	continuedEl := template.Cues[1].Actions[0].ActionTypeData.(*protos.Action_Slide).Slide.Slide.(*protos.Action_SlideType_Presentation).Presentation.BaseSlide.Elements[1]

	template.Uuid.String_ = uuid.NewString()
	template.Name = input.ShortName
	template.LastModifiedDate.Seconds = time.Now().Unix()
	template.LastModifiedDate.Nanos = 0
	template.LastDateUsed.Seconds = template.LastModifiedDate.Seconds
	template.LastDateUsed.Nanos = template.LastModifiedDate.Nanos
	template.Cues = nil
	template.CueGroups = nil

	addBlankSlide(template, blank)

	sections := splitSlides(face, input.Content, 8, func(value string) int {
		// if there's a recent break, use the last little bit of text on the next section
		i := strings.LastIndex(value, "\n\n")
		if i == -1 {
			return 0
		}

		if len(value)-i < 50 {
			return -(len(value) - i)
		}

		return 0
	})

	for i, section := range sections {
		addTextSlide(template, text, continuedEl, &TextInput{
			Title:     input.DisplayName,
			Content:   section,
			Continued: i+1 < len(sections),
		})
	}

	raw, err := proto.Marshal(template)
	if err != nil {
		log.Fatalln(err)
	}

	fname := "./output/" + input.FileName
	ioutil.WriteFile(fname, raw, os.ModePerm)
	fmt.Println("wrote: ", fname)
}

func splitSlides(face font.Face, value string, nLinesPerSlide int, lookBack func(chunk string) int) []string {
	oneLine := `iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii`

	iWidth, _ := face.GlyphAdvance(rune(oneLine[0]))
	boxWidth := iWidth * fixed.Int26_6(len(oneLine))

	lineCt := 0
	var width fixed.Int26_6
	startI := 0
	section := []string{}
	sectionHasNonWhiteSpace := false

	checkSection := func(i int) int {
		if lineCt < nLinesPerSlide {
			return i
		}

		endI := i + 1
		proposedChunk := value[startI:endI]
		updateEndIndex := lookBack(proposedChunk)
		endI += updateEndIndex

		section = append(section, value[startI:endI])
		startI = endI
		lineCt = 0
		sectionHasNonWhiteSpace = false
		return startI
	}

	for i := 0; i < len(value); i++ {
		r := rune(value[i])

		if r == '\n' {
			if !sectionHasNonWhiteSpace {
				startI = i + 1
				width = 0
				continue
			}

			lineCt++
			width = 0
			i = checkSection(i)
			continue
		}

		if r != ' ' && string(r) != "\t" {
			if !sectionHasNonWhiteSpace {
				sectionHasNonWhiteSpace = true
			}
		}

		w, ok := face.GlyphAdvance(r)
		if !ok {
			log.Fatalln("font missing glyph for rune '"+string(r)+"' (", int(r), ")")
		}

		width += w

		if width > boxWidth {
			lineCt++
			width = 0
			i = checkSection(i)
		}
	}

	if startI+1 < len(value) {
		section = append(section, value[startI:])
	}

	return section
}

func cloneSlide(template *protos.Cue) *protos.Cue {
	cue := &protos.Cue{}
	bytes, _ := proto.Marshal(template)
	proto.Unmarshal(bytes, cue)

	act := cue.Actions[0]
	act.Uuid.String_ = uuid.NewString()

	sl := act.ActionTypeData.(*protos.Action_Slide)
	slide := sl.Slide.Slide.(*protos.Action_SlideType_Presentation)
	base := slide.Presentation.BaseSlide
	base.Uuid.String_ = uuid.NewString()

	for _, el := range base.Elements {
		el.Element.Uuid.String_ = uuid.NewString()
	}

	return cue
}

type TextInput struct {
	Title     string
	Content   string
	Continued bool
}

func addTextSlide(pres *protos.Presentation, template *protos.Cue, continuedEl *protos.Slide_Element, input *TextInput) {
	cue := cloneSlide(template)
	slide := cue.Actions[0].ActionTypeData.(*protos.Action_Slide).Slide.Slide.(*protos.Action_SlideType_Presentation).Presentation.BaseSlide
	setTitle(slide.Elements[0], input.Title)
	setContent(slide.Elements[1], input.Content)

	if input.Continued {
		continuedEl.Element.Uuid.String_ = uuid.NewString()
		slide.Elements = append(slide.Elements, continuedEl)
	}

	addSlideToPresentation(pres, cue)
}

func setContent(el *protos.Slide_Element, text string) {
	el.Element.Text.RtfData = []byte(`{\rtf0\ansi\ansicpg1252{\fonttbl\f0\fnil ArialMT;}{\colortbl;\red255\green255\blue255;\red255\green255\blue255;}{\*\expandedcolortbl;\csgenericrgb\c100000\c100000\c100000\c100000;\csgenericrgb\c100000\c100000\c100000\c0;}{\*\listtable}{\*\listoverridetable}\uc1\paperw36948\margl0\margr0\margt0\margb0\pard\li0\fi0\ri0\ql\sb0\sa0\sl240\slmult1\slleading0\f0\b0\i0\ul0\strike0\fs180\expnd0\expndtw0\cf1\strokewidth0\strokec1\nosupersub\ulc0\highlight2\cb2\par\pard\li0\fi0\ri0\ql\sb0\sa0\sl240\slmult1\slleading0\f0\b0\i0\ul0\strike0\fs180\expnd0\expndtw0\cf1\strokewidth0\strokec1\nosupersub\ulc0\highlight2\cb2 ` + sanitizeRtf(text) + `}`)
}

func sanitizeRtf(text string) string {
	text = strings.ReplaceAll(text, "\t", `     `)
	text = strings.ReplaceAll(text, "\n", `\par `)
	return text
}

func setTitle(el *protos.Slide_Element, text string) {
	el.Element.Text.RtfData = []byte(`{\rtf0\ansi\ansicpg1252{\fonttbl\f0\fnil Arial-BoldMT;}{\colortbl;\red255\green255\blue255;\red255\green255\blue255;}{\*\expandedcolortbl;\csgenericrgb\c100000\c100000\c100000\c100000;\csgenericrgb\c100000\c100000\c100000\c0;}{\*\listtable}{\*\listoverridetable}\uc1\paperw26074\margl0\margr0\margt0\margb0\pard\li0\fi0\ri0\ql\sb0\sa0\sl240\slmult1\slleading0\f0\b\i0\ul0\strike0\fs200\expnd0\expndtw0\cf1\strokewidth0\strokec1\nosupersub\ulc0\highlight2\cb2 ` + sanitizeRtf(text) + `}`)
}

func addBlankSlide(pres *protos.Presentation, template *protos.Cue) {
	addSlideToPresentation(pres, cloneSlide(template))
}

func addSlideToPresentation(pres *protos.Presentation, item *protos.Cue) {
	item.Uuid.String_ = uuid.NewString()

	pres.SelectedArrangement.String_ = uuid.NewString() // random
	pres.Cues = append(pres.Cues, item)
	pres.CueGroups = append(pres.CueGroups, &protos.Presentation_CueGroup{
		Group: &protos.Group{
			Name: "Group",
			Uuid: &protos.UUID{
				String_: uuid.NewString(),
			},
		},
		CueIdentifiers: []*protos.UUID{{
			String_: item.Uuid.String_,
		}},
	})
}

type Input struct {
	FileName    string
	ShortName   string
	DisplayName string
	Content     string
}

func getLDs(dir string) []*Input {
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	out := []*Input{}

	for _, item := range list {
		if item.IsDir() {
			continue
		}

		if path.Ext(item.Name()) == ".html" {
			value := parseLD(filepath.Join(dir, item.Name()))
			fmt.Println("> ", value.ShortName)
			fmt.Println("> ", value.FileName)
			fmt.Println("> ", value.DisplayName)
			fmt.Println("> ", value.Content)
			out = append(out, value)
		}
	}

	return out
}

func parseLD(file string) *Input {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		log.Fatalln(err)
	}

	tbl := firstChildWithClassName(doc, "question-and-answer")
	row := firstChildWithTagName(tbl, "tr")
	buf := &bytes.Buffer{}

	for row != nil {
		if row.Type != html.ElementNode {
			row = row.NextSibling
			continue
		}

		if firstChildWithClassName(row, "ld-footnotes") != nil {
			// ignore footnotes
			row = row.NextSibling
			continue
		}

		numberEl := firstChildWithClassName(row, "number")
		if numberEl != nil {
			n := strings.TrimSuffix(numberEl.FirstChild.Data, ".")
			buf.WriteString(n + ". ")
		}

		titleEl := firstChildWithTagName(row, "h2")
		if titleEl != nil {
			title := extractText(titleEl, nil, nil).String()
			title = strings.ReplaceAll(strings.ReplaceAll(title, "\n", " "), "  ", " ")
			buf.WriteString(title + "\n")
		} else {

			buf.WriteString("\n")
			lastIndent := false

			for _, item := range childrenWithClass(row, "line") {
				isIndent := hasClass(item, "indented")

				if lastIndent && !isIndent {
					// add extra line break after paragraph
					buf.WriteString("\n")
				}

				lastIndent = isIndent

				extractText(item, buf, func(item *html.Node) bool {
					// ignore superscripts
					if item.Type == html.ElementNode && item.Data == "sup" {
						return false
					}

					return true
				})
			}
		}

		row = row.NextSibling
	}

	name := filepath.Base(file)
	number := strings.ReplaceAll(strings.ReplaceAll(name, "LD-", ""), ".html", "")

	return &Input{
		ShortName:   fmt.Sprintf("LD-%s", number),
		FileName:    fmt.Sprintf("LD-%s.pro", number),
		DisplayName: fmt.Sprintf("Lord's Day %s", number),
		Content:     buf.String(),
	}
}

func childrenWithClass(node *html.Node, search string) []*html.Node {
	var list []*html.Node

	if hasClass(node, search) {
		list = append(list, node)
		return list
	}

	cur := node.FirstChild
	for cur != nil {
		list = append(list, childrenWithClass(cur, search)...)
		cur = cur.NextSibling
	}

	return list
}

func extractText(node *html.Node, buf *bytes.Buffer, filter func(node *html.Node) bool) *bytes.Buffer {
	if buf == nil {
		buf = &bytes.Buffer{}
	}

	if filter != nil && !filter(node) {
		return buf
	}

	if node.Type == html.ElementNode {
		cur := node.FirstChild

		if hasClass(node, "indented") {
			buf.WriteString("\t")
		}

		for cur != nil {
			extractText(cur, buf, filter)
			cur = cur.NextSibling
		}

		if hasClass(node, "line") {
			buf.WriteString("\n")
		}
	}

	if node.Type == html.TextNode {
		buf.WriteString(node.Data)
	}

	return buf
}

func firstChildWithTagName(node *html.Node, search string) *html.Node {
	return firstChildWithMatch(node, func(node *html.Node) bool {
		return node.Type == html.ElementNode && node.Data == search
	})
}

func firstChildWithClassName(node *html.Node, search string) *html.Node {
	return firstChildWithMatch(node, func(node *html.Node) bool {
		return hasClass(node, search)
	})
}

func firstChildWithMatch(node *html.Node, matcher func(node *html.Node) bool) *html.Node {

	if matcher(node) {
		return node
	}

	cur := node.FirstChild
	if cur == nil {
		return nil
	}

	for {
		v := firstChildWithMatch(cur, matcher)
		if v != nil {
			return v
		}

		if cur.NextSibling != nil {
			cur = cur.NextSibling
		} else {
			break
		}
	}

	return nil
}

func hasClass(c *html.Node, search string) bool {
	for _, attr := range c.Attr {
		if attr.Key == "class" {
			if contains(strings.Split(attr.Val, " "), search) {
				return true
			}
		}
	}

	return false
}

func contains(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}
