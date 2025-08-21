// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSecurityPolicyDeploymentManagementResourceConfig = DataSafeSecurityPolicyDeploymentManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment_management", "test_security_policy_deployment_management", acctest.Optional, acctest.Create, DataSafeSecurityPolicyDeploymentManagementRepresentation)

	DataSafeSecurityPolicyDeploymentManagementRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"security_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_id}`},
		"target_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"target_type":        acctest.Representation{RepType: acctest.Required, Create: `TARGET_DATABASE`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"deploy_trigger":     acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"refresh_trigger":    acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSecurityPolicyDeploymentManagementTagsChangesRep},
	}
	ignoreSecurityPolicyDeploymentManagementTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSecurityPolicyDeploymentManagementResourceDependencies = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyDeploymentManagementResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the target ocid & security policy ocid are hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSecurityPolicyDeploymentManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	securityPolicyId := utils.GetEnvSettingWithBlankDefault("security_policy_ocid")
	securityPolicyIdVariableStr := fmt.Sprintf("variable \"security_policy_id\" { default = \"%s\" }\n", securityPolicyId)

	resourceName := "oci_data_safe_security_policy_deployment_management.test_security_policy_deployment_management"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeSecurityPolicyDeploymentManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment_management", "test_security_policy_deployment_management", acctest.Optional, acctest.Create, DataSafeSecurityPolicyDeploymentManagementRepresentation), "datasafe", "securityPolicyDeploymentManagement", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSecurityPolicyDeploymentManagementDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyDeploymentManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment_management", "test_security_policy_deployment_management", acctest.Required, acctest.Create, DataSafeSecurityPolicyDeploymentManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSecurityPolicyDeploymentManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyDeploymentManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment_management", "test_security_policy_deployment_management", acctest.Optional, acctest.Create, DataSafeSecurityPolicyDeploymentManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + securityPolicyIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyDeploymentManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment_management", "test_security_policy_deployment_management", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyDeploymentManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyDeploymentManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment_management", "test_security_policy_deployment_management", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify deploy
		{
			Config: config + compartmentIdVariableStr + DataSafeSecurityPolicyDeploymentManagementResourceDependencies + securityPolicyIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment_management", "test_security_policy_deployment_management", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be deployed.")
					}
					return err
				},
			),
		},
		// verify refresh
		{
			Config: config + compartmentIdVariableStr + DataSafeSecurityPolicyDeploymentManagementResourceDependencies + securityPolicyIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment_management", "test_security_policy_deployment_management", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be refreshed.")
					}
					return err
				},
			),
		},
	})
}

func testAccCheckDataSafeSecurityPolicyDeploymentManagementDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_security_policy_deployment_management" {
			noResourceFound = false
			request := oci_data_safe.GetSecurityPolicyDeploymentRequest{}

			tmp := rs.Primary.ID
			request.SecurityPolicyDeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSecurityPolicyDeployment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SecurityPolicyDeploymentLifecycleStateDeleted): true,
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
