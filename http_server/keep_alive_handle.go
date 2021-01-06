package http_server

import (
	"fmt"
	"net/http"
	"time"
)

func KeepAliveHandle(rsp http.ResponseWriter, req *http.Request) {
	ch := time.NewTicker(time.Second  * 5)
	finishTK := time.NewTimer(time.Second * 20)
	round := 0
		select {
		case <-ch.C:
			log := fmt.Sprintf("round:%v status:%v\n", round, GetSingleTon().GetStatus())
			_, err := rsp.Write([]byte(log))
			if err != nil {
				fmt.Println(err)
			}
			round += 1
			fmt.Println(log)
		case <-finishTK.C:
			ch.Stop()
			finishTK.Stop()
			_, _ = rsp.Write([]byte("tickFinish"))
			fmt.Println("finish")
			return
		}
}

