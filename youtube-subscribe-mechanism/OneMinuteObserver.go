package youtube_subscribe_mechanism

type OneMinuteObserver struct {
	*YoutubeChannelSubscriber
}

func NewOneMinuteObserver(name string) YoutubeChannelObserver {
	observer := NewSubscriber(name)
	return &OneMinuteObserver{
		observer,
	}
}

func (oneMinuteObserver *OneMinuteObserver) Update(video *Video, channel *YoutubeChannel) {
	if video.TimeLengthInSecond <= 60 {
		channel.UnSubscribe(oneMinuteObserver)
	}
}
