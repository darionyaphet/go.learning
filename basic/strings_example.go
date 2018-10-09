package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	//Fields are: ["foo" "bar" "baz"]

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo1;bar2,baz3...", f))
	// Fields are: ["foo1" "bar2" "baz3"]

	//s := "<p>会有很多人问我，为什么我画的画总不像呢？或者一张画画像了很重要吗？" +
	//	"其实这些答案取决于你目前的需求，可以把画面画得写实具象，也可以很抽象地去传达什么。" +
	//	"\r</p><p>但大家如果处于绘画初级阶段或是把绘画当做一个爱好去培养，那么【画像了】就是一种绘画基本功能力的体现。" +
	//	"所以本次live也是对绘画基本功的讲课，会从起形构图，如何上调子，整体画面的处理等方面去讲解；" +
	//	"包括在使用不同的材料时，如水彩、水粉、彩铅、针管笔等等，如何去将画面画得写实，画得很像。" +
	//	"\r</p><p>观察方法很重要，这可能也是一个画不像的原因。" +
	//	"所以如果有对绘画感兴趣的或者正在学画的朋友，本次live将会很适合</p>"

	s := "<p><img src=\"v2 42d2ebb073269856a72f922e9d6be2dc.jpg\">高一的你，如何为每一科做好积累。</p>"
	s = strings.Replace(s, "\r", "", -1)
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	s = re.ReplaceAllString(s, "")
	//s = strings.Replace(s, "<p>", "", -1)
	//s = strings.Replace(s, "</p>", "", -1)
	fmt.Println(s)

	//fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	//fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

}
