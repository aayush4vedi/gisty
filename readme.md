# Gisty
A naive clone of github's [gists](https://gist.github.com/), created to act as standard guidelines when developing any Go app.

## Usage
* To run`$ go run cmd/web/* `
* To check code at any step, see the attached files(@Code).It has only those files which were updated, along with the comments for what, why & Huhs.This is also acheivable by matching the commits id+1(starting from section #2.3)

## Demo
* To play live, click [here](https://aayush4vedi.github.io/gisty)

## Index
Taken from [Book](https://drive.google.com/file/d/1tfDmBhzW9BZOP5RicCwhSTZ2i59IiQLl/view?usp=sharing) Let's Go by Alex Edwards
1. [Basic Setup](https://github.com/aayush4vedi/gisty#1-basic-setup)
    * 1.1 [Basic Web App](https://github.com/aayush4vedi/gisty#11-basic-web-app)
    * 1.2 [Add more routes](https://github.com/aayush4vedi/gisty#12-add-more-routes)
    * 1.3 [Customize Headers](https://github.com/aayush4vedi/gisty#13-customize-headers)
    * 1.4 [Adding URL query parameter(for string)](https://github.com/aayush4vedi/gisty#14-adding-url-query-parameterfor-string)
    * 1.6 [Add HTML Template & Inheritence](https://github.com/aayush4vedi/gisty#16-add-html-template--inheritence)
    * 1.7 [Serving Static Files](https://github.com/aayush4vedi/gisty#17-serving-static-files)
2. [Configuration & Error Handling](https://github.com/aayush4vedi/gisty#2-configuration--error-handling)
    * 2.1 [Better Logging](https://github.com/aayush4vedi/gisty#21-better-logging)
    * 2.2 [Dependency injection(Methodising the app)](https://github.com/aayush4vedi/gisty#22-dependency-injectionmethodising-the-app)
    * 2.3 [Centralized Error Handling](https://github.com/aayush4vedi/gisty#23-centralized-error-handling)
    * 2.4 [Isolating the Application Routes ](https://github.com/aayush4vedi/gisty#24-isolating-the-application-routes)
3. [DB-driven Responses](https://github.com/aayush4vedi/gisty#3-db-driven-responses)
    * 3.1 [Setting up MySQL](https://github.com/aayush4vedi/gisty#31-setting-up-mysql)
    * 3.2 [Installing a DB Driver](https://github.com/aayush4vedi/gisty#32-installing-a-db-driver)
    * 3.3 [Creating a Database Connection Pool ](https://github.com/aayush4vedi/gisty#33-creating-a-database-connection-pool)
    * 3.4 [Designing a Database Model](https://github.com/aayush4vedi/gisty#34-designing-a-database-model)
    * 3.5 [Executing SQL Statements(`C`)](https://github.com/aayush4vedi/gisty#35-executing-sql-statementsc)
    * 3.6 [Single-record SQL Queries(`R`)](https://github.com/aayush4vedi/gisty#36-single-record-sql-queriesr)
    * 3.7 [Multiple-record SQL Queries](https://github.com/aayush4vedi/gisty#37-multiple-record-sql-queries)
    * 3.8. [Transactions, #Connections & Prepared Statement](https://github.com/aayush4vedi/gisty#38-transactions-connections--prepared-statement)
 4. [Dynamic HTML Templates](https://github.com/aayush4vedi/gisty#4-dynamic-html-templates)
    * 4.1 [Displaying Dynamic Data](https://github.com/aayush4vedi/gisty#41-displaying-dynamic-data)
    * 4.2 [Template Actions and Functions](https://github.com/aayush4vedi/gisty#42-template-actions-and-functions)
    * 4.3 [Caching Templates](https://github.com/aayush4vedi/gisty#43-caching-templates)
    * 4.4 [Catching Runtime Errors](https://github.com/aayush4vedi/gisty#44-catching-runtime-errors)
    * 4.5 [Common Dynamic Data](https://github.com/aayush4vedi/gisty#45-common-dynamic-data)
    * 4.6 [Custom Template Functions](https://github.com/aayush4vedi/gisty#46-custom-template-functions)
5. [Middlewares](https://github.com/aayush4vedi/gisty#5-middlewares)
    * 5.1 [About Middleware](https://github.com/aayush4vedi/gisty#51-about-middleware)
    * 5.2 [Creating beforeMux-middleware: Set-Security-Headers](https://github.com/aayush4vedi/gisty#52-creating-beforemux-middleware-set-security-headers)
    * 5.3 [Creating beforeMux-middleware: Request-Logger](https://github.com/aayush4vedi/gisty#53-creating-beforemux-middleware-request-logger)
    * 5.4 [Create Panic Recovery Middleware](https://github.com/aayush4vedi/gisty#54-create-panic-recovery-middleware)
    * 5.5 [Composable Middleware Chains- pkg 'alice'](https://github.com/aayush4vedi/gisty#55-composable-middleware-chains--pkg-alice)
6. [RESTful Routing](https://github.com/aayush4vedi/gisty#6-restful-routing)
    * 6.1 [About Routers](https://github.com/aayush4vedi/gisty#61-about-routers)
    * 6.2 [Implement RESTful Routes in app- Using `Pat`](https://github.com/aayush4vedi/gisty#62-implement-restful-routes-in-app--using-pat) -TODO: use `gin`
7. [Processing Forms](https://github.com/aayush4vedi/gisty#7-processing-forms)
    * 7.1 [Setting Up a Form & Parsing Data to It](https://github.com/aayush4vedi/gisty#71-setting-up-a-form--parsing-data-to-it)
    * 7.2 [Data Validation](https://github.com/aayush4vedi/gisty#72-data-validation)
    * 7.3 [Scaling Data Validation(Refactor)](https://github.com/aayush4vedi/gisty#73-scaling-data-validationrefactor)
8. [Stateful HTTP](https://github.com/aayush4vedi/gisty#8-stateful-http)
    * 8.1 [About Session Managers](https://github.com/aayush4vedi/gisty#81-about-session-managers)
    * 8.2 [Working with Session Data](https://github.com/aayush4vedi/gisty#82-working-with-session-data)
9. [Security](https://github.com/aayush4vedi/gisty#9-security)
    * 9.1 [TLS Certificate](https://github.com/aayush4vedi/gisty#91-tls-certificate)
    * 9.2 [Timeouts](https://github.com/aayush4vedi/gisty#92-timeouts)
10. [Auth](https://github.com/aayush4vedi/gisty#10-auth)
    * 10.1. [Routes Setup](https://github.com/aayush4vedi/gisty#101-routes-setup)
    * 10.2. [Creating a Users Model](https://github.com/aayush4vedi/gisty#102-creating-a-users-model)
    * 10.3. [User Signup and Password Encryption](https://github.com/aayush4vedi/gisty#103-user-signup-and-password-encryption)
    * 10.4. [User Login](https://github.com/aayush4vedi/gisty#104-user-login)
    * 10.5. [User Logout](https://github.com/aayush4vedi/gisty#105-user-logout)
    * 10.6. [User Authorization](https://github.com/aayush4vedi/gisty#106-user-authorization)
    * 10.7. [CSRF Protection](https://github.com/aayush4vedi/gisty#107-csrf-protection)
11. [Using Request Context](https://github.com/aayush4vedi/gisty#11-using-request-context)
12. [Testing](https://github.com/aayush4vedi/gisty#12-testing)
    * 12.1 [Testing HTTP Handlers](https://github.com/aayush4vedi/gisty#121-testing-http-handlers)
    * 12.2 [Testing Middleware](https://github.com/aayush4vedi/gisty#122-testing-middleware)
    * 12.3 [End-To-End Testing](https://github.com/aayush4vedi/gisty#123-end-to-end-testing)
    * 12.4 [Mocking Dependencies](https://github.com/aayush4vedi/gisty#124-mocking-dependencies)
    * 12.5 [Testing HTML Forms](https://github.com/aayush4vedi/gisty#124-mocking-dependencies)
    * 12.6 [Integration Testing](https://github.com/aayush4vedi/gisty#124-mocking-dependencies)
    * 12.7 [Profiling Test Coverage](https://github.com/aayush4vedi/gisty#124-mocking-dependencies)

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
        * Sending request from terminal w/o Postman: `$ curl -i -X POST http://localhost:4000/gist/create `
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
* @Code: [routes.go](https://github.com/aayush4vedi/gisty/blob/9262412f6fea9a52b047a148af80e9b9e21a9c45/cmd/web/routes.go), [main.go](https://github.com/aayush4vedi/gisty/blob/9262412f6fea9a52b047a148af80e9b9e21a9c45/cmd/web/main.go#L30)


## 3. DB-driven Responses

### 3.1 Setting up MySQL
* Setup a DB in mysql using UTF8 encoding: `CREATE DATABASE gisty CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci; `
* Switch to it: `USE gisty; `
* create table: 
    ```sql
    CREATE TABLE gists ( 
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, 
    title VARCHAR(100) NOT NULL, 
    content TEXT NOT NULL, 
    created DATETIME NOT NULL, 
    expires DATETIME NOT NULL 
    );
    ```
* Create index on the created column: `CREATE INDEX idx_gists_created ON gists(created);`
* Seed table with dummmy data: 
    ```sql
    INSERT INTO gists (title, content, created, expires) VALUES ( 
    'An old silent pond', 
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.
    UTC_TIMESTAMP(), 
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY) 
    ); 
    ```
* Display table with entries: `SELECT id, title, expires FROM gists;`
* Delete table(if wish): `DROP TABLE gists;`

### 3.2 Installing a DB Driver
* **DB Driver** acts as a middleman, translating commands between Go and the MySQL database itself. 
* List of all available drivers for Mysql on [go-page](https://github.com/golang/go/wiki/SQLDrivers).
* Used here:  [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql):
    * `go get github.com/go-sql-driver/mysql `
    * This creates a new file `go.sum` which contains the cryptographic checksums representing the content of the required packages.
* @Notes: About verifing the checksums & dependencies
    * Run the `go mod verify`, this will verify that the checksums of the downloaded packages on your machine match the entries in `go.sum`, so you can be confident that they haven’t been altered. 
    * To install the dependencies for the project — run `go mod download`
    * To uninstall a package: `go get github.com/foo/bar@none `

###  3.3 Creating a Database Connection Pool 
* @What: Connect the driver with app(& verify connection) using `sql.Open()` function:
    * db, err := sql.Open(<driver name>, <DSN>) 
    * DSN(Data Source Name): depends on the type [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
    * Update `main.go`
* @Code: [main.go](https://github.com/aayush4vedi/gisty/blob/23b8cf9d983934da702a8c0a4f910c75a22bafd9/cmd/web/main.go#L10)

### 3.4 Designing a Database Model
* @What: Create model & connect with app:
    * make dir `pkg/models` & `pkg/models/sql`
    * Create model `Gist` in `pkg/models/models.go`:
    ```go
    type Gist struct {
        ID      int
        Title   string
        Content string
        Created time.Time
        Expires time.Time
    }
    ```
    * Write CRUD methods for Mysql DB in `pkg/models/mysql/gisty.go` on model `GistModel`
    * Use `GistModel` in `main.go`(~~what's the use of `Gist` then???~~ => It's used when yielding templates with dynamic data, all the fields in template needs to be in this): `gists: &mysql.GistModel{DB: db}`
* @Code: [models.go](https://github.com/aayush4vedi/gisty/blob/2d9cb7d12a195f1e492826e95c189dbc8b35785f/pkg/models/models.go), [gisty.go](https://github.com/aayush4vedi/gisty/blob/2d9cb7d12a195f1e492826e95c189dbc8b35785f/pkg/models/mysql/gisty.go), [main.go](https://github.com/aayush4vedi/gisty/blob/2d9cb7d12a195f1e492826e95c189dbc8b35785f/cmd/web/main.go#L20)
* @Notes: On directory structure
    * the directory structure scales nicely if your project has multiple back ends. For example, if some of your data is held in `Redis` you could put all the models for it in a `pkg/models/redis` package.

### 3.5 Executing SQL Statements(`C`)
* Go provides **3 methods for executing database queries**: 
    * `DB.QueryRow()` is used for `SELECT` queries which *return a single row*.
    * `DB.Query()` is used for `SELECT` queries which *return multiple rows*. 
    * `DB.Exec()` is used for statements which don’t return rows (like `INSERT` and `DELETE`). 
* @What#1: Update the `create` method: i.e. `GisttModel.Insert()` using Go's `DB.Exec()`
    * Write query statement(with `?` as `placeholders`): 
    ```go
    stmt := `INSERT INTO gists (title, content, created, expires) 
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
    ```
    * Execute it: `result, err := m.DB.Exec(stmt, title, content, expires)`
    * Get the ID of entry in table: `id, err := result.LastInsertId()`
    * Return ID,convert from int64 to int:	`return int(id), nil`
* @Code: [gisty.go](https://github.com/aayush4vedi/gisty/blob/da31cd65c658b9f1680ddcf406875bd9778fc2c9/pkg/models/mysql/gisty.go#L14)
* @Notes: On `sql.Result` interface returned by `DB.Exec()`.This provides two methods: 
    * 1.`LastInsertId()` — which returns the integer (an int64) generated by the database in response to a command. 
        * [Not suported by](https://github.com/lib/pq/issues/24) postgresql
    * 2.`RowsAffected()` — which returns the number of rows (as an int64) affected by the statement. 
    * It is perfectly acceptable (and common) to ignore the `sql.Result` return value if you don’t need it:`_, err := m.DB.Exec("INSERT INTO ...", ...) `
* @What#2: Using the Model in Our Handlers 
    * Declare some random data(in vars)
    * Insert: `id, err := app.gists.Insert(title, content, expires)`
    * Redirect: `http.Redirect(w, r, fmt.Sprintf("/gist?id=%d", id), http.StatusSeeOther)`
* @Code: [handlers.go](https://github.com/aayush4vedi/gisty/blob/da31cd65c658b9f1680ddcf406875bd9778fc2c9/cmd/web/handlers.go#L50)
* @Notes:
    * The reason for using placeholder parameters(`?`) to construct our query (rather than string interpolation) is to help avoid SQL injection attacks from any untrusted user-provided input. 
    * Behind the scenes, the DB.Exec() method works in three steps: 
        * 1. It creates a new prepared statement on the database using the provided SQL statement. The database parses and compiles the statement, then stores it ready for execution.
        * 2. In a second separate step, Exec() passes the parameter values to the database. The database then executes the prepared statement using these parameters. Because the parameters are transmitted later, after the statement has been compiled, the  database treats them as pure data. They can’t change the intent of the statement. So long as the original statement is not derived from an untrusted data, injection cannot occur. 
        * 3. It then closes (or deallocates) the prepared statement on the database. 

### 3.6 Single-record SQL Queries(`R`)
* @What#1: update `get` function for single record
    * Write the SQL statement we want to execute:
	```go
    stmt := `SELECT id, title, content, created, expires FROM gisty 
    WHERE expires > UTC_TIMESTAMP() AND id = ?`
    ```
    * Use `QueryRow()` to execute the SQL statement.This returns a pointer to a `sql.Row` object:
	`row := m.DB.QueryRow(stmt, id) `
	
    * Initialize a pointer to a new zeroed Gist struct	`s := &models.Gist{} `
	
	* Use `row.Scan()` to copy the values from each field in `sql.Row `to corr. fields in struct:
	`err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires) `
* @Code: [gisty.go](https://github.com/aayush4vedi/gisty/blob/80d93b96234382e884a3f1fba4c2ebccfec07623/pkg/models/mysql/gisty.go#L38), [handlers.go](https://github.com/aayush4vedi/gisty/blob/80d93b96234382e884a3f1fba4c2ebccfec07623/cmd/web/handlers.go#L35)
    * Shoorter code: 
        * @Code:[gisty.go](https://github.com/aayush4vedi/gisty/blob/baf6c41ba19eb498d1e1f812f7985ae80410d7ee/pkg/models/mysql/gisty.go#L31)

* @What#1: use it in handlers: `s, err := app.gists.Get(id)`

### 3.7 Multiple-record SQL Queries
* @What: Useing `DB.Query(stmt)` fetch all recently created 10 gists
* @Code: [gisty.go](https://github.com/aayush4vedi/gisty/blob/b2654366f2558b9f518a015fd4a6db0723047eb4/pkg/models/mysql/gisty.go#L41) , [handlers.go](https://github.com/aayush4vedi/gisty/blob/b2654366f2558b9f518a015fd4a6db0723047eb4/cmd/web/handlers.go#L16)

### 3.8. Transactions, #Connections & Prepared Statement
* How to maintain **atomicity** (Ans: Use mysql's transactions):
    * The Problem: 
        * The calls to `Exec()`, `Query()` and `QueryRow()` can use any connection from the `sql.DB` pool.Even if you have two calls to `Exec()` immediately next to each other in your code, there is no guarantee that they will use the same database connection. 
        * Sometimes this isn’t acceptable. For instance, if you lock a table with MySQL’s `LOCK TABLES` command you must call `UNLOCK TABLES` on exactly the same connection to avoid a **deadlock**. 
    * The Solution: 
        * To guarantee that the same connection is used you can wrap multiple statements in a transaction using `Begin()` which creates a new transaction obj `sql.Tx`
        * It implements `tx.Rollback()` method in the event of any errors, the transaction ensures that either: 
            * 1.All statements are executed successfully; or
            * 2.No statements are executed and the database remains unchanged. 
    * Demo: [transaction_demo](https://github.com/aayush4vedi/gisty/blob/baf6c41ba19eb498d1e1f812f7985ae80410d7ee/examples/transaction_demo.go)
* Managing **Number of Connections**:
    * The `sql.DB` connection pool is made up of connections which are either idle or open (in use).
    * By default, there is no limit on the maximum number of open connections at one time, but the default maximum number of idle connections in the pool is 2
    * You can set them like this:
        * `db.SetMaxOpenConns(95) `
        * `db.SetMaxIdleConns(5)`
    * Limitation: your database itself probably has a hard limit on the maximum number of connections. E.g.> default limit for MySQL is 151.Beyond that you'd get `"too many connections"` error.
* Prepared Statement:
    *  `Exec()`, `Query()` and `QueryRow()` methods all use prepared statements behind the scenes to help prevent SQL injection  attacks.
    * This might feel rather inefficient because we are creating and recreating the same prepared statements every single time. 
    * Better Approach:  use of the `DB.Prepare()` method to create our own prepared statement once, and reuse that instead.
    * Demo: [prepared_stmt_demo.go](https://github.com/aayush4vedi/gisty/blob/baf6c41ba19eb498d1e1f812f7985ae80410d7ee/examples/prepared_stmt_demo.go)

## 4. Dynamic HTML Templates

### 4.1 Displaying Dynamic Data
* @What#1: Displaying dynamic data in `showGist()` handler function(Single piece of data)
    * Import `html/template` into `handlers.go`
    * Declare a template `show.page.tmpl` with `dots(.)` - as placeholders for dynamic data
        * Because our `models.Gist` struct has a `Title` field, we could yield the gist title by writing `{{.Title}}` in our templates(this is where)
    * Get data from mysql, using `s, err := app.gists.Get(id)`
    * Render template: `err = ts.Execute(w, s)`
* @Code: [show.page.tmpl](https://github.com/aayush4vedi/gisty/blob/0168474fc9429eb8a7c2bacab00349c897e4ba36/ui/html/show.page.tmpl), [handler.go](https://github.com/aayush4vedi/gisty/blob/0168474fc9429eb8a7c2bacab00349c897e4ba36/cmd/web/handlers.go#L35)

* @What#2: Displaying Multiple Pieces of Data
    * Create `template.go` to act as the holding structure for any dynamic data that we want to pass to our HTML template.
    * Update the `showGist` handler to use this new struct when executing our templates
    * In `show.page.tmpl` chain the appropriate field names together like `{{.Gist.Title}}`

* @Code: [template.go](https://github.com/aayush4vedi/gisty/blob/8b4b0573f203d23e8f4192f127eae9b5ac1963f7/cmd/web/templates.go) , [show.page.tmpl](https://github.com/aayush4vedi/gisty/blob/8b4b0573f203d23e8f4192f127eae9b5ac1963f7/ui/html/show.page.tmpl) , [handlers.go](https://github.com/aayush4vedi/gisty/blob/8b4b0573f203d23e8f4192f127eae9b5ac1963f7/cmd/web/handlers.go#L41)

* @Notes: On `html/template`:
    * Issue with multiple data:
        * It allows you to pass in one — and only one — item of dynamic data when rendering a template.
        * How to overcome it for multiple data: wrap your dynamic data in a struct which acts like a single ‘holding structure’ for your data. 
    * Inbuilt **Escaping**:     
        * The pkg automatically escapes any data that is yielded between `{{ }}` tags. This behavior is hugely helpful in avoiding  cross-site scripting (**XSS**) attacks, and is the reason that you should use the `html/template` package instead of the more  eneric `text/template` package that Go also provides. 
        * E.g.: if your dynamic data is: `<span>{{"<script>alert('xss attack')</script>"}}</span>`<br> it would be rendered harmless as: `<span>&lt;script&gt;alert(&#39;xss attack&#39;)&lt;/script&gt;</span> ` 
        * **HTML Comments:** Go simply strips out all HTML comments.
    * **Nested Templates:** It’s really important to note that when you’re invoking one template from another template, dot needs to  be explicitly passed or pipelined to the template being invoked.
        * You do this by including it at the end of each `{{template}}` or `{{block}}` action, like so: 
        * ```go
            {{template "base" .}} 
            {{template "body" .}} 
            {{template "footer" .}} 
            {{block "sidebar" .}}{{end}}
        ```
    * **Calling Methods:**
        * If the object that you’re yielding has methods defined against it, you can call them (so long as they are exported and they return only a single value — or a single value and an error).
        * **E.g.**:
        * For example, if `.Gist.Created` has the underlying type `time.Time` you could render the name of the weekday by calling its `Weekday()` method like so: `<span>{{.Gist.Created.Weekday}}</span> `
        * **Passing Parameters:**  Use the `AddDate()` method to add six months to a time like so: `<span>{{.Gist.Created.AddDate 0 6 0}}</span> `
            * Notice that this is different syntax to calling functions in Go — the parameters are not surrounded by parentheses and are separated by a single space character, not a comma. 

### 4.2 Template Actions and Functions
The `html/template` package provides following:
* **Actions**

|Action | Description|
| ----  | -----------|
|`{{define}}`  |  |
|`{{template}}`  |  |
|`{{block}}`  |  |
|`{{if .Foo}} C1 {{else}} C2 {{end}} `  | If `.Foo` is not empty then render the content C1, else render the content C2 |
|`{{with .Foo}} C1 {{else}} C2 {{end}}`  | If `.Foo` is not empty, then set `.` to the value of `.Foo` and render the content C1, else render the content C2  |
|`{{range .Foo}} C1 {{else}} C2 {{end}}`  | if len(`.Foo`) > 0 then loop over each element, setting `.` to the value of each element and render content C1. If len(`.Foo`) == 0 then render content C2.*Note:* `.Foo` must be an array, slice, map, or channel.  |

* **Functions**

|Function | Description|
| ----  | -----------|
|`{{eq .Foo .Bar}}`  | Yields true if `.Foo` == `.Bar`  |
|`{{ne .Foo .Bar}}`  | Yields true if `.Foo` != `.Bar`  |
|`{{not .Foo}}`  | Yields the boolean negation of `.Foo`  |
|`{{or .Foo .Bar}} `  | Yields `.Foo` if `.Foo` is not empty; else yields `.Bar`  |
|`{{index .Foo i}} `  | Yields value of `.Foo` @ index `i`.`.Foo` must be a map, slice or array |
|`{{printf "%s-%s" .Foo .Bar}}`  | Yields a string containing the `.Foo` & `.Bar` values.Same as `fmt.Sprintf()` |
|`{{len .Foo}}`  | Yields the length of `.Foo` as an integer |
|`{{$bar := len .Foo}} `  | Assign the length of `.Foo` to the template variable `$bar`  |

@Code: [handlers.go](https://github.com/aayush4vedi/gisty/blob/e69196a9853ee8816f5853e8b3363bb33eeca48e/cmd/web/handlers.go#L25) , [templates.go](https://github.com/aayush4vedi/gisty/blob/e69196a9853ee8816f5853e8b3363bb33eeca48e/cmd/web/templates.go) , [home.page.tmpl](https://github.com/aayush4vedi/gisty/blob/e69196a9853ee8816f5853e8b3363bb33eeca48e/ui/html/home.page.tmpl) , [show.page.tmpl](https://github.com/aayush4vedi/gisty/blob/e69196a9853ee8816f5853e8b3363bb33eeca48e/ui/html/show.page.tmpl)


### 4.3 Caching Templates

* @What#1: On each load,template files are read from disk.Speeded it up by **caching the templates in memory**. 
    * Create an in-memory map with the type `map[string]*template.Template` to cache the templates
    * initialize this cache in the `main()` function & make it available to `handlers` as a dependency via the `App` struct
* @Code: [templates.go](https://github.com/aayush4vedi/gisty/blob/6ebcc843fad101b39ceee1db9bb0e45fb66598c4/cmd/web/templates.go#L17) , [main.go](https://github.com/aayush4vedi/gisty/blob/6ebcc843fad101b39ceee1db9bb0e45fb66598c4/cmd/web/main.go#L37)

* @What#2: Create helper function to reduce code duplicationin the `home` and `showGist` handlers.
    * Create a `render` method in `handlers.go` which takes in `td = map[string]*template.Template` & renders like: `err := ts.Execute(w, td) `
    * Use `render` in other methods in `handler.go` to shorten the code.

* @Code: [handlers.go](https://github.com/aayush4vedi/gisty/blob/0acff887af38613b257a88c197cfc599b86786b6/cmd/web/handlers.go#L68)


### 4.4 Catching Runtime Errors
* @What: If the app throws error, it's sending `200 OK` response along with a half-complete HTML page.This is bad.**To fix:**
    * Make the template render a two-stage process:(in `render` method)
    * 1.Make a ‘trial’ render by writing the template into a `buffer`. 
    * 2.If this fails: respond with an error message.But if it works:write the contents of the buffer to the usual `http.ResponseWriter` 
* @Code: [helpers.go]()

### 4.5 Common Dynamic Data
In some web applications there may be common dynamic data that you want to include on multiple webpages.E.g.: name,profile picture,CSRF token(in all pages with forms). 
* @What: Include the current year in the footer on every page
    * Add a new CurrentYear field to the templateData struct
    * create a new `addDefaultData()` helper method to inject the current year into an instance of a `templateData` struct
    * Call this method from `render() `helper to add the default data automatically for each page.
    * Update `footer.tmpl` to display `{{.CurrentYear}}`
* @Code: [templates.go](https://github.com/aayush4vedi/gisty/blob/7575d06c6fc0c3884a651b0575a933195233411e/cmd/web/templates.go#L15), [helpers.go](https://github.com/aayush4vedi/gisty/blob/7575d06c6fc0c3884a651b0575a933195233411e/cmd/web/helpers.go#L30), [footer.partial.tmpl](https://github.com/aayush4vedi/gisty/blob/7575d06c6fc0c3884a651b0575a933195233411e/ui/html/footer.partial.tmpl)

### 4.6 Custom Template Functions
* @What: Create a custom `humanDate()` function.Steps:
    * 1.create a `template.FuncMap` object containing the custom `humanDate()` function<br>
    ```go
    func humanDate(t time.Time) string {
        return t.Format("02 Jan 2006 at 15:04")
    }
    ```
    * 2.Use the `template.Funcs()` method to register this before parsing the templates. <br>`ts, err := template.New(name).Funcs(datefunc).ParseFiles(page) `
    * Call it in templates: `{{.Created | humanDate}}`
* @Code: [template.go](https://github.com/aayush4vedi/gisty/blob/89529adde28555a086ea076e009cfa997bd82717/cmd/web/templates.go#L19), [home.page.tmpl](https://github.com/aayush4vedi/gisty/blob/f422a94507445194e4e435ab71418bab57e18288/ui/html/home.page.tmpl#L17)
* @Notes: On custom template functions:
    * It can accept as many parameters as they need to, but they must *return one value* only. 
    * The only exception to this is if you want to return an error as the second value, in which case that’s OK too.


## 5 Middlewares

### 5.1 About Middleware
* **Middleware**: is some self-contained code which independently acts on a request before or after normal app handlers. E.g.:  log every request, compress every response, or check a cache before passing to req.
* Format of declaring middleware:<br>
    * Notes for `myMiddleware` function:
        * This middlware function is essentially a wrapper around the `next` handler. 
        * It establishes an anonymous function(returning a `http.Handler`) which closes over the next handler to form a closure.
```go
func myMiddleware(next http.Handler) http.Handler { 
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
        //...  if you want only login-uses to pass this
        if !isAuthorized(r) { 
            w.WriteHeader(http.StatusForbidden) 
            return 
        } 
        // ... do epic shit
        next.ServeHTTP(w, r) 
    }) 
} 
```
* Positioning the Middleware 
    * If before-mux(`myMiddleware → servemux →  handler`): it will act on incoming requests.Eg: Request-logger.
    * If before-mux(`servemux → myMiddleware →  handler `):act on specefic routes.Eg: auth middleware


* @Note: 
    * **On Actual FLow**:It’s important to know that when the last handler in the chain returns, *control is passed back up the chain in the reverse direction*. So, actual flow of control(while execution): <br>
    `secureHeaders → servemux → application handler → servemux → secureHeaders `
    * Why do we need headers: 
        * they essentially instruct the user’s web browser to implement some additional security measures to help prevent XSS and Clickjacking attacks.


### 5.2 Creating beforeMux-middleware: Set-Security-Headers
* @What:
    * Create reqd function `secureHeaders` in `middleware.go` (new file to hold all middlewares)
    * Reqd flow(so that it can act on every request): `secureHeaders → servemux → handler`
    * To do this we’ll need the `secureHeaders` middleware function to wrap our `servemux`, like so in `routes.go` : <br>`return secureHeaders(mux)` 

* @Code: [middleware.go](https://github.com/aayush4vedi/gisty/blob/09a4529397cd73789fe355b984ca40a2229b80e0/cmd/web/middleware.go#L8), [routes.go](https://github.com/aayush4vedi/gisty/blob/09a4529397cd73789fe355b984ca40a2229b80e0/cmd/web/routes.go#L16)

### 5.3 Creating beforeMux-middleware: Request-Logger
* @What: Create a middleware to record the IP address of the user, and which URL and method are being requested.
    * Create `logRequest` middleware func in `middleware.go`
    * Reqd flow: `logRequest ↔ secureHeaders ↔ servemux ↔ application handler `
    * Wrap `ecureHeaders(mux)` using this `logRequest`  in `routes.go` like: `return app.logRequest(secureHeaders(mux))`
* @Code: [middleware.go](https://github.com/aayush4vedi/gisty/blob/f910e67acaa4d29f274de8ee2e4379cefafc3107/cmd/web/middleware.go#L17), [routes.go](https://github.com/aayush4vedi/gisty/blob/f910e67acaa4d29f274de8ee2e4379cefafc3107/cmd/web/routes.go#L15)

### 5.4 Create Panic Recovery Middleware
* @What:Create a middleware to set `HTTP/1.1 500 Internal Server Error` in case of panic, called `recoverPanic`
    * Wrap `logRequest` around it: `return app.recoverPanic(app.logRequest(secureHeaders(mux)))`
    * @Code: [middleware.go](https://github.com/aayush4vedi/gisty/blob/220c7b4d4c2701a21a853ed302ff736dc1821c93/cmd/web/middleware.go#L24) , [routes.go](https://github.com/aayush4vedi/gisty/blob/220c7b4d4c2701a21a853ed302ff736dc1821c93/cmd/web/routes.go#L14)
* @Notes: On general Panic recovery in goroutine:
    * If not done by the middleware... and not by the panic recovery built into the Go HTTP server. They will cause your application  to *exit and bring down the server*. 


### 5.5 Composable Middleware Chains- pkg 'alice'
* @pkg: `justinas/alice` to manage middleware/handler chains. [doc](https://github.com/justinas/alice).Eg:
    * Better chaining code: 
        * w/o pkg: `return myMiddleware1(myMiddleware2(myMiddleware3(myHandler))) `
        * w/ pkg: `return alice.New(myMiddleware1, myMiddleware2, myMiddleware3).Then(myHandler) `
    * Creating new chains as *variables*:
    ```go
    myChain := alice.New(myMiddlewareOne, myMiddlewareTwo) 
    myOtherChain := myChain.Append(myMiddleware3) 
    return myOtherChain.Then(myHandler) 
    ```
* @What: Use `justinas/alice` in `routes.go` for better middleware chianing
* @Code: [routes.go](https://github.com/aayush4vedi/gisty/blob/67e9320ab22a1084ad3062c173b898b9797f9dcf/cmd/web/routes.go#L13)


## 6. RESTful Routing

### 6.1 About Routers
* @What: Move to **semantic/clean URLs**(i.e. RESTful)

|Method |Pattern|Handler|Action |
| ------|-------|-------|-------|
|GET    | /     | home  |Display the home page |
|GET |/gist/:id |showGist|Display a specific gist |
|GET |/gist/create |createGistForm |Display the new gist form |
|POST |/gist/create |createGist| Create a new gist |
|GET |/static/|http.FileServer| Serve a specific static file |
* @Issue:  Go’s servemux doesn’t support method based routing or semantic URLs with variables in them.
    * How to solve it: 
    * 1.w/o any 3rd party router: some tricks as talked about in [dotGo2014](https://www.youtube.com/watch?v=yi5A3cK1LNA&feature=youtu.be&t=11m44s)
    * 2.Use 3rd party package for routing

* List of some good 3rd party routers:
    * [Gin](https://github.com/gin-gonic/gin) :The fastest full-featured web framework for Golang.It features a Martini-like API with much better performance -- up to 40X faster.
        * Good: Minimalist, Usable-easy to learn and debug ,agille-very fast
        * Bad: Not for big load
    * [go-martini/martini](https://github.com/go-martini/martini) :Classy web framework for Go
        * Good:  lean,great third party support base, making it modular and scalable.
        * Bad: 40x slower than Gin.
    * [gorilla/mux](https://github.com/gorilla/mux): full featured, method-based routing,support for semantic URLs, REGEX urls.
        * Con: comparatively slow and memory hungry 
    * [Web.go](https://github.com/hoisie/web) :The easiest way to create web applications with Go
        * Good: very lightweight ,offers additional functionality over Go due to a tree routing system
        * Bad: Doesn’t extend Go that much: There’s extremely little what Web.go does that the standard Go framework cannot do on its own.
    * [Beego](https://beego.me/)
        * Good: Full featured, Solid built-in ORM for application database organization
        * Does automatic caching(higher chance of build failure),  extremely verbose
    * FastHTTP
    * [Revel](https://github.com/revel/revel) :
        * Good: Full featured, self-containted
        * Bad: Larger Codebase, No MongoDB
    * [ bmizerany/pat](https://github.com/bmizerany/pat):  more focused,lightweight. Provides method-based routing and support for semantic URLs... and not much else.
        * Con: pkg isn’t really maintained anymore. 
    * [go-chi/chi](https://github.com/go-chi/chi) :  uses a radix-tree for pattern matching and has a nice, flexible API. 
        * It also includes a [go-chi/chi/middleware](https://godoc.org/github.com/go-chi/chi/middleware) sub-package containing a range of useful middleware.
### 6.2 Implement RESTful Routes in app- Using Pat
//TODO: IMPLEMENT EVERYTHING WITH GIN LATERRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR.
* Basic Pat Syntax:
```go
mux := pat.New() 
mux.Get("/gist/:id", http.HandlerFunc(app.showGist)) 
```
* @What: 
    * Update `routes.go` using gorilla/mux
    * Remove redundant checkers in `handlers.go`
    * Update `<a href>` routes in templates
* @Code: [routes.go](https://github.com/aayush4vedi/gisty/blob/a0bd0684fec039877e4c9202610ba44c1ab5c973/cmd/web/routes.go#L13), [handlers.go](https://github.com/aayush4vedi/gisty/blob/a0bd0684fec039877e4c9202610ba44c1ab5c973/cmd/web/handlers.go)



## 7. Processing Forms

### 7.1 Setting Up a Form & Parsing Data to It
* @What: 
    * Create template
    * Create method in handler
    * Parse Form data
        * use `r.ParseForm()` method to parse the request body=>get to the form data `r.PostForm.Get()`
* @Code: [handlers.go](https://github.com/aayush4vedi/gisty/blob/28ae442274062243bca8e11f0a204c8c694f6916/cmd/web/handlers.go#L43)
* @Notes: On multiple fields:
```go
for i, item := range r.PostForm["items"] { 
fmt.Fprintf(w, "%d: Item %s\n", i, item) 
} 
```

### 7.2 Data Validation
* @What#1: Show all the validation errors
    * Make an error map `errors := make(map[string]string)` & fill it with all validation errors(along wiht err msg), ig exist.
* @What#2 : Displaying Validation Errors & Repopulating Fields 
    * Add new fields to `templateData`struct: 
        * `FormErrors` to hold any validation errors and 
        * `FormData` to hold any previously submitted data
    * update the `createGist` handler, so that if any validation errors are encountered the form is re-displayed with the relevant errors and form data passed to the template.
    * Updat `create.template` page to display form data if not null(using: `{{with .FormErrors.<field_name>}} `) for all the field.
* @Code: [handlers.go](https://github.com/aayush4vedi/gisty/blob/3f433490ef94b7c0d39211e2469724e8e9909daf/cmd/web/handlers.go#L75), [templates.go](https://github.com/aayush4vedi/gisty/blob/3f433490ef94b7c0d39211e2469724e8e9909daf/cmd/web/templates.go#L16), [create.page.tmpl](https://github.com/aayush4vedi/gisty/blob/3f433490ef94b7c0d39211e2469724e8e9909daf/ui/html/create.page.tmpl)

### 7.3 Scaling Data Validation(Refactor)
* @Issue: if your application has many forms then you can end up with quite a lot of repetition in your code and validation rules.
    * **Solution** : Creating a `forms` pkg to abstract some of this behavior and reduce the boilerplate code in our handler.
        * `gisty/pkg/forms/errors.go` :for Add() & Get() errors
        * `gisty/pkg/forms/forms.go`  : keeps all type of validators
    * Update `templateData` strcut, by adding `Form     *forms.Form` field
    * Update handlers
    * Update templates
* @Code: [pkg/forms/errors.go](https://github.com/aayush4vedi/gisty/blob/3cfd3dfa005f258feb7de1975a0110775703824b/pkg/forms/errors.go), [pkg/forms/forms.go](https://github.com/aayush4vedi/gisty/blob/3cfd3dfa005f258feb7de1975a0110775703824b/pkg/forms/forms.go), [templates.go](https://github.com/aayush4vedi/gisty/blob/3cfd3dfa005f258feb7de1975a0110775703824b/cmd/web/templates.go#L16), [handlers.go](https://github.com/aayush4vedi/gisty/blob/3cfd3dfa005f258feb7de1975a0110775703824b/cmd/web/handlers.go#L18), [create.page.tmpl](https://github.com/aayush4vedi/gisty/blob/3cfd3dfa005f258feb7de1975a0110775703824b/ui/html/create.page.tmpl#L4)

## 8. Stateful HTTP
* ### 8.1 About Session Managers
* List of **3-rd party session managers**:
    * [gorilla/sessions](https://github.com/gorilla/sessions)
    * [alexedwards/scs](https://github.com/alexedwards/scs)
    * [golangcollege/sessions](https://github.com/golangcollege/sessions) --used here

* @What: display a one-time confirmation message which the user sees after they’ve added a new gist.
    * To make this work, we need to start sharing data (or state) between HTTP requests for the same user.The most common way to do that is to implement a session for the user
    * Import session in App in `main.go`
    * Create a new `dynamicMiddleware` chain containing the middleware appropriate for our dynamic application routes only in `routes.go`
* ### 8.2 Working with Session Data
    * Use `session.Put()` in `handlers.go` to add a string value ("Your gist was saved successfully!") and the corresponding key ("flash") to the session data. Note that if there's no existing session for the current user (or their session has expired) then a new, empty, session for them will automatically be created by the session middleware.
    * Make `showGist` handler to retrieve the confirmation message
    * Update `templateData` in `templates.go` to contain `Flash       string `
    * update our `base.layout.tmpl` file to display the flash message, if one exists
* @Code: [main.go](https://github.com/aayush4vedi/gisty/blob/c9587eb4efc0eb0a5c2d9173d08809e925b65410/cmd/web/main.go#L23), [routes.go](https://github.com/aayush4vedi/gisty/blob/c9587eb4efc0eb0a5c2d9173d08809e925b65410/cmd/web/routes.go#L12) ,[handlers.go](https://github.com/aayush4vedi/gisty/blob/c9587eb4efc0eb0a5c2d9173d08809e925b65410/cmd/web/handlers.go#L73), [templates.go](https://github.com/aayush4vedi/gisty/blob/c9587eb4efc0eb0a5c2d9173d08809e925b65410/cmd/web/templates.go#L17), [base.layout.tmpl](https://github.com/aayush4vedi/gisty/blob/c9587eb4efc0eb0a5c2d9173d08809e925b65410/ui/html/base.layout.tmpl#L19)



## 9. Security
### 9.1 TLS Certificate
* `crypto/tls package` in Go’s standard library includes a `generate_cert.go` tool that we can use to easily create our own self-signed  certificate.
    * In `main.go` Use the `ListenAndServeTLS()` method to start the HTTPS server by passing in the paths to the TLS certificate and corresponding private key as its parameters
    * Not implemented here
### 9.2 Timeouts:
* Add timeouts in `main.go`:
    ```go
    srv := &http.Server{ 
        ...
        // Add Idle, Read and Write timeouts to the server. 
        IdleTimeout:  time.Minute, 
        ReadTimeout:  5 * time.Second, 
        WriteTimeout: 10 * time.Second, 
    } 
    ```
* @Code: [main.go](https://github.com/aayush4vedi/gisty/blob/609769d9d66af65354e4c8e281cf12b9cc7ad3a6/cmd/web/main.go#L59)

## 10. Auth

### 10.1. Routes Setup

| Method | Pattern| Handler | Action |
| -------|--------|---------|--------|
| GET | / |home|Display the home page |
|GET |/gist/:id |showGist |Display a specific gist |
|GET |/gist/create |createGistForm| Display the new gist form |
|POST |/gist/create |createGist |Create a new gist| 
|GET |/user/signup |signupUserForm |Display the user signup form |
|POST |/user/signup| signupUser|Create a new user |
|GET |/user/login |loginUserForm |Display the user login form |
|POST| /user/login |loginUser||Authenticate and login the user |
|POST| /user/logout| logoutUser||Logout the user |
|GET |/static/|http.FileServer |Serve a specific static file |

* @What:
    * Add new handlers for these auth routes in `handlers.go`
    * Add new routes in `routes.go`
    * Update `base.layout.tmpl`
* @Code: [handlers.go](https://github.com/aayush4vedi/gisty/blob/6fa2b792e3362c39dcda28bfdd0ddecfca84bcc6/cmd/web/handlers.go#L85), [routes.go](https://github.com/aayush4vedi/gisty/blob/6fa2b792e3362c39dcda28bfdd0ddecfca84bcc6/cmd/web/routes.go#L22), [base.layout.tmpl](https://github.com/aayush4vedi/gisty/blob/6fa2b792e3362c39dcda28bfdd0ddecfca84bcc6/ui/html/base.layout.tmpl#L19)

### 10.2. Creating a Users Model
* Run these mysql commands:
```mysql
USE gisty; 

CREATE TABLE users ( 
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, 
    name VARCHAR(255) NOT NULL, 
    email VARCHAR(255) NOT NULL, 
    hashed_password CHAR(60) NOT NULL, 
    created DATETIME NOT NULL 
); 

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email); 
```
* In `pkg/models/models.go` add a new `User` struct to hold the data for each user, plus a couple of new error types.
* Create a new file at `pkg/models/mysql/users.go` And then create a new UserModel type with some placeholder function
* In `main.go`:add a new field to our application struct so that we can make this model available to our handlers with `users         *mysql.UserModel `

### 10.3. User Signup and Password Encryption
* Create signup form `signup.page.tmpl`
    * Importantly, notice that we’re not re-displaying the password if the form fails validation — we don’t want there to be any risk of the browser (or other intermediary) caching the plain-text password entered by the user.
*  hook this up to the `signupUserForm` handler

* #### Validating & Encryption of the User Input:
    * Things to check:
        1. Check that the user’s name, email address and password are not blank. 
        2. Sanity check the format of the email address. 
        3. Ensure that the password is at least 10 characters long. 
        4. Make sure that the email address isn’t already in use.
    *  creating two helper new methods —  `MinLength()` and `MatchesPattern()` & `MatchesPattern` for sanity checking an email address
    * In `handlers.go` add some code to process the form and run the validation checks
    * Now, all that remains is to make sure that the email address isn’t already in use.HOW???
        * Because we’ve got a `UNIQUE` constraint on the `email` field of our users table, it’s already guaranteed that we won’t end up with two users in our database who have the same email address. So from a business logic and data integrity point of view we are already OK. 
        * But the question remains about how we communicate any email already in use problem to a user.
  
* @Notes: On Encryption
    * Go provides strong implementation of these encryption algos: `Argon2`, `scrypt` & `bcrypt`
    * Picked: [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt) because: it includes some helper functions specifically designed for hashing and checking passwords.
    * `bcrypt` provides 2 functions:
    * 1. Create a hash of a given plain-text password: `hash, err := bcrypt.GenerateFromPassword([]byte("my plain text password"), 12)`
        * Here `12` is called `cost` :means that that 4096 `(2^12)` *bcrypt iterations* will be used to hash the password
        * it will return `hash` value as a 60-character long hash which looks a bit like this: <br>`$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG`
        * This function also adds **some random salt** to the password to help avoid **rainbow-table attacks**
    * 2.  To check if a plain-text password matches a particular hash:<br>
        ```go
        hash := []byte("$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6GzGJSWG") 
        err := bcrypt.CompareHashAndPassword(hash, []byte("my plain text password"))
        ```
        * return `nil` if the plain-text password matches a particular hash, or an error if they don’t match. 
* #### Storing the User Details 
    * @What#1:  store the bcrypt hash of the password
    * @What#2: manage the potential error caused by a duplicate email violating the `UNIQUE` constraint that we added to the table
        * Update code for `Insert` method in `pkg/models/mysql/users.go`
        * Update `signupUser` handler

* @Code: [signup.page.tmpl](https://github.com/aayush4vedi/gisty/blob/03743b04c6a612355235096111398732218157fd/ui/html/signup.page.tmpl)  ,[models.go](https://github.com/aayush4vedi/gisty/blob/03743b04c6a612355235096111398732218157fd/pkg/models/models.go#L28) ,[users.go](https://github.com/aayush4vedi/gisty/blob/03743b04c6a612355235096111398732218157fd/pkg/models/mysql/users.go) ,  [forms.go](https://github.com/aayush4vedi/gisty/blob/03743b04c6a612355235096111398732218157fd/pkg/forms/forms.go#L11),  [main.go](https://github.com/aayush4vedi/gisty/blob/03743b04c6a612355235096111398732218157fd/cmd/web/main.go#L24) ,[handlers.go](https://github.com/aayush4vedi/gisty/blob/03743b04c6a612355235096111398732218157fd/cmd/web/handlers.go#L78)

### 10.4. User Login
* Create template `login.page.tmpl`
* hook this up so it’s rendered by our `loginUserForm` handler
* #### Verifying user details
    * @What: The core part of this verification logic will take place in the `UserModel.Authenticate()` method of our user model. Specifically,   we’ll need it to do two things: 
        * 1. First it should retrieve the hashed password associated with the email address from our MySQL `users` table. If the email doesn’t exist in the database, we want to return the `ErrInvalidCredentials` error that we made earlier. 
        * 2. If the email does exist, we want to compare the bcrypt-hashed password to the plain-text password that the user provided when 
        logging in. If they don’t match, we want to return the `ErrInvalidCredentials` error again. But if they do match, we want to return the user’s id value from the database.
    * Update `Authenticate()` method in `pkg/models/mysql/users.go`
    * Update the `loginUser` handler so that it parses the submitted login form data and calls this `UserModel.Authenticate()` method. 
* @Code: [login.page.tmpl](https://github.com/aayush4vedi/gisty/blob/23bdce5b22008091c2b5d3e9dfeb7ace3e31cead/ui/html/login.page.tmpl), [users.go](https://github.com/aayush4vedi/gisty/blob/23bdce5b22008091c2b5d3e9dfeb7ace3e31cead/pkg/models/mysql/users.go#L35) ,[handlers.go](https://github.com/aayush4vedi/gisty/blob/23bdce5b22008091c2b5d3e9dfeb7ace3e31cead/cmd/web/handlers.go#L115)

### 10.5. User Logout
* To logout:all we need to do is remove the `userID` value from the session
*  update the `logoutUser` hander to do exactly that
* @Code: [hanlders.go](https://github.com/aayush4vedi/gisty/blob/36ffeddf68f62b2bb0333641fbdd3de909184d95/cmd/web/handlers.go#L136)

### 10.6. User Authorization
* @What#1: Auth
    * 1. Only authenticated (logged in) users can create a new gist
    * 2. The contents of the navigation bar changes depending on whether a user is authenticated or not. Authenticated users should see links to ‘Home’, ‘Create gist’ and ‘Logout’. Unauthenticated users should see links to ‘Home’, ‘Signup’ and ‘Login’.
* How to check if current user is authenticated user or not by checking for the existence of a `userID` value in their session data.
* In `helpers.go` file and add `authenticatedUser()` helper function to retrieve and return the `userID` value from the session
* pass the `userID` value from `authenticatedUser()` to our HTML templates, so that we can toggle the contents of the navigation bar based on whether it is zero or not
    * add a new `AuthenticatedUser` field to our `templateData` struct in `templates.go`
    * In `helpers.go` update our `addDefaultData()` helper method so that the user ID is automatically added to the templateData struct every time we render a template
    * update the `base.layout.tmpl` file to toggle the navigation links using the `{{if .AuthenticatedUser}}` action

* @Code: [ helpers.go](https://github.com/aayush4vedi/gisty/blob/36ffeddf68f62b2bb0333641fbdd3de909184d95/cmd/web/helpers.go#L41), [base.layout.tmpl](https://github.com/aayush4vedi/gisty/blob/36ffeddf68f62b2bb0333641fbdd3de909184d95/ui/html/base.layout.tmpl#L17) , [templates.go](https://github.com/aayush4vedi/gisty/blob/36ffeddf68f62b2bb0333641fbdd3de909184d95/cmd/web/templates.go#L18)

* @What#2: **Restriction Url Access** : hiding the `Create gist` navigation link for any user that isn’t logged in.
    * so that if an unauthenticated user tries to visit any routes: `/gist/create` they are redirected to `/user/login`
    * How: In `middleware.go` file and create a new middleware function `requireAuthenticatedUser()`
    * Chain this middleware on all routes having `create` in them: in `routes.go`

* @Code: [middleware.go](https://github.com/aayush4vedi/gisty/blob/f3826d35bed792bf3c49f00b2fdff0c9bad4323c/cmd/web/middleware.go#L37), [routes.go](https://github.com/aayush4vedi/gisty/blob/f3826d35bed792bf3c49f00b2fdff0c9bad4323c/cmd/web/routes.go#L17)


### 10.7. CSRF Protection
* **How can a CSRF attack happen here**
    * A user logs into our application. Our session cookie is set to persist for 12 hours, so they will remain logged in even if they  navigate away from the application. 
    * The user then goes to a malicious website which contains some code that sends a request to POST /snippets/create to add a new snippet to our database. 
    * Since the user is still logged in to our application, the request is processed with their privileges. Completely unknown to them, a new snippet will be added to our database. 
* For Token based mitigation check: use `gorilla/csrf` or `justinas/nosurf` pkgs.
    * Both use `Double Submit Cookie` to prevent attacks. 
    * In this pattern a random CSRF token is generated and sent to the user in a CSRF cookie. This CSRF token is then added to a hidden field in each form that’s vulnerable to CSRF. When the form is submitted, both packages use some middleware to check that the hidden field value and cookie value match. 


## 11. Using Request Context
* **What is context:**
    * In essence every `http.Request` that our handlers process has a `context.Context` object embedded in it, which we can use to store  information during the lifetime of the request. 
* **Useage:**
    * to pass information between your pieces of middleware and other handlers. 
* **Syntax**:
    * we use the `r.Context()` method to retrieve the existing context from a request and assign it to the `ctx` variable <br> `ctx := r.Context() `
    * Then we use the `context.WithValue()` method to create a new copy of the existing context, with a `*models.User` struct added to it. In this case, we’ve used the string `user` as the key for this data in the context. <br>
    `ctx = context.WithValue(ctx, "user", &models.User{Name: "Bob Jones"})` 
    * Then finally we use the `r.WithContext()` method to create a copy of the request containing our new context`r = r.WithContext(ctx) `
* Shorthand delcaratin:
```go
ctx = context.WithValue(r.Context(), "user", &models.User{Name: "Bob Jones"}) 
r = r.WithContext(ctx) 
```
* **Retrieving values from context**
    * The important thing to explain here is that, behind the scenes, request context values are stored with the type `interface{}`. And that meansthat, after retrieving them from the context, you’ll need to assert them to their original type before you use them.
    * Use the `r.Context().Value()` method<br>
    ```go
    user, ok := r.Context().Value("user").(*models.User) 
    if !ok { 
        return fmt.Errorf("could not convert %T to *models.User", user) 
    } 
    fmt.Println(user.Name) 
    ```

## 12. Testing

* To run tests: `go test -v ./cmd/web/ `
* To run the app, **excluding** a file(`*_test.go`): `$ go run cmd/web/!(*_test).go `
* ### 12.1 Testing HTTP Handlers
    * In `handlers.go` file and create a new `ping` handler function:
        ```go
        func ping(w http.ResponseWriter, r *http.Request) { 
            w.Write([]byte("OK")) 
        } 
        ```
    * create a new `handlers_test.go` file to hold the test
    * Run the test: `go test -v ./cmd/web/ `
    * @Code: [handlers.go](https://github.com/aayush4vedi/gisty/blob/17e1198d3576afb59e92732c31e7af0bb1677ac4/cmd/web/handlers.go#L142), [handlers_test.go](https://github.com/aayush4vedi/gisty/blob/17e1198d3576afb59e92732c31e7af0bb1677ac4/cmd/web/handlers_test.go)

* ### 12.2 Testing Middleware
* To test the middleware `secureHeaders` in `middleware.go`, create `cmd/web/middleware_test.go` to containg the test ``
* @Code: [middleware_test.go]()

* ### 12.3 End-To-End Testing

* ### 12.4 Mocking Dependencies

* ### 12.5 Testing HTML Forms

* ### 12.6 Integration Testing

* ### 12.7 Profiling Test Coverage
(in book @chapter #13)
