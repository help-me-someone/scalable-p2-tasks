package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeVideoSave = "video:save"
)

type VideoSavePayload struct {
	UserID    string
	VideoName string
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewVideoSaveTask(userID string, videoName string) (*asynq.Task, error) {
	payload, err := json.Marshal(VideoSavePayload{UserID: userID, VideoName: videoName})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeVideoSave, payload), nil
}

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface.
//---------------------------------------------------------------

func HandleVideoSaveTask(ctx context.Context, t *asynq.Task) error {
	var p VideoSavePayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// Logic goes here...

	return nil
}
