---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_advanced_feature_configuration"
sidebar_current: "docs-oci-resource-jms-fleet_advanced_feature_configuration"
description: |-
  Provides the Fleet Advanced Feature Configuration resource in Oracle Cloud Infrastructure Jms service
---

# oci_jms_fleet_advanced_feature_configuration
This resource provides the Fleet Advanced Feature Configuration resource in Oracle Cloud Infrastructure Jms service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/jms/latest/FleetAdvancedFeatureConfiguration

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/jms

Update advanced feature configurations for the Fleet.
Ensure that the namespace and bucket storage are created prior to turning on the JfrRecording or CryptoEventAnalysis feature.


## Example Usage

```hcl
resource "oci_jms_fleet_advanced_feature_configuration" "test_fleet_advanced_feature_configuration" {

	#Required
	fleet_id = oci_jms_fleet.test_fleet.id

	#Optional
	advanced_usage_tracking {

		#Optional
		is_enabled = true
	}
	analytic_bucket_name = oci_objectstorage_bucket.test_bucket.name
	analytic_namespace = "example-bucket-namespace"
	crypto_event_analysis {

		#Optional
		is_enabled = true
		summarized_events_log {
			#Required
			log_group_id = oci_logging_log_group.test_log_group.id
			log_id = oci_logging_log.test_log.id
		}
	}
	java_migration_analysis {

		#Optional
		is_enabled = true
	}
	jfr_recording {

		#Optional
		is_enabled = true
	}
	lcm {

		#Optional
		is_enabled = true
		post_installation_actions {

			#Optional
			add_logging_handler   = false
      disabled_tls_versions = ["TLS_1_0"]
      global_logging_level  = "ALL"
			minimum_key_size_settings {

				#Optional
				certpath {

					#Optional
					key_size = "2048"
          name     = "RSA"
				}
				jar {

					#Optional
					key_size = "2048"
          name     = "RSA"
				}
				tls {

					#Optional
					key_size = "2048"
          name     = "RSA"
				}
			}
			proxies {

				#Optional
				ftp_proxy_host     = "example-ftp-proxy-host"
        ftp_proxy_port     = "10"
        http_proxy_host    = "example-http-proxy-host"
        http_proxy_port    = "10"
        https_proxy_host   = "example-https-proxy-host"
        https_proxy_port   = "10"
        socks_proxy_host   = "example-socks-proxy-host"
        socks_proxy_port   = "10"
        use_system_proxies = "false"
			}
			should_replace_certificates_operating_system = false
		}
	}
	performance_tuning_analysis {

		#Optional
		is_enabled = true
	}
}
```

## Argument Reference

The following arguments are supported:

* `advanced_usage_tracking` - (Optional) (Updatable) AdvancedUsageTracking configuration
	* `is_enabled` - (Optional) (Updatable) AdvancedUsageTracking flag to store enabled or disabled status.
* `analytic_bucket_name` - (Optional) (Updatable) Bucket name required to store JFR and related data.
* `analytic_namespace` - (Optional) (Updatable) Namespace for the Fleet advanced feature.
* `crypto_event_analysis` - (Optional) (Updatable) CryptoEventAnalysis configuration
	* `is_enabled` - (Optional) (Updatable) CryptoEventAnalysis flag to store enabled or disabled status.
	* `summarized_events_log` - (Optional) (Updatable) Summarized events log for advanced feature. 
		* `log_group_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
		* `log_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `java_migration_analysis` - (Optional) (Updatable) JavaMigrationAnalysis configuration
	* `is_enabled` - (Optional) (Updatable) JavaMigrationAnalysis flag to store enabled or disabled status.
* `jfr_recording` - (Optional) (Updatable) JfrRecording configuration
	* `is_enabled` - (Optional) (Updatable) JfrRecording flag to store enabled or disabled status.
* `lcm` - (Optional) (Updatable) Enable lifecycle management and set post action configurations.
	* `is_enabled` - (Optional) (Updatable) Lifecycle management flag to store enabled or disabled status.
	* `post_installation_actions` - (Optional) (Updatable) List of available post actions you can execute after the successful Java installation. 
		* `add_logging_handler` - (Optional) (Updatable) Sets FileHandler and ConsoleHandler as handlers in logging.properties file. 
		* `disabled_tls_versions` - (Optional) (Updatable) The following post JRE installation actions are supported by the field:
			* Disable TLS 1.0 , TLS 1.1 
		* `global_logging_level` - (Optional) (Updatable) Sets the logging level in logging.properties file. 
		* `minimum_key_size_settings` - (Optional) (Updatable) test
			* `certpath` - (Optional) (Updatable) Updates the minimum key size for the specified encryption algorithm. The JDK property jdk.certpath.disabledAlgorithms will be updated with the following supported actions:
				* Changing minimum key length for RSA signed jars
				* Changing minimum key length for EC
				* Changing minimum key length for DSA 
				* `key_size` - (Optional) (Updatable) Key size for the encryption algorithm. Allowed values: 256 for EC, 2048 for DH/DSA/RSA 
				* `name` - (Optional) (Updatable) The algorithm name.
			* `jar` - (Optional) (Updatable) Updates the minimum key size for the specified encryption algorithm. The JDK property jdk.jar.disabledAlgorithms will be updated with the following supported actions:
				* Changing minimum key length for RSA signed jars
				* Changing minimum key length for EC
				* Changing minimum key length for DSA 
				* `key_size` - (Optional) (Updatable) Key size for the encryption algorithm. Allowed values: 256 for EC, 2048 for DH/DSA/RSA 
				* `name` - (Optional) (Updatable) The algorithm name.
			* `tls` - (Optional) (Updatable) Updates the minimum key size for the specified encryption algorithm. The JDK property jdk.tls.disabledAlgorithms will be updated with the following supported actions:
				* Changing minimum key length for Diffie-Hellman 
				* `key_size` - (Optional) (Updatable) Key size for the encryption algorithm. Allowed values: 256 for EC, 2048 for DH/DSA/RSA 
				* `name` - (Optional) (Updatable) The algorithm name.
		* `proxies` - (Optional) (Updatable) List of proxy properties to be configured in net.properties file. 
			* `ftp_proxy_host` - (Optional) (Updatable) Ftp host to be set in net.properties file. 
			* `ftp_proxy_port` - (Optional) (Updatable) Ftp port number to be set in net.properties file. 
			* `http_proxy_host` - (Optional) (Updatable) Http host to be set in net.properties file. 
			* `http_proxy_port` - (Optional) (Updatable) Http port number to be set in net.properties file. 
			* `https_proxy_host` - (Optional) (Updatable) Https host to be set in net.properties file. 
			* `https_proxy_port` - (Optional) (Updatable) Https port number to be set in net.properties file. 
			* `socks_proxy_host` - (Optional) (Updatable) Socks host to be set in net.properties file. 
			* `socks_proxy_port` - (Optional) (Updatable) Socks port number to be set in net.properties file. 
			* `use_system_proxies` - (Optional) (Updatable) Sets "java.net.useSystemProxies=true" in net.properties when they exist. 
		* `should_replace_certificates_operating_system` - (Optional) (Updatable) Restores JDK root certificates with the certificates that are available in the operating system. The following action is supported by the field:
			* Replace JDK root certificates with a list provided by the operating system. 
* `performance_tuning_analysis` - (Optional) (Updatable) Performance tuning analysis configuration
	* `is_enabled` - (Optional) (Updatable) PerformanceTuningAnalysis flag to store enabled or disabled status


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `advanced_usage_tracking` - AdvancedUsageTracking configuration
	* `is_enabled` - AdvancedUsageTracking flag to store enabled or disabled status.
* `analytic_bucket_name` - Bucket name required to store JFR and related data.
* `analytic_namespace` - Namespace for the Fleet advanced feature.
* `crypto_event_analysis` - CryptoEventAnalysis configuration
	* `is_enabled` - CryptoEventAnalysis flag to store enabled or disabled status.
	* `summarized_events_log` - Summarized events log for advanced feature. 
		* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `java_migration_analysis` - JavaMigrationAnalysis configuration
	* `is_enabled` - JavaMigrationAnalysis flag to store enabled or disabled status.
* `jfr_recording` - JfrRecording configuration
	* `is_enabled` - JfrRecording flag to store enabled or disabled status.
* `lcm` - Enable lifecycle management and set post action configurations.
	* `is_enabled` - Lifecycle management flag to store enabled or disabled status.
	* `post_installation_actions` - List of available post actions you can execute after the successful Java installation. 
		* `add_logging_handler` - Sets FileHandler and ConsoleHandler as handlers in logging.properties file. 
		* `disabled_tls_versions` - The following post JRE installation actions are supported by the field:
			* Disable TLS 1.0 , TLS 1.1 
		* `global_logging_level` - Sets the logging level in logging.properties file. 
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
		* `proxies` - List of proxy properties to be configured in net.properties file. 
			* `ftp_proxy_host` - Ftp host to be set in net.properties file. 
			* `ftp_proxy_port` - Ftp port number to be set in net.properties file. 
			* `http_proxy_host` - Http host to be set in net.properties file. 
			* `http_proxy_port` - Http port number to be set in net.properties file. 
			* `https_proxy_host` - Https host to be set in net.properties file. 
			* `https_proxy_port` - Https port number to be set in net.properties file. 
			* `socks_proxy_host` - Socks host to be set in net.properties file. 
			* `socks_proxy_port` - Socks port number to be set in net.properties file. 
			* `use_system_proxies` - Sets "java.net.useSystemProxies=true" in net.properties when they exist. 
		* `should_replace_certificates_operating_system` - Restores JDK root certificates with the certificates that are available in the operating system. The following action is supported by the field:
			* Replace JDK root certificates with a list provided by the operating system. 
* `performance_tuning_analysis` - Performance tuning analysis configuration
	* `is_enabled` - PerformanceTuningAnalysis flag to store enabled or disabled status
* `time_last_modified` - The date and time of the last modification to the Fleet Agent Configuration (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fleet Advanced Feature Configuration
	* `update` - (Defaults to 20 minutes), when updating the Fleet Advanced Feature Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Fleet Advanced Feature Configuration


## Import

FleetAdvancedFeatureConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_jms_fleet_advanced_feature_configuration.test_fleet_advanced_feature_configuration "fleets/{fleetId}/advancedFeatureConfiguration" 
```

