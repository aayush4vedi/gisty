# Gisty
A naive clone of github's [gists](https://gist.github.com/), created to act as standard guidelines when developing any Go app.




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
 
### 1.3 Adding URL query parameter(for string)
 * @What: update the `showGist` handler so that it accepts an id query string parameter using `r.URL.Query().Get()` method
 * @Code: [main.go]()>showGist
 * @Notes: `io.Writer` == `http.ResponseWriter`
     * `io.Writer` type is an interface, and the `http.ResponseWriter` object satisfies the interface because it has a `w.Write()` method

 



## 2. Configuration & Error Handling


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







