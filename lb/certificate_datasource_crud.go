// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package lb

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

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
