package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/aaryan003/GoFeed/internal/database"
)

func startScarping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequests time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Printf("Failed to get feeds to fetch: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	// Implementation for scraping a single feed

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID), 
	if err != nil {
		log.Println("Failed to mark feed as fetched: %v", err)
	}

	rss, err := urlToFeed(feed.URL)
	if err != nil {
		log.Printf("Failed to fetch feed %s: %v", feed.URL, err)
		return
	}
	
	for _, item := range rssFeed.Channe.Item {
		log.SPrintf("Found post: %s", item.Title)
	}
	log.Printf("")

}
