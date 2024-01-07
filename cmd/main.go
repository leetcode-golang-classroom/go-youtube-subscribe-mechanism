package main

import (
	"bufio"
	"os"

	youtube_subscribe_mechanism "github.com/leetcode-golang-classroom/go-youtube-subscribe-mechanism/youtube-subscribe-mechanism"
)

func main() {
	ioWriter := bufio.NewWriter(os.Stdout)
	waterballChannel := youtube_subscribe_mechanism.NewYoutubeChannel("水球軟體學院", ioWriter)
	waterballObserver := youtube_subscribe_mechanism.NewThreeMinuteObserver("水球")
	pewDieDieChannel := youtube_subscribe_mechanism.NewYoutubeChannel("PewDiePie", ioWriter)
	fireballObserver := youtube_subscribe_mechanism.NewOneMinuteObserver("火球")
	waterballChannel.Subscribe(waterballObserver)
	pewDieDieChannel.Subscribe(waterballObserver)
	waterballChannel.Subscribe(fireballObserver)
	pewDieDieChannel.Subscribe(fireballObserver)
	waterballChannel.Upload(youtube_subscribe_mechanism.NewVideo("C1M1S2", "個世界正是物件導向的呢！", 4*60, ioWriter))
	pewDieDieChannel.Upload(youtube_subscribe_mechanism.NewVideo("Hello guys", "Clickbait", 30, ioWriter))
	waterballChannel.Upload(youtube_subscribe_mechanism.NewVideo("C1M1S3", "物件 vs. 類別", 1*60, ioWriter))
	pewDieDieChannel.Upload(youtube_subscribe_mechanism.NewVideo("Minecraft", "Let's play Minecraft", 30*60, ioWriter))
}
