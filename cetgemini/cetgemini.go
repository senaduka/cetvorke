package cetgemini

import (
	"context"
	"fmt"
	"github.com/a-h/gemini"
	"github.com/a-h/gemini/mux"
	"github.com/senaduka/cetvorke/database"
	"log"
	"strconv"
	"strings"
	"sync"
)

func StartServer(waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()

	// Create the handlers for a domain (a.gemini).
	index := gemini.HandlerFunc(indexHandler)
	post := gemini.HandlerFunc(postHandler)

	// Create a router for gemini://a.gemini/require_cert and gemini://a.gemini/public
	routerA := mux.NewMux()
	routerA.AddRoute("/clanak/{postID}/*", post)
	routerA.AddRoute("/", index)

	// Set up the domain handlers.
	ctx := context.Background()
	a, err := gemini.NewDomainHandlerFromFiles("localhost", "server.crt", "server.key", routerA)
	if err != nil {
		log.Fatal("error creating domain handler A:", err)
	}

	err = gemini.ListenAndServe(ctx, ":1965", a)
	if err != nil {
		log.Fatal("error:", err)
	}
}

// write a gemini handler function
func indexHandler(w gemini.ResponseWriter, r *gemini.Request) {

	links, err := database.GetRecentLinks()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	response := "# Najnoviji članci\n\n"

	for _, link := range links {
		// append links to response in gemini feed format
		response += link.GemtextLink()
	}

	fmt.Println(response)

	w.Write([]byte(response))
}

func postHandler(w gemini.ResponseWriter, r *gemini.Request) {
	post_id_str := strings.Split(r.URL.Path, "/")[2]
	post_id, err := strconv.Atoi(post_id_str)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	post, err := database.GetPost(post_id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	response := post.GemtextPage()
	w.Write([]byte(response))
}
