// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsAccountRecoverySettingRequiredOnlyResource = IdentityDomainsAccountRecoverySettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_account_recovery_setting", "test_account_recovery_setting", acctest.Required, acctest.Create, IdentityDomainsAccountRecoverySettingRepresentation)

	IdentityDomainsAccountRecoverySettingResourceConfig = IdentityDomainsAccountRecoverySettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_account_recovery_setting", "test_account_recovery_setting", acctest.Optional, acctest.Update, IdentityDomainsAccountRecoverySettingRepresentation)

	IdentityDomainsIdentityDomainsAccountRecoverySettingSingularDataSourceRepresentation = map[string]interface{}{
		"account_recovery_setting_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_account_recovery_setting.test_account_recovery_setting.id}`},
		"idcs_endpoint":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets":              acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsAccountRecoverySettingDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsAccountRecoverySettingRepresentation = map[string]interface{}{
		"account_recovery_setting_id": acctest.Representation{RepType: acctest.Required, Create: `AccountRecoverySettings`}, // AccountRecoverySettings
		"factors":                     acctest.Representation{RepType: acctest.Required, Create: []string{`email`}, Update: []string{`secquestions`}},
		"idcs_endpoint":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"lockout_duration":            acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"max_incorrect_attempts":      acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"schemas":                     acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:AccountRecoverySettings`}},
		"attribute_sets":              acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"external_id":                 acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		//"id":                          acctest.Representation{RepType: acctest.Optional, Create: `id`, Update: `id2`},
		//"ocid":                        acctest.Representation{RepType: acctest.Optional, Create: `ocid`, Update: `ocid2`},
		"tags": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAccountRecoverySettingTagsRepresentation},
	}
	IdentityDomainsAccountRecoverySettingTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsAccountRecoverySettingResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsAccountRecoverySettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsAccountRecoverySettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_account_recovery_setting.test_account_recovery_setting"
	datasourceName := "data.oci_identity_domains_account_recovery_settings.test_account_recovery_settings"
	singularDatasourceName := "data.oci_identity_domains_account_recovery_setting.test_account_recovery_setting"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsAccountRecoverySettingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_account_recovery_setting", "test_account_recovery_setting", acctest.Optional, acctest.Create, IdentityDomainsAccountRecoverySettingRepresentation), "identitydomains", "accountRecoverySetting", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAccountRecoverySettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_account_recovery_setting", "test_account_recovery_setting", acctest.Required, acctest.Create, IdentityDomainsAccountRecoverySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "account_recovery_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "factors.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "lockout_duration", "10"),
				resource.TestCheckResourceAttr(resourceName, "max_incorrect_attempts", "10"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAccountRecoverySettingResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAccountRecoverySettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_account_recovery_setting", "test_account_recovery_setting", acctest.Optional, acctest.Create, IdentityDomainsAccountRecoverySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "account_recovery_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "factors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "id", "AccountRecoverySettings"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "lockout_duration", "10"),
				resource.TestCheckResourceAttr(resourceName, "max_incorrect_attempts", "10"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "accountRecoverySettings", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAccountRecoverySettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_account_recovery_setting", "test_account_recovery_setting", acctest.Optional, acctest.Update, IdentityDomainsAccountRecoverySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "account_recovery_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "factors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "id", "AccountRecoverySettings"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "lockout_duration", "11"),
				resource.TestCheckResourceAttr(resourceName, "max_incorrect_attempts", "11"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_account_recovery_settings", "test_account_recovery_settings", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsAccountRecoverySettingDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAccountRecoverySettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_account_recovery_setting", "test_account_recovery_setting", acctest.Optional, acctest.Update, IdentityDomainsAccountRecoverySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "account_recovery_settings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "account_recovery_settings.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_account_recovery_setting", "test_account_recovery_setting", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsAccountRecoverySettingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAccountRecoverySettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "account_recovery_setting_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "factors.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "id", "AccountRecoverySettings"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lockout_duration", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_incorrect_attempts", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsAccountRecoverySettingRequiredOnlyResource,
			ImportState:       true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_account_recovery_setting", "accountRecoverySettings"),
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"idcs_last_upgraded_in_release",
				"idcs_prevented_operations",
				"account_recovery_setting_id",
				"tags",
			},
			ResourceName: resourceName,
		},
	})
}
