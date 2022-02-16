// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// MaskingReport A masking report contains information about a completed masking request. It includes details such as the target database masked,
// masking policy used, masking start and finish time, total number of schemas, tables, columns and values masked, masked columns, and the masking formats used.
type MaskingReport struct {

	// The OCID of the masking report.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the masking report.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the masking work request that resulted in this masking report.
	MaskingWorkRequestId *string `mandatory:"true" json:"maskingWorkRequestId"`

	// The OCID of the masking policy used.
	MaskingPolicyId *string `mandatory:"true" json:"maskingPolicyId"`

	// The OCID of the target database masked.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The total number of unique sensitive types associated with the masked columns.
	TotalMaskedSensitiveTypes *int64 `mandatory:"true" json:"totalMaskedSensitiveTypes"`

	// The total number of unique schemas that contain the masked columns.
	TotalMaskedSchemas *int64 `mandatory:"true" json:"totalMaskedSchemas"`

	// The total number of unique objects (tables and editioning views) that contain the masked columns.
	TotalMaskedObjects *int64 `mandatory:"true" json:"totalMaskedObjects"`

	// The total number of masked columns.
	TotalMaskedColumns *int64 `mandatory:"true" json:"totalMaskedColumns"`

	// The total number of masked values.
	TotalMaskedValues *int64 `mandatory:"true" json:"totalMaskedValues"`

	// The date and time data masking started, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeMaskingStarted *common.SDKTime `mandatory:"true" json:"timeMaskingStarted"`

	// The date and time data masking finished, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeMaskingFinished *common.SDKTime `mandatory:"true" json:"timeMaskingFinished"`
}

func (m MaskingReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
