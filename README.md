<a href="https://travis-ci.org/iris-framework/iris"><img src="https://api.travis-ci.org/iris-framework/iris.svg?branch=v6&style=flat-square" alt="Build Status"></a>
<a href="http://goreportcard.com/report/iris-framework/iris"><img src="https://img.shields.io/badge/report%20card%20-a%2B-F44336.svg?style=flat-square" alt="http://goreportcard.com/report/iris-framework/iris"></a>
<a href="http://support.iris-go.com"><img src="https://img.shields.io/badge/support-page-ec2eb4.svg?style=flat-square" alt="Iris support forum"></a> <!-- <a href="https://github.com/iris-framework/iris/blob/v6/HISTORY.md"><img src="https://img.shields.io/badge/codename-√Νεxτ%20-blue.svg?style=flat-square" alt="CHANGELOG/HISTORY"></a> --><a href="https://github.com/iris-framework/iris/tree/v6/_examples#table-of-contents"><img src="https://img.shields.io/badge/examples-%20repository-3362c2.svg?style=flat-square" alt="Examples for new Gophers"></a>
<a href="https://godoc.org/gopkg.in/iris-framework/iris.v6"><img src="https://img.shields.io/badge/docs-%20reference-5272B4.svg?style=flat-square" alt="Docs"></a>
<a href="https://iris-framework.rocket.chat/channel/iris"><img src="https://img.shields.io/badge/community-%20chat-00BCD4.svg?style=flat-square" alt="Chat"></a>
<br/>
Iris is an <b>efficient</b> and well-designed, <b>cross-platform, web framework</b> with robust set of <b>features</b>.<br/>Build your own <b>high-performance</b> web applications and <b>APIs</b> powered by unlimited <b>potentials and portability</b>.

If you're coming from <a href="https://nodejs.org/en/">Node.js</a> world, this is the <a href="https://github.com/expressjs/express">expressjs</a> equivalent for the <a href="https://golang.org">Go Programming Language.</a>
</p>


Feature Overview
-----------

- Focus on high performance
- Automatically install and serve certificates from https://letsencrypt.org
- Robust routing and middleware ecosystem
- Build RESTful APIs
- Choose your favorite routes' path syntax between [httprouter](https://github.com/iris-framework/iris/blob/v6/_examples/beginner/routes-using-httprouter/main.go) and [gorillamux](https://github.com/iris-framework/iris/blob/v6/_examples/beginner/routes-using-gorillamux/main.go)
- Request-Scoped Transactions
- Group API's and subdomains with wildcard support
- Body binding for JSON, XML, Forms, can be extended to use your own custom binders
- More than 50 handy functions to send HTTP responses
- View system: supporting more than 6+ template engines, with prerenders. You can still use your favorite
- Define virtual hosts and (wildcard) subdomains with path level routing
- Graceful shutdown
- Limit request body
- Localization i18N
- Serve static files
- Cache
- Log requests
- Customizable format and output for the logger
- Customizable HTTP errors
- Compression (Gzip)
- Authentication
 - OAuth, OAuth2 supporting 27+ popular websites
 - JWT
 - Basic Authentication
 - HTTP Sessions
- Add / Remove trailing slash from the URL with option to redirect
- Redirect requests
 - HTTP to HTTPS
 - HTTP to HTTPS WWW
 - HTTP to HTTPS non WWW
 - Non WWW to WWW
 - WWW to non WWW
- Highly scalable rich content render (Markdown, JSON, JSONP, XML...)
- Websocket-only API similar to socket.io
- Hot Reload on source code changes
- Typescript integration + Web IDE
- Checks for updates at startup
- Highly customizable
- Feels like you used iris forever, thanks to its Fluent API
- And many others...


Table of Contents
-----------

<a href="https://github.com/iris-framework/iris/tree/v6/_examples#table-of-contents"><img align="right" width="265" src="https://raw.githubusercontent.com/iris-contrib/website/gh-pages/assets/book/cover_4.jpg"></a>


* [Level: Beginner](_examples/beginner)
    * [Hello World](_examples/beginner/hello-world/main.go)
    * [Routes (using httprouter)](_examples/beginner/routes-using-httprouter/main.go)
    * [Routes (using gorillamux)](_examples/beginner/routes-using-gorillamux/main.go)
    * [Internal Application File Logger](_examples/beginner/file-logger/main.go)
    * [Write JSON](_examples/beginner/write-json/main.go)
    * [Read JSON](_examples/beginner/read-json/main.go)
    * [Read Form](_examples/beginner/read-form/main.go)
    * [Favicon](_examples/beginner/favicon/main.go)
    * [File Server](_examples/beginner/file-server/main.go)
    * [Send Files](_examples/beginner/send-files/main.go)
    * [Stream Writer](_examples/beginner/stream-writer/main.go)
    * [Listen UNIX Socket](_examples/beginner/listen-unix/main.go)
    * [Listen TLS](_examples/beginner/listen-tls/main.go)
    * [Listen Letsencrypt (Automatic Certifications)](_examples/beginner/listen-letsencrypt/main.go)
* [Level: Intermediate](_examples/intermediate)
    * [Send An E-mail](_examples/intermediate/e-mail/main.go)
    * [Upload/Read Files](_examples/intermediate/upload-files/main.go)
    * [Request Logger](_examples/intermediate/request-logger/main.go)
    * [Profiling (pprof)](_examples/intermediate/pprof/main.go)
    * [Basic Authentication](_examples/intermediate/basicauth/main.go)
    * [HTTP Access Control](_examples/intermediate/cors/main.go)
    * [Cache Markdown](_examples/intermediate/cache-markdown/main.go)
    * [Localization and Internationalization](_examples/intermediate/i18n/main.go)
    * [Recovery](_examples/intermediate/recover/main.go)
    * [Graceful Shutdown](_examples/intermediate/graceful-shutdown/main.go)
    * [Custom TCP Listener](_examples/intermediate/custom-listener/main.go)
    * [Custom HTTP Server](_examples/intermediate/custom-httpserver/main.go)
    * [View Engine](_examples/intermediate/view)
        * [Overview](_examples/intermediate/view/overview/main.go)
        * [Template HTML: Part Zero](_examples/intermediate/view/template_html_0/main.go)
        * [Template HTML: Part One](_examples/intermediate/view/template_html_1/main.go)
        * [Template HTML: Part Two](_examples/intermediate/view/template_html_2/main.go)
        * [Template HTML: Part Three](_examples/intermediate/view/template_html_3/main.go)
        * [Template HTML: Part Four](_examples/intermediate/view/template_html_4/main.go)
        * [Inject Data Between Handlers](_examples/intermediate/view/context-view-data/main.go)
        * [Embedding Templates Into Executable](_examples/intermediate/view/embedding-templates-into-app)
        * [Custom Renderer](_examples/intermediate/view/custom-renderer/main.go)
    * [Password Hashing](_examples/intermediate/password-hashing/main.go)
    * [Sessions](_examples/intermediate/sessions)
        * [Overview](_examples/intermediate/sessions/overview/main.go)
        * [Encoding & Decoding the Session ID: Secure Cookie](_examples/intermediate/sessions/securecookie/main.go)
        * [Standalone](_examples/intermediate/sessions/standalone/main.go)
        * [With A Back-End Database](_examples/intermediate/sessions/database/main.go)
    * [Flash Messages](_examples/intermediate/flash-messages/main.go)
    * [Websockets](_examples/intermediate/websockets)
        * [Ridiculous Simple](_examples/intermediate/websockets/ridiculous-simple/main.go)
        * [Overview](_examples/intermediate/websockets/overview/main.go)
        * [Connection List](_examples/intermediate/websockets/connectionlist/main.go)
        * [Native Messages](_examples/intermediate/websockets/naive-messages/main.go)
        * [Secure](_examples/intermediate/websockets/secure/main.go)
        * [Custom Go Client](_examples/intermediate/websockets/custom-go-client/main.go)
* [Level: Advanced](_examples/advanced)
    * [Transactions](_examples/advanced/transactions/main.go)
    * [HTTP Testing](_examples/advanced/httptest/main_test.go)
    * [Watch & Compile Typescript source files](_examples/advanced/typescript/main.go)
    * [Cloud Editor](_examples/advanced/cloud-editor/main.go)
    * [Online Visitors](_examples/advanced/online-visitors/main.go)
    * [URL Shortener using BoltDB](_examples/advanced/url-shortener/main.go)
    * [Subdomains](_examples/advanced/subdomains)
        * [Single](_examples/advanced/subdomains/single/main.go)
        * [Multi](_examples/advanced/subdomains/multi/main.go)
        * [Wildcard](_examples/advanced/subdomains/wildcard/main.go)

Installation
-----------

The only requirement is the [Go Programming Language](https://golang.org/dl/), at least 1.8

```sh
$ go get gopkg.in/iris-framework/iris.v6
```

For further installation support, navigate [here](http://support.iris-go.com/d/16-how-to-install-iris-web-framework).


Overview
-----------

```go
package main

import (
	"gopkg.in/iris-framework/iris.v6"
	"gopkg.in/iris-framework/iris.v6/adaptors/cors"
	"gopkg.in/iris-framework/iris.v6/adaptors/httprouter"
	"gopkg.in/iris-framework/iris.v6/adaptors/view"
)

func main() {
	// Receives optional iris.Configuration{}, see ./configuration.go
	// for more.
	app := iris.New()

	// Order doesn't matter,
	// You can split it to different .Adapt calls.
	// See ./adaptors folder for more.
	app.Adapt(
		// adapt a logger which prints all errors to the os.Stdout
		iris.DevLogger(),
		// adapt the adaptors/httprouter or adaptors/gorillamux
		httprouter.New(),
		// 5 template engines are supported out-of-the-box:
		//
		// - standard html/template
		// - amber
		// - django
		// - handlebars
		// - pug(jade)
		//
		// Use the html standard engine for all files inside "./views" folder with extension ".html"
		view.HTML("./views", ".html"),
		// Cors wrapper to the entire application, allow all origins.
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}))

	// http://localhost:6300
	// Method: "GET"
	// Render ./views/index.html
	app.Get("/", func(ctx *iris.Context) {
		ctx.Render("index.html", iris.Map{"Title": "Page Title"}, iris.RenderOptions{"gzip": true})
	})

	// Group routes, optionally: share middleware, template layout and custom http errors.
	userAPI := app.Party("/users", userAPIMiddleware).
		Layout("layouts/userLayout.html")
	{
		// Fire userNotFoundHandler when Not Found
		// inside http://localhost:6300/users/*anything
		userAPI.OnError(404, userNotFoundHandler)

		// http://localhost:6300/users
		// Method: "GET"
		userAPI.Get("/", getAllHandler)

		// http://localhost:6300/users/42
		// Method: "GET"
		userAPI.Get("/:id", getByIDHandler)

		// http://localhost:6300/users
		// Method: "POST"
		userAPI.Post("/", saveUserHandler)
	}

	// Start the server at 127.0.0.1:6300
	app.Listen(":6300")
}

func userAPIMiddleware(ctx *iris.Context) {
	// your code here...
	println("Request: " + ctx.Path())
	ctx.Next() // go to the next handler(s)
}

func userNotFoundHandler(ctx *iris.Context) {
	// your code here...
	ctx.HTML(iris.StatusNotFound, "<h1> User page not found </h1>")
}

func getAllHandler(ctx *iris.Context) {
	// your code here...
}

func getByIDHandler(ctx *iris.Context) {
	// take the :id from the path, parse to integer
	// and set it to the new userID local variable.
	userID, _ := ctx.ParamInt("id")

	// userRepo, imaginary database service <- your only job.
	user := userRepo.GetByID(userID)

	// send back a response to the client,
	// .JSON: content type as application/json; charset="utf-8"
	// iris.StatusOK: with 200 http status code.
	//
	// send user as it is or make use of any json valid golang type,
	// like the iris.Map{"username" : user.Username}.
	ctx.JSON(iris.StatusOK, user)
}

func saveUserHandler(ctx *iris.Context) {
	// your code here...
}
```

### Reload on source code changes

```sh
$ go get -u github.com/kataras/rizla
$ cd $GOPATH/src/mywebapp
$ rizla main.go
```

### Reload templates on each incoming request

```go
app.Adapt(view.HTML("./views", ".html").Reload(true))
```


FAQ & Documentation
-----------

 <a href="https://github.com/iris-framework/iris/tree/v6/_examples#table-of-contents"><img align="right" width="125" src="https://raw.githubusercontent.com/iris-contrib/website/gh-pages/assets/book/cover_4.jpg"></a>

1. [Getting Started with Go+Iris](http://gopherbook.iris-go.com)

2. Official small but practical [examples](https://github.com/iris-framework/iris/tree/v6/_examples#table-of-contents)

3. Navigate through [community examples](https://github.com/iris-contrib/examples) too

4. [Creating A URL Shortener Service Using Go, Iris, and Bolt](https://medium.com/@kataras/a-url-shortener-service-using-go-iris-and-bolt-4182f0b00ae7)

5. [Godocs](https://godoc.org/gopkg.in/iris-framework/iris.v6) for deep documentation

6. [HISTORY.md](https://github.com//iris-framework/iris/tree/v6/HISTORY.md) is your best friend, version migrations are released there


I'll be glad to talk with you about **your awesome feature requests**, 
open a new [discussion](http://support.iris-go.com), you will be heard!


Third Party Middleware
------------

Iris has its own middleware form of `func(ctx *iris.Context)` but it's also compatible with all `net/http` middleware forms using [iris.ToHandler](https://github.com/iris-contrib/middleware/blob/master/cors/cors.go#L33), i.e Negroni's middleware form of `func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)`.

Here is a small list of Iris compatible middleware, I'm sure you can find more:

| Middleware | Author | Description |
| -----------|--------|-------------|
| [binding](https://github.com/mholt/binding) | [Matt Holt](https://github.com/mholt) | Data binding from HTTP requests into structs |
| [cloudwatch](https://github.com/cvillecsteele/negroni-cloudwatch) | [Colin Steele](https://github.com/cvillecsteele) | AWS cloudwatch metrics middleware |
| [csp](https://github.com/awakenetworks/csp) | [Awake Networks](https://github.com/awakenetworks) | [Content Security Policy](https://www.w3.org/TR/CSP2/) (CSP) support |
| [delay](https://github.com/jeffbmartinez/delay) | [Jeff Martinez](https://github.com/jeffbmartinez) | Add delays/latency to endpoints. Useful when testing effects of high latency |
| [New Relic Go Agent](https://github.com/yadvendar/negroni-newrelic-go-agent) | [Yadvendar Champawat](https://github.com/yadvendar) | Official [New Relic Go Agent](https://github.com/newrelic/go-agent) (currently in beta)  |
| [gorelic](https://github.com/jingweno/negroni-gorelic) | [Jingwen Owen Ou](https://github.com/jingweno) | New Relic agent for Go runtime |
| [JWT Middleware](https://github.com/auth0/go-jwt-middleware) | [Auth0](https://github.com/auth0) | Middleware checks for a JWT on the `Authorization` header on incoming requests and decodes it|
| [logrus](https://github.com/meatballhat/negroni-logrus) | [Dan Buch](https://github.com/meatballhat) | Logrus-based logger |
| [onthefly](https://github.com/xyproto/onthefly) | [Alexander Rødseth](https://github.com/xyproto) | Generate TinySVG, HTML and CSS on the fly |
| [permissions2](https://github.com/xyproto/permissions2) | [Alexander Rødseth](https://github.com/xyproto) | Cookies, users and permissions |
| [prometheus](https://github.com/zbindenren/negroni-prometheus) | [Rene Zbinden](https://github.com/zbindenren) | Easily create metrics endpoint for the [prometheus](http://prometheus.io) instrumentation tool |
| [render](https://github.com/unrolled/render) | [Cory Jacobsen](https://github.com/unrolled) | Render JSON, XML and HTML templates |
| [RestGate](https://github.com/pjebs/restgate) | [Prasanga Siripala](https://github.com/pjebs) | Secure authentication for REST API endpoints |
| [secure](https://github.com/unrolled/secure) | [Cory Jacobsen](https://github.com/unrolled) | Middleware that implements a few quick security wins |
| [stats](https://github.com/thoas/stats) | [Florent Messa](https://github.com/thoas) | Store information about your web application (response time, etc.) |
| [VanGoH](https://github.com/auroratechnologies/vangoh) | [Taylor Wrobel](https://github.com/twrobel3) | Configurable [AWS-Style](http://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html) HMAC authentication middleware |
| [xrequestid](https://github.com/pilu/xrequestid) | [Andrea Franz](https://github.com/pilu) | Middleware that assigns a random X-Request-Id header to each request |
| [digits](https://github.com/bamarni/digits) | [Bilal Amarni](https://github.com/bamarni) | Middleware that handles [Twitter Digits](https://get.digits.com/) authentication |

Feel free to put up a [PR](https://github.com/iris-contrib/middleware) your middleware!

<!-- 

FAQ
-----------

<p>



</p>

> Q: OK Iris is really fast, but my current website does not need that performance at the moment, are there other reasons to move into Iris?


Iris is fully vendored. That means it is independent of any API changes in the used libraries and **will work seamlessly in the future**!

The size of the executable file is a critical part of the Application Deployment Process.

Two very simple identical applications, the first was written with `iris` and the second with a simple golang router.

 - _iris_ had `8.505 KB` overall file size
 - _gin_ had `9.029 KB` overall file size
 - _net/http_ had produced an executable file with `5.380 KB` size.


> Iris has built'n support for the most of the features that you will use to craft your perfect web application, while the golang router(gin & httprouter alone) doesn't. Imagine what would happened if the simple app we created would use `sessions`, `websockets`, `view engine`... I tested that too, `gin` and `net/http` had produced the x3 of their original size, **while `iris application`' overall executable filesize remained stable**!


**Applications that are written using Iris produce smaller file size even if they use more features** than a simple router library!


Iris always follows the latest trends and best practices. Iris is the **Secret To Staying One Step Ahead of Your Competition**.


Iris is **a high-performance** tool, but it doesn't stops there. Performance depends on your application too, **Iris helps you to make the right choices** on every step.

**Familiar** and easy **API**. Sinatra-like REST API.

Contains examples and documentation for all its features.

Iris is a `low-level access` web framework, you always know what you're doing.

You'll **never miss a thing** from `net/http`, but if you do on some point, no problem because Iris is fully compatible with stdlib, you still have access to `http.ResponseWriter` and `http.Request`, you can adapt any third-party middleware of form `func(http.ResponseWriter, *http.Request, next http.HandlerFunc)` as well.

Iris is a community-driven project, **you suggest and I code**.

Unlike other repositories, this one is **very active**. When you post an issue, you get an answer at the next couple of minutes(hours at the worst). If you find a bug, **I am obliged to fix** that on the same day.


Click the below animation to see by your self what people, like you, say about Iris.

<a href="https://www.youtube.com/watch?v=jGx0LkuUs4A">
<img src="https://github.com/iris-contrib/website/raw/gh-pages/assets/gif_link_to_yt2.gif" alt="What people say" />
</a>



> Q: Why no `serverless`?

New web developers are so enthusiastic about the idea of `serverless` and `AWS`. Most of the experienced developers we already know that we shouldn't use these things for our critical parts of our application.


**`Serverless and AWS` Are Wonderful—Until They Go Wrong.**  There was a flash-point (at 28 February of 2017) where the 'internet was offline' and most of the sites, including isitdownrightnow.com,  were down or operated very slow! Why? Because of `serverless` and `AmazonS3`.
Please think twice before moving your code into `serverless`, **instead, use web frameworks that are created for servers that you control**, i.e Iris.

Proof of concept:

- [Washington Post](https://www.washingtonpost.com/news/the-switch/wp/2017/02/28/why-a-whole-slew-of-websites-are-suddenly-down-or-working-slowly)
- [CNN](http://money.cnn.com/2017/02/28/technology/amazon-web-services-outages/index.html)
- [CNET](https://www.cnet.com/news/no-the-internet-is-not-broken-amazon-web-services-is-just-having-issues/?ftag=COS-05-10-aa0a&linkId=34980800)
- [MIT Technology Review](https://www.technologyreview.com/s/603738/centralized-web-services-are-wonderful-until-they-go-wrong/?_ga=1.82562070.1263144274.1488319022)
- [GolangNews](https://golangnews.com/stories/1835-serverless-is-wonderfuluntil-they-go-wrong.)





Explore [these questions](http://support.iris-go.com/) and join to our [community chat][Chat]!

-->

Testing
------------

The `httptest` package is a simple Iris helper for the httpexpect, a new library for End-to-end HTTP and REST API testing for Go.

You can find tests by navigating to the source code,
i.e:

- [context_test.go](https://github.com/iris-framework/iris/blob/v6/context_test.go)
- [handler_test.go](https://github.com/iris-framework/iris/blob/v6/handler_test.go)
- [policy_gorillamux_test.go](https://github.com/iris-framework/iris/blob/v6/policy_gorillamux_test.go)
- [policy_httprouter_test.go](https://github.com/iris-framework/iris/blob/v6/policy_httprouter_test.go)
- [policy_nativerouter_test.go](https://github.com/iris-framework/iris/blob/v6/policy_nativerouter_test.go)
- [policy_routerwrapper_test.go](https://github.com/iris-framework/iris/blob/v6/policy_routerwrapper_test.go)
- [policy_sessions_test.go](https://github.com/iris-framework/iris/blob/v6/policy_sessions_test.go)
- [response_writer_test.go](https://github.com/iris-framework/iris/blob/v6/response_writer_test.go)
- [route_test.go](https://github.com/iris-framework/iris/blob/v6/route_test.go)
- [status_test.go](https://github.com/iris-framework/iris/blob/v6/status_test.go)
- [transaction_test.go](https://github.com/iris-framework/iris/blob/v6/transaction_test.go)
- [serializer_test.go](https://github.com/iris-framework/iris/blob/v6/serializer_test.go)

A simple test is located to [./_examples/advanced/httptest/main_test.go](https://github.com/iris-framework/iris/blob/v6/_examples/advanced/httptest/main_test.go)

Philosophy
------------

The Iris philosophy is to provide robust tooling for HTTP, making it a great solution for single page applications, web sites, hybrids, or public HTTP APIs. Keep note that, today, iris is faster than nginx itself.

Iris does not force you to use any specific ORM or template engine. Iris is routerless which means you can adapt any router you like, [httprouter](https://github.com/iris-framework/iris/blob/v6/_examples/beginner/routes-using-httprouter/main.go) is the fastest, [gorillamux](https://github.com/iris-framework/iris/blob/v6/_examples/beginner/routes-using-gorillamux/main.go) has more features. With support for the most used template engines (5), you can quickly craft the perfect application.

License
------------

Unless otherwise noted, the source files are distributed
under the MIT License found in the [LICENSE file](LICENSE).

Note that some optional components that you may use with Iris requires
different license agreements.
