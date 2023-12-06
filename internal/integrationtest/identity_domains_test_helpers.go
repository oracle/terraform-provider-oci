package integrationtest

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var domainId = utils.GetEnvSettingWithBlankDefault("identity_domain_id")
var domainIdVariableStr = fmt.Sprintf("variable \"domain_id\" { default = \"%s\" }\n", domainId)
var domainIdForMyEndpoint = utils.GetEnvSettingWithBlankDefault("identity_domain_id_for_my_endpoint")
var domainIdForMyEndpointVariableStr = fmt.Sprintf("variable \"domain_id\" { default = \"%s\" }\n", domainIdForMyEndpoint)
var testDomainDataSourceStr = `
	data "oci_identity_domain" "test_domain" {
		domain_id = "${var.domain_id}"
	}
`

var TestDomainDependencies = domainIdVariableStr + testDomainDataSourceStr
var TestDomainForMyEndpointDependencies = domainIdForMyEndpointVariableStr + testDomainDataSourceStr

// User dependency
var IdentityDomainsUserManager = `
resource "oci_identity_domains_user" "test_user_manager" {
	# Required
	emails {
		value = "value@email.com"
		type = "work"
		primary = "true"
	}
	idcs_endpoint = "${data.oci_identity_domain.test_domain.url}"
	name {
		family_name = "managerFamilyName"
	}
	schemas = ["urn:ietf:params:scim:schemas:core:2.0:User"]
	user_name = "managerUserName"
	lifecycle {
		ignore_changes = ["urnietfparamsscimschemasoracleidcsextension_oci_tags[0].defined_tags", "emails", "schemas"]
	}
}
`

func getIdentityDomainsCompositeId(idcsEndpoint string, resourceName string, resId string) string {
	// e.g. idcsEndpoint/https://something.com/groups/{groupId}
	return fmt.Sprintf("idcsEndpoint/%s/%s/%s", idcsEndpoint, resourceName, resId)
}

func getIdentityDomainsImportIdFn(typeName string, resourceName string) func(*terraform.State) (string, error) {
	return func(state *terraform.State) (string, error) {
		for _, rs := range state.RootModule().Resources {
			idcsEndpoint := rs.Primary.Attributes["idcs_endpoint"]
			id := rs.Primary.ID
			if rs.Type == typeName {
				return getIdentityDomainsCompositeId(idcsEndpoint, resourceName, id), nil
			}
		}

		return "", fmt.Errorf("[ERROR] unable to Create import id as no resource of type %s in state", typeName)
	}
}

var SamlServiceProviderDependencies = `
	resource "oci_identity_domains_app" "test_saml_app" {
		idcs_endpoint = data.oci_identity_domain.test_domain.url
		based_on_template {
			value = "CustomWebAppTemplateId"
		}
		display_name = "TestSamlApp"
		schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:App"]
		lifecycle {
			ignore_changes = [schemas]
		}
	}
`

var GrantTestAppDependencies = `
	data "oci_identity_domains_apps" "test_grant_apps" {
	  idcs_endpoint = data.oci_identity_domain.test_domain.url
	  app_filter                   = "displayName sw \"GrantTestApp\""
	  attribute_sets               = ["all"]
	}
`

var testUserOcid = utils.GetEnvSettingWithBlankDefault("user_ocid")
var testUserOcidVarStr = fmt.Sprintf("variable \"my_user_ocid\" { default = \"%s\" }\n", testUserOcid)

func GenerateMyAppTestApp(activeValue string) string {
	return fmt.Sprintf("resource \"oci_identity_domains_app\" \"my_app_test_app\" {\n\t\tidcs_endpoint = data.oci_identity_domain.test_domain.url\n\t\tbased_on_template {\n\t\t\tvalue = \"CustomWebAppTemplateId\"\n\t\t}\n\t\tdisplay_name = \"MyAppTestApp\"\n\t\tschemas = [\"urn:ietf:params:scim:schemas:oracle:idcs:App\"]\n\t\tlifecycle {\n\t\t\tignore_changes = [schemas]\n\t\t}\n\t\tshow_in_my_apps = true\n\t\tactive = %s\n\t}\n", activeValue)
}

var MyAppTestGrantDependencies = testUserOcidVarStr + `
	data "oci_identity_domains_user" "my_user" {
		idcs_endpoint = data.oci_identity_domain.test_domain.url
		user_id = var.my_user_ocid
	}
	resource "oci_identity_domains_grant" "my_app_test_grant" {
		idcs_endpoint = data.oci_identity_domain.test_domain.url
		grant_mechanism = "ADMINISTRATOR_TO_USER"
		schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:Grant"]
		grantee {
			type = "User"
			value = data.oci_identity_domains_user.my_user.id
		}
		app {
			value = oci_identity_domains_app.my_app_test_app.id
		}
	}
`
