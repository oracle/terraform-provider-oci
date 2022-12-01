---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_advanced_feature_configuration"
sidebar_current: "docs-oci-datasource-jms-fleet_advanced_feature_configuration"
description: |-
  Provides details about a specific Fleet Advanced Feature Configuration in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_advanced_feature_configuration
This data source provides details about a specific Fleet Advanced Feature Configuration resource in Oracle Cloud Infrastructure Jms service.

Returns fleet level advanced feature configuration


## Example Usage

```hcl
data "oci_jms_fleet_advanced_feature_configuration" "test_fleet_advanced_feature_configuration" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.


## Attributes Reference

The following attributes are exported:

* `advanced_usage_tracking` - AdvancedUsageTracking configuration
	* `is_enabled` - AdvancedUsageTracking flag to store enabled or disabled status
* `analytic_bucket_name` - Bucket name required to store jfr and related data
* `analytic_namespace` - Namespace for the fleet advanced feature
* `crypto_event_analysis` - CryptoEventAnalysis configuration
	* `is_enabled` - CryptoEventAnalysis flag to store enabled or disabled status
	* `summarized_events_log` - Summarized events log for advanced feature. 
		* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `jfr_recording` - JfrRecording configuration
	* `is_enabled` - JfrRecording flag to store enabled or disabled status
* `lcm` - Enable lifecycle management and set post action configurations
	* `is_enabled` - Lcm flag to store enabled or disabled status
	* `post_installation_actions` - List of available post actions you can execute after the successful Java installation. 
		* `disabled_tls_versions` - The following post JRE installation actions are supported by the field:
			* Disable TLS 1.0 , TLS 1.1 
		* `minimum_key_size_settings` - test
			* `certpath` - Updates the minimum key size for the specified encryption algorithm. The JDK property jdk.certpath.disabledAlgorithms will be updated with the following supported actions:
				* Changing minimum key length for RSA signed jars
				* Changing minimum key length for EC
				* Changing minimum key length for DSA 
				* `key_size` - Key size for the encryption algorithm. Allowed values: 256 for EC, 2048 for DH/DSA/RSA 
				* `name` - The algorithm name.
			* `jar` - Updates the minimum key size for the specified encryption algorithm. The JDK property jdk.jar.disabledAlgorithms will be updated with the following supported actions:
				* Changing minimum key length for RSA signed jars
				* Changing minimum key length for EC
				* Changing minimum key length for DSA 
				* `key_size` - Key size for the encryption algorithm. Allowed values: 256 for EC, 2048 for DH/DSA/RSA 
				* `name` - The algorithm name.
			* `tls` - Updates the minimum key size for the specified encryption algorithm. The JDK property jdk.tls.disabledAlgorithms will be updated with the following supported actions:
				* Changing minimum key length for Diffie-Hellman 
				* `key_size` - Key size for the encryption algorithm. Allowed values: 256 for EC, 2048 for DH/DSA/RSA 
				* `name` - The algorithm name.
		* `should_replace_certificates_operating_system` - Restores JDK root certificates with the certificates that are available in the operating system. The following action is supported by the field:
			* Replace JDK root certificates with a list provided by the operating system 
* `time_last_modified` - The date and time of the last modification to the Fleet Agent Configuration (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 

