package tasks

import (
	"fmt"
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestCRUD(t *testing.T) {
	//connect to server
	sess, err := mgo.Dial("localhost:27017")
	if err != nil {
		t.Fatalf("error dialing Mango: %v", err)
	}
	defer sess.Close()
	store := &MongoStore{
		Session:        sess,
		DatabaseName:   "test",
		CollectionName: "tasks",
	}
	//make new task
	newtask := &NewTask{
		Title: "Learn MongoDB",
		Tags:  []string{"mongo", "info344", "hello world"},
	}
	//insert into mango db
	task, err := store.Insert(newtask)
	if err != nil {
		t.Errorf("error inserting new task: %v", err)
	}
	fmt.Println(task.ID)
	//how to Get
	task2, err := store.Get(task.ID)
	if err != nil {
		t.Errorf("error fetching task: %v", err)
	}
	if task2.Title != task.Title {
		t.Errorf("task title did not match, expected %s but got %s", task.Title, task2.Title)
	}
	sess.DB(store.DatabaseName).C(store.CollectionName).RemoveAll(nil)
}
