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
    * Go’s servemux treats the URL pattern `/` like a catch-all(as it's a Subtree path patterns).i.e. even though there's single url declared rn, every url will get redirecter here.Try hitting `/nonfuckingexist` ,it will get redirected to `/`
### 1.2 Add more routes
  * | url  |  handler | action   |
    | -----| -------- | ---------|
    | /    | home     | Display home page|
    | /gist| showGist | Display a specific gist|
    | /gist/create | createGist | Create new gist |

  * @Code: [main.go]()
  * @Notes:
    * Go’s servemux supports 2 different types of URL patterns: `fixed paths` and `subtree paths`. 
       * Fixed path patterns: are only matched when the request URL path **exactly** matches the fixed path.
       * Subtree path patterns: are matched whenever the start of a request URL path matches the subtree path.
          * For understanding, you can think of them as they have a wildcard at the end, like `"/**"` or `"/static/**"`
   
    

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







