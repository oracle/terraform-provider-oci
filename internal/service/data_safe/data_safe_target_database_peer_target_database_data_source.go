// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeTargetDatabasePeerTargetDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["peer_target_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["target_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeTargetDatabasePeerTargetDatabaseResource(), fieldMap, readSingularDataSafeTargetDatabasePeerTargetDatabase)
}

func readSingularDataSafeTargetDatabasePeerTargetDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabasePeerTargetDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeTargetDatabasePeerTargetDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetPeerTargetDatabaseResponse
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseDataSourceCrud) Get() error {
	request := oci_data_safe.GetPeerTargetDatabaseRequest{}

	if peerTargetDatabaseId, ok := s.D.GetOkExists("peer_target_database_id"); ok {
		tmp := peerTargetDatabaseId.(int)
		request.PeerTargetDatabaseId = &tmp
	}

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetPeerTargetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeTargetDatabasePeerTargetDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeTargetDatabasePeerTargetDatabaseDataSource-", DataSafeTargetDatabasePeerTargetDatabaseDataSource(), s.D))

	if s.Res.DatabaseDetails != nil {
		databaseDetailsArray := []interface{}{}
		if databaseDetailsMap := DatabaseDetailsToMap(&s.Res.DatabaseDetails); databaseDetailsMap != nil {
			databaseDetailsArray = append(databaseDetailsArray, databaseDetailsMap)
		}
		s.D.Set("database_details", databaseDetailsArray)
	} else {
		s.D.Set("database_details", nil)
	}

	if s.Res.DatabaseUniqueName != nil {
		s.D.Set("database_unique_name", *s.Res.DatabaseUniqueName)
	}

	if s.Res.DataguardAssociationId != nil {
		s.D.Set("dataguard_association_id", *s.Res.DataguardAssociationId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Role != nil {
		s.D.Set("role", *s.Res.Role)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TlsConfig != nil {
		s.D.Set("tls_config", []interface{}{TlsConfigToMap(s.Res.TlsConfig)})
	} else {
		s.D.Set("tls_config", nil)
	}

	return nil
}
