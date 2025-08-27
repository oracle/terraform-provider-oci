// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScanResultSummary The result of a server check in a managed instance.
type ScanResultSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	WlsDomainId *string `mandatory:"true" json:"wlsDomainId"`

	// The name of the WebLogic server to which the server check belongs.
	ServerName *string `mandatory:"true" json:"serverName"`

	// The date when the WebLogic server health check is performed (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeOfServerCheck *common.SDKTime `mandatory:"true" json:"timeOfServerCheck"`

	// The status of the server check which is OK, FAILURE, or WARNING.
	ServerCheckStatus ScanResultSummaryServerCheckStatusEnum `mandatory:"true" json:"serverCheckStatus"`

	// The name of the check performed.
	ServerCheckName *string `mandatory:"true" json:"serverCheckName"`

	// The result of the server check.
	ServerCheckResult *string `mandatory:"true" json:"serverCheckResult"`

	// The identifier of the the server check result.
	ServerCheckResultId *string `mandatory:"false" json:"serverCheckResultId"`
}

func (m ScanResultSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScanResultSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScanResultSummaryServerCheckStatusEnum(string(m.ServerCheckStatus)); !ok && m.ServerCheckStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServerCheckStatus: %s. Supported values are: %s.", m.ServerCheckStatus, strings.Join(GetScanResultSummaryServerCheckStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScanResultSummaryServerCheckStatusEnum Enum with underlying type: string
type ScanResultSummaryServerCheckStatusEnum string

// Set of constants representing the allowable values for ScanResultSummaryServerCheckStatusEnum
const (
	ScanResultSummaryServerCheckStatusOk      ScanResultSummaryServerCheckStatusEnum = "OK"
	ScanResultSummaryServerCheckStatusWarning ScanResultSummaryServerCheckStatusEnum = "WARNING"
	ScanResultSummaryServerCheckStatusFailure ScanResultSummaryServerCheckStatusEnum = "FAILURE"
)

var mappingScanResultSummaryServerCheckStatusEnum = map[string]ScanResultSummaryServerCheckStatusEnum{
	"OK":      ScanResultSummaryServerCheckStatusOk,
	"WARNING": ScanResultSummaryServerCheckStatusWarning,
	"FAILURE": ScanResultSummaryServerCheckStatusFailure,
}

var mappingScanResultSummaryServerCheckStatusEnumLowerCase = map[string]ScanResultSummaryServerCheckStatusEnum{
	"ok":      ScanResultSummaryServerCheckStatusOk,
	"warning": ScanResultSummaryServerCheckStatusWarning,
	"failure": ScanResultSummaryServerCheckStatusFailure,
}

// GetScanResultSummaryServerCheckStatusEnumValues Enumerates the set of values for ScanResultSummaryServerCheckStatusEnum
func GetScanResultSummaryServerCheckStatusEnumValues() []ScanResultSummaryServerCheckStatusEnum {
	values := make([]ScanResultSummaryServerCheckStatusEnum, 0)
	for _, v := range mappingScanResultSummaryServerCheckStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetScanResultSummaryServerCheckStatusEnumStringValues Enumerates the set of values in String for ScanResultSummaryServerCheckStatusEnum
func GetScanResultSummaryServerCheckStatusEnumStringValues() []string {
	return []string{
		"OK",
		"WARNING",
		"FAILURE",
	}
}

// GetMappingScanResultSummaryServerCheckStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScanResultSummaryServerCheckStatusEnum(val string) (ScanResultSummaryServerCheckStatusEnum, bool) {
	enum, ok := mappingScanResultSummaryServerCheckStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
