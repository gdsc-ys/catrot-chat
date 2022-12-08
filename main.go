package main

import (
	// db "catrot-chat/services/mysql"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	var apiEnv string = os.Getenv("API_ENV")
	if apiEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	// initialLoad()               // 초기에 로드 해야 할 것들 ex) 디비, s3 등등
	// middleware.SetMiddleWare(r) // 미들웨어 설정
	// scheduler.GoScheduler()     // 스케줄러 설정 (크론)
	// router.SetRoutes(r)         // 라우트 설정

	if apiEnv == "production" {
		_ = r.Run(":5000")
	} else {
		_ = r.Run(":38001")
	}
}

func initialLoad()  {
	var wg sync.WaitGroup
	
	// wg.Add(3)
	// go db.SetDBConnection(&wg)                                  // 디비 로드
	// go s3Client.SetS3Client(&wg)                                // s3 클라이언트 로드
	// go firebase.SetFirebaseApp(&wg)                             // 파이어베이스 클라이언트 로드
	wg.Wait()
}
