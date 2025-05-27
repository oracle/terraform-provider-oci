---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_wls_domain_agreement_records"
sidebar_current: "docs-oci-datasource-wlms-wls_domain_agreement_records"
description: |-
  Provides the list of Wls Domain Agreement Records in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_wls_domain_agreement_records
This data source provides the list of Wls Domain Agreement Records in Oracle Cloud Infrastructure Wlms service.

List the terms of use agreement record for the WebLogic domain.


## Example Usage

```hcl
data "oci_wlms_wls_domain_agreement_records" "test_wls_domain_agreement_records" {
	#Required
	wls_domain_id = oci_wlms_wls_domain.test_wls_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `wls_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.


## Attributes Reference

The following attributes are exported:

* `agreement_record_collection` - The list of agreement_record_collection.

### WlsDomainAgreementRecord Reference

The following attributes are exported:

* `items` - List of agreement records.
	* `agreement_signature` - The agreement signature.
	* `agreement_uuid` - The ID of the accepted agreement.
	* `time_accepted` - The accepted time for the agreement record.

