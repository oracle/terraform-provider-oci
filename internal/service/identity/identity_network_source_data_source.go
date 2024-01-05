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

func IdentityNetworkSourceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_source_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityNetworkSourceResource(), fieldMap, readSingularIdentityNetworkSource)
}

func readSingularIdentityNetworkSource(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityNetworkSourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityNetworkSourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetNetworkSourceResponse
}

func (s *IdentityNetworkSourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityNetworkSourceDataSourceCrud) Get() error {
	request := oci_identity.GetNetworkSourceRequest{}

	if networkSourceId, ok := s.D.GetOkExists("network_source_id"); ok {
		tmp := networkSourceId.(string)
		request.NetworkSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.GetNetworkSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityNetworkSourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("public_source_list", s.Res.PublicSourceList)

	s.D.Set("services", s.Res.Services)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	virtualSourceList := []interface{}{}
	for _, item := range s.Res.VirtualSourceList {
		virtualSourceList = append(virtualSourceList, networkSourcesVirtualSourceListToMap(item))
	}
	s.D.Set("virtual_source_list", virtualSourceList)

	return nil
}
