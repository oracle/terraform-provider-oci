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

// EntityLineage Lineage for a data entity.
type EntityLineage struct {

	// Object level at which the lineage is returned.
	Level *int `mandatory:"true" json:"level"`

	// Direction of the lineage returned.
	Direction LineageDirectionEnum `mandatory:"true" json:"direction"`

	// Set of objects that are involved in the lineage.
	Objects []LineageObject `mandatory:"false" json:"objects"`

	// Set of relationships between the objects in the 'objects' set.
	Relationships []LineageRelationship `mandatory:"false" json:"relationships"`

	// A map of maps that contains additional information in explanation of the lineage returned. The map keys are
	// categories of information and the values are maps of annotation names to their corresponding values.
	// Every annotation is contained inside a category.
	// Example: `{"annotations": { "category": { "key": "value"}}}`
	Annotations map[string]map[string]string `mandatory:"false" json:"annotations"`
}

func (m EntityLineage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EntityLineage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLineageDirectionEnum(string(m.Direction)); !ok && m.Direction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Direction: %s. Supported values are: %s.", m.Direction, strings.Join(GetLineageDirectionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
