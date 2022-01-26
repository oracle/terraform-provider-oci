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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ModelTrainingDetails Specifies the details of the MSET model during the create call.
type ModelTrainingDetails struct {

	// The list of OCIDs of the data assets to train the model. The dataAssets have to be in the same project where the ai model would reside.
	DataAssetIds []string `mandatory:"true" json:"dataAssetIds"`

	// A target model accuracy metric user provides as their requirement
	TargetFap *float32 `mandatory:"false" json:"targetFap"`

	// Fraction of total data that is used for training the model. The remaining is used for validation of the model.
	TrainingFraction *float32 `mandatory:"false" json:"trainingFraction"`
}

func (m ModelTrainingDetails) String() string {
	return common.PointerString(m)
}
