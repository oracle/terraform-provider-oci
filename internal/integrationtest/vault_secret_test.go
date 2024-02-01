// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	VaultSecretRequiredOnlyResource = VaultSecretResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation)

	VaultSecretResourceConfig = VaultSecretResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Optional, acctest.Update, VaultSecretRepresentation)

	VaultVaultSecretSingularDataSourceRepresentation = map[string]interface{}{
		"secret_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_secret.id}`},
	}
	secretName  = utils.RandomString(10, utils.CharsetWithoutDigits)
	secretName2 = utils.RandomString(10, utils.CharsetWithoutDigits)

	VaultVaultSecretDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: secretName2},
		"vault_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: VaultSecretDataSourceFilterRepresentation}}
	VaultSecretDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_vault_secret.test_secret.id}`}},
	}

	VaultSecretRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"secret_content": acctest.RepresentationGroup{RepType: acctest.Required, Group: VaultSecretSecretContentRepresentation},
		"secret_name":    acctest.Representation{RepType: acctest.Required, Create: secretName},
		"vault_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.vault_id}`},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `my test secret`, Update: `description2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"key_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.key_id}`},
		"metadata":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"metadata": "metadata"}, Update: map[string]string{"metadata2": "metadata2"}},
		"secret_rules":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: VaultSecretSecretRulesRepresentation},
		"rotation_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: VaultSecretRotationConfigRepresentation},
	}
	VaultSecretRotationConfigRepresentation = map[string]interface{}{
		"target_system_details":         acctest.RepresentationGroup{RepType: acctest.Required, Group: VaultSecretRotationConfigTargetSystemDetailsRepresentation},
		"is_scheduled_rotation_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"rotation_interval":             acctest.Representation{RepType: acctest.Optional, Create: `P30D`, Update: `P90D`},
	}
	VaultSecretSecretContentRepresentation = map[string]interface{}{
		"content_type": acctest.Representation{RepType: acctest.Required, Create: `BASE64`},
		"content":      acctest.Representation{RepType: acctest.Required, Create: `PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg==`},
		"name":         acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"stage":        acctest.Representation{RepType: acctest.Optional, Create: `CURRENT`},
	}
	VaultSecretSecretRulesRepresentation = map[string]interface{}{
		"rule_type":                                     acctest.Representation{RepType: acctest.Required, Create: `SECRET_EXPIRY_RULE`, Update: `SECRET_REUSE_RULE`},
		"is_enforced_on_deleted_secret_versions":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_secret_content_retrieval_blocked_on_expiry": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"secret_version_expiry_interval":                acctest.Representation{RepType: acctest.Optional, Create: `P3D`},
		"time_of_absolute_expiry":                       acctest.Representation{RepType: acctest.Optional, Create: deletionTime.Format(time.RFC3339)},
	}
	VaultSecretRotationConfigTargetSystemDetailsRepresentation = map[string]interface{}{
		"target_system_type": acctest.Representation{RepType: acctest.Required, Create: `ADB`, Update: `FUNCTION`},
		"adb_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.adb_id}`},
		"function_id":        acctest.Representation{RepType: acctest.Optional, Update: `${var.function_id}`},
	}

	VaultSecretResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: vault/default
func TestVaultSecretResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestVaultSecretResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	vaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_ocid")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	keyId := utils.GetEnvSettingWithBlankDefault("kms_key_ocid")
	keyIdVariableStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)

	adbId := utils.GetEnvSettingWithBlankDefault("adb_ocid")
	adbIdVariableStr := fmt.Sprintf("variable \"adb_id\" { default = \"%s\" }\n", adbId)

	functionId := utils.GetEnvSettingWithBlankDefault("function_ocid")
	functionIdVariableStr := fmt.Sprintf("variable \"function_id\" { default = \"%s\" }\n", functionId)

	resourceName := "oci_vault_secret.test_secret"
	datasourceName := "data.oci_vault_secrets.test_secrets"
	singularDatasourceName := "data.oci_vault_secret.test_secret"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+vaultIdVariableStr+keyIdVariableStr+VaultSecretResourceDependencies+adbIdVariableStr+functionIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(VaultSecretRepresentation, map[string]interface{}{
			"secret_name": acctest.Representation{RepType: acctest.Required, Create: secretName2},
		})), "vault", "secret", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + VaultSecretResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "secret_name"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + VaultSecretResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + VaultSecretResourceDependencies + adbIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(VaultSecretRepresentation, map[string]interface{}{
					"secret_name": acctest.Representation{RepType: acctest.Required, Create: secretName2},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "my test secret"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.is_scheduled_rotation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.rotation_interval", "P30D"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.target_system_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "rotation_config.0.target_system_details.0.adb_id"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.target_system_details.0.target_system_type", "ADB"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.content", "PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg=="),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.content_type", "BASE64"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.stage", "CURRENT"),
				resource.TestCheckResourceAttrSet(resourceName, "secret_name"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.is_enforced_on_deleted_secret_versions", "false"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.is_secret_content_retrieval_blocked_on_expiry", "false"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.rule_type", "SECRET_EXPIRY_RULE"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.secret_version_expiry_interval", "P3D"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.time_of_absolute_expiry", deletionTime.Format(time.RFC3339)),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + vaultIdVariableStr + keyIdVariableStr + VaultSecretResourceDependencies + adbIdVariableStr + functionIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(VaultSecretRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"secret_name":    acctest.Representation{RepType: acctest.Required, Create: secretName2},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "my test secret"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.is_scheduled_rotation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.rotation_interval", "P30D"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.target_system_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "rotation_config.0.target_system_details.0.adb_id"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.target_system_details.0.target_system_type", "ADB"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.content", "PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg=="),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.content_type", "BASE64"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.stage", "CURRENT"),
				resource.TestCheckResourceAttrSet(resourceName, "secret_name"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.is_enforced_on_deleted_secret_versions", "false"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.is_secret_content_retrieval_blocked_on_expiry", "false"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.rule_type", "SECRET_EXPIRY_RULE"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.secret_version_expiry_interval", "P3D"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.time_of_absolute_expiry", deletionTime.Format(time.RFC3339)),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
			Config: config + compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + VaultSecretResourceDependencies + adbIdVariableStr + functionIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Optional, acctest.Update,
					acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
						"secret_rules.secret_version_expiry_interval",
						"secret_rules.time_of_absolute_expiry",
					}, acctest.RepresentationCopyWithNewProperties(VaultSecretRepresentation, map[string]interface{}{
						"secret_name": acctest.Representation{RepType: acctest.Required, Create: secretName2},
					}))),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.is_scheduled_rotation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.rotation_interval", "P90D"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.target_system_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "rotation_config.0.target_system_details.0.function_id"),
				resource.TestCheckResourceAttr(resourceName, "rotation_config.0.target_system_details.0.target_system_type", "FUNCTION"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.content", "PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg=="),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.content_type", "BASE64"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "secret_content.0.stage", "CURRENT"),
				resource.TestCheckResourceAttrSet(resourceName, "secret_name"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.is_enforced_on_deleted_secret_versions", "true"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.is_secret_content_retrieval_blocked_on_expiry", "false"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.rule_type", "SECRET_REUSE_RULE"),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.secret_version_expiry_interval", ""),
				resource.TestCheckResourceAttr(resourceName, "secret_rules.0.time_of_absolute_expiry", ""),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_vault_secrets", "test_secrets", acctest.Optional, acctest.Update, VaultVaultSecretDataSourceRepresentation) +
				vaultIdVariableStr + keyIdVariableStr + VaultSecretResourceDependencies + compartmentIdVariableStr + adbIdVariableStr + functionIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Optional, acctest.Update, acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
					"secret_rules.secret_version_expiry_interval",
					"secret_rules.time_of_absolute_expiry",
				}, acctest.RepresentationCopyWithNewProperties(VaultSecretRepresentation, map[string]interface{}{
					"secret_name": acctest.Representation{RepType: acctest.Required, Create: secretName2},
				}))),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", secretName2),
				resource.TestCheckResourceAttrSet(datasourceName, "vault_id"),

				resource.TestCheckResourceAttr(datasourceName, "secrets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "secrets.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "secrets.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "secrets.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.last_rotation_time"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.next_rotation_time"),
				resource.TestCheckResourceAttr(datasourceName, "secrets.0.rotation_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "secrets.0.rotation_config.0.is_scheduled_rotation_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "secrets.0.rotation_config.0.rotation_interval", "P90D"),
				resource.TestCheckResourceAttr(datasourceName, "secrets.0.rotation_config.0.target_system_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.rotation_config.0.target_system_details.0.function_id"),
				resource.TestCheckResourceAttr(datasourceName, "secrets.0.rotation_config.0.target_system_details.0.target_system_type", "FUNCTION"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.rotation_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.secret_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.vault_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultVaultSecretSingularDataSourceRepresentation) +
				compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr + VaultSecretResourceDependencies + functionIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Optional, acctest.Update, acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
					"secret_rules.secret_version_expiry_interval",
					"secret_rules.time_of_absolute_expiry",
				}, acctest.RepresentationCopyWithNewProperties(VaultSecretRepresentation, map[string]interface{}{
					"secret_name": acctest.Representation{RepType: acctest.Required, Create: secretName2},
				}))),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "secret_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_version_number"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_rotation_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "next_rotation_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rotation_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rotation_config.0.is_scheduled_rotation_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rotation_config.0.rotation_interval", "P90D"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rotation_config.0.target_system_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rotation_config.0.target_system_details.0.target_system_type", "FUNCTION"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "rotation_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "secret_rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "secret_rules.0.is_enforced_on_deleted_secret_versions", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "secret_rules.0.is_secret_content_retrieval_blocked_on_expiry", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "secret_rules.0.rule_type", "SECRET_REUSE_RULE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "secret_rules.0.secret_version_expiry_interval", ""),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_current_version_expiry"),
			),
		},
		// verify resource import
		{
			Config:            config + VaultSecretRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"secret_content",
			},
			ResourceName: resourceName,
		},
	})
}
