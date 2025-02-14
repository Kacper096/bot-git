package tests

import (
	"bot-git/bot/commands"
	"bot-git/config"
	"fmt"
	"testing"
)

var msgVersion = config.Msg{commands.VER, config.Image{}, false}

func TestVersion(t *testing.T) {
	msg := commands.VersionHandler.Handle("version")
	if msg != msgVersion {
		t.Error(fmt.Sprintf("Wrong response to 'version'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl, msg.IsFunnyMessage))
	}
}

func TestWersja(t *testing.T) {
	msg := commands.VersionHandler.Handle("wersja")
	if msg != msgVersion {
		t.Error(fmt.Sprintf("Wrong response to 'wersja'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl, msg.IsFunnyMessage))
	}
}

func TestVer(t *testing.T) {
	msg := commands.VersionHandler.Handle("ver")
	if msg != msgVersion {
		t.Error(fmt.Sprintf("Wrong response to 'ver'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl, msg.IsFunnyMessage))
	}
}
