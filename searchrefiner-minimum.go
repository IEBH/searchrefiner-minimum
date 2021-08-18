package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	var c Config
	err = json.NewDecoder(f).Decode(&c)
	if err != nil {
		log.Fatalln(err)
	}

	g := gin.Default()

	// Writer / Logger
	err = os.MkdirAll("logs", 0777)
	if err != nil {
		log.Fatalln(err)
	}
	t := time.Now().Unix()
	ginLf, err := os.OpenFile(fmt.Sprintf("logs/sr-gin-%d.log", t), os.O_WRONLY|os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	gin.DefaultWriter = io.MultiWriter(ginLf, os.Stdout)
	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://searchrefinery.sr-accelerator.com", "http://localhost:8080"}
	// OPTIONS method for preflight request
	corsConfig.AddAllowMethods("OPTIONS")
	g.Use(cors.New(corsConfig))

	g.POST("/api/queryvis", handleTree)

	log.Fatalln(g.Run(c.Host))

}