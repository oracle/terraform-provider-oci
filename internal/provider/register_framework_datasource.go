package provider

import (
	"github.com/oracle/oci-go-sdk/v65/common"
	tf_vault "github.com/oracle/terraform-provider-oci/internal/service-framework/vault"
)

func init() {
	if common.CheckForEnabledServices("vault") {
		tf_vault.RegisterFrameworkDataSource()
	}
}
