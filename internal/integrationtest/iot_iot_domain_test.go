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
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IotIotDomainRequiredOnlyResource = IotIotDomainResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain", "test_iot_domain", acctest.Required, acctest.Create, IotIotDomainRepresentation)

	IotIotDomainResourceConfig = IotIotDomainResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain", "test_iot_domain", acctest.Optional, acctest.Update, IotIotDomainRepresentation)

	IotIotDomainSingularDataSourceRepresentation = map[string]interface{}{
		"iot_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_iot_domain.test_iot_domain.id}`},
	}
	ignoreIotDomainDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	IotIotDomainDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_iot_domain.test_iot_domain.id}`},
		"iot_domain_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.iot_domain_group_id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: IotIotDomainDataSourceFilterRepresentation}}
	IotIotDomainDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_iot_iot_domain.test_iot_domain.id}`}},
	}

	IotIotDomainRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"iot_domain_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_group_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Protocol": "Mqtt"}, Update: map[string]string{"Protocol": "MQTT"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreIotDomainDefinedTagsChangesRepresentation},
	}

	IotIotDomainResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: iot/default
func TestIotIotDomainResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotIotDomainResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	iotDomainGroupId := utils.GetEnvSettingWithBlankDefault("iot_domain_group_ocid")
	iotDomainGroupIdVariableStr := fmt.Sprintf("variable \"iot_domain_group_id\" { default = \"%s\" }\n", iotDomainGroupId)

	resourceName := "oci_iot_iot_domain.test_iot_domain"
	datasourceName := "data.oci_iot_iot_domains.test_iot_domains"
	singularDatasourceName := "data.oci_iot_iot_domain.test_iot_domain"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+iotDomainGroupIdVariableStr+IotIotDomainResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain", "test_iot_domain", acctest.Optional, acctest.Create, IotIotDomainRepresentation), "iot", "iotDomain", t)

	acctest.ResourceTest(t, testAccCheckIotIotDomainDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + iotDomainGroupIdVariableStr + IotIotDomainResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain", "test_iot_domain", acctest.Required, acctest.Create, IotIotDomainRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_group_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + iotDomainGroupIdVariableStr + IotIotDomainResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + iotDomainGroupIdVariableStr + IotIotDomainResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain", "test_iot_domain", acctest.Optional, acctest.Create, IotIotDomainRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + iotDomainGroupIdVariableStr + IotIotDomainResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain", "test_iot_domain", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(IotIotDomainRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + iotDomainGroupIdVariableStr + IotIotDomainResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain", "test_iot_domain", acctest.Optional, acctest.Update, IotIotDomainRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_iot_domains", "test_iot_domains", acctest.Optional, acctest.Update, IotIotDomainDataSourceRepresentation) +
				compartmentIdVariableStr + iotDomainGroupIdVariableStr + IotIotDomainResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain", "test_iot_domain", acctest.Optional, acctest.Update, IotIotDomainRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "iot_domain_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "iot_domain_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "iot_domain_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_iot_domain", "test_iot_domain", acctest.Required, acctest.Create, IotIotDomainSingularDataSourceRepresentation) +
				compartmentIdVariableStr + iotDomainGroupIdVariableStr + IotIotDomainResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "iot_domain_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_retention_periods_in_days.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "device_host"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + IotIotDomainRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIotIotDomainDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IotClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_iot_iot_domain" {
			noResourceFound = false
			request := oci_iot.GetIotDomainRequest{}

			tmp := rs.Primary.ID
			request.IotDomainId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")

			response, err := client.GetIotDomain(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_iot.IotDomainLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("IotIotDomain") {
		resource.AddTestSweepers("IotIotDomain", &resource.Sweeper{
			Name:         "IotIotDomain",
			Dependencies: acctest.DependencyGraph["iotDomain"],
			F:            sweepIotIotDomainResource,
		})
	}
}

func sweepIotIotDomainResource(compartment string) error {
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()
	iotDomainIds, err := getIotIotDomainIds(compartment)
	if err != nil {
		return err
	}
	for _, iotDomainId := range iotDomainIds {
		if ok := acctest.SweeperDefaultResourceId[iotDomainId]; !ok {
			deleteIotDomainRequest := oci_iot.DeleteIotDomainRequest{}

			deleteIotDomainRequest.IotDomainId = &iotDomainId

			deleteIotDomainRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")
			_, error := iotClient.DeleteIotDomain(context.Background(), deleteIotDomainRequest)
			if error != nil {
				fmt.Printf("Error deleting IotDomain %s %s, It is possible that the resource is already deleted. Please verify manually \n", iotDomainId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &iotDomainId, IotIotDomainSweepWaitCondition, time.Duration(3*time.Minute),
				IotIotDomainSweepResponseFetchOperation, "iot", true)
		}
	}
	return nil
}

func getIotIotDomainIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IotDomainId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()

	listIotDomainsRequest := oci_iot.ListIotDomainsRequest{}
	listIotDomainsRequest.CompartmentId = &compartmentId
	listIotDomainsRequest.LifecycleState = oci_iot.IotDomainLifecycleStateActive
	listIotDomainsResponse, err := iotClient.ListIotDomains(context.Background(), listIotDomainsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IotDomain list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, iotDomain := range listIotDomainsResponse.Items {
		id := *iotDomain.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IotDomainId", id)
	}
	return resourceIds, nil
}

func IotIotDomainSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if iotDomainResponse, ok := response.Response.(oci_iot.GetIotDomainResponse); ok {
		return iotDomainResponse.LifecycleState != oci_iot.IotDomainLifecycleStateDeleted
	}
	return false
}

func IotIotDomainSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.IotClient().GetIotDomain(context.Background(), oci_iot.GetIotDomainRequest{
		IotDomainId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
