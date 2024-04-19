---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_vcn_dns_resolver_association"
sidebar_current: "docs-oci-datasource-core-vcn_dns_resolver_association"
description: |-
  Provides details about a specific Vcn Dns Resolver Association in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_vcn_dns_resolver_association
This data source provides details about a specific Vcn Dns Resolver Association resource in Oracle Cloud Infrastructure Core service.

Get the associated DNS resolver information with a vcn

## Example Usage

```hcl
data "oci_core_vcn_dns_resolver_association" "test_vcn_dns_resolver_association" {
	#Required
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `vcn_id` - (Required) Specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


## Attributes Reference

The following attributes are exported:

* `dns_resolver_id` - The OCID of the DNS resolver in the association. We won't have the DNS resolver id as soon as vcn 
is created, we will create it asynchronously. It would be null until it is actually created.
* `state` - The current state of the association. Resolver will be created when vcn is created. But the creation happens 
asynchronously and may take longer because it is a background event that needs to run. The state will be PROVISIONING 
until the resolver is actually created. After the resolver is actually created, the state will be set to AVAILABLE. Users 
need to do a terraform refresh to poll and update the state file after sometime to get the dns_resolver_id and state AVAILABLE.
* `vcn_id` - The OCID of the VCN in the association.

