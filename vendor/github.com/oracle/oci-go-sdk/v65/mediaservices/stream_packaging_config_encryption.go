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

// StreamPackagingConfigEncryption The encryption used by the stream packaging configuration.
type StreamPackagingConfigEncryption interface {
}

type streampackagingconfigencryption struct {
	JsonData  []byte
	Algorithm string `json:"algorithm"`
}

// UnmarshalJSON unmarshals json
func (m *streampackagingconfigencryption) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstreampackagingconfigencryption streampackagingconfigencryption
	s := struct {
		Model Unmarshalerstreampackagingconfigencryption
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Algorithm = s.Model.Algorithm

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *streampackagingconfigencryption) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Algorithm {
	case "AES128":
		mm := StreamPackagingConfigEncryptionAes128{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := StreamPackagingConfigEncryptionNone{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for StreamPackagingConfigEncryption: %s.", m.Algorithm)
		return *m, nil
	}
}

func (m streampackagingconfigencryption) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m streampackagingconfigencryption) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamPackagingConfigEncryptionAlgorithmEnum Enum with underlying type: string
type StreamPackagingConfigEncryptionAlgorithmEnum string

// Set of constants representing the allowable values for StreamPackagingConfigEncryptionAlgorithmEnum
const (
	StreamPackagingConfigEncryptionAlgorithmNone   StreamPackagingConfigEncryptionAlgorithmEnum = "NONE"
	StreamPackagingConfigEncryptionAlgorithmAes128 StreamPackagingConfigEncryptionAlgorithmEnum = "AES128"
)

var mappingStreamPackagingConfigEncryptionAlgorithmEnum = map[string]StreamPackagingConfigEncryptionAlgorithmEnum{
	"NONE":   StreamPackagingConfigEncryptionAlgorithmNone,
	"AES128": StreamPackagingConfigEncryptionAlgorithmAes128,
}

var mappingStreamPackagingConfigEncryptionAlgorithmEnumLowerCase = map[string]StreamPackagingConfigEncryptionAlgorithmEnum{
	"none":   StreamPackagingConfigEncryptionAlgorithmNone,
	"aes128": StreamPackagingConfigEncryptionAlgorithmAes128,
}

// GetStreamPackagingConfigEncryptionAlgorithmEnumValues Enumerates the set of values for StreamPackagingConfigEncryptionAlgorithmEnum
func GetStreamPackagingConfigEncryptionAlgorithmEnumValues() []StreamPackagingConfigEncryptionAlgorithmEnum {
	values := make([]StreamPackagingConfigEncryptionAlgorithmEnum, 0)
	for _, v := range mappingStreamPackagingConfigEncryptionAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamPackagingConfigEncryptionAlgorithmEnumStringValues Enumerates the set of values in String for StreamPackagingConfigEncryptionAlgorithmEnum
func GetStreamPackagingConfigEncryptionAlgorithmEnumStringValues() []string {
	return []string{
		"NONE",
		"AES128",
	}
}

// GetMappingStreamPackagingConfigEncryptionAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamPackagingConfigEncryptionAlgorithmEnum(val string) (StreamPackagingConfigEncryptionAlgorithmEnum, bool) {
	enum, ok := mappingStreamPackagingConfigEncryptionAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
