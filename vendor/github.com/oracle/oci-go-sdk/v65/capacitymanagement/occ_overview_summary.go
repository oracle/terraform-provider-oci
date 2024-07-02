// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccOverviewSummary The overview summary is used to represent an array item that shall be used to represent the overview of the catalog resources along with their corresponding capacity requests.
type OccOverviewSummary struct {

	// The OCID of the compartment from which the api call is made. This will be used for authorizing the request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the month along with year for which this summary corresponds to.
	PeriodValue *string `mandatory:"true" json:"periodValue"`

	// The name of the resource for which we have aggregated the value.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The quantity of the resource which is available at the end of the period of aggregationDetails model in consideration.
	TotalAvailable *int64 `mandatory:"true" json:"totalAvailable"`

	// The quantity of the resource which is demanded by customers via capacity requests against the resource name at the end of the time period in consideration for overview.
	TotalDemanded *int64 `mandatory:"true" json:"totalDemanded"`

	// The quantity of the resource which is supplied by Oracle to the customer against the resource name at the end of the time period in consideration.
	TotalSupplied *int64 `mandatory:"true" json:"totalSupplied"`

	// The quantity of the resource which is rejected by Oracle.
	TotalRejected *int64 `mandatory:"true" json:"totalRejected"`

	// The quantity of the resource which is cancelled by the customer. Once the capacity request was submitted, the customer can still cancel it. This field sums up those values.
	TotalCancelled *int64 `mandatory:"true" json:"totalCancelled"`

	// The quantity of the resource which Oracle was unable to supply. For a given capacity request, Oracle sometimes cannot supply the entire value demanded by the customer. In such cases a partial value is provided, thereby leaving behind a portion of unfulfilled values. This field sums that up.
	TotalUnfulfilled *int64 `mandatory:"true" json:"totalUnfulfilled"`

	// A raw json blob containing breakdown of totalAvailable, totalDemanded, totalSupplied, totalRejected, totalCancelled and totalUnfulfilled by workload types
	WorkloadTypeBreakdownBlob *string `mandatory:"true" json:"workloadTypeBreakdownBlob"`

	// The unit e.g SERVER in which the above values like totalAvailable, totalSupplied etc is measured.
	Unit *string `mandatory:"true" json:"unit"`

	// A raw json blob containing all the capacity requests corresponding to the resource name
	CapacityRequestsBlob *string `mandatory:"false" json:"capacityRequestsBlob"`
}

func (m OccOverviewSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccOverviewSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
