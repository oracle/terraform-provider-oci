---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_security_features"
sidebar_current: "docs-oci-datasource-data_safe-security_assessment_security_features"
description: |-
  Provides the list of Security Assessment Security Features in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessment_security_features
This data source provides the list of Security Assessment Security Features in Oracle Cloud Infrastructure Data Safe service.

Lists the usage of Database security features for a given compartment or a target level, based on the filters provided.


## Example Usage

```hcl
data "oci_data_safe_security_assessment_security_features" "test_security_assessment_security_features" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.security_assessment_security_feature_access_level
	compartment_id_in_subtree = var.security_assessment_security_feature_compartment_id_in_subtree
	target_id = oci_cloud_guard_target.test_target.id
	targets_with_column_encryption = var.security_assessment_security_feature_targets_with_column_encryption
	targets_with_database_vault = var.security_assessment_security_feature_targets_with_database_vault
	targets_with_external_authentication = var.security_assessment_security_feature_targets_with_external_authentication
	targets_with_fine_grained_audit = var.security_assessment_security_feature_targets_with_fine_grained_audit
	targets_with_global_authentication = var.security_assessment_security_feature_targets_with_global_authentication
	targets_with_network_encryption = var.security_assessment_security_feature_targets_with_network_encryption
	targets_with_password_authentication = var.security_assessment_security_feature_targets_with_password_authentication
	targets_with_privilege_analysis = var.security_assessment_security_feature_targets_with_privilege_analysis
	targets_with_tablespace_encryption = var.security_assessment_security_feature_targets_with_tablespace_encryption
	targets_with_traditional_audit = var.security_assessment_security_feature_targets_with_traditional_audit
	targets_with_unified_audit = var.security_assessment_security_feature_targets_with_unified_audit
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `targets_with_column_encryption` - (Optional) A filter to return only the targets that enable the DB security feature - Column Encryption enabled/disabled.
* `targets_with_database_vault` - (Optional) A filter to return only the targets with the DB security feature - Database Vault enabled/disabled.
* `targets_with_external_authentication` - (Optional) A filter to return only the targets with the DB security feature - External Authentication enabled/disabled.
* `targets_with_fine_grained_audit` - (Optional) A filter to return only the targets with the DB security feature - Fine Grained Audit enabled/disabled.
* `targets_with_global_authentication` - (Optional) A filter to return only the targets with the DB security feature - Global Authentication enabled/disabled.
* `targets_with_network_encryption` - (Optional) A filter to return only the targets with the DB security feature - Network Encryption enabled/disabled.
* `targets_with_password_authentication` - (Optional) A filter to return only the targets with the DB security feature - Password Authentication enabled/disabled.
* `targets_with_privilege_analysis` - (Optional) A filter to return only the targets with the DB security feature - Privilege Analysis enabled/disabled.
* `targets_with_tablespace_encryption` - (Optional) A filter to return only the targets with the DB security feature - Tablespace Encryption enabled/disabled.
* `targets_with_traditional_audit` - (Optional) A filter to return only the targets with the DB security feature - Traditional Audit enabled/disabled.
* `targets_with_unified_audit` - (Optional) A filter to return only the targets with the DB security feature - Unified Audit enabled/disabled.


## Attributes Reference

The following attributes are exported:

* `security_feature_collection` - The list of security_feature_collection.

### SecurityAssessmentSecurityFeature Reference

The following attributes are exported:

* `items` - Array of database security feature summary.
	* `assessment_id` - The OCID of the assessment that generates this security feature usage result.
	* `column_encryption` - The usage of security feature - Column Encryption.
	* `compartment_id` - The OCID of the compartment.
	* `database_vault` - The usage of security feature - Database Vault.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
	* `external_authentication` - The usage of security feature - External Authentication.
	* `fine_grained_audit` - The usage of security feature - Fine Grained Audit.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
	* `global_authentication` - The usage of security feature - Global Authentication.
	* `network_encryption` - The usage of security feature - Network Encryption.
	* `password_authentication` - The usage of security feature - Password Authentication.
	* `privilege_analysis` - The usage of security feature - Privilege Analysis.
	* `tablespace_encryption` - The usage of security feature - Tablespace Encryption.
	* `target_id` - The OCID of the target database.
	* `traditional_audit` - The usage of security feature - Traditional Audit.
	* `unified_audit` - The usage of security feature - Unified Audit.

