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
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

const (
	DefinedTagsDependencies_fusionapps = `
variable defined_tag_namespace_name { default = "" }
resource "oci_identity_tag_namespace" "tag-namespace-terraform" {
  		#Required
		compartment_id = "${var.tenancy_ocid}"
  		description = "example tag namespace"
  		name = "example-tag-namespace-all-terraform"

		is_retired = false
}

resource "oci_identity_tag" "tag1" {
  		#Required
  		description = "example tag"
  		name = "example-tag"
        tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace-terraform.id}"

		is_retired = false
}
`
)

var (
	KeyResourceDependencyConfig_fusionapps = KmsKeyResourceDependencies + `
	data "oci_kms_keys" "test_keys_dependency_fusionapps" {
		#Required
		compartment_id = "${data.oci_kms_vault.test_vault.compartment_id}"
		management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
		algorithm = "AES"

		filter {
    		name = "state"
    		values = ["ENABLED", "UPDATING"]
        }
	}
	`

	definedTagsIgnoreRepresentation_fusionapps = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{`defined_tags`}},
	}

	FusionAppsFusionEnvironmentRequiredOnlyResource = FusionAppsFusionEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentRepresentation)

	FusionAppsFusionEnvironmentResourceConfig = FusionAppsFusionEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentRepresentation)

	FusionAppsFusionAppsFusionEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
	}

	FusionAppsFusionAppsFusionEnvironmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fusion_environment_family_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id}`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: FusionAppsFusionEnvironmentDataSourceFilterRepresentation}}
	FusionAppsFusionEnvironmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`}},
	}

	FusionAppsFusionEnvironmentRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"create_fusion_environment_admin_user_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FusionAppsFusionEnvironmentCreateFusionEnvironmentAdminUserDetailsRepresentation},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"fusion_environment_family_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id}`},
		"fusion_environment_type":      acctest.Representation{RepType: acctest.Required, Create: `TEST`},
		"additional_language_packs":    acctest.Representation{RepType: acctest.Optional, Create: []string{`en`}, Update: []string{`ar`}},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace-terraform.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace-terraform.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns_prefix":                   acctest.Representation{RepType: acctest.Optional, Create: `dnsPrefix`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_fusionapps.keys[0], "id")}`},
		"maintenance_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: FusionAppsFusionEnvironmentMaintenancePolicyRepresentation},
		"rules":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: FusionAppsFusionEnvironmentRulesRepresentation},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreRepresentation_fusionapps},
	}
	FusionAppsFusionEnvironmentCreateFusionEnvironmentAdminUserDetailsRepresentation = map[string]interface{}{
		"email_address": acctest.Representation{RepType: acctest.Required, Create: `JohnSmith@example.com`},
		"first_name":    acctest.Representation{RepType: acctest.Required, Create: `firstName`},
		"last_name":     acctest.Representation{RepType: acctest.Required, Create: `lastName`},
		"password":      acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"username":      acctest.Representation{RepType: acctest.Required, Create: `username_test`},
	}
	FusionAppsFusionEnvironmentMaintenancePolicyRepresentation = map[string]interface{}{
		"environment_maintenance_override": acctest.Representation{RepType: acctest.Optional, Create: `PROD`, Update: `NON_PROD`},
		"monthly_patching_override":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
	}
	FusionAppsFusionEnvironmentRulesRepresentation = map[string]interface{}{
		"action":      acctest.Representation{RepType: acctest.Required, Create: `ALLOW`},
		"conditions":  acctest.RepresentationGroup{RepType: acctest.Required, Group: FusionAppsFusionEnvironmentRulesConditionsRepresentation},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
	}
	FusionAppsFusionEnvironmentRulesConditionsRepresentation = map[string]interface{}{
		"attribute_name":  acctest.Representation{RepType: acctest.Required, Create: `SOURCE_IP_ADDRESS`, Update: `SOURCE_IP_ADDRESS`},
		"attribute_value": acctest.Representation{RepType: acctest.Required, Create: `208.128.0.0/10`, Update: `209.128.0.0/10`},
	}

	FusionAppsFusionEnvironmentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation) +
		DefinedTagsDependencies_fusionapps +
		KeyResourceDependencyConfig_fusionapps
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_fusion_apps_fusion_environment.test_fusion_environment"
	datasourceName := "data.oci_fusion_apps_fusion_environments.test_fusion_environments"
	singularDatasourceName := "data.oci_fusion_apps_fusion_environment.test_fusion_environment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FusionAppsFusionEnvironmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Optional, acctest.Create, FusionAppsFusionEnvironmentRepresentation), "fusionapps", "fusionEnvironment", t)

	acctest.ResourceTest(t, testAccCheckFusionAppsFusionEnvironmentDestroy, []resource.TestStep{

		/**
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.email_address", "JohnSmith@example.com"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.first_name", "firstName"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.last_name", "lastName"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.username", "username_test"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "fusion_environment_family_id"),
				resource.TestCheckResourceAttr(resourceName, "fusion_environment_type", "TEST"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentResourceDependencies,
		},
		**/

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Optional, acctest.Create, FusionAppsFusionEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "additional_language_packs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.email_address", "JohnSmith@example.com"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.first_name", "firstName"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.last_name", "lastName"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.username", "username_test"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_prefix", "dnsPrefix"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fusion_environment_family_id"),
				resource.TestCheckResourceAttr(resourceName, "fusion_environment_type", "TEST"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_policy.0.environment_maintenance_override", "PROD"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_policy.0.monthly_patching_override", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.conditions.0.attribute_name", "SOURCE_IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.conditions.0.attribute_value", "208.128.0.0/10"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FusionAppsFusionEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FusionAppsFusionEnvironmentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "additional_language_packs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.email_address", "JohnSmith@example.com"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.first_name", "firstName"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.last_name", "lastName"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.username", "username_test"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_prefix", "dnsPrefix"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fusion_environment_family_id"),
				resource.TestCheckResourceAttr(resourceName, "fusion_environment_type", "TEST"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_policy.0.environment_maintenance_override", "PROD"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_policy.0.monthly_patching_override", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.conditions.0.attribute_name", "SOURCE_IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.conditions.0.attribute_value", "208.128.0.0/10"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "additional_language_packs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.email_address", "JohnSmith@example.com"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.first_name", "firstName"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.last_name", "lastName"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "create_fusion_environment_admin_user_details.0.username", "username_test"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_prefix", "dnsPrefix"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fusion_environment_family_id"),
				resource.TestCheckResourceAttr(resourceName, "fusion_environment_type", "TEST"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_policy.0.environment_maintenance_override", "NON_PROD"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_policy.0.monthly_patching_override", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.conditions.0.attribute_name", "SOURCE_IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.conditions.0.attribute_value", "209.128.0.0/10"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environments", "test_fusion_environments", acctest.Optional, acctest.Update, FusionAppsFusionAppsFusionEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "fusion_environment_family_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "fusion_environment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fusion_environment_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "additional_language_packs.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "applied_patch_bundles.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_prefix", "dnsPrefix"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "domain_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fusion_environment_type", "TEST"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_domain_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_break_glass_enabled"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kms_key_info.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lockbox_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_policy.0.environment_maintenance_override", "NON_PROD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_policy.0.monthly_patching_override", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_policy.0.quarterly_upgrade_begin_times.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "refresh.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.conditions.0.attribute_name", "SOURCE_IP_ADDRESS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.conditions.0.attribute_value", "209.128.0.0/10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subscription_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_upcoming_maintenance"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// verify resource import
		{
			Config:            config + FusionAppsFusionEnvironmentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"create_fusion_environment_admin_user_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckFusionAppsFusionEnvironmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FusionApplicationsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fusion_apps_fusion_environment" {
			noResourceFound = false
			request := oci_fusion_apps.GetFusionEnvironmentRequest{}

			tmp := rs.Primary.ID
			request.FusionEnvironmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fusion_apps")

			response, err := client.GetFusionEnvironment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fusion_apps.FusionEnvironmentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FusionAppsFusionEnvironment") {
		resource.AddTestSweepers("FusionAppsFusionEnvironment", &resource.Sweeper{
			Name:         "FusionAppsFusionEnvironment",
			Dependencies: acctest.DependencyGraph["fusionEnvironment"],
			F:            sweepFusionAppsFusionEnvironmentResource,
		})
	}
}

func sweepFusionAppsFusionEnvironmentResource(compartment string) error {
	fusionApplicationsClient := acctest.GetTestClients(&schema.ResourceData{}).FusionApplicationsClient()
	fusionEnvironmentIds, err := getFusionAppsFusionEnvironmentIds(compartment)
	if err != nil {
		return err
	}
	for _, fusionEnvironmentId := range fusionEnvironmentIds {
		if ok := acctest.SweeperDefaultResourceId[fusionEnvironmentId]; !ok {
			deleteFusionEnvironmentRequest := oci_fusion_apps.DeleteFusionEnvironmentRequest{}

			deleteFusionEnvironmentRequest.FusionEnvironmentId = &fusionEnvironmentId

			deleteFusionEnvironmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fusion_apps")
			_, error := fusionApplicationsClient.DeleteFusionEnvironment(context.Background(), deleteFusionEnvironmentRequest)
			if error != nil {
				fmt.Printf("Error deleting FusionEnvironment %s %s, It is possible that the resource is already deleted. Please verify manually \n", fusionEnvironmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fusionEnvironmentId, FusionAppsFusionEnvironmentSweepWaitCondition, time.Duration(3*time.Minute),
				FusionAppsFusionEnvironmentSweepResponseFetchOperation, "fusion_apps", true)
		}
	}
	return nil
}

func getFusionAppsFusionEnvironmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FusionEnvironmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fusionApplicationsClient := acctest.GetTestClients(&schema.ResourceData{}).FusionApplicationsClient()

	listFusionEnvironmentsRequest := oci_fusion_apps.ListFusionEnvironmentsRequest{}
	listFusionEnvironmentsRequest.CompartmentId = &compartmentId
	listFusionEnvironmentsRequest.LifecycleState = oci_fusion_apps.FusionEnvironmentLifecycleStateActive
	listFusionEnvironmentsResponse, err := fusionApplicationsClient.ListFusionEnvironments(context.Background(), listFusionEnvironmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting FusionEnvironment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, fusionEnvironment := range listFusionEnvironmentsResponse.Items {
		id := *fusionEnvironment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FusionEnvironmentId", id)
	}
	return resourceIds, nil
}

func FusionAppsFusionEnvironmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fusionEnvironmentResponse, ok := response.Response.(oci_fusion_apps.GetFusionEnvironmentResponse); ok {
		return fusionEnvironmentResponse.LifecycleState != oci_fusion_apps.FusionEnvironmentLifecycleStateDeleted
	}
	return false
}

func FusionAppsFusionEnvironmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FusionApplicationsClient().GetFusionEnvironment(context.Background(), oci_fusion_apps.GetFusionEnvironmentRequest{
		FusionEnvironmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
