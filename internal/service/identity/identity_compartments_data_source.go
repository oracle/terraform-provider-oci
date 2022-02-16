// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityCompartmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityCompartments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityCompartmentResource()),
			},
		},
	}
}

func readIdentityCompartments(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCompartmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityCompartmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListCompartmentsResponse
}

func (s *IdentityCompartmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityCompartmentsDataSourceCrud) Get() error {
	request := oci_identity.ListCompartmentsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_identity.ListCompartmentsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_identity.CompartmentLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListCompartments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCompartments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityCompartmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityCompartmentsDataSource-", IdentityCompartmentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		compartment := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			compartment["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			compartment["description"] = *r.Description
		}

		compartment["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			compartment["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			compartment["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
		}

		if r.IsAccessible != nil {
			compartment["is_accessible"] = *r.IsAccessible
		}

		if r.Name != nil {
			compartment["name"] = *r.Name
		}

		compartment["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			compartment["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, compartment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityCompartmentsDataSource().Schema["compartments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("compartments", resources); err != nil {
		return err
	}

	return nil
}
