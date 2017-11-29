// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func APIKeyDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readAPIKeys,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     APIKeyResource(),
			},
		},
	}
}

func readAPIKeys(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &APIKeyDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type APIKeyDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListAPIKeyResponses
}

func (s *APIKeyDatasourceCrud) Get() (e error) {
	userID := s.D.Get("user_id").(string)
	s.Res, e = s.Client.ListAPIKeys(userID)
	return
}

func (s *APIKeyDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Keys {
			res := map[string]interface{}{
				"fingerprint":  v.Fingerprint,
				"id":           v.KeyID,
				"key_value":    v.KeyValue,
				"state":        v.State,
				"time_created": v.TimeCreated.String(),
				"user_id":      v.UserID,
			}
			resources = append(resources, res)
		}
		s.D.Set("api_keys", resources)
	}
	return
}
