package main

import (
	"fmt"
	"flag"
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/xml"
	"strings"
	"strconv"
	"time"
	"sync"
)

var MaxDuration time.Duration

func get_resp_time(url string) {
	time_start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching: %v", err)
	} else {
		defer resp.Body.Close()
	}
	since := time.Since(time_start)

	if since > MaxDuration {
		fmt.Println(time.Since(time_start), url)
	}
}

func main() {
	var url_flag = flag.String("u", "", "Streaming URL to download")
	var max_duration_flag = flag.String("d", "", "Max duration to wait for")
	var num_workers_flag = flag.Int("w", 2, "Number of workers to run")
	flag.Parse()

	if *url_flag == "" || *max_duration_flag == "" {
		fmt.Println("Usage: streamdl -d 200ms -u <http://example.com/stream>")
		os.Exit(0)
	}

	var stream_url string = *url_flag
	var max_duration_string string = *max_duration_flag
	var max_duration time.Duration
	var num_workers int = *num_workers_flag

	max_duration, err := time.ParseDuration(max_duration_string)
	if err != nil {
		log.Fatal(err)
	}

	MaxDuration = max_duration

	response, err := http.Get(stream_url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		manifest := new(SmoothStreamingMedia)

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Print(err)
		}

		err = xml.Unmarshal(body, &manifest)
		if err != nil {
			log.Print(err)
		}

		manifest_index := strings.LastIndex(stream_url, "/Manifest")
		stream_base_url := stream_url[:manifest_index] + "/"

		fragment_urls := make([]string, 0)

		for _, streamIndex := range manifest.StreamIndex {
			for _, qualityLevel := range streamIndex.QualityLevel {
				var timestamp int = 0
				for _, fragment := range streamIndex.Fragment{
					if fragment.Timestamp != 0 {
						timestamp = fragment.Timestamp
					}
					fragment_url := stream_base_url + streamIndex.Url
					fragment_url = strings.Replace(fragment_url, "{bitrate}", strconv.Itoa(qualityLevel.Bitrate), 1)
					fragment_url = strings.Replace(fragment_url, "{start time}", strconv.Itoa(timestamp), 1)
					fragment_urls = append(fragment_urls, fragment_url)
					timestamp = timestamp + fragment.Duration
				}
			}
		}

		// TODO: Remove temporary truncation
		// fragment_urls = fragment_urls[:50]

		wg := new(sync.WaitGroup)
		in := make(chan string, 2*num_workers)

		for i := 0; i < num_workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for url := range in {
					get_resp_time(url)
				}
			}()
		}

		for _, url := range fragment_urls {
			if url != "" {
				in <- url
			}
		}
		close(in)
		wg.Wait()
	}
}