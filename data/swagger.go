package data

import "time"

type ResponseHashedData struct {
	ID           uint      `json:"id" example:"123"`
	Alias        string    `json:"alias,omitempty" example:"ffsecurity"`
	OriginalUrl  string    `json:"url,required" example:"https://github.com/swaggo/echo-swagger"`
	HashedUrl    string    `json:"hashed" example:"someHash"`
	CounterUrl   int       `json:"count" example:"123"`
	TTL     	 time.Time `json:"ttl" example:"2009-11-10T23:00:00-00:00"`
	CreateDate   time.Time `json:"created_date" example:"2009-11-10T23:00:00-00:00"`
	UpdateDate   time.Time `json:"updated_date" example:"2009-11-10T23:00:00-00:00"`
}


type HashedDataPutRequest struct {
	Alias        string    `json:"alias,omitempty" example:"ffsecurity"`
	OriginalUrl  string    `json:"url,required" example:"https://github.com/swaggo/echo-swagger"`
	HashedUrl    string    `json:"hashed" example:"someHash"`
	CounterUrl   int       `json:"count" example:"123"`
	TTL     	 time.Time `json:"ttl" example:"2009-11-10T23:00:00-00:00"`
	CreateDate   time.Time `json:"created_date" example:"2009-11-10T23:00:00-00:00"`
	UpdateDate   time.Time `json:"updated_date" example:"2009-11-10T23:00:00-00:00"`

}

type HashedDataCreateRequest struct {
	Alias        string    `json:"alias,omitempty" example:"ffsecurity"`
	OriginalUrl  string    `json:"url,required" example:"https://github.com/swaggo/echo-swagger"`
}

