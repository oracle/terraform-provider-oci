// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ObjectHeadDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readObjectHead,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"object": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content-length": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"content-type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func readObjectHead(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	reader := &ObjectHeadDatasourceCrud{}
	reader.D = d
	reader.Client = client.client

	return crud.ReadResource(reader)
}

type ObjectHeadDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.HeadObject
}

func (s *ObjectHeadDatasourceCrud) Get() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	object := s.D.Get("object").(string)

	s.Res, e = s.Client.HeadObject(baremetal.Namespace(namespace), bucket, object, &baremetal.HeadObjectOptions{})
	return
}

func (s *ObjectHeadDatasourceCrud) SetData() {
	// Important, if you don't have an ID, make one up for your datasource
	// or things will end in tears
	s.D.SetId(time.Now().UTC().String())
	s.D.Set("metadata", s.Res.Metadata)
	s.D.Set("content-length", string(s.Res.ContentLength))
	s.D.Set("content-type", s.Res.ContentType)
	return
}
