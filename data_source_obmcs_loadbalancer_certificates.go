// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func CertificateDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readCertificate,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"certificates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerCertificateResource(),
			},
		},
	}
}

func readCertificate(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &CertificateDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type CertificateDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListCertificates
}

func (s *CertificateDatasourceCrud) Get() (e error) {
	lbID := s.D.Get("load_balancer_id").(string)
	s.Res, e = s.Client.ListCertificates(lbID, nil)
	return
}

func (s *CertificateDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Certificates {
			res := map[string]interface{}{
				"ca_certificate":     v.CACertificate,
				"certificate_name":   v.CertificateName,
				"public_certificate": v.PublicCertificate,
			}
			resources = append(resources, res)
		}
		s.D.Set("certificates", resources)
	}
	return
}
