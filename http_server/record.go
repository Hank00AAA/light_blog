package http_server

import (
	"fmt"
	"go.uber.org/atomic"
	"net"
	"net/http"
	"os"
	"sync"
)

type Record struct {
	listener net.Listener
	svr *http.Server
	isGracefulRestart bool
	childPid int
	linkMap sync.Map
	isShutDownNow atomic.Bool
	pid int
}

var record *Record = NewRecord()

func NewRecord() *Record {
	r := &Record{
		linkMap:sync.Map{},
		pid:os.Getpid(),
	}
	return r
}

func GetSingleTon() *Record {
	return record
}

func (r *Record)GetLinkNum() uint64 {
	var i uint64
	r.linkMap.Range(func(key, value interface{}) bool {
		i += 1
		return true
	})

	return i
}

func (r *Record)GetStatus() string {
	linkLog := ""
	r.linkMap.Range(func(key, value interface{}) bool {
		linkLog += fmt.Sprintf(" %v ", key)
		return true
	})

	return fmt.Sprintf("data:%+v linkMap:%+v", *r, linkLog)
}

func (r *Record)AddLink(key string) {
	fmt.Println(fmt.Sprintf("pid:%v AddLink:%v", r.pid, key))
	r.linkMap.Store(key,true)
}

func (r *Record)DelLink(key string) {
	fmt.Println(fmt.Sprintf("pid:%v DelLink:%v", r.pid, key))
	r.linkMap.Delete(key)
}

func (r *Record)SetListener(l net.Listener) {
	r.listener = l
}

func (r *Record)GetListener()net.Listener {
	return r.listener
}

func (r *Record)GracefulRestart() {
	r.isGracefulRestart = true
}

func (r *Record)SetChildPid(pid int) {
	r.childPid = pid
}

func (r *Record)SetHttpServer(s *http.Server) {
	r.svr = s
}

func (r *Record)GetHttpServer() *http.Server {
	return r.svr
}

func (r *Record)GetPID() int {
	return os.Getpid()
}

func (r *Record)ShutDown() {
	r.isShutDownNow.Store(true)
}

func (r *Record)GetShutDown() bool {
	return r.isShutDownNow.Load()
}