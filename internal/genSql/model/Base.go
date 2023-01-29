package model

import "time"

type BaseModel struct {
	GenTime        time.Time
	ModifiedTime   time.Time
	GenUserId      int
	ModifiedUserId int
	DeleteStatus   int
	Status         int
	OrgId          int
}
