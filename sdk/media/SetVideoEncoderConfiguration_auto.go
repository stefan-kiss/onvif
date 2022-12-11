// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/stefan-kiss/onvif"
	"github.com/stefan-kiss/onvif/sdk"
	"github.com/stefan-kiss/onvif/media"
)

// Call_SetVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoEncoderConfigurationResponse.
func Call_SetVideoEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.SetVideoEncoderConfiguration) (media.SetVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoEncoderConfigurationResponse media.SetVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetVideoEncoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetVideoEncoderConfiguration")
		return reply.Body.SetVideoEncoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
