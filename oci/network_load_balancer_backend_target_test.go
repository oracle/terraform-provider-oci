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
		generateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", Required, Create, nlbBackendTargetRepresentation)

	NlbBackendTargetResourceConfig = NlbBackendTargetResourceDependencies +
		generateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", Optional, Update, nlbBackendTargetRepresentation)

	nlbBackendTargetRepresentation = map[string]interface{}{
		"backend_set_name":         Representation{repType: Required, create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"network_load_balancer_id": Representation{repType: Required, create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"port":                     Representation{repType: Required, create: `10`},
		"target_id":                Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"is_backup":                Representation{repType: Optional, create: `false`, update: `true`},
		"is_drain":                 Representation{repType: Optional, create: `false`, update: `true`},
		"is_offline":               Representation{repType: Optional, create: `false`, update: `true`},
		"name":                     Representation{repType: Optional, create: `name`},
		"weight":                   Representation{repType: Required, create: `10`, update: `11`},
	}

	NlbBackendTargetResourceDependencies = OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		generateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", Required, Create, nlbBackendSetRepresentation) +
		generateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Required, Create, networkLoadBalancerRepresentation) +
		AvailabilityDomainConfig
)

func TestNetworkLoadBalancerBackendTargetIdResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerBackendTargetIdResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_load_balancer_backend.test_backend"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckNetworkLoadBalancerBackendDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + NlbBackendTargetResourceDependencies +
					generateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", Required, Create, nlbBackendTargetRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "backend_set_name"),
					resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "target_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + NlbBackendTargetResourceDependencies,
			},

			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + NlbBackendTargetResourceDependencies +
					generateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", Optional, Create, nlbBackendTargetRepresentation),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", Optional, Update, nlbBackendTargetRepresentation),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
		initDependencyGraph()
	}
}
