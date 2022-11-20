package examplemodel

import (
	"embed"
	"errors"
	"io/fs"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/rmatsuoka/myvideos/model"
)

//go:embed videos/* thumb/*
var exampleFS embed.FS

type exampleModel int

var errNotImplement = errors.New("not implement")

const loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func randID() string {
	return strconv.Itoa(rand.Int() % 10000)
}

var exampleVideoInfo = &model.VideoInfo{
	ID:          randID(),
	Title:       "example video",
	FilePath:    []string{"videos/sample.mp4"},
	ThumbPath:   "thumb/thumb.jpg",
	Likes:       1145141919810,
	Views:       114514,
	PostDate:    time.Time{},
	Description: []byte(loremIpsum + "\n\n" + loremIpsum),
}

func (exampleModel) FS() fs.FS {
	return exampleFS
}

var tags map[string]*model.Tag

func init() {
	initTags := []*model.Tag{
		{ID: "fakeTag", Name: "偽のタグです", N: 3},
		{ID: "sampleTag", Name: "サンプルです", N: 14},
		{ID: "Yes", Name: "そうです", N: 24},
	}
	tags = make(map[string]*model.Tag)
	for _, t := range initTags {
		tags[t.ID] = t
	}
}

// AddTag implements Model
func (exampleModel) AddTag(videoID string, tagname string) (*model.Tag, error) {
	newtag := &model.Tag{
		ID:   randID(),
		Name: tagname,
		N:    rand.Int() % 100,
	}
	tags[newtag.ID] = newtag
	log.Printf("AddTag(videoID = %s, tagname = %s) => (%+v, nil)", videoID, tagname, newtag)
	return newtag, nil
}

// DeleteTag implements model.Model
func (exampleModel) DeleteTag(videoID string, tagID string) error {
	delete(tags, tagID)
	log.Printf("DeleteTag(videoID = %s, tagID = %s) => nil", videoID, tagID)
	return nil
}

// RelatedVideos implements Model
func (exampleModel) RelatedVideos(string) ([]*model.VideoInfo, error) {
	return []*model.VideoInfo{
		exampleVideoInfo,
		exampleVideoInfo,
		exampleVideoInfo,
		exampleVideoInfo,
	}, nil
}

func (exampleModel) SetVideoInfo(*model.VideoInfo) (string, error) {
	return randID(), errNotImplement
}

func mapToList[K comparable, V any](m map[K]V) []V {
	l := make([]V, 0, len(m))
	for _, v := range m {
		l = append(l, v)
	}
	return l
}

func (exampleModel) TagsWithVideo(videoID string) ([]*model.Tag, error) {
	l := mapToList(tags)
	log.Printf("TagsWithVideo(videoID = %s) => %+v, nil", videoID, l)
	return l, nil
}

// Tags implements Model
func (exampleModel) Tags() ([]*model.Tag, error) {
	return mapToList(tags), nil
}

// Updatemodel.VideoInfo implements Model
func (exampleModel) UpdateVideoInfo(string, *model.VideoInfo) error {
	return errNotImplement
}

// model.VideoInfo implements Model
func (exampleModel) VideoInfo(string) (*model.VideoInfo, error) {
	return exampleVideoInfo, nil
}
func (exampleModel) VideosWithTag(tagName string) ([]*model.VideoInfo, error) {
	return []*model.VideoInfo{
		exampleVideoInfo,
		exampleVideoInfo,
		exampleVideoInfo,
		exampleVideoInfo,
	}, nil
}

func (exampleModel) LookUpVideosBy(model.Attr) ([]*model.VideoInfo, error) {
	return []*model.VideoInfo{
		exampleVideoInfo,
		exampleVideoInfo,
		exampleVideoInfo,
		exampleVideoInfo,
	}, nil
}

func (exampleModel) IncrementLike(string) (int, error) {
	return 0, errNotImplement
}

var Model model.Model = exampleModel(0)
