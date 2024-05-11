package main

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"
	"time"
)

func Test001(t *testing.T) {
	filename := time.Now().Format(time.RFC3339) + filepath.Ext("3_pms.sql")
	dst := fmt.Sprintf("./static/migrations/%s", filename)
	log.Println(dst)
}

func Test002(t *testing.T) {
}
