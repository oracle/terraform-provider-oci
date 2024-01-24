// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	NlbBackendTargetRequiredOnlyResource = NlbBackendTargetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", acctest.Required, acctest.Create, nlbBackendTargetRepresentation)

	NlbBackendTargetResourceConfig = NlbBackendTargetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", acctest.Optional, acctest.Update, nlbBackendTargetRepresentation)

	nlbBackendTargetRepresentation = map[string]interface{}{
		"backend_set_name":         acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"port":                     acctest.Representation{RepType: acctest.Required, Create: `10`},
		"target_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"is_backup":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_drain":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_offline":               acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"name":                     acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"weight":                   acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	NlbBackendTargetResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, NetworkLoadBalancerBackendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerBackendTargetIdResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerBackendTargetIdResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_load_balancer_backend.test_backend"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckNetworkLoadBalancerBackendDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + NlbBackendTargetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", acctest.Required, acctest.Create, nlbBackendTargetRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "backend_set_name"),
					resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "target_id"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
					acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", acctest.Optional, acctest.Create, nlbBackendTargetRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + NlbBackendTargetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", acctest.Optional, acctest.Update, nlbBackendTargetRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
}
