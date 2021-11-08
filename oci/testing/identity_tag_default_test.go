// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package testing

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v49/common"
	oci_identity "github.com/oracle/oci-go-sdk/v49/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	TagDefaultRequiredOnlyResource = TagDefaultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Required, Create, tagDefaultRepresentation)

	TagDefaultResourceConfig = TagDefaultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Optional, Update, tagDefaultRepresentation)

	tagDefaultSingularDataSourceRepresentation = map[string]interface{}{
		"tag_default_id": Representation{RepType: Required, Create: `${oci_identity_tag_default.test_tag_default.id}`},
	}

	tagDefaultDataSourceRepresentationWithCompartmentIdFilter = map[string]interface{}{
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{Required, tagDefaultDataSourceFilterRepresentation}}
	tagDefaultDataSourceRepresentationWithIdFilter = map[string]interface{}{
		"id":     Representation{RepType: Optional, Create: `${oci_identity_tag_default.test_tag_default.id}`},
		"filter": acctest.RepresentationGroup{Required, tagDefaultDataSourceFilterRepresentation}}
	tagDefaultDataSourceRepresentationWithStateFilter = map[string]interface{}{
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{Required, tagDefaultDataSourceFilterRepresentation}}
	tagDefaultDataSourceRepresentationWithTagDefinitionIdFilter = map[string]interface{}{
		"tag_definition_id": Representation{RepType: Optional, Create: `${oci_identity_tag.test_tag.id}`},
		"filter":            acctest.RepresentationGroup{Required, tagDefaultDataSourceFilterRepresentation}}
	tagDefaultDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_identity_tag_default.test_tag_default.id}`}},
	}

	tagDefaultRepresentation = map[string]interface{}{
		"compartment_id":    Representation{RepType: Required, Create: `${var.compartment_id}`},
		"tag_definition_id": Representation{RepType: Required, Create: `${oci_identity_tag.test_tag.id}`},
		"value":             Representation{RepType: Required, Create: `value1`, Update: `value2`},
		"is_required":       Representation{RepType: Optional, Create: `true`, Update: `false`},
	}

	TagDefaultResourceDependencies = TagRequiredOnlyResource
)

// issue-routing-tag: identity/default
func TestIdentityTagDefaultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTagDefaultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentIdCreate := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id_for_update", compartmentIdCreate)
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_tag_default.test_tag_default"
	datasourceName := "data.oci_identity_tag_defaults.test_tag_defaults"
	singularDatasourceName := "data.oci_identity_tag_default.test_tag_default"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+TagDefaultResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Optional, Create, tagDefaultRepresentation), "identity", "tagDefault", t)

	acctest.ResourceTest(t, testAccCheckIdentityTagDefaultDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + TagDefaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Required, Create, tagDefaultRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "tag_definition_id"),
				resource.TestCheckResourceAttr(resourceName, "value", "value1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + TagDefaultResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + TagDefaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Optional, Create, tagDefaultRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_required", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "tag_definition_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tag_definition_name"),
				resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "value", "value1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithBlankDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + TagDefaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Optional, Update, tagDefaultRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_required", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "tag_definition_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tag_definition_name"),
				resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "value", "value2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_tag_defaults", "test_tag_defaults_with_compartment_id_filter", Optional, Update, tagDefaultDataSourceRepresentationWithCompartmentIdFilter) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_tag_defaults", "test_tag_defaults_with_id_filter", Optional, Update, tagDefaultDataSourceRepresentationWithIdFilter) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_tag_defaults", "test_tag_defaults_with_state_filter", Optional, Update, tagDefaultDataSourceRepresentationWithStateFilter) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_tag_defaults", "test_tag_defaults_with_tag_definition_id_filter", Optional, Update, tagDefaultDataSourceRepresentationWithTagDefinitionIdFilter) +
				compartmentIdVariableStr + TagDefaultResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Optional, Update, tagDefaultRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName+"_with_compartment_id_filter", "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.#"),
				resource.TestCheckResourceAttr(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.tag_definition_id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.tag_definition_name"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.tag_namespace_id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.value", "value2"),

				resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "id"),

				resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.#"),
				resource.TestCheckResourceAttr(datasourceName+"_with_id_filter", "tag_defaults.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.0.tag_definition_id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.0.tag_definition_name"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.0.tag_namespace_id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName+"_with_id_filter", "tag_defaults.0.value", "value2"),

				resource.TestCheckResourceAttr(datasourceName+"_with_state_filter", "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName+"_with_state_filter", "state", "AVAILABLE"),

				resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.#"),
				resource.TestCheckResourceAttr(datasourceName+"_with_state_filter", "tag_defaults.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.0.tag_definition_id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.0.tag_definition_name"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.0.tag_namespace_id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName+"_with_state_filter", "tag_defaults.0.value", "value2"),

				resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_definition_id"),

				resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.#"),
				resource.TestCheckResourceAttr(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.tag_definition_id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.tag_definition_name"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.tag_namespace_id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.value", "value2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Required, Create, tagDefaultSingularDataSourceRepresentation) +
				compartmentIdVariableStr + TagDefaultResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tag_default_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tag_definition_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_required", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tag_definition_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tag_namespace_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "value", "value2"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + TagDefaultResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIdentityTagDefaultDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_tag_default" {
			noResourceFound = false
			request := oci_identity.GetTagDefaultRequest{}

			tmp := rs.Primary.ID
			request.TagDefaultId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "identity")

			_, err := client.GetTagDefault(context.Background(), request)

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
