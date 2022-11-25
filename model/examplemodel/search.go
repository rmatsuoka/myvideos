package examplemodel

import (
	"log"

	"github.com/rmatsuoka/myvideos/model"
)

type searchModel int

var videoInfoList = []*model.VideoInfo{
	exampleVideoInfo,
	exampleVideoInfo,
	exampleVideoInfo,
	exampleVideoInfo,
}

// BelongingTo implements model.SearchModel
func (searchModel) BelongingTo(tagID string) (infos []*model.VideoInfo, err error) {
	defer func() {
		log.Printf("(SearchModel).BelongingTo(tagID = %s) => (infos = %+v, err = %v)", tagID, infos, err)
	}()
	return videoInfoList, nil
}

// LookUpBy implements model.SearchModel
func (searchModel) LookUpBy(attr model.Attr) (infos []*model.VideoInfo, err error) {
	defer func() {
		log.Printf("(SearchModel).LookUpBy(attr = %v) => (infos = %+v, err = %v)", attr, infos, err)
	}()
	return videoInfoList, nil
}

// RelatedTo implements model.SearchModel
func (searchModel) RelatedTo(videoID string) (infos []*model.VideoInfo, err error) {
	defer func() {
		log.Printf("(SearchModel).RelatedTo(videoID = %s) => (infos = %+v, err = %v)", videoID, infos, err)
	}()
	return videoInfoList, nil
}

var _ model.SearchModel = searchModel(0)
