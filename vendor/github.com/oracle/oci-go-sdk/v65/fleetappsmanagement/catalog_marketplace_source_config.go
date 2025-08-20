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

// CatalogMarketplaceSourceConfig Catalog market source config.
type CatalogMarketplaceSourceConfig struct {

	// File path to the directory to use for running Terraform. If not specified, the root directory is used.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// This listing Id parameter of Payload.
	ListingId *string `mandatory:"false" json:"listingId"`

	// This version parameter of Payload.
	Version *string `mandatory:"false" json:"version"`
}

// GetWorkingDirectory returns WorkingDirectory
func (m CatalogMarketplaceSourceConfig) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m CatalogMarketplaceSourceConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CatalogMarketplaceSourceConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CatalogMarketplaceSourceConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCatalogMarketplaceSourceConfig CatalogMarketplaceSourceConfig
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeCatalogMarketplaceSourceConfig
	}{
		"MARKETPLACE_CATALOG_SOURCE",
		(MarshalTypeCatalogMarketplaceSourceConfig)(m),
	}

	return json.Marshal(&s)
}
