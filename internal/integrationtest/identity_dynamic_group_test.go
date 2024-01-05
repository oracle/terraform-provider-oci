// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDynamicGroupRequiredOnlyResource = IdentityDynamicGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", acctest.Required, acctest.Create, IdentityDynamicGroupRepresentation)

	IdentityIdentityDynamicGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `DevCompartmentDynamicGroup`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDynamicGroupDataSourceFilterRepresentation}}
	IdentityDynamicGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_dynamic_group.test_dynamic_group.id}`}},
	}

	IdentityDynamicGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `Instance Group for dev compartment`, Update: `description2`},
		"matching_rule":  acctest.Representation{RepType: acctest.Required, Create: `${var.dynamic_group_matching_rule}`, Update: `${var.dynamic_group_matching_rule}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `DevCompartmentDynamicGroup`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	IdentityDynamicGroupResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: identity/default
func TestIdentityDynamicGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDynamicGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	matchingRuleValueStr := fmt.Sprintf("instance.compartment_id='%s'", compartmentId)
	matchingRuleVariableStr := fmt.Sprintf("variable \"dynamic_group_matching_rule\" {default = \"%s\" }\n", matchingRuleValueStr)

	matchingRule2ValueStr := fmt.Sprintf("instance.compartment_id='%s'", compartmentId)
	matchingRule2VariableStr := fmt.Sprintf("variable \"dynamic_group_matching_rule\" {default = \"%s\" }\n", matchingRule2ValueStr)
	resourceName := "oci_identity_dynamic_group.test_dynamic_group"
	datasourceName := "data.oci_identity_dynamic_groups.test_dynamic_groups"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+matchingRuleVariableStr+IdentityDynamicGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", acctest.Optional, acctest.Create, IdentityDynamicGroupRepresentation), "identity", "dynamicGroup", t)

	acctest.ResourceTest(t, testAccCheckIdentityDynamicGroupDestroy, []resource.TestStep{
		// verify matching rule syntax
		{
			Config: config + `
variable "dynamic_group_description" { default = "description2" }
variable "dynamic_group_matching_rule" { default = "bad_matching_rule" }
variable "dynamic_group_name" { default = "DevCompartmentDynamicGroup" }
` + compartmentIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", acctest.Required, acctest.Create, IdentityDynamicGroupRepresentation),
			ExpectError: regexp.MustCompile("Unable to parse matching rule"),
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + matchingRuleVariableStr + IdentityDynamicGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", acctest.Required, acctest.Create, IdentityDynamicGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Instance Group for dev compartment"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule", matchingRuleValueStr),
				resource.TestCheckResourceAttr(resourceName, "name", "DevCompartmentDynamicGroup"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + matchingRuleVariableStr + IdentityDynamicGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + matchingRuleVariableStr + IdentityDynamicGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", acctest.Optional, acctest.Create, IdentityDynamicGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Instance Group for dev compartment"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule", matchingRuleValueStr),
				resource.TestCheckResourceAttr(resourceName, "name", "DevCompartmentDynamicGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + matchingRule2VariableStr + IdentityDynamicGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", acctest.Optional, acctest.Update, IdentityDynamicGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule", matchingRule2ValueStr),
				resource.TestCheckResourceAttr(resourceName, "name", "DevCompartmentDynamicGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + matchingRule2VariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_dynamic_groups", "test_dynamic_groups", acctest.Optional, acctest.Update, IdentityIdentityDynamicGroupDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDynamicGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", acctest.Optional, acctest.Update, IdentityDynamicGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "DevCompartmentDynamicGroup"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.matching_rule", matchingRule2ValueStr),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.name", "DevCompartmentDynamicGroup"),
				resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + IdentityDynamicGroupRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIdentityDynamicGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_dynamic_group" {
			noResourceFound = false
			request := oci_identity.GetDynamicGroupRequest{}

			tmp := rs.Primary.ID
			request.DynamicGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")

			response, err := client.GetDynamicGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.DynamicGroupLifecycleStateDeleted): true,
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
