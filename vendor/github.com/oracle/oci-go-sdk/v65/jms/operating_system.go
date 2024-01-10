// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OperatingSystem Operating System of the platform on which the Java Runtime was reported.
type OperatingSystem struct {

	// The operating system type, such as Windows or Linux
	Family OsFamilyEnum `mandatory:"true" json:"family"`

	// The name of the operating system as provided by the Java system property os.name.
	Name *string `mandatory:"true" json:"name"`

	// The version of the operating system as provided by the Java system property os.version.
	Version *string `mandatory:"true" json:"version"`

	// The architecture of the operating system as provided by the Java system property os.arch.
	Architecture *string `mandatory:"true" json:"architecture"`

	// Number of instances running the operating system.
	ManagedInstanceCount *int `mandatory:"false" json:"managedInstanceCount"`
}

func (m OperatingSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OperatingSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOsFamilyEnum(string(m.Family)); !ok && m.Family != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Family: %s. Supported values are: %s.", m.Family, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
