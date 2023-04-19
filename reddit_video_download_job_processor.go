package redditjobprocessor

import (
	"fmt"
	jobprocessor "github.com/vaiojarsad/job-processor"
)

var (
	RedditVideoDownloadJobType jobprocessor.JobType = "reddit-download-video-job"
)

func init() {
	jobprocessor.RegisterFactoryForJobType(RedditVideoDownloadJobType, NewRedditVideoDownloadJobProcessor)
}

type RedditVideoDownloadJobProcessor struct {
}

func NewRedditVideoDownloadJobProcessor(ctx jobprocessor.ExecutionContext) (jobprocessor.JobProcessor, error) {
	fmt.Println(ctx)
	return &RedditVideoDownloadJobProcessor{}, nil
}

func (jp *RedditVideoDownloadJobProcessor) Process(job jobprocessor.Job) error {
	fmt.Println(job)
	return nil
}
