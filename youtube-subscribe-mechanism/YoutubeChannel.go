package youtube_subscribe_mechanism

import (
	"bufio"
	"fmt"
)

type YoutubeChannel struct {
	YoutubeChannelInterface
	Name        string
	subscribers []YoutubeChannelObserver
	videos      []*Video
	IoWriter    *bufio.Writer
}
type YoutubeChannelInterface interface {
	Subscribe(subscriber YoutubeChannelObserver)
	UnSubscribe(subscriber YoutubeChannelObserver)
	Upload(video *Video)
	Notify(video *Video)
}

func NewYoutubeChannel(name string, ioWriter *bufio.Writer) *YoutubeChannel {
	return &YoutubeChannel{
		Name:        name,
		IoWriter:    ioWriter,
		subscribers: []YoutubeChannelObserver{},
		videos:      []*Video{},
	}
}
func (youtubeChannel *YoutubeChannel) Subscribe(subscriber YoutubeChannelObserver) {
	youtubeChannel.subscribers = append(youtubeChannel.subscribers, subscriber)
	youtubeChannel.IoWriter.WriteString(fmt.Sprintf("%v 訂閱了 %v。\n", subscriber.GetName(), youtubeChannel.Name))
	youtubeChannel.IoWriter.Flush()
}
func (youtubeChannel *YoutubeChannel) UnSubscribe(subscriber YoutubeChannelObserver) {
	youtubeChannel.IoWriter.WriteString(fmt.Sprintf("%v 解除訂閱了 %v。\n", subscriber.GetName(), youtubeChannel.Name))
	youtubeChannel.IoWriter.Flush()
	newSubscribers := []YoutubeChannelObserver{}
	for _, currentSubscriber := range youtubeChannel.subscribers {
		if currentSubscriber != subscriber {
			newSubscribers = append(newSubscribers, currentSubscriber)
		}
	}
	youtubeChannel.subscribers = newSubscribers
}
func (youtubeChannel *YoutubeChannel) Upload(video *Video) {
	youtubeChannel.IoWriter.WriteString(fmt.Sprintf("頻道 %v 上架了一則新影片 \"%v\"。\n", youtubeChannel.Name, video.Name))
	youtubeChannel.IoWriter.Flush()
	youtubeChannel.videos = append(youtubeChannel.videos, video)
	youtubeChannel.Notify(video)
}

func (youtubeChannel *YoutubeChannel) Notify(video *Video) {
	for _, subscriber := range youtubeChannel.subscribers {
		subscriber.Update(video, youtubeChannel)
	}
}
