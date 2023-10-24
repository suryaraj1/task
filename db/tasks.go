package main

import (
	"encoding/binary"
	"task/models"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var taskBucket = []byte("tasks")

func ListTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, models.Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil // to complete transaction
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
