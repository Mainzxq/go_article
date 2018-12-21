package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func MakeUUIDS() (string, error) {

	var (
		myUUID uuid.UUID
		err error
	)

	myUUID, err = uuid.NewV4()

	if err != nil {
		fmt.Printf("got some err here: %s", err)
		return "", err
	}

	return myUUID.String(), nil
}