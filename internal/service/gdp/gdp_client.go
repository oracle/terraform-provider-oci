package gdp

import (
	"strings"

	oci_gdp "github.com/oracle/oci-go-sdk/v65/gdp"

	"github.com/oracle/terraform-provider-oci/internal/client"
)

func getGdpClient(m interface{}, useCommercialEndpoint bool) *oci_gdp.GuardedDataPipelineClient {
	defaultClient := m.(*client.OracleClients).GuardedDataPipelineClient()
	if !useCommercialEndpoint {
		return defaultClient
	}

	commercialClient := *defaultClient
	commercialClient.Host = strings.Replace(defaultClient.Host, "://gdp.", "://"+commercialSubdomain+".", 1)
	return &commercialClient
}
