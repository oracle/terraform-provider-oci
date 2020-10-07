// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v26/identity"
)

func init() {
	RegisterDatasource("oci_identity_api_keys", IdentityApiKeysDataSource())
}

func IdentityApiKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityApiKeys,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(IdentityApiKeyResource()),
			},
		},
	}
}

func readIdentityApiKeys(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityApiKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityApiKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListApiKeysResponse
}

func (s *IdentityApiKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityApiKeysDataSourceCrud) Get() error {
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

func (s *IdentityApiKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
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
			apiKey["inactive_status"] = strconv.FormatInt(*r.InactiveStatus, 10)
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
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityApiKeysDataSource().Schema["api_keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("api_keys", resources); err != nil {
		return err
	}

	return nil
}
