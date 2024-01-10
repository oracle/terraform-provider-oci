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

// StreamCdnConfigSection Base fields of the StreamCdnConfig configuration object.
type StreamCdnConfigSection interface {
}

type streamcdnconfigsection struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *streamcdnconfigsection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstreamcdnconfigsection streamcdnconfigsection
	s := struct {
		Model Unmarshalerstreamcdnconfigsection
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *streamcdnconfigsection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "AKAMAI_MANUAL":
		mm := AkamaiManualStreamCdnConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EDGE":
		mm := EdgeStreamCdnConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for StreamCdnConfigSection: %s.", m.Type)
		return *m, nil
	}
}

func (m streamcdnconfigsection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m streamcdnconfigsection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamCdnConfigSectionTypeEnum Enum with underlying type: string
type StreamCdnConfigSectionTypeEnum string

// Set of constants representing the allowable values for StreamCdnConfigSectionTypeEnum
const (
	StreamCdnConfigSectionTypeEdge         StreamCdnConfigSectionTypeEnum = "EDGE"
	StreamCdnConfigSectionTypeAkamaiManual StreamCdnConfigSectionTypeEnum = "AKAMAI_MANUAL"
)

var mappingStreamCdnConfigSectionTypeEnum = map[string]StreamCdnConfigSectionTypeEnum{
	"EDGE":          StreamCdnConfigSectionTypeEdge,
	"AKAMAI_MANUAL": StreamCdnConfigSectionTypeAkamaiManual,
}

var mappingStreamCdnConfigSectionTypeEnumLowerCase = map[string]StreamCdnConfigSectionTypeEnum{
	"edge":          StreamCdnConfigSectionTypeEdge,
	"akamai_manual": StreamCdnConfigSectionTypeAkamaiManual,
}

// GetStreamCdnConfigSectionTypeEnumValues Enumerates the set of values for StreamCdnConfigSectionTypeEnum
func GetStreamCdnConfigSectionTypeEnumValues() []StreamCdnConfigSectionTypeEnum {
	values := make([]StreamCdnConfigSectionTypeEnum, 0)
	for _, v := range mappingStreamCdnConfigSectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamCdnConfigSectionTypeEnumStringValues Enumerates the set of values in String for StreamCdnConfigSectionTypeEnum
func GetStreamCdnConfigSectionTypeEnumStringValues() []string {
	return []string{
		"EDGE",
		"AKAMAI_MANUAL",
	}
}

// GetMappingStreamCdnConfigSectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamCdnConfigSectionTypeEnum(val string) (StreamCdnConfigSectionTypeEnum, bool) {
	enum, ok := mappingStreamCdnConfigSectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
