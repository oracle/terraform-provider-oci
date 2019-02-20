// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentitySwiftPasswordsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentitySwiftPasswords,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"passwords": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(IdentitySwiftPasswordResource()),
			},
		},
	}
}

func readIdentitySwiftPasswords(d *schema.ResourceData, m interface{}) error {
	sync := &IdentitySwiftPasswordsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type IdentitySwiftPasswordsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListSwiftPasswordsResponse
}

func (s *IdentitySwiftPasswordsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentitySwiftPasswordsDataSourceCrud) Get() error {
	request := oci_identity.ListSwiftPasswordsRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListSwiftPasswords(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentitySwiftPasswordsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		swiftPassword := map[string]interface{}{
			"user_id": *r.UserId,
		}

		if r.Description != nil {
			swiftPassword["description"] = *r.Description
		}

		if r.ExpiresOn != nil {
			swiftPassword["expires_on"] = *r.ExpiresOn
		}

		if r.Id != nil {
			swiftPassword["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			swiftPassword["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
		}

		if r.Password != nil {
			swiftPassword["password"] = *r.Password
		}

		swiftPassword["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			swiftPassword["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, swiftPassword)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentitySwiftPasswordsDataSource().Schema["passwords"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("passwords", resources); err != nil {
		return err
	}

	return nil
}
