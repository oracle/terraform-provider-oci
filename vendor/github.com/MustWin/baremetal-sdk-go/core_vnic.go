package baremetal

// Vnic describes a virtual network interface.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vnic/
type Vnic struct {
	ETaggedResource
	AvailabilityDomain string `json:"availabilityDomain"`
	CompartmentID      string `json:"compartmentId"`
	DisplayName        string `json:"displayName"`
	ID                 string `json:"id"`
	State              string `json:"lifecycleState"`
	PrivateIPAddress   string `json:"privateIp"`
	PublicIPAddress    string `json:"publicIp"`
	SubnetID           string `json:"subnetId"`
	TimeCreated        Time   `json:"timeCreated"`
}

// GetVnic retrieves information about a virtual network interface identified
// by vnicID. ListVnicAttachments can be used to retrieve Vnic IDs.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vnic/GetVnic
func (c *Client) GetVnic(id string) (vnic *Vnic, e error) {
	details := &requestDetails{
		name: resourceVnics,
		ids:  urlParts{id},
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	vnic = &Vnic{}
	e = resp.unmarshal(vnic)
	return
}
