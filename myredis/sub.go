package main

import (

	"github.com/garyburd/redigo/redis"
	"log"

)

func main() {
	server := "172.20.0.54:6379"

	//option := redis.DialPassword("123456")
	c, err := redis.Dial("tcp", server)
	if err != nil {
		log.Println("connect server failed:", err)
		return
	}

	defer c.Close()

	c.Send("SUBSCRIBE", "TongTianDaDao")
	c.Flush()
	for {
		reply, err := redis.Values(c.Receive())
		if err != nil {
			log.Println("Receive failed:", err)
		}

		log.Println("reply:", reply)

		for i,v := range reply {

			switch vv:= v.(type){
			case int64:
				log.Println(i,":", vv)
			case []byte:
				log.Println(i,":", string(vv))
			case string:
				log.Println(i,":", vv)
			default:
				log.Println("unknown:",v)
			}

		}
	}

}
