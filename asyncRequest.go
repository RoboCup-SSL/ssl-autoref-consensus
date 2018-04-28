package main

import (
	"github.com/RoboCup-SSL/ssl-go-tools/sslproto"
	"time"
)

type ASyncRequest struct {
	request     *sslproto.SSL_RefereeRemoteControlRequest
	outcome     chan<- sslproto.SSL_RefereeRemoteControlReply_Outcome
	receiveTime time.Time
	replied     bool
}

func (r *ASyncRequest) Reply(outcome sslproto.SSL_RefereeRemoteControlReply_Outcome) {
	if !r.replied {
		r.outcome <- outcome
		r.replied = true
	}
}

func (r *ASyncRequest) TimedOut() bool {
	return r.receiveTime.Add(time.Second * time.Duration(*majorityTimeout)).Before(time.Now())
}

func findMatchingRequests(reqBuffer []*ASyncRequest, request *sslproto.SSL_RefereeRemoteControlRequest) []*ASyncRequest {
	matchingRequests := make([]*ASyncRequest, 0)
	for _, otherRequest := range reqBuffer {
		if otherRequest.request.Equals(request) {
			matchingRequests = append(matchingRequests, otherRequest)
		}
	}
	return matchingRequests
}
