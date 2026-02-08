package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		s := `<html>
    <head>
        <title>busycard</title>
    </head>
    <body>
    	<h1>hello!<h1>
        <h2>let's create a business card on this weak!</h2>
        <section>
            <h3>projects</h3>
            <ul>
				<li>
					<p>
						<h4>games</h4>
						<ul>
							<li>
								<a href="https://github.com/kra53n/pytouch">pytouch</a>
							</li>
						</ul>
					</p>
				</li>
            </ul>
            <p>
                <h4>other</h4>
                <ul>
                    <li>
                        <a href="https://github.com/kra53n/curve">curve</a>
                    </li>
                </ul>
            </p>
            <p>
                <h4>web</h4>
                <ul>
                    <li>
                        <a href="https://github.com/TrondinDS/SNGGame">snggame</a>
                    </li>
                </ul>
            </p>
        </section>
        <footer>
            <a href="https://github.com/kra53n">github</a>
            |
            <a href="https://github.com/kra53n/busycard">src</a>
        </footer>
    </body>
</html>`
		io.WriteString(w, s)
	})

	http.ListenAndServe(":8080", nil)
}
