// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LiveKitWebrtcSourceDetails Details of Live kit Peer
type LiveKitWebrtcSourceDetails struct {
	StreamNetworkAccessDetails StreamNetworkAccessDetails `mandatory:"true" json:"streamNetworkAccessDetails"`

	// peer id of device
	PeerId *string `mandatory:"true" json:"peerId"`

	// Url for room
	RoomUrl *string `mandatory:"true" json:"roomUrl"`

	// name of the room
	RoomName *string `mandatory:"true" json:"roomName"`

	// User generated auth token to access the stream
	Token *string `mandatory:"false" json:"token"`
}

// GetStreamNetworkAccessDetails returns StreamNetworkAccessDetails
func (m LiveKitWebrtcSourceDetails) GetStreamNetworkAccessDetails() StreamNetworkAccessDetails {
	return m.StreamNetworkAccessDetails
}

func (m LiveKitWebrtcSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LiveKitWebrtcSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LiveKitWebrtcSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLiveKitWebrtcSourceDetails LiveKitWebrtcSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeLiveKitWebrtcSourceDetails
	}{
		"LIVEKIT_WEBRTC",
		(MarshalTypeLiveKitWebrtcSourceDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *LiveKitWebrtcSourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Token                      *string                    `json:"token"`
		StreamNetworkAccessDetails streamnetworkaccessdetails `json:"streamNetworkAccessDetails"`
		PeerId                     *string                    `json:"peerId"`
		RoomUrl                    *string                    `json:"roomUrl"`
		RoomName                   *string                    `json:"roomName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Token = model.Token

	nn, e = model.StreamNetworkAccessDetails.UnmarshalPolymorphicJSON(model.StreamNetworkAccessDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StreamNetworkAccessDetails = nn.(StreamNetworkAccessDetails)
	} else {
		m.StreamNetworkAccessDetails = nil
	}

	m.PeerId = model.PeerId

	m.RoomUrl = model.RoomUrl

	m.RoomName = model.RoomName

	return
}
