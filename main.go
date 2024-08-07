package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

const baseURL = "https://api.cloudflare.com/client/v4/accounts/%s/stream"

type Video struct {
	UID string `json:"uid"`
}

type VideosResponse struct {
	Result []Video `json:"result"`
}

func main() {
	token := flag.String("token", "", "Cloudflare streams API token")
	accountID := flag.String("accountid", "", "Cloudflare account ID")
	flag.Parse()
	deleteCount := 0

	if *token == "" || *accountID == "" {
		fmt.Println("Both --token and --accountid are required")
		os.Exit(1)
	}

	videos, err := fetchVideos(*token, *accountID)

	if err != nil {
		log.Fatalf("Error fetching videos: %v", err)
	}

	if len(videos) == 0 {
		fmt.Println("No videos found in this account")
		return
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10) // Limit concurrent deletions

	for _, video := range videos {
		wg.Add(1)
		go func(videoID string) {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire semaphore
			defer func() { <-semaphore }() // Release semaphore

			if err := deleteVideo(*token, *accountID, videoID); err != nil {
				log.Printf("Error deleting video %s: %v", videoID, err)
			} else {
				deleteCount++
				fmt.Printf("Deleted video: %s\n", videoID)
			}
		}(video.UID)
	}

	wg.Wait()
	fmt.Printf("All %d videos deleted\n", deleteCount)
}

func fetchVideos(token, accountID string) ([]Video, error) {
	url := fmt.Sprintf(baseURL, accountID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var videosResp VideosResponse
	if err := json.Unmarshal(body, &videosResp); err != nil {
		return nil, err
	}

	return videosResp.Result, nil
}

func deleteVideo(token, accountID, videoID string) error {
	url := fmt.Sprintf(baseURL+"/%s", accountID, videoID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete video, status code: %d", resp.StatusCode)
	}

	return nil
}
