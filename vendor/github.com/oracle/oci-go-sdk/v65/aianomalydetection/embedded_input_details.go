// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmbeddedInputDetails The request body when byte data is provided in detect call, which is Base64 encoded.
// The default type of the data is CSV and can be JSON by setting the 'contentType'.
type EmbeddedInputDetails struct {
	Content []byte `mandatory:"true" json:"content"`

	ContentType EmbeddedInputDetailsContentTypeEnum `mandatory:"true" json:"contentType"`
}

func (m EmbeddedInputDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmbeddedInputDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmbeddedInputDetailsContentTypeEnum(string(m.ContentType)); !ok && m.ContentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContentType: %s. Supported values are: %s.", m.ContentType, strings.Join(GetEmbeddedInputDetailsContentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EmbeddedInputDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEmbeddedInputDetails EmbeddedInputDetails
	s := struct {
		DiscriminatorParam string `json:"inputType"`
		MarshalTypeEmbeddedInputDetails
	}{
		"BASE64_ENCODED",
		(MarshalTypeEmbeddedInputDetails)(m),
	}

	return json.Marshal(&s)
}

// EmbeddedInputDetailsContentTypeEnum Enum with underlying type: string
type EmbeddedInputDetailsContentTypeEnum string

// Set of constants representing the allowable values for EmbeddedInputDetailsContentTypeEnum
const (
	EmbeddedInputDetailsContentTypeCsv  EmbeddedInputDetailsContentTypeEnum = "CSV"
	EmbeddedInputDetailsContentTypeJson EmbeddedInputDetailsContentTypeEnum = "JSON"
)

var mappingEmbeddedInputDetailsContentTypeEnum = map[string]EmbeddedInputDetailsContentTypeEnum{
	"CSV":  EmbeddedInputDetailsContentTypeCsv,
	"JSON": EmbeddedInputDetailsContentTypeJson,
}

var mappingEmbeddedInputDetailsContentTypeEnumLowerCase = map[string]EmbeddedInputDetailsContentTypeEnum{
	"csv":  EmbeddedInputDetailsContentTypeCsv,
	"json": EmbeddedInputDetailsContentTypeJson,
}

// GetEmbeddedInputDetailsContentTypeEnumValues Enumerates the set of values for EmbeddedInputDetailsContentTypeEnum
func GetEmbeddedInputDetailsContentTypeEnumValues() []EmbeddedInputDetailsContentTypeEnum {
	values := make([]EmbeddedInputDetailsContentTypeEnum, 0)
	for _, v := range mappingEmbeddedInputDetailsContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEmbeddedInputDetailsContentTypeEnumStringValues Enumerates the set of values in String for EmbeddedInputDetailsContentTypeEnum
func GetEmbeddedInputDetailsContentTypeEnumStringValues() []string {
	return []string{
		"CSV",
		"JSON",
	}
}

// GetMappingEmbeddedInputDetailsContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmbeddedInputDetailsContentTypeEnum(val string) (EmbeddedInputDetailsContentTypeEnum, bool) {
	enum, ok := mappingEmbeddedInputDetailsContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
