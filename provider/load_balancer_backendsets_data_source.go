// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func BackendSetDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readBackendSets,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backendsets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerBackendSetResource(),
			},
		},
	}
}

func readBackendSets(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &BackendSetDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

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
	if s.Res == nil {
		panic("LoadBalancer Backend Resource is nil, cannot SetData")
	}
	s.D.SetId(time.Now().UTC().String())
	resources := []map[string]interface{}{}
	for _, v := range s.Res.BackendSets {
		var healthChecker []map[string]interface{}
		if hc := v.HealthChecker; hc != nil {
			healthChecker = []map[string]interface{}{
				{
					"interval_ms":         hc.IntervalInMS,
					"port":                hc.Port,
					"protocol":            hc.Protocol,
					"response_body_regex": hc.ResponseBodyRegex,
				},
			}
		}

		var sslConfig []map[string]interface{}
		if ssl := v.SSLConfig; ssl != nil {
			sslConfig = []map[string]interface{}{
				{
					"certificate_name":        ssl.CertificateName,
					"verify_depth":            ssl.VerifyDepth,
					"verify_peer_certificate": ssl.VerifyPeerCertificate,
				},
			}
		}

		var sessionConfig []map[string]interface{}
		if session := v.SessionPersistenceConfig; session != nil {
			sessionConfig = []map[string]interface{}{
				{
					"cookie_name":      session.CookieName,
					"disable_fallback": session.DisableFallback,
				},
			}
		}

		backends := []map[string]interface{}{}
		for _, backend := range v.Backends {
			res := map[string]interface{}{
				"ip_address": backend.IPAddress,
				"port":       backend.Port,
				"backup":     backend.Backup,
				"drain":      backend.Drain,
				"offline":    backend.Offline,
				"weight":     backend.Weight,
			}
			backends = append(backends, res)
		}
		res := map[string]interface{}{
			"name":                              v.Name,
			"policy":                            v.Policy,
			"health_checker":                    healthChecker,
			"ssl_configuration":                 sslConfig,
			"session_persistence_configuration": sessionConfig,
			"backend":                           backends,
		}
		resources = append(resources, res)
	}
	err := s.D.Set("backendsets", resources)
	if err != nil {
		log.Printf("[ERROR] Failed to set load_balancers: %v", err)
	}
	return
}
