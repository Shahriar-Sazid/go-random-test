package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

func httpConnectionTest() {
	// client trace to log whether the request's underlying tcp connection was re-used
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { log.Printf("conn was reused: %t", info.Reused) },
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	start := time.Now()
	// 1st request
	req, err := http.NewRequestWithContext(traceCtx, http.MethodGet, "https://loop.p-stageenv.xyz/api/v1/data/params", nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	req2, err := http.NewRequestWithContext(traceCtx, http.MethodGet, "https://api.publicapis.org/entries", nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = http.DefaultClient.Do(req2)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 50; i++ {
		req2, err = http.NewRequestWithContext(traceCtx, http.MethodGet, "https://api.publicapis.org/entries", nil)
		if err != nil {
			log.Fatal(err)
		}
		_, err = http.DefaultClient.Do(req2)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Print("Time elapsed: ")
	fmt.Println(time.Since(start))
}

func httpConnectionTestMultiThread() {
	// client trace to log whether the request's underlying tcp connection was re-used
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { log.Printf("conn was reused: %t", info.Reused) },
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	start := time.Now()
	// 1st request
	// req, err := http.NewRequestWithContext(traceCtx, http.MethodGet, "https://loop.p-stageenv.xyz/api/v1/data/params", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	req2, err := http.NewRequestWithContext(traceCtx, http.MethodGet, "https://api.publicapis.org/entries", nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = http.DefaultClient.Do(req2)
	if err != nil {
		log.Fatal(err)
	}
	respChan := make(chan http.Response)
	for i := 0; i < 50; i++ {
		func() {
			req2, err = http.NewRequestWithContext(traceCtx, http.MethodGet, "https://loop.p-stageenv.xyz/api/v1/data/params", nil)
			if err != nil {
				log.Fatal(err)
			}
			resp, err := http.DefaultClient.Do(req2)
			respChan <- *resp
			if err != nil {
				log.Fatal(err)
			}
		}()

		// req2, err = http.NewRequestWithContext(traceCtx, http.MethodGet, "https://api.publicapis.org/entries", nil)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// _, err = http.DefaultClient.Do(req2)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
	for i := 0; i < 50; i++ {
		<-respChan
	}
	fmt.Print("Time elapsed: ")
	fmt.Println(time.Since(start))
}
