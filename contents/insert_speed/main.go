package insert_speed

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"time"
	"vs/pkg/orm"
	"vs/pkg/tests"
	"vs/pkg/utils"
)

func runRDB(totalCnt, rps int) {
	// 초기화
	db := orm.GetDB(os.Args[1])
	Init(db)

	// 호출
	start := time.Now()
	status := tests.LoopFunction(db, totalCnt, rps, createRDB)
	fmt.Println("[ 완료 ]\n▶️ 동작 시간 : ", time.Since(start), "\n▶️ 성공 : ", status.Success, "\n▶️ 실패 : ", status.Failed)
}

// createRDB 새로운 데이터 생성
func createRDB(arg interface{}) error {
	db := arg.(*gorm.DB)
	return db.Model(&User{}).Create(&User{
		Name: utils.RandomString(30),
		Age:  rand.Intn(100) + 1,
	}).Error
}

func Run() {
	totalCnt := 1000000
	switch os.Args[1] {
	case "mongo":
		runMongo(totalCnt, 0)
	case "sqlite":
		runRDB(totalCnt, 10000)
	case "mysql":
		runRDB(totalCnt, 10000)
	case "postgres":
		runRDB(totalCnt, 10000)
	}
}
