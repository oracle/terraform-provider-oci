package integrationtest

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	namespaceResourceName = "oci_identity_tag_namespace.test-tag-namespace"
	tagResourceNames      = [2]string{"oci_identity_tag.test-tag1", "oci_identity_tag.test-tag2"}
	costTagResourceNames  = [2]string{"oci_identity_tag.test-cost-tag1", "oci_identity_tag.test-cost-tag2"}

	namespaceResourceValue = utils.RandomString(5, utils.CharsetWithoutDigits) + "-delete-namespace"
	tagResourceValues      = [2]string{utils.RandomString(5, utils.CharsetWithoutDigits) + "-tag-1", utils.RandomString(5, utils.CharsetWithoutDigits) + "-tag-2"}
	costTagResourceValues  = [2]string{utils.RandomString(5, utils.CharsetWithoutDigits) + "-cost-tag-1", utils.RandomString(5, utils.CharsetWithoutDigits) + "-cost-tag-2"}
)

// This test will be executed in a separate suite with 'tags_import_if_exists = false'
// issue-routing-tag: identity/default
func TestResourceIdentityTagDeletion(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTagDeletion")
	defer httpreplay.SaveScenario()

	importIfExists, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("tags_import_if_exists", "false"))
	if importIfExists {
		t.Skip("[WARN] TestIdentityTagDeletion requires 'tags_import_if_exists' to be set to false ")
	}

	provider := acctest.TestAccProvider
	config := acctest.LegacyTestProviderConfig() + `
	resource "oci_identity_tag_namespace" "test-tag-namespace" {
		compartment_id = "${var.compartment_id}"
		description    = "test tag namespace"
		name           = "` + namespaceResourceValue + `"
						
		freeform_tags = {
			"Department" = "Finance"
		}
	}`

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify deletion of simple tags
			{
				Config: config + `
					resource "oci_identity_tag" "test-tag1" {
						description      = "tf deletion example tag-1"
  						name             = "` + tagResourceValues[0] + `"
  						tag_namespace_id = "${oci_identity_tag_namespace.test-tag-namespace.id}"
					}
					resource "oci_identity_tag" "test-tag2" {
						description      = "tf deletion example tag-2"
  						name             = "` + tagResourceValues[1] + `"
  						tag_namespace_id = "${oci_identity_tag_namespace.test-tag-namespace.id}"
					}
					`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(namespaceResourceName, "name", namespaceResourceValue),
					resource.TestCheckResourceAttr(tagResourceNames[0], "description", "tf deletion example tag-1"),
					resource.TestCheckResourceAttr(tagResourceNames[0], "name", tagResourceValues[0]),
					resource.TestCheckResourceAttr(tagResourceNames[1], "description", "tf deletion example tag-2"),
					resource.TestCheckResourceAttr(tagResourceNames[1], "name", tagResourceValues[1]),
				),
			},
			// verify deletion of cost tracking tags
			{
				Config: config + `
					resource "oci_identity_tag" "test-cost-tag1" {
						description      = "tf cost tracking deletion example tag-1"
  						name             = "` + costTagResourceValues[0] + `"
  						tag_namespace_id = "${oci_identity_tag_namespace.test-tag-namespace.id}"
						is_cost_tracking = true 	
					}
					resource "oci_identity_tag" "test-cost-tag2" {
						description      = "tf cost tracking deletion example tag-2"
  						name             = "` + costTagResourceValues[1] + `"
  						tag_namespace_id = "${oci_identity_tag_namespace.test-tag-namespace.id}"
						is_cost_tracking = true 	
					}
					`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(namespaceResourceName, "name", namespaceResourceValue),
					resource.TestCheckResourceAttr(costTagResourceNames[0], "description", "tf cost tracking deletion example tag-1"),
					resource.TestCheckResourceAttr(costTagResourceNames[0], "name", costTagResourceValues[0]),
					resource.TestCheckResourceAttr(costTagResourceNames[1], "description", "tf cost tracking deletion example tag-2"),
					resource.TestCheckResourceAttr(costTagResourceNames[1], "name", costTagResourceValues[1]),
				),
			},
		},
	})
}

// execute this test in identity compartment only and not on root compartment
// issue-routing-tag: identity/default
func TestResourceIdentityDefaultTag_required(t *testing.T) {
	httpreplay.SetScenario("TestResourceIdentityDefaultTag_required")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	config := acctest.LegacyTestProviderConfig() + `
	variable defined_tag_namespace_name { default = "" }
	resource "oci_identity_tag_namespace" "tag-namespace1" {
  		#Required
		compartment_id = "${var.compartment_id}"
  		description = "Just a test"
  		name = "example-tag-default-namespace"

		is_retired = false
	}
	resource "oci_identity_tag" "tag1" {
  		#Required
  		description = "tf example tag 2"
  		name = "tf-example-tag-default"
        tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"

		is_retired = false
	}
	`

	resourceName := "oci_identity_tag_default.test_tag_default"
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: config +
					`
					resource "oci_identity_tag_default" "test_tag_default" {
						compartment_id = "${var.compartment_id}"
						is_required = "true"
						value="value1" 
						tag_definition_id = "${oci_identity_tag.tag1.id}"
					}
					resource "oci_identity_tag_namespace" "tag-namespace-test-1" {
  						compartment_id = "${var.compartment_id}"
  						description = "test namespace 1"
  						name = "example-test-delete-default-namespace-1"

						defined_tags   = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${oci_identity_tag_default.test_tag_default.value}")}"
						is_retired = false
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "tag_definition_id"),
					resource.TestCheckResourceAttr(resourceName, "value", "value1"),
					// stream
					resource.TestCheckResourceAttr("oci_identity_tag_namespace.tag-namespace-test-1", "compartment_id", compartmentId),
				),
			},
			{
				Config: config +
					`
				resource "oci_identity_tag_default" "test_tag_default" {
					compartment_id = "${var.compartment_id}"
					is_required = "false"
					value="value1" 
					tag_definition_id = "${oci_identity_tag.tag1.id}"
				}
				resource "oci_identity_tag_namespace" "tag-namespace-test-2" {
  						compartment_id = "${var.compartment_id}"
  						description = "test namespace 2"
  						name = "example-test-delete-default-namespace-2"

						is_retired = false
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "tag_definition_id"),
					resource.TestCheckResourceAttr(resourceName, "value", "value1"),
					// stream
					resource.TestCheckResourceAttr("oci_identity_tag_namespace.tag-namespace-test-2", "compartment_id", compartmentId),
				),
			},
		},
	})
}
