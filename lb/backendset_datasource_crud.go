// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package lb

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type BackendSetDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListBackendSets
}

func (s *BackendSetDatasourceCrud) Get() (e error) {
	lbID := s.D.Get("load_balancer_id").(string)
	s.Res, e = s.Client.ListBackendSets(lbID, nil)
	return
}

func (s *BackendSetDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.BackendSets {
			backends := []map[string]interface{}{}
			for _, backend := range v.Backends {
				res := map[string]interface{}{
					"name":       backend.Name,
					"ip_address": backend.IPAddress,
					"port":       backend.Port,
					"backup":     backend.Backup,
					"drain":      backend.Drain,
					"offline":    backend.Offline,
					"weight":     backend.Weight,
				}
				backends = append(backends, res)
			}

			hcState := s.D.Get("health_checker").(map[string]interface{})
			healthChecker := baremetal.HealthChecker{
				IntervalInMS:      hcState["interval_ms"].(int),
				Port:              hcState["port"].(int),
				Protocol:          hcState["protocol"].(string),
				ResponseBodyRegex: hcState["response_body_regex"].(string),
			}

			sslState := s.D.Get("ssl_configuration").(map[string]interface{})
			sslConfig := baremetal.SSLConfiguration{
				CertificateName:       sslState["certificate_name"].(string),
				VerifyDepth:           sslState["verify_depth"].(int),
				VerifyPeerCertificate: sslState["verify_peer_certificate"].(bool),
			}
			res := map[string]interface{}{
				"name":              v.Name,
				"policy":            v.Policy,
				"health_checker":    healthChecker,
				"ssl_configuration": sslConfig,
				"backends":          backends,
			}
			resources = append(resources, res)
		}
		s.D.Set("backendsets", resources)
	}
	return
}
