// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateStreamPackagingConfigDetails The information about the new Packaging Configuration.
type CreateStreamPackagingConfigDetails struct {

	// Unique identifier of the Distribution Channel that this stream packaging configuration belongs to.
	DistributionChannelId *string `mandatory:"true" json:"distributionChannelId"`

	// The name of the stream Packaging Configuration. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The output format for the package.
	StreamPackagingFormat CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum `mandatory:"true" json:"streamPackagingFormat"`

	// The duration in seconds for each fragment.
	SegmentTimeInSeconds *int `mandatory:"true" json:"segmentTimeInSeconds"`

	Encryption StreamPackagingConfigEncryption `mandatory:"false" json:"encryption"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateStreamPackagingConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateStreamPackagingConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateStreamPackagingConfigDetailsStreamPackagingFormatEnum(string(m.StreamPackagingFormat)); !ok && m.StreamPackagingFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StreamPackagingFormat: %s. Supported values are: %s.", m.StreamPackagingFormat, strings.Join(GetCreateStreamPackagingConfigDetailsStreamPackagingFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateStreamPackagingConfigDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Encryption            streampackagingconfigencryption                             `json:"encryption"`
		FreeformTags          map[string]string                                           `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}                           `json:"definedTags"`
		DistributionChannelId *string                                                     `json:"distributionChannelId"`
		DisplayName           *string                                                     `json:"displayName"`
		StreamPackagingFormat CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum `json:"streamPackagingFormat"`
		SegmentTimeInSeconds  *int                                                        `json:"segmentTimeInSeconds"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Encryption.UnmarshalPolymorphicJSON(model.Encryption.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Encryption = nn.(StreamPackagingConfigEncryption)
	} else {
		m.Encryption = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DistributionChannelId = model.DistributionChannelId

	m.DisplayName = model.DisplayName

	m.StreamPackagingFormat = model.StreamPackagingFormat

	m.SegmentTimeInSeconds = model.SegmentTimeInSeconds

	return
}

// CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum Enum with underlying type: string
type CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum string

// Set of constants representing the allowable values for CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum
const (
	CreateStreamPackagingConfigDetailsStreamPackagingFormatHls  CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum = "HLS"
	CreateStreamPackagingConfigDetailsStreamPackagingFormatDash CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum = "DASH"
)

var mappingCreateStreamPackagingConfigDetailsStreamPackagingFormatEnum = map[string]CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum{
	"HLS":  CreateStreamPackagingConfigDetailsStreamPackagingFormatHls,
	"DASH": CreateStreamPackagingConfigDetailsStreamPackagingFormatDash,
}

var mappingCreateStreamPackagingConfigDetailsStreamPackagingFormatEnumLowerCase = map[string]CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum{
	"hls":  CreateStreamPackagingConfigDetailsStreamPackagingFormatHls,
	"dash": CreateStreamPackagingConfigDetailsStreamPackagingFormatDash,
}

// GetCreateStreamPackagingConfigDetailsStreamPackagingFormatEnumValues Enumerates the set of values for CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum
func GetCreateStreamPackagingConfigDetailsStreamPackagingFormatEnumValues() []CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum {
	values := make([]CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum, 0)
	for _, v := range mappingCreateStreamPackagingConfigDetailsStreamPackagingFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateStreamPackagingConfigDetailsStreamPackagingFormatEnumStringValues Enumerates the set of values in String for CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum
func GetCreateStreamPackagingConfigDetailsStreamPackagingFormatEnumStringValues() []string {
	return []string{
		"HLS",
		"DASH",
	}
}

// GetMappingCreateStreamPackagingConfigDetailsStreamPackagingFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateStreamPackagingConfigDetailsStreamPackagingFormatEnum(val string) (CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum, bool) {
	enum, ok := mappingCreateStreamPackagingConfigDetailsStreamPackagingFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
