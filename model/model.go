package model

import (
	"io/fs"
	"time"
)

type Tag struct {
	ID   string
	Name string
	N    int
}

type VideoInfo struct {
	ID          string
	Title       string
	FilePath    []string
	ThumbPath   string
	Likes       int
	Views       int
	PostDate    time.Time
	Description []byte
}

type Model interface {
	VideoInfo(videoID string) (*VideoInfo, error)
	RelatedVideos(videoID string) ([]*VideoInfo, error)

	IncrementLike(videoID string) (int, error)
	UpdateVideoInfo(videoID string, info *VideoInfo) error
	SetVideoInfo(*VideoInfo) (videoID string, err error)

	TagsWithVideo(videoID string) ([]*Tag, error)
	DeleteTag(videoID string, tagID string) error
	AddTag(videoID string, tagName string) (*Tag, error)

	// return all tags
	Tags() ([]*Tag, error)
	VideosWithTag(tagID string) ([]*VideoInfo, error)

	// return 30 videos with Attr
	LookUpVideosBy(Attr) ([]*VideoInfo, error)

	// return fs.FS which opens (VideoInfo).FilePath[n]
	FS() fs.FS
}

type Attr int

const (
	AttrDate = iota
	AttrLikes
	AttrViews
	AttrRandom
)
