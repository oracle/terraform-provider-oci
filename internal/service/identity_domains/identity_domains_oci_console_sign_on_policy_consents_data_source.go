// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsOciConsoleSignOnPolicyConsentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomainsOciConsoleSignOnPolicyConsents,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"oci_console_sign_on_policy_consent_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oci_console_sign_on_policy_consent_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Computed
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"client_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"change_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"justification": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"notification_recipients": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"time_consent_signed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reason": {
							Type:     schema.TypeString,
							Computed: true,
						},
						// Optional
						"tenancy_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"compartment_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"consent_signed_by": {
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
									"ocid": {
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

						"domain_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"idcs_created_by": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ocid": {
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

						"idcs_last_modified_by": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ocid": {
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

						"idcs_last_upgraded_in_release": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_prevented_operations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"meta": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_modified": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"location": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"modified_resource": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"ocid": {
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

						"ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"policy_resource": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"ocid": {
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
						// Computed
					},
				},
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"items_per_page": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"key": {
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
			"total_results": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readIdentityDomainsOciConsoleSignOnPolicyConsents(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsOciConsoleSignOnPolicyConsentsDataSourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

type IdentityDomainsOciConsoleSignOnPolicyConsentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.ListOciConsoleSignOnPolicyConsentsResponse
}

func (s *IdentityDomainsOciConsoleSignOnPolicyConsentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsOciConsoleSignOnPolicyConsentsDataSourceCrud) Get() error {
	request := oci_identity_domains.ListOciConsoleSignOnPolicyConsentsRequest{}

	if ociConsoleSignOnPolicyConsentCount, ok := s.D.GetOkExists("oci_console_sign_on_policy_consent_count"); ok {
		tmp := ociConsoleSignOnPolicyConsentCount.(int)
		request.Count = &tmp
	}

	if ociConsoleSignOnPolicyConsentFilter, ok := s.D.GetOkExists("oci_console_sign_on_policy_consent_filter"); ok {
		tmp := ociConsoleSignOnPolicyConsentFilter.(string)
		request.Filter = &tmp
	}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if startIndex, ok := s.D.GetOkExists("start_index"); ok {
		tmp := startIndex.(int)
		request.StartIndex = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.ListOciConsoleSignOnPolicyConsents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsOciConsoleSignOnPolicyConsentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsOciConsoleSignOnPolicyConsentsDataSource-", IdentityDomainsOciConsoleSignOnPolicyConsentsDataSource(), s.D))

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, OciConsoleSignOnPolicyConsentToMap(item))
	}
	s.D.Set("resources", resources)

	if s.Res.ItemsPerPage != nil {
		s.D.Set("items_per_page", *s.Res.ItemsPerPage)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.StartIndex != nil {
		s.D.Set("start_index", *s.Res.StartIndex)
	}

	if s.Res.TotalResults != nil {
		s.D.Set("total_results", *s.Res.TotalResults)
	}

	return nil
}

func OciConsoleSignOnPolicyConsentToMap(obj oci_identity_domains.OciConsoleSignOnPolicyConsent) map[string]interface{} {
	result := map[string]interface{}{}

	log.Printf("[DEBUG] Skip sweeper for %s", string(*obj.ClientIp))
	log.Printf("[DEBUG] Skip sweeper for %s", string(*obj.Justification))

	result["change_type"] = string(obj.ChangeType)

	if obj.ClientIp != nil {
		result["client_ip"] = string(*obj.ClientIp)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Justification != nil {
		result["justification"] = string(*obj.Justification)
	}

	if obj.Reason != nil {
		result["reason"] = string(*obj.Reason)
	}

	if obj.TimeConsentSigned != nil {
		result["time_consent_signed"] = string(*obj.TimeConsentSigned)
	}

	result["notification_recipients"] = obj.NotificationRecipients

	if obj.ConsentSignedBy != nil {
		result["consent_signed_by"] = []interface{}{OciConsoleSignOnPolicyConsentConsentSignedByToMap(obj.ConsentSignedBy)}
	}

	if obj.ModifiedResource != nil {
		result["modified_resource"] = []interface{}{OciConsoleSignOnPolicyConsentModifiedResourceToMap(obj.ModifiedResource)}
	}

	if obj.PolicyResource != nil {
		result["policy_resource"] = []interface{}{OciConsoleSignOnPolicyConsentPolicyResourceToMap(obj.PolicyResource)}
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	return result
}

func OciConsoleSignOnPolicyConsentConsentSignedByToMap(obj *oci_identity_domains.OciConsoleSignOnPolicyConsentConsentSignedBy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func OciConsoleSignOnPolicyConsentModifiedResourceToMap(obj *oci_identity_domains.OciConsoleSignOnPolicyConsentModifiedResource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func OciConsoleSignOnPolicyConsentPolicyResourceToMap(obj *oci_identity_domains.OciConsoleSignOnPolicyConsentPolicyResource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
