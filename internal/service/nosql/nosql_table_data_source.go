// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package nosql

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_nosql "github.com/oracle/oci-go-sdk/v56/nosql"
)

func NosqlTableDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["table_name_or_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NosqlTableResource(), fieldMap, readSingularNosqlTable)
}

func readSingularNosqlTable(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlTableDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.ReadResource(sync)
}

type NosqlTableDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_nosql.NosqlClient
	Res    *oci_nosql.GetTableResponse
}

func (s *NosqlTableDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NosqlTableDataSourceCrud) Get() error {
	request := oci_nosql.GetTableRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "nosql")

	response, err := s.Client.GetTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NosqlTableDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DdlStatement != nil {
		s.D.Set("ddl_statement", *s.Res.DdlStatement)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoReclaimable != nil {
		s.D.Set("is_auto_reclaimable", *s.Res.IsAutoReclaimable)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Schema != nil {
		s.D.Set("schema", []interface{}{SchemaToMap(s.Res.Schema)})
	} else {
		s.D.Set("schema", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TableLimits != nil {
		s.D.Set("table_limits", []interface{}{TableLimitsToMap(s.Res.TableLimits)})
	} else {
		s.D.Set("table_limits", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfExpiration != nil {
		s.D.Set("time_of_expiration", s.Res.TimeOfExpiration.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
