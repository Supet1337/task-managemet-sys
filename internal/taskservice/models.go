package models

import "time"

type Task struct {
	TaskId      int64     `json:"taskId"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	TimeStart   time.Time `json:"timeStart"`
	TimeEnd     time.Time `json:"timeEnd"`
	Description string    `json:"description"`
	//MembersId   []int64   `json:"membersId"`
}
