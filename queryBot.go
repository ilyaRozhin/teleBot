package main

import "fmt"

type nodeBotTypes[T resultTypes] struct {
	next *nodeBotTypes[T]
	info updateInfo
}

func (a *nodeBotTypes[T]) initNode(upd updateInfo) *nodeBotTypes[T] {
	var n nodeBotTypes[T]
	n.info = upd
	return &n
}

type queryBotTypes[T resultTypes] struct {
	start        *nodeBotTypes[T]
	end          *nodeBotTypes[T]
	length       int
	lastUpdateID int64
}

func (q *queryBotTypes[T]) initQuery(upd []updateInfo) {
	q.length = len(upd)
	q.start = q.start.initNode(upd[0])
	q.end = q.start
	q.lastUpdateID = q.start.info.UpdateId
	for i := 1; i < q.length-1; i++ {
		q.append(upd[i])
	}
}

func (q *queryBotTypes[T]) append(upd updateInfo) {
	nd := q.end
	nd.next = nd.next.initNode(upd)
	q.end = nd.next
	q.lastUpdateID = q.end.info.UpdateId
}

func (q *queryBotTypes[T]) deleteFirst() {
	if q.length > 1 {
		q.start = q.start.next
		q.length--
	} else if q.length == 1 {
		fmt.Println("all, updates processed...")
	}
}
