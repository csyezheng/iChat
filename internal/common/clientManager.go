package common

// Management of long polling chients

import (
	"container/list"
	"sync"
	"time"
)

type clientStoreElement struct {
	key string
	val *Client
}

type ClientStore struct {
	rw       sync.RWMutex
	clients  map[string]*list.Element
	lru      *list.List
	lifeTime time.Duration
}

// NewClient creates a new client and saves it to the client store.
//func (cs *ClientStore) NewClient(conn interface{}, appid uint32) *Client {
//	var cli Client
//
//	switch conn.(type) {
//	case *websocket.Conn:
//		cli.proto = WEBSOCK
//		cli.wsConn, _ = conn.(*websocket.Conn)
//	case http.ResponseWriter:
//		cli.proto = LPOLL
//		cli.wrt, _ = conn.(http.ResponseWriter)
//	default:
//		cli.proto = NONE
//	}
//
//	if cli.proto != NONE {
//		// cli.subs = make(map[string]*Subscription)
//		cli.send = make(chan []byte, 64) // buffered
//	}
//
//	cli.lastTouched = time.Now()
//
//	if cli.proto != WEBSOCK {
//		// Websocket connections are not managed by SessionStore
//		cs.rw.Lock()
//
//		elem := cs.lru.PushFront(&clientStoreElement{cli.cid, &cli})
//		cs.clients[cli.cid] = elem
//		log.Println("cid: ", cli.cid)
//		log.Println("push success")
//
//		// Remove expired sessions
//		expire := cli.lastTouched.Add(-cs.lifeTime)
//		for elem = cs.lru.Back(); elem != nil; elem = cs.lru.Back() {
//			if elem.Value.(*clientStoreElement).val.lastTouched.Before(expire) {
//				cs.lru.Remove(elem)
//				delete(cs.clients, elem.Value.(*clientStoreElement).key)
//			} else {
//				break // don't need to traverse further
//			}
//		}
//		cs.rw.Unlock()
//	}
//
//	return &cli
//}

func (cs *ClientStore) Get(sid string) *Client {
	cs.rw.Lock()
	defer cs.rw.Unlock()

	if elem := cs.clients[sid]; elem != nil {
		cs.lru.MoveToFront(elem)
		//elem.Value.(*clientStoreElement).val.lastTouched = time.Now()
		return elem.Value.(*clientStoreElement).val
	}

	return nil
}

func (cs *ClientStore) Delete(sid string) *Client {
	cs.rw.Lock()
	defer cs.rw.Unlock()

	if elem := cs.clients[sid]; elem != nil {
		cs.lru.Remove(elem)
		delete(cs.clients, sid)

		return elem.Value.(*clientStoreElement).val
	}

	return nil
}

func NewClientStore(lifetime time.Duration) *ClientStore {
	store := &ClientStore{
		clients:  make(map[string]*list.Element),
		lru:      list.New(),
		lifeTime: lifetime,
	}

	return store
}
