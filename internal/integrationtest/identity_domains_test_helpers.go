package integrationtest

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// by default is terraformtest-donottouch domain in domainsuxtest tenancy (YYZ)
var domainId = utils.GetEnvSettingWithDefault("identity_domain_id", "ocid1.domain.oc1..aaaaaaaanjdr33uufgat3t7mftmfo6qeggs3mgu43whydson2eeqtqoijaxq")
var domainIdVariableStr = fmt.Sprintf("variable \"domain_id\" { default = \"%s\" }\n", domainId)
var TestDomainDependencies = domainIdVariableStr + `
		data "oci_identity_domain" "test_domain" {
			domain_id = "${var.domain_id}"
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
