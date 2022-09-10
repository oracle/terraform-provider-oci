// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PredictionUseCase Prediction contract for particular use-case
type PredictionUseCase interface {

	// Name of use case
	GetName() *string
}

type predictionusecase struct {
	JsonData       []byte
	Name           *string `mandatory:"true" json:"name"`
	PredictionType string  `json:"predictionType"`
}

// UnmarshalJSON unmarshals json
func (m *predictionusecase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpredictionusecase predictionusecase
	s := struct {
		Model Unmarshalerpredictionusecase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.PredictionType = s.Model.PredictionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *predictionusecase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PredictionType {
	case "BATCH":
		mm := BatchPredictionUseCase{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ONLINE":
		mm := OnlinePredictionUseCase{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PredictionUseCase: %s.", m.PredictionType)
		return *m, nil
	}
}

//GetName returns Name
func (m predictionusecase) GetName() *string {
	return m.Name
}

func (m predictionusecase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m predictionusecase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PredictionUseCasePredictionTypeEnum Enum with underlying type: string
type PredictionUseCasePredictionTypeEnum string

// Set of constants representing the allowable values for PredictionUseCasePredictionTypeEnum
const (
	PredictionUseCasePredictionTypeOnline PredictionUseCasePredictionTypeEnum = "ONLINE"
	PredictionUseCasePredictionTypeBatch  PredictionUseCasePredictionTypeEnum = "BATCH"
	PredictionUseCasePredictionTypeStream PredictionUseCasePredictionTypeEnum = "STREAM"
)

var mappingPredictionUseCasePredictionTypeEnum = map[string]PredictionUseCasePredictionTypeEnum{
	"ONLINE": PredictionUseCasePredictionTypeOnline,
	"BATCH":  PredictionUseCasePredictionTypeBatch,
	"STREAM": PredictionUseCasePredictionTypeStream,
}

var mappingPredictionUseCasePredictionTypeEnumLowerCase = map[string]PredictionUseCasePredictionTypeEnum{
	"online": PredictionUseCasePredictionTypeOnline,
	"batch":  PredictionUseCasePredictionTypeBatch,
	"stream": PredictionUseCasePredictionTypeStream,
}

// GetPredictionUseCasePredictionTypeEnumValues Enumerates the set of values for PredictionUseCasePredictionTypeEnum
func GetPredictionUseCasePredictionTypeEnumValues() []PredictionUseCasePredictionTypeEnum {
	values := make([]PredictionUseCasePredictionTypeEnum, 0)
	for _, v := range mappingPredictionUseCasePredictionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPredictionUseCasePredictionTypeEnumStringValues Enumerates the set of values in String for PredictionUseCasePredictionTypeEnum
func GetPredictionUseCasePredictionTypeEnumStringValues() []string {
	return []string{
		"ONLINE",
		"BATCH",
		"STREAM",
	}
}

// GetMappingPredictionUseCasePredictionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPredictionUseCasePredictionTypeEnum(val string) (PredictionUseCasePredictionTypeEnum, bool) {
	enum, ok := mappingPredictionUseCasePredictionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
