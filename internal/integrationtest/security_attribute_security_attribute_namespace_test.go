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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_security_attribute "github.com/oracle/oci-go-sdk/v65/securityattribute"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ignoreChangesSecurityAttributeNamespaceRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"defined_tags"}},
	}

	SecurityAttributeSecurityAttributeNamespaceRequiredOnlyResource = SecurityAttributeSecurityAttributeNamespaceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute_namespace", "test_security_attribute_namespace", acctest.Required, acctest.Create, SecurityAttributeSecurityAttributeNamespaceRepresentation)

	SecurityAttributeSecurityAttributeNamespaceResourceConfig = SecurityAttributeSecurityAttributeNamespaceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute_namespace", "test_security_attribute_namespace", acctest.Optional, acctest.Update, SecurityAttributeSecurityAttributeNamespaceRepresentation)

	SecurityAttributeSecurityAttributeNamespaceSingularDataSourceRepresentation = map[string]interface{}{
		"security_attribute_namespace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace.id}`},
	}

	SecurityAttributeSecurityAttributeNamespaceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `example-security-attribute-namespace`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: SecurityAttributeSecurityAttributeNamespaceDataSourceFilterRepresentation}}
	SecurityAttributeSecurityAttributeNamespaceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace.id}`}},
	}

	SecurityAttributeSecurityAttributeNamespaceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `This is the Zero Trust Packet Routing security attribute namespace.`, Update: `description2`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `example-security-attribute-namespace`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesSecurityAttributeNamespaceRepresentation},
	}

	SecurityAttributeSecurityAttributeNamespaceResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: security_attribute/default
func TestSecurityAttributeSecurityAttributeNamespaceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSecurityAttributeSecurityAttributeNamespaceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace"
	datasourceName := "data.oci_security_attribute_security_attribute_namespaces.test_security_attribute_namespaces"
	singularDatasourceName := "data.oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SecurityAttributeSecurityAttributeNamespaceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute_namespace", "test_security_attribute_namespace", acctest.Optional, acctest.Create, SecurityAttributeSecurityAttributeNamespaceRepresentation), "securityattribute", "securityAttributeNamespace", t)

	acctest.ResourceTest(t, testAccCheckSecurityAttributeSecurityAttributeNamespaceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SecurityAttributeSecurityAttributeNamespaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute_namespace", "test_security_attribute_namespace", acctest.Required, acctest.Create, SecurityAttributeSecurityAttributeNamespaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "This is the Zero Trust Packet Routing security attribute namespace."),
				resource.TestCheckResourceAttr(resourceName, "name", "example-security-attribute-namespace"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SecurityAttributeSecurityAttributeNamespaceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SecurityAttributeSecurityAttributeNamespaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute_namespace", "test_security_attribute_namespace", acctest.Optional, acctest.Create, SecurityAttributeSecurityAttributeNamespaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "This is the Zero Trust Packet Routing security attribute namespace."),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example-security-attribute-namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SecurityAttributeSecurityAttributeNamespaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute_namespace", "test_security_attribute_namespace", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(SecurityAttributeSecurityAttributeNamespaceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "This is the Zero Trust Packet Routing security attribute namespace."),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example-security-attribute-namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + SecurityAttributeSecurityAttributeNamespaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute_namespace", "test_security_attribute_namespace", acctest.Optional, acctest.Update, SecurityAttributeSecurityAttributeNamespaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example-security-attribute-namespace"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_security_attribute_security_attribute_namespaces", "test_security_attribute_namespaces", acctest.Optional, acctest.Update, SecurityAttributeSecurityAttributeNamespaceDataSourceRepresentation) +
				compartmentIdVariableStr + SecurityAttributeSecurityAttributeNamespaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute_namespace", "test_security_attribute_namespace", acctest.Optional, acctest.Update, SecurityAttributeSecurityAttributeNamespaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "name", "example-security-attribute-namespace"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "security_attribute_namespaces.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_attribute_namespaces.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "security_attribute_namespaces.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "security_attribute_namespaces.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_attribute_namespaces.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "security_attribute_namespaces.0.is_retired", "false"),
				resource.TestCheckResourceAttr(datasourceName, "security_attribute_namespaces.0.mode.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "security_attribute_namespaces.0.name", "example-security-attribute-namespace"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_attribute_namespaces.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_attribute_namespaces.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_security_attribute_security_attribute_namespace", "test_security_attribute_namespace", acctest.Required, acctest.Create, SecurityAttributeSecurityAttributeNamespaceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SecurityAttributeSecurityAttributeNamespaceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_attribute_namespace_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mode.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "example-security-attribute-namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + SecurityAttributeSecurityAttributeNamespaceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"defined_tags"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckSecurityAttributeSecurityAttributeNamespaceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).SecurityAttributeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_security_attribute_security_attribute_namespace" {
			noResourceFound = false
			request := oci_security_attribute.GetSecurityAttributeNamespaceRequest{}

			tmp := rs.Primary.ID
			request.SecurityAttributeNamespaceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "security_attribute")

			response, err := client.GetSecurityAttributeNamespace(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_security_attribute.SecurityAttributeNamespaceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("SecurityAttributeSecurityAttributeNamespace") {
		resource.AddTestSweepers("SecurityAttributeSecurityAttributeNamespace", &resource.Sweeper{
			Name:         "SecurityAttributeSecurityAttributeNamespace",
			Dependencies: acctest.DependencyGraph["securityAttributeNamespace"],
			F:            sweepSecurityAttributeSecurityAttributeNamespaceResource,
		})
	}
}

func sweepSecurityAttributeSecurityAttributeNamespaceResource(compartment string) error {
	securityAttributeClient := acctest.GetTestClients(&schema.ResourceData{}).SecurityAttributeClient()
	securityAttributeNamespaceIds, err := getSecurityAttributeSecurityAttributeNamespaceIds(compartment)
	if err != nil {
		return err
	}
	for _, securityAttributeNamespaceId := range securityAttributeNamespaceIds {
		if ok := acctest.SweeperDefaultResourceId[securityAttributeNamespaceId]; !ok {
			deleteSecurityAttributeNamespaceRequest := oci_security_attribute.DeleteSecurityAttributeNamespaceRequest{}

			deleteSecurityAttributeNamespaceRequest.SecurityAttributeNamespaceId = &securityAttributeNamespaceId

			deleteSecurityAttributeNamespaceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "security_attribute")
			_, error := securityAttributeClient.DeleteSecurityAttributeNamespace(context.Background(), deleteSecurityAttributeNamespaceRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityAttributeNamespace %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityAttributeNamespaceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &securityAttributeNamespaceId, SecurityAttributeSecurityAttributeNamespaceSweepWaitCondition, time.Duration(3*time.Minute),
				SecurityAttributeSecurityAttributeNamespaceSweepResponseFetchOperation, "security_attribute", true)
		}
	}
	return nil
}

func getSecurityAttributeSecurityAttributeNamespaceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SecurityAttributeNamespaceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	securityAttributeClient := acctest.GetTestClients(&schema.ResourceData{}).SecurityAttributeClient()

	listSecurityAttributeNamespacesRequest := oci_security_attribute.ListSecurityAttributeNamespacesRequest{}
	listSecurityAttributeNamespacesRequest.CompartmentId = &compartmentId
	listSecurityAttributeNamespacesRequest.LifecycleState = oci_security_attribute.SecurityAttributeNamespaceLifecycleStateActive
	listSecurityAttributeNamespacesResponse, err := securityAttributeClient.ListSecurityAttributeNamespaces(context.Background(), listSecurityAttributeNamespacesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityAttributeNamespace list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityAttributeNamespace := range listSecurityAttributeNamespacesResponse.Items {
		id := *securityAttributeNamespace.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityAttributeNamespaceId", id)
	}
	return resourceIds, nil
}

func SecurityAttributeSecurityAttributeNamespaceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if securityAttributeNamespaceResponse, ok := response.Response.(oci_security_attribute.GetSecurityAttributeNamespaceResponse); ok {
		return securityAttributeNamespaceResponse.LifecycleState != oci_security_attribute.SecurityAttributeNamespaceLifecycleStateDeleted
	}
	return false
}

func SecurityAttributeSecurityAttributeNamespaceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.SecurityAttributeClient().GetSecurityAttributeNamespace(context.Background(), oci_security_attribute.GetSecurityAttributeNamespaceRequest{
		SecurityAttributeNamespaceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
