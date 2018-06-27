// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func IdentityPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityPolicies,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     PolicyResource(),
			},
		},
	}
}

func readIdentityPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

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

func (s *IdentityPoliciesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		policy := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			policy["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			policy["description"] = *r.Description
		}

		policy["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			policy["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			policy["inactive_state"] = *r.InactiveStatus
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
			policy["version_date"] = *r.VersionDate
		}

		resources = append(resources, policy)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityPoliciesDataSource().Schema["policies"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("policies", resources); err != nil {
		panic(err)
	}

	return
}
