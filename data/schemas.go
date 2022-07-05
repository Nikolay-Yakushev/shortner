package data

import "time"


type JSONResult struct {
	Code    int          `json:"code" `
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

type PostgresConfigSection struct {
	PostgresHost                 string    `ini:"host"`
	PostgresPort                 string    `ini:"port"`
	PostgresUser                 string    `ini:"user"`
	PostgresUserPassword         string    `ini:"passwd"`
	PostgresDatabaseName         string    `ini:"dbname"`
}

type ServerConfigSection struct {
	ServerHost     string `ini:"host"`
	ServerPort     string `ini:"port"`
	ServerUseHTTPs bool   `ini:"https"`
}

type Services struct {
	TTLCleanupServiceEnabled bool `ini:"cleanup_ttl"`
}

type Config struct {
	AppMode  string                `ini:"app_mode"`
	Postgres PostgresConfigSection `ini:"database"`
	Server   ServerConfigSection   `ini:"server"`
	Services Services              `ini:"services"`
}


type HashedData struct {
	ID           uint      `json:"id" gorm:"primaryKey,priority:3"`
	Alias        string    `json:"alias,omitempty" example:"ffsecurity" gorm:"index:alias_idx,unique,priority:2;default:null"`
	OriginalUrl  string    `json:"url,required" example:"https://github.com/swaggo/echo-swagger"`
	HashedUrl    string    `json:"hashed" gorm:"index:hashed_idx,unique,priority:1"`
	CounterUrl   int       `json:"count"`
	TTL     	 time.Time `json:"ttl,omitempty"`
	CreateDate   time.Time `json:"created_date,omitempty" gorm:"autoCreateTime:true"`
	UpdateDate   time.Time `json:"updated_date,omitempty" gorm:"autoCreateTime:true"`
}

