// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityPolicies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityPolicyResource()),
			},
		},
	}
}

func readIdentityPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListPoliciesResponse
}

func (s *IdentityPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityPoliciesDataSourceCrud) Get() error {
	request := oci_identity.ListPoliciesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_identity.PolicyLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityPoliciesDataSource-", IdentityPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		policy := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			policy["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			policy["description"] = *r.Description
		}

		policy["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			policy["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			policy["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
		}

		if r.Name != nil {
			policy["name"] = *r.Name
		}

		policy["state"] = r.LifecycleState

		policy["statements"] = r.Statements

		if r.TimeCreated != nil {
			policy["time_created"] = r.TimeCreated.String()
		}

		// TODO: see comment "pending spec/sdk versionDate solution" in identity_policy_resource.go
		if r.VersionDate != nil {
			policy["version_date"] = r.VersionDate.String()
		}

		resources = append(resources, policy)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityPoliciesDataSource().Schema["policies"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("policies", resources); err != nil {
		return err
	}

	return nil
}
