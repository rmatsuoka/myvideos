package examplemodel

import (
	"encoding/base32"
	"fmt"
	"log"
	"math/rand"

	"github.com/rmatsuoka/myvideos/model"
)

func mapToList[K comparable, V any](m map[K]V) []V {
	l := make([]V, 0, len(m))
	for _, v := range m {
		l = append(l, v)
	}
	return l
}

// key is tagID which is of encoding tag's name in base32.
type tagModel map[string]*model.Tag

var tagMap model.TagModel

func init() {
	initTagName := []string{"偽のタグ1", "偽のタグ2", "偽のタグ3"}
	tagMap = make(tagModel)
	for _, t := range initTagName {
		tagMap.Add("", t)
	}
}

// Add implements model.TagModel
func (t tagModel) Add(videoID string, tagName string) (tag *model.Tag, err error) {
	defer func() {
		log.Printf("(TagModel).Tag(videoID = %s, tagName = %s) => (tag = %+v, err = %v)", videoID, tagName, tag, err)
	}()

	tagID := base32.StdEncoding.EncodeToString([]byte(tagName))
	tag, ok := t[tagID]
	if ok {
		return tag, fmt.Errorf(`add a tag: the video has already the tag "%s"`, tagName)
	}
	newTag := &model.Tag{
		ID:   tagID,
		Name: tagName,
		N:    rand.Int() % 100,
	}
	t[tagID] = newTag
	return newTag, nil
}

// All implements model.TagModel
func (t tagModel) All() (tags []*model.Tag, err error) {
	defer func() {
		log.Printf("(TagModel).All() => (tags = %+v, err = %v)", tags, err)
	}()
	return mapToList(t), nil
}

// Delete implements model.TagModel
func (t tagModel) Delete(videoID string, tagID string) (err error) {
	defer func() {
		log.Printf("(TagModel).Delete(videoID = %s, tagID = %s) => (err = %v)", videoID, tagID, err)
	}()
	delete(t, tagID)
	return nil
}

// WithTagID implements model.TagModel
func (t tagModel) WithTagID(tagID string) (tag *model.Tag, err error) {
	defer func() {
		log.Printf("(TagModel).WithTagID(tagID = %s) => (tag = %+v, err = %v)", tagID, tag, err)
	}()
	tag, ok := t[tagID]
	if !ok {
		return nil, fmt.Errorf(`get a tag with %s: no such a tag`, tagID)
	}
	return tag, nil
}

// WithVideoID implements model.TagModel
func (t tagModel) WithVideoID(videoID string) (tags []*model.Tag, err error) {
	defer func() {
		log.Printf("(TagModel).WithVideoID(videoID = %s) => (tags = %+v, err = %v)", videoID, tags, err)
	}()
	return mapToList(t), nil
}
