// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"
)

func IdentityApiKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityApiKeys,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityApiKeyResource()),
			},
		},
	}
}

func readIdentityApiKeys(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityApiKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityApiKeysDataSource-", IdentityApiKeysDataSource(), s.D))
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityApiKeysDataSource().Schema["api_keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("api_keys", resources); err != nil {
		return err
	}

	return nil
}
