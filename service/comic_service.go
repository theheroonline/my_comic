package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	db "my_comic/dao"
	"my_comic/model"
	"sort"
	"strings"
)

var ComicService = new(comicService)
var m = make(map[string][]string, 0)

type comicService struct{}

func (s *comicService) Page() (u model.Comic) {
	var comic model.Comic
	d := db.Db
	d.First(&comic, 10)
	return comic
}

// func HaveToken() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		fmt.Println("-----")
// 	}
// }

// 列表
func (s *comicService) GetComicList(pageno, pagesize int, search string) (comics []model.Comic, count int) {
	//fmt.Println("搜索关键词",search)
	comics = make([]model.Comic, 0)
	db := db.Db.Model(&model.Comic{})
	var total int64
	db.Count(&total)
	//sql搜索分页查询语句
	if search != `null` {
		db.Limit(pagesize).Offset((pageno-1)*pagesize).Where("name like ?", "%"+search+"%").Find(&comics)
	} else {
		db.Limit(pagesize).Offset((pageno - 1) * pagesize).Find(&comics)
	}
	return comics, int(total)
}

// 获取单条数据
func (s *comicService) GetComicById(pageno, pagesize, id int) (pics []model.ComicDetail, count int) {
	pics = make([]model.ComicDetail, 0)
	d := db.Db.Model(&model.ComicDetail{}).Where("comic_id = ?", id)
	var total int64
	d.Count(&total)
	d.Limit(pagesize).Offset((pageno - 1) * pagesize).Find(&pics)
	return pics, int(total)
}

// 删除
func (s *comicService) Delete(id int) (nub int64) {
	d := db.Db.Model(&model.Comic{}).Where("comic_id =?", id).Delete(new(model.Comic))
	return d.RowsAffected
}

func (s *comicService) Shuaxin() {
	// path := "/usr/local/comic"
	path := "D:/personal/my_comic"
	readFile(path)

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	b, _ := json.Marshal(m)
	fmt.Printf("m: %v\n", string(b))
}

func readFile(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			readFile(path + "/" + f.Name())
		} else {
			s := strings.Split(path, "/")
			name := s[len(s)-1]
			l := m[name]
			if len(l) > 0 {
				l = append(l, f.Name())
				m[name] = l
			} else {
				l := make([]string, 0)
				l = append(l, f.Name())
				m[name] = l
			}
		}
	}
}
