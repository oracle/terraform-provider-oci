// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreInstancePoolRequiredOnlyResource = CoreInstancePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Required, acctest.Create, CoreInstancePoolRepresentation)

	CoreInstancePoolResourceConfig = CoreInstancePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, CoreInstancePoolRepresentation)

	CoreCoreInstancePoolSingularDataSourceRepresentation = map[string]interface{}{
		"instance_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_pool.test_instance_pool.id}`},
	}

	CoreCoreInstancePoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `RUNNING`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstancePoolDataSourceFilterRepresentation}}
	CoreInstancePoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_instance_pool.test_instance_pool.id}`}},
	}

	CoreInstancePoolRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_configuration_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":        acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstancePoolPlacementConfigurationsRepresentation},
		"size":                            acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"state":                           acctest.Representation{RepType: acctest.Optional, Create: `Running`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`, Update: `displayName2`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"load_balancers":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstancePoolLoadBalancersRepresentation},
		"instance_display_name_formatter": acctest.Representation{RepType: acctest.Optional, Create: `host-$${launchCount}`, Update: `host2-$${launchCount}`},
		"instance_hostname_formatter":     acctest.Representation{RepType: acctest.Optional, Create: `host-$${launchCount}`, Update: `host2-$${launchCount}`},
	}

	CoreInstancePoolRepresentationIPv6 = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":  acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstancePoolPlacementConfigurationsRepresentationIpv6},
		"size":                      acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `Running`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`, Update: `displayName2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"load_balancers":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstancePoolLoadBalancersRepresentation},
	}

	CoreInstancePoolRepresentationIPv6WithPrimaryVnicSubnets = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":  acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstancePoolPlacementConfigurationsRepresentationIpv6WithPrimaryVnicSubnets},
		"size":                      acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `Running`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`, Update: `displayName2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"load_balancers":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstancePoolLoadBalancersRepresentation},
	}

	CoreInstancePoolRepresentationWithLifecycleSizeIgnoreChanges = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":  acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstancePoolPlacementConfigurationsRepresentation},
		"size":                      acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `Running`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`, Update: `displayName2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"load_balancers":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstancePoolLoadBalancersRepresentation},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstancePoolSizeIgnoreChangesRepresentation},
	}
	CoreInstancePoolSizeIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`size`}},
	}

	CoreSubnetRepresentationIpv6 = map[string]interface{}{
		"cidr_block":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/24`, Update: "10.0.0.0/16"},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain":        acctest.Representation{RepType: acctest.Optional, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dhcp_options_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`, Update: `${oci_core_dhcp_options.test_dhcp_options.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `IpvSubnet1`, Update: `IpvSubnet2`},
		"dns_label":                  acctest.Representation{RepType: acctest.Optional, Create: `dnslabel`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"prohibit_internet_ingress":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"ipv6cidr_blocks":            acctest.Representation{RepType: acctest.Optional, Create: []string{`${substr(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn.ipv6cidr_blocks[0]) - 2)}${64}`}},
		"route_table_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`, Update: `${oci_core_route_table.test_route_table.id}`},
		"security_list_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}, Update: []string{`${oci_core_security_list.test_security_list.id}`}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	CoreInstancePoolPlacementConfigurationsRepresentation = map[string]interface{}{
		"availability_domain":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"fault_domains":          acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-1`}, Update: []string{`FAULT-DOMAIN-2`}},
		"primary_subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"secondary_vnic_subnets": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentation},
	}
	CoreInstancePoolPlacementConfigurationsRepresentationIpv6 = map[string]interface{}{
		"availability_domain":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"fault_domains":          acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-1`}, Update: []string{`FAULT-DOMAIN-2`}},
		"primary_subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"secondary_vnic_subnets": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentationIpv6},
	}

	CoreInstancePoolPlacementConfigurationsRepresentationIpv6WithPrimaryVnicSubnets = map[string]interface{}{
		"availability_domain":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"fault_domains":          acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-1`}, Update: []string{`FAULT-DOMAIN-2`}},
		"primary_vnic_subnets":   acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstancePoolPlacementConfigurationsPrimaryVnicSubnetsRepresentation},
		"secondary_vnic_subnets": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentationIpv6},
	}

	CoreInstancePoolPlacementConfigurationsPrimaryVnicSubnetsRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"ipv6address_ipv6subnet_cidr_pair_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreIpv6AddressIpv6SubnetCidrPairRepresentation},
		"is_assign_ipv6ip":                         acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	CoreInstancePoolLoadBalancersRepresentation = map[string]interface{}{
		"backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"port":             acctest.Representation{RepType: acctest.Required, Create: `10`},
		"vnic_selection":   acctest.Representation{RepType: acctest.Required, Create: `PrimaryVnic`},
	}
	CoreInstancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentation = map[string]interface{}{
		"subnet_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `backend-servers-pool`},
	}
	CoreInstancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentationIpv6 = map[string]interface{}{
		"subnet_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`, Update: `displayName2`},
		"ipv6address_ipv6subnet_cidr_pair_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreIpv6AddressIpv6SubnetCidrPairRepresentation},
		"is_assign_ipv6ip":                         acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	CoreInstancePoolLoadBalancers2Representation = map[string]interface{}{
		"backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set2.name}`},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer2.id}`},
		"port":             acctest.Representation{RepType: acctest.Required, Create: `10`},
		"vnic_selection":   acctest.Representation{RepType: acctest.Required, Create: `PrimaryVnic`},
	}

	CoreInstancePoolConfigurationPoolRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstancePoolInstanceConfigurationInstanceDetailsPoolRepresentation},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	CoreInstancePoolInstanceConfigurationInstanceDetailsPoolRepresentation = map[string]interface{}{
		"instance_type":   acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"secondary_vnics": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstancePoolInstanceConfigurationInstanceDetailsSecondaryVnicsPoolRepresentation},
		"launch_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstancePoolInstanceConfigurationInstanceDetailsLaunchDetailsPoolRepresentation},
	}
	CoreInstancePoolInstanceConfigurationInstanceDetailsSecondaryVnicsPoolRepresentation = map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstancePoolInstanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation},
		//the display_name should be the same as in the secondary_vnic_subnets
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`},
	}
	CoreInstancePoolInstanceConfigurationInstanceDetailsLaunchDetailsPoolRepresentation = map[string]interface{}{
		"compartment_id":                      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"create_vnic_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstancePoolInstanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"extended_metadata":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"extendedMetadata": "extendedMetadata"}, Update: map[string]string{"extendedMetadata2": "extendedMetadata2"}},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"metadata": "metadata"}, Update: map[string]string{"metadata2": "metadata2"}},
		"shape":                               acctest.Representation{RepType: acctest.Optional, Create: InstanceConfigurationVmShape},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation},
		"agent_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceLaunchOptionsRepresentationOnlyNetworkType},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_mode":                         acctest.Representation{RepType: acctest.Optional, Create: `NATIVE`},
		"preferred_maintenance_action":        acctest.Representation{RepType: acctest.Optional, Create: `LIVE_MIGRATE`},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
	}
	CoreInstancePoolInstanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation = map[string]interface{}{
		"assign_public_ip":       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"skip_source_dest_check": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	CoreVcnRepresentationIpv6 = map[string]interface{}{
		"cidr_block":     acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `NewVcn`, Update: `LatestVcn`},
		"dns_label":      acctest.Representation{RepType: acctest.Optional, Create: `dnslabel`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
		"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	CoreInstancePoolResourceDependenciesWithoutSecondaryVnic = CoreSubnetResourceConfig + utils.OciImageIdsVariable + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}` +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
			acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("launch_details.launch_options", instanceLaunchOptionsRepresentationForInstanceConfiguration, CoreInstancePoolInstanceConfigurationInstanceDetailsPoolRepresentation), []string{"secondary_vnics"})}, CoreInstancePoolConfigurationPoolRepresentation))

	CoreInstancePoolResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, acctest.GetUpdatedRepresentationCopy("instance_details.launch_details.launch_options", instanceLaunchOptionsRepresentationForInstanceConfiguration, CoreInstancePoolConfigurationPoolRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, CoreVcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set2", acctest.Required, acctest.Create, backendSet2Representation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer2", acctest.Required, acctest.Create, loadBalancer2Representation) +
		LoadBalancerSubnetDependencies

	CoreInstancePoolResourceDependenciesIpv6 = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, acctest.GetUpdatedRepresentationCopy("instance_details.launch_details.launch_options", instanceLaunchOptionsRepresentationForInstanceConfiguration, CoreInstancePoolConfigurationPoolRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, CoreSubnetRepresentationIpv6) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, CoreVcnRepresentationIpv6) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set2", acctest.Required, acctest.Create, backendSet2Representation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer2", acctest.Required, acctest.Create, loadBalancer2Representation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: core/computeManagement
func TestCoreInstancePoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance_pool.test_instance_pool"
	datasourceName := "data.oci_core_instance_pools.test_instance_pools"
	singularDatasourceName := "data.oci_core_instance_pool.test_instance_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreInstancePoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Create, CoreInstancePoolRepresentation), "core", "instancePool", t)

	acctest.ResourceTest(t, testAccCheckCoreInstancePoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreInstancePoolResourceDependenciesWithoutSecondaryVnic +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Required, acctest.Create, CoreInstancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "2"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreInstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Create, CoreInstancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_display_name_formatter", "host-${launchCount}"),
				resource.TestCheckResourceAttr(resourceName, "instance_hostname_formatter", "host-${launchCount}"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreInstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreInstancePoolRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_display_name_formatter", "host-${launchCount}"),
				resource.TestCheckResourceAttr(resourceName, "instance_hostname_formatter", "host-${launchCount}"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + CoreInstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, CoreInstancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_display_name_formatter", "host2-${launchCount}"),
				resource.TestCheckResourceAttr(resourceName, "instance_hostname_formatter", "host2-${launchCount}"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "3"),
				resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify attach
		{
			Config: config + compartmentIdVariableStr + CoreInstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreInstancePoolRepresentation, map[string]interface{}{
					"load_balancers": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: CoreInstancePoolLoadBalancersRepresentation}, {RepType: acctest.Optional, Group: CoreInstancePoolLoadBalancers2Representation}},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.1.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.1.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.1.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.1.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.1.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.1.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.1.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "3"),
				resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify detach
		{
			Config: config + compartmentIdVariableStr + CoreInstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, CoreInstancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "3"),
				resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify stop the Instance Pool
		{
			Config: config + compartmentIdVariableStr + CoreInstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("state", acctest.Representation{RepType: acctest.Optional, Create: "Stopped"}, CoreInstancePoolRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "3"),
				resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify start the Instance Pool
		{
			Config: config + compartmentIdVariableStr + CoreInstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, CoreInstancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "3"),
				resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource the state will be updated to RUNNING
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_pools", "test_instance_pools", acctest.Optional, acctest.Update, CoreCoreInstancePoolDataSourceRepresentation) +
				compartmentIdVariableStr + CoreInstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, CoreInstancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "instance_pools.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.instance_configuration_id"),
				resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.size", "3"),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Required, acctest.Create, CoreCoreInstancePoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreInstancePoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_display_name_formatter", "host2-${launchCount}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_hostname_formatter", "host2-${launchCount}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.secondary_vnic_subnets.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.secondary_vnic_subnets.0.display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttr(singularDatasourceName, "size", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreInstancePoolRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreInstancePoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_instance_pool" {
			noResourceFound = false
			request := oci_core.GetInstancePoolRequest{}

			tmp := rs.Primary.ID
			request.InstancePoolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetInstancePool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.InstancePoolLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreInstancePool") {
		resource.AddTestSweepers("CoreInstancePool", &resource.Sweeper{
			Name:         "CoreInstancePool",
			Dependencies: acctest.DependencyGraph["instancePool"],
			F:            sweepCoreInstancePoolResource,
		})
	}
}

func sweepCoreInstancePoolResource(compartment string) error {
	computeManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeManagementClient()
	instancePoolIds, err := getInstancePoolIds(compartment)
	if err != nil {
		return err
	}
	for _, instancePoolId := range instancePoolIds {
		if ok := acctest.SweeperDefaultResourceId[instancePoolId]; !ok {
			terminateInstancePoolRequest := oci_core.TerminateInstancePoolRequest{}

			terminateInstancePoolRequest.InstancePoolId = &instancePoolId

			terminateInstancePoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeManagementClient.TerminateInstancePool(context.Background(), terminateInstancePoolRequest)
			if error != nil {
				fmt.Printf("Error deleting InstancePool %s %s, It is possible that the resource is already deleted. Please verify manually \n", instancePoolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &instancePoolId, instancePoolSweepWaitCondition, time.Duration(3*time.Minute),
				instancePoolSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getInstancePoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InstancePoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeManagementClient()

	listInstancePoolsRequest := oci_core.ListInstancePoolsRequest{}
	listInstancePoolsRequest.CompartmentId = &compartmentId
	listInstancePoolsRequest.LifecycleState = oci_core.InstancePoolSummaryLifecycleStateRunning
	listInstancePoolsResponse, err := computeManagementClient.ListInstancePools(context.Background(), listInstancePoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting InstancePool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, instancePool := range listInstancePoolsResponse.Items {
		id := *instancePool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InstancePoolId", id)
	}
	return resourceIds, nil
}

// issue-routing-tag: core/computeManagement
func TestCoreInstancePoolResourceIpv6_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolResourceIpv6_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance_pool.test_instance_pool"

	singularDatasourceName := "data.oci_core_instance_pool.test_instance_pool"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreInstancePoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Create, CoreInstancePoolRepresentationIPv6), "core", "instancePool", t)

	acctest.ResourceTest(t, testAccCheckCoreInstancePoolDestroy, []resource.TestStep{
		// verify Create with ipv6 supported primarySubnets and secondarySubnets
		{
			Config: config + compartmentIdVariableStr + CoreInstancePoolResourceDependenciesIpv6 +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Create, CoreInstancePoolRepresentationIPv6WithPrimaryVnicSubnets),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.primary_vnic_subnets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_vnic_subnets.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.primary_vnic_subnets.0.ipv6address_ipv6subnet_cidr_pair_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_vnic_subnets.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6subnet_cidr"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.primary_vnic_subnets.0.is_assign_ipv6ip", "true"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.ipv6address_ipv6subnet_cidr_pair_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6subnet_cidr"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.is_assign_ipv6ip", "true"),
				resource.TestCheckResourceAttr(resourceName, "size", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Required, acctest.Create, CoreCoreInstancePoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreInstancePoolResourceDependenciesIpv6 +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Required, acctest.Create, CoreInstancePoolRepresentationIPv6WithPrimaryVnicSubnets),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.primary_vnic_subnets.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configurations.0.primary_vnic_subnets.0.subnet_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.primary_vnic_subnets.0.ipv6address_ipv6subnet_cidr_pair_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configurations.0.primary_vnic_subnets.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6subnet_cidr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.primary_vnic_subnets.0.is_assign_ipv6ip", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.secondary_vnic_subnets.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.secondary_vnic_subnets.0.display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configurations.0.secondary_vnic_subnets.0.subnet_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.secondary_vnic_subnets.0.ipv6address_ipv6subnet_cidr_pair_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configurations.0.secondary_vnic_subnets.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6subnet_cidr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.secondary_vnic_subnets.0.is_assign_ipv6ip", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "size", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}

func instancePoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if instancePoolResponse, ok := response.Response.(oci_core.GetInstancePoolResponse); ok {
		return instancePoolResponse.LifecycleState != oci_core.InstancePoolLifecycleStateTerminated
	}
	return false
}

func instancePoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeManagementClient().GetInstancePool(context.Background(), oci_core.GetInstancePoolRequest{
		InstancePoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
