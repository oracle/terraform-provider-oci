// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v58/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CloudGuardManagedListDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["managed_list_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudGuardManagedListResource(), fieldMap, readSingularCloudGuardManagedList)
}

func readSingularCloudGuardManagedList(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardManagedListDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardManagedListDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetManagedListResponse
}

func (s *CloudGuardManagedListDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardManagedListDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetManagedListRequest{}

	if managedListId, ok := s.D.GetOkExists("managed_list_id"); ok {
		tmp := managedListId.(string)
		request.ManagedListId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetManagedList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardManagedListDataSourceCrud) SetData() error {
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("feed_provider", s.Res.FeedProvider)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEditable != nil {
		s.D.Set("is_editable", *s.Res.IsEditable)
	}

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	s.D.Set("list_items", s.Res.ListItems)

	s.D.Set("list_type", s.Res.ListType)

	if s.Res.SourceManagedListId != nil {
		s.D.Set("source_managed_list_id", *s.Res.SourceManagedListId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
