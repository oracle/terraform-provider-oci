// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	DefinedTagsDependencies = `
variable defined_tag_namespace_name { default = "" }
resource "oci_identity_tag_namespace" "tag-namespace1" {
  		#Required
		compartment_id = "${var.tenancy_ocid}"
  		description = "example tag namespace"
  		name = "${var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"}"

		is_retired = false
}

resource "oci_identity_tag" "tag1" {
  		#Required
  		description = "example tag"
  		name = "example-tag"
        tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"

		is_retired = false
}
`
)

var (
	TagRequiredOnlyResource = TagResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_tag", "test_tag", Required, Create, tagRepresentation)

	tagDataSourceRepresentation = map[string]interface{}{
		"tag_namespace_id": Representation{repType: Required, create: `${oci_identity_tag_namespace.test_tag_namespace.id}`},
		"filter":           RepresentationGroup{Required, tagDataSourceFilterRepresentation}}
	tagDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_tag.test_tag.id}`}},
	}

	tagRepresentation = map[string]interface{}{
		"description":      Representation{repType: Required, create: `This tag will show the cost center that will be used for billing of associated resources.`, update: `description2`},
		"name":             Representation{repType: Required, create: `CostCenter`},
		"tag_namespace_id": Representation{repType: Required, create: `${oci_identity_tag_namespace.test_tag_namespace.id}`},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	TagResourceDependencies = TagNamespaceRequiredOnlyResource
)

func TestIdentityTagResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_tag.test_tag"
	datasourceName := "data.oci_identity_tags.test_tags"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + TagResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag", "test_tag", Required, Create, tagRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "This tag will show the cost center that will be used for billing of associated resources."),
					resource.TestCheckResourceAttr(resourceName, "name", "CostCenter"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + TagResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + TagResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag", "test_tag", Optional, Create, tagRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "This tag will show the cost center that will be used for billing of associated resources."),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_retired"),
					resource.TestCheckResourceAttr(resourceName, "name", "CostCenter"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + TagResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag", "test_tag", Optional, Update, tagRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_retired"),
					resource.TestCheckResourceAttr(resourceName, "name", "CostCenter"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),
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
					generateDataSourceFromRepresentationMap("oci_identity_tags", "test_tags", Optional, Update, tagDataSourceRepresentation) +
					compartmentIdVariableStr + TagResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag", "test_tag", Optional, Update, tagRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "tag_namespace_id"),

					resource.TestCheckResourceAttr(datasourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.is_retired"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.name", "CostCenter"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.time_created"),
				),
			},
		},
	})
}
