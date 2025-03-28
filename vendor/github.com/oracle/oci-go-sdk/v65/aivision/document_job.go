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

// DocumentJob The job details for a batch document analysis.
type DocumentJob struct {

	// The job id.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that starts the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of requested document analysis types.
	Features []DocumentFeature `mandatory:"true" json:"features"`

	// The job acceptance time.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// The current state of the batch document job.
	LifecycleState DocumentJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The document job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The document language, abbreviated according to ISO 639-2.
	Language DocumentLanguageEnum `mandatory:"false" json:"language,omitempty"`

	// The type of document.
	DocumentType DocumentTypeEnum `mandatory:"false" json:"documentType,omitempty"`

	InputLocation InputLocation `mandatory:"false" json:"inputLocation"`

	// The job start time.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The job finish time.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// How much progress the operation has made, compared to the total amount of work to be performed.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// Whether or not to generate a ZIP file containing the results.
	IsZipOutputEnabled *bool `mandatory:"false" json:"isZipOutputEnabled"`

	// The detailed status of FAILED state.
	LifecycleDetails DocumentJobLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

func (m DocumentJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DocumentJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDocumentJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDocumentJobLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDocumentLanguageEnum(string(m.Language)); !ok && m.Language != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Language: %s. Supported values are: %s.", m.Language, strings.Join(GetDocumentLanguageEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDocumentTypeEnum(string(m.DocumentType)); !ok && m.DocumentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DocumentType: %s. Supported values are: %s.", m.DocumentType, strings.Join(GetDocumentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDocumentJobLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetDocumentJobLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DocumentJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName        *string                         `json:"displayName"`
		Language           DocumentLanguageEnum            `json:"language"`
		DocumentType       DocumentTypeEnum                `json:"documentType"`
		InputLocation      inputlocation                   `json:"inputLocation"`
		TimeStarted        *common.SDKTime                 `json:"timeStarted"`
		TimeFinished       *common.SDKTime                 `json:"timeFinished"`
		PercentComplete    *float32                        `json:"percentComplete"`
		IsZipOutputEnabled *bool                           `json:"isZipOutputEnabled"`
		LifecycleDetails   DocumentJobLifecycleDetailsEnum `json:"lifecycleDetails"`
		Id                 *string                         `json:"id"`
		CompartmentId      *string                         `json:"compartmentId"`
		Features           []documentfeature               `json:"features"`
		TimeAccepted       *common.SDKTime                 `json:"timeAccepted"`
		OutputLocation     *OutputLocation                 `json:"outputLocation"`
		LifecycleState     DocumentJobLifecycleStateEnum   `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Language = model.Language

	m.DocumentType = model.DocumentType

	nn, e = model.InputLocation.UnmarshalPolymorphicJSON(model.InputLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InputLocation = nn.(InputLocation)
	} else {
		m.InputLocation = nil
	}

	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.PercentComplete = model.PercentComplete

	m.IsZipOutputEnabled = model.IsZipOutputEnabled

	m.LifecycleDetails = model.LifecycleDetails

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.Features = make([]DocumentFeature, len(model.Features))
	for i, n := range model.Features {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Features[i] = nn.(DocumentFeature)
		} else {
			m.Features[i] = nil
		}
	}
	m.TimeAccepted = model.TimeAccepted

	m.OutputLocation = model.OutputLocation

	m.LifecycleState = model.LifecycleState

	return
}

// DocumentJobLifecycleStateEnum Enum with underlying type: string
type DocumentJobLifecycleStateEnum string

// Set of constants representing the allowable values for DocumentJobLifecycleStateEnum
const (
	DocumentJobLifecycleStateSucceeded  DocumentJobLifecycleStateEnum = "SUCCEEDED"
	DocumentJobLifecycleStateFailed     DocumentJobLifecycleStateEnum = "FAILED"
	DocumentJobLifecycleStateAccepted   DocumentJobLifecycleStateEnum = "ACCEPTED"
	DocumentJobLifecycleStateCanceled   DocumentJobLifecycleStateEnum = "CANCELED"
	DocumentJobLifecycleStateInProgress DocumentJobLifecycleStateEnum = "IN_PROGRESS"
	DocumentJobLifecycleStateCanceling  DocumentJobLifecycleStateEnum = "CANCELING"
)

var mappingDocumentJobLifecycleStateEnum = map[string]DocumentJobLifecycleStateEnum{
	"SUCCEEDED":   DocumentJobLifecycleStateSucceeded,
	"FAILED":      DocumentJobLifecycleStateFailed,
	"ACCEPTED":    DocumentJobLifecycleStateAccepted,
	"CANCELED":    DocumentJobLifecycleStateCanceled,
	"IN_PROGRESS": DocumentJobLifecycleStateInProgress,
	"CANCELING":   DocumentJobLifecycleStateCanceling,
}

var mappingDocumentJobLifecycleStateEnumLowerCase = map[string]DocumentJobLifecycleStateEnum{
	"succeeded":   DocumentJobLifecycleStateSucceeded,
	"failed":      DocumentJobLifecycleStateFailed,
	"accepted":    DocumentJobLifecycleStateAccepted,
	"canceled":    DocumentJobLifecycleStateCanceled,
	"in_progress": DocumentJobLifecycleStateInProgress,
	"canceling":   DocumentJobLifecycleStateCanceling,
}

// GetDocumentJobLifecycleStateEnumValues Enumerates the set of values for DocumentJobLifecycleStateEnum
func GetDocumentJobLifecycleStateEnumValues() []DocumentJobLifecycleStateEnum {
	values := make([]DocumentJobLifecycleStateEnum, 0)
	for _, v := range mappingDocumentJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDocumentJobLifecycleStateEnumStringValues Enumerates the set of values in String for DocumentJobLifecycleStateEnum
func GetDocumentJobLifecycleStateEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"ACCEPTED",
		"CANCELED",
		"IN_PROGRESS",
		"CANCELING",
	}
}

// GetMappingDocumentJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDocumentJobLifecycleStateEnum(val string) (DocumentJobLifecycleStateEnum, bool) {
	enum, ok := mappingDocumentJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DocumentJobLifecycleDetailsEnum Enum with underlying type: string
type DocumentJobLifecycleDetailsEnum string

// Set of constants representing the allowable values for DocumentJobLifecycleDetailsEnum
const (
	DocumentJobLifecycleDetailsPartiallySucceeded DocumentJobLifecycleDetailsEnum = "PARTIALLY_SUCCEEDED"
	DocumentJobLifecycleDetailsCompletelyFailed   DocumentJobLifecycleDetailsEnum = "COMPLETELY_FAILED"
)

var mappingDocumentJobLifecycleDetailsEnum = map[string]DocumentJobLifecycleDetailsEnum{
	"PARTIALLY_SUCCEEDED": DocumentJobLifecycleDetailsPartiallySucceeded,
	"COMPLETELY_FAILED":   DocumentJobLifecycleDetailsCompletelyFailed,
}

var mappingDocumentJobLifecycleDetailsEnumLowerCase = map[string]DocumentJobLifecycleDetailsEnum{
	"partially_succeeded": DocumentJobLifecycleDetailsPartiallySucceeded,
	"completely_failed":   DocumentJobLifecycleDetailsCompletelyFailed,
}

// GetDocumentJobLifecycleDetailsEnumValues Enumerates the set of values for DocumentJobLifecycleDetailsEnum
func GetDocumentJobLifecycleDetailsEnumValues() []DocumentJobLifecycleDetailsEnum {
	values := make([]DocumentJobLifecycleDetailsEnum, 0)
	for _, v := range mappingDocumentJobLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetDocumentJobLifecycleDetailsEnumStringValues Enumerates the set of values in String for DocumentJobLifecycleDetailsEnum
func GetDocumentJobLifecycleDetailsEnumStringValues() []string {
	return []string{
		"PARTIALLY_SUCCEEDED",
		"COMPLETELY_FAILED",
	}
}

// GetMappingDocumentJobLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDocumentJobLifecycleDetailsEnum(val string) (DocumentJobLifecycleDetailsEnum, bool) {
	enum, ok := mappingDocumentJobLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
