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

func (a *nodeBotTypes[T]) print() {
	fmt.Print("UpdateID: ", a.info.UpdateId, ", ")
	fmt.Print("Text: ", a.info.Message.Text, ", ")
	fmt.Print("ChatPerson: ", a.info.Message.FromWho.UserName)
}

type queryBotTypes[T resultTypes] struct {
	start        *nodeBotTypes[T]
	end          *nodeBotTypes[T]
	length       int
	lastUpdateID int64
}

func (q *queryBotTypes[T]) initQuery(upd []updateInfo) {
	q.length = len(upd)
	for i := 0; i < q.length; i++ {
		if i == 0 {
			var newNode nodeBotTypes[T]
			q.start = newNode.initNode(upd[0])
			q.end = q.start
		} else {
			q.append(upd[i])
		}
	}
}

func updateQ(updNew []updateInfo, updOld []updateInfo) (bool, []updateInfo) {
	if len(updNew[len(updOld):len(updNew)-1]) != 0 {
		return true, updNew[len(updOld):]
	}
	return false, updNew
}

func (q *queryBotTypes[T]) append(upd updateInfo) {
	var newNode nodeBotTypes[T]
	q.end.next = newNode.initNode(upd)
	q.end = q.end.next
	q.lastUpdateID = q.end.info.UpdateId
}

func (q *queryBotTypes[T]) deleteFirst() {
	if q.start != q.end {
		q.start = q.start.next
		q.length--
	} else {
		q.start = nil
		q.end = nil
	}
}

func (q *queryBotTypes[T]) print() {
	var startNode = q.start
	fmt.Print("Query: [")
	for i := 0; i < q.length; i++ {
		fmt.Print("{")
		startNode.print()
		startNode = startNode.next
		fmt.Print("},")
	}
	fmt.Print("]")
}

func (q *queryBotTypes[T]) isEmpty() bool {
	if q.start == nil {
		return true
	}
	return false
}
