// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DHCPOptionsResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createDHCPOptions,
		Read:     readDHCPOptions,
		Update:   updateDHCPOptions,
		Delete:   deleteDHCPOptions,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"options": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"custom_dns_servers": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"server_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"search_domain_names": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.CreateResource(d, crd)
}

func readDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.ReadResource(crd)
}

func updateDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.UpdateResource(d, crd)
}

func deleteDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.DeleteResource(d, crd)
}

type DHCPOptionsResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.DHCPOptions
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
			"type":                val.Type,
			"custom_dns_servers":  val.CustomDNSServers,
			"server_type":         val.ServerType,
			"search_domain_names": val.SearchDomainNames,
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
		if len(servers) == 0 {
			servers = nil
		}
		searchDomains := []string{}
		for _, val := range data["search_domain_names"].([]interface{}) {
			searchDomains = append(searchDomains, val.(string))
		}
		if len(searchDomains) == 0 {
			searchDomains = nil
		}
		entity := baremetal.DHCPDNSOption{
			Type:              data["type"].(string),
			CustomDNSServers:  servers,
			ServerType:        data["server_type"].(string),
			SearchDomainNames: searchDomains,
		}
		entities = append(entities, entity)
	}
	return
}
