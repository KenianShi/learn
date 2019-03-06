package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"html/template"
)

var addr = flag.String("addr","localhost:8080","http service address")

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter,r *http.Request){
	c,err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Println("upgrader err:",err)
		return
	}
	defer c.Close()
	count := 0
	for{
		count++
		mt,message,err := c.ReadMessage()
		if err != nil {
			log.Println("readMessage err:",err)
			break
		}
		log.Printf("recv:%s \n",message)
		messageStr := string(message)
		messageStr = "Server receive " + string(count)+ " times: "+messageStr
		err = c.WriteMessage(mt,[]byte(messageStr))
		if err != nil {
			log.Println("writeMessage err:",err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(3)
	http.HandleFunc("/",home)
	http.HandleFunc("/echo",echo)
	log.Fatal(http.ListenAndServe(*addr,nil))
}



func home(w http.ResponseWriter,r *http.Request){
	homeTemplate.Execute(w,"ws://"+r.Host+"/echo")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
