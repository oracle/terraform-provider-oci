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
	oci_psa "github.com/oracle/oci-go-sdk/v65/psa"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PrivateServiceAccessIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	PsaPrivateServiceAccesRequiredOnlyResource = PsaPrivateServiceAccesResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_psa_private_service_access", "test_private_service_access", acctest.Required, acctest.Create, PsaPrivateServiceAccesRepresentation)

	PsaPrivateServiceAccesResourceConfig = PsaPrivateServiceAccesResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_psa_private_service_access", "test_private_service_access", acctest.Optional, acctest.Update, PsaPrivateServiceAccesRepresentation)

	PsaPrivateServiceAccesSingularDataSourceRepresentation = map[string]interface{}{
		"private_service_access_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_psa_private_service_access.test_private_service_access.id}`},
	}

	PsaPrivateServiceAccesDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"service_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_psa_psa_services.test_psa_services.psa_service_collection.0.items.0.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: PsaPrivateServiceAccesDataSourceFilterRepresentation}}

	PsaPrivateServiceAccesDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_psa_private_service_access.test_private_service_access.id}`}},
	}

	PsaPrivateServiceAccesRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"service_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_psa_psa_services.test_psa_services.psa_service_collection.0.items.0.id}`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ipv4ip":         acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.100`},
		"nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{`${oci_core_network_security_group.test_network_security_group1.id}`}},
		"security_attributes": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"oracle-zpr.sensitivity.value": "medium",
			"oracle-zpr.sensitivity.mode":  "enforce",
		}},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: PrivateServiceAccessIgnoreChangesRepresentation},
	}
	subnetCidrBlock = `${substr(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn.ipv6cidr_blocks[0]) - 2)}${64}`

	PsaPrivateServiceAccesResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group1", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
				"ipv6cidr_blocks": acctest.Representation{RepType: acctest.Optional, Create: []string{subnetCidrBlock}},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreIpv6VcnRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_psa_psa_services", "test_psa_services", acctest.Required, acctest.Create, PsaPsaServiceDataSourceRepresentation)
)

// issue-routing-tag: psa/default
func TestPsaPrivateServiceAccesResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsaPrivateServiceAccesResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_psa_private_service_access.test_private_service_access"
	datasourceName := "data.oci_psa_private_service_accesses.test_private_service_access"
	singularDatasourceName := "data.oci_psa_private_service_access.test_private_service_access"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PsaPrivateServiceAccesResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_psa_private_service_access", "test_private_service_access", acctest.Optional, acctest.Create, PsaPrivateServiceAccesRepresentation), "psa", "privateServiceAcces", t)

	acctest.ResourceTest(t, testAccCheckPsaPrivateServiceAccesDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PsaPrivateServiceAccesResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psa_private_service_access", "test_private_service_access", acctest.Required, acctest.Create, PsaPrivateServiceAccesRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "service_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PsaPrivateServiceAccesResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PsaPrivateServiceAccesResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psa_private_service_access", "test_private_service_access", acctest.Optional, acctest.Create, PsaPrivateServiceAccesRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdns.#"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4ip", "10.0.0.100"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "service_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PsaPrivateServiceAccesResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psa_private_service_access", "test_private_service_access", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(PsaPrivateServiceAccesRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdns.#"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4ip", "10.0.0.100"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "service_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + PsaPrivateServiceAccesResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psa_private_service_access", "test_private_service_access", acctest.Optional, acctest.Update, PsaPrivateServiceAccesRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdns.#"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4ip", "10.0.0.100"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "service_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psa_private_service_accesses", "test_private_service_access", acctest.Optional, acctest.Update, PsaPrivateServiceAccesDataSourceRepresentation) +
				compartmentIdVariableStr + PsaPrivateServiceAccesResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psa_private_service_access", "test_private_service_access", acctest.Optional, acctest.Update, PsaPrivateServiceAccesRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "service_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "private_service_access_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "private_service_access_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "private_service_access_collection.0.items.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "private_service_access_collection.0.items.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_service_access_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_service_access_collection.0.items.0.service_id"),
				resource.TestCheckResourceAttr(datasourceName, "private_service_access_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_service_access_collection.0.items.0.vcn_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psa_private_service_access", "test_private_service_access", acctest.Required, acctest.Create, PsaPrivateServiceAccesSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PsaPrivateServiceAccesResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_service_access_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fqdns.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ipv4ip", "10.0.0.100"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vnic_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + PsaPrivateServiceAccesRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckPsaPrivateServiceAccesDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).PrivateServiceAccessClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_psa_private_service_access" {
			noResourceFound = false
			request := oci_psa.GetPrivateServiceAccessRequest{}

			if value, ok := rs.Primary.Attributes["private_service_access_id"]; ok {
				request.PrivateServiceAccessId = &value
			} else {
				tmp := rs.Primary.ID
				request.PrivateServiceAccessId = &tmp
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "psa")

			response, err := client.GetPrivateServiceAccess(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_psa.PrivateServiceAccessLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("PsaPrivateServiceAcces") {
		resource.AddTestSweepers("PsaPrivateServiceAcces", &resource.Sweeper{
			Name:         "PsaPrivateServiceAcces",
			Dependencies: acctest.DependencyGraph["privateServiceAcces"],
			F:            sweepPsaPrivateServiceAccesResource,
		})
	}
}

func sweepPsaPrivateServiceAccesResource(compartment string) error {
	privateServiceAccessClient := acctest.GetTestClients(&schema.ResourceData{}).PrivateServiceAccessClient()
	privateServiceAccesIds, err := getPsaPrivateServiceAccesIds(compartment)
	if err != nil {
		return err
	}
	for _, privateServiceAccesId := range privateServiceAccesIds {
		if ok := acctest.SweeperDefaultResourceId[privateServiceAccesId]; !ok {
			deletePrivateServiceAccessRequest := oci_psa.DeletePrivateServiceAccessRequest{}
			deletePrivateServiceAccessRequest.PrivateServiceAccessId = &privateServiceAccesId

			deletePrivateServiceAccessRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "psa")
			_, error := privateServiceAccessClient.DeletePrivateServiceAccess(context.Background(), deletePrivateServiceAccessRequest)
			if error != nil {
				fmt.Printf("Error deleting PrivateServiceAcces %s %s, It is possible that the resource is already deleted. Please verify manually \n", privateServiceAccesId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &privateServiceAccesId, PsaPrivateServiceAccesSweepWaitCondition, time.Duration(3*time.Minute),
				PsaPrivateServiceAccesSweepResponseFetchOperation, "psa", true)
		}
	}
	return nil
}

func getPsaPrivateServiceAccesIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PrivateServiceAccesId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	privateServiceAccessClient := acctest.GetTestClients(&schema.ResourceData{}).PrivateServiceAccessClient()

	listPrivateServiceAccessesRequest := oci_psa.ListPrivateServiceAccessesRequest{}
	listPrivateServiceAccessesRequest.CompartmentId = &compartmentId
	listPrivateServiceAccessesRequest.LifecycleState = oci_psa.PrivateServiceAccessLifecycleStateActive
	listPrivateServiceAccessesResponse, err := privateServiceAccessClient.ListPrivateServiceAccesses(context.Background(), listPrivateServiceAccessesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PrivateServiceAcces list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, privateServiceAcces := range listPrivateServiceAccessesResponse.Items {
		id := *privateServiceAcces.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PrivateServiceAccesId", id)
	}
	return resourceIds, nil
}

func PsaPrivateServiceAccesSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if privateServiceAccesResponse, ok := response.Response.(oci_psa.GetPrivateServiceAccessResponse); ok {
		return privateServiceAccesResponse.LifecycleState != oci_psa.PrivateServiceAccessLifecycleStateDeleted
	}
	return false
}

func PsaPrivateServiceAccesSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.PrivateServiceAccessClient().GetPrivateServiceAccess(context.Background(), oci_psa.GetPrivateServiceAccessRequest{
		PrivateServiceAccessId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
