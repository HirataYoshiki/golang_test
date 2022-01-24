package dbinterface

import (
	"fmt"
	"time"
)


type clientResponse interface {}
type Data interface {}

type resourceClient interface {
	Search(table string, id string) (clientResponse, error)
	GetAll(table string) (clientResponse, error)
	Get(table string, id string) (clientResponse, error)
	Post(table string, data Data) (clientResponse, error)
	Update(table string, data Data) (clientResponse, error)
	Delete(table string, data Data) (clientResponse, error)
}

func SearchRecord (client resourceClient, id string) (response clientResponse, err error) {
	response, err = client.Search(id)
	if err != nil {
		fmt.Errorf("record (id: %s) not found.",id)
		return nil, err
	}
	return response, err
}

func getAllRecords (client resourceClient) (response clientResponse, err error) {
	response, err = client.GetAll()
	if err != nil {
		fmt.Errorf("Error occured %s",time.Now())
		return nil, err
	}
	return response, nil
}

func getRecord (client resourceClient,table string, id string) (response clientResponse, err error) {
	response, err = client.Get(table, id)
	if err != nil {
		fmt.Errorf("Error occured %s",time.Now())
		return nil, err
	}
	return response, nil
}
