package types

type IRoom interface {
	Renew(filters interface{}, realtimeNotificationChannel chan<- KuzzleNotification, subscribeResponseChan chan<- SubscribeResponse)
	Unsubscribe()
	GetRealtimeChannel() chan<- KuzzleNotification
	GetResponseChannel() chan<- SubscribeResponse
	GetRoomId() string
	GetFilters() interface{}
}

type SubscribeResponse struct {
	Room  IRoom
	Error error
}

type RoomList map[string]map[string]IRoom
