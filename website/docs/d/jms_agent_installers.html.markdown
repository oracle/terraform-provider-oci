---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_agent_installers"
sidebar_current: "docs-oci-datasource-jms-agent_installers"
description: |-
  Provides the list of Agent Installers in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_agent_installers
This data source provides the list of Agent Installers in Oracle Cloud Infrastructure Jms service.

Returns a list of the agent installer information.


## Example Usage

```hcl
data "oci_jms_agent_installers" "test_agent_installers" {

	#Optional
	compartment_id = var.compartment_id
	fleet_id = oci_jms_fleet.test_fleet.id
	os_family = var.agent_installer_os_family
	platform_architecture = var.agent_installer_platform_architecture
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `fleet_id` - (Optional) The ID of the Fleet.
* `os_family` - (Optional) The OS family for the agent installer.
* `platform_architecture` - (Optional) The platform architecture for the agent installer.


## Attributes Reference

The following attributes are exported:

* `agent_installer_collection` - The list of agent_installer_collection.

### AgentInstaller Reference

The following attributes are exported:

* `items` - A list of the agent installer summaries.
	* `agent_installer_description` - Description of the agent installer artifact. The description typically includes the OS, architecture, and agent installer type.
	* `agent_installer_id` - Unique identifier for the agent installer.
	* `agent_installer_version` - Agent installer version.
	* `agent_version` - Agent image version.
	* `approximate_file_size_in_bytes` - Approximate compressed file size in bytes.
	* `java_version` - Java version.
	* `os_family` - The target operating system family for the agent installer.
	* `package_type` - The package type (typically the file extension) of the agent software included in the installer.
	* `platform_architecture` - The target operating system architecture for the installer.
	* `sha256` - SHA256 checksum of the agent installer.

