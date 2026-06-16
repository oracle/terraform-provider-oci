// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataFileSummary Properties related to the file.
type DataFileSummary struct {

	// The name to use as identifier for the data file.
	Name *string `mandatory:"true" json:"name"`

	// The type of the data file, indicating its intended use
	// Example: `source_map`
	ApmType *string `mandatory:"true" json:"apmType"`

	// Size of the object in bytes.
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// Base64-encoded MD5 hash of the object data.
	Md5 *string `mandatory:"false" json:"md5"`

	// The last time the object was modified, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.29).
	// Expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2020-02-19T22:47:12.613Z`
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// Metadata properties related to the data file.
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m DataFileSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataFileSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
