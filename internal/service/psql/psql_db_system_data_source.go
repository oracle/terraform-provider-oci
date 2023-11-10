// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlDbSystemDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["db_system_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["excluded_fields"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(PsqlDbSystemResource(), fieldMap, readSingularPsqlDbSystem)
}

func readSingularPsqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlDbSystemDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

type PsqlDbSystemDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psql.PostgresqlClient
	Res    *oci_psql.GetDbSystemResponse
}

func (s *PsqlDbSystemDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsqlDbSystemDataSourceCrud) Get() error {
	request := oci_psql.GetDbSystemRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if excludedFields, ok := s.D.GetOkExists("excluded_fields"); ok {
		interfaces := excludedFields.([]interface{})
		tmp := make([]oci_psql.GetDbSystemExcludedFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_psql.GetDbSystemExcludedFieldsEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("excluded_fields") {
			request.ExcludedFields = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psql")

	response, err := s.Client.GetDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PsqlDbSystemDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdminUsername != nil {
		s.D.Set("admin_username", *s.Res.AdminUsername)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigId != nil {
		s.D.Set("config_id", *s.Res.ConfigId)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
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
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceCount != nil {
		s.D.Set("instance_count", *s.Res.InstanceCount)
	}

	if s.Res.InstanceMemorySizeInGBs != nil {
		s.D.Set("instance_memory_size_in_gbs", *s.Res.InstanceMemorySizeInGBs)
	}

	if s.Res.InstanceOcpuCount != nil {
		s.D.Set("instance_ocpu_count", *s.Res.InstanceOcpuCount)
	}

	instances := []interface{}{}
	for _, item := range s.Res.Instances {
		instances = append(instances, DbInstanceToMap(item))
	}
	s.D.Set("instances", instances)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ManagementPolicy != nil {
		s.D.Set("management_policy", []interface{}{ManagementPolicyToMap(s.Res.ManagementPolicy)})
	} else {
		s.D.Set("management_policy", nil)
	}

	if s.Res.NetworkDetails != nil {
		s.D.Set("network_details", []interface{}{NetworkDetailsToMap(s.Res.NetworkDetails, true)})
	} else {
		s.D.Set("network_details", nil)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	if s.Res.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := SourceDetailsToMap(&s.Res.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		s.D.Set("source", sourceArray)
	} else {
		s.D.Set("source", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageDetails != nil {
		storageDetailsArray := []interface{}{}
		if storageDetailsMap := StorageDetailsToMap(&s.Res.StorageDetails); storageDetailsMap != nil {
			storageDetailsArray = append(storageDetailsArray, storageDetailsMap)
		}
		s.D.Set("storage_details", storageDetailsArray)
	} else {
		s.D.Set("storage_details", nil)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("system_type", s.Res.SystemType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
