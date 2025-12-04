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

func DatascienceMlApplicationImplementationVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatascienceMlApplicationImplementationVersion,
		Schema: map[string]*schema.Schema{
			"ml_application_implementation_version_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"allowed_migration_destinations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"application_components": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"application_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"component_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"job_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pipeline_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"configuration_schema": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"default_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_mandatory": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"key_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sample_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"validation_regexp": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ml_application_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ml_application_implementation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ml_application_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ml_application_package_arguments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"arguments": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_mandatory": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"package_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatascienceMlApplicationImplementationVersion(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationImplementationVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceMlApplicationImplementationVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetMlApplicationImplementationVersionResponse
}

func (s *DatascienceMlApplicationImplementationVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceMlApplicationImplementationVersionDataSourceCrud) Get() error {
	request := oci_datascience.GetMlApplicationImplementationVersionRequest{}

	if mlApplicationImplementationVersionId, ok := s.D.GetOkExists("ml_application_implementation_version_id"); ok {
		tmp := mlApplicationImplementationVersionId.(string)
		request.MlApplicationImplementationVersionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetMlApplicationImplementationVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceMlApplicationImplementationVersionDataSourceCrud) SetData() error {
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

	if s.Res.MlApplicationId != nil {
		s.D.Set("ml_application_id", *s.Res.MlApplicationId)
	}

	if s.Res.MlApplicationImplementationId != nil {
		s.D.Set("ml_application_implementation_id", *s.Res.MlApplicationImplementationId)
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

	return nil
}
