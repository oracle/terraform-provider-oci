// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateSourceScanHost Private source Scan Hostname model.
type PrivateSourceScanHost struct {

	// Private Source Scan hostname. Ex: db01-scan.corp.example.com, prd-db01-scan.mycompany.com.
	ScanHostname *string `mandatory:"true" json:"scanHostname"`

	// Private Source Scan host port. This is the source port where SCAN protocol will get connected (e.g. 1521).
	ScanPort *int `mandatory:"true" json:"scanPort"`

	// Description of private source scan host zone.
	Description *string `mandatory:"false" json:"description"`
}

func (m PrivateSourceScanHost) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateSourceScanHost) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
