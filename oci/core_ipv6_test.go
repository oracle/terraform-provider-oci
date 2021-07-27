// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v45/common"
	oci_core "github.com/oracle/oci-go-sdk/v45/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	Ipv6RequiredOnlyResource = Ipv6ResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", Required, Create, ipv6Representation)

	Ipv6ResourceConfig = Ipv6ResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", Optional, Update, ipv6Representation)

	ipv6SingularDataSourceRepresentation = map[string]interface{}{
		"ipv6id": Representation{repType: Required, create: `${oci_core_ipv6.test_ipv6.id}`},
	}

	ipv6DataSourceRepresentation = map[string]interface{}{
		"vnic_id": Representation{repType: Optional, create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
		"filter":  RepresentationGroup{Required, ipv6DataSourceFilterRepresentation}}
	ipv6DataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_ipv6.test_ipv6.id}`}},
	}

	ipv6Representation = map[string]interface{}{
		"vnic_id":       Representation{repType: Required, create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
		"defined_tags":  Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags": Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"ip_address":    Representation{repType: Optional, create: `${substr(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn.ipv6cidr_blocks[0]) - 4)}5901:cede:a617:8bba`},
	}

	Ipv6ResourceDependencies = OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Optional, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label":      Representation{repType: Required, create: `dnslabel`},
			"ipv6cidr_block": Representation{repType: Optional, create: `${substr(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn.ipv6cidr_blocks[0]) - 2)}${64}`},
		})) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label":      Representation{repType: Required, create: `dnslabel`},
			"is_ipv6enabled": Representation{repType: Optional, create: `true`},
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

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_ipv6.test_ipv6"
	datasourceName := "data.oci_core_ipv6s.test_ipv6s"
	singularDatasourceName := "data.oci_core_ipv6.test_ipv6"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+Ipv6ResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", Optional, Create, ipv6Representation), "core", "ipv6", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreIpv6Destroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + Ipv6ResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", Required, Create, ipv6Representation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + Ipv6ResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + Ipv6ResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", Optional, Create, ipv6Representation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

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
				Config: config + compartmentIdVariableStr + Ipv6ResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", Optional, Update, ipv6Representation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_ipv6s", "test_ipv6s", Optional, Update, ipv6DataSourceRepresentation) +
					compartmentIdVariableStr + Ipv6ResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", Optional, Update, ipv6Representation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_id"),

					resource.TestCheckResourceAttr(datasourceName, "ipv6s.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "ipv6s.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "ipv6s.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "ipv6s.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.ip_address"),
					resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "ipv6s.0.vnic_id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_ipv6", "test_ipv6", Required, Create, ipv6SingularDataSourceRepresentation) +
					compartmentIdVariableStr + Ipv6ResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ipv6id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + Ipv6ResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreIpv6Destroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_ipv6" {
			noResourceFound = false
			request := oci_core.GetIpv6Request{}

			tmp := rs.Primary.ID
			request.Ipv6Id = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreIpv6") {
		resource.AddTestSweepers("CoreIpv6", &resource.Sweeper{
			Name:         "CoreIpv6",
			Dependencies: DependencyGraph["ipv6"],
			F:            sweepCoreIpv6Resource,
		})
	}
}

func sweepCoreIpv6Resource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	ipv6Ids, err := getIpv6Ids(compartment)
	if err != nil {
		return err
	}
	for _, ipv6Id := range ipv6Ids {
		if ok := SweeperDefaultResourceId[ipv6Id]; !ok {
			deleteIpv6Request := oci_core.DeleteIpv6Request{}

			deleteIpv6Request.Ipv6Id = &ipv6Id

			deleteIpv6Request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteIpv6(context.Background(), deleteIpv6Request)
			if error != nil {
				fmt.Printf("Error deleting Ipv6 %s %s, It is possible that the resource is already deleted. Please verify manually \n", ipv6Id, error)
				continue
			}
			waitTillCondition(testAccProvider, &ipv6Id, ipv6SweepWaitCondition, time.Duration(3*time.Minute),
				ipv6SweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getIpv6Ids(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "Ipv6Id")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listIpv6sRequest := oci_core.ListIpv6sRequest{}
	listIpv6sResponse, err := virtualNetworkClient.ListIpv6s(context.Background(), listIpv6sRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Ipv6 list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ipv6 := range listIpv6sResponse.Items {
		id := *ipv6.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "Ipv6Id", id)
	}
	return resourceIds, nil
}

func ipv6SweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ipv6Response, ok := response.Response.(oci_core.GetIpv6Response); ok {
		return ipv6Response.LifecycleState != oci_core.Ipv6LifecycleStateTerminated
	}
	return false
}

func ipv6SweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetIpv6(context.Background(), oci_core.GetIpv6Request{
		Ipv6Id: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
