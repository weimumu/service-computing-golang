# cloudgo
A homework for Golang(Third Homework) 

##### About Martini

Martini is a powerful package for quickly writing modular web applications/services in Golang.

Its features can be show as followed:

- Extremely simple to use.
- Non-intrusive design.
- Plays nice with other Golang packages.
- Awesome path matching and routing.
- Modular design - Easy to add functionality, easy to rip stuff out.
- Lots of good handlers/middlewares to use.
- Great 'out of the box' feature set.
- **Fully compatible with the http.HandlerFunc interface.**
- Default document serving (e.g., for serving AngularJS apps in HTML5 mode).

#####How to use Martini 

1. Download the Martini to your GoPath by use

   ```
   go get github.com/go-martini/martini
   ```

2. Finish the server.go like:

   ```
   package main

   import "github.com/go-martini/martini"

   func main() {
   	m := martini.Classic()
   	m.Get("/", func() string {
   		return "Hello! This is my first cloud web server"
   	})
   	m.Run()
   }
   ```

3. Run the server.go ![run](/image/run.png)

4. Use the curl tool to test ![curl](/image/curl.png)

5. Use the browser to see ![brower](/image/brower.png)

6. Use the ab to test ![ab](/image/ab.png)

 As we can see in the screenshot:

* Document length is 40 bytes
* 50% of the request were serverd after 9ms. 100% of the request were serverd after 38ms
* Transfer rate: 1372.33[Kbytes/sec]



##### Thanks!

