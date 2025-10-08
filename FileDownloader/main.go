package main

import(
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func downloadfile(url string, wg *sync.WaitGroup, results chan string){
	defer wg.Done()

	// extract file name from url
	filename := filepath.Base(url)

	// create file on local disk
	out, err := os.Create(filename)

	if err != nil{
		results <- fmt.Sprintf("failed to create %s: %v", filename, err)
		return
	}

	defer out.Close()

	// get file data

	resp, err := http.Get(url)
	if err != nil{
		results <- fmt.Sprintf("failed to download %s : %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		results <- fmt.Sprintf("Bad status for %s: %v", url, err)
		return
	}

	// copy reponse body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Error saving %s: %v", filename, err)
		return
	}

	results <- fmt.Sprintf("Downloaded %s successfully", filename)

}

func main() {

	urls := []string{
		// "https://file-examples.com/storage/fe3b4a4baf9eabc7a3f18b7/2017/04/file_example_MP4_480_1_5MG.mp4",
		// "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
		"https://images.pexels.com/photos/33598068/pexels-photo-33598068.jpeg",
		"https://www.pexels.com/download/video/9974341/",
	}

	var wg sync.WaitGroup
	results := make(chan string)

	for _, url := range urls{
		wg.Add(1)
		go downloadfile(url, &wg, results)
	}

	go func ()  {
		wg.Wait()
		close(results)	
	}()

	for res := range results{
		fmt.Println(res)
	}
}