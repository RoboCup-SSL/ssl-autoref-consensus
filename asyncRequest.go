package main

import "time"

type ASyncRequest struct {
	request     *SSL_RefereeRemoteControlRequest
	outcome     chan<- SSL_RefereeRemoteControlReply_Outcome
	receiveTime time.Time
	replied     bool
}

func (r *ASyncRequest) Reply(outcome SSL_RefereeRemoteControlReply_Outcome) {
	if !r.replied {
		r.outcome <- outcome
		r.replied = true
	}
}

func (r *ASyncRequest) TimedOut() bool {
	return r.receiveTime.Add(majorityTimeout).Before(time.Now())
}
