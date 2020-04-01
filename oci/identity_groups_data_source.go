// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func init() {
	RegisterDatasource("oci_identity_groups", IdentityGroupsDataSource())
}

func IdentityGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityGroups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(IdentityGroupResource()),
			},
		},
	}
}

func readIdentityGroups(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type IdentityGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListGroupsResponse
}

func (s *IdentityGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityGroupsDataSourceCrud) Get() error {
	request := oci_identity.ListGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		group := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			group["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			group["description"] = *r.Description
		}

		group["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			group["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			group["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
		}

		if r.Name != nil {
			group["name"] = *r.Name
		}

		group["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			group["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, group)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityGroupsDataSource().Schema["groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("groups", resources); err != nil {
		return err
	}

	return nil
}
