// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func CompartmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCompartments,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     CompartmentResource(),
			},
		},
	}
}

func readCompartments(d *schema.ResourceData, m interface{}) error {
	sync := &CompartmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type CompartmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListCompartmentsResponse
}

func (s *CompartmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CompartmentsDataSourceCrud) Get() error {
	request := oci_identity.ListCompartmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

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

func (s *CompartmentsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		compartment := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			compartment["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			compartment["description"] = *r.Description
		}

		compartment["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			compartment["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			compartment["inactive_state"] = *r.InactiveStatus
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
		resources = ApplyFilters(f.(*schema.Set), resources, CompartmentsDataSource().Schema["compartments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("compartments", resources); err != nil {
		panic(err)
	}

	return
}
