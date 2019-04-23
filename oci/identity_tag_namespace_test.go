// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	TagNamespaceRequiredOnlyResource = TagNamespaceResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Required, Create, tagNamespaceRepresentation)

	tagNamespaceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"include_subcompartments": Representation{repType: Optional, create: `false`},
		"filter":                  RepresentationGroup{Required, tagNamespaceDataSourceFilterRepresentation}}
	tagNamespaceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_tag_namespace.test_tag_namespace.id}`}},
	}

	tagNamespaceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"description":    Representation{repType: Required, create: `This namespace contains tags that will be used in billing.`, update: `description2`},
		"name":           Representation{repType: Required, create: `BillingTags`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	TagNamespaceResourceDependencies = DefinedTagsDependencies
)

func TestIdentityTagNamespaceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTagNamespaceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_identity_tag_namespace.test_tag_namespace"
	datasourceName := "data.oci_identity_tag_namespaces.test_tag_namespaces"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + TagNamespaceResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Required, Create, tagNamespaceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "This namespace contains tags that will be used in billing."),
					resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + TagNamespaceResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + TagNamespaceResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Optional, Create, tagNamespaceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "This namespace contains tags that will be used in billing."),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_retired"),
					resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + TagNamespaceResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Optional, Create,
						representationCopyWithNewProperties(tagNamespaceRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "This namespace contains tags that will be used in billing."),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_retired"),
					resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + TagNamespaceResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Optional, Update, tagNamespaceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_retired"),
					resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),
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
					generateDataSourceFromRepresentationMap("oci_identity_tag_namespaces", "test_tag_namespaces", Optional, Update, tagNamespaceDataSourceRepresentation) +
					compartmentIdVariableStr + TagNamespaceResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Optional, Update, tagNamespaceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "include_subcompartments", "false"),

					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "tag_namespaces.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "tag_namespaces.0.is_retired"),
					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.name", "BillingTags"),
					resource.TestCheckResourceAttrSet(datasourceName, "tag_namespaces.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
