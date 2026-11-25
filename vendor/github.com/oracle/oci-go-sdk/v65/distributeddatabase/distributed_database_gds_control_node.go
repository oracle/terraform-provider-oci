// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DistributedDatabaseGdsControlNode Details of Global Database Services Control(GDS CTL) instances for the Globally distributed database.
type DistributedDatabaseGdsControlNode struct {

	// Name of the Global Database Services Control instance
	Name *string `mandatory:"true" json:"name"`

	// The time the Global Database Services Control instance was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Global Database Services Control instance was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Global Database Services Control instance.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The name of the availability domain where the Global Database Services Control instance is located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// IP address of the Global Database Services Control instance.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	GdsControlNodeImageDetails *DistributedDbGdsControlNodeImage `mandatory:"false" json:"gdsControlNodeImageDetails"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`
}

func (m DistributedDatabaseGdsControlNode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseGdsControlNode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
