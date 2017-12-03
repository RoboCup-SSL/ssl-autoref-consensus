package main

import "math"

func findMatchingRequests(reqBuffer []*ASyncRequest, request *SSL_RefereeRemoteControlRequest) []*ASyncRequest {
	matchingRequests := make([]*ASyncRequest, 0)
	for _, otherRequest := range reqBuffer {
		if otherRequest.request.equals(request) {
			matchingRequests = append(matchingRequests, otherRequest)
		}
	}
	return matchingRequests
}

func (m *SSL_RefereeRemoteControlRequest) equals(o *SSL_RefereeRemoteControlRequest) bool {

	if m.LastCommandCounter != nil || o.LastCommandCounter != nil {
		if m.LastCommandCounter == nil || o.LastCommandCounter == nil || m.GetLastCommandCounter() != o.GetLastCommandCounter() {
			return false
		}
	}

	if m.Stage != nil || o.Stage != nil {
		if m.Stage == nil || o.Stage == nil || m.GetStage() != o.GetStage() {
			return false
		}
	}

	if m.Command != nil || o.Command != nil {
		if m.Command == nil || o.Command == nil || m.GetCommand() != o.GetCommand() {
			return false
		}
	}

	if m.Card != nil || o.Card != nil {
		if m.Card == nil || o.Card == nil || !m.GetCard().equals(o.GetCard()) {
			return false
		}
	}

	if m.DesignatedPosition != nil || o.DesignatedPosition != nil {
		if m.DesignatedPosition == nil || o.DesignatedPosition == nil || m.GetDesignatedPosition().similar(o.GetDesignatedPosition()) {
			return false
		}
	}

	return true
}

func (m *SSL_RefereeRemoteControlRequest_CardInfo) equals(o *SSL_RefereeRemoteControlRequest_CardInfo) bool {
	if o == nil {
		return false
	}
	if m.Team == nil || o.Team == nil || m.Type == nil || o.Type == nil {
		return false
	}
	if m.GetType() != o.GetType() || m.GetTeam() != o.GetTeam() {
		return false
	}
	return true
}

func (m *SSL_Referee_Point) similar(o *SSL_Referee_Point) bool {
	if o == nil {
		return false
	}

	dx := m.GetX() - o.GetX()
	dy := m.GetY() - o.GetY()
	diff := math.Sqrt(float64(dx*dx + dy*dy))

	if diff > positionToleranceMM {
		return false
	}
	return true
}
