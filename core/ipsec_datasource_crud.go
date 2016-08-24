package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type IPSecDatasourceCrud struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.ListIPSecConnections
}

func (s *IPSecDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	opts := getCoreOptionsFromResourceData(s.D, "drg_id", "cpe_id")

	s.Resource, e = s.Client.ListIPSecConnections(compartmentID, opts...)
	return

}

func (s IPSecDatasourceCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Resource.Connections {

			resource := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"drg_id":         v.DrgID,
				"cpe_id":         v.CpeID,
				"display_name":   v.DisplayName,
				"id":             v.ID,
				"state":          v.State,
				"static_routes":  v.StaticRoutes,
				"time_created":   v.TimeCreated.String(),
			}

			resources = append(resources, resource)
		}

		s.D.Set("connections", resources)

	}

	return
}

type IPSecDatasourceStatusCrud struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.IPSecConnectionDeviceStatus
}

func (s *IPSecDatasourceStatusCrud) Get() (e error) {
	ipsecID := s.D.Get("ipsec_id").(string)
	s.Resource, e = s.Client.GetIPSecConnectionDeviceStatus(ipsecID)
	return
}

func (s *IPSecDatasourceStatusCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(s.Resource.ID)
		s.D.Set("compartment_id", s.Resource.CompartmentID)
		s.D.Set("id", s.Resource.ID)
		s.D.Set("time_created", s.Resource.TimeCreated)

		tunnels := []map[string]interface{}{}

		for _, val := range s.Resource.Tunnels {
			tunnel := map[string]interface{}{
				"ip_address":         val.IPAddress,
				"state":              val.State,
				"time_created":       val.TimeCreated.String(),
				"time_state_modifed": val.TimeStateModified.String(),
			}

			tunnels = append(tunnels, tunnel)
		}

		s.D.Set("tunnels", tunnels)

	}
}

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
