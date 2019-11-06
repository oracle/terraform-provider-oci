// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VmClusterNetworkValidatedResourceConfig = VmClusterNetworkValidateResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateRepresentation)

	vmClusterNetworkValidateRepresentation = map[string]interface{}{
		"compartment_id":              Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                Representation{repType: Required, create: `testVmClusterNw`},
		"exadata_infrastructure_id":   Representation{repType: Required, create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"scans":                       RepresentationGroup{Required, vmClusterNetworkScansRepresentation},
		"vm_networks":                 []RepresentationGroup{{Required, vmClusterNetworkBackupVmNetworkRepresentation}, {Required, vmClusterNetworkClientVmNetworkRepresentation}},
		"defined_tags":                Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                         Representation{repType: Optional, create: []string{`192.168.10.10`}, update: []string{`192.168.10.12`}},
		"freeform_tags":               Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"ntp":                         Representation{repType: Optional, create: []string{`192.168.10.20`}, update: []string{`192.168.10.22`}},
		"validate_vm_cluster_network": Representation{repType: Optional, create: "true"},
	}

	VmClusterNetworkValidateResourceDependencies = DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
			representationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": Representation{repType: Optional, update: activationFilePath}}))
)

func TestResourceDatabaseVmClusterNetwork_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseVmClusterNetwork_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_vm_cluster_network.test_vm_cluster_network"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with validation
			{
				Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
					resource.TestCheckResourceAttr(resourceName, "dns.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "ntp.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
						"hostname": "myprefix2-ivmmj-scan",
						"ips.#":    "3",
						"port":     "1522",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
						"domain_name":  "oracle.com",
						"gateway":      "192.169.20.2",
						"netmask":      "255.255.0.1",
						"network_type": "BACKUP",
						"nodes.#":      "2",
					},
						[]string{
							"vlan_id",
						}),
					resource.TestCheckResourceAttr(resourceName, "state", "VALIDATED"),
				),
			},
			//  delete before next create
			{
				Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies,
			},
			// verify create without validation
			{
				Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update,
						representationCopyWithRemovedProperties(vmClusterNetworkValidateRepresentation, []string{`validate_vm_cluster_network`})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
					resource.TestCheckResourceAttr(resourceName, "dns.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "ntp.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
						"hostname": "myprefix2-ivmmj-scan",
						"ips.#":    "3",
						"port":     "1522",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
						"domain_name":  "oracle.com",
						"gateway":      "192.169.20.2",
						"netmask":      "255.255.0.1",
						"network_type": "BACKUP",
						"nodes.#":      "2",
					},
						[]string{
							"vlan_id",
						}),
					resource.TestCheckResourceAttr(resourceName, "state", "REQUIRES_VALIDATION"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify validation
			{
				Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
					resource.TestCheckResourceAttr(resourceName, "dns.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "ntp.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
						"hostname": "myprefix2-ivmmj-scan",
						"ips.#":    "3",
						"port":     "1522",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
						"domain_name":  "oracle.com",
						"gateway":      "192.169.20.2",
						"netmask":      "255.255.0.1",
						"network_type": "BACKUP",
						"nodes.#":      "2",
					},
						[]string{
							"vlan_id",
						}),
					resource.TestCheckResourceAttr(resourceName, "state", "VALIDATED"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// verify update after validation
			{
				Config: config + compartmentIdVariableStr + VmClusterNetworkResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Create, vmClusterNetworkRepresentation),
				ExpectError: regexp.MustCompile("update not allowed on validated vm cluster network"),
			},
		},
	})
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseValidatedVmClusterNetwork") {
		resource.AddTestSweepers("DatabaseValidatedVmClusterNetwork", &resource.Sweeper{
			Name:         "DatabaseValidatedVmClusterNetwork",
			Dependencies: DependencyGraph["vmClusterNetwork"],
			F:            sweepDatabaseValidatedVmClusterNetworkResource,
		})
	}
}

func sweepDatabaseValidatedVmClusterNetworkResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient
	vmClusterNetworkIds, err := getVmClusterNetworkIds(compartment)
	if err != nil {
		return err
	}
	for _, vmClusterNetworkId := range vmClusterNetworkIds {
		if ok := SweeperDefaultResourceId[vmClusterNetworkId]; !ok {
			deleteVmClusterNetworkRequest := oci_database.DeleteVmClusterNetworkRequest{}

			deleteVmClusterNetworkRequest.VmClusterNetworkId = &vmClusterNetworkId

			deleteVmClusterNetworkRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteVmClusterNetwork(context.Background(), deleteVmClusterNetworkRequest)
			if error != nil {
				fmt.Printf("Error deleting VmClusterNetwork %s %s, It is possible that the resource is already deleted. Please verify manually \n", vmClusterNetworkId, error)
				continue
			}
			waitTillCondition(testAccProvider, &vmClusterNetworkId, vmClusterNetworkSweepWaitCondition, time.Duration(3*time.Minute),
				vmClusterNetworkSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getValidatedVmClusterNetworkIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VmClusterNetworkId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient

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
			addResourceIdToSweeperResourceIdMap(compartmentId, "VmClusterNetworkId", id)
		}

	}
	return resourceIds, nil
}
