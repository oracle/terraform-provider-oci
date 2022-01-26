// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// EmbeddedDetectAnomaliesRequest The request body when the user selects to provide byte data in detect call which is Base64 encoded.
// The default type of the data is CSV and can be JSON by setting the 'contentType'.
type EmbeddedDetectAnomaliesRequest struct {

	// The OCID of the trained model。
	ModelId *string `mandatory:"true" json:"modelId"`

	Content []byte `mandatory:"true" json:"content"`

	ContentType EmbeddedDetectAnomaliesRequestContentTypeEnum `mandatory:"false" json:"contentType,omitempty"`
}

//GetModelId returns ModelId
func (m EmbeddedDetectAnomaliesRequest) GetModelId() *string {
	return m.ModelId
}

func (m EmbeddedDetectAnomaliesRequest) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m EmbeddedDetectAnomaliesRequest) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEmbeddedDetectAnomaliesRequest EmbeddedDetectAnomaliesRequest
	s := struct {
		DiscriminatorParam string `json:"requestType"`
		MarshalTypeEmbeddedDetectAnomaliesRequest
	}{
		"BASE64_ENCODED",
		(MarshalTypeEmbeddedDetectAnomaliesRequest)(m),
	}

	return json.Marshal(&s)
}

// EmbeddedDetectAnomaliesRequestContentTypeEnum Enum with underlying type: string
type EmbeddedDetectAnomaliesRequestContentTypeEnum string

// Set of constants representing the allowable values for EmbeddedDetectAnomaliesRequestContentTypeEnum
const (
	EmbeddedDetectAnomaliesRequestContentTypeCsv  EmbeddedDetectAnomaliesRequestContentTypeEnum = "CSV"
	EmbeddedDetectAnomaliesRequestContentTypeJson EmbeddedDetectAnomaliesRequestContentTypeEnum = "JSON"
)

var mappingEmbeddedDetectAnomaliesRequestContentType = map[string]EmbeddedDetectAnomaliesRequestContentTypeEnum{
	"CSV":  EmbeddedDetectAnomaliesRequestContentTypeCsv,
	"JSON": EmbeddedDetectAnomaliesRequestContentTypeJson,
}

// GetEmbeddedDetectAnomaliesRequestContentTypeEnumValues Enumerates the set of values for EmbeddedDetectAnomaliesRequestContentTypeEnum
func GetEmbeddedDetectAnomaliesRequestContentTypeEnumValues() []EmbeddedDetectAnomaliesRequestContentTypeEnum {
	values := make([]EmbeddedDetectAnomaliesRequestContentTypeEnum, 0)
	for _, v := range mappingEmbeddedDetectAnomaliesRequestContentType {
		values = append(values, v)
	}
	return values
}
