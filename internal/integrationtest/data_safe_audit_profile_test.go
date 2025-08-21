// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

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
	DataSafeAuditProfileRequiredOnlyResource = DataSafeAuditProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Create, DataSafeAuditProfileRepresentation)

	DataSafeAuditProfileRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_database_group.test_target_database_group.id}`},
		"target_type":                   acctest.Representation{RepType: acctest.Required, Create: `TARGET_DATABASE_GROUP`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `Audit_1`, Update: `displayName2`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `Description`, Update: `description2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_override_global_paid_usage": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_paid_usage_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"offline_months":                acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"online_months":                 acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DataSafeAuditProfileTargetRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.target_ocid}`},
		"target_type":                   acctest.Representation{RepType: acctest.Required, Create: `TARGET_DATABASE`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `Audit_1`, Update: `displayName2`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `Description`, Update: `description2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_override_global_paid_usage": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_paid_usage_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"offline_months":                acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"online_months":                 acctest.Representation{RepType: acctest.Optional, Create: `10`},
		//"change_retention_trigger":             acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"is_override_global_retention_setting": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
	}
	DataSafeAuditProfileDataSourceRepresentation = map[string]interface{}{
		"audit_profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_audit_profile.test_audit_profile.id}`},
	}

	DataSafeAuditProfileResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_group", "test_target_database_group", acctest.Required, acctest.Create, DataSafeTargetDatabaseGroupRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAuditProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_audit_profile.test_audit_profile"
	datasourceName := "data.oci_data_safe_audit_profile.test_audit_profile"
	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeAuditProfileResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Create, DataSafeAuditProfileRepresentation), "datasafe", "auditProfile", t)

	acctest.ResourceTest(t, testAccCheckDataSafeAuditProfileDestroy, []resource.TestStep{
		// verify Create
		// Audit profile for target group creation with required parameters only
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Required, acctest.Create, DataSafeAuditProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE_GROUP"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					fmt.Print(resId)
					return err
				},
			),
		},

		// delete before next Create
		// Delete the created audit profile
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditProfileResourceDependencies,
		},
		// verify Create with optionals
		// Audit profile for target group creation with Optional parameters
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Create, DataSafeAuditProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Audit_1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_override_global_paid_usage", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "is_override_global_retention_setting"),
				resource.TestCheckResourceAttr(resourceName, "is_paid_usage_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "offline_months", "10"),
				resource.TestCheckResourceAttr(resourceName, "online_months", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE_GROUP"),
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
		// Update the compartment
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeAuditProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeAuditProfileRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Audit_1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_override_global_paid_usage", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "is_override_global_retention_setting"),
				resource.TestCheckResourceAttr(resourceName, "is_paid_usage_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "offline_months", "10"),
				resource.TestCheckResourceAttr(resourceName, "online_months", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE_GROUP"),
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
		// Update the Audit profile for target group
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Update, DataSafeAuditProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_override_global_paid_usage", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "is_override_global_retention_setting"),
				resource.TestCheckResourceAttr(resourceName, "is_paid_usage_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "offline_months", "10"),
				resource.TestCheckResourceAttr(resourceName, "online_months", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE_GROUP"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Update, DataSafeAuditProfileDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAuditProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Update, DataSafeAuditProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "audit_profile_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_paid_usage_enabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttr(datasourceName, "target_type", "TARGET_DATABASE_GROUP"),
			),
		},
		{
			Config:            config + compartmentIdVariableStr + DataSafeAuditProfileRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_override_global_paid_usage",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataSafeAuditProfileDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "oci_data_safe_audit_profile" {
			noResourceFound = false

			request := oci_data_safe.GetAuditProfileRequest{}

			tmp := rs.Primary.ID
			request.AuditProfileId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetAuditProfile(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.AuditProfileLifecycleStateDeleted): true,
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
