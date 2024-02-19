package ws

type Trainer struct {
	Content string `bson:"content"`
	StartTime int64 `bson:"startTime"`
	EndTime int64 `bson:"endTime"`
	Read uint `bson:"read"`
}

type Result struct{
	StartTime int64
	Msg string
	Content interface{}
	From string
}
