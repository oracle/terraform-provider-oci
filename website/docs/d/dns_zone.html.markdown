---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_zone"
sidebar_current: "docs-oci-datasource-dns-zone"
description: |-
  Provides details about a specific Zone in Oracle Cloud Infrastructure DNS service
---

# Data Source: oci_dns_zone
This data source provides details about a specific Zone in Oracle Cloud Infrastructure DNS service.

Gets information about a specific zone by name or OCID.

Notes:
- When accessing a private zone by name, the `view_id` must be provided and `scope` must be `PRIVATE`.
- The `compartment_id` request parameter is deprecated by the service for GetZone and can generally be omitted.

## Example Usage

Using zone OCID:
```hcl
data "oci_dns_zone" "target" {
  # Required
  zone_name_or_id = var.zone_ocid

  # Optional (recommended for private zones)
  scope   = "PRIVATE"
  view_id = oci_dns_view.test_view.id
}
```

Using zone name (private zone):
```hcl
data "oci_dns_zone" "target_by_name" {
  # Required
  zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.example.internal."

  # Required for private zones accessed by name
  scope   = "PRIVATE"
  view_id = oci_dns_view.test_view.id
}
```

## Argument Reference

The following arguments are supported:

- `zone_name_or_id` - (Required) The name or OCID of the target zone.
- `scope` - (Optional) Specifies to operate only on resources that have a matching DNS scope. Allowed values: `GLOBAL`, `PRIVATE`. Required to access private zones by name.
- `view_id` - (Optional) The OCID of the view the zone is associated with. Required when accessing a private zone by name.
- `compartment_id` - (Optional) The OCID of the compartment the zone belongs to. This parameter is deprecated by the service for `GetZone` and should be omitted in most cases.

## Attributes Reference

The following attributes are exported:

- `id` - The OCID of the zone.
- `compartment_id` - The OCID of the compartment containing the zone.
- `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

   **Example:** `{"Operations": {"CostCenter": "42"}}`
- `dnssec_config` - DNSSEC configuration data.
  - `ksk_dnssec_key_versions` - A read-only array of key signing key (KSK) versions.
    - `algorithm` - The signing algorithm used for the key.
    - `ds_data` - An array of data for DS records corresponding with this key version. An entry will exist for each supported DS digest algorithm.
      - `digest_type` - The type of the digest associated with the rdata.
      - `rdata` - Presentation-format DS record data that must be added to the parent zone.
    - `key_tag` - The key tag associated with the `DnssecKeyVersion`.
    - `length_in_bytes` - The length of the corresponding private key in bytes.
    - `predecessor_dnssec_key_version_uuid` - UUID of the key version this will replace or has replaced.
    - `successor_dnssec_key_version_uuid` - UUID of the key version that will replace or has replaced this key version.
    - `time_activated` - RFC 3339 timestamp when the key version went or will go active.
    - `time_created` - RFC 3339 timestamp when the key version was created.
    - `time_expired` - RFC 3339 timestamp for end of recommended lifetime.
    - `time_inactivated` - RFC 3339 timestamp when the key version went or will go inactive.
    - `time_promoted` - RFC 3339 timestamp when the key version was promoted.
    - `time_published` - RFC 3339 timestamp when the zone contents include a DNSKEY for the key material.
    - `time_unpublished` - RFC 3339 timestamp when the DNSKEY is removed from the zone contents.
    - `uuid` - The UUID of the `DnssecKeyVersion`.
  - `zsk_dnssec_key_versions` - A read-only array of zone signing key (ZSK) versions.
    - `algorithm` - The signing algorithm used for the key.
    - `key_tag` - The key tag associated with the `DnssecKeyVersion`.
    - `length_in_bytes` - The length of the corresponding private key in bytes.
    - `predecessor_dnssec_key_version_uuid` - UUID of the key version this will replace or has replaced.
    - `successor_dnssec_key_version_uuid` - UUID of the key version that will replace or has replaced this key version.
    - `time_activated` - RFC 3339 timestamp when the key version went or will go active.
    - `time_created` - RFC 3339 timestamp when the key version was created.
    - `time_expired` - RFC 3339 timestamp for end of recommended lifetime.
    - `time_inactivated` - RFC 3339 timestamp when the key version went or will go inactive.
    - `time_promoted` - RFC 3339 timestamp when the key version was promoted.
    - `time_published` - RFC 3339 timestamp when the zone contents include a DNSKEY for the key material.
    - `time_unpublished` - RFC 3339 timestamp when the DNSKEY is removed from the zone contents.
    - `uuid` - The UUID of the `DnssecKeyVersion`.
- `dnssec_state` - The state of DNSSEC on the zone.
- `external_downstreams` - External secondary servers for the zone.
  - `address` - The server's IP address (IPv4 or IPv6).
  - `port` - The server's port. Must be 53 if provided; otherwise omit.
  - `tsig_key_id` - The OCID of the TSIG key used to secure zone transfers.
- `external_masters` - External master servers for the zone (required when `zone_type` is `SECONDARY`).
  - `address` - The server's IP address (IPv4 or IPv6).
  - `port` - The server's port. Must be 53 if provided; otherwise omit.
  - `tsig_key_id` - The OCID of the TSIG key.
- `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair.
- `is_protected` - Whether parts of the resource are protected from explicit management.
- `name` - The name of the zone.
- `nameservers` - Nameservers for the zone.
  - `hostname` - Hostname of the nameserver.
- `resolution_mode` - The resolution mode for the zone.
- `scope` - The scope of the zone.
- `self` - The canonical absolute URL of the resource.
- `serial` - The current serial of the zone (from the SOA record).
- `state` - The current lifecycle state of the zone.
- `time_created` - The RFC 3339 timestamp when the zone was created.
- `version` - The version of the zone from which the SOA serial is derived.
- `view_id` - The OCID of the private view containing the zone (null for global zones).
- `zone_transfer_servers` - The OCI nameservers that transfer zone data with external nameservers.
  - `address` - The server's IP address (IPv4 or IPv6).
  - `is_transfer_destination` - Whether the server is a zone data transfer destination.
  - `is_transfer_source` - Whether the server is a zone data transfer source.
  - `port` - The server's port.
- `zone_type` - The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is supported only for `GLOBAL` zones.
