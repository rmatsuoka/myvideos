package examplemodel

import (
	"fmt"
	"log"
	"time"

	"github.com/rmatsuoka/myvideos/model"
)

const loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

var exampleVideoInfo = &model.VideoInfo{
	Title:     "example video",
	FilePath:  "videos/sample.mp4",
	ThumbPath: "thumb/thumb.jpg",
	Likes:     123,
	Views:     123456,
	PostDate:  time.Time{},
}

type videoInfoModel int

// Description implements model.VideoInfoModel
func (videoInfoModel) Description(videoID string) ([]byte, error) {
	panic("unimplemented")
}

// SetDescription implements model.VideoInfoModel
func (videoInfoModel) SetDescription(videoID string, desc []byte) error {
	panic("unimplemented")
}

func (videoInfoModel) Info(videoID string) (info *model.VideoInfo, err error) {
	defer func() {
		log.Printf("(VideoInfoModel).Get(videoID = %s) => (videoInfo = %+v, err = %v)", videoID, info, err)
	}()
	return exampleVideoInfo, nil
}

// Increment implements model.VideoInfoModel
func (videoInfoModel) Increment(videoID string, attr model.Attr) (incremented int, err error) {
	defer func() {
		log.Printf("(VideoInfoModel).Increment(videoID = %s, attr = %v) => (incremented = %d, err = %v)", videoID, attr, incremented, err)
	}()
	switch attr {
	case model.AttrLikes:
		exampleVideoInfo.Likes++
		incremented = exampleVideoInfo.Likes
	case model.AttrViews:
		exampleVideoInfo.Views++
		incremented = exampleVideoInfo.Views
	default:
		err = fmt.Errorf("increment a value of attr %v: not support or not exist such attribution", attr)
	}
	return
}

// Set implements model.VideoInfoModel
func (videoInfoModel) SetInfo(*model.VideoInfo) (videoID string, err error) {
	panic("unimplemented")
}

// Update implements model.VideoInfoModel
func (videoInfoModel) UpdateInfo(videoID string, info *model.VideoInfo) (err error) {
	defer func() {
		log.Printf("(VideoInfoModel).Update(videoID = %s, info = %+v) => (err = %v)", videoID, info, err)
	}()
	if info.Title != "" {
		exampleVideoInfo.Title = info.Title
	}
	return nil
}

var _ model.VideoModel = videoInfoModel(0)
