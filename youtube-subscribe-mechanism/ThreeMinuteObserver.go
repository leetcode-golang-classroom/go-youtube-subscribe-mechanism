package youtube_subscribe_mechanism

type ThreeMinuteObserver struct {
	*YoutubeChannelSubscriber
}

func NewThreeMinuteObserver(name string) YoutubeChannelObserver {
	observer := NewSubscriber(name)
	return &ThreeMinuteObserver{
		observer,
	}
}

func (threeMinuteObserver *ThreeMinuteObserver) Update(video *Video, channel *YoutubeChannel) {
	if video.TimeLengthInSecond >= 180 {
		video.AddLike(threeMinuteObserver)
	}
}
