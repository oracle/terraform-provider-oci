---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_zone"
sidebar_current: "docs-oci-resource-dns-zone"
description: |-
  Provides the Zone resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_zone
This resource provides the Zone resource in Oracle Cloud Infrastructure DNS service.

Creates a new zone in the specified compartment.

Private zones must have a zone type of `PRIMARY`. Creating a private zone at or under `oraclevcn.com`
within the default protected view of a VCN-dedicated resolver is not permitted.


## Example Usage

```hcl
resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = var.compartment_id
	name = var.zone_name
	zone_type = var.zone_zone_type

	#Optional
	defined_tags = var.zone_defined_tags
	dnssec_state = var.zone_dnssec_state
	external_downstreams {
		#Required
		address = var.zone_external_downstreams_address

		#Optional
		port = var.zone_external_downstreams_port
		tsig_key_id = oci_dns_tsig_key.test_tsig_key.id
	}
	external_masters {
		#Required
		address = var.zone_external_masters_address

		#Optional
		port = var.zone_external_masters_port
		tsig_key_id = oci_dns_tsig_key.test_tsig_key.id
	}
	freeform_tags = var.zone_freeform_tags
	scope = var.zone_scope
	view_id = oci_dns_view.test_view.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment containing the zone.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `dnssec_state` - (Optional) (Updatable) The state of DNSSEC on the zone.

	For DNSSEC to function, every parent zone in the DNS tree up to the top-level domain (or an independent trust anchor) must also have DNSSEC correctly set up. After enabling DNSSEC, you must add a DS record to the zone's parent zone containing the `KskDnssecKeyVersion` data. You can find the DS data in the `dsData` attribute of the `KskDnssecKeyVersion`. Then, use the `PromoteZoneDnssecKeyVersion` operation to promote the `KskDnssecKeyVersion`.

	New `KskDnssecKeyVersion`s are generated annually, a week before the existing `KskDnssecKeyVersion`'s expiration. To rollover a `KskDnssecKeyVersion`, you must replace the parent zone's DS record containing the old `KskDnssecKeyVersion` data with the data from the new `KskDnssecKeyVersion`.

	To remove the old DS record without causing service disruption, wait until the old DS record's TTL has expired, and the new DS record has propagated. After the DS replacement has been completed, then the `PromoteZoneDnssecKeyVersion` operation must be called.

	Metrics are emitted in the `oci_dns` namespace daily for each `KskDnssecKeyVersion` indicating how many days are left until expiration. We recommend that you set up alarms and notifications for KskDnssecKeyVersion expiration so that the necessary parent zone updates can be made and the `PromoteZoneDnssecKeyVersion` operation can be called.

	Enabling DNSSEC results in additional records in DNS responses which increases their size and can cause higher response latency.

	For more information, see [DNSSEC](https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnssec.htm). 
* `external_downstreams` - (Optional) (Updatable) External secondary servers for the zone. This field is currently not supported when `zoneType` is `SECONDARY` or `scope` is `PRIVATE`. 
	* `address` - (Required) (Updatable) The server's IP address (IPv4 or IPv6).
	* `port` - (Optional) (Updatable) The server's port. Port value must be a value of 53, otherwise omit the port value. 
	* `tsig_key_id` - (Optional) (Updatable) The OCID of the TSIG key. A TSIG key is used to secure DNS messages (in this case, zone transfers) between two systems that both have the (shared) secret. 
* `external_masters` - (Optional) (Updatable) External master servers for the zone. `externalMasters` becomes a required parameter when the `zoneType` value is `SECONDARY`. 
	* `address` - (Required) (Updatable) The server's IP address (IPv4 or IPv6).
	* `port` - (Optional) (Updatable) The server's port. Port value must be a value of 53, otherwise omit the port value. 
	* `tsig_key_id` - (Optional) (Updatable) The OCID of the TSIG key.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `name` - (Required) The name of the zone.
* `scope` - (Optional) Specifies to operate only on resources that have a matching DNS scope. 
This value will be null for zones in the global DNS and `PRIVATE` when creating a private zone.
* `view_id` - (Optional) The OCID of the private view containing the zone. This value will be null for zones in the global DNS, which are publicly resolvable and not part of a private view. 
* `zone_type` - (Required) The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL zones. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the zone.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `dnssec_config` - DNSSEC configuration data.

	A zone may have a maximum of 10 `DnssecKeyVersions`, regardless of signing key type. 
	* `ksk_dnssec_key_versions` - A read-only array of key signing key (KSK) versions. 
		* `algorithm` - The signing algorithm used for the key. 
		* `ds_data` - An array of data for DS records corresponding with this key version. An entry will exist for each supported DS digest algorithm. 
			* `digest_type` - The type of the digest associated with the rdata. 
			* `rdata` - Presentation-format DS record data that must be added to the parent zone. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm) 
		* `key_tag` - The key tag associated with the `DnssecKeyVersion`. This key tag will be present in the RRSIG and DS records associated with the key material for this `DnssecKeyVersion`. For more information about key tags, see [RFC 4034](https://tools.ietf.org/html/rfc4034). 
		* `length_in_bytes` - The length of the corresponding private key in bytes, expressed as an integer. 
		* `predecessor_dnssec_key_version_uuid` - When populated, this is the UUID of the `DnssecKeyVersion` that this `DnssecKeyVersion` will replace or has replaced. 
		* `successor_dnssec_key_version_uuid` - When populated, this is the UUID of the `DnssecKeyVersion` that will replace, or has replaced, this `DnssecKeyVersion`. 
		* `time_activated` - The date and time the key version went, or will go, active, expressed in RFC 3339 timestamp format. This is when the key material will be used to generate RRSIGs.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_created` - The date and time the key version was created, expressed in RFC 3339 timestamp format.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_expired` - The date and time at which the recommended key version publication/activation lifetime ends, expressed in RFC 3339 timestamp format. This is when the corresponding DNSKEY should no longer exist in zone contents and no longer be used to generate RRSIGs. For a key sigining key (KSK), if `PromoteZoneDnssecKeyVersion` has not been called on this `DnssecKeyVersion`'s successor then it will remain active for arbitrarily long past its recommended lifetime. This prevents service disruption at the potential increased risk of key compromise.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_inactivated` - The date and time the key version went, or will go, inactive, expressed in RFC 3339 timestamp format. This is when the key material will no longer be used to generate RRSIGs. For a key signing key (KSK) `DnssecKeyVersion`, this is populated after `PromoteZoneDnssecKeyVersion` has been called on its successor `DnssecKeyVersion`.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_promoted` - The date and time the key version was promoted expressed in RFC 3339 timestamp format.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_published` - The date and time the key version was, or will be, published, expressed in RFC 3339 timestamp format. This is when the zone contents will include a DNSKEY record corresponding to the key material.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_unpublished` - The date and time the key version was, or will be, unpublished, expressed in RFC 3339 timestamp format. This is when the corresponding DNSKEY will be removed from zone contents. For a key signing key (KSK) `DnssecKeyVersion`, this is populated after `PromoteZoneDnssecKeyVersion` has been called on its successor `DnssecKeyVersion`.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `uuid` - The UUID of the `DnssecKeyVersion`. 
	* `zsk_dnssec_key_versions` - A read-only array of zone signing key (ZSK) versions. 
		* `algorithm` - The signing algorithm used for the key. 
		* `key_tag` - The key tag associated with the `DnssecKeyVersion`. This key tag will be present in the RRSIG and DS records associated with the key material for this `DnssecKeyVersion`. For more information about key tags, see [RFC 4034](https://tools.ietf.org/html/rfc4034). 
		* `length_in_bytes` - The length of the corresponding private key in bytes, expressed as an integer. 
		* `predecessor_dnssec_key_version_uuid` - When populated, this is the UUID of the `DnssecKeyVersion` that this `DnssecKeyVersion` will replace or has replaced. 
		* `successor_dnssec_key_version_uuid` - When populated, this is the UUID of the `DnssecKeyVersion` that will replace, or has replaced, this `DnssecKeyVersion`. 
		* `time_activated` - The date and time the key version went, or will go, active, expressed in RFC 3339 timestamp format. This is when the key material will be used to generate RRSIGs.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_created` - The date and time the key version was created, expressed in RFC 3339 timestamp format.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_expired` - The date and time at which the recommended key version publication/activation lifetime ends, expressed in RFC 3339 timestamp format. This is when the corresponding DNSKEY should no longer exist in zone contents and no longer be used to generate RRSIGs. For a key sigining key (KSK), if `PromoteZoneDnssecKeyVersion` has not been called on this `DnssecKeyVersion`'s successor then it will remain active for arbitrarily long past its recommended lifetime. This prevents service disruption at the potential increased risk of key compromise.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_inactivated` - The date and time the key version went, or will go, inactive, expressed in RFC 3339 timestamp format. This is when the key material will no longer be used to generate RRSIGs. For a key signing key (KSK) `DnssecKeyVersion`, this is populated after `PromoteZoneDnssecKeyVersion` has been called on its successor `DnssecKeyVersion`.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_promoted` - The date and time the key version was promoted expressed in RFC 3339 timestamp format.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_published` - The date and time the key version was, or will be, published, expressed in RFC 3339 timestamp format. This is when the zone contents will include a DNSKEY record corresponding to the key material.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `time_unpublished` - The date and time the key version was, or will be, unpublished, expressed in RFC 3339 timestamp format. This is when the corresponding DNSKEY will be removed from zone contents. For a key signing key (KSK) `DnssecKeyVersion`, this is populated after `PromoteZoneDnssecKeyVersion` has been called on its successor `DnssecKeyVersion`.

			**Example:** `2016-07-22T17:23:59:00Z` 
		* `uuid` - The UUID of the `DnssecKeyVersion`. 
* `dnssec_state` - The state of DNSSEC on the zone.

	For DNSSEC to function, every parent zone in the DNS tree up to the top-level domain (or an independent trust anchor) must also have DNSSEC correctly set up. After enabling DNSSEC, you must add a DS record to the zone's parent zone containing the `KskDnssecKeyVersion` data. You can find the DS data in the `dsData` attribute of the `KskDnssecKeyVersion`. Then, use the `PromoteZoneDnssecKeyVersion` operation to promote the `KskDnssecKeyVersion`.

	New `KskDnssecKeyVersion`s are generated annually, a week before the existing `KskDnssecKeyVersion`'s expiration. To rollover a `KskDnssecKeyVersion`, you must replace the parent zone's DS record containing the old `KskDnssecKeyVersion` data with the data from the new `KskDnssecKeyVersion`.

	To remove the old DS record without causing service disruption, wait until the old DS record's TTL has expired, and the new DS record has propagated. After the DS replacement has been completed, then the `PromoteZoneDnssecKeyVersion` operation must be called.

	Metrics are emitted in the `oci_dns` namespace daily for each `KskDnssecKeyVersion` indicating how many days are left until expiration. We recommend that you set up alarms and notifications for KskDnssecKeyVersion expiration so that the necessary parent zone updates can be made and the `PromoteZoneDnssecKeyVersion` operation can be called.

	Enabling DNSSEC results in additional records in DNS responses which increases their size and can cause higher response latency.

	For more information, see [DNSSEC](https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnssec.htm). 
* `external_downstreams` - External secondary servers for the zone. This field is currently not supported when `zoneType` is `SECONDARY` or `scope` is `PRIVATE`. 
	* `address` - The server's IP address (IPv4 or IPv6).
	* `port` - The server's port. Port value must be a value of 53, otherwise omit the port value. 
	* `tsig_key_id` - The OCID of the TSIG key. A TSIG key is used to secure DNS messages (in this case, zone transfers) between two systems that both have the (shared) secret. 
* `external_masters` - External master servers for the zone. `externalMasters` becomes a required parameter when the `zoneType` value is `SECONDARY`. 
	* `address` - The server's IP address (IPv4 or IPv6).
	* `port` - The server's port. Port value must be a value of 53, otherwise omit the port value. 
	* `tsig_key_id` - The OCID of the TSIG key.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `id` - The OCID of the zone.
* `is_protected` - A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed. 
* `name` - The name of the zone.
* `nameservers` - The authoritative nameservers for the zone.
	* `hostname` - The hostname of the nameserver.
* `scope` - The scope of the zone.
* `self` - The canonical absolute URL of the resource.
* `serial` - The current serial of the zone. As seen in the zone's SOA record. 
* `state` - The current state of the zone resource.
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `version` - Version is the never-repeating, totally-orderable, version of the zone, from which the serial field of the zone's SOA record is derived. 
* `view_id` - The OCID of the private view containing the zone. This value will be null for zones in the global DNS, which are publicly resolvable and not part of a private view. 
* `zone_transfer_servers` - The Oracle Cloud Infrastructure nameservers that transfer the zone data with external nameservers. 
	* `address` - The server's IP address (IPv4 or IPv6).
	* `is_transfer_destination` - A Boolean flag indicating whether or not the server is a zone data transfer destination. 
	* `is_transfer_source` - A Boolean flag indicating whether or not the server is a zone data transfer source. 
	* `port` - The server's port. 
* `zone_type` - The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL zones. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Zone
	* `update` - (Defaults to 20 minutes), when updating the Zone
	* `delete` - (Defaults to 20 minutes), when destroying the Zone


## Import

Zones can be imported using the `id`, e.g.

```
$ terraform import oci_dns_zone.test_zone "id"
```
