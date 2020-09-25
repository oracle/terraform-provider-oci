// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// OutboundConnectorTestResults Diagnostic test results between mount target and LDAP server using outbound connector.
type OutboundConnectorTestResults struct {

	// Status of overall connection test against all endpoints.
	// Success represents all endpoints tested and functioning.
	// PartialSuccess represents all endpoints tested, some function properly, some do not.
	// Fail represents no endpoints are functioning properly, MT cannot communicate with LDAP server.
	OverallResult OutboundConnectorTestResultsOverallResultEnum `mandatory:"false" json:"overallResult,omitempty"`

	// Array of result of connecting to each endpoint
	ResultsDetail []TestOutboundConnectorResult `mandatory:"false" json:"resultsDetail"`
}

func (m OutboundConnectorTestResults) String() string {
	return common.PointerString(m)
}

// OutboundConnectorTestResultsOverallResultEnum Enum with underlying type: string
type OutboundConnectorTestResultsOverallResultEnum string

// Set of constants representing the allowable values for OutboundConnectorTestResultsOverallResultEnum
const (
	OutboundConnectorTestResultsOverallResultSuccess        OutboundConnectorTestResultsOverallResultEnum = "SUCCESS"
	OutboundConnectorTestResultsOverallResultPartialsuccess OutboundConnectorTestResultsOverallResultEnum = "PARTIALSUCCESS"
	OutboundConnectorTestResultsOverallResultFail           OutboundConnectorTestResultsOverallResultEnum = "FAIL"
)

var mappingOutboundConnectorTestResultsOverallResult = map[string]OutboundConnectorTestResultsOverallResultEnum{
	"SUCCESS":        OutboundConnectorTestResultsOverallResultSuccess,
	"PARTIALSUCCESS": OutboundConnectorTestResultsOverallResultPartialsuccess,
	"FAIL":           OutboundConnectorTestResultsOverallResultFail,
}

// GetOutboundConnectorTestResultsOverallResultEnumValues Enumerates the set of values for OutboundConnectorTestResultsOverallResultEnum
func GetOutboundConnectorTestResultsOverallResultEnumValues() []OutboundConnectorTestResultsOverallResultEnum {
	values := make([]OutboundConnectorTestResultsOverallResultEnum, 0)
	for _, v := range mappingOutboundConnectorTestResultsOverallResult {
		values = append(values, v)
	}
	return values
}
