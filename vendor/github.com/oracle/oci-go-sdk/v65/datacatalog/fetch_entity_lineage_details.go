// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FetchEntityLineageDetails The information needed to obtain desired lineage.
type FetchEntityLineageDetails struct {

	// Object level at which the lineage is returned.
	Level *int `mandatory:"false" json:"level"`

	// Direction of the lineage returned.
	Direction LineageDirectionEnum `mandatory:"false" json:"direction,omitempty"`

	// Intra-lineages are drill down lineages. This field indicates whether all intra-lineages need to be
	// expanded inline in the lineage returned.
	IsIntraLineage *bool `mandatory:"false" json:"isIntraLineage"`

	// Unique object key for which intra-lineage needs to be fetched. Only drill-down lineage corresponding
	// to the object whose object key is passed is returned.
	IntraLineageObjectKey *string `mandatory:"false" json:"intraLineageObjectKey"`
}

func (m FetchEntityLineageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FetchEntityLineageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLineageDirectionEnum(string(m.Direction)); !ok && m.Direction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Direction: %s. Supported values are: %s.", m.Direction, strings.Join(GetLineageDirectionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
