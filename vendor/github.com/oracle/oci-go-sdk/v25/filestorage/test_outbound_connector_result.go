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

// TestOutboundConnectorResult Result of connecting to a single endpoint
type TestOutboundConnectorResult struct {

	// Endpoint used while trying to connect while using LDAP bind account
	Endpoint *interface{} `mandatory:"false" json:"endpoint"`

	// Status of connection to given endpoint
	Result TestOutboundConnectorResultResultEnum `mandatory:"false" json:"result,omitempty"`

	// Error from the perspective of the mount target.
	MountTargetErrorString *string `mandatory:"false" json:"mountTargetErrorString"`

	// Error string from underlying component that failed.
	EndpointErrorString *string `mandatory:"false" json:"endpointErrorString"`
}

func (m TestOutboundConnectorResult) String() string {
	return common.PointerString(m)
}

// TestOutboundConnectorResultResultEnum Enum with underlying type: string
type TestOutboundConnectorResultResultEnum string

// Set of constants representing the allowable values for TestOutboundConnectorResultResultEnum
const (
	TestOutboundConnectorResultResultSuccess TestOutboundConnectorResultResultEnum = "SUCCESS"
	TestOutboundConnectorResultResultFail    TestOutboundConnectorResultResultEnum = "FAIL"
	TestOutboundConnectorResultResultNotrun  TestOutboundConnectorResultResultEnum = "NOTRUN"
)

var mappingTestOutboundConnectorResultResult = map[string]TestOutboundConnectorResultResultEnum{
	"SUCCESS": TestOutboundConnectorResultResultSuccess,
	"FAIL":    TestOutboundConnectorResultResultFail,
	"NOTRUN":  TestOutboundConnectorResultResultNotrun,
}

// GetTestOutboundConnectorResultResultEnumValues Enumerates the set of values for TestOutboundConnectorResultResultEnum
func GetTestOutboundConnectorResultResultEnumValues() []TestOutboundConnectorResultResultEnum {
	values := make([]TestOutboundConnectorResultResultEnum, 0)
	for _, v := range mappingTestOutboundConnectorResultResult {
		values = append(values, v)
	}
	return values
}
