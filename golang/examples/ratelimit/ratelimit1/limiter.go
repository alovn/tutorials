package main

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan byte
}

//New
func NewConnLimiter(limit int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: limit,
		bucket: make(chan byte, limit),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Println("reached rate limit.")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	<-cl.bucket
	log.Printf("release")
}
