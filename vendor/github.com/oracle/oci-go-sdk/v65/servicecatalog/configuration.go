// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Use the Service Catalog API to manage solutions in Oracle Cloud Infrastructure Service Catalog.
// For more information, see Overview of Service Catalog (https://docs.oracle.com/iaas/Content/service-catalog/overview_of_service_catalog.htm).
//

package servicecatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Configuration returns the service catalog mode
type Configuration struct {

	// mode of tenancy
	IsServiceCatalogMode ServiceCatalogModeEnumEnum `mandatory:"true" json:"isServiceCatalogMode"`
}

func (m Configuration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Configuration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingServiceCatalogModeEnumEnum(string(m.IsServiceCatalogMode)); !ok && m.IsServiceCatalogMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsServiceCatalogMode: %s. Supported values are: %s.", m.IsServiceCatalogMode, strings.Join(GetServiceCatalogModeEnumEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
