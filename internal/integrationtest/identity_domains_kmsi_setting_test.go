// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsKmsiSettingRequiredOnlyResource = IdentityDomainsKmsiSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_kmsi_setting", "test_kmsi_setting", acctest.Required, acctest.Create, IdentityDomainsKmsiSettingRepresentation)

	IdentityDomainsKmsiSettingResourceConfig = IdentityDomainsKmsiSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_kmsi_setting", "test_kmsi_setting", acctest.Optional, acctest.Update, IdentityDomainsKmsiSettingRepresentation)

	IdentityDomainsIdentityDomainsKmsiSettingSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"kmsi_setting_id": acctest.Representation{RepType: acctest.Required, Create: `KmsiSettings`},
		"attribute_sets":  acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsKmsiSettingDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsKmsiSettingRepresentation = map[string]interface{}{
		"idcs_endpoint":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"kmsi_setting_id":            acctest.Representation{RepType: acctest.Required, Create: `KmsiSettings`},
		"schemas":                    acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:KmsiSettings`}},
		"attribute_sets":             acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"external_id":                acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"kmsi_feature_enabled":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"kmsi_prompt_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"last_used_validity_in_days": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"max_allowed_sessions":       acctest.Representation{RepType: acctest.Required, Create: `9`, Update: `10`},
		"tags":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsKmsiSettingTagsRepresentation},
		"token_validity_in_days":     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"tou_prompt_disabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsKmsiSettingTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsKmsiSettingResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsKmsiSettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsKmsiSettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_kmsi_setting.test_kmsi_setting"
	datasourceName := "data.oci_identity_domains_kmsi_settings.test_kmsi_settings"
	singularDatasourceName := "data.oci_identity_domains_kmsi_setting.test_kmsi_setting"

	c := config + compartmentIdVariableStr + IdentityDomainsKmsiSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_kmsi_setting", "test_kmsi_setting", acctest.Required, acctest.Create, IdentityDomainsKmsiSettingRepresentation)
	println(c)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsKmsiSettingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_kmsi_setting", "test_kmsi_setting", acctest.Optional, acctest.Create, IdentityDomainsKmsiSettingRepresentation), "identitydomains", "kmsiSetting", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create (Create with PUT)
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsKmsiSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_kmsi_setting", "test_kmsi_setting", acctest.Required, acctest.Create, IdentityDomainsKmsiSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "kmsi_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsKmsiSettingResourceDependencies,
		},
		// verify Create with optionals (Create with PUT)
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsKmsiSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_kmsi_setting", "test_kmsi_setting", acctest.Optional, acctest.Create, IdentityDomainsKmsiSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "id", "KmsiSettings"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "kmsi_feature_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "kmsi_prompt_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "kmsi_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "last_used_validity_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "max_allowed_sessions", "9"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "token_validity_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "tou_prompt_disabled", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "kmsiSettings", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsKmsiSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_kmsi_setting", "test_kmsi_setting", acctest.Optional, acctest.Update, IdentityDomainsKmsiSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "id", "KmsiSettings"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "kmsi_feature_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "kmsi_prompt_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "kmsi_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "last_used_validity_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "max_allowed_sessions", "10"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "token_validity_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "tou_prompt_disabled", "true"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_kmsi_settings", "test_kmsi_settings", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsKmsiSettingDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsKmsiSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_kmsi_setting", "test_kmsi_setting", acctest.Optional, acctest.Update, IdentityDomainsKmsiSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "kmsi_settings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "kmsi_settings.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_kmsi_setting", "test_kmsi_setting", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsKmsiSettingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsKmsiSettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kmsi_setting_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "id", "KmsiSettings"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kmsi_feature_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kmsi_prompt_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "last_used_validity_in_days", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_allowed_sessions", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tags.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "token_validity_in_days", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tou_prompt_disabled", "true"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsKmsiSettingRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_kmsi_setting", "kmsiSettings"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"resource_type_schema_version",
				"idcs_last_upgraded_in_release",
				"tags",
				"kmsi_setting_id",
			},
			ResourceName: resourceName,
		},
	})
}
