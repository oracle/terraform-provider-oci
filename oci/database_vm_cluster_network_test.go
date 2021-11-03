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
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_database "github.com/oracle/oci-go-sdk/v54/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VmClusterNetworkRequiredOnlyResource = VmClusterNetworkResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Required, Create, vmClusterNetworkRepresentation)

	VmClusterNetworkResourceConfig = VmClusterNetworkResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkRepresentation)

	vmClusterNetworkSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"vm_cluster_network_id":     Representation{RepType: Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
	}

	vmClusterNetworkDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"exadata_infrastructure_id": Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"display_name":              Representation{RepType: Optional, Create: `testVmClusterNw`},
		"state":                     Representation{RepType: Optional, Create: `REQUIRES_VALIDATION`},
		"filter":                    RepresentationGroup{Required, vmClusterNetworkDataSourceFilterRepresentation}}
	vmClusterNetworkDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_database_vm_cluster_network.test_vm_cluster_network.id}`}},
	}

	vmClusterNetworkRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":              Representation{RepType: Required, Create: `testVmClusterNw`},
		"exadata_infrastructure_id": Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"scans":                     RepresentationGroup{Required, vmClusterNetworkScansRepresentation},
		"vm_networks":               []RepresentationGroup{{Required, vmClusterNetworkBackupVmNetworkRepresentation}, {Required, vmClusterNetworkClientVmNetworkRepresentation}},
		"defined_tags":              Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                       Representation{RepType: Optional, Create: []string{`192.168.10.10`}, Update: []string{`192.168.10.12`}},
		"freeform_tags":             Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ntp":                       Representation{RepType: Optional, Create: []string{`192.168.10.20`}, Update: []string{`192.168.10.22`}},
	}
	vmClusterNetworkScansRepresentation = map[string]interface{}{
		"hostname":                   Representation{RepType: Required, Create: `myprefix1-ivmmj-scan`, Update: `myprefix2-ivmmj-scan`},
		"ips":                        Representation{RepType: Required, Create: []string{`192.168.19.7`, `192.168.19.6`, `192.168.19.8`}, Update: []string{`192.168.19.7`, `192.168.19.8`, `192.168.19.9`}},
		"port":                       Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"scan_listener_port_tcp":     Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"scan_listener_port_tcp_ssl": Representation{RepType: Optional, Create: `2484`, Update: `2484`},
	}
	vmClusterNetworkBackupVmNetworkRepresentation = map[string]interface{}{
		"domain_name":  Representation{RepType: Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      Representation{RepType: Required, Create: `192.169.20.1`, Update: `192.169.20.2`},
		"netmask":      Representation{RepType: Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": Representation{RepType: Required, Create: `BACKUP`, Update: `BACKUP`},
		"nodes":        []RepresentationGroup{{Required, vmClusterNetworkVmNetworksBackupNodes1Representation}, {Required, vmClusterNetworkVmNetworksBackupNodes2Representation}},
		"vlan_id":      Representation{RepType: Required, Create: `100`},
	}
	vmClusterNetworkClientVmNetworkRepresentation = map[string]interface{}{
		"domain_name":  Representation{RepType: Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      Representation{RepType: Required, Create: `192.168.20.1`, Update: `192.168.20.2`},
		"netmask":      Representation{RepType: Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": Representation{RepType: Required, Create: `CLIENT`, Update: `CLIENT`},
		"nodes":        []RepresentationGroup{{Required, vmClusterNetworkVmNetworksClientNodes1Representation}, {Required, vmClusterNetworkVmNetworksClientNodes2Representation}},
		"vlan_id":      Representation{RepType: Required, Create: `101`},
	}
	vmClusterNetworkVmNetworksClientNodes1Representation = map[string]interface{}{
		"hostname":     Representation{RepType: Required, Create: `myprefix2-xapb21`, Update: `myprefix2-xapb22`},
		"ip":           Representation{RepType: Required, Create: `192.168.19.10`, Update: `192.168.19.11`},
		"vip":          Representation{RepType: Optional, Create: `192.168.19.12`, Update: `192.168.19.13`},
		"vip_hostname": Representation{RepType: Optional, Create: `myprefix2-xapb21-vip`, Update: `myprefix2-xapb22-vip`},
	}
	vmClusterNetworkVmNetworksClientNodes2Representation = map[string]interface{}{
		"hostname":     Representation{RepType: Required, Create: `myprefix2-xapb25`, Update: `myprefix2-xapb26`},
		"ip":           Representation{RepType: Required, Create: `192.168.19.14`, Update: `192.168.19.15`},
		"vip":          Representation{RepType: Optional, Create: `192.168.19.16`, Update: `192.168.19.17`},
		"vip_hostname": Representation{RepType: Optional, Create: `myprefix2-xapb25-vip`, Update: `myprefix2-xapb26-vip`},
	}
	vmClusterNetworkVmNetworksBackupNodes1Representation = map[string]interface{}{
		"hostname": Representation{RepType: Required, Create: `myprefix2-xapb23`, Update: `myprefix2-xapb24`},
		"ip":       Representation{RepType: Required, Create: `192.169.19.18`, Update: `192.169.19.19`},
	}
	vmClusterNetworkVmNetworksBackupNodes2Representation = map[string]interface{}{
		"hostname": Representation{RepType: Required, Create: `myprefix2-xapb27`, Update: `myprefix2-xapb28`},
		"ip":       Representation{RepType: Required, Create: `192.169.19.20`, Update: `192.169.19.21`},
	}

	activationFilePath, _ = createTmpActivationFile()

	VmClusterNetworkResourceDependencies = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
			RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
				"activation_file":    Representation{RepType: Optional, Update: activationFilePath},
				"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
			}))
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterNetworkResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterNetworkResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_vm_cluster_network.test_vm_cluster_network"
	datasourceName := "data.oci_database_vm_cluster_networks.test_vm_cluster_networks"
	singularDatasourceName := "data.oci_database_vm_cluster_network.test_vm_cluster_network"

	var resId, resId2, compositeId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+VmClusterNetworkResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Create, vmClusterNetworkRepresentation), "database", "vmClusterNetwork", t)

	ResourceTest(t, testAccCheckDatabaseVmClusterNetworkDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Required, Create, vmClusterNetworkRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
					"hostname": "myprefix1-ivmmj-scan",
					"ips.#":    "3",
					"port":     "1521",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
				CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
					"domain_name":  "oracle.com",
					"gateway":      "192.168.20.1",
					"netmask":      "255.255.0.0",
					"network_type": "CLIENT",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Create, vmClusterNetworkRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(resourceName, "dns.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
					"hostname":                   "myprefix1-ivmmj-scan",
					"ips.#":                      "3",
					"port":                       "1521",
					"scan_listener_port_tcp":     "1521",
					"scan_listener_port_tcp_ssl": "2484",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
				CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
					"domain_name":  "oracle.com",
					"gateway":      "192.168.20.1",
					"netmask":      "255.255.0.0",
					"network_type": "CLIENT",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					exadataInfrastructureId, _ := FromInstanceState(s, resourceName, "exadata_infrastructure_id")
					compositeId = "exadataInfrastructures/" + exadataInfrastructureId + "/vmClusterNetworks/" + resId
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(resourceName, "dns.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
					"hostname":                   "myprefix2-ivmmj-scan",
					"ips.#":                      "3",
					"port":                       "1522",
					"scan_listener_port_tcp":     "1522",
					"scan_listener_port_tcp_ssl": "2484",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
				CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
					"domain_name":  "oracle.com",
					"gateway":      "192.169.20.2",
					"netmask":      "255.255.192.0",
					"network_type": "BACKUP",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_networks", "test_vm_cluster_networks", Optional, Update, vmClusterNetworkDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "REQUIRES_VALIDATION"),

				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.dns.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_networks.0.exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_networks.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.ntp.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.scans.#", "1"),
				CheckResourceSetContainsElementWithProperties(datasourceName, "vm_cluster_networks.0.scans", map[string]string{
					"hostname":                   "myprefix2-ivmmj-scan",
					"ips.#":                      "3",
					"port":                       "1522",
					"scan_listener_port_tcp":     "1522",
					"scan_listener_port_tcp_ssl": "2484",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_networks.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_networks.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.vm_networks.#", "2"),
				CheckResourceSetContainsElementWithProperties(datasourceName, "vm_cluster_networks.0.vm_networks", map[string]string{
					"domain_name":  "oracle.com",
					"gateway":      "192.169.20.2",
					"netmask":      "255.255.192.0",
					"network_type": "BACKUP",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Required, Create, vmClusterNetworkSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterNetworkResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_network_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scans.#", "1"),
				CheckResourceSetContainsElementWithProperties(singularDatasourceName, "scans", map[string]string{
					"hostname":                   "myprefix2-ivmmj-scan",
					"ips.#":                      "3",
					"port":                       "1522",
					"scan_listener_port_tcp":     "1522",
					"scan_listener_port_tcp_ssl": "2484",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.#", "2"),
				CheckResourceSetContainsElementWithProperties(singularDatasourceName, "vm_networks", map[string]string{
					"domain_name":  "oracle.com",
					"gateway":      "192.169.20.2",
					"netmask":      "255.255.192.0",
					"network_type": "BACKUP",
					"nodes.#":      "2",
				},
					[]string{}),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + VmClusterNetworkResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateIdFunc: getVmClusterNetworkImportId(resourceName),
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"validate_vm_cluster_network",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseVmClusterNetworkDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_vm_cluster_network" {
			noResourceFound = false
			request := oci_database.GetVmClusterNetworkRequest{}

			if value, ok := rs.Primary.Attributes["exadata_infrastructure_id"]; ok {
				request.ExadataInfrastructureId = &value
			}

			tmp := rs.Primary.ID
			request.VmClusterNetworkId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")

			response, err := client.GetVmClusterNetwork(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.VmClusterNetworkLifecycleStateTerminated): true,
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
		InitDependencyGraph()
	}
	if !InSweeperExcludeList("DatabaseVmClusterNetwork") {
		resource.AddTestSweepers("DatabaseVmClusterNetwork", &resource.Sweeper{
			Name:         "DatabaseVmClusterNetwork",
			Dependencies: DependencyGraph["vmClusterNetwork"],
			F:            sweepDatabaseVmClusterNetworkResource,
		})
	}
}

func sweepDatabaseVmClusterNetworkResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	vmClusterNetworkIds, err := getVmClusterNetworkIds(compartment)
	if err != nil {
		return err
	}
	for _, vmClusterNetworkId := range vmClusterNetworkIds {
		if ok := SweeperDefaultResourceId[vmClusterNetworkId]; !ok {
			deleteVmClusterNetworkRequest := oci_database.DeleteVmClusterNetworkRequest{}

			deleteVmClusterNetworkRequest.VmClusterNetworkId = &vmClusterNetworkId

			deleteVmClusterNetworkRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteVmClusterNetwork(context.Background(), deleteVmClusterNetworkRequest)
			if error != nil {
				fmt.Printf("Error deleting VmClusterNetwork %s %s, It is possible that the resource is already deleted. Please verify manually \n", vmClusterNetworkId, error)
				continue
			}
			WaitTillCondition(TestAccProvider, &vmClusterNetworkId, vmClusterNetworkSweepWaitCondition, time.Duration(3*time.Minute),
				vmClusterNetworkSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getVmClusterNetworkIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "VmClusterNetworkId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listVmClusterNetworksRequest := oci_database.ListVmClusterNetworksRequest{}
	listVmClusterNetworksRequest.CompartmentId = &compartmentId

	exadataInfrastructureIds, error := getExadataInfrastructureIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting exadataInfrastructureId required for VmClusterNetwork resource requests \n")
	}
	for _, exadataInfrastructureId := range exadataInfrastructureIds {
		listVmClusterNetworksRequest.ExadataInfrastructureId = &exadataInfrastructureId

		listVmClusterNetworksRequest.LifecycleState = oci_database.VmClusterNetworkSummaryLifecycleStateRequiresValidation
		listVmClusterNetworksResponse, err := databaseClient.ListVmClusterNetworks(context.Background(), listVmClusterNetworksRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting VmClusterNetwork list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, vmClusterNetwork := range listVmClusterNetworksResponse.Items {
			id := *vmClusterNetwork.Id
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "VmClusterNetworkId", id)
		}

	}
	return resourceIds, nil
}

func vmClusterNetworkSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vmClusterNetworkResponse, ok := response.Response.(oci_database.GetVmClusterNetworkResponse); ok {
		return vmClusterNetworkResponse.LifecycleState != oci_database.VmClusterNetworkLifecycleStateTerminated
	}
	return false
}

func vmClusterNetworkSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetVmClusterNetwork(context.Background(), oci_database.GetVmClusterNetworkRequest{
		VmClusterNetworkId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func getVmClusterNetworkImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("exadataInfrastructures/" + rs.Primary.Attributes["exadata_infrastructure_id"] + "/vmClusterNetworks/" + rs.Primary.Attributes["id"]), nil
	}
}
