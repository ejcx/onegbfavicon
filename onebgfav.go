package main
 
import (
        "crypto/rand"
        "fmt"
        "log"
        "net/http"
        "os"
)
 
func main() {
        fmt.Println("Started:")
	webpage := `
		<html>
		<head>
		<link id="fav" rel="shortcut icon" href="/empty"/>
		<link href='http://fonts.googleapis.com/css?family=Open+Sans' rel='stylesheet' type='text/css'>
		</head>
		<body>
			<style>
				body {
					margin:0 auto;
					background-color:#cbeaf8;
				}
				.content {
					margin:0 auto;
					height:100%;
					width:100%;
					text-align:center;
				}
				h1 {
					padding-top:200px;
					font-family: 'Open Sans', sans-serif;
				}
				h2 {
					font-family: 'Open Sans', sans-serif;
				}
				#noscript {
					text-align:center;
					font-family: 'Open Sans', sans-serif;
					font-size:18px;
					color:white;
					padding:15px;
				}
				#noscript a {
					text-decoration:none;
				}
				.btn {
					font-family: 'Open Sans', sans-serif;
					padding:5px;
					color:white;
					text-decoration:none;
					border-radius:5px;
				}
				#dlbtn {
					border:1px solid #d32d27;
				}
				#stopbtn {
					border:1px solid #6acc3d;
				}
				.green {
					background-color:#6acc3d;
				}
				.red {
					background-color:#d32d27;
				}
			</style>
			<script>
				function stopit(){
					document.location = document.location;
				}
				function faviconfact(where){
					var link = document.createElement("link");		
					link.id = "fav";
					link.href=where;
					link.rel="shortcut icon";
					return link;
				}
				function doit(){
					var fav = document.getElementById("fav");
					if (fav){
						fav.href = "/favicon.ico";
						var prim = document.getElementById("prim");
						if (prim) {
							prim.textContent = "You are now downloading a 1GB favicon";
						}
						var sec = document.getElementById("sec");
						if (sec) {
							sec.textContent = "That was a bad idea. If you want to stop...";
						}
						var dlbtn = document.getElementById("dlbtn");
						var stopbtn = document.getElementById("stopbtn");
						if (stopbtn){
							stopbtn.style.display="inline-block";
						}
						if (dlbtn){
							dlbtn.style.display="none";
						}
						fav.parentNode.removeChild(fav);
						var heads = document.getElementsByTagName("head");
						if (heads[0]){
							var head = heads[0];
							var newfav = faviconfact("/favicon.ico");
							head.appendChild(newfav);
						}
						
						return true;
					}
					return false;
				}
			</script>
			<div class="red" id="noscript">
			You have noscript turned on. <a href="/indexns.html">Click Here</a> to download a 1gb favicon
			</div>
			<div class="content">
				<h1 id="prim">Want to download a 1gb favicon? </h1>
				<h2 id="sec">Most phone browsers don't download favicons?</h2><h2> 
				<a class="btn red" id="dlbtn" href="" onclick="doit();return false">Click Here</a>
				<a class="btn green" id="stopbtn" style="display:none" href="">Click Here</a>
				</h2>
			</div>
		</body>
		</html>
		<script>
			var noscript = document.getElementById("noscript");
			noscript.style.display="none";
		</script>
	`
	webpagens := `
		<html>
		<head>
		<link href='http://fonts.googleapis.com/css?family=Open+Sans' rel='stylesheet' type='text/css'>
		</head>
		<body>
			<style>
				body {
					margin:0 auto;
					background-color:#cbeaf8;
				}
				.content {
					margin:0 auto;
					height:100%;
					width:100%;
					text-align:center;
				}
				h1 {
					padding-top:200px;
					font-family: 'Open Sans', sans-serif;
				}
				h2 {
					font-family: 'Open Sans', sans-serif;
				}
				.btn {
					font-family: 'Open Sans', sans-serif;
					padding:5px;
					color:white;
					text-decoration:none;
					border-radius:5px;
				}
				.green {
					background-color:#6acc3d;
				}
			</style>
			<div class="content">
				<h1 id="prim">You are now downloading a 1GB favicon</h1>
				<h2 id="sec">That was a bad idea. If you want to stop...</h2>
				<h2>
				<a class="btn green" id="stopbtn" href="/">Click Here</a>
				</h2>
			</div>
		</body>
		</html>
	`

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                w.Write([]byte(webpage))
        })
        http.HandleFunc("/indexns.html", func(w http.ResponseWriter, r *http.Request) {
                w.Write([]byte(webpagens))
        })
        http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 10000000; i++ {
			b := make([]byte, 4096)
			rand.Read(b)
			if b[0] == 0 {
				fmt.Printf("Writing to %s\n", r.RemoteAddr)
			}
			_, err := w.Write([]byte(b))
			if err != nil {
				return
			}
		}
        })
        err := http.ListenAndServe(":80", nil)
        if err != nil {
                fmt.Println("Server did not start")
                log.Println(err)
                os.Exit(1)
        }
}
