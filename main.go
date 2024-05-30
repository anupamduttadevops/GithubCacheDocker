package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/alecthomas/kingpin"
	_ "github.com/aws/aws-sdk-go"
	_ "github.com/dgraph-io/badger/v3"
	_ "github.com/elastic/go-elasticsearch/v7"
	_ "github.com/fsnotify/fsnotify"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocql/gocql"
	_ "github.com/gocraft/work"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/googleapis/gax-go/v2"
	_ "github.com/gorilla/websocket"
	_ "github.com/grpc-ecosystem/go-grpc-middleware"
	_ "github.com/honeycombio/beeline-go"
	_ "github.com/honeycombio/libhoney-go"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jessevdk/go-flags"
	_ "github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv"

	// _ "github.com/kataras/iris"
	_ "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "github.com/minio/minio-go"
	_ "github.com/mitchellh/mapstructure"
	_ "github.com/natefinch/lumberjack"
	_ "github.com/nats-io/nats.go"
	_ "github.com/olivere/elastic/v7"
	_ "github.com/patrickmn/go-cache"
	_ "github.com/pborman/uuid"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/robfig/cron/v3"
	_ "github.com/shirou/gopsutil"
	_ "github.com/sirupsen/logrus"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
	_ "github.com/streadway/amqp"
	_ "github.com/urfave/cli/v2"
	_ "github.com/valyala/fasthttp"
	_ "github.com/vektra/mockery"
	_ "go.etcd.io/etcd/client/v3"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	_ "golang.org/x/image/draw"
	_ "gonum.org/v1/gonum"
	_ "gonum.org/v1/plot"
	_ "gopkg.in/cheggaaa/pb.v1"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
