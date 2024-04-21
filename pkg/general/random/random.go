package random

import (
	"github.com/google/uuid"
)


func MakeRandomId() string {
	randomVal := uuid.New()
	
	return randomVal.String()
}