// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceMlApplicationInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ml_application_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatascienceMlApplicationInstanceResource(), fieldMap, readSingularDatascienceMlApplicationInstance)
}

func readSingularDatascienceMlApplicationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceMlApplicationInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetMlApplicationInstanceResponse
}

func (s *DatascienceMlApplicationInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceMlApplicationInstanceDataSourceCrud) Get() error {
	request := oci_datascience.GetMlApplicationInstanceRequest{}

	if mlApplicationInstanceId, ok := s.D.GetOkExists("ml_application_instance_id"); ok {
		tmp := mlApplicationInstanceId.(string)
		request.MlApplicationInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetMlApplicationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceMlApplicationInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AuthConfiguration != nil {
		authConfigurationArray := []interface{}{}
		if authConfigurationMap := AuthConfigurationToMap(&s.Res.AuthConfiguration); authConfigurationMap != nil {
			authConfigurationArray = append(authConfigurationArray, authConfigurationMap)
		}
		s.D.Set("auth_configuration", authConfigurationArray)
	} else {
		s.D.Set("auth_configuration", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	configuration := []interface{}{}
	for _, item := range s.Res.Configuration {
		configuration = append(configuration, ConfigurationPropertyToMap(item))
	}
	s.D.Set("configuration", configuration)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_substate", s.Res.LifecycleSubstate)

	if s.Res.MlApplicationId != nil {
		s.D.Set("ml_application_id", *s.Res.MlApplicationId)
	}

	if s.Res.MlApplicationImplementationId != nil {
		s.D.Set("ml_application_implementation_id", *s.Res.MlApplicationImplementationId)
	}

	if s.Res.MlApplicationImplementationName != nil {
		s.D.Set("ml_application_implementation_name", *s.Res.MlApplicationImplementationName)
	}

	if s.Res.MlApplicationName != nil {
		s.D.Set("ml_application_name", *s.Res.MlApplicationName)
	}

	if s.Res.PredictionEndpointDetails != nil {
		s.D.Set("prediction_endpoint_details", []interface{}{PredictionEndpointDetailsToMap(s.Res.PredictionEndpointDetails)})
	} else {
		s.D.Set("prediction_endpoint_details", nil)
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
