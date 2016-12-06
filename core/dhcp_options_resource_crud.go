package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type DHCPOptionsResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.DHCPOptions
}

func (s *DHCPOptionsResourceCrud) ID() string {
	return s.Res.ID
}

func (s *DHCPOptionsResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *DHCPOptionsResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *DHCPOptionsResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *DHCPOptionsResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *DHCPOptionsResourceCrud) State() string {
	return s.Res.State
}

func (s *DHCPOptionsResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = s.D.Get("display_name").(string)

	s.Res, e = s.Client.CreateDHCPOptions(compartmentID, vcnID, s.buildEntities(), opts)

	return
}

func (s *DHCPOptionsResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetDHCPOptions(s.D.Id())
	return
}

func (s *DHCPOptionsResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateDHCPDNSOptions{}
	opts.Options = s.buildEntities()

	s.Res, e = s.Client.UpdateDHCPOptions(s.D.Id(), opts)
	return
}

func (s *DHCPOptionsResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)

	entities := []map[string]interface{}{}
	for _, val := range s.Res.Options {
		entity := map[string]interface{}{
			"type":               val.Type,
			"custom_dns_servers": val.CustomDNSServers,
			"server_type":        val.ServerType,
		}
		entities = append(entities, entity)
	}
	s.D.Set("options", entities)

	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *DHCPOptionsResourceCrud) Delete() (e error) {
	return s.Client.DeleteDHCPOptions(s.D.Id(), nil)
}

func (s *DHCPOptionsResourceCrud) buildEntities() (entities []baremetal.DHCPDNSOption) {
	entities = []baremetal.DHCPDNSOption{}
	for _, val := range s.D.Get("options").([]interface{}) {
		data := val.(map[string]interface{})

		servers := []string{}
		for _, val := range data["custom_dns_servers"].([]interface{}) {
			servers = append(servers, val.(string))
		}

		entity := baremetal.DHCPDNSOption{
			Type:             data["type"].(string),
			CustomDNSServers: servers,
			ServerType:       data["server_type"].(string),
		}
		entities = append(entities, entity)
	}
	return
}
