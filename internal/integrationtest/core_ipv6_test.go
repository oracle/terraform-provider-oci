// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
	CoreIpv6RequiredOnlyResource = CoreIpv6ResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", acctest.Required, acctest.Create, CoreIpv6Representation)

	CoreIpv6ResourceConfig = CoreIpv6ResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", acctest.Optional, acctest.Update, CoreIpv6Representation)

	CoreCoreIpv6SingularDataSourceRepresentation = map[string]interface{}{
		"ipv6id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipv6.test_ipv6.id}`},
	}

	CoreIpv6DataSourceRepresentation = map[string]interface{}{
		"ip_state":  acctest.Representation{RepType: acctest.Optional, Create: `ipState`},
		"lifetime":  acctest.Representation{RepType: acctest.Optional, Create: `EPHEMERAL`, Update: `RESERVED`},
		"subnet_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"vnic_id":   acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
		"filter":    acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreIpv6DataSourceFilterRepresentation}}

	CoreCoreIpv6DataSourceRepresentation = map[string]interface{}{
		"vnic_id": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
		"filter":  acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreIpv6DataSourceFilterRepresentation}}

	CoreIpv6DataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_ipv6.test_ipv6.id}`}},
	}

	CoreIpv6Representation = map[string]interface{}{
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifetime":       acctest.Representation{RepType: acctest.Optional, Create: `EPHEMERAL`, Update: `RESERVED`},
		"route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.test_route_table.id}`},
		"vnic_id":        acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
	}

	CoreIpv6Representation2 = map[string]interface{}{
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifetime":       acctest.Representation{RepType: acctest.Optional, Create: `EPHEMERAL`, Update: `RESERVED`},
		"route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.test_route_table.id}`},
	}

	CoreIpv6ResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
			"ipv6cidr_block": acctest.Representation{RepType: acctest.Required, Create: `${substr(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn.ipv6cidr_blocks[0]) - 2)}${64}`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
			"is_ipv6enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies + `
	data "oci_core_vnic_attachments" "t" {
		availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		instance_id = "${oci_core_instance.test_instance.id}"
	}

`
)

// issue-routing-tag: core/virtualNetwork
func TestCoreIpv6Resource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpv6Resource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_ipv6.test_ipv6"
	datasourceName := "data.oci_core_ipv6s.test_ipv6s"
	singularDatasourceName := "data.oci_core_ipv6.test_ipv6"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreIpv6ResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", acctest.Optional, acctest.Create, CoreIpv6Representation), "core", "ipv6", t)

	acctest.ResourceTest(t, testAccCheckCoreIpv6Destroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreIpv6ResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", acctest.Required, acctest.Create, CoreIpv6Representation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreIpv6ResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreIpv6ResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", acctest.Optional, acctest.Create, CoreIpv6Representation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", "EPHEMERAL"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

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
			Config: config + compartmentIdVariableStr + CoreIpv6ResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", acctest.Optional, acctest.Update, CoreIpv6Representation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", "RESERVED"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify detach operation
		{
			Config: config + compartmentIdVariableStr + CoreIpv6ResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", acctest.Optional, acctest.Update, CoreIpv6Representation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", "RESERVED"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipv6s", "test_ipv6s", acctest.Optional, acctest.Update, CoreCoreIpv6DataSourceRepresentation) +
				compartmentIdVariableStr + CoreIpv6ResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", acctest.Optional, acctest.Update, CoreIpv6Representation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(datasourceName, "ip_address", "ipAddress"),
				//resource.TestCheckResourceAttr(datasourceName, "ip_state", "ipState"),
				//resource.TestCheckResourceAttr(datasourceName, "lifetime", "lifetime2"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subnet_id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "vnic_id"),

				resource.TestCheckResourceAttr(datasourceName, "ipv6s.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ipv6s.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "ipv6s.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.ip_state"),
				resource.TestCheckResourceAttr(datasourceName, "ipv6s.0.lifetime", "RESERVED"),
				resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.ip_address"),
				resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.vnic_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", acctest.Required, acctest.Create, CoreCoreIpv6SingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreIpv6ResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ipv6id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lifetime", "RESERVED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "route_table_id"),
			),
		},
		// verify resource import
		{
			Config:            config + CoreIpv6RequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"ipv6subnet_cidr",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCoreIpv6Destroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_ipv6" {
			noResourceFound = false
			request := oci_core.GetIpv6Request{}

			tmp := rs.Primary.ID
			request.Ipv6Id = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetIpv6(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.Ipv6LifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreIpv6") {
		resource.AddTestSweepers("CoreIpv6", &resource.Sweeper{
			Name:         "CoreIpv6",
			Dependencies: acctest.DependencyGraph["ipv6"],
			F:            sweepCoreIpv6Resource,
		})
	}
}

func sweepCoreIpv6Resource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	ipv6Ids, err := getCoreIpv6Ids(compartment)
	if err != nil {
		return err
	}
	for _, ipv6Id := range ipv6Ids {
		if ok := acctest.SweeperDefaultResourceId[ipv6Id]; !ok {
			deleteIpv6Request := oci_core.DeleteIpv6Request{}

			deleteIpv6Request.Ipv6Id = &ipv6Id

			deleteIpv6Request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteIpv6(context.Background(), deleteIpv6Request)
			if error != nil {
				fmt.Printf("Error deleting Ipv6 %s %s, It is possible that the resource is already deleted. Please verify manually \n", ipv6Id, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &ipv6Id, CoreIpv6SweepWaitCondition, time.Duration(3*time.Minute),
				CoreIpv6SweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreIpv6Ids(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "Ipv6Id")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listIpv6sRequest := oci_core.ListIpv6sRequest{}
	listIpv6sResponse, err := virtualNetworkClient.ListIpv6s(context.Background(), listIpv6sRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Ipv6 list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ipv6 := range listIpv6sResponse.Items {
		id := *ipv6.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "Ipv6Id", id)
	}
	return resourceIds, nil
}

func CoreIpv6SweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ipv6Response, ok := response.Response.(oci_core.GetIpv6Response); ok {
		return ipv6Response.LifecycleState != oci_core.Ipv6LifecycleStateTerminated
	}
	return false
}

func CoreIpv6SweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetIpv6(context.Background(), oci_core.GetIpv6Request{
		Ipv6Id: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
