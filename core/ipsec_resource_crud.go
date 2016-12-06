package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type IPSecResourceCrud struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.IPSecConnection
}

func (s *IPSecResourceCrud) ID() string {
	return s.Resource.ID
}

// TODO: I'm not sure whether we need to
//       manage delete states for IPSec tunnels.
//       We'll need to determine this in testing.
// func (s *IPSecSync) CreatedTarget() []string {
// 	return []string{baremetal.ResourceUp,
//   baremetal.ResourceDownForMaintenance}
// }
//
// func (s *IPSecSync) DeletedTarget() []string {
// 	return []string{baremetal.ResourceDown}
// }

func (s *IPSecResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	cpeID := s.D.Get("cpe_id").(string)
	drgID := s.D.Get("drg_id").(string)

	staticRoutes := []string{}
	for _, route := range s.D.Get("static_routes").([]interface{}) {
		staticRoutes = append(staticRoutes, route.(string))
	}

	opts := &baremetal.CreateOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.CreateIPSecConnection(
		compartmentID,
		cpeID,
		drgID,
		staticRoutes,
		opts,
	)

	return
}

func (s *IPSecResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetIPSecConnection(s.D.Id())
	return
}

func (s *IPSecResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("cpe_id", s.Resource.CpeID)
	s.D.Set("drg_id", s.Resource.DrgID)
	s.D.Set("static_routes", s.Resource.StaticRoutes)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("state", s.Resource.State)
	s.D.Set("time_created", s.Resource.TimeCreated.String())

}

func (s *IPSecResourceCrud) Delete() (e error) {
	return s.Client.DeleteIPSecConnection(s.D.Id(), nil)
}
