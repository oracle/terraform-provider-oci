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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ModelTrainingResults Specifies the details for an Anomaly Detection model trained with MSET.
type ModelTrainingResults struct {

	// The final-achieved model accuracy metric on individual value level
	Fap *float32 `mandatory:"true" json:"fap"`

	// The model accuracy metric on timestamp level.
	MultivariateFap *float32 `mandatory:"false" json:"multivariateFap"`

	// A boolean value to indicate if train goal/targetFap is achieved for trained model
	IsTrainingGoalAchieved *bool `mandatory:"false" json:"isTrainingGoalAchieved"`

	// A warning message to explain the reason when targetFap cannot be achieved for trained model
	Warning *string `mandatory:"false" json:"warning"`

	// The list of signal details.
	SignalDetails []PerSignalDetails `mandatory:"false" json:"signalDetails"`

	RowReductionDetails *RowReductionDetails `mandatory:"false" json:"rowReductionDetails"`
}

func (m ModelTrainingResults) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelTrainingResults) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
