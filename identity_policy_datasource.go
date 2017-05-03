// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func IdentityPolicyDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityPolicies,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     PolicyResource(),
			},
		},
	}
}

func readIdentityPolicies(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &IdentityPolicyDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type IdentityPolicyDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListPolicies
}

func (s *IdentityPolicyDatasourceCrud) Get() (e error) {
	compartment_id := s.D.Get("compartment_id").(string)

	s.Res, e = s.Client.ListPolicies(compartment_id, nil)
	return
}

func (s *IdentityPolicyDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Policies {
			res := map[string]interface{}{
				"id":             v.ID,
				"compartment_id": v.CompartmentID,
				"name":           v.Name,
				"statements":     v.Statements,
				"description":    v.Description,
				"time_created":   v.TimeCreated.String(),
				"state":          v.State,
				"inactive_state": v.InactiveStatus,
				"version_date":   v.VersionDate.String(),
			}
			resources = append(resources, res)
		}
		if err := s.D.Set("policies", resources); err != nil {
			panic(err)
		}
	}
	return
}
