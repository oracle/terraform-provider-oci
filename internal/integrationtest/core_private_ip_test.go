// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

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
	CorePrivateIpRequiredOnlyResource = CorePrivateIpResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", acctest.Required, acctest.Create, CorePrivateIpRepresentation)

	CorePrivateIpResourceConfig = CorePrivateIpResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", acctest.Optional, acctest.Update, CorePrivateIpRepresentation)

	privateIpSingularDataSourceRepresentation = map[string]interface{}{
		"private_ip_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_private_ip.test_private_ip.id}`},
	}

	CorePrivateIpDataSourceRepresentation = map[string]interface{}{
		"ip_address": acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.5`},
		"ip_state":   acctest.Representation{RepType: acctest.Optional, Create: `ipState`},
		"lifetime":   acctest.Representation{RepType: acctest.Optional, Create: `EPHEMERAL`, Update: `RESERVED`},
		"subnet_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"vnic_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vnic_attachment.test_vnic_attachment.id}`},
		"filter":     acctest.RepresentationGroup{RepType: acctest.Required, Group: CorePrivateIpDataSourceFilterRepresentation}}
	CoreCorePrivateIpDataSourceRepresentation = map[string]interface{}{
		"vnic_id": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
		"filter":  acctest.RepresentationGroup{RepType: acctest.Required, Group: CorePrivateIpDataSourceFilterRepresentation}}
	CorePrivateIpDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_private_ip.test_private_ip.id}`}},
	}

	CorePrivateIpRepresentation = map[string]interface{}{
		"vnic_id":        acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label": acctest.Representation{RepType: acctest.Optional, Create: `privateiptestinstance`, Update: `privateiptestinstance2`},
		"ip_address":     acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.5`},
		"lifetime":       acctest.Representation{RepType: acctest.Optional, Create: `EPHEMERAL`, Update: `RESERVED`},
		"route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.test_route_table.id}`},
	}

	CorePrivateIpRepresentation2 = map[string]interface{}{
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label": acctest.Representation{RepType: acctest.Optional, Create: `privateiptestinstance`, Update: `privateiptestinstance2`},
		"ip_address":     acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.5`},
		"lifetime":       acctest.Representation{RepType: acctest.Optional, Create: `EPHEMERAL`, Update: `RESERVED`},
		"route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.test_route_table.id}`},
	}

	CorePrivateIpResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
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
func TestCorePrivateIpResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePrivateIpResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_private_ip.test_private_ip"
	datasourceName := "data.oci_core_private_ips.test_private_ips"
	singularDatasourceName := "data.oci_core_private_ip.test_private_ip"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CorePrivateIpResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", acctest.Optional, acctest.Create, CorePrivateIpRepresentation), "core", "privateIp", t)

	acctest.ResourceTest(t, testAccCheckCorePrivateIpDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CorePrivateIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", acctest.Required, acctest.Create, CorePrivateIpRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CorePrivateIpResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CorePrivateIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", acctest.Optional, acctest.Create, CorePrivateIpRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "privateiptestinstance"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", "EPHEMERAL"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
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
			Config: config + compartmentIdVariableStr + CorePrivateIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", acctest.Optional, acctest.Update, CorePrivateIpRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "privateiptestinstance2"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", "RESERVED"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
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

		// test detach
		{
			Config: config + compartmentIdVariableStr + CorePrivateIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", acctest.Optional, acctest.Update, CorePrivateIpRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "privateiptestinstance2"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", "RESERVED"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_private_ips", "test_private_ips", acctest.Optional, acctest.Update, CoreCorePrivateIpDataSourceRepresentation) +
				compartmentIdVariableStr + CorePrivateIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", acctest.Optional, acctest.Update, CorePrivateIpRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(datasourceName, "ip_address", "ipAddress"),
				//resource.TestCheckResourceAttr(datasourceName, "ip_state", "ipState"),
				//resource.TestCheckResourceAttr(datasourceName, "lifetime", "lifetime2"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subnet_id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "vnic_id"),

				resource.TestCheckResourceAttr(datasourceName, "private_ips.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "private_ips.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "private_ips.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "private_ips.0.hostname_label", "privateiptestinstance2"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "private_ips.0.ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.ip_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.is_primary"),
				resource.TestCheckResourceAttr(datasourceName, "private_ips.0.lifetime", "RESERVED"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.route_table_id"),
				//commenting until service issue resolved
				//resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.is_reserved"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.vnic_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", acctest.Required, acctest.Create, privateIpSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CorePrivateIpResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hostname_label", "privateiptestinstance2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lifetime", "RESERVED"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "is_reserved"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "route_table_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + CorePrivateIpRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCorePrivateIpDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_private_ip" {
			noResourceFound = false
			request := oci_core.GetPrivateIpRequest{}

			tmp := rs.Primary.ID
			request.PrivateIpId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			_, err := client.GetPrivateIp(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("CorePrivateIp") {
		resource.AddTestSweepers("CorePrivateIp", &resource.Sweeper{
			Name:         "CorePrivateIp",
			Dependencies: acctest.DependencyGraph["privateIp"],
			F:            sweepCorePrivateIpResource,
		})
	}
}

func sweepCorePrivateIpResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	privateIpIds, err := getCorePrivateIpIds(compartment)
	if err != nil {
		return err
	}
	for _, privateIpId := range privateIpIds {
		if ok := acctest.SweeperDefaultResourceId[privateIpId]; !ok {
			deletePrivateIpRequest := oci_core.DeletePrivateIpRequest{}

			deletePrivateIpRequest.PrivateIpId = &privateIpId

			deletePrivateIpRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeletePrivateIp(context.Background(), deletePrivateIpRequest)
			if error != nil {
				fmt.Printf("Error deleting PrivateIp %s %s, It is possible that the resource is already deleted. Please verify manually \n", privateIpId, error)
				continue
			}
		}
	}
	return nil
}

func getCorePrivateIpIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PrivateIpId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listPrivateIpsRequest := oci_core.ListPrivateIpsRequest{}

	subnetIds, err := getCoreSubnetIds(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SubnetId list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subnetId := range subnetIds {
		listPrivateIpsRequest.SubnetId = &subnetId
		listPrivateIpsResponse, err := virtualNetworkClient.ListPrivateIps(context.Background(), listPrivateIpsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting PrivateIp list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, privateIp := range listPrivateIpsResponse.Items {
			id := *privateIp.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PrivateIpId", id)
		}
	}
	return resourceIds, nil
}
