package common

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type Job struct {
	ID     string
	Status string
	Err    error
}

func NewJob() *Job {
	u1 := uuid.NewV4()
	return &Job{ID: u1.String(), Status: "running", Err: nil}
}

func (j *Job) GetStatus() (string, error) {
	if j != nil {
		return j.Status, j.Err
	}
	return "", nil
}

func (j *Job) SetFinish() {
	if j != nil {
		j.Status = "finish"
		j.Err = nil
	}
}

func (j *Job) SetErr(err error) {
	if j != nil {
		j.Status = "err"
		j.Err = err
	}
}

func (j *Job) CacheJob() (string, error) {
	var cache = make(map[string]interface{})
	cache["id"] = j.ID
	cache["status"] = j.Status
	if j.Err != nil {
		cache["err"] = j.Err.Error()
	} else {
		cache["err"] = nil
	}
	return j.ID, CacheJob(cache)
}

func (j *Job) GetCacheJob() error {
	data, err := GetJob(j.ID)
	if err != nil {
		return err
	}
	j.Status = data["status"]
	j.Err = errors.New(data["err"])
	return nil
}
