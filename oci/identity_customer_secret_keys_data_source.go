// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v44/identity"
)

func init() {
	RegisterDatasource("oci_identity_customer_secret_keys", IdentityCustomerSecretKeysDataSource())
}

func IdentityCustomerSecretKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityCustomerSecretKeys,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"customer_secret_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(IdentityCustomerSecretKeyResource()),
			},
		},
	}
}

func readIdentityCustomerSecretKeys(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCustomerSecretKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityCustomerSecretKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListCustomerSecretKeysResponse
}

func (s *IdentityCustomerSecretKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityCustomerSecretKeysDataSourceCrud) Get() error {
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

func (s *IdentityCustomerSecretKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("IdentityCustomerSecretKeysDataSource-", IdentityCustomerSecretKeysDataSource(), s.D))
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
			customerSecretKey["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
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
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityCustomerSecretKeysDataSource().Schema["customer_secret_keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("customer_secret_keys", resources); err != nil {
		return err
	}

	return nil
}
