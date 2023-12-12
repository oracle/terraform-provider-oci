// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_container_instances "github.com/oracle/oci-go-sdk/v65/containerinstances"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerInstancesContainerInstanceRequiredOnlyResource = ContainerInstancesContainerInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Required, acctest.Create, ContainerInstancesContainerInstanceRepresentation)

	ContainerInstancesContainerInstanceResourceConfig = ContainerInstancesContainerInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Optional, acctest.Update, ContainerInstancesContainerInstanceRepresentation)

	ContainerInstancesContainerInstancesContainerInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"container_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_container_instances_container_instance.test_container_instance.id}`},
	}

	ContainerInstancesContainerInstancesContainerInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `INACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceDataSourceFilterRepresentation},
	}
	ContainerInstancesContainerInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_container_instances_container_instance.test_container_instance.id}`}},
	}

	ContainerInstancesContainerInstanceRepresentation = map[string]interface{}{
		"availability_domain":                  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"containers":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceContainersRepresentation},
		"shape":                                acctest.Representation{RepType: acctest.Required, Create: `CI.Standard.E4.Flex`},
		"shape_config":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceShapeConfigRepresentation},
		"vnics":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceVnicsRepresentation},
		"container_restart_policy":             acctest.Representation{RepType: acctest.Optional, Create: `ALWAYS`},
		"defined_tags":                         acctest.Representation{RepType: acctest.Optional, Create: `${map("tf_test_namespace.test_tag", "value")}`, Update: `${map("tf_test_namespace.test_tag", "updatedValue")}`},
		"display_name":                         acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"dns_config":                           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceDnsConfigRepresentation},
		"freeform_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"graceful_shutdown_timeout_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"volumes":                              acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceEmptyDirVolumesRepresentation},
		"state":                                acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `INACTIVE`},
		"lifecycle":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesCIDefinedTagsRepresentation},
	}
	ContainerInstancesContainerInstanceContainersRepresentation = map[string]interface{}{
		"image_url":                      acctest.Representation{RepType: acctest.Required, Create: `busybox`},
		"arguments":                      acctest.Representation{RepType: acctest.Optional, Create: []string{`-c`, `sleep 24h`}},
		"command":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`/bin/sh`}},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"environment_variables":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"environmentVariables": "environmentVariables"}},
		"health_checks":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersHealthChecksRepresentation},
		"is_resource_principal_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"resource_config":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersResourceConfigRepresentation},
		"security_context":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersSecurityContextRepresentation},
		"volume_mounts":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersEmptyDirVolumeMountsRepresentation},
		"working_directory":              acctest.Representation{RepType: acctest.Optional, Create: `/mnt`},
	}
	ContainerInstancesContainerInstanceContainersSecondRepresentation = map[string]interface{}{
		"image_url":                      acctest.Representation{RepType: acctest.Required, Create: `busybox`},
		"arguments":                      acctest.Representation{RepType: acctest.Optional, Create: []string{`-c`, `sleep 24h`}},
		"command":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`/bin/sh`}},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `secondDisplayName`, Update: `secondDisplayName2`},
		"environment_variables":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"environmentVariables": "environmentVariables"}},
		"health_checks":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersHealthChecksRepresentation},
		"is_resource_principal_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"resource_config":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersResourceConfigRepresentation},
		"volume_mounts":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersEmptyDirVolumeMountsRepresentation},
		"working_directory":              acctest.Representation{RepType: acctest.Optional, Create: `/mnt`},
	}
	ContainerInstancesContainerInstanceContainersGoodConfigFileRepresentation = map[string]interface{}{
		"image_url":                      acctest.Representation{RepType: acctest.Required, Create: `busybox`},
		"arguments":                      acctest.Representation{RepType: acctest.Optional, Create: []string{`-c`, `sleep 24h`}},
		"command":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`/bin/sh`}},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `secondDisplayName`, Update: `secondDisplayName2`},
		"environment_variables":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"environmentVariables": "environmentVariables"}},
		"health_checks":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersHealthChecksRepresentation},
		"is_resource_principal_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"resource_config":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersResourceConfigRepresentation},
		"volume_mounts":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersGoodConfigFileVolumeMountsRepresentation},
		"working_directory":              acctest.Representation{RepType: acctest.Optional, Create: `/mnt`},
	}
	ContainerInstancesContainerInstanceShapeConfigRepresentation = map[string]interface{}{
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `4`},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `8`},
	}
	ContainerInstancesContainerInstanceVnicsRepresentation = map[string]interface{}{
		"subnet_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("tf_test_namespace.test_tag", "value")}`, Update: `${map("tf_test_namespace.test_tag", "updatedValue")}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"hostname_label":         acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`, Update: `hostnamelabel2`},
		"is_public_ip_assigned":  acctest.Representation{RepType: acctest.Required, Create: `true`},
		"nsg_ids":                acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		"private_ip":             acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.7`},
		"skip_source_dest_check": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	ContainerInstancesContainerInstanceDnsConfigRepresentation = map[string]interface{}{
		"nameservers": acctest.Representation{RepType: acctest.Optional, Create: []string{`8.8.8.8`}},
		"options":     acctest.Representation{RepType: acctest.Optional, Create: []string{`options`}},
		"searches":    acctest.Representation{RepType: acctest.Optional, Create: []string{`search domain`}},
	}
	ContainerInstancesContainerInstanceEmptyDirVolumesRepresentation = map[string]interface{}{
		"name":          acctest.Representation{RepType: acctest.Required, Create: `volumeName`},
		"volume_type":   acctest.Representation{RepType: acctest.Required, Create: `EMPTYDIR`},
		"backing_store": acctest.Representation{RepType: acctest.Optional, Create: `EPHEMERAL_STORAGE`},
	}
	ContainerInstancesContainerInstanceGoodConfigFileVolumesRepresentation = map[string]interface{}{
		"name":        acctest.Representation{RepType: acctest.Required, Create: `volumeGoodConfigFile`},
		"volume_type": acctest.Representation{RepType: acctest.Required, Create: `CONFIGFILE`},
		"configs":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceGoodConfigFileVolumesConfigsRepresentation},
	}
	ContainerInstancesContainerInstanceGoodConfigFileVolumesConfigsRepresentation = map[string]interface{}{
		"data":      acctest.Representation{RepType: acctest.Required, Create: `T0NJ`},
		"file_name": acctest.Representation{RepType: acctest.Required, Create: `my_file`},
	}
	ContainerInstancesContainerInstanceContainersHealthChecksRepresentation = map[string]interface{}{
		"health_check_type":        acctest.Representation{RepType: acctest.Required, Create: `HTTP`},
		"failure_action":           acctest.Representation{RepType: acctest.Optional, Create: `KILL`},
		"failure_threshold":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"headers":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceContainersHealthChecksHeadersRepresentation},
		"initial_delay_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"interval_in_seconds":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"name":                     acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"path":                     acctest.Representation{RepType: acctest.Optional, Create: `path`},
		"port":                     acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"success_threshold":        acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"timeout_in_seconds":       acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	ContainerInstancesContainerInstanceContainersResourceConfigRepresentation = map[string]interface{}{
		"memory_limit_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"vcpus_limit":         acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}
	ContainerInstancesContainerInstanceContainersSecurityContextRepresentation = map[string]interface{}{
		"is_non_root_user_check_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_root_file_system_readonly":   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"run_as_group":                   acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"run_as_user":                    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"security_context_type":          acctest.Representation{RepType: acctest.Optional, Create: `LINUX`},
	}
	ContainerInstancesContainerInstanceContainersEmptyDirVolumeMountsRepresentation = map[string]interface{}{
		"mount_path":   acctest.Representation{RepType: acctest.Required, Create: `/mnt`},
		"volume_name":  acctest.Representation{RepType: acctest.Required, Create: `volumeName`},
		"is_read_only": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"partition":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"sub_path":     acctest.Representation{RepType: acctest.Optional, Create: `/subPath`},
	}
	ContainerInstancesContainerInstanceContainersGoodConfigFileVolumeMountsRepresentation = map[string]interface{}{
		"mount_path":  acctest.Representation{RepType: acctest.Required, Create: `/mnt`},
		"volume_name": acctest.Representation{RepType: acctest.Required, Create: `volumeGoodConfigFile`},
	}

	//check how this works for multiple containers
	ignoreChangesCIDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `vnics[0].defined_tags`}},
	}

	CISubnetRepresentation = map[string]interface{}{
		"cidr_block":        acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/24`},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"dns_label":         acctest.Representation{RepType: acctest.Required, Create: `testsubnet`},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
		"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_sec_list.id}`}},
		"route_table_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`},
	}

	CIVcnRepresentation = map[string]interface{}{
		"cidr_block":     acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `testvcn`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	CISecurityListRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"egress_security_rules":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: CISecurityListTCPEgressSecurityRulesRepresentation}},
		"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: CISecurityListSSHIngressSecurityRulesRepresentation}, {RepType: acctest.Required, Group: CISecurityListHTTPIngressSecurityRulesRepresentation}, {RepType: acctest.Required, Group: CISecurityListICMPIngressSecurityRulesRepresentation}, {RepType: acctest.Required, Group: CISecurityListICMPVcnCidrIngressSecurityRulesRepresentation}},
	}
	ContainerInstancesContainerInstanceContainersHealthChecksHeadersRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`},
	}

	CISecurityListTCPEgressSecurityRulesRepresentation = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `all`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	CISecurityListSSHIngressSecurityRulesRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: CISecurityListSSHIngressSecurityRulesTcpOptionsRepresentation},
	}

	CISecurityListSSHIngressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `22`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `22`},
	}

	CISecurityListHTTPIngressSecurityRulesRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: CISecurityListHTTPIngressSecurityRulesTcpOptionsRepresentation},
	}

	CISecurityListHTTPIngressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `80`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `80`},
	}

	CISecurityListICMPIngressSecurityRulesRepresentation = map[string]interface{}{
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `1`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"stateless":    acctest.Representation{RepType: acctest.Required, Create: `true`},
		"icmp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: CISecurityListICMPIngressSecurityRulesTcpOptionsRepresentation},
	}

	CISecurityListICMPIngressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `3`},
		"code": acctest.Representation{RepType: acctest.Required, Create: `4`},
	}

	CISecurityListICMPVcnCidrIngressSecurityRulesRepresentation = map[string]interface{}{
		"protocol":  acctest.Representation{RepType: acctest.Required, Create: `1`},
		"source":    acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"stateless": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	CIInternetGatewayRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	CIRouteTableRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_rules":    acctest.RepresentationGroup{RepType: acctest.Required, Group: CIRouteTableRouteRulesRepresentation},
	}

	CIRouteTableRouteRulesRepresentation = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_internet_gateway.test_ig.id}`},
		"destination":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
	}

	ContainerInstancesContainerInstanceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CIVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CISubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_sec_list", acctest.Required, acctest.Create, CISecurityListRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_ig", acctest.Required, acctest.Create, CIInternetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CIRouteTableRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: container_instances/default
func TestContainerInstancesContainerInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerInstancesContainerInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_container_instances_container_instance.test_container_instance"
	datasourceName := "data.oci_container_instances_container_instances.test_container_instances"
	singularDatasourceName := "data.oci_container_instances_container_instance.test_container_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerInstancesContainerInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Optional, acctest.Create, ContainerInstancesContainerInstanceRepresentation), "containerinstances", "containerInstance", t)

	acctest.ResourceTest(t, testAccCheckContainerInstancesContainerInstanceDestroy, []resource.TestStep{
		// Create dependencies and wait
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies,
			Check:  delayAndReturnNil(),
		},

		// verify default create and check power on
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Required, acctest.Create, ContainerInstancesContainerInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "containers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.image_url", "busybox"),
				resource.TestCheckResourceAttr(resourceName, "shape", "CI.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "4"),
				resource.TestCheckResourceAttr(resourceName, "vnics.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vnics.0.subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies,
		},

		// verify create with power on
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ContainerInstancesContainerInstanceRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies,
		},

		// verify create with power off
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ContainerInstancesContainerInstanceRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Required, Create: `INACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify update to power on
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ContainerInstancesContainerInstanceRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify update to power off
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ContainerInstancesContainerInstanceRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Required, Create: `INACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Optional, acctest.Create, ContainerInstancesContainerInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "container_count"),
				resource.TestCheckResourceAttr(resourceName, "container_restart_policy", "ALWAYS"),
				resource.TestCheckResourceAttr(resourceName, "containers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.arguments.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.command.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.0.container_id"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.health_checks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.image_url", "busybox"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.is_resource_principal_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.0.memory_limit_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.0.vcpus_limit", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.is_non_root_user_check_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.is_root_file_system_readonly", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.run_as_group", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.run_as_user", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.security_context_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.is_read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.mount_path", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.partition", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.sub_path", "/subPath"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.0.volume_mounts.0.volume_name"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.working_directory", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.0.nameservers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.0.options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.0.searches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "graceful_shutdown_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "CI.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "8"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "4"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.processor_description"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vnics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.is_public_ip_assigned", "true"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.private_ip", "10.0.0.7"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "vnics.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "volumes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.backing_store", "EPHEMERAL_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.name", "volumeName"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.volume_type", "EMPTYDIR"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ContainerInstancesContainerInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "container_count"),
				resource.TestCheckResourceAttr(resourceName, "container_restart_policy", "ALWAYS"),
				resource.TestCheckResourceAttr(resourceName, "containers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.arguments.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.command.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.0.container_id"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.health_checks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.image_url", "busybox"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.is_resource_principal_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.0.memory_limit_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.0.vcpus_limit", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.is_non_root_user_check_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.is_root_file_system_readonly", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.run_as_group", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.run_as_user", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.security_context_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.is_read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.mount_path", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.partition", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.sub_path", "/subPath"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.0.volume_mounts.0.volume_name"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.working_directory", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.0.nameservers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.0.options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.0.searches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "graceful_shutdown_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "CI.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "8"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "4"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.processor_description"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vnics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.is_public_ip_assigned", "true"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.private_ip", "10.0.0.7"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "vnics.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "volumes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.backing_store", "EPHEMERAL_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.name", "volumeName"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.volume_type", "EMPTYDIR"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Optional, acctest.Update, ContainerInstancesContainerInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "container_count"),
				resource.TestCheckResourceAttr(resourceName, "container_restart_policy", "ALWAYS"),
				resource.TestCheckResourceAttr(resourceName, "containers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.arguments.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.command.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.0.container_id"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.health_checks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.image_url", "busybox"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.is_resource_principal_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.0.memory_limit_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.0.vcpus_limit", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.is_non_root_user_check_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.is_root_file_system_readonly", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.run_as_group", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.run_as_user", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.security_context.0.security_context_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.is_read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.mount_path", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.partition", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.sub_path", "/subPath"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.0.volume_mounts.0.volume_name"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.working_directory", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.0.nameservers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.0.options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dns_config.0.searches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "graceful_shutdown_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "CI.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "8"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "4"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.processor_description"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vnics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.hostname_label", "hostnamelabel2"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.is_public_ip_assigned", "true"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.private_ip", "10.0.0.7"),
				resource.TestCheckResourceAttr(resourceName, "vnics.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "vnics.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "volumes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.backing_store", "EPHEMERAL_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.name", "volumeName"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.volume_type", "EMPTYDIR"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_container_instances_container_instances", "test_container_instances", acctest.Optional, acctest.Update, ContainerInstancesContainerInstancesContainerInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Optional, acctest.Update, ContainerInstancesContainerInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "container_instance_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "container_instance_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Required, acctest.Create, ContainerInstancesContainerInstancesContainerInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "container_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "container_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "container_restart_policy", "ALWAYS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "containers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "containers.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "containers.0.container_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_config.0.nameservers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_config.0.options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_config.0.searches.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "graceful_shutdown_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape", "CI.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.memory_in_gbs", "8"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.ocpus", "4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vnics.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vnics.0.vnic_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "volume_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "volumes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "volumes.0.backing_store", "EPHEMERAL_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "volumes.0.name", "volumeName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "volumes.0.volume_type", "EMPTYDIR"),
			),
		},
		// verify resource import
		{
			Config:                  config + ContainerInstancesContainerInstanceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies,
		},
		// verify default create with multiple containers
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ContainerInstancesContainerInstanceRepresentation, map[string]interface{}{
						"containers": []acctest.RepresentationGroup{
							{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceContainersRepresentation},
							{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceContainersSecondRepresentation},
							{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceContainersGoodConfigFileRepresentation},
						},
						"volumes": []acctest.RepresentationGroup{
							{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceEmptyDirVolumesRepresentation},
							{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceGoodConfigFileVolumesRepresentation},
						},
					}),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "container_count"),
				resource.TestCheckResourceAttr(resourceName, "container_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "containers.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.arguments.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.command.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.0.container_id"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.image_url", "busybox"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.is_resource_principal_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.0.memory_limit_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.0.vcpus_limit", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.is_read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.mount_path", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.partition", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.sub_path", "/subPath"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.0.volume_mounts.0.volume_name"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.working_directory", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.arguments.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.command.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.1.container_id"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.display_name", "secondDisplayName"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.image_url", "busybox"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.is_resource_principal_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.resource_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.resource_config.0.memory_limit_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.resource_config.0.vcpus_limit", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.volume_mounts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.volume_mounts.0.is_read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.volume_mounts.0.mount_path", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.volume_mounts.0.partition", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.volume_mounts.0.sub_path", "/subPath"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.1.volume_mounts.0.volume_name"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.working_directory", "/mnt"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "volumes.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.backing_store", "EPHEMERAL_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.name", "volumeName"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.volume_type", "EMPTYDIR"),
				resource.TestCheckResourceAttr(resourceName, "volumes.1.name", "volumeGoodConfigFile"),
				resource.TestCheckResourceAttr(resourceName, "volumes.1.volume_type", "CONFIGFILE"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify update to multiple containers
		{
			Config: config + compartmentIdVariableStr + ContainerInstancesContainerInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_container_instances_container_instance", "test_container_instance", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(ContainerInstancesContainerInstanceRepresentation, map[string]interface{}{
						"containers": []acctest.RepresentationGroup{
							{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceContainersRepresentation},
							{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceContainersSecondRepresentation},
							{RepType: acctest.Required, Group: ContainerInstancesContainerInstanceContainersGoodConfigFileRepresentation},
						},
						"volumes": []acctest.RepresentationGroup{
							{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceEmptyDirVolumesRepresentation},
							{RepType: acctest.Optional, Group: ContainerInstancesContainerInstanceGoodConfigFileVolumesRepresentation},
						},
					}),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "container_count"),
				resource.TestCheckResourceAttr(resourceName, "container_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "containers.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.arguments.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.command.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.0.container_id"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.image_url", "busybox"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.is_resource_principal_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.0.memory_limit_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.resource_config.0.vcpus_limit", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.is_read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.mount_path", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.partition", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.volume_mounts.0.sub_path", "/subPath"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.0.volume_mounts.0.volume_name"),
				resource.TestCheckResourceAttr(resourceName, "containers.0.working_directory", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.arguments.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.command.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.1.container_id"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.display_name", "secondDisplayName2"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.image_url", "busybox"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.is_resource_principal_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.resource_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.resource_config.0.memory_limit_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.resource_config.0.vcpus_limit", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.volume_mounts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.volume_mounts.0.is_read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.volume_mounts.0.mount_path", "/mnt"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.volume_mounts.0.partition", "10"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.volume_mounts.0.sub_path", "/subPath"),
				resource.TestCheckResourceAttrSet(resourceName, "containers.1.volume_mounts.0.volume_name"),
				resource.TestCheckResourceAttr(resourceName, "containers.1.working_directory", "/mnt"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "volumes.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.backing_store", "EPHEMERAL_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.name", "volumeName"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.volume_type", "EMPTYDIR"),
				resource.TestCheckResourceAttr(resourceName, "volumes.1.name", "volumeGoodConfigFile"),
				resource.TestCheckResourceAttr(resourceName, "volumes.1.volume_type", "CONFIGFILE"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
	})
}

func testAccCheckContainerInstancesContainerInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerInstanceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_container_instances_container_instance" {
			noResourceFound = false
			request := oci_container_instances.GetContainerInstanceRequest{}

			tmp := rs.Primary.ID
			request.ContainerInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerinstance")

			response, err := client.GetContainerInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_container_instances.ContainerInstanceLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ContainerInstancesContainerInstance") {
		resource.AddTestSweepers("ContainerInstancesContainerInstance", &resource.Sweeper{
			Name:         "ContainerInstancesContainerInstance",
			Dependencies: acctest.DependencyGraph["containerInstance"],
			F:            sweepContainerInstancesContainerInstanceResource,
		})
	}
}

func sweepContainerInstancesContainerInstanceResource(compartment string) error {
	containerInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerInstanceClient()
	containerInstanceIds, err := getContainerInstancesContainerInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, containerInstanceId := range containerInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[containerInstanceId]; !ok {
			deleteContainerInstanceRequest := oci_container_instances.DeleteContainerInstanceRequest{}

			deleteContainerInstanceRequest.ContainerInstanceId = &containerInstanceId

			deleteContainerInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerinstance")
			_, error := containerInstanceClient.DeleteContainerInstance(context.Background(), deleteContainerInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting ContainerInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", containerInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &containerInstanceId, ContainerInstancesContainerInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				ContainerInstancesContainerInstanceSweepResponseFetchOperation, "containerinstance", true)
		}
	}
	return nil
}

func getContainerInstancesContainerInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ContainerInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerInstanceClient()

	listContainerInstancesRequest := oci_container_instances.ListContainerInstancesRequest{}
	listContainerInstancesRequest.CompartmentId = &compartmentId
	listContainerInstancesRequest.LifecycleState = oci_container_instances.ContainerInstanceLifecycleStateActive
	listContainerInstancesResponse, err := containerInstanceClient.ListContainerInstances(context.Background(), listContainerInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ContainerInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, containerInstance := range listContainerInstancesResponse.Items {
		id := *containerInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ContainerInstanceId", id)
	}
	return resourceIds, nil
}

func ContainerInstancesContainerInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if containerInstanceResponse, ok := response.Response.(oci_container_instances.GetContainerInstanceResponse); ok {
		return containerInstanceResponse.LifecycleState != oci_container_instances.ContainerInstanceLifecycleStateDeleted
	}
	return false
}

func ContainerInstancesContainerInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerInstanceClient().GetContainerInstance(context.Background(), oci_container_instances.GetContainerInstanceRequest{
		ContainerInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func delayAndReturnNil() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Beginning wait ...")
		time.Sleep(90 * time.Second)
		log.Println("Ending wait ...")
		return nil
	}
}
