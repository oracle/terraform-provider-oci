// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_identity "github.com/oracle/oci-go-sdk/v53/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	GroupRequiredOnlyResource = GroupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create, groupRepresentation)

	groupSingularDataSourceRepresentation = map[string]interface{}{
		"group_id": Representation{RepType: Required, Create: `${oci_identity_group.test_group.id}`},
	}

	groupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"name":           Representation{RepType: Optional, Create: `NetworkAdmins`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, groupDataSourceFilterRepresentation}}
	groupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_identity_group.test_group.id}`}},
	}

	groupRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"description":    Representation{RepType: Required, Create: `Group for network administrators`, Update: `description2`},
		"name":           Representation{RepType: Required, Create: `NetworkAdmins`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	GroupResourceDependencies = DefinedTagsDependencies
	GroupResourceConfig       = GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create, groupRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_group.test_group"
	datasourceName := "data.oci_identity_groups.test_groups"
	singularDatasourceName := "data.oci_identity_group.test_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+GroupResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Create, groupRepresentation), "identity", "Group", t)

	ResourceTest(t, testAccCheckIdentityGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create, groupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Group for network administrators"),
				resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Create, groupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "Group for network administrators"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + GroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Update, groupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_identity_groups", "test_groups", Optional, Update, groupDataSourceRepresentation) +
				compartmentIdVariableStr + GroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Update, groupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "NetworkAdmins"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "groups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "groups.0.compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "groups.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "groups.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "groups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "groups.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "groups.0.name", "NetworkAdmins"),
				resource.TestCheckResourceAttrSet(datasourceName, "groups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "groups.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create, groupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Update, groupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "NetworkAdmins"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + GroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Update, groupRepresentation),
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

func testAccCheckIdentityGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_group" {
			noResourceFound = false
			request := oci_identity.GetGroupRequest{}

			tmp := rs.Primary.ID
			request.GroupId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "identity")

			response, err := client.GetGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.GroupLifecycleStateDeleted): true,
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
