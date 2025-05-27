// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentNamedCredentialsMetadataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularManagementAgentManagementAgentNamedCredentialsMetadata,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_agent_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum_agent_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"properties": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"allowed_values": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"default_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_required": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"regex": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value_category": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularManagementAgentManagementAgentNamedCredentialsMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentNamedCredentialsMetadataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentNamedCredentialsMetadataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.GetNamedCredentialsMetadatumResponse
}

func (s *ManagementAgentManagementAgentNamedCredentialsMetadataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentNamedCredentialsMetadataDataSourceCrud) Get() error {
	request := oci_management_agent.GetNamedCredentialsMetadatumRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.GetNamedCredentialsMetadatum(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagementAgentManagementAgentNamedCredentialsMetadataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentManagementAgentNamedCredentialsMetadataDataSource-", ManagementAgentManagementAgentNamedCredentialsMetadataDataSource(), s.D))

	metadata := []interface{}{}
	for _, item := range s.Res.Metadata {
		metadata = append(metadata, NamedCredentialMetadataDefinitionToMap(item))
	}
	s.D.Set("metadata", metadata)

	return nil
}

func NamedCredentialFieldDefinitionToMap(obj oci_management_agent.NamedCredentialFieldDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_values"] = obj.AllowedValues

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.IsRequired != nil {
		result["is_required"] = bool(*obj.IsRequired)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Regex != nil {
		result["regex"] = string(*obj.Regex)
	}

	result["value_category"] = obj.ValueCategory

	return result
}

func NamedCredentialMetadataDefinitionToMap(obj oci_management_agent.NamedCredentialMetadataDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.MinimumAgentVersion != nil {
		result["minimum_agent_version"] = string(*obj.MinimumAgentVersion)
	}

	properties := []interface{}{}
	for _, item := range obj.Properties {
		properties = append(properties, NamedCredentialFieldDefinitionToMap(item))
	}
	result["properties"] = properties

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}
