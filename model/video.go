package model

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type VideoInfo struct {
	Title     string        `db:"title"`
	FilePath  string        `db:"file_path"`
	ThumbPath string        `db:"thumb_path"`
	CreateAt  time.Time     `db:"create_at"`
	Duration  time.Duration `db:"duration"`
}

type VideoModel interface {
	Info(videoID string) (*VideoInfo, error)
	SetInfo(*VideoInfo) (videoID string, err error)
	UpdateInfo(videoID string, info *VideoInfo) error
}

func (m *Model) VideoInfo(videoID string) (*VideoInfo, error) {
	fmterr := func(err error) error {
		if err != nil {
			return fmt.Errorf("videoinfo %s: %w", videoID, err)
		}
		return nil
	}
	id, err := strconv.ParseInt(videoID, 10, 64)
	if err != nil {
		return nil, fmterr(err)
	}

	info := new(VideoInfo)
	row := m.db.QueryRow("SELECT title, file_path, thumb_path, create_at, duration FROM videos WHERE id = ?", id)
	err = row.Scan(&info.Title, &info.FilePath, &info.ThumbPath, &info.CreateAt, &info.Duration)
	return info, err
}

func (m *Model) CreateVideoInfo(info *VideoInfo) (videoID string, err error) {
	fmterr := func(err error) error {
		if err != nil {
			return fmt.Errorf(`createVideoInfo "%s": %w`, info.Title, err)
		}
		return nil
	}
	result, err := m.db.Exec(
		"INSERT INTO videos (title, file_path, thumb_path, duration) VALUES (?, ?, ?, ?, ?)",
		info.Title,
		info.FilePath,
		info.ThumbPath,
		info.Duration,
	)
	if err != nil {
		return "", fmterr(err)
	}
	id, err := result.LastInsertId()
	return strconv.FormatInt(id, 10), fmterr(err)
}

func (m *Model) UpdateVideoInfo(videoID string, info *VideoInfo) error {
	fmterr := func(err error) error {
		if err != nil {
			return fmt.Errorf("videoinfo %s: %w", videoID, err)
		}
		return nil
	}
	id, err := strconv.ParseInt(videoID, 10, 64)
	if err != nil {
		return fmterr(err)
	}

	valuesStmt := ""
	values := []any{}
	first := true
	for _, v := range updatedValues(*info) {
		if first {
			valuesStmt = v.name + " = ?"
			first = false
		} else {
			valuesStmt = valuesStmt + " ," + v.name + " = ?"
		}
		values = append(values, v.value)
	}

	values = append(values, id)
	_, err = m.db.Exec("UPDATE videos SET "+valuesStmt+" WHERE id = ?", values...)
	return fmterr(err)
}

type values struct {
	value any
	name  string
}

func updatedValues(s any) []values {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		panic("updateValues: type of s should be struct:" + v.Kind().String())
	}
	getName := func(sf reflect.StructField) string {
		name := sf.Tag.Get("db")
		if name == "" {
			name = sf.Name
		}
		return name
	}

	list := []values{}
	for i := 0; i < v.NumField(); i++ {
		sf := v.Type().Field(i)
		if !sf.IsExported() {
			continue
		}

		vf := v.Field(i)
		if !vf.IsZero() {
			list = append(list, values{v.Field(i).Interface(), getName(sf)})
		}
	}
	return list
}
