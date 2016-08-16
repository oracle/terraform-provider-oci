package baremetal

// AvailablityDomain contains name and then tenancy ID that an
import (
	"bytes"
	"encoding/json"
)

// availability domain belongs to.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#AvailabilityDomain
type AvailabilityDomain struct {
	Name          string `json:"name"`
	CompartmentID string `json:"compartmentId"`
}

// ListAvailablityDomains lists availability domains in a user's root tenancy.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listAvailabilityDomains
func (c *Client) ListAvailablityDomains(compartmentID string) (ads []AvailabilityDomain, e error) {
	reqOpts := &sdkRequestOptions{
		name: resourceAvailabilityDomains,
		ocid: compartmentID,
	}

	var getResp *requestResponse
	if getResp, e = c.identityApi.getRequest(reqOpts); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	e = decoder.Decode(&ads)
	return
}
