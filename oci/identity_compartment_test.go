// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CompartmentRequiredOnlyResource = CompartmentResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", Required, Create, compartmentRepresentation)

	CompartmentResourceConfig = CompartmentResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", Optional, Update, compartmentRepresentation)

	compartmentSingularDataSourceRepresentation = map[string]interface{}{
		"id": Representation{repType: Required, create: `${oci_identity_compartment.test_compartment.id}`},
	}

	compartmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"access_level":              Representation{repType: Optional, create: `ANY`},
		"compartment_id_in_subtree": Representation{repType: Optional, create: `false`},
		"filter":                    RepresentationGroup{Required, compartmentDataSourceFilterRepresentation}}
	compartmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_compartment.test_compartment.id}`}},
	}

	compartmentRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"description":    Representation{repType: Required, create: `For network components`, update: `description2`},
		"name":           Representation{repType: Required, create: `Network`, update: `name2`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	CompartmentResourceDependencies = DefinedTagsDependencies
)

func TestIdentityCompartmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityCompartmentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_compartment.test_compartment"
	datasourceName := "data.oci_identity_compartments.test_compartments"
	singularDatasourceName := "data.oci_identity_compartment.test_compartment"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityCompartmentDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + CompartmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", Required, Create, compartmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "For network components"),
					resource.TestCheckResourceAttr(resourceName, "name", "Network"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CompartmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", Optional, Create, compartmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "For network components"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "Network"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CompartmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", Optional, Update, compartmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_identity_compartments", "test_compartments", Optional, Update, compartmentDataSourceRepresentation) +
					compartmentIdVariableStr + CompartmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", Optional, Update, compartmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "access_level", "ANY"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),

					resource.TestCheckResourceAttr(datasourceName, "compartments.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "compartments.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "compartments.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "compartments.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "compartments.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "compartments.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "compartments.0.name", "name2"),
					resource.TestCheckResourceAttrSet(datasourceName, "compartments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "compartments.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_compartment", "test_compartment", Required, Create, compartmentSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CompartmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + CompartmentResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enable_delete",
				},
				ResourceName: resourceName,
			},
			// restore name of compartment
			{
				Config: config + compartmentIdVariableStr + CompartmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", Required, Create, compartmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Network"),
				),
			},
			// verify error on existing compartment name where automatic import fails due to enable_delete
			{
				Config: config + compartmentIdVariableStr + CompartmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", Required, Create, compartmentRepresentation) +
					generateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment2", Required, Create,
						representationCopyWithNewProperties(compartmentRepresentation, map[string]interface{}{
							"enable_delete": Representation{repType: Required, create: `true`}})),
				ExpectError: regexp.MustCompile("If you intended to manage an existing compartment, use terraform import instead."),
			},
		},
	})
}

func testAccCheckIdentityCompartmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_compartment" {
			noResourceFound = false
			request := oci_identity.GetCompartmentRequest{}

			tmp := rs.Primary.ID
			request.CompartmentId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")

			response, err := client.GetCompartment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.CompartmentLifecycleStateDeleted): true, // target state when delete_enabled = true
					string(oci_identity.CompartmentLifecycleStateActive):  true, // target state when delete_enabled = false or ""
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
