package integrationtest

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/stretchr/testify/assert"
)

func TestProviderServiceEndpointOverridePrecedenceOverDomainOverrideWhenBothSet(t *testing.T) {
	domainOverride := "orac.cloud80.com"
	clientHostOverrides := "oci_identity.IdentityClient=identity.us-mars-1.dont.change.domain.com"
	expectedHost := strings.Split(clientHostOverrides, "=")[1]

	t.Setenv(globalvar.DomainNameOverrideEnv, domainOverride)
	t.Setenv(globalvar.ClientHostOverridesEnv, clientHostOverrides)

	client := acctest.GetTestClients(&schema.ResourceData{}).IdentityClient()
	assert.Equal(t, expectedHost, client.Host)
}
