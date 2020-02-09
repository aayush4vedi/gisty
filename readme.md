# Gisty
A naive clone of github's [gists](https://gist.github.com/), created to act as standard guidelines when developing any Go app.

## Usage
* To run`$ go run cmd/web/* `
* To check code at any step, see the attached files(@Code).It has only those files which were updated, along with the comments for what, why & Huhs.This is also acheivable by matching the commits(starting from section #2.3)



## 1. Basic Setup

### 1.1 Basic Web App
 * @What: Writing the 3 essential components:
    * #1 **Handler**: similar to Controllers(for MVC guys)
        * responsible for executing your application logic and for writing HTTP response headers and bodies. 
    * #2 **Servemux**: i.e. Router(for MVC guys)
        * stores a mapping between urls & the corresponding handlers.
    * #3 **Web Server**: to listen for incoming requests 
        * Go lets you make it as part of your application itself(unlike external 3rd-party server like Nginx or Apache.)
* @Code: [main.go](https://github.com/aayush4vedi/gisty/blob/48aa717e7ecb487f7b48e0aa6a21e00798331c93/main.go) 
* Notes:
    * Go’s **servemux** treats the URL pattern `/` like a catch-all(as it's a Subtree path patterns).i.e. even though there's single url declared rn, every url will get redirecter here.Try hitting `/nonfuckingexist` ,it will get redirected to `/`
    * Go’s **servemux** supports 2 different types of URL patterns: `fixed paths` and `subtree paths`. 
       * Fixed path patterns: are only matched when the request URL path **exactly** matches the fixed path.
       * Subtree path patterns: are matched whenever the start of a request URL path matches the subtree path.
          * For understanding, you can think of them as they have a wildcard at the end, like `"/**"` or `"/static/**"`
    * A faster+smaller_code but wrong approach for Servemux by using `HandleFunc`: 
      * `http.HandleFunc("/", home)` :internally registers routes using `DefaultServeMux`
      * What's `DefaultServeMux`?: it’s just regular servemux, but which is initialized by default and stored in a `net/http` global variable. 
      * Eevn though this Why not to use it? => Security
        * Because `DefaultServeMux` is a global variable, any package can access it and register a route — including any third-party packages that your application imports. If one of those third-party packages is compromised, they could use `DefaultServeMux` to expose a malicious handler to the web. 
### 1.2 Add more routes
  * |Method| url  |  handler | action   |
    |--| -----| -------- | ---------|
    |Any| /    | home     | Display home page|
    |Any| /gist?id=1| showGist | Display a specific gist|
    |Post| /gist/create | createGist | Create new gist |
    |ANY  |/static/|http.FileServer |Serve a specific static file
  * @Code: [main.go](https://github.com/aayush4vedi/gisty/blob/950013e3f6e6caa46d1764618458de2c8320d9e0/main.go)
  * @Notes:
    * In **servemux** longer URL patterns always take precedence over shorter ones. So, if a servemux contains multiple patterns.nice side-effect => you can register patterns in any order and it won’t change how the servemux behaves. 
    * 301:Permanent Redirect
      * If the request path contains any `.` or `..` , it will automatically redirect the user to an equivalent clean URL. E.g.: `/foo/bar/..//baz` => 301 Permanent Redirect => `/foo/baz` 
      * Subtree path redirection: `/foo/bar/` => 301 Permanent Redirect => `/foo/bar` 
    * Host Name matching: any host-specific patterns will be checked before the regular urls. e.g: `mux.HandleFunc("foo.example.org/", fooHandler) ` will be matched first

### 1.3 Customize Headers
 * @What: Make handler `createGist` send `405` (method not allowed) HTTP status code unless the request method is POST
   * Use: `r.Method != "POST"` for checking method & `w.WriteHeader()` to send needed status code.
 * @Code: [main.go](https://github.com/aayush4vedi/gisty/blob/844036144582c45f480751859725740bf90255d2/main.go#L22)>createGist
 * @Notes:
   * ABout `w.WriteHeader()`:
        * It’s only possible to call `w.WriteHeader()` once per response, and after the status code has been written it can’t be changed. If you try to call `w.WriteHeader()` a second time Go will log a warning message. 
        * If you don’t call w.WriteHeader() explicitly, then the first call to `w.Write()` will automatically send a 200 OK status code to the user. So, if you want to send a non-200 status code, you must call `w.WriteHeader()` before any call to `w.Write()`. 
   * About `curl`
        * Sending request from terminal w/o Postman: `$ curl -i -X POST http://localhost:4000/snippet/create `
 * #Another(& better) approach:
   * Use: `w.Header().Set()` method to add a new header to the response header map & `http.Error()` shortcut
   * @Code: [main.go](https://github.com/aayush4vedi/gisty/blob/c4b4a899aa9eef020bbfae187ddab97513356e8b/main.go#L23)>createGist
        * `w.Header().Set("Allow", "POST") `<br>
          `http.Error(w, "Method Not Allowed", 405) `
   * @Notes: Manipulating the Header Map :
        * Similar to `w.Header().Set()` there are also `Add()`,`Del()` and `Get()` methods that you can use to read and manipulate the header map:
            * `Get`: Retrieve the first value for the "Cache-Control" header<br>`w.Header().Get("Cache-Control")`
            * `Set` a new/Overwrite cache-control header: <br>`w.Header().Set("Content-Type", "application/json")` 
            * `Add()` method appends a new "Cache-Control" header & can be called multiple times:<br>`w.Header().Add("Cache-Control", "public") `<br>`w.Header().Add("Cache-Control", "max-age=31536000") `
            * `Delete` all values for the "Cache-Control" header: <br>`w.Header().Del("Cache-Control")`
 
### 1.4 Adding URL query parameter(for string)
 * @What: update the `showGist` handler so that it accepts an id query string parameter using `r.URL.Query().Get()` method
 * @Code: [main.go](https://github.com/aayush4vedi/gisty/blob/2684831ac7c4447bc9e102006a9f1c734c56f28e/main.go#L20)>showGist
 * @Notes: `io.Writer` == `http.ResponseWriter`
     * `io.Writer` type is an interface, and the `http.ResponseWriter` object satisfies the interface because it has a `w.Write()` method
### 1.5 Project Structure
* Read: [Setting up a Go project](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
* Structure:(Clear separation of GO & non-Go assets)
  * **cmd** : application-specific code(web, cli)
    * cmd/web : contains `main.go` & `handlers.go` or this project
  * **pkg** : reusable non-application-specific code(validation helpers, SQL models)
  * **ui**       
    * ui/html : html templates
    * ui/static : static files(css, imgs)
 
### 1.6 Add HTML Template & Inheritence
* @What: Add a template file & use in headers using `html/template` pkg
    * In handler:
        * Read template file:`ts, err := template.ParseFiles("./ui/html/home.page.tmpl")`
        * Write template content as response body.(`nil` is in place for dynamic data part):<br>`err = ts.Execute(w, nil)`
    * Smart templating(inheritence):
        * Define custom templates: `title`, `base`, `body` & import them in your templates
        * `{{define "base"}}...{{end}}` to define a named template called `base`
            * You can define multiple templates inside a single files
        * `{{template "title" .}}` to invoke other a named templates called `title` 
        * The dot at the end of the `{{template "title" .}}` action represents any dynamic data that you want to pass to the invoked template(later)
    * Use the `template.ParseFiles()` function to read the files and store the templates in a template set(as variadic param...)
* @Code: [home.page.tmpl](https://github.com/aayush4vedi/gisty/blob/b2b1ffadd001ac638ec6be420adbcbc0183367c4/ui/html/home.page.tmpl), [base.layout.tmpl](https://github.com/aayush4vedi/gisty/blob/b2b1ffadd001ac638ec6be420adbcbc0183367c4/ui/html/base.layout.tmpl) [handlers.go](https://github.com/aayush4vedi/gisty/blob/b2b1ffadd001ac638ec6be420adbcbc0183367c4/cmd/web/handlers.go#L25)
* @Notes:
  * On naming template files: `<name>.<role>.tmpl`, where <role> is either `page`, `partial` or `layout`.
    * Being able to distinguish the role of the template will help us when it comes to creating a cache of templates(later)
  * It doesn’t matter so much what naming convention or file extension you use. Go is flexible about this. 

### 1.7 Serving Static Files
* @What: Add a new route so that all requests which begin with `/static/`(subtree path pattern) are handled using this.Using:
    * `net/http` package ships with a built-in `http.FileServer` handler which you can use to serve files over HTTP from a specific directory.
    * To create a new http.FileServer handler, we need to use the `http.FileServer()` function like this: <br>
    `fileServer := http.FileServer(http.Dir("./ui/static")) `
* @Code: [main.go](https://github.com/aayush4vedi/gisty/blob/b135cc9d77630f09377a19f9ff8c76abcc83d9e6/cmd/web/main.go#L15), [static files](https://github.com/aayush4vedi/gisty/tree/b135cc9d77630f09377a19f9ff8c76abcc83d9e6/ui)
* @Notes: Features of Go’s file server:
    * It sanitizes all request paths by running them through the `path.Clean()` function before searching for a file. This removes any  `.` and `..` elements from the URL path, which helps to stop directory traversal attacks.
    * But be aware:`http.ServeFile()` does not automatically sanitize the file path. If you’re constructing a file path from untrusted user input, to avoid directory traversal attacks you must sanitize the input with `filepath.Clean()` before using it.
    * The `Last-Modified` and `If-Modified-Since` headers are transparently supported. If a file hasn’t changed since the user last  requested it, then `http.FileServer` will send a `304 Not Modified` status code instead of the file itself. This helps reduce latency and processing overhead for both the client and server. 
    * The `Content-Type` is automatically set from the file extension using the `mime.TypeByExtension()` function. You can add your own custom extensions and content types using the `mime.AddExtensionType()` function if necessary. 
* All incoming HTTP requests are served in their **own goroutine**. For busy servers, this means it’s very likely that the code in or called by your handlers will berunning concurrently. **While this helps make Go blazingly fast**, the downside is that you need to be aware of (and protect against) race conditions when accessing shared resources from your handlers.

## 2. Configuration & Error Handling
### 2.1 Better Logging
* @What: 
    * #1: use the `log.New()` function to create two new custom loggers
        * Declaring logs:<br>`infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) `<br>`errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)`
        * Writing logs using them `infoLog.Printf("Starting server on %s", *addr) `<br>  `errorLog.Fatal(err) `
    * #2: http.Server Error Log
        * Declaring serv log struct: <br>
        ```go
        srv := &http.Server{ 
            Addr:     *addr, 
            ErrorLog: errorLog, 
            Handler:  mux, 
        } 
        ```
        * Call the ListenAndServe() method on our new http.Server struct: `err := srv.ListenAndServe()`
* @code: [main.go](https://github.com/aayush4vedi/gisty/blob/c09bef2b40e706db5a3520231f2568bbd41f4e03/cmd/web/main.go#L14)

### 2.2 Dependency injection(Methodising the app)
>>"Make everything into methods first"
* @What: Define an application struct to hold the application-wide dependencies for the app & make it available for handler
    * Defining the struct <br>
    ```go
    type application struct { 
        errorLog *log.Logger 
        infoLog  *log.Logger 
    }
    ```
    * Import them in `handler.go` to make all the functions as methods
    * Swap the route declarations to use the application struct's methods as handler functio
* Code: [main.go](https://github.com/aayush4vedi/gisty/blob/a3a837961ca64776e318f9253010600844125c9a/cmd/web/main.go) , [handlers.go](https://github.com/aayush4vedi/gisty/blob/a3a837961ca64776e318f9253010600844125c9a/cmd/web/handlers.go)

* @Notes: **On Closure**
    * This pattern for inject dependencies won’t work if your handlers are spread across multiple packages. 
    * **Better approacb**: to create a `config` package exporting an `Application` struct and have your handler functions close over this to form a closure.
    * ```go
        func main() { 
            app := &config.Application{ 
                ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
             }...
        ```
    
### 2.3 Centralized Error Handling
* @What: Move the error handling code into helper methods. This will help separate our concerns and stop repitition.
    * In `cmd/web/errors.go`:
        * `serverError` : sends a generic 500 Internal Server Error response to the user. 
            * Uses the `debug.Stack()` function to get a stack trace for the current goroutine and append it to the log message.
        * `clientError` : helper sends a specific status code and corresponding description
            * Uses the  http.StatusText()` function to automatically generate a human-friendly text representation of a given HTTP status code
        * `notFound` : sends a 404 Not Found 
            *  Uses the constant `http.StatusNotFound` instead of writing 404. 
    * Update errors in handler

* @Code: [errors.go](https://github.com/aayush4vedi/gisty/blob/d83906e46feeb2edbcc670f8a0d86a13e9d6b769/cmd/web/errors.go) , [handlers.go](https://github.com/aayush4vedi/gisty/blob/d83906e46feeb2edbcc670f8a0d86a13e9d6b769/cmd/web/handlers.go)

### 2.4 Isolating the Application Routes 
* @What: Move routes to `routes.go`
* @Code: [routes.go](), [main.go]()


## 3. DB-driven Responses


## 4. Dynamic HTML Templates


## 5. Middlewares


## 6. RESTful Routing


## 7. Processing Forms


## 8. Stateful HTTP


## 9. Security


## 10. Auth


## 11. Using Request Context


## 12. Testing







