package nats

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber-schemas/build/go/protos/records"

	"github.com/batchcorp/plumber/util"
)

func (n *Nats) Write(_ context.Context, writeOpts *opts.WriteOptions, errorCh chan *records.ErrorRecord, messages ...*records.WriteRecord) error {
	if err := validateWriteOptions(writeOpts); err != nil {
		return errors.Wrap(err, "unable to validate write options")
	}

	defer n.Client.Close()

	subject := writeOpts.Nats.Args.Subject

	for _, msg := range messages {
		if err := n.Client.Publish(subject, []byte(msg.Input)); err != nil {
			util.WriteError(n.log, errorCh, fmt.Errorf("unable to publish message to subject '%s': %s", subject, err))
			continue
		}

		n.log.Infof("Successfully wrote message to '%s'", subject)
		return nil
	}

	return nil
}

func validateWriteOptions(writeOpts *opts.WriteOptions) error {
	if writeOpts == nil {
		return errors.New("write options cannot be nil")
	}

	if writeOpts.Nats == nil {
		return errors.New("backend group options cannot be nil")
	}

	if writeOpts.Nats.Args == nil {
		return errors.New("backend arg options cannot be nil")
	}

	if writeOpts.Nats.Args.Subject == "" {
		return ErrMissingSubject
	}

	return nil
}