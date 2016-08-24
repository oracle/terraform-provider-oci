package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type IPSecDatasourceConfigCrud struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.IPSecConnectionDeviceConfig
}

func (s *IPSecDatasourceConfigCrud) Get() (e error) {
	ipsecID := s.D.Get("ipsec_id").(string)
	s.Resource, e = s.Client.GetIPSecConnectionDeviceConfig(ipsecID)
	return
}

func (s *IPSecDatasourceConfigCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(s.Resource.ID)
		s.D.Set("compartment_id", s.Resource.CompartmentID)
		s.D.Set("id", s.Resource.ID)
		s.D.Set("time_created", s.Resource.TimeCreated)

		tunnels := []map[string]interface{}{}

		for _, val := range s.Resource.Tunnels {
			tunnel := map[string]interface{}{
				"ip_address":    val.IPAddress,
				"shared_secret": val.SharedSecret,
				"time_created":  val.TimeCreated.String(),
			}

			tunnels = append(tunnels, tunnel)
		}

		s.D.Set("tunnels", tunnels)

	}
}
