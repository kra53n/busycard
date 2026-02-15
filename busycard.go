package main

import (
	// "io"
	"net/http"
	"bytes"
)

const Title string = "busycard"

func head(b *bytes.Buffer) {
	b.WriteString("<head>")
	b.WriteString("<title>")
	b.WriteString(Title)
	b.WriteString("</title>")
	b.WriteString("</head>")
}

func body(b *bytes.Buffer) {
	b.WriteString("<body>")
	tag(b, "h2", "let's create a business card on this weak!")
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

func main() {

	var b bytes.Buffer
	b.WriteString("<html>")
	head(&b)
	body(&b)
	b.WriteString("</html>")

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
// 		s := `<html>
//     <head>
//         <title>busycard</title>
//     </head>
//     <body>
//     	<h1>hello!<h1>
//         <h2>let's create a business card on this weak!</h2>
//         <section>
//             <h3>projects</h3>
//             <ul>
// 				<li>
// 					<p>
// 						<h4>games</h4>
// 						<ul>
// 							<li>
// 								<a href="https://github.com/kra53n/pytouch">pytouch</a>
// 							</li>
// 						</ul>
// 					</p>
// 				</li>
//             </ul>
//             <p>
//                 <h4>other</h4>
//                 <ul>
//                     <li>
//                         <a href="https://github.com/kra53n/curve">curve</a>
//                     </li>
//                 </ul>
//             </p>
//             <p>
//                 <h4>web</h4>
//                 <ul>
//                     <li>
//                         <a href="https://github.com/TrondinDS/SNGGame">snggame</a>
//                     </li>
//                 </ul>
//             </p>
//         </section>
//         <footer>
//             <a href="https://github.com/kra53n">github</a>
//             |
//             <a href="https://github.com/kra53n/busycard">src</a>
//         </footer>
//     </body>
// </html>`
		// io.WriteString(w, s)
		w.Write(b.Bytes())
	})

	println(b.String())
	println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
