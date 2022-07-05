package pkg

import (
	"crypto/sha256"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	resp "mango/data"
)

type IStorage interface {
	SetItem(resp.HashedData)(resp.HashedData, error)
	GetItem(int)(resp.HashedData, error)
	UpdateItem(int, resp.HashedData)(resp.HashedData, error)
	GetListItems()([]resp.HashedData, error)
	DeleteItem(int)(resp.HashedData, error)
	Check(string)(resp.HashedData, error)
}

type DatabaseStorage struct {
	Conn gorm.DB
}

func to_hash(url string) string{
	return fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:6]
}

func NewDatabaseStorage() (*DatabaseStorage, error) {
	Conn, err := GetDatabaseConnection()
	if err != nil {
		log.Fatalf("Error occurred while getting a DB connection from the connection pool %s", err)
		return nil, err
	}
	return &DatabaseStorage{Conn: *Conn}, nil
}

func(dbs *DatabaseStorage) DeleteItem(id int)(resp.HashedData, error){
	var data resp.HashedData
	result:= dbs.Conn.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&data)

	if err:=result.Error; err!=nil{
		return data, err
	}
	if err:=result.RowsAffected;err == 0{
		return data, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (dbs *DatabaseStorage) SetItem(data resp.HashedData)(resp.HashedData, error){
	data.HashedUrl = to_hash(data.OriginalUrl)
	data.CounterUrl = 0
	result := dbs.Conn.Create(&data)

	if err:=result.Error; err!=nil{
		return data, err
	}
	return data, nil
}

func (dbs *DatabaseStorage) GetItem(id int)(resp.HashedData, error){
	var data resp.HashedData
	result :=dbs.Conn.Where("id = ?", id).Find(&data)
	if err:=result.Error; err!=nil{
		return data, err
	}
	if err:=result.RowsAffected;err == 0{
		return data, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (dbs *DatabaseStorage) Check(hash_or_url string)(resp.HashedData, error){
	var data resp.HashedData
	result :=dbs.Conn.Where("alias = ?", hash_or_url).Or(
		"hashed_url = ?", hash_or_url).Find(&data)
	if err:=result.Error; err!=nil{
		return data, err
	}
	if err:=result.RowsAffected;err == 0{
		return data, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (dbs *DatabaseStorage) UpdateItem(id int, data resp.HashedData)(resp.HashedData, error){
	result :=dbs.Conn.Clauses(
		clause.Returning{}).Where("id = ?", id).Updates(&data)

	if err:=result.Error; err!=nil{
		return data, err
	}
	return data, nil
}

func (dbs *DatabaseStorage) GetListItems()([]resp.HashedData, error){
	var data []resp.HashedData
	result :=dbs.Conn.Find(&data)
	if err:=result.Error; err!=nil{
		return data, err
	}
	return data, nil
}
