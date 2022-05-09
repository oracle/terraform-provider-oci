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

// OnlinePredictionUseCase Prediction contract for online prediction use-case
type OnlinePredictionUseCase struct {

	// Name of use case
	Name *string `mandatory:"true" json:"name"`

	// Type of authN/Z used for prediction endpoint.
	AuthType OnlinePredictionUseCaseAuthTypeEnum `mandatory:"true" json:"authType"`
}

//GetName returns Name
func (m OnlinePredictionUseCase) GetName() *string {
	return m.Name
}

func (m OnlinePredictionUseCase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OnlinePredictionUseCase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOnlinePredictionUseCaseAuthTypeEnum(string(m.AuthType)); !ok && m.AuthType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthType: %s. Supported values are: %s.", m.AuthType, strings.Join(GetOnlinePredictionUseCaseAuthTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OnlinePredictionUseCase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOnlinePredictionUseCase OnlinePredictionUseCase
	s := struct {
		DiscriminatorParam string `json:"predictionType"`
		MarshalTypeOnlinePredictionUseCase
	}{
		"ONLINE",
		(MarshalTypeOnlinePredictionUseCase)(m),
	}

	return json.Marshal(&s)
}

// OnlinePredictionUseCaseAuthTypeEnum Enum with underlying type: string
type OnlinePredictionUseCaseAuthTypeEnum string

// Set of constants representing the allowable values for OnlinePredictionUseCaseAuthTypeEnum
const (
	OnlinePredictionUseCaseAuthTypeIdcs OnlinePredictionUseCaseAuthTypeEnum = "IDCS"
)

var mappingOnlinePredictionUseCaseAuthTypeEnum = map[string]OnlinePredictionUseCaseAuthTypeEnum{
	"IDCS": OnlinePredictionUseCaseAuthTypeIdcs,
}

var mappingOnlinePredictionUseCaseAuthTypeEnumLowerCase = map[string]OnlinePredictionUseCaseAuthTypeEnum{
	"idcs": OnlinePredictionUseCaseAuthTypeIdcs,
}

// GetOnlinePredictionUseCaseAuthTypeEnumValues Enumerates the set of values for OnlinePredictionUseCaseAuthTypeEnum
func GetOnlinePredictionUseCaseAuthTypeEnumValues() []OnlinePredictionUseCaseAuthTypeEnum {
	values := make([]OnlinePredictionUseCaseAuthTypeEnum, 0)
	for _, v := range mappingOnlinePredictionUseCaseAuthTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOnlinePredictionUseCaseAuthTypeEnumStringValues Enumerates the set of values in String for OnlinePredictionUseCaseAuthTypeEnum
func GetOnlinePredictionUseCaseAuthTypeEnumStringValues() []string {
	return []string{
		"IDCS",
	}
}

// GetMappingOnlinePredictionUseCaseAuthTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOnlinePredictionUseCaseAuthTypeEnum(val string) (OnlinePredictionUseCaseAuthTypeEnum, bool) {
	enum, ok := mappingOnlinePredictionUseCaseAuthTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
