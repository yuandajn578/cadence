// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dynamodb

import (
	"context"

	"github.com/uber/cadence/common/persistence"
	"github.com/uber/cadence/common/persistence/nosql/nosqlplugin"
)

//Insert message into queue, return error if failed or already exists
// Return ConditionFailure if the condition doesn't meet
func (db *ddb) InsertIntoQueue(
	ctx context.Context,
	row *nosqlplugin.QueueMessageRow,
) error {
	panic("TODO")
}

// Get the ID of last message inserted into the queue
func (db *ddb) SelectLastEnqueuedMessageID(
	ctx context.Context,
	queueType persistence.QueueType,
) (int64, error) {
	panic("TODO")
}

// Read queue messages starting from the exclusiveBeginMessageID
func (db *ddb) SelectMessagesFrom(
	ctx context.Context,
	queueType persistence.QueueType,
	exclusiveBeginMessageID int64,
	maxRows int,
) ([]*nosqlplugin.QueueMessageRow, error) {
	panic("TODO")
}

// Read queue message starting from exclusiveBeginMessageID int64, inclusiveEndMessageID int64
func (db *ddb) SelectMessagesBetween(
	ctx context.Context,
	request nosqlplugin.SelectMessagesBetweenRequest,
) (*nosqlplugin.SelectMessagesBetweenResponse, error) {
	panic("TODO")
}

// Delete all messages before exclusiveBeginMessageID
func (db *ddb) DeleteMessagesBefore(
	ctx context.Context,
	queueType persistence.QueueType,
	exclusiveBeginMessageID int64,
) error {
	panic("TODO")
}

// Delete all messages in a range between exclusiveBeginMessageID and inclusiveEndMessageID
func (db *ddb) DeleteMessagesInRange(
	ctx context.Context,
	queueType persistence.QueueType,
	exclusiveBeginMessageID int64,
	inclusiveEndMessageID int64,
) error {
	panic("TODO")
}

// Delete one message
func (db *ddb) DeleteMessage(
	ctx context.Context,
	queueType persistence.QueueType,
	messageID int64,
) error {
	panic("TODO")
}

// Insert an empty metadata row, starting from a version
func (db *ddb) InsertQueueMetadata(
	ctx context.Context,
	queueType persistence.QueueType,
	version int64,
) error {
	panic("TODO")
}

// **Conditionally** update a queue metadata row, if current version is matched(meaning current == row.Version - 1),
// then the current version will increase by one when updating the metadata row
// Return ConditionFailure if the condition doesn't meet
func (db *ddb) UpdateQueueMetadataCas(
	ctx context.Context,
	row nosqlplugin.QueueMetadataRow,
) error {
	panic("TODO")
}

// Read a QueueMetadata
func (db *ddb) SelectQueueMetadata(
	ctx context.Context,
	queueType persistence.QueueType,
) (*nosqlplugin.QueueMetadataRow, error) {
	panic("TODO")
}

func (db *ddb) GetQueueSize(
	ctx context.Context,
	queueType persistence.QueueType,
) (int64, error) {
	panic("TODO")
}
