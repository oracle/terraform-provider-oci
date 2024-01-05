// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModuleStream A module stream provided by a software source.
type ModuleStream struct {

	// The name of the module that contains the stream.
	ModuleName *string `mandatory:"true" json:"moduleName"`

	// The name of the stream.
	Name *string `mandatory:"true" json:"name"`

	// Indicates if this stream is the default for its module.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// The OCID of the software source that provides this module stream.
	SoftwareSourceId *string `mandatory:"false" json:"softwareSourceId"`

	// The architecture for which the packages in this module stream were built.
	ArchType ArchTypeEnum `mandatory:"false" json:"archType,omitempty"`

	// A description of the contents of the module stream.
	Description *string `mandatory:"false" json:"description"`

	// A list of profiles that are part of the stream.  Each element in
	// the list is the name of a profile.  The name is suitable to use as
	// an argument to other OS Management Hub APIs that interact directly with
	// module stream profiles.  However, it is not URL encoded.
	Profiles []string `mandatory:"false" json:"profiles"`

	// A list of packages that are contained by the stream.  Each element
	// in the list is the name of a package.  The name is suitable to use
	// as an argument to other OS Management Hub APIs that interact directly
	// with packages.
	Packages []string `mandatory:"false" json:"packages"`

	// Indicates whether this module stream is the latest.
	IsLatest *bool `mandatory:"false" json:"isLatest"`
}

func (m ModuleStream) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModuleStream) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArchTypeEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
