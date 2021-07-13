// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v44/common"
	oci_core "github.com/oracle/oci-go-sdk/v44/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PrivateIpRequiredOnlyResource = PrivateIpResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Required, Create, privateIpRepresentation)

	PrivateIpResourceConfig = PrivateIpResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Optional, Update, privateIpRepresentation)

	privateIpSingularDataSourceRepresentation = map[string]interface{}{
		"private_ip_id": Representation{repType: Required, create: `${oci_core_private_ip.test_private_ip.id}`},
	}

	privateIpDataSourceRepresentation = map[string]interface{}{
		"vnic_id": Representation{repType: Optional, create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
		"filter":  RepresentationGroup{Required, privateIpDataSourceFilterRepresentation}}
	privateIpDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_private_ip.test_private_ip.id}`}},
	}

	privateIpRepresentation = map[string]interface{}{
		"vnic_id":        Representation{repType: Required, create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"hostname_label": Representation{repType: Optional, create: `privateiptestinstance`, update: `privateiptestinstance2`},
		"ip_address":     Representation{repType: Optional, create: `10.0.0.5`},
	}

	PrivateIpResourceDependencies = OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
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

func TestCorePrivateIpResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePrivateIpResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_private_ip.test_private_ip"
	datasourceName := "data.oci_core_private_ips.test_private_ips"
	singularDatasourceName := "data.oci_core_private_ip.test_private_ip"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+PrivateIpResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Optional, Create, privateIpRepresentation), "core", "privateIp", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCorePrivateIpDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + PrivateIpResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Required, Create, privateIpRepresentation),
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
				Config: config + compartmentIdVariableStr + PrivateIpResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + PrivateIpResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Optional, Create, privateIpRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "privateiptestinstance"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
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
				Config: config + compartmentIdVariableStr + PrivateIpResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Optional, Update, privateIpRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "privateiptestinstance2"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
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
					generateDataSourceFromRepresentationMap("oci_core_private_ips", "test_private_ips", Optional, Update, privateIpDataSourceRepresentation) +
					compartmentIdVariableStr + PrivateIpResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Optional, Update, privateIpRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_id"),

					resource.TestCheckResourceAttr(datasourceName, "private_ips.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.hostname_label", "privateiptestinstance2"),
					resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.ip_address", "10.0.0.5"),
					resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.is_primary"),
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
					generateDataSourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Required, Create, privateIpSingularDataSourceRepresentation) +
					compartmentIdVariableStr + PrivateIpResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "hostname_label", "privateiptestinstance2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ip_address", "10.0.0.5"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "is_reserved"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + PrivateIpResourceConfig,
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

func testAccCheckCorePrivateIpDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_private_ip" {
			noResourceFound = false
			request := oci_core.GetPrivateIpRequest{}

			tmp := rs.Primary.ID
			request.PrivateIpId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CorePrivateIp") {
		resource.AddTestSweepers("CorePrivateIp", &resource.Sweeper{
			Name:         "CorePrivateIp",
			Dependencies: DependencyGraph["privateIp"],
			F:            sweepCorePrivateIpResource,
		})
	}
}

func sweepCorePrivateIpResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	privateIpIds, err := getPrivateIpIds(compartment)
	if err != nil {
		return err
	}
	for _, privateIpId := range privateIpIds {
		if ok := SweeperDefaultResourceId[privateIpId]; !ok {
			deletePrivateIpRequest := oci_core.DeletePrivateIpRequest{}

			deletePrivateIpRequest.PrivateIpId = &privateIpId

			deletePrivateIpRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeletePrivateIp(context.Background(), deletePrivateIpRequest)
			if error != nil {
				fmt.Printf("Error deleting PrivateIp %s %s, It is possible that the resource is already deleted. Please verify manually \n", privateIpId, error)
				continue
			}
		}
	}
	return nil
}

func getPrivateIpIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "PrivateIpId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listPrivateIpsRequest := oci_core.ListPrivateIpsRequest{}

	subnetIds, err := getSubnetIds(compartment)
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
			addResourceIdToSweeperResourceIdMap(compartmentId, "PrivateIpId", id)
		}
	}
	return resourceIds, nil
}
