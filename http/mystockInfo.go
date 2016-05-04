package http

import (
	"github.com/open-falcon/agent/funcs"
	"net/http"
	"os"
	"strings"
	"github.com/toolkits/nux"
	"log"
)

type ServerInfo struct {
	Id   string
	IP   string
	Name string
	CPU  float64
	Mem  float64
}

func configMyStockRoutes() {
	http.HandleFunc("/emstock/tryget", HandleTryConnectStockAgent)
	http.HandleFunc("/emstock/deploy", HandleDeployStockAgent)
}

func HandleTryConnectStockAgent(w http.ResponseWriter, r *http.Request) {
	sourceAddress := strings.TrimSpace(r.FormValue("ip"))
	hostName, _ := os.Hostname()
	busy := 100.0 - funcs.CpuIdle()

	info := &ServerInfo{IP: sourceAddress, Name: hostName, CPU: busy, Mem: getPMemory()}
	RenderJson(w, Dto{Msg: "success", Data: info})
}

func HandleDeployStockAgent(w http.ResponseWriter, r *http.Request) {
	RenderJson(w, Dto{Msg: "success", Data: "deploy function under devlopment"})
}

func getPMemory() float64{
	m, err := nux.MemInfo()
	if err != nil {
		log.Println(err)
		return 0.0
	}
	memFree := m.MemFree + m.Buffers + m.Cached
	memUsed := m.MemTotal - memFree
	pMemUsed := 0.0
	if m.MemTotal != 0 {
		pMemUsed = float64(memUsed) * 100.0 / float64(m.MemTotal)
	}

	return pMemUsed
}
