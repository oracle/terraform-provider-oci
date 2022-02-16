// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Zone A DNS zone.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type Zone struct {

	// The name of the zone.
	Name *string `mandatory:"true" json:"name"`

	// The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL zones.
	ZoneType ZoneZoneTypeEnum `mandatory:"true" json:"zoneType"`

	// The OCID of the compartment containing the zone.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The scope of the zone.
	Scope ScopeEnum `mandatory:"true" json:"scope"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The state of DNSSEC on the zone.
	// In order to benefit from utilizing DNSSEC, every parent zone in the DNS tree, up to the TLD or an
	// independent trust anchor, must also have DNSSEC correctly set up. After enabling DNSSEC, a DS record must be
	// added to this zone's parent zone containing data corresponding to the KskDnssecKeyVersion that gets created.
	// New KskDnssecKeyVersions are generated annually, a week before the existing KskDnssecKeyVersion's expiration.
	// KskDnssecKeyVersion rollover requires replacing the parent zone's DS record, corresponding to the current
	// KskDnssecKeyVersion, using the data from its successor KskDnssecKeyVersion. To prevent service disruption
	// from resolver caches including signatures using only the old KSK version, that DS record should not be
	// replaced until the new version has been active for at least the DNSKEY TTL. After the DS replacement has been
	// completed then the PromoteZoneDnssecKeyVersion operation must be called. Metrics are emitted in the oci_dns
	// namespace daily for each KskDnssecKeyVersion indicating how many days are left until expiration. Alarms and
	// notifications should be set up in order to be notified of the KskDnssecKeyVersion expiration so that the
	// necessary parent zone updates can be made and the PromoteZoneDnssecKeyVersion operation can be called.
	// Zones with DNSSEC enabled are subject to a maximum allowed TTL on records of 1 day (86400 seconds). Enabling
	// DNSSEC will result in additional records in DNS responses which will increase their size and can cause higher
	// response latency. Re-enabling DNSSEC on a zone shortly after it being disabled will restore the previous
	// DnssecKeyVersions.
	// TODO: Add link to DNSSEC docs covering: how to set up alarms/notifications, warnings about enabling/disabling,
	// warnings about timing and impacts, how to handle automatic rollover, how to handle manual rollover, and how
	// to handle emergency rollover.
	DnssecState ZoneDnssecStateEnum `mandatory:"true" json:"dnssecState"`

	// External master servers for the zone. `externalMasters` becomes a
	// required parameter when the `zoneType` value is `SECONDARY`.
	ExternalMasters []ExternalMaster `mandatory:"true" json:"externalMasters"`

	// External secondary servers for the zone.
	// This field is currently not supported when `zoneType` is `SECONDARY` or `scope` is `PRIVATE`.
	ExternalDownstreams []ExternalDownstream `mandatory:"true" json:"externalDownstreams"`

	// The canonical absolute URL of the resource.
	Self *string `mandatory:"true" json:"self"`

	// The OCID of the zone.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format
	// with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Version is the never-repeating, totally-orderable, version of the
	// zone, from which the serial field of the zone's SOA record is
	// derived.
	Version *string `mandatory:"true" json:"version"`

	// The current serial of the zone. As seen in the zone's SOA record.
	Serial *int64 `mandatory:"true" json:"serial"`

	// The current state of the zone resource.
	LifecycleState ZoneLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.
	IsProtected *bool `mandatory:"true" json:"isProtected"`

	// The authoritative nameservers for the zone.
	Nameservers []Nameserver `mandatory:"true" json:"nameservers"`

	// The OCID of the private view containing the zone. This value will
	// be null for zones in the global DNS, which are publicly resolvable and
	// not part of a private view.
	ViewId *string `mandatory:"false" json:"viewId"`

	DnssecConfig *DnssecConfig `mandatory:"false" json:"dnssecConfig"`

	// The OCI nameservers that transfer the zone data with external nameservers.
	ZoneTransferServers []ZoneTransferServer `mandatory:"false" json:"zoneTransferServers"`
}

func (m Zone) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Zone) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingZoneZoneTypeEnum[string(m.ZoneType)]; !ok && m.ZoneType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ZoneType: %s. Supported values are: %s.", m.ZoneType, strings.Join(GetZoneZoneTypeEnumStringValues(), ",")))
	}
	if _, ok := mappingScopeEnum[string(m.Scope)]; !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetScopeEnumStringValues(), ",")))
	}
	if _, ok := mappingZoneDnssecStateEnum[string(m.DnssecState)]; !ok && m.DnssecState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnssecState: %s. Supported values are: %s.", m.DnssecState, strings.Join(GetZoneDnssecStateEnumStringValues(), ",")))
	}
	if _, ok := mappingZoneLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetZoneLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ZoneZoneTypeEnum Enum with underlying type: string
type ZoneZoneTypeEnum string

// Set of constants representing the allowable values for ZoneZoneTypeEnum
const (
	ZoneZoneTypePrimary   ZoneZoneTypeEnum = "PRIMARY"
	ZoneZoneTypeSecondary ZoneZoneTypeEnum = "SECONDARY"
)

var mappingZoneZoneTypeEnum = map[string]ZoneZoneTypeEnum{
	"PRIMARY":   ZoneZoneTypePrimary,
	"SECONDARY": ZoneZoneTypeSecondary,
}

// GetZoneZoneTypeEnumValues Enumerates the set of values for ZoneZoneTypeEnum
func GetZoneZoneTypeEnumValues() []ZoneZoneTypeEnum {
	values := make([]ZoneZoneTypeEnum, 0)
	for _, v := range mappingZoneZoneTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetZoneZoneTypeEnumStringValues Enumerates the set of values in String for ZoneZoneTypeEnum
func GetZoneZoneTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"SECONDARY",
	}
}

// ZoneLifecycleStateEnum Enum with underlying type: string
type ZoneLifecycleStateEnum string

// Set of constants representing the allowable values for ZoneLifecycleStateEnum
const (
	ZoneLifecycleStateActive   ZoneLifecycleStateEnum = "ACTIVE"
	ZoneLifecycleStateCreating ZoneLifecycleStateEnum = "CREATING"
	ZoneLifecycleStateDeleted  ZoneLifecycleStateEnum = "DELETED"
	ZoneLifecycleStateDeleting ZoneLifecycleStateEnum = "DELETING"
	ZoneLifecycleStateFailed   ZoneLifecycleStateEnum = "FAILED"
	ZoneLifecycleStateUpdating ZoneLifecycleStateEnum = "UPDATING"
)

var mappingZoneLifecycleStateEnum = map[string]ZoneLifecycleStateEnum{
	"ACTIVE":   ZoneLifecycleStateActive,
	"CREATING": ZoneLifecycleStateCreating,
	"DELETED":  ZoneLifecycleStateDeleted,
	"DELETING": ZoneLifecycleStateDeleting,
	"FAILED":   ZoneLifecycleStateFailed,
	"UPDATING": ZoneLifecycleStateUpdating,
}

// GetZoneLifecycleStateEnumValues Enumerates the set of values for ZoneLifecycleStateEnum
func GetZoneLifecycleStateEnumValues() []ZoneLifecycleStateEnum {
	values := make([]ZoneLifecycleStateEnum, 0)
	for _, v := range mappingZoneLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetZoneLifecycleStateEnumStringValues Enumerates the set of values in String for ZoneLifecycleStateEnum
func GetZoneLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}
