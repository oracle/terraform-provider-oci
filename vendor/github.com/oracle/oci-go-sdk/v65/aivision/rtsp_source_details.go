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

// RtspSourceDetails Details of RtspDevice
type RtspSourceDetails struct {
	StreamNetworkAccessDetails StreamNetworkAccessDetails `mandatory:"true" json:"streamNetworkAccessDetails"`

	// url of camera
	CameraUrl *string `mandatory:"true" json:"cameraUrl"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of secret where credentials are stored in username:password format.
	SecretId *string `mandatory:"false" json:"secretId"`
}

// GetStreamNetworkAccessDetails returns StreamNetworkAccessDetails
func (m RtspSourceDetails) GetStreamNetworkAccessDetails() StreamNetworkAccessDetails {
	return m.StreamNetworkAccessDetails
}

func (m RtspSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RtspSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RtspSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRtspSourceDetails RtspSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeRtspSourceDetails
	}{
		"RTSP",
		(MarshalTypeRtspSourceDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *RtspSourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SecretId                   *string                    `json:"secretId"`
		StreamNetworkAccessDetails streamnetworkaccessdetails `json:"streamNetworkAccessDetails"`
		CameraUrl                  *string                    `json:"cameraUrl"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SecretId = model.SecretId

	nn, e = model.StreamNetworkAccessDetails.UnmarshalPolymorphicJSON(model.StreamNetworkAccessDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StreamNetworkAccessDetails = nn.(StreamNetworkAccessDetails)
	} else {
		m.StreamNetworkAccessDetails = nil
	}

	m.CameraUrl = model.CameraUrl

	return
}
