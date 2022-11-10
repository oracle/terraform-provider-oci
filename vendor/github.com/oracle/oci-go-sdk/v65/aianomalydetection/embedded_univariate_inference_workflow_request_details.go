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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmbeddedUnivariateInferenceWorkflowRequestDetails The request body when the user selects to provide byte data in detect call which is Base64 encoded.
// The default type of the data is CSV and can be JSON by setting the 'contentType'.
type EmbeddedUnivariateInferenceWorkflowRequestDetails struct {

	// Choose whether you'd like the service to return all datapoints or just anomlies
	AreAllDataPointsRequired *bool `mandatory:"false" json:"areAllDataPointsRequired"`

	TrainingRequestDetails *UnivariateModelTrainingRequestDetails `mandatory:"false" json:"trainingRequestDetails"`

	// tune between precision and recall
	Sensitivity *float32 `mandatory:"false" json:"sensitivity"`

	// List of byte encoded files.
	Content [][]byte `mandatory:"false" json:"content"`

	ContentType EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum `mandatory:"false" json:"contentType,omitempty"`
}

//GetAreAllDataPointsRequired returns AreAllDataPointsRequired
func (m EmbeddedUnivariateInferenceWorkflowRequestDetails) GetAreAllDataPointsRequired() *bool {
	return m.AreAllDataPointsRequired
}

//GetTrainingRequestDetails returns TrainingRequestDetails
func (m EmbeddedUnivariateInferenceWorkflowRequestDetails) GetTrainingRequestDetails() *UnivariateModelTrainingRequestDetails {
	return m.TrainingRequestDetails
}

//GetSensitivity returns Sensitivity
func (m EmbeddedUnivariateInferenceWorkflowRequestDetails) GetSensitivity() *float32 {
	return m.Sensitivity
}

func (m EmbeddedUnivariateInferenceWorkflowRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmbeddedUnivariateInferenceWorkflowRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum(string(m.ContentType)); !ok && m.ContentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContentType: %s. Supported values are: %s.", m.ContentType, strings.Join(GetEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EmbeddedUnivariateInferenceWorkflowRequestDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEmbeddedUnivariateInferenceWorkflowRequestDetails EmbeddedUnivariateInferenceWorkflowRequestDetails
	s := struct {
		DiscriminatorParam string `json:"requestType"`
		MarshalTypeEmbeddedUnivariateInferenceWorkflowRequestDetails
	}{
		"BASE64_ENCODED",
		(MarshalTypeEmbeddedUnivariateInferenceWorkflowRequestDetails)(m),
	}

	return json.Marshal(&s)
}

// EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum Enum with underlying type: string
type EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum string

// Set of constants representing the allowable values for EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum
const (
	EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeCsv  EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum = "CSV"
	EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeJson EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum = "JSON"
)

var mappingEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum = map[string]EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum{
	"CSV":  EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeCsv,
	"JSON": EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeJson,
}

var mappingEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnumLowerCase = map[string]EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum{
	"csv":  EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeCsv,
	"json": EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeJson,
}

// GetEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnumValues Enumerates the set of values for EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum
func GetEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnumValues() []EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum {
	values := make([]EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum, 0)
	for _, v := range mappingEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnumStringValues Enumerates the set of values in String for EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum
func GetEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnumStringValues() []string {
	return []string{
		"CSV",
		"JSON",
	}
}

// GetMappingEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum(val string) (EmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnum, bool) {
	enum, ok := mappingEmbeddedUnivariateInferenceWorkflowRequestDetailsContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
