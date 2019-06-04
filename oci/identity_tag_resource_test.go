package provider

import (
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	namespaceResourceName = "oci_identity_tag_namespace.test-tag-namespace"
	tagResourceNames      = [2]string{"oci_identity_tag.test-tag1", "oci_identity_tag.test-tag2"}
	costTagResourceNames  = [2]string{"oci_identity_tag.test-cost-tag1", "oci_identity_tag.test-cost-tag2"}

	namespaceResourceValue = randomString(5, charsetWithoutDigits) + "-delete-namespace"
	tagResourceValues      = [2]string{randomString(5, charsetWithoutDigits) + "-tag-1", randomString(5, charsetWithoutDigits) + "-tag-2"}
	costTagResourceValues  = [2]string{randomString(5, charsetWithoutDigits) + "-cost-tag-1", randomString(5, charsetWithoutDigits) + "-cost-tag-2"}
)

// This test will be executed in a separate suite with 'tags_import_if_exists = false'
func TestIdentityTagDeletion(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTagDeletion")
	defer httpreplay.SaveScenario()

	importIfExists, _ := strconv.ParseBool(getEnvSettingWithDefault("tags_import_if_exists", "false"))
	if importIfExists {
		t.Skip("[WARN] TestIdentityTagDeletion requires 'tags_import_if_exists' to be set to false ")
	}

	provider := testAccProvider
	config := legacyTestProviderConfig() + `
	resource "oci_identity_tag_namespace" "test-tag-namespace" {
		compartment_id = "${var.compartment_id}"
		description    = "test tag namespace"
		name           = "` + namespaceResourceValue + `"
						
		freeform_tags = {
			"Department" = "Finance"
		}
	}`

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
