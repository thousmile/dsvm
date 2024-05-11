package core

import (
	"dsvm/global"
	"errors"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"strconv"
)

const MigrationsIndex = StaticDir + "/index"

const MigrationsDir = StaticDir + "/migrations"

func ReadMigrationsIndex() int {
	bytes, _ := os.ReadFile(MigrationsIndex)
	index, _ := strconv.Atoi(string(bytes))
	return index
}

func WriteMigrationsIndex(index int) {
	indexStr := strconv.FormatInt(int64(index), 10)
	_ = os.WriteFile(MigrationsIndex, []byte(indexStr), os.ModePerm)
}

func UpBySchema(ds, schema string) error {
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
		if err := goose.SetDialect(source.Dialect); err != nil {
			return err
		}
		if err := goose.Up(source.Db, "."); err != nil {
			return err
		}
	} else {
		return errors.New(fmt.Sprintf("no data source found : %s", ds))
	}
	return nil
}

func DownBySchema(ds string, version int64) error {
	if source, ok := global.DataSources[ds]; ok {
		if err := goose.SetDialect(source.Dialect); err != nil {
			return err
		}
		if err := goose.DownTo(source.Db, ".", version); err != nil {
			return err
		}
	} else {
		return errors.New(fmt.Sprintf("no data source found : %s", ds))
	}
	return nil
}
