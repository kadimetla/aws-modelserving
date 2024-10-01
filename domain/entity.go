package domain

import "time"

// SQSMessage represents the core entity of a message received from SQS.
type SQSMessage struct {
    ID      string
    Body    string
    SentAt  time.Time
}
