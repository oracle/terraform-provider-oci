// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ApiKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApiKeys,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ApiKeyResource(),
			},
		},
	}
}

func readApiKeys(d *schema.ResourceData, m interface{}) error {
	sync := &ApiKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type ApiKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListApiKeysResponse
}

func (s *ApiKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApiKeysDataSourceCrud) Get() error {
	request := oci_identity.ListApiKeysRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListApiKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApiKeysDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		apiKey := map[string]interface{}{
			"user_id": *r.UserId,
		}

		if r.Fingerprint != nil {
			apiKey["fingerprint"] = *r.Fingerprint
		}

		if r.KeyId != nil {
			apiKey["id"] = *r.KeyId
		}

		if r.InactiveStatus != nil {
			apiKey["inactive_status"] = *r.InactiveStatus
		}

		if r.KeyValue != nil {
			apiKey["key_value"] = *r.KeyValue
		}

		apiKey["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			apiKey["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, apiKey)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ApiKeysDataSource().Schema["api_keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("api_keys", resources); err != nil {
		panic(err)
	}

	return
}
