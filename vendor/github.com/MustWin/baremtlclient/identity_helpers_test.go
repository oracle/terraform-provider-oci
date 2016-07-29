package baremtlclient

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildIdentityURL(t *testing.T) {
	expected := "https://identity.us-az-phoenix-1.OracleIaaS.com/v1/policies/"
	actual := buildIdentityURL(resourcePolicies, nil)
	assert.Equal(t, expected, actual)

	expected = "https://identity.us-az-phoenix-1.OracleIaaS.com/v1/policies?foo=bar%2Fbaz"
	actual = buildIdentityURL(resourcePolicies, &url.Values{"foo": []string{"bar/baz"}})
	assert.Equal(t, expected, actual)

	expected = "https://identity.us-az-phoenix-1.OracleIaaS.com/v1/policies/one/two?foo=bar%2Fbaz"
	actual = buildIdentityURL(resourcePolicies, &url.Values{"foo": []string{"bar/baz"}}, "one", "two")
	assert.Equal(t, expected, actual)

	expected = "https://identity.us-az-phoenix-1.OracleIaaS.com/v1/policies/one/two"
	actual = buildIdentityURL(resourcePolicies, nil, "one", "two")
	assert.Equal(t, expected, actual)

}
