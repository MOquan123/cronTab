package master

import (
	"net"
	"net/http"
	"time"
)

//任务的HTTP接口
type ApiServer struct {
	httpServer *http.Server
}

var (
	//单例对象
	G_apiServer *ApiServer
)

//保存任务接口
func handleJobServer(http.ResponseWriter, *http.Request) {

}

//初始化服务

func InitApiServer() (err error) {
	var (
		mux        *http.ServeMux
		erro       error
		httpServer *http.Server
		Listener   net.Listener
	)

	//配置路由(转接口)
	mux = http.NewServeMux()
	mux.HandleFunc("/Job/Save", handleJobServer)

	//启动监听
	if Listener, erro = net.Listen("tcp", ":8070"); erro != nil {
		return erro
	}

	//创建一个HTTP服务
	httpServer = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      mux,
	}

	//初始化单例对象
	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	//启动服务端
	go httpServer.Serve(Listener)

}
