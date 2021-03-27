package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	//Create a unbuffered channel to receive match results
	results := make(chan *Result)

	//Setup a wait group so we can process all the feeds
	//The wait group helps prevent the program from terminating before
	// all search processing is completed.
	var waitGroup sync.WaitGroup

	//Set the gorountines we need, while we process 
	// the individual feeds
	waitGroup.Add(len(feeds))
	for _, feed := range feeds {
		//Retrieve a matcher for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the gorountines to perform the search
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	//Launch a gorountine to monitor when all the work is done
	go func() {
		//Wait for everything to be processed.
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program
		close(results)
	}()

	// Start displaying results, since they are now available
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
