// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NlbBackendTargetRequiredOnlyResource = NlbBackendTargetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", Required, Create, nlbBackendTargetRepresentation)

	NlbBackendTargetResourceConfig = NlbBackendTargetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", Optional, Update, nlbBackendTargetRepresentation)

	nlbBackendTargetRepresentation = map[string]interface{}{
		"backend_set_name":         Representation{RepType: Required, Create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"network_load_balancer_id": Representation{RepType: Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"port":                     Representation{RepType: Required, Create: `10`},
		"target_id":                Representation{RepType: Required, Create: `${oci_core_instance.test_instance.id}`},
		"is_backup":                Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_drain":                 Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_offline":               Representation{RepType: Optional, Create: `false`, Update: `true`},
		"name":                     Representation{RepType: Optional, Create: `name`},
		"weight":                   Representation{RepType: Required, Create: `10`, Update: `11`},
	}

	NlbBackendTargetResourceDependencies = OciImageIdsVariable +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, RepresentationCopyWithNewProperties(SubnetRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		})) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, RepresentationCopyWithNewProperties(VcnRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		})) +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", Required, Create, nlbBackendSetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Required, Create, networkLoadBalancerRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerBackendTargetIdResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerBackendTargetIdResource_basic")
	defer httpreplay.SaveScenario()

	provider := TestAccProvider
	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_load_balancer_backend.test_backend"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { PreCheck() },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckNetworkLoadBalancerBackendDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + NlbBackendTargetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", Required, Create, nlbBackendTargetRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "backend_set_name"),
					resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "target_id"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + NlbBackendTargetResourceDependencies,
			},

			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + NlbBackendTargetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", Optional, Create, nlbBackendTargetRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "backend_set_name"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "target_id"),
					resource.TestCheckResourceAttr(resourceName, "is_backup", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_drain", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_offline", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "weight", "10"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + NlbBackendTargetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", Optional, Update, nlbBackendTargetRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "backend_set_name"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "target_id"),
					resource.TestCheckResourceAttr(resourceName, "is_backup", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_drain", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_offline", "true"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "weight", "11"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

func init() {
	if DependencyGraph == nil {
		InitDependencyGraph()
	}
}
