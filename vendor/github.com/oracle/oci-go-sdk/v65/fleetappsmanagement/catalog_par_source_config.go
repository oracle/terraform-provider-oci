// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CatalogParSourceConfig Catalog par source config.
type CatalogParSourceConfig struct {

	// File path to the directory to use for running Terraform. If not specified, the root directory is used.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// nameSpace
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// bucket name
	BucketName *string `mandatory:"false" json:"bucketName"`

	// object name
	ObjectName *string `mandatory:"false" json:"objectName"`

	// access uri
	AccessUri *string `mandatory:"false" json:"accessUri"`

	// The date and time expires, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`
}

// GetWorkingDirectory returns WorkingDirectory
func (m CatalogParSourceConfig) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m CatalogParSourceConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CatalogParSourceConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CatalogParSourceConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCatalogParSourceConfig CatalogParSourceConfig
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeCatalogParSourceConfig
	}{
		"PAR_CATALOG_SOURCE",
		(MarshalTypeCatalogParSourceConfig)(m),
	}

	return json.Marshal(&s)
}
