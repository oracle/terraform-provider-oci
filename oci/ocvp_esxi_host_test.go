// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v31/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v31/ocvp"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	EsxiHostRequiredOnlyResource = EsxiHostResourceDependencies +
		generateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", Required, Create, esxiHostRepresentation)

	EsxiHostResourceConfig = EsxiHostResourceDependencies +
		generateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", Optional, Update, esxiHostRepresentation)

	esxiHostSingularDataSourceRepresentation = map[string]interface{}{
		"esxi_host_id": Representation{repType: Required, create: `${oci_ocvp_esxi_host.test_esxi_host.id}`},
	}
	esxiHostDataSourceRepresentation = map[string]interface{}{
		"compute_instance_id": Representation{repType: Optional, create: `${oci_ocvp_esxi_host.test_esxi_host.compute_instance_id}`},
		"display_name":        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"sddc_id":             Representation{repType: Optional, create: `${oci_ocvp_sddc.test_sddc.id}`},
		"state":               Representation{repType: Optional, create: `ACTIVE`},
		"filter":              RepresentationGroup{Required, esxiHostDataSourceFilterRepresentation},
	}
	esxiHostDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_ocvp_esxi_host.test_esxi_host.id}`}},
	}

	esxiHostRepresentation = map[string]interface{}{
		"sddc_id":       Representation{repType: Required, create: `${oci_ocvp_sddc.test_sddc.id}`},
		"defined_tags":  Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags": Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	EsxiHostResourceDependencies = SddcResourceDependencies + `

resource "oci_ocvp_sddc" "test_sddc" {
    compartment_id = "${var.compartment_id}"
    compute_availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    esxi_hosts_count = "3"
    nsx_edge_uplink1vlan_id = "${oci_core_vlan.test_nsx_edge_uplink1_vlan.id}"
    nsx_edge_uplink2vlan_id = "${oci_core_vlan.test_nsx_edge_uplink2_vlan.id}"
    nsx_edge_vtep_vlan_id = "${oci_core_vlan.test_nsx_edge_vtep_vlan.id}"
    nsx_vtep_vlan_id = "${oci_core_vlan.test_nsx_vtep_vlan.id}"
    provisioning_subnet_id = "${oci_core_subnet.test_provisioning_subnet.id}"
    ssh_authorized_keys = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
    vmotion_vlan_id = "${oci_core_vlan.test_vmotion_net_vlan.id}"
    vmware_software_version = "${lookup(data.oci_ocvp_supported_vmware_software_versions.test_supported_vmware_software_versions.items[1], "version")}"
    vsan_vlan_id = "${oci_core_vlan.test_vsan_net_vlan.id}"
    vsphere_vlan_id = "${oci_core_vlan.test_vsphere_net_vlan.id}"
}

`
)

func TestOcvpEsxiHostResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpEsxiHostResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_ocvp_esxi_host.test_esxi_host"
	datasourceName := "data.oci_ocvp_esxi_hosts.test_esxi_hosts"
	singularDatasourceName := "data.oci_ocvp_esxi_host.test_esxi_host"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOcvpEsxiHostDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies +
					generateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", Required, Create, esxiHostRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies +
					generateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", Optional, Create, esxiHostRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),

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
				Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies +
					generateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", Optional, Update, esxiHostRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),

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
					generateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "test_esxi_hosts", Optional, Update, esxiHostDataSourceRepresentation) +
					compartmentIdVariableStr + EsxiHostResourceDependencies +
					generateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", Optional, Update, esxiHostRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "sddc_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.sddc_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.compute_instance_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.time_updated"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", Required, Create, esxiHostSingularDataSourceRepresentation) +
					compartmentIdVariableStr + EsxiHostResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "esxi_host_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + EsxiHostResourceConfig,
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

func testAccCheckOcvpEsxiHostDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).esxiHostClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_esxi_host" {
			noResourceFound = false
			request := oci_ocvp.GetEsxiHostRequest{}

			tmp := rs.Primary.ID
			request.EsxiHostId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "ocvp")

			response, err := client.GetEsxiHost(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ocvp.LifecycleStatesDeleted): true,
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
	if !inSweeperExcludeList("OcvpEsxiHost") {
		resource.AddTestSweepers("OcvpEsxiHost", &resource.Sweeper{
			Name:         "OcvpEsxiHost",
			Dependencies: DependencyGraph["esxiHost"],
			F:            sweepOcvpEsxiHostResource,
		})
	}
}

func sweepOcvpEsxiHostResource(compartment string) error {
	esxiHostClient := GetTestClients(&schema.ResourceData{}).esxiHostClient()
	esxiHostIds, err := getEsxiHostIds(compartment)
	if err != nil {
		return err
	}
	for _, esxiHostId := range esxiHostIds {
		if ok := SweeperDefaultResourceId[esxiHostId]; !ok {
			deleteEsxiHostRequest := oci_ocvp.DeleteEsxiHostRequest{}

			deleteEsxiHostRequest.EsxiHostId = &esxiHostId

			deleteEsxiHostRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "ocvp")
			_, error := esxiHostClient.DeleteEsxiHost(context.Background(), deleteEsxiHostRequest)
			if error != nil {
				fmt.Printf("Error deleting EsxiHost %s %s, It is possible that the resource is already deleted. Please verify manually \n", esxiHostId, error)
				continue
			}
			waitTillCondition(testAccProvider, &esxiHostId, esxiHostSweepWaitCondition, time.Duration(3*time.Minute),
				esxiHostSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getEsxiHostIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "EsxiHostId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	esxiHostClient := GetTestClients(&schema.ResourceData{}).esxiHostClient()

	listEsxiHostsRequest := oci_ocvp.ListEsxiHostsRequest{}
	listEsxiHostsRequest.LifecycleState = oci_ocvp.ListEsxiHostsLifecycleStateActive
	listEsxiHostsResponse, err := esxiHostClient.ListEsxiHosts(context.Background(), listEsxiHostsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting EsxiHost list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, esxiHost := range listEsxiHostsResponse.Items {
		id := *esxiHost.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "EsxiHostId", id)
	}
	return resourceIds, nil
}

func esxiHostSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if esxiHostResponse, ok := response.Response.(oci_ocvp.GetEsxiHostResponse); ok {
		return esxiHostResponse.LifecycleState != oci_ocvp.LifecycleStatesDeleted
	}
	return false
}

func esxiHostSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.esxiHostClient().GetEsxiHost(context.Background(), oci_ocvp.GetEsxiHostRequest{
		EsxiHostId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
