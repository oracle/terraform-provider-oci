// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CryptoAssessmentFindingSummaryMetrics Aggregate finding counts for the specified crypto assessment.
type CryptoAssessmentFindingSummaryMetrics struct {

	// Total findings across all statuses.
	TotalChecks *int `mandatory:"true" json:"totalChecks"`

	// Findings with FAIL or EVALUATE status.
	TotalFindings *int `mandatory:"true" json:"totalFindings"`

	// FAIL or EVALUATE findings with priority 1.
	Critical *int `mandatory:"true" json:"critical"`

	// FAIL or EVALUATE findings with priority 2.
	High *int `mandatory:"true" json:"high"`

	// FAIL or EVALUATE findings with priority 3.
	Med *int `mandatory:"true" json:"med"`

	// FAIL or EVALUATE findings with priority 4.
	Low *int `mandatory:"true" json:"low"`

	// Counts keyed by finding status. All supported statuses are included with a zero count when absent.
	StatusCounts map[string]int `mandatory:"true" json:"statusCounts"`

	NetworkEncryptionStatus *CryptoAssessmentFindingCategorySummary `mandatory:"true" json:"networkEncryptionStatus"`

	DataEncryptionStatus *CryptoAssessmentFindingCategorySummary `mandatory:"true" json:"dataEncryptionStatus"`

	WalletStatus *CryptoAssessmentFindingCategorySummary `mandatory:"true" json:"walletStatus"`

	BackupStatus *CryptoAssessmentFindingCategorySummary `mandatory:"true" json:"backupStatus"`
}

func (m CryptoAssessmentFindingSummaryMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentFindingSummaryMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
