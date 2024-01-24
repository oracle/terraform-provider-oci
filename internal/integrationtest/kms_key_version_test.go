// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	KmsKeyVersionRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation)

	KmsKeyVersionResourceConfig = KmsKeyVersionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation)

	KmsExternalKeyVersionResourceConfig = KmsKeyVersionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_ext_key_version", acctest.Required, acctest.Create, KmsExternalKeyVersionRepresentation)

	KmsKmsKeyVersionSingularDataSourceRepresentation = map[string]interface{}{
		"key_id":              acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"key_version_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key_version.test_key_version.key_version_id}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
	}

	KmsKmsExternalKeyVersionSingularDataSourceRepresentation = map[string]interface{}{
		"key_id":              acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_ext_keys_dependency.keys[0], "id")}`},
		"key_version_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key_version.test_ext_key_version.key_version_id}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_ext_vault.management_endpoint}`},
	}

	KmsKmsKeyVersionDataSourceRepresentation = map[string]interface{}{
		"key_id":              acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsKeyVersionDataSourceFilterRepresentation}}
	KmsKeyVersionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key_version_id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_kms_key_version.test_key_version.key_version_id}`}},
	}

	KmsKmsExternalKeyVersionDataSourceRepresentation = map[string]interface{}{
		"key_id":              acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_ext_keys_dependency.keys[0], "id")}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_ext_vault.management_endpoint}`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsExternalKeyVersionDataSourceFilterRepresentation}}
	KmsExternalKeyVersionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key_version_id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_kms_key_version.test_ext_key_version.key_version_id}`}},
	}

	keyVersionDeletionTime = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)

	keyId            = utils.GetEnvSettingWithBlankDefault("key_id")
	keyIdVariableStr = fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)

	KmsKeyVersionRepresentation = map[string]interface{}{
		"key_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.key_id}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"time_of_deletion":    acctest.Representation{RepType: acctest.Required, Create: keyVersionDeletionTime.Format(time.RFC3339Nano)},
	}

	KmsExternalKeyVersionRepresentation = map[string]interface{}{
		"key_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.ext_key_id}`},
		"management_endpoint":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_ext_vault.management_endpoint}`},
		"time_of_deletion":        acctest.Representation{RepType: acctest.Required, Create: keyVersionDeletionTime.Format(time.RFC3339Nano)},
		"external_key_version_id": acctest.Representation{RepType: acctest.Required, Create: `204026f5-9798-46da-9b9e-0c99689b599c`},
	}

	//KmsKeyVersionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation) +
	//	KeyResourceDependencyConfig
	KmsKeyVersionResourceDependencies = KeyResourceDependencyConfig
)

func TestKmsExternalKeyVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsExternalKeyVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	os.Setenv("disable_kms_version_delete", "true")

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	extKeyId := utils.GetEnvSettingWithBlankDefault("ext_key_id")
	extKeyIdVariableStr := fmt.Sprintf("variable \"ext_key_id\" { default = \"%s\" }\n", extKeyId)

	resourceName := "oci_kms_key_version.test_ext_key_version"
	datasourceName := "data.oci_kms_key_versions.test_ext_key_versions"
	singularDatasourceName := "data.oci_kms_key_version.test_ext_key_version"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+KmsKeyVersionResourceDependencies+extKeyIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_ext_key_version", acctest.Optional, acctest.Create, KmsExternalKeyVersionRepresentation), "keymanagement", "keyVersion", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		{
			Config: config + compartmentIdVariableStr + KmsKeyVersionResourceDependencies + extKeyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_ext_key_version", acctest.Required, acctest.Create, KmsExternalKeyVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "external_key_version_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + KmsKeyVersionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + KmsKeyVersionResourceDependencies + extKeyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_ext_key_version", acctest.Optional, acctest.Create, KmsExternalKeyVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_key_version_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_key_versions", "test_ext_key_versions", acctest.Optional, acctest.Update, KmsKmsExternalKeyVersionDataSourceRepresentation) +
				compartmentIdVariableStr + KmsKeyVersionResourceDependencies + extKeyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_ext_key_version", acctest.Optional, acctest.Update, KmsExternalKeyVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_endpoint"),

				resource.TestCheckResourceAttr(datasourceName, "key_versions.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.compartment_id"),
				//resource.TestCheckResourceAttr(datasourceName, "key_versions.0.external_key_reference_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.key_version_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_key_version_id"),
			),
		},

		// verify singular datasource
		{
			Config: config + extKeyIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_key_version", "test_ext_key_version", acctest.Required, acctest.Create, KmsKmsExternalKeyVersionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + KmsExternalKeyVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_version_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replica_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_key_version_id"),
			),
		},
	})
}

// issue-routing-tag: kms/default
func TestKmsKeyVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsKeyVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	os.Setenv("disable_kms_version_delete", "true")

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	//tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_kms_key_version.test_key_version"
	datasourceName := "data.oci_kms_key_versions.test_key_versions"
	singularDatasourceName := "data.oci_kms_key_version.test_key_version"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+KmsKeyVersionResourceDependencies+keyIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Optional, acctest.Create, KmsKeyVersionRepresentation), "keymanagement", "keyVersion", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + KmsKeyVersionResourceDependencies + keyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + KmsKeyVersionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + KmsKeyVersionResourceDependencies + keyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Optional, acctest.Create, KmsKeyVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_key_versions", "test_key_versions", acctest.Optional, acctest.Update, KmsKmsKeyVersionDataSourceRepresentation) +
				compartmentIdVariableStr + KmsKeyVersionResourceDependencies + keyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Optional, acctest.Update, KmsKeyVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_endpoint"),

				resource.TestCheckResourceAttr(datasourceName, "key_versions.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.compartment_id"),
				//resource.TestCheckResourceAttr(datasourceName, "key_versions.0.external_key_reference_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.key_version_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.vault_id"),
			),
		},
		// verify singular datasource
		{
			Config: config + keyIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKmsKeyVersionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + KmsKeyVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_version_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replica_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),
			),
		},
		// verify resource import
		{
			Config:            config + KmsKeyVersionRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: keyVersionImportId,
			ImportStateVerifyIgnore: []string{
				//"external_key_version_id",
				"management_endpoint",
				"time_of_deletion",
				"replica_details",
			},
			ResourceName: resourceName,
		},
	})
}

func keyVersionImportId(state *terraform.State) (string, error) {
	for _, rs := range state.RootModule().Resources {
		if rs.Type == "oci_kms_key_version" {
			return fmt.Sprintf("managementEndpoint/%s/%s", rs.Primary.Attributes["management_endpoint"], rs.Primary.ID), nil
		}
	}

	return "", fmt.Errorf("unable to Create import id as no resource of type oci_kms_key_version in state")
}
