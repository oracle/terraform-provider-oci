// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsMappedAttributeRequiredOnlyResource = IdentityDomainsMappedAttributeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_mapped_attribute", "test_mapped_attribute", acctest.Required, acctest.Create, IdentityDomainsMappedAttributeRepresentation)

	IdentityDomainsMappedAttributeResourceConfig = IdentityDomainsMappedAttributeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_mapped_attribute", "test_mapped_attribute", acctest.Optional, acctest.Update, IdentityDomainsMappedAttributeRepresentation)

	IdentityDomainsMappedAttributeSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"mapped_attribute_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_mapped_attributes.test_mapped_attributes.mapped_attributes.0.id}`},
		"attribute_sets":      acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsMappedAttributeDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"mapped_attribute_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mapped_attribute_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":          acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":             acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMappedAttributeRepresentation = map[string]interface{}{
		"direction":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_mapped_attributes.test_mapped_attributes.mapped_attributes.0.direction}`},
		"idcs_endpoint":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"idcs_resource_type":  acctest.Representation{RepType: acctest.Required, Create: `User`},
		"mapped_attribute_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_mapped_attributes.test_mapped_attributes.mapped_attributes.0.id}`},
		"ref_resource_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_mapped_attributes.test_mapped_attributes.mapped_attributes.0.ref_resource_id}`},
		"ref_resource_type":   acctest.Representation{RepType: acctest.Required, Create: `App`},
		"schemas":             acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:MappedAttribute`}},
		"attribute_mappings":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsMappedAttributeAttributeMappingsRepresentation},
		"attribute_sets":      acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"tags":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsMappedAttributeTagsRepresentation},
	}
	IdentityDomainsMappedAttributeAttributeMappingsRepresentation = map[string]interface{}{
		"idcs_attribute_name":           acctest.Representation{RepType: acctest.Required, Create: `$(user.userName)`, Update: `$(user.name.givenName)`},
		"managed_object_attribute_name": acctest.Representation{RepType: acctest.Required, Create: `name`},
		"applies_to_actions":            acctest.Representation{RepType: acctest.Optional, Create: []string{`create`}, Update: []string{`update`}},
	}
	IdentityDomainsMappedAttributeTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsMappedAttributeResourceDependencies = TestDomainDependencies + MappedAttributeRefResourceAppDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMappedAttributeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMappedAttributeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_mapped_attribute.test_mapped_attribute"
	datasourceName := "data.oci_identity_domains_mapped_attributes.test_mapped_attributes"
	singularDatasourceName := "data.oci_identity_domains_mapped_attribute.test_mapped_attribute"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsMappedAttributeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_mapped_attribute", "test_mapped_attribute", acctest.Optional, acctest.Create, IdentityDomainsMappedAttributeRepresentation), "identitydomains", "mappedAttribute", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Prerequisite: MappedAttributes cannot be created. It's auto seeded when there exists e.g. SAML app.
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMappedAttributeResourceDependencies,
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMappedAttributeResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_mapped_attributes", "test_mapped_attributes", acctest.Required, acctest.Create, IdentityDomainsMappedAttributeDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_mapped_attribute", "test_mapped_attribute", acctest.Optional, acctest.Create, IdentityDomainsMappedAttributeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "direction", "outbound"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idcs_resource_type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "mapped_attribute_id"),
				resource.TestCheckResourceAttr(resourceName, "ref_resource_type", "App"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMappedAttributeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMappedAttributeResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_mapped_attributes", "test_mapped_attributes", acctest.Required, acctest.Create, IdentityDomainsMappedAttributeDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_mapped_attribute", "test_mapped_attribute", acctest.Optional, acctest.Create, IdentityDomainsMappedAttributeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_mappings.0.applies_to_actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_mappings.0.idcs_attribute_name", "$(user.userName)"),
				resource.TestCheckResourceAttr(resourceName, "attribute_mappings.0.managed_object_attribute_name", "name"),
				resource.TestCheckResourceAttr(resourceName, "attribute_mappings.0.required", "false"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "direction", "outbound"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idcs_resource_type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "mapped_attribute_id"),
				resource.TestCheckResourceAttr(resourceName, "ref_resource_type", "App"),
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

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "mappedAttributes", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsMappedAttributeResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_mapped_attributes", "test_mapped_attributes", acctest.Required, acctest.Create, IdentityDomainsMappedAttributeDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_mapped_attribute", "test_mapped_attribute", acctest.Optional, acctest.Update, IdentityDomainsMappedAttributeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_mappings.0.applies_to_actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_mappings.0.idcs_attribute_name", "$(user.name.givenName)"),
				resource.TestCheckResourceAttr(resourceName, "attribute_mappings.0.managed_object_attribute_name", "name"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "direction", "outbound"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idcs_resource_type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "mapped_attribute_id"),
				resource.TestCheckResourceAttr(resourceName, "ref_resource_type", "App"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_mapped_attributes", "test_mapped_attributes", acctest.Optional, acctest.Update, IdentityDomainsMappedAttributeDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMappedAttributeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_mapped_attribute", "test_mapped_attribute", acctest.Optional, acctest.Update, IdentityDomainsMappedAttributeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "mapped_attribute_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),
				resource.TestCheckResourceAttr(datasourceName, "mapped_attributes.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_mapped_attributes", "test_mapped_attributes", acctest.Required, acctest.Create, IdentityDomainsMappedAttributeDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_mapped_attribute", "test_mapped_attribute", acctest.Required, acctest.Create, IdentityDomainsMappedAttributeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMappedAttributeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mapped_attribute_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "direction", "outbound"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_resource_type", "User"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ref_resource_type", "App"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsMappedAttributeRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_mapped_attribute", "mappedAttributes"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"tags",
				"mapped_attribute_id",
			},
			ResourceName: resourceName,
		},
	})
}
