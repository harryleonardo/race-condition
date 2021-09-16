package transporter

import "github.com/nats-io/nats.go"

// NatsClient ...
type NatsClient struct {
	EncodedConnection *nats.EncodedConn
}
