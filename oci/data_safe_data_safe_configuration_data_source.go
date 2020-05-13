// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/datasafe"
)

func init() {
	RegisterDatasource("oci_data_safe_data_safe_configuration", DataSafeDataSafeConfigurationDataSource())
}

func DataSafeDataSafeConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DataSafeDataSafeConfigurationResource(), fieldMap, readSingularDataSafeDataSafeConfiguration)
}

func readSingularDataSafeDataSafeConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDataSafeConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataSafeClient()

	return ReadResource(sync)
}

type DataSafeDataSafeConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetDataSafeConfigurationResponse
}

func (s *DataSafeDataSafeConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeDataSafeConfigurationDataSourceCrud) Get() error {
	request := oci_data_safe.GetDataSafeConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "data_safe")

	response, err := s.Client.GetDataSafeConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeDataSafeConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeEnabled != nil {
		s.D.Set("time_enabled", s.Res.TimeEnabled.String())
	}

	if s.Res.Url != nil {
		s.D.Set("url", *s.Res.Url)
	}

	return nil
}
