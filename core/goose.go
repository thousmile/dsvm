package core

import (
	"dsvm/global"
	"dsvm/model"
	"errors"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"
)

func UpBySchema(ds, schema string, version int64) error {
	if source, ok := global.DataSources[ds]; ok {
		// 创建 模式
		sql1 := fmt.Sprintf(source.CreateSchemaSql, schema)
		if _, err := source.Db.Exec(sql1); err != nil {
			log.Println(err)
		}
		// 切换 模式
		sql2 := fmt.Sprintf(source.SwitchSchemaSql, schema)
		if _, err := source.Db.Exec(sql2); err != nil {
			log.Println(err)
		}
		return updateBySchema(source, version, true)
	} else {
		return errors.New(fmt.Sprintf("no data source found : %s", ds))
	}
}

func DownBySchema(ds string, version int64) error {
	if source, ok := global.DataSources[ds]; ok {
		return updateBySchema(source, version, false)
	} else {
		return errors.New(fmt.Sprintf("no data source found : %s", ds))
	}
}

func updateBySchema(source *model.GooseDataSource, version int64, f1 bool) error {
	if err := goose.SetDialect(source.Dialect); err != nil {
		return err
	}
	if f1 {
		if err := goose.UpTo(source.Db, ".", version); err != nil {
			return err
		}
	} else {
		if err := goose.DownTo(source.Db, ".", version); err != nil {
			return err
		}
	}
	return nil
}
