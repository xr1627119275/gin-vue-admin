package main

//
//import (
//	"fmt"
//	"github.com/shadow1ng/fscan/Plugins"
//	"github.com/shadow1ng/fscan/common"
//	"net/http"
//)
//
//func main() {
//	// Serve static files from the current directory
//	http.Handle("/", http.FileServer(http.Dir("./")))
//
//	// SSE endpoint
//	http.HandleFunc("/events", sseHandler)
//
//	// Start the server
//
//	if err := http.ListenAndServe(":8081", nil); err != nil {
//		fmt.Println("Error starting server:", err)
//	}
//	//
//	//
//	//// Print the captured output
//	//fmt.Print(outputBuffer.String())
//	//
//	//fmt.Printf("[*] 扫描结束,耗时: %s\n", time.Since(start))
//}
//
//func sseHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/event-stream")
//	w.Header().Set("Cache-Control", "no-cache")
//	w.Header().Set("Connection", "keep-alive")
//
//	// Simulate sending data every second
//
//	var Info common.HostInfo
//	common.Flag(&Info)
//	Info.Host = "6.6.6.3"
//	common.Ports = "9000"
//	Info.Printf = func(format string, v ...any) (int, error) {
//		fmt.Fprintf(w, "data: %s \n\n", fmt.Sprintf(format, v...))
//		flusher, ok := w.(http.Flusher)
//		if ok {
//			flusher.Flush() // Flush the data to the client
//		}
//		return 0, nil
//	}
//	Info.Println = func(v ...interface{}) (int, error) {
//		fmt.Fprintf(w, "data: %s \n\n", fmt.Sprint(v...))
//		flusher, ok := w.(http.Flusher)
//		if ok {
//			flusher.Flush() // Flush the data to the client
//		}
//		return 0, nil
//	}
//
//	//go func() {
//	//	for result := range common.Results {
//	//
//	//		if strings.HasPrefix(*result, "[+] InfoScan") {
//	//			//color.Green(*result)
//	//		} else if strings.HasPrefix(*result, "[+]") {
//	//			//color.Red(*result)
//	//		} else {
//	//			fmt.Println(*result)
//	//		}
//	//		fmt.Fprintf(w, "data: %s \n\n", *result)
//	//		flusher, ok := w.(http.Flusher)
//	//		if ok {
//	//			flusher.Flush() // Flush the data to the client
//	//		}
//	//		common.LogWG.Done()
//	//	}
//	//}()
//	common.CurrHostInstance = &Info
//	Plugins.Scan(Info)
//	//flusher, ok := w.(http.Flusher)
//	//if ok {
//	//	flusher.Flush() // Flush the data to the client
//	//}
//}
