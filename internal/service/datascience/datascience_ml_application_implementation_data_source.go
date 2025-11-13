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

func DatascienceMlApplicationImplementationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ml_application_implementation_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatascienceMlApplicationImplementationResource(), fieldMap, readSingularDatascienceMlApplicationImplementation)
}

func readSingularDatascienceMlApplicationImplementation(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationImplementationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceMlApplicationImplementationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetMlApplicationImplementationResponse
}

func (s *DatascienceMlApplicationImplementationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceMlApplicationImplementationDataSourceCrud) Get() error {
	request := oci_datascience.GetMlApplicationImplementationRequest{}

	if mlApplicationImplementationId, ok := s.D.GetOkExists("ml_application_implementation_id"); ok {
		tmp := mlApplicationImplementationId.(string)
		request.MlApplicationImplementationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetMlApplicationImplementation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceMlApplicationImplementationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("allowed_migration_destinations", s.Res.AllowedMigrationDestinations)

	applicationComponents := []interface{}{}
	for _, item := range s.Res.ApplicationComponents {
		applicationComponents = append(applicationComponents, ApplicationComponentToMap(item))
	}
	s.D.Set("application_components", applicationComponents)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	configurationSchema := []interface{}{}
	for _, item := range s.Res.ConfigurationSchema {
		configurationSchema = append(configurationSchema, ConfigurationPropertySchemaToMap(item))
	}
	s.D.Set("configuration_schema", configurationSchema)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Logging != nil {
		s.D.Set("logging", []interface{}{ImplementationLoggingToMap(s.Res.Logging)})
	} else {
		s.D.Set("logging", nil)
	}

	if s.Res.MlApplicationId != nil {
		s.D.Set("ml_application_id", *s.Res.MlApplicationId)
	}

	if s.Res.MlApplicationName != nil {
		s.D.Set("ml_application_name", *s.Res.MlApplicationName)
	}

	if s.Res.MlApplicationPackageArguments != nil {
		s.D.Set("ml_application_package_arguments", []interface{}{MlApplicationPackageArgumentsToMap(s.Res.MlApplicationPackageArguments)})
	} else {
		s.D.Set("ml_application_package_arguments", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PackageVersion != nil {
		s.D.Set("package_version", *s.Res.PackageVersion)
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
