package deqdb

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"gitlab.com/katcheCode/deq"
)

func newTestDB() (*Store, func()) {
	dir, err := ioutil.TempDir("", "test-pub")
	if err != nil {
		panic("create temp dir: " + err.Error())
	}

	db, err := Open(Options{
		Dir:             dir,
		UpgradeIfNeeded: true,
	})
	if err != nil {
		panic("open db: " + err.Error())
	}

	return db, func() {
		db.Close()
		os.RemoveAll(dir)
	}
}

func TestDel(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	db, discard := newTestDB()
	defer discard()

	expected := deq.Event{
		ID:           "event1",
		Topic:        "topic",
		CreateTime:   time.Now(),
		DefaultState: deq.StateQueued,
		State:        deq.StateQueued,
	}

	_, err := db.Pub(ctx, expected)
	if err != nil {
		t.Fatalf("pub: %v", err)
	}

	err = db.Del(ctx, expected.Topic, expected.ID)
	if err != nil {
		t.Fatalf("del: %v", err)
	}

	channel := db.Channel("channel", expected.Topic)
	defer channel.Close()

	_, err = channel.Get(ctx, expected.ID)
	if err == nil {
		t.Fatalf("returned deleted event")
	}
	if err != deq.ErrNotFound {
		t.Fatalf("get deleted: %v", err)
	}
}

func TestPub(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	db, discard := newTestDB()
	defer discard()

	expected := deq.Event{
		ID:           "event1",
		Topic:        "topic",
		CreateTime:   time.Now(),
		DefaultState: deq.StateQueued,
		State:        deq.StateQueued,
	}

	channel := db.Channel("channel", expected.Topic)
	defer channel.Close()

	_, err := db.Pub(ctx, expected)
	if err != nil {
		t.Fatalf("pub: %v", err)
	}

	event, err := channel.Next(context.Background())
	if err != nil {
		t.Fatalf("get next: %v", err)
	}
	if !cmp.Equal(event, expected) {
		t.Errorf("get next:\n%s", cmp.Diff(expected, event))
	}

	expected.State = deq.StateQueued

	event, err = channel.Get(ctx, expected.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if !cmp.Equal(event, expected) {
		t.Errorf("get:\n%s", cmp.Diff(expected, event))
	}
}

func TestMassPub(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	db, discard := newTestDB()
	defer discard()

	topic := "topic"
	// Round(0) to discard's leap-second info that's lost in serialization
	createTime := time.Now().Round(0)

	expected := make([]deq.Event, 500)
	for i := 0; i < 500; i++ {
		expected[i] = deq.Event{
			ID:           fmt.Sprintf("event%03d", i),
			Topic:        topic,
			CreateTime:   createTime,
			DefaultState: deq.StateQueued,
			State:        deq.StateQueued,
		}
	}

	channel := db.Channel("channel", topic)
	defer channel.Close()

	for i, e := range expected {
		_, err := db.Pub(ctx, e)
		if err != nil {
			t.Fatalf("pub %d: %v", i, err)
		}
	}

	var actual []deq.Event
	iter := channel.NewEventIter(nil)
	for iter.Next(ctx) {
		actual = append(actual, iter.Event())
	}
	if iter.Err() != nil {
		t.Fatalf("iterate: %v", iter.Err())
	}
	if !cmp.Equal(expected, actual) {
		t.Errorf(":\n%s", cmp.Diff(expected, actual))
	}
}

func TestPubDuplicate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	db, discard := newTestDB()
	defer discard()

	expected := deq.Event{
		ID:           "event1",
		Topic:        "topic",
		CreateTime:   time.Now(),
		DefaultState: deq.StateQueued,
		State:        deq.StateQueued,
	}

	channel := db.Channel("channel", expected.Topic)
	defer channel.Close()

	// Publish the event
	_, err := db.Pub(ctx, expected)
	if err != nil {
		t.Fatalf("pub: %v", err)
	}

	// Publish and verify event with same id and payload
	_, err = db.Pub(ctx, expected)
	if err != nil {
		t.Fatalf("identifical duplicate pub: %v", err)
	}

	// Publish and verify event with same id and different payload
	expected.Payload = []byte{1}
	_, err = db.Pub(ctx, expected)
	if err != deq.ErrAlreadyExists {
		t.Fatalf("modified duplicate pub: %v", err)
	}
	expected.Payload = nil

	event, err := channel.Next(context.Background())
	if err != nil {
		t.Fatalf("get next: %v", err)
	}
	if !cmp.Equal(event, expected) {
		t.Errorf("get next:\n%s", cmp.Diff(expected, event))
	}

	expected.State = deq.StateQueued

	event, err = channel.Get(ctx, expected.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if !cmp.Equal(event, expected) {
		t.Errorf("get:\n%s", cmp.Diff(expected, event))
	}
}
