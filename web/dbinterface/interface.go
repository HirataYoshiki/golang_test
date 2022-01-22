package dbinterface

import (
	"fmt"
	"time"
)


type clientResponse interface {}

type resourceClient interface {
	Search(table string, id string) (clientResponse, error)
	GetAll(table string) (clientResponse, error)
	Get(table string, id string) (clientResponse, error)
	Post(table string) (clientResponse, error)
	Update() (clientResponse, error)
	Delete() (clientResponse, error)
}

func getAllRecord (client resourceClient) (response clientResponse, err error) {
	response, err = client.GetAll()
	if err != nil {
		fmt.Errorf("Error occured %s",time.Now())
		return nil, err
	}
	return response, err
}


func SearchRecord (client resourceClient, id string) (response clientResponse, err error) {
	response, err = client.Search(id)
	if err != nil {
		fmt.Errorf("record (id: %s) not found.",id)
		return nil, err
	}
	return response, err
}