// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_certificates_management "github.com/oracle/oci-go-sdk/v65/certificatesmanagement"
)

func CertificatesManagementCertificateAuthorityResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCertificatesManagementCertificateAuthority,
		Read:     readCertificatesManagementCertificateAuthority,
		Update:   updateCertificatesManagementCertificateAuthority,
		Delete:   deleteCertificatesManagementCertificateAuthority,
		Schema: map[string]*schema.Schema{
			// Required
			"certificate_authority_config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"config_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ROOT_CA_GENERATED_INTERNALLY",
								"SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA",
							}, true),
						},
						"subject": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"common_name": {
										Type:     schema.TypeString,
										Required: true,
										Computed: false,
										ForceNew: true,
									},
									// Optional
									"country": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"distinguished_name_qualifier": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"domain_component": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"generation_qualifier": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"given_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"initials": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"locality_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"organization": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"organizational_unit": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"pseudonym": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"serial_number": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"state_or_province_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"street": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"surname": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"title": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"user_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Optional
						"issuer_certificate_authority_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
							ForceNew: true,
						},
						"signing_algorithm": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
							ForceNew: true,
						},
						"validity": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"time_of_validity_not_after": {
										Type:             schema.TypeString,
										Required:         true,
										Computed:         false,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									// Optional
									"time_of_validity_not_before": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         false,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},

									// Computed
								},
							},
						},
						"version_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},

						// Computed
					},
				},
			},
			// end certificate_authority_config
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"certificate_authority_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"rule_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE",
							}, true),
						},

						// Optional
						"certificate_authority_max_validity_duration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"leaf_certificate_max_validity_duration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},

						// Computed
					},
				},
			},
			"certificate_revocation_list_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: false,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"object_storage_config": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"object_storage_bucket_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"object_storage_object_name_format": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"object_storage_namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Optional
						"custom_formatted_urls": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"config_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_version": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"certificate_authority_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"issuer_ca_version_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"revocation_status": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"revocation_reason": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_of_revocation": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"stages": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_deletion": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"validity": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"time_of_validity_not_after": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_of_validity_not_before": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			// end current_version
			"issuer_certificate_authority_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signing_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subject": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"common_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"distinguished_name_qualifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain_component": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"generation_qualifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"given_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"initials": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"locality_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organization": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organizational_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pseudonym": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state_or_province_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"street": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"surname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCertificatesManagementCertificateAuthority(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateAuthorityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readCertificatesManagementCertificateAuthority(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateAuthorityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

func updateCertificatesManagementCertificateAuthority(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateAuthorityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCertificatesManagementCertificateAuthority(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateAuthorityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.DeleteResource(d, sync)
}

type CertificatesManagementCertificateAuthorityResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_certificates_management.CertificatesManagementClient
	Res                    *oci_certificates_management.CertificateAuthority
	DisableNotFoundRetries bool
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_certificates_management.CertificateAuthorityLifecycleStateCreating),
	}
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_certificates_management.CertificateAuthorityLifecycleStateActive),
	}
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_certificates_management.CertificateAuthorityLifecycleStateUpdating),
	}
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_certificates_management.CertificateAuthorityLifecycleStateActive),
	}
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_certificates_management.CertificateAuthorityLifecycleStateDeleting),
		string(oci_certificates_management.CertificateAuthorityLifecycleStateSchedulingDeletion),
	}
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_certificates_management.CertificateAuthorityLifecycleStateDeleted),
		string(oci_certificates_management.CertificateAuthorityLifecycleStatePendingDeletion),
	}
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) Create() error {
	request := oci_certificates_management.CreateCertificateAuthorityRequest{}

	if certificateAuthorityConfig, ok := s.D.GetOkExists("certificate_authority_config"); ok {
		if tmpList := certificateAuthorityConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_authority_config", 0)
			tmp, err := s.mapToCreateCertificateAuthorityConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CertificateAuthorityConfig = tmp
		}
	}

	if certificateAuthorityRules, ok := s.D.GetOkExists("certificate_authority_rules"); ok {
		interfaces := certificateAuthorityRules.([]interface{})
		tmp := make([]oci_certificates_management.CertificateAuthorityRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_authority_rules", stateDataIndex)
			converted, err := s.mapToCertificateAuthorityRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("certificate_authority_rules") {
			request.CertificateAuthorityRules = tmp
		}
	}

	if certificateRevocationListDetails, ok := s.D.GetOkExists("certificate_revocation_list_details"); ok {
		if tmpList := certificateRevocationListDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_revocation_list_details", 0)
			tmp, err := s.mapToCertificateRevocationListDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CertificateRevocationListDetails = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	response, err := s.Client.CreateCertificateAuthority(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CertificateAuthority
	return nil
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) Get() error {
	request := oci_certificates_management.GetCertificateAuthorityRequest{}

	tmp := s.D.Id()
	request.CertificateAuthorityId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	response, err := s.Client.GetCertificateAuthority(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CertificateAuthority
	return nil
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_certificates_management.UpdateCertificateAuthorityRequest{}

	if certificateAuthorityConfig, ok := s.D.GetOkExists("certificate_authority_config"); ok && s.D.HasChange("certificate_authority_config") {
		if tmpList := certificateAuthorityConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_authority_config", 0)
			tmp, err := s.mapToUpdateCertificateAuthorityConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CertificateAuthorityConfig = tmp
		}
	}

	if certificateRevocationListDetails, ok := s.D.GetOkExists("certificate_revocation_list_details"); ok {
		if tmpList := certificateRevocationListDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_revocation_list_details", 0)
			tmp, err := s.mapToCertificateRevocationListDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			if len(tmpList) != 0 || s.D.HasChange("certificate_revocation_list_details") {
				request.CertificateRevocationListDetails = &tmp
			}
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	// only update if request has updates
	if request.CertificateAuthorityConfig != nil || request.CertificateRevocationListDetails != nil ||
		request.FreeformTags != nil || request.Description != nil || request.DefinedTags != nil {
		tmp := s.D.Id()
		request.CertificateAuthorityId = &tmp

		response, err := s.Client.UpdateCertificateAuthority(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res = &response.CertificateAuthority
	}

	// update rules
	if _, ok := s.D.GetOkExists("certificate_authority_rules"); ok && s.D.HasChange("certificate_authority_rules") {
		err := s.UpdateRules()
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) UpdateRules() error {
	request := oci_certificates_management.UpdateCertificateAuthorityRequest{}

	tmp := s.D.Id()
	request.CertificateAuthorityId = &tmp

	if certificateAuthorityRules, ok := s.D.GetOkExists("certificate_authority_rules"); ok {
		interfaces := certificateAuthorityRules.([]interface{})
		tmp := make([]oci_certificates_management.CertificateAuthorityRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_authority_rules", stateDataIndex)
			converted, err := s.mapToCertificateAuthorityRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("certificate_authority_rules") {
			request.CertificateAuthorityRules = tmp
		} else {
			return nil
		}
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	response, err := s.Client.UpdateCertificateAuthority(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CertificateAuthority
	return s.Get()
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) Delete() error {
	request := oci_certificates_management.ScheduleCertificateAuthorityDeletionRequest{}

	if timeOfDeletion, ok := s.D.GetOkExists("time_of_deletion"); ok {
		tmpTime, err := time.Parse(time.RFC3339Nano, timeOfDeletion.(string))
		if err != nil {
			return err
		}
		request.TimeOfDeletion = &oci_common.SDKTime{Time: tmpTime}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")
	tmp := s.D.Id()
	request.CertificateAuthorityId = &tmp

	_, err := s.Client.ScheduleCertificateAuthorityDeletion(context.Background(), request)
	return err
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) SetData() error {
	certificateAuthorityRules := []interface{}{}
	for _, item := range s.Res.CertificateAuthorityRules {
		certificateAuthorityRules = append(certificateAuthorityRules, CertificateAuthorityRuleToMap(item))
	}
	s.D.Set("certificate_authority_rules", certificateAuthorityRules)

	if s.Res.CertificateRevocationListDetails != nil {
		s.D.Set("certificate_revocation_list_details", []interface{}{CertificateRevocationListDetailsToMap(s.Res.CertificateRevocationListDetails)})
	} else {
		s.D.Set("certificate_revocation_list_details", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("config_type", s.Res.ConfigType)

	if s.Res.CurrentVersion != nil {
		s.D.Set("current_version", []interface{}{CertificateAuthorityVersionSummaryToMap(s.Res.CurrentVersion)})
	} else {
		s.D.Set("current_version", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IssuerCertificateAuthorityId != nil {
		s.D.Set("issuer_certificate_authority_id", *s.Res.IssuerCertificateAuthorityId)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("signing_algorithm", s.Res.SigningAlgorithm)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Subject != nil {
		s.D.Set("subject", []interface{}{CertificateSubjectToMap(s.Res.Subject)})
	} else {
		s.D.Set("subject", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	return nil
}

func CertificateAuthorityRuleToMap(obj oci_certificates_management.CertificateAuthorityRule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_certificates_management.CertificateAuthorityIssuanceExpiryRule:
		result["rule_type"] = "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"

		if v.CertificateAuthorityMaxValidityDuration != nil {
			result["certificate_authority_max_validity_duration"] = string(*v.CertificateAuthorityMaxValidityDuration)
		}

		if v.LeafCertificateMaxValidityDuration != nil {
			result["leaf_certificate_max_validity_duration"] = string(*v.LeafCertificateMaxValidityDuration)
		}
	default:
		log.Printf("[WARN] Received 'rule_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func CertificateAuthoritySummaryToMap(obj oci_certificates_management.CertificateAuthoritySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = *obj.Id
	}

	if obj.IssuerCertificateAuthorityId != nil {
		result["issuer_certificate_authority_id"] = *obj.IssuerCertificateAuthorityId
	}

	result["name"] = *obj.Name

	if obj.Description != nil {
		result["description"] = *obj.Description
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeOfDeletion != nil {
		result["time_of_deletion"] = obj.TimeOfDeletion.String()
	}

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = *obj.KmsKeyId
	}

	result["state"] = obj.LifecycleState

	result["compartment_id"] = *obj.CompartmentId

	if obj.CertificateAuthorityRules != nil {
		certificateAuthorityRules := []interface{}{}
		for _, item := range obj.CertificateAuthorityRules {
			certificateAuthorityRules = append(certificateAuthorityRules, CertificateAuthorityRuleToMap(item))
		}
		result["certificate_authority_rules"] = certificateAuthorityRules
	}

	if obj.CurrentVersionSummary != nil {
		result["current_version"] = []interface{}{CertificateAuthorityVersionSummaryToMap(obj.CurrentVersionSummary)}
	} else {
		result["current_version"] = nil
	}

	if obj.Subject != nil {
		result["subject"] = []interface{}{CertificateSubjectToMap(obj.Subject)}
	}

	result["config_type"] = string(obj.ConfigType)

	result["signing_algorithm"] = obj.SigningAlgorithm

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.FreeformTags != nil {
		result["freeform_tags"] = obj.FreeformTags

	}

	return result
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) mapToCertificateRevocationListDetails(fieldKeyFormat string) (oci_certificates_management.CertificateRevocationListDetails, error) {
	result := oci_certificates_management.CertificateRevocationListDetails{}

	if customFormattedUrls, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_formatted_urls")); ok {
		interfaces := customFormattedUrls.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "custom_formatted_urls")) {
			result.CustomFormattedUrls = tmp
		}
	}

	if objectStorageConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_config")); ok {
		if tmpList := objectStorageConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_config"), 0)
			tmp, err := s.mapToObjectStorageBucketConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert object_storage_config, encountered error: %v", err)
			}
			result.ObjectStorageConfig = &tmp
		}
	}

	return result, nil
}

func CertificateRevocationListDetailsToMap(obj *oci_certificates_management.CertificateRevocationListDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["custom_formatted_urls"] = obj.CustomFormattedUrls

	if obj.ObjectStorageConfig != nil {
		result["object_storage_config"] = []interface{}{ObjectStorageBucketConfigDetailsToMap(obj.ObjectStorageConfig)}
	}

	return result
}

func CertificateSubjectToMap(obj *oci_certificates_management.CertificateSubject) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CommonName != nil {
		result["common_name"] = string(*obj.CommonName)
	}

	if obj.Country != nil {
		result["country"] = string(*obj.Country)
	}

	if obj.DistinguishedNameQualifier != nil {
		result["distinguished_name_qualifier"] = string(*obj.DistinguishedNameQualifier)
	}

	if obj.DomainComponent != nil {
		result["domain_component"] = string(*obj.DomainComponent)
	}

	if obj.GenerationQualifier != nil {
		result["generation_qualifier"] = string(*obj.GenerationQualifier)
	}

	if obj.GivenName != nil {
		result["given_name"] = string(*obj.GivenName)
	}

	if obj.Initials != nil {
		result["initials"] = string(*obj.Initials)
	}

	if obj.LocalityName != nil {
		result["locality_name"] = string(*obj.LocalityName)
	}

	if obj.Organization != nil {
		result["organization"] = string(*obj.Organization)
	}

	if obj.OrganizationalUnit != nil {
		result["organizational_unit"] = string(*obj.OrganizationalUnit)
	}

	if obj.Pseudonym != nil {
		result["pseudonym"] = string(*obj.Pseudonym)
	}

	if obj.SerialNumber != nil {
		result["serial_number"] = string(*obj.SerialNumber)
	}

	if obj.StateOrProvinceName != nil {
		result["state_or_province_name"] = string(*obj.StateOrProvinceName)
	}

	if obj.Street != nil {
		result["street"] = string(*obj.Street)
	}

	if obj.Surname != nil {
		result["surname"] = string(*obj.Surname)
	}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	if obj.UserId != nil {
		result["user_id"] = string(*obj.UserId)
	}

	return result
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) mapToCertificateAuthoritySubject(fieldKeyFormat string) (oci_certificates_management.CertificateSubject, error) {
	result := oci_certificates_management.CertificateSubject{}

	if commonName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "common_name")); ok {
		tmp := commonName.(string)
		result.CommonName = &tmp
	}

	if country, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "country")); ok {
		tmp := country.(string)
		result.Country = &tmp
	}

	if distinguishedNameQualifier, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "distinguished_name_qualifier")); ok {
		tmp := distinguishedNameQualifier.(string)
		result.DistinguishedNameQualifier = &tmp
	}

	if domainComponent, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "domain_component")); ok {
		tmp := domainComponent.(string)
		result.DomainComponent = &tmp
	}

	if generationQualifier, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "generation_qualifier")); ok {
		tmp := generationQualifier.(string)
		result.GenerationQualifier = &tmp
	}

	if givenName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "given_name")); ok {
		tmp := givenName.(string)
		result.GivenName = &tmp
	}

	if initials, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "initials")); ok {
		tmp := initials.(string)
		result.Initials = &tmp
	}

	if localityName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "locality_name")); ok {
		tmp := localityName.(string)
		result.LocalityName = &tmp
	}

	if organization, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "organization")); ok {
		tmp := organization.(string)
		result.Organization = &tmp
	}

	if organizationalUnit, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "organizational_unit")); ok {
		tmp := organizationalUnit.(string)
		result.OrganizationalUnit = &tmp
	}

	if pseudonym, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "pseudonym")); ok {
		tmp := pseudonym.(string)
		result.Pseudonym = &tmp
	}

	if serialNumber, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "serial_number")); ok {
		tmp := serialNumber.(string)
		result.SerialNumber = &tmp
	}

	if stateOrProvinceName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "state_or_province_name")); ok {
		tmp := stateOrProvinceName.(string)
		result.StateOrProvinceName = &tmp
	}

	if street, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "street")); ok {
		tmp := street.(string)
		result.Street = &tmp
	}

	if surname, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "surname")); ok {
		tmp := surname.(string)
		result.Surname = &tmp
	}

	if title, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "title")); ok {
		tmp := title.(string)
		result.Title = &tmp
	}

	if userId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
		tmp := userId.(string)
		result.UserId = &tmp
	}

	return result, nil
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) mapToCreateCertificateAuthorityConfigDetails(fieldKeyFormat string) (oci_certificates_management.CreateCertificateAuthorityConfigDetails, error) {
	var baseObject oci_certificates_management.CreateCertificateAuthorityConfigDetails
	//discriminator
	configTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_type"))
	var configType string
	if ok {
		configType = configTypeRaw.(string)
	} else {
		configType = "" // default value
	}
	switch strings.ToLower(configType) {
	case strings.ToLower("ROOT_CA_GENERATED_INTERNALLY"):
		details := oci_certificates_management.CreateRootCaByGeneratingInternallyConfigDetails{}
		if validity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validity")); ok {
			if tmpList := validity.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "validity"), 0)
				tmp, err := s.mapToValidity(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert validity, encountered error: %v", err)
				}
				details.Validity = &tmp
			}
		}
		if versionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_name")); ok {
			tmp := versionName.(string)
			details.VersionName = &tmp
		}
		if signingAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "signing_algorithm")); ok {
			details.SigningAlgorithm = oci_certificates_management.SignatureAlgorithmEnum(signingAlgorithm.(string))
		}

		if subject, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subject")); ok {
			if tmpList := subject.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "subject"), 0)
				tmp, err := s.mapToCertificateAuthoritySubject(fieldKeyFormatNextLevel)
				if err != nil {
					return baseObject, err
				}
				details.Subject = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"):
		details := oci_certificates_management.CreateSubordinateCaIssuedByInternalCaConfigDetails{}
		if issuerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "issuer_certificate_authority_id")); ok {
			tmp := issuerId.(string)
			details.IssuerCertificateAuthorityId = &tmp
		}

		if validity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validity")); ok {
			if tmpList := validity.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "validity"), 0)
				tmp, err := s.mapToValidity(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert validity, encountered error: %v", err)
				}
				details.Validity = &tmp
			}
		}
		if versionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_name")); ok {
			tmp := versionName.(string)
			details.VersionName = &tmp
		}
		if signingAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "signing_algorithm")); ok {
			details.SigningAlgorithm = oci_certificates_management.SignatureAlgorithmEnum(signingAlgorithm.(string))
		}

		if subject, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subject")); ok {
			if tmpList := subject.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "subject"), 0)
				tmp, err := s.mapToCertificateAuthoritySubject(fieldKeyFormatNextLevel)
				if err != nil {
					return baseObject, err
				}
				details.Subject = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown config_type '%v' was specified", configType)
	}
	return baseObject, nil
}

func CreateCertificateAuthorityConfigDetailsToMap(obj *oci_certificates_management.CreateCertificateAuthorityConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_certificates_management.UpdateRootCaByGeneratingInternallyConfigDetails:
		result["config_type"] = "ROOT_CA_GENERATED_INTERNALLY"

		if v.Validity != nil {
			result["validity"] = []interface{}{ValidityToMap(v.Validity)}
		}
	case oci_certificates_management.UpdateSubordinateCaIssuedByInternalCaConfigDetails:
		result["config_type"] = "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"

		if v.Validity != nil {
			result["validity"] = []interface{}{ValidityToMap(v.Validity)}
		}
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) mapToObjectStorageBucketConfigDetails(fieldKeyFormat string) (oci_certificates_management.ObjectStorageBucketConfigDetails, error) {
	result := oci_certificates_management.ObjectStorageBucketConfigDetails{}

	if objectStorageBucketName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_bucket_name")); ok {
		tmp := objectStorageBucketName.(string)
		result.ObjectStorageBucketName = &tmp
	}

	if objectStorageNamespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_namespace")); ok {
		tmp := objectStorageNamespace.(string)
		result.ObjectStorageNamespace = &tmp
	}

	if objectStorageObjectNameFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_object_name_format")); ok {
		tmp := objectStorageObjectNameFormat.(string)
		result.ObjectStorageObjectNameFormat = &tmp
	}

	return result, nil
}

func ObjectStorageBucketConfigDetailsToMap(obj *oci_certificates_management.ObjectStorageBucketConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectStorageBucketName != nil {
		result["object_storage_bucket_name"] = string(*obj.ObjectStorageBucketName)
	}

	if obj.ObjectStorageNamespace != nil {
		result["object_storage_namespace"] = string(*obj.ObjectStorageNamespace)
	}

	if obj.ObjectStorageObjectNameFormat != nil {
		result["object_storage_object_name_format"] = string(*obj.ObjectStorageObjectNameFormat)
	}

	return result
}

func RevocationStatusToMap(obj *oci_certificates_management.RevocationStatus) map[string]interface{} {
	result := map[string]interface{}{}

	result["revocation_reason"] = string(obj.RevocationReason)

	if obj.TimeOfRevocation != nil {
		result["time_of_revocation"] = obj.TimeOfRevocation.String()
	}

	return result
}

func ValidityToMap(obj *oci_certificates_management.Validity) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TimeOfValidityNotAfter != nil {
		result["time_of_validity_not_after"] = obj.TimeOfValidityNotAfter.String()
	}

	if obj.TimeOfValidityNotBefore != nil {
		result["time_of_validity_not_before"] = obj.TimeOfValidityNotBefore.String()
	}

	return result
}

func VersionStageToMap(obj oci_certificates_management.VersionStageEnum) map[string]interface{} {
	result := map[string]interface{}{}

	return result
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_certificates_management.ChangeCertificateAuthorityCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CertificateAuthorityId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	_, err := s.Client.ChangeCertificateAuthorityCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) mapToValidity(fieldKeyFormat string) (oci_certificates_management.Validity, error) {
	result := oci_certificates_management.Validity{}

	if timeOfValidityNotAfter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_of_validity_not_after")); ok {
		tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotAfter.(string))
		if err != nil {
			return result, err
		}
		result.TimeOfValidityNotAfter = &oci_common.SDKTime{Time: tmp}
	}

	if timeOfValidityNotBefore, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_of_validity_not_before")); ok {
		tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotBefore.(string))
		if err != nil {
			return result, err
		}
		result.TimeOfValidityNotBefore = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) mapToCertificateAuthorityRule(fieldKeyFormat string) (oci_certificates_management.CertificateAuthorityRule, error) {
	var baseObject oci_certificates_management.CertificateAuthorityRule
	//discriminator
	ruleTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_type"))
	var ruleType string
	if ok {
		ruleType = ruleTypeRaw.(string)
	} else {
		ruleType = "" // default value
	}
	switch strings.ToLower(ruleType) {
	case strings.ToLower("CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"):
		details := oci_certificates_management.CertificateAuthorityIssuanceExpiryRule{}
		if certificateAuthorityMaxValidityDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_authority_max_validity_duration")); ok {
			tmp := certificateAuthorityMaxValidityDuration.(string)
			details.CertificateAuthorityMaxValidityDuration = &tmp
		}
		if leafCertificateMaxValidityDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "leaf_certificate_max_validity_duration")); ok {
			tmp := leafCertificateMaxValidityDuration.(string)
			details.LeafCertificateMaxValidityDuration = &tmp
		}
		baseObject = &details
	default:
		log.Printf("[WARN] Received 'rule_type' of unknown type %v", ruleType)
		return baseObject, nil
	}

	return baseObject, nil
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) mapToUpdateCertificateAuthorityConfigDetails(fieldKeyFormat string) (oci_certificates_management.UpdateCertificateAuthorityConfigDetails, error) {
	var baseObject oci_certificates_management.UpdateCertificateAuthorityConfigDetails
	//discriminator
	configTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_type"))
	var configType string
	if ok {
		configType = configTypeRaw.(string)
	} else {
		configType = "" // default value
	}
	switch strings.ToLower(configType) {
	case strings.ToLower("ROOT_CA_GENERATED_INTERNALLY"):
		details := oci_certificates_management.UpdateRootCaByGeneratingInternallyConfigDetails{}
		if validity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validity")); ok {
			if tmpList := validity.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "validity"), 0)
				tmp, err := s.mapToValidity(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert validity, encountered error: %v", err)
				}
				details.Validity = &tmp
			}
		}

		if versionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_name")); ok {
			tmp := versionName.(string)
			details.VersionName = &tmp
		}

		baseObject = details
	case strings.ToLower("SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"):
		details := oci_certificates_management.UpdateSubordinateCaIssuedByInternalCaConfigDetails{}
		if validity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validity")); ok {
			if tmpList := validity.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "validity"), 0)
				tmp, err := s.mapToValidity(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert validity, encountered error: %v", err)
				}
				details.Validity = &tmp
			}
		}
		if versionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_name")); ok {
			tmp := versionName.(string)
			details.VersionName = &tmp
		}

		baseObject = details
	default:
		return nil, fmt.Errorf("unknown config_type '%v' was specified", configType)
	}

	return baseObject, nil
}
