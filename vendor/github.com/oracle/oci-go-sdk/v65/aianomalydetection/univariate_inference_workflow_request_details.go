// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// UnivariateInferenceWorkflowRequestDetails The base class for the Univariate Inference Workflow call.
type UnivariateInferenceWorkflowRequestDetails interface {

	// Choose whether you would like the service to return all data points or just anomalies.
	GetAreAllDataPointsRequired() *bool

	GetTrainingRequestDetails() *UnivariateModelTrainingRequestDetails

	// Tune between precision and recall.
	GetSensitivity() *float32
}

type univariateinferenceworkflowrequestdetails struct {
	JsonData                 []byte
	AreAllDataPointsRequired *bool                                  `mandatory:"false" json:"areAllDataPointsRequired"`
	TrainingRequestDetails   *UnivariateModelTrainingRequestDetails `mandatory:"false" json:"trainingRequestDetails"`
	Sensitivity              *float32                               `mandatory:"false" json:"sensitivity"`
	RequestType              string                                 `json:"requestType"`
}

// UnmarshalJSON unmarshals json
func (m *univariateinferenceworkflowrequestdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerunivariateinferenceworkflowrequestdetails univariateinferenceworkflowrequestdetails
	s := struct {
		Model Unmarshalerunivariateinferenceworkflowrequestdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AreAllDataPointsRequired = s.Model.AreAllDataPointsRequired
	m.TrainingRequestDetails = s.Model.TrainingRequestDetails
	m.Sensitivity = s.Model.Sensitivity
	m.RequestType = s.Model.RequestType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *univariateinferenceworkflowrequestdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RequestType {
	case "BASE64_ENCODED":
		mm := EmbeddedUnivariateInferenceWorkflowRequestDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INLINE":
		mm := InlineUnivariateInferenceWorkflowRequestDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UnivariateInferenceWorkflowRequestDetails: %s.", m.RequestType)
		return *m, nil
	}
}

//GetAreAllDataPointsRequired returns AreAllDataPointsRequired
func (m univariateinferenceworkflowrequestdetails) GetAreAllDataPointsRequired() *bool {
	return m.AreAllDataPointsRequired
}

//GetTrainingRequestDetails returns TrainingRequestDetails
func (m univariateinferenceworkflowrequestdetails) GetTrainingRequestDetails() *UnivariateModelTrainingRequestDetails {
	return m.TrainingRequestDetails
}

//GetSensitivity returns Sensitivity
func (m univariateinferenceworkflowrequestdetails) GetSensitivity() *float32 {
	return m.Sensitivity
}

func (m univariateinferenceworkflowrequestdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m univariateinferenceworkflowrequestdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum Enum with underlying type: string
type UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum string

// Set of constants representing the allowable values for UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum
const (
	UnivariateInferenceWorkflowRequestDetailsRequestTypeInline        UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum = "INLINE"
	UnivariateInferenceWorkflowRequestDetailsRequestTypeBase64Encoded UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum = "BASE64_ENCODED"
)

var mappingUnivariateInferenceWorkflowRequestDetailsRequestTypeEnum = map[string]UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum{
	"INLINE":         UnivariateInferenceWorkflowRequestDetailsRequestTypeInline,
	"BASE64_ENCODED": UnivariateInferenceWorkflowRequestDetailsRequestTypeBase64Encoded,
}

var mappingUnivariateInferenceWorkflowRequestDetailsRequestTypeEnumLowerCase = map[string]UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum{
	"inline":         UnivariateInferenceWorkflowRequestDetailsRequestTypeInline,
	"base64_encoded": UnivariateInferenceWorkflowRequestDetailsRequestTypeBase64Encoded,
}

// GetUnivariateInferenceWorkflowRequestDetailsRequestTypeEnumValues Enumerates the set of values for UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum
func GetUnivariateInferenceWorkflowRequestDetailsRequestTypeEnumValues() []UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum {
	values := make([]UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum, 0)
	for _, v := range mappingUnivariateInferenceWorkflowRequestDetailsRequestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnivariateInferenceWorkflowRequestDetailsRequestTypeEnumStringValues Enumerates the set of values in String for UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum
func GetUnivariateInferenceWorkflowRequestDetailsRequestTypeEnumStringValues() []string {
	return []string{
		"INLINE",
		"BASE64_ENCODED",
	}
}

// GetMappingUnivariateInferenceWorkflowRequestDetailsRequestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnivariateInferenceWorkflowRequestDetailsRequestTypeEnum(val string) (UnivariateInferenceWorkflowRequestDetailsRequestTypeEnum, bool) {
	enum, ok := mappingUnivariateInferenceWorkflowRequestDetailsRequestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
