package youtube_subscribe_mechanism

import (
	"bufio"
	"fmt"
)

type Video struct {
	Name               string
	Title              string
	TimeLengthInSecond int
	Likers             []YoutubeChannelObserver
	ioWriter           *bufio.Writer
}

func NewVideo(name string, title string, timeLengthInSecond int, ioWriter *bufio.Writer) *Video {
	return &Video{
		Name:               name,
		Title:              title,
		TimeLengthInSecond: timeLengthInSecond,
		ioWriter:           ioWriter,
	}
}

func (video *Video) AddLike(liker YoutubeChannelObserver) {
	video.ioWriter.WriteString(fmt.Sprintf("%v 對影片 \"%v\" 按讚。\n", liker.GetName(), video.Name))
	video.ioWriter.Flush()
	video.Likers = append(video.Likers, liker)
}
