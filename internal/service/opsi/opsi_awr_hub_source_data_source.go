// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiAwrHubSourceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["awr_hub_source_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpsiAwrHubSourceResource(), fieldMap, readSingularOpsiAwrHubSource)
}

func readSingularOpsiAwrHubSource(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubSourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiAwrHubSourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetAwrHubSourceResponse
}

func (s *OpsiAwrHubSourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiAwrHubSourceDataSourceCrud) Get() error {
	request := oci_opsi.GetAwrHubSourceRequest{}

	if awrHubSourceId, ok := s.D.GetOkExists("awr_hub_source_id"); ok {
		tmp := awrHubSourceId.(string)
		request.AwrHubSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetAwrHubSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiAwrHubSourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AssociatedOpsiId != nil {
		s.D.Set("associated_opsi_id", *s.Res.AssociatedOpsiId)
	}

	if s.Res.AssociatedResourceId != nil {
		s.D.Set("associated_resource_id", *s.Res.AssociatedResourceId)
	}

	if s.Res.AwrHubId != nil {
		s.D.Set("awr_hub_id", *s.Res.AwrHubId)
	}

	if s.Res.AwrHubOpsiSourceId != nil {
		s.D.Set("awr_hub_opsi_source_id", *s.Res.AwrHubOpsiSourceId)
	}

	if s.Res.AwrSourceDatabaseId != nil {
		s.D.Set("awr_source_database_id", *s.Res.AwrSourceDatabaseId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HoursSinceLastImport != nil {
		s.D.Set("hours_since_last_import", *s.Res.HoursSinceLastImport)
	}

	if s.Res.IsRegisteredWithAwrHub != nil {
		s.D.Set("is_registered_with_awr_hub", *s.Res.IsRegisteredWithAwrHub)
	}

	if s.Res.MaxSnapshotIdentifier != nil {
		s.D.Set("max_snapshot_identifier", *s.Res.MaxSnapshotIdentifier)
	}

	if s.Res.MinSnapshotIdentifier != nil {
		s.D.Set("min_snapshot_identifier", *s.Res.MinSnapshotIdentifier)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.SourceMailBoxUrl != nil {
		s.D.Set("source_mail_box_url", *s.Res.SourceMailBoxUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeFirstSnapshotGenerated != nil {
		s.D.Set("time_first_snapshot_generated", s.Res.TimeFirstSnapshotGenerated.String())
	}

	if s.Res.TimeLastSnapshotGenerated != nil {
		s.D.Set("time_last_snapshot_generated", s.Res.TimeLastSnapshotGenerated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
