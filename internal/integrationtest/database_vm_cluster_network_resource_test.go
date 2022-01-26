// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	oci_database "github.com/oracle/oci-go-sdk/v56/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	VmClusterNetworkValidatedResourceConfig = VmClusterNetworkValidateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, vmClusterNetworkValidateRepresentation)

	vmClusterNetworkValidateRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `testVmClusterNw`},
		"exadata_infrastructure_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"scans":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: vmClusterNetworkScansRepresentation},
		"vm_networks":                 []acctest.RepresentationGroup{{RepType: acctest.Required, Group: vmClusterNetworkBackupVmNetworkRepresentation}, {RepType: acctest.Required, Group: vmClusterNetworkClientVmNetworkRepresentation}},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.10`}, Update: []string{`192.168.10.12`}},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ntp":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.20`}, Update: []string{`192.168.10.22`}},
		"validate_vm_cluster_network": acctest.Representation{RepType: acctest.Optional, Create: "true", Update: "true"},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: vmClusterNetworkIgnoreChangesRepresentation},
	}
	vmClusterNetworkIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`validate_vm_cluster_network`}},
	}
	vmClusterNetworkValidateUpdateRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `testVmClusterNw`},
		"exadata_infrastructure_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"scans":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: vmClusterNetworkScansRepresentation},
		"vm_networks":                 []acctest.RepresentationGroup{{RepType: acctest.Required, Group: vmClusterNetworkBackupVmNetworkRepresentation}, {RepType: acctest.Required, Group: vmClusterNetworkClientVmNetworkRepresentation}},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.10`}, Update: []string{`192.168.10.12`}},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ntp":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.20`}, Update: []string{`192.168.10.22`}},
		"validate_vm_cluster_network": acctest.Representation{RepType: acctest.Optional, Create: "true", Update: "true"},
	}

	VmClusterNetworkValidateResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
				"activation_file":    acctest.Representation{RepType: acctest.Optional, Update: activationFilePath},
				"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
			}))
)

// issue-routing-tag: database/ExaCC
func TestResourceDatabaseVmClusterNetwork_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseVmClusterNetwork_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_vm_cluster_network.test_vm_cluster_network"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with validation
		{
			Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, vmClusterNetworkValidateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(resourceName, "dns.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
					"hostname": "myprefix2-ivmmj-scan",
					"ips.#":    "3",
					"port":     "1522",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
					"domain_name":  "oracle.com",
					"gateway":      "192.169.20.2",
					"netmask":      "255.255.192.0",
					"network_type": "BACKUP",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),
				resource.TestCheckResourceAttr(resourceName, "state", "VALIDATED"),
			),
		},
		//  delete before next Create
		{
			Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies,
		},
		// verify Create without validation
		{
			Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithRemovedProperties(vmClusterNetworkValidateRepresentation, []string{`validate_vm_cluster_network`})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(resourceName, "dns.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
					"hostname": "myprefix1-ivmmj-scan",
					"ips.#":    "3",
					"port":     "1521",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
					"domain_name":  "oracle.com",
					"gateway":      "192.169.20.1",
					"netmask":      "255.255.0.0",
					"network_type": "BACKUP",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),
				resource.TestCheckResourceAttr(resourceName, "state", "REQUIRES_VALIDATION"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify validation
		{
			Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, vmClusterNetworkValidateUpdateRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(resourceName, "dns.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
					"hostname": "myprefix2-ivmmj-scan",
					"ips.#":    "3",
					"port":     "1522",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
					"domain_name":  "oracle.com",
					"gateway":      "192.169.20.2",
					"netmask":      "255.255.192.0",
					"network_type": "BACKUP",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),
				resource.TestCheckResourceAttr(resourceName, "state", "VALIDATED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify Update after validation
		{
			Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Create, vmClusterNetworkRepresentation),
			ExpectError: regexp.MustCompile("Update not allowed on validated vm cluster network"),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseValidatedVmClusterNetwork") {
		resource.AddTestSweepers("DatabaseValidatedVmClusterNetwork", &resource.Sweeper{
			Name:         "DatabaseValidatedVmClusterNetwork",
			Dependencies: acctest.DependencyGraph["vmClusterNetwork"],
			F:            sweepDatabaseValidatedVmClusterNetworkResource,
		})
	}
}

func sweepDatabaseValidatedVmClusterNetworkResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	vmClusterNetworkIds, err := getVmClusterNetworkIds(compartment)
	if err != nil {
		return err
	}
	for _, vmClusterNetworkId := range vmClusterNetworkIds {
		if ok := acctest.SweeperDefaultResourceId[vmClusterNetworkId]; !ok {
			deleteVmClusterNetworkRequest := oci_database.DeleteVmClusterNetworkRequest{}

			deleteVmClusterNetworkRequest.VmClusterNetworkId = &vmClusterNetworkId

			deleteVmClusterNetworkRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteVmClusterNetwork(context.Background(), deleteVmClusterNetworkRequest)
			if error != nil {
				fmt.Printf("Error deleting VmClusterNetwork %s %s, It is possible that the resource is already deleted. Please verify manually \n", vmClusterNetworkId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vmClusterNetworkId, vmClusterNetworkSweepWaitCondition, time.Duration(3*time.Minute),
				vmClusterNetworkSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getValidatedVmClusterNetworkIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VmClusterNetworkId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listVmClusterNetworksRequest := oci_database.ListVmClusterNetworksRequest{}
	listVmClusterNetworksRequest.CompartmentId = &compartmentId

	exadataInfrastructureIds, error := getExadataInfrastructureIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting exadataInfrastructureId required for VmClusterNetwork resource requests \n")
	}
	for _, exadataInfrastructureId := range exadataInfrastructureIds {
		listVmClusterNetworksRequest.ExadataInfrastructureId = &exadataInfrastructureId

		listVmClusterNetworksRequest.LifecycleState = oci_database.VmClusterNetworkSummaryLifecycleStateValidated
		listVmClusterNetworksResponse, err := databaseClient.ListVmClusterNetworks(context.Background(), listVmClusterNetworksRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting VmClusterNetwork list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, vmClusterNetwork := range listVmClusterNetworksResponse.Items {
			id := *vmClusterNetwork.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VmClusterNetworkId", id)
		}

	}
	return resourceIds, nil
}
