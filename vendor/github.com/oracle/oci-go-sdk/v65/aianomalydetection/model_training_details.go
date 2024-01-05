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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelTrainingDetails Specifies the details of the MSET model during the create call.
type ModelTrainingDetails struct {

	// The list of OCIDs of the data assets to train the model. The dataAssets have to be in the same project where the ai model would reside.
	DataAssetIds []string `mandatory:"true" json:"dataAssetIds"`

	// User can choose specific algorithm for training.
	AlgorithmHint ModelTrainingDetailsAlgorithmHintEnum `mandatory:"false" json:"algorithmHint,omitempty"`

	// A target model accuracy metric user provides as their requirement
	TargetFap *float32 `mandatory:"false" json:"targetFap"`

	// Fraction of total data that is used for training the model. The remaining is used for validation of the model.
	TrainingFraction *float32 `mandatory:"false" json:"trainingFraction"`

	// This value would determine the window size of the training algorithm.
	WindowSize *int `mandatory:"false" json:"windowSize"`
}

func (m ModelTrainingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelTrainingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingModelTrainingDetailsAlgorithmHintEnum(string(m.AlgorithmHint)); !ok && m.AlgorithmHint != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AlgorithmHint: %s. Supported values are: %s.", m.AlgorithmHint, strings.Join(GetModelTrainingDetailsAlgorithmHintEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelTrainingDetailsAlgorithmHintEnum Enum with underlying type: string
type ModelTrainingDetailsAlgorithmHintEnum string

// Set of constants representing the allowable values for ModelTrainingDetailsAlgorithmHintEnum
const (
	ModelTrainingDetailsAlgorithmHintMultivariateMset ModelTrainingDetailsAlgorithmHintEnum = "MULTIVARIATE_MSET"
	ModelTrainingDetailsAlgorithmHintUnivariateOcsvm  ModelTrainingDetailsAlgorithmHintEnum = "UNIVARIATE_OCSVM"
)

var mappingModelTrainingDetailsAlgorithmHintEnum = map[string]ModelTrainingDetailsAlgorithmHintEnum{
	"MULTIVARIATE_MSET": ModelTrainingDetailsAlgorithmHintMultivariateMset,
	"UNIVARIATE_OCSVM":  ModelTrainingDetailsAlgorithmHintUnivariateOcsvm,
}

var mappingModelTrainingDetailsAlgorithmHintEnumLowerCase = map[string]ModelTrainingDetailsAlgorithmHintEnum{
	"multivariate_mset": ModelTrainingDetailsAlgorithmHintMultivariateMset,
	"univariate_ocsvm":  ModelTrainingDetailsAlgorithmHintUnivariateOcsvm,
}

// GetModelTrainingDetailsAlgorithmHintEnumValues Enumerates the set of values for ModelTrainingDetailsAlgorithmHintEnum
func GetModelTrainingDetailsAlgorithmHintEnumValues() []ModelTrainingDetailsAlgorithmHintEnum {
	values := make([]ModelTrainingDetailsAlgorithmHintEnum, 0)
	for _, v := range mappingModelTrainingDetailsAlgorithmHintEnum {
		values = append(values, v)
	}
	return values
}

// GetModelTrainingDetailsAlgorithmHintEnumStringValues Enumerates the set of values in String for ModelTrainingDetailsAlgorithmHintEnum
func GetModelTrainingDetailsAlgorithmHintEnumStringValues() []string {
	return []string{
		"MULTIVARIATE_MSET",
		"UNIVARIATE_OCSVM",
	}
}

// GetMappingModelTrainingDetailsAlgorithmHintEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelTrainingDetailsAlgorithmHintEnum(val string) (ModelTrainingDetailsAlgorithmHintEnum, bool) {
	enum, ok := mappingModelTrainingDetailsAlgorithmHintEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
