// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateZoneDetails The body for updating a zone.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateZoneDetails struct {

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The state of DNSSEC on the zone.
	// In order to benefit from utilizing DNSSEC, every parent zone in the DNS tree, up to the TLD or an
	// independent trust anchor, must also have DNSSEC correctly set up. After enabling DNSSEC, a DS record must be
	// added to this zone's parent zone containing data corresponding to the KskDnssecKeyVersion that gets created,
	// and then the KskDnssecKeyVersion must be promoted via the PromoteZoneDnssecKeyVersion operation.
	// New KskDnssecKeyVersions are generated annually, a week before the existing KskDnssecKeyVersion's expiration.
	// KskDnssecKeyVersion rollover requires replacing the parent zone's DS record, corresponding to the current
	// KskDnssecKeyVersion, using the data from its successor KskDnssecKeyVersion. To prevent service disruption
	// from resolver caches including signatures using only the old KSK version, that DS record should not be
	// replaced until the new version has been active for at least the DNSKEY TTL. After the DS replacement has been
	// completed then the PromoteZoneDnssecKeyVersion operation must be called. Metrics are emitted in the oci_dns
	// namespace daily for each KskDnssecKeyVersion indicating how many days are left until expiration. Alarms and
	// notifications should be set up in order to be notified of the KskDnssecKeyVersion expiration so that the
	// necessary parent zone updates can be made and the PromoteZoneDnssecKeyVersion operation can be called.
	// Enabling DNSSEC will result in additional records in DNS responses which will increase their size and can
	// cause higher response latency.
	// For more information, see the DNS docs (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
	DnssecState ZoneDnssecStateEnum `mandatory:"false" json:"dnssecState,omitempty"`

	// External master servers for the zone. `externalMasters` becomes a
	// required parameter when the `zoneType` value is `SECONDARY`.
	ExternalMasters []ExternalMaster `mandatory:"false" json:"externalMasters"`

	// External secondary servers for the zone.
	// This field is currently not supported when `zoneType` is `SECONDARY` or `scope` is `PRIVATE`.
	ExternalDownstreams []ExternalDownstream `mandatory:"false" json:"externalDownstreams"`
}

func (m UpdateZoneDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateZoneDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingZoneDnssecStateEnum(string(m.DnssecState)); !ok && m.DnssecState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnssecState: %s. Supported values are: %s.", m.DnssecState, strings.Join(GetZoneDnssecStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
