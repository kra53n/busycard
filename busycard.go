package main

import (
	// "io"
	"bytes"
	"net/http"
	"os"
)

const Title string = "busycard"

func main() {

	var b bytes.Buffer
	b.WriteString("<html>")
	head(&b)
	body(&b)
	b.WriteString("</html>")

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		// io.WriteString(w, s)
		w.Write(b.Bytes())
	})

	//println(b.String())
	println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func cube(b *bytes.Buffer) {
	content, err := os.ReadFile("cube.js")
	if err != nil {
		panic("can't open cube.js")
	}
	b.WriteString(`<div id="scene-container">
<div class="corner tl"></div>
<div class="corner tr"></div>
<div class="corner bl"></div>
<div class="corner br"></div>
</div>`)
	b.WriteString(`<div id="controls">
<button id="btn-wireframe">Wireframe</button>
<button id="btn-solid" class="active">Solid</button>
<button id="btn-pause">Pause</button>
<button id="btn-explode">Explode</button>
</div>`)
	b.WriteString(`<script src="https://cdnjs.cloudflare.com/ajax/libs/three.js/r128/three.min.js"></script>`)
	tagfunc(b, "script", func(b *bytes.Buffer) {
		b.Write(content)
	})
}

func style(b *bytes.Buffer) {
	content, err := os.ReadFile("style.css")
	if err != nil {
		panic("can't open style.css")
	}
	tagfunc(b, "style", func(b *bytes.Buffer) {
		b.Write(content)
	})
}

func head(b *bytes.Buffer) {
	b.WriteString("<head>")

	b.WriteString("<title>")
	b.WriteString(Title)
	b.WriteString("</title>")

	style(b)

	b.WriteString("</head>")
}

func body(b *bytes.Buffer) {
	b.WriteString("<body>")
	tag(b, "h2", "let's create a business card on this weak!")
	cube(b)
	projects(b)
	footer(b)
	b.WriteString("</body>")
}

func projects(b *bytes.Buffer) {
	type Item struct {
		name string
		href string
	}

	type Section struct {
		name  string
		items []Item
	}

	sections := []Section{
		{
			name: "somesection",
			items: []Item{
				{"curve", "https://github.com/kra53n/curve"},
				{"project2", "https://github.com/kra53n/project2"},
			},
		},
		{
			name: "anothersection",
			items: []Item{
				{"tool", "https://github.com/kra53n/tool"},
				{"lib", "https://github.com/kra53n/lib"},
			},
		},
	}

	tag(b, "h3", "projects")
	tagfunc(b, "ul", func(b *bytes.Buffer) {
		for _, section := range sections {
			tagfunc(b, "li", func(b *bytes.Buffer) {
				// tagfunc(b, "p", func(b *bytes.Buffer) {

				tag(b, "h4", section.name)
				tagfunc(b, "ul", func(b *bytes.Buffer) {
					for _, elem := range section.items {
						tagfunc(b, "li", func(b *bytes.Buffer) {
							a(b, elem.name, elem.href)
						})
					}
				})

				// })
			})
		}
	})
}

func footer(b *bytes.Buffer) {
	b.WriteString("<footer>")
	a(b, "github", "https://github.com/kra53n")
	b.WriteString(" | ")
	a(b, "src", "https://github.com/kra53n/busycard")
	b.WriteString("</footer>")
}

func tag(b *bytes.Buffer, name, val string) {
	b.WriteByte('<')
	b.WriteString(name)
	b.WriteByte('>')
	b.WriteString(val)
	b.WriteString("</")
	b.WriteString(name)
	b.WriteByte('>')
}

func tagfunc(b *bytes.Buffer, name string, f func(b *bytes.Buffer)) {
	b.WriteByte('<')
	b.WriteString(name)
	b.WriteByte('>')
	f(b)
	b.WriteString("</")
	b.WriteString(name)
	b.WriteByte('>')
}

func a(b *bytes.Buffer, val, href string) {
	b.WriteString("<a ")
	b.WriteString("href=\"")
	b.WriteString(href)
	b.WriteString("\">")
	b.WriteString(val)
	b.WriteString("</a>")
}
