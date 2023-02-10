package model

import (
	"io/fs"
	"time"
)

//go:generate stringer -type Attr

type Tag struct {
	ID   string
	Name string
	N    int
}

type VideoInfo struct {
	Title     string
	FilePath  string
	ThumbPath string
	Likes     int
	Views     int
	PostDate  time.Time
	Duration  time.Duration
}

type VideoModel interface {
	Info(videoID string) (*VideoInfo, error)
	SetInfo(*VideoInfo) (videoID string, err error)
	UpdateInfo(videoID string, info *VideoInfo) error
	Description(videoID string) ([]byte, error)
	SetDescription(videoID string, desc []byte) error
	Increment(videoID string, attr Attr) (int, error)
}

type TagModel interface {
	All() ([]*Tag, error)
	Delete(videoID string, tagID string) error
	Add(videoID string, tagName string) (*Tag, error)
	WithVideoID(videoID string) ([]*Tag, error)
	WithTagID(tagID string) (*Tag, error)
}

type SearchModel interface {
	RelatedTo(videoID string) ([]*VideoInfo, error)
	LookUpBy(Attr) ([]*VideoInfo, error)
	BelongingTo(tagID string) ([]*VideoInfo, error)
}

type Model interface {
	VideoInfo() VideoModel
	Tag() TagModel
	Search() SearchModel
	// return fs.FS which opens (VideoInfo).FilePath[n]
	FS() fs.FS
}

type Attr int

const (
	AttrDate Attr = iota
	AttrLikes
	AttrViews
	AttrRandom
)
