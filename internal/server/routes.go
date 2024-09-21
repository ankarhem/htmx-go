package server

import (
	"htmx/cmd/web"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/coder/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(web.Home()))
	mux.HandleFunc("/api/random_number", web.RandomNumberHandler)

	mux.HandleFunc("/livereload", s.livereloadHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)

	return mux
}

func (s *Server) livereloadHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		err := socket.Write(socketCtx, websocket.MessageText, []byte("keepalive"))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 10)
	}
}
