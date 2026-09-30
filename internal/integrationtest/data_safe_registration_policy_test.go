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
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeRegistrationPolicyRequiredOnlyResource = DataSafeRegistrationPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_registration_policy", "test_registration_policy", acctest.Required, acctest.Create, DataSafeRegistrationPolicyRepresentation)

	DataSafeRegistrationPolicyResourceConfig = DataSafeRegistrationPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_registration_policy", "test_registration_policy", acctest.Optional, acctest.Update, DataSafeRegistrationPolicyRepresentation)

	DataSafeRegistrationPolicySingularDataSourceRepresentation = map[string]interface{}{
		"registration_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_registration_policy.test_registration_policy.id}`},
	}

	DataSafeRegistrationPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                          acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"enablement_level":                      acctest.Representation{RepType: acctest.Optional, Create: `DATABASE`},
		"registration_policy_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_registration_policy.test_registration_policy.id}`},
		"resource_id":                           acctest.Representation{RepType: acctest.Optional, Create: `${var.registration_policy_resource_id}`},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeRegistrationPolicyDataSourceFilterRepresentation}}
	DataSafeRegistrationPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_registration_policy.test_registration_policy.id}`}},
	}

	DataSafeRegistrationPolicyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		// ASSESSMENT is always enabled; updates add features alongside it.
		"features":              acctest.Representation{RepType: acctest.Required, Create: []string{`ASSESSMENT`}, Update: []string{`ASSESSMENT`, `AUDIT_COLLECTION`}},
		"resource_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.registration_policy_resource_id}`},
		"can_override_features": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DataSafeRegistrationPolicyResourceDependencies = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeRegistrationPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeRegistrationPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	registrationPolicyResourceID := utils.GetEnvSettingWithBlankDefault("registration_policy_resource_id")
	registrationPolicyResourceIDVariableStr := fmt.Sprintf("variable \"registration_policy_resource_id\" { default = \"%s\" }\n", registrationPolicyResourceID)
	compartmentIdVariableStr += registrationPolicyResourceIDVariableStr

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	registrationPolicyProvisioningWaitSeconds, err := strconv.Atoi(utils.GetEnvSettingWithDefault("registration_policy_provisioning_wait_seconds", "0"))
	if err != nil || registrationPolicyProvisioningWaitSeconds < 0 {
		t.Fatalf("registration_policy_provisioning_wait_seconds must be a non-negative integer")
	}
	registrationPolicyUpdateWaitSeconds, err := strconv.Atoi(utils.GetEnvSettingWithDefault("registration_policy_update_wait_seconds", "360"))
	if err != nil || registrationPolicyUpdateWaitSeconds < 0 {
		t.Fatalf("registration_policy_update_wait_seconds must be a non-negative integer")
	}

	resourceName := "oci_data_safe_registration_policy.test_registration_policy"
	datasourceName := "data.oci_data_safe_registration_policies.test_registration_policies"
	singularDatasourceName := "data.oci_data_safe_registration_policy.test_registration_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeRegistrationPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_registration_policy", "test_registration_policy", acctest.Optional, acctest.Create, DataSafeRegistrationPolicyRepresentation), "datasafe", "registrationPolicy", t)

	acctest.ResourceTest(t, testAccCheckDataSafeRegistrationPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeRegistrationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_registration_policy", "test_registration_policy", acctest.Required, acctest.Create, DataSafeRegistrationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "features.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			PreConfig: func() {
				if registrationPolicyProvisioningWaitSeconds > 0 {
					t.Logf("waiting %d seconds for registration policy target provisioning before delete", registrationPolicyProvisioningWaitSeconds)
					time.Sleep(time.Duration(registrationPolicyProvisioningWaitSeconds) * time.Second)
				}
			},
			Config: config + compartmentIdVariableStr + DataSafeRegistrationPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeRegistrationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_registration_policy", "test_registration_policy", acctest.Optional, acctest.Create, DataSafeRegistrationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "can_override_features", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "features.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeRegistrationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_registration_policy", "test_registration_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeRegistrationPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "can_override_features", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "features.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			PreConfig: func() {
				if registrationPolicyUpdateWaitSeconds > 0 {
					t.Logf("waiting %d seconds for registration policy target provisioning before update", registrationPolicyUpdateWaitSeconds)
					time.Sleep(time.Duration(registrationPolicyUpdateWaitSeconds) * time.Second)
				}
			},
			Config: config + compartmentIdVariableStr + DataSafeRegistrationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_registration_policy", "test_registration_policy", acctest.Optional, acctest.Update, DataSafeRegistrationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "can_override_features", "true"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "features.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "features.0", "ASSESSMENT"),
				resource.TestCheckResourceAttr(resourceName, "features.1", "AUDIT_COLLECTION"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_registration_policies", "test_registration_policies", acctest.Optional, acctest.Update, DataSafeRegistrationPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeRegistrationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_registration_policy", "test_registration_policy", acctest.Optional, acctest.Update, DataSafeRegistrationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "enablement_level", "DATABASE"),
				resource.TestCheckResourceAttrSet(datasourceName, "registration_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "registration_policy_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "registration_policy_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_registration_policy", "test_registration_policy", acctest.Required, acctest.Create, DataSafeRegistrationPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeRegistrationPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "registration_policy_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "enablement_level"),
				resource.TestCheckResourceAttr(singularDatasourceName, "features.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "features.0", "ASSESSMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "features.1", "AUDIT_COLLECTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + compartmentIdVariableStr + DataSafeRegistrationPolicyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ResourceName:      resourceName,
		},
	})
}

func testAccCheckDataSafeRegistrationPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_registration_policy" {
			noResourceFound = false
			request := oci_data_safe.GetRegistrationPolicyRequest{}

			tmp := rs.Primary.ID
			request.RegistrationPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetRegistrationPolicy(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataSafeRegistrationPolicy") {
		resource.AddTestSweepers("DataSafeRegistrationPolicy", &resource.Sweeper{
			Name:         "DataSafeRegistrationPolicy",
			Dependencies: acctest.DependencyGraph["registrationPolicy"],
			F:            sweepDataSafeRegistrationPolicyResource,
		})
	}
}

func sweepDataSafeRegistrationPolicyResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	registrationPolicyIds, err := getDataSafeRegistrationPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, registrationPolicyId := range registrationPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[registrationPolicyId]; !ok {
			deleteRegistrationPolicyRequest := oci_data_safe.DeleteRegistrationPolicyRequest{}

			deleteRegistrationPolicyRequest.RegistrationPolicyId = &registrationPolicyId

			deleteRegistrationPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteRegistrationPolicy(context.Background(), deleteRegistrationPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting RegistrationPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", registrationPolicyId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeRegistrationPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RegistrationPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listRegistrationPoliciesRequest := oci_data_safe.ListRegistrationPoliciesRequest{}
	listRegistrationPoliciesRequest.CompartmentId = &compartmentId
	listRegistrationPoliciesResponse, err := dataSafeClient.ListRegistrationPolicies(context.Background(), listRegistrationPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting RegistrationPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, registrationPolicy := range listRegistrationPoliciesResponse.Items {
		id := *registrationPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RegistrationPolicyId", id)
	}
	return resourceIds, nil
}
