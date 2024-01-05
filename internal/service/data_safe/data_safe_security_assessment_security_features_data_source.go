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

func DataSafeSecurityAssessmentSecurityFeaturesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityAssessmentSecurityFeatures,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_column_encryption": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_database_vault": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_external_authentication": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_fine_grained_audit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_global_authentication": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_network_encryption": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_password_authentication": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_privilege_analysis": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_tablespace_encryption": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_traditional_audit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"targets_with_unified_audit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_feature_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"column_encryption": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_vault": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"external_authentication": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fine_grained_audit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"global_authentication": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_encryption": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"password_authentication": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"privilege_analysis": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tablespace_encryption": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"traditional_audit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unified_audit": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDataSafeSecurityAssessmentSecurityFeatures(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentSecurityFeaturesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentSecurityFeaturesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSecurityFeaturesResponse
}

func (s *DataSafeSecurityAssessmentSecurityFeaturesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentSecurityFeaturesDataSourceCrud) Get() error {
	request := oci_data_safe.ListSecurityFeaturesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSecurityFeaturesAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if targetsWithColumnEncryption, ok := s.D.GetOkExists("targets_with_column_encryption"); ok {
		request.TargetsWithColumnEncryption = oci_data_safe.ListSecurityFeaturesTargetsWithColumnEncryptionEnum(targetsWithColumnEncryption.(string))
	}

	if targetsWithDatabaseVault, ok := s.D.GetOkExists("targets_with_database_vault"); ok {
		request.TargetsWithDatabaseVault = oci_data_safe.ListSecurityFeaturesTargetsWithDatabaseVaultEnum(targetsWithDatabaseVault.(string))
	}

	if targetsWithExternalAuthentication, ok := s.D.GetOkExists("targets_with_external_authentication"); ok {
		request.TargetsWithExternalAuthentication = oci_data_safe.ListSecurityFeaturesTargetsWithExternalAuthenticationEnum(targetsWithExternalAuthentication.(string))
	}

	if targetsWithFineGrainedAudit, ok := s.D.GetOkExists("targets_with_fine_grained_audit"); ok {
		request.TargetsWithFineGrainedAudit = oci_data_safe.ListSecurityFeaturesTargetsWithFineGrainedAuditEnum(targetsWithFineGrainedAudit.(string))
	}

	if targetsWithGlobalAuthentication, ok := s.D.GetOkExists("targets_with_global_authentication"); ok {
		request.TargetsWithGlobalAuthentication = oci_data_safe.ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum(targetsWithGlobalAuthentication.(string))
	}

	if targetsWithNetworkEncryption, ok := s.D.GetOkExists("targets_with_network_encryption"); ok {
		request.TargetsWithNetworkEncryption = oci_data_safe.ListSecurityFeaturesTargetsWithNetworkEncryptionEnum(targetsWithNetworkEncryption.(string))
	}

	if targetsWithPasswordAuthentication, ok := s.D.GetOkExists("targets_with_password_authentication"); ok {
		request.TargetsWithPasswordAuthentication = oci_data_safe.ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum(targetsWithPasswordAuthentication.(string))
	}

	if targetsWithPrivilegeAnalysis, ok := s.D.GetOkExists("targets_with_privilege_analysis"); ok {
		request.TargetsWithPrivilegeAnalysis = oci_data_safe.ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum(targetsWithPrivilegeAnalysis.(string))
	}

	if targetsWithTablespaceEncryption, ok := s.D.GetOkExists("targets_with_tablespace_encryption"); ok {
		request.TargetsWithTablespaceEncryption = oci_data_safe.ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum(targetsWithTablespaceEncryption.(string))
	}

	if targetsWithTraditionalAudit, ok := s.D.GetOkExists("targets_with_traditional_audit"); ok {
		request.TargetsWithTraditionalAudit = oci_data_safe.ListSecurityFeaturesTargetsWithTraditionalAuditEnum(targetsWithTraditionalAudit.(string))
	}

	if targetsWithUnifiedAudit, ok := s.D.GetOkExists("targets_with_unified_audit"); ok {
		request.TargetsWithUnifiedAudit = oci_data_safe.ListSecurityFeaturesTargetsWithUnifiedAuditEnum(targetsWithUnifiedAudit.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSecurityFeatures(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSecurityFeatures(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityAssessmentSecurityFeaturesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityAssessmentSecurityFeaturesDataSource-", DataSafeSecurityAssessmentSecurityFeaturesDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityAssessmentSecurityFeature := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SecurityFeatureSummaryToMap(item))
	}
	securityAssessmentSecurityFeature["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSecurityAssessmentSecurityFeaturesDataSource().Schema["security_feature_collection"].Elem.(*schema.Resource).Schema)
		securityAssessmentSecurityFeature["items"] = items
	}

	resources = append(resources, securityAssessmentSecurityFeature)
	if err := s.D.Set("security_feature_collection", resources); err != nil {
		return err
	}

	return nil
}

func SecurityFeatureSummaryToMap(obj oci_data_safe.SecurityFeatureSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssessmentId != nil {
		result["assessment_id"] = string(*obj.AssessmentId)
	}

	result["column_encryption"] = string(obj.ColumnEncryption)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["database_vault"] = string(obj.DatabaseVault)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["external_authentication"] = string(obj.ExternalAuthentication)

	result["fine_grained_audit"] = string(obj.FineGrainedAudit)

	result["freeform_tags"] = obj.FreeformTags

	result["global_authentication"] = string(obj.GlobalAuthentication)

	result["network_encryption"] = string(obj.NetworkEncryption)

	result["password_authentication"] = string(obj.PasswordAuthentication)

	result["privilege_analysis"] = string(obj.PrivilegeAnalysis)

	result["tablespace_encryption"] = string(obj.TablespaceEncryption)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	result["traditional_audit"] = string(obj.TraditionalAudit)

	result["unified_audit"] = string(obj.UnifiedAudit)

	return result
}
