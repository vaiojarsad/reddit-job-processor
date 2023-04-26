package redditjobprocessor

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	jobprocessor "github.com/vaiojarsad/job-processor"
	"github.com/vaiojarsad/reddit-video-downloader/reddit"
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
	return &RedditVideoDownloadJobProcessor{}, nil
}

type RedditVideoDownloadJob struct {
	VideoURL       string
	OutputFileName string
	NoAudio        bool
	NoVideo        bool
}

func (jp *RedditVideoDownloadJobProcessor) Process(job jobprocessor.Job) error {
	var parsedJob RedditVideoDownloadJob
	err := mapstructure.Decode(job, &parsedJob)
	if err != nil {
		return fmt.Errorf("error parsing job: %w", err)
	}
	err = reddit.DownloadVideo(parsedJob.VideoURL, parsedJob.OutputFileName, parsedJob.NoAudio, parsedJob.NoVideo)
	if err != nil {
		return fmt.Errorf("error downloading video: %w", err)
	}
	return nil
}
