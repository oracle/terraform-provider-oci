// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ExternalPublicationValidationSummary The external publication validation summary contains the validation summary information and the definition of the external object.
type ExternalPublicationValidationSummary struct {

	// Total number of validation messages.
	TotalMessageCount *int `mandatory:"false" json:"totalMessageCount"`

	// Total number of validation error messages.
	ErrorMessageCount *int `mandatory:"false" json:"errorMessageCount"`

	// Total number of validation warning messages.
	WarnMessageCount *int `mandatory:"false" json:"warnMessageCount"`

	// Total number of validation information messages.
	InfoMessageCount *int `mandatory:"false" json:"infoMessageCount"`

	// Detailed information of the data flow object validation.
	ValidationMessages map[string][]ValidationMessage `mandatory:"false" json:"validationMessages"`

	// Objects use a 36 character key as unique ID. It is system generated and cannot be modified.
	Key *string `mandatory:"false" json:"key"`
}

func (m ExternalPublicationValidationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalPublicationValidationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
