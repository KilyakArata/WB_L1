package main

import "fmt"

// MediaPlayer - старый интерфейс
type MediaPlayer interface {
	PlayAudio(file string)
}

// AudioPlayer - реализация старого интерфейса
type AudioPlayer struct{}

func (a *AudioPlayer) PlayAudio(file string) {
	fmt.Println("Playing audio file:", file)
}

// VideoPlayer - новый интерфейс
type VideoPlayer interface {
	PlayVideo(file string)
}

// AdvancedVideoPlayer - реализация нового интерфейса
type AdvancedVideoPlayer struct{}

func (a *AdvancedVideoPlayer) PlayVideo(file string) {
	fmt.Println("Playing video file:", file)
}

// VideoToAudioAdapter - адаптер, который позволяет использовать VideoPlayer как MediaPlayer
type VideoToAudioAdapter struct {
	videoPlayer VideoPlayer
}

func (v *VideoToAudioAdapter) PlayAudio(file string) {
	v.videoPlayer.PlayVideo(file)
}

func main() {
	// Использование старого интерфейса
	audioPlayer := &AudioPlayer{}
	audioPlayer.PlayAudio("song.mp3")

	// Использование нового интерфейса через адаптер
	advancedPlayer := &AdvancedVideoPlayer{}
	adapter := &VideoToAudioAdapter{videoPlayer: advancedPlayer}
	adapter.PlayAudio("video.mp4")
}
