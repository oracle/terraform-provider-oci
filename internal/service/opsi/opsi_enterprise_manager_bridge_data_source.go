// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v56/opsi"
)

func OpsiEnterpriseManagerBridgeDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["enterprise_manager_bridge_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpsiEnterpriseManagerBridgeResource(), fieldMap, readSingularOpsiEnterpriseManagerBridge)
}

func readSingularOpsiEnterpriseManagerBridge(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiEnterpriseManagerBridgeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiEnterpriseManagerBridgeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetEnterpriseManagerBridgeResponse
}

func (s *OpsiEnterpriseManagerBridgeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiEnterpriseManagerBridgeDataSourceCrud) Get() error {
	request := oci_opsi.GetEnterpriseManagerBridgeRequest{}

	if enterpriseManagerBridgeId, ok := s.D.GetOkExists("enterprise_manager_bridge_id"); ok {
		tmp := enterpriseManagerBridgeId.(string)
		request.EnterpriseManagerBridgeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetEnterpriseManagerBridge(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiEnterpriseManagerBridgeDataSourceCrud) SetData() error {
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ObjectStorageBucketName != nil {
		s.D.Set("object_storage_bucket_name", *s.Res.ObjectStorageBucketName)
	}

	if s.Res.ObjectStorageBucketStatusDetails != nil {
		s.D.Set("object_storage_bucket_status_details", *s.Res.ObjectStorageBucketStatusDetails)
	}

	if s.Res.ObjectStorageNamespaceName != nil {
		s.D.Set("object_storage_namespace_name", *s.Res.ObjectStorageNamespaceName)
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
