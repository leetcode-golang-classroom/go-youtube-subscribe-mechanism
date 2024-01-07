package youtube_subscribe_mechanism

type YoutubeChannelSubscriber struct {
	YoutubeChannelObserver
	Name string
}

type YoutubeChannelObserver interface {
	Update(video *Video, channel *YoutubeChannel)
	GetName() string
}

func NewSubscriber(name string) *YoutubeChannelSubscriber {
	return &YoutubeChannelSubscriber{
		Name: name,
	}
}

func (subscriber *YoutubeChannelSubscriber) GetName() string {
	return subscriber.Name
}
