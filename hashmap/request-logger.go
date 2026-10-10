package hashmap

type RequestLogger struct {
	messages  map[string]int
	timeLimit int
}

func (this *RequestLogger) requestLoggerInit(timeLimit int) {
	this.messages = make(map[string]int)
	this.timeLimit = timeLimit
}

func (this *RequestLogger) messageRequestDecision(timestamp int, request string) bool {
	prevTimestamp := this.messages[request]
	if prevTimestamp > 0 && (timestamp-prevTimestamp) >= this.timeLimit {
		this.messages[request] = timestamp
		return true
	} else if prevTimestamp > 0 {
		return false
	} else {
		this.messages[request] = timestamp
	}
	return true
}
