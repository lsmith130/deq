package eventstore

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/katcheCode/deq/api/v1/deq"
	"gitlab.com/katcheCode/deq/pkg/eventstore/data"
)

func TestWriteEvent(t *testing.T) {

	dir, err := ioutil.TempDir("", "test-write-event")
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	opts := badger.DefaultOptions
	opts.Dir = dir
	opts.ValueDir = dir
	db, err := badger.Open(opts)
	defer db.Close()
	if err != nil {
		t.Fatal("open db: ", err)
	}

	txn := db.NewTransaction(true)
	defer txn.Discard()

	// Setup existing channels - currently we have to ack an existing event on the
	// channels we want
	err = writeEvent(txn, &deq.Event{
		Topic:      "topic",
		Id:         "event0",
		CreateTime: time.Now().UnixNano(),
		Payload:    []byte{1},
	})
	if err != nil {
		t.Fatal("write event: ", err)
	}
	err = writeEvent(txn, &deq.Event{
		Topic:      "topic",
		Id:         "event00",
		CreateTime: time.Now().UnixNano(),
		Payload:    []byte{1},
	})
	if err != nil {
		t.Fatal("write event: ", err)
	}

	dequeuePayload := data.ChannelPayload{
		EventState: deq.EventState_DEQUEUED_OK,
	}

	channelKey := data.ChannelKey{
		ID:      "event0",
		Topic:   "topic",
		Channel: "channel",
	}
	err = setChannelEvent(txn, channelKey, dequeuePayload)
	if err != nil {
		t.Fatalf("set event state: %v", err)
	}
	channelKey.ID = "event00"
	err = setChannelEvent(txn, channelKey, dequeuePayload)
	if err != nil {
		t.Fatalf("set event state: %v", err)
	}
	channelKey.Channel = "channel2"
	err = setChannelEvent(txn, channelKey, dequeuePayload)
	if err != nil {
		t.Fatalf("set event state: %v", err)
	}

	// Write actual event
	expected := deq.Event{
		Topic:        "topic",
		Id:           "event1",
		CreateTime:   time.Now().UnixNano(),
		Payload:      []byte{1, 2, 3},
		DefaultState: deq.EventState_DEQUEUED_OK,
		// Should be ignored.
		State: deq.EventState_DEQUEUED_ERROR,
	}

	err = writeEvent(txn, &expected)
	if err != nil {
		t.Fatal("write event: ", err)
	}

	expected.State = deq.EventState_DEQUEUED_OK

	e, err := getEvent(txn, expected.Topic, expected.Id, "channel")
	if err != nil {
		t.Fatalf("get event on channel: %v", err)
	}
	if !proto.Equal(e, &expected) {
		t.Errorf("expected %v, got %v", expected, e)
	}
	e, err = getEvent(txn, expected.Topic, expected.Id, "channel2")
	if err != nil {
		t.Fatalf("get event on channel2: %v", err)
	}
	if !proto.Equal(e, &expected) {
		t.Errorf("expected %v, got %v", expected, e)
	}

	expected.State = deq.EventState_QUEUED

	e, err = getEvent(txn, expected.Topic, expected.Id, "newchannel")
	if err != nil {
		t.Fatalf("get event on newchannel: %v", err)
	}
	if !proto.Equal(e, &expected) {
		t.Errorf("expected %v, got %v", expected, e)
	}

	err = txn.Commit(nil)
	if err != nil {
		t.Error("commit: ", err)
	}
}

func BenchmarkWriteEvent(b *testing.B) {

	dir, err := ioutil.TempDir("", "test-write-event")
	if err != nil {
		b.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	opts := badger.DefaultOptions
	opts.Dir = dir
	opts.ValueDir = dir
	db, err := badger.Open(opts)
	defer db.Close()
	if err != nil {
		b.Fatal("open db: ", err)
	}

	txn := db.NewTransaction(true)
	defer txn.Discard()

	expected := &deq.Event{
		Topic:        "topic",
		Id:           "event1",
		CreateTime:   time.Now().UnixNano(),
		Payload:      []byte{1, 2, 3},
		DefaultState: deq.EventState_DEQUEUED_OK,
		// Should be ignored.
		State: deq.EventState_DEQUEUED_ERROR,
	}

	err = writeEvent(txn, expected)
	if err != nil {
		b.Fatal("write event: ", err)
	}

	err = txn.Commit(nil)
	if err != nil {
		b.Error("commit: ", err)
	}
}
