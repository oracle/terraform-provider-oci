// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func CustomerSecretKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCustomerSecretKeys,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"customer_secret_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     CustomerSecretKeyResource(),
			},
		},
	}
}

func readCustomerSecretKeys(d *schema.ResourceData, m interface{}) error {
	sync := &CustomerSecretKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type CustomerSecretKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListCustomerSecretKeysResponse
}

func (s *CustomerSecretKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CustomerSecretKeysDataSourceCrud) Get() error {
	request := oci_identity.ListCustomerSecretKeysRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListCustomerSecretKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CustomerSecretKeysDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		customerSecretKey := map[string]interface{}{
			"user_id": *r.UserId,
		}

		if r.DisplayName != nil {
			customerSecretKey["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			customerSecretKey["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			customerSecretKey["inactive_state"] = *r.InactiveStatus
		}

		customerSecretKey["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			customerSecretKey["time_created"] = r.TimeCreated.String()
		}

		if r.TimeExpires != nil {
			customerSecretKey["time_expires"] = r.TimeExpires.String()
		}

		resources = append(resources, customerSecretKey)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CustomerSecretKeysDataSource().Schema["customer_secret_keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("customer_secret_keys", resources); err != nil {
		panic(err)
	}

	return
}
