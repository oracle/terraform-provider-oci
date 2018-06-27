// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func SwiftPasswordsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSwiftPasswords,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"passwords": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     SwiftPasswordResource(),
			},
		},
	}
}

func readSwiftPasswords(d *schema.ResourceData, m interface{}) error {
	sync := &SwiftPasswordsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type SwiftPasswordsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListSwiftPasswordsResponse
}

func (s *SwiftPasswordsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SwiftPasswordsDataSourceCrud) Get() error {
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

func (s *SwiftPasswordsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
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
			swiftPassword["inactive_state"] = *r.InactiveStatus
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
		resources = ApplyFilters(f.(*schema.Set), resources, SwiftPasswordsDataSource().Schema["passwords"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("passwords", resources); err != nil {
		panic(err)
	}

	return
}
