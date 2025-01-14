// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PipelineInitializationStep The step and its progress based on the recipe type.
type PipelineInitializationStep struct {

	// An object's Display Name.
	Name *string `mandatory:"true" json:"name"`

	// Status of the steps in a recipe. This option applies during pipeline initialization.
	Status StepStatusTypeEnum `mandatory:"true" json:"status"`

	// Shows the percentage complete of each recipe step during pipeline initialization.
	PercentComplete *int `mandatory:"true" json:"percentComplete"`

	// The date and time the request was started. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The date and time the request was finished. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeFinished *common.SDKTime `mandatory:"true" json:"timeFinished"`

	// The list of messages for each step while running.
	Messages []StepMessage `mandatory:"false" json:"messages"`
}

func (m PipelineInitializationStep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineInitializationStep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStepStatusTypeEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetStepStatusTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
