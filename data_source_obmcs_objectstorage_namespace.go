// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func NamespaceDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readNamespace,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readNamespace(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	reader := &NamespaceDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}

type NamespaceDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.Namespace
}

func (s *NamespaceDatasourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetNamespace()
	return
}

func (s *NamespaceDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		s.D.Set("namespace", string(*s.Res))
	}
	return
}
