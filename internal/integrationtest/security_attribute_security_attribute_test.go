// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	SecurityAttributeSecurityAttributeRequiredOnlyResource = SecurityAttributeSecurityAttributeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute", "test_security_attribute", acctest.Required, acctest.Create, SecurityAttributeSecurityAttributeRepresentation)

	SecurityAttributeSecurityAttributeResourceConfig = SecurityAttributeSecurityAttributeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute", "test_security_attribute", acctest.Optional, acctest.Update, SecurityAttributeSecurityAttributeRepresentation)

	SecurityAttributeSecurityAttributeSingularDataSourceRepresentation = map[string]interface{}{
		"security_attribute_name":         acctest.Representation{RepType: acctest.Required, Create: `${oci_security_attribute_security_attribute.test_security_attribute.name}`},
		"security_attribute_namespace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_security_attribute_security_attribute_namespace.security_attribute_namespace1.id}`},
	}

	SecurityAttributeSecurityAttributeDataSourceRepresentation = map[string]interface{}{
		"security_attribute_namespace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_security_attribute_security_attribute_namespace.security_attribute_namespace1.id}`},
		"state":                           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: SecurityAttributeSecurityAttributeDataSourceFilterRepresentation}}
	SecurityAttributeSecurityAttributeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_security_attribute_security_attribute.test_security_attribute.name}`}},
	}

	SecurityAttributeSecurityAttributeRetireRepresentation = map[string]interface{}{
		"description":                     acctest.Representation{RepType: acctest.Required, Create: `This security attribute will be used for billing of associated resources.`, Update: `description2`},
		"name":                            acctest.Representation{RepType: acctest.Required, Create: `TFTestSecurityAttribute`},
		"security_attribute_namespace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_security_attribute_security_attribute_namespace.security_attribute_namespace1.id}`},
		"is_retired":                      acctest.Representation{RepType: acctest.Optional, Create: "false", Update: "true"},
	}

	SecurityAttributeSecurityAttributeRepresentation = map[string]interface{}{
		"description":                     acctest.Representation{RepType: acctest.Required, Create: `This security attribute will be used for billing of associated resources.`, Update: `description2`},
		"name":                            acctest.Representation{RepType: acctest.Required, Create: `TFTestSecurityAttribute`},
		"security_attribute_namespace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_security_attribute_security_attribute_namespace.security_attribute_namespace1.id}`},
		"is_retired":                      acctest.Representation{RepType: acctest.Optional, Create: "false", Update: "true"},
		"validator":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: SecurityAttributeSecurityAttributeValidatorRepresentation},
	}
	SecurityAttributeSecurityAttributeValidatorRepresentation = map[string]interface{}{
		"validator_type": acctest.Representation{RepType: acctest.Required, Create: `ENUM`},
		"values":         acctest.Representation{RepType: acctest.Required, Create: []string{`value1`, `value2`}},
	}

	SecurityAttributeSecurityAttributeResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute_namespace", "security_attribute_namespace1", acctest.Required, acctest.Create, SecurityAttributeSecurityAttributeNamespaceRepresentation)
)

// issue-routing-tag: security_attribute/default
func TestSecurityAttributeSecurityAttributeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSecurityAttributeSecurityAttributeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_security_attribute_security_attribute.test_security_attribute"
	datasourceName := "data.oci_security_attribute_security_attributes.test_security_attributes"
	singularDatasourceName := "data.oci_security_attribute_security_attribute.test_security_attribute"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SecurityAttributeSecurityAttributeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute", "test_security_attribute", acctest.Optional, acctest.Create, SecurityAttributeSecurityAttributeRepresentation), "securityattribute", "securityAttribute", t)

	acctest.ResourceTest(t, testAccCheckSecurityAttributeSecurityAttributeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SecurityAttributeSecurityAttributeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute", "test_security_attribute", acctest.Required, acctest.Create, SecurityAttributeSecurityAttributeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "This security attribute will be used for billing of associated resources."),
				resource.TestCheckResourceAttr(resourceName, "name", "TFTestSecurityAttribute"),
				resource.TestCheckResourceAttrSet(resourceName, "security_attribute_namespace_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//Update is_retired to true before deleting
		{
			Config: config + compartmentIdVariableStr + SecurityAttributeSecurityAttributeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute", "test_security_attribute", acctest.Optional, acctest.Update, SecurityAttributeSecurityAttributeRetireRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "TFTestSecurityAttribute"),
				resource.TestCheckResourceAttrSet(resourceName, "security_attribute_namespace_id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_attribute_namespace_name"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "validator.#", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SecurityAttributeSecurityAttributeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SecurityAttributeSecurityAttributeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute", "test_security_attribute", acctest.Optional, acctest.Create, SecurityAttributeSecurityAttributeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "This security attribute will be used for billing of associated resources."),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "TFTestSecurityAttribute"),
				resource.TestCheckResourceAttrSet(resourceName, "security_attribute_namespace_id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_attribute_namespace_name"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "validator.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "validator.0.validator_type", "ENUM"),
				resource.TestCheckResourceAttr(resourceName, "validator.0.values.#", "2"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + SecurityAttributeSecurityAttributeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute", "test_security_attribute", acctest.Optional, acctest.Update, SecurityAttributeSecurityAttributeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_retired", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "TFTestSecurityAttribute"),
				resource.TestCheckResourceAttrSet(resourceName, "security_attribute_namespace_id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_attribute_namespace_name"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "validator.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "validator.0.validator_type", "ENUM"),
				resource.TestCheckResourceAttr(resourceName, "validator.0.values.#", "2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_security_attribute_security_attributes", "test_security_attributes", acctest.Optional, acctest.Update, SecurityAttributeSecurityAttributeDataSourceRepresentation) +
				compartmentIdVariableStr + SecurityAttributeSecurityAttributeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_security_attribute_security_attribute", "test_security_attribute", acctest.Optional, acctest.Update, SecurityAttributeSecurityAttributeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "security_attribute_namespace_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "security_attributes.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_attributes.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "security_attributes.0.description", "description2"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_attributes.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "security_attributes.0.is_retired", "true"),
				resource.TestCheckResourceAttr(datasourceName, "security_attributes.0.name", "TFTestSecurityAttribute"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_attributes.0.security_attribute_namespace_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_attributes.0.security_attribute_namespace_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_attributes.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_attributes.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_security_attribute_security_attribute", "test_security_attribute", acctest.Required, acctest.Create, SecurityAttributeSecurityAttributeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SecurityAttributeSecurityAttributeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_attribute_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_attribute_namespace_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_retired", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TFTestSecurityAttribute"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_attribute_namespace_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "validator.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "validator.0.validator_type", "ENUM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "validator.0.values.#", "2"),
			),
		},
		// verify resource import
		{
			Config:                  config + SecurityAttributeSecurityAttributeRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckSecurityAttributeSecurityAttributeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).SecurityAttributeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_security_attribute_security_attribute" {
			noResourceFound = false
			request := oci_security_attribute.GetSecurityAttributeRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.SecurityAttributeName = &value
			}

			if value, ok := rs.Primary.Attributes["security_attribute_namespace_id"]; ok {
				request.SecurityAttributeNamespaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "security_attribute")

			response, err := client.GetSecurityAttribute(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_security_attribute.SecurityAttributeLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("SecurityAttributeSecurityAttribute") {
		resource.AddTestSweepers("SecurityAttributeSecurityAttribute", &resource.Sweeper{
			Name:         "SecurityAttributeSecurityAttribute",
			Dependencies: acctest.DependencyGraph["securityAttribute"],
			F:            sweepSecurityAttributeSecurityAttributeResource,
		})
	}
}

func sweepSecurityAttributeSecurityAttributeResource(compartment string) error {
	securityAttributeClient := acctest.GetTestClients(&schema.ResourceData{}).SecurityAttributeClient()
	securityAttributeIds, err := getSecurityAttributeSecurityAttributeIds(compartment)
	if err != nil {
		return err
	}
	for _, securityAttributeId := range securityAttributeIds {
		if ok := acctest.SweeperDefaultResourceId[securityAttributeId]; !ok {
			deleteSecurityAttributeRequest := oci_security_attribute.DeleteSecurityAttributeRequest{}

			deleteSecurityAttributeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "security_attribute")

			_, error := securityAttributeClient.DeleteSecurityAttribute(context.Background(), deleteSecurityAttributeRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityAttribute %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityAttributeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &securityAttributeId, SecurityAttributeSecurityAttributeSweepWaitCondition, time.Duration(3*time.Minute),
				SecurityAttributeSecurityAttributeSweepResponseFetchOperation, "security_attribute", true)
		}
	}
	return nil
}

func getSecurityAttributeSecurityAttributeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SecurityAttributeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	securityAttributeClient := acctest.GetTestClients(&schema.ResourceData{}).SecurityAttributeClient()

	listSecurityAttributesRequest := oci_security_attribute.ListSecurityAttributesRequest{}

	securityAttributeNamespaceIds, error := getSecurityAttributeSecurityAttributeNamespaceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting securityAttributeNamespaceId required for SecurityAttribute resource requests \n")
	}
	for _, securityAttributeNamespaceId := range securityAttributeNamespaceIds {
		listSecurityAttributesRequest.SecurityAttributeNamespaceId = &securityAttributeNamespaceId

		listSecurityAttributesRequest.LifecycleState = oci_security_attribute.SecurityAttributeLifecycleStateActive
		listSecurityAttributesResponse, err := securityAttributeClient.ListSecurityAttributes(context.Background(), listSecurityAttributesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting SecurityAttribute list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, securityAttribute := range listSecurityAttributesResponse.Items {
			id := *securityAttribute.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityAttributeId", id)
		}

	}
	return resourceIds, nil
}

func SecurityAttributeSecurityAttributeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if securityAttributeResponse, ok := response.Response.(oci_security_attribute.GetSecurityAttributeResponse); ok {
		return securityAttributeResponse.LifecycleState != oci_security_attribute.SecurityAttributeLifecycleStateDeleted
	}
	return false
}

func SecurityAttributeSecurityAttributeSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.SecurityAttributeClient().GetSecurityAttribute(context.Background(), oci_security_attribute.GetSecurityAttributeRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
