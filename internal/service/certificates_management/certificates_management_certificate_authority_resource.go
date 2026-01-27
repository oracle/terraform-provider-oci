// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"
	"encoding/json"
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
								"ROOT_CA_MANAGED_EXTERNALLY",
								"SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA",
								"SUBORDINATE_CA_MANAGED_INTERNALLY_ISSUED_BY_EXTERNAL_CA",
							}, true),
						},

						// Optional
						"certificate_pem": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"issuer_certificate_authority_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
							ForceNew: true,
						},
						"signing_algorithm": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"subject": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
						"action_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action_type": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"GENERATE_CSR",
											"UPDATE_CERTIFICATE",
										}, true),
									},

									// Optional
									"certificate_pem": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
									},
								},
							},
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
								"CERTIFICATE_AUTHORITY_ISSUANCE_RULE",
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
						"name_constraint": {
							Type:             schema.TypeList,
							Optional:         true,
							Computed:         true,
							MaxItems:         1,
							MinItems:         1,
							DiffSuppressFunc: nameOrPathConstraintDiffSuppress,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"excluded_subtree": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"permitted_subtree": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"path_length_constraint": {
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: nameOrPathConstraintDiffSuppress,
						},

						// Computed
					},
				},
				DiffSuppressFunc: certificateAuthorityRulesDiffSuppress,
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
			"external_key_description": {
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
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
		string(oci_certificates_management.CertificateAuthorityLifecycleStatePendingActivation),
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

	if externalKeyDescription, ok := s.D.GetOkExists("external_key_description"); ok {
		tmp := externalKeyDescription.(string)
		request.ExternalKeyDescription = &tmp
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

	if externalKeyDescription, ok := s.D.GetOkExists("external_key_description"); ok {
		tmp := externalKeyDescription.(string)
		request.ExternalKeyDescription = &tmp
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
		newCARules := certificateAuthorityRules.([]interface{})
		var tmpRules []oci_certificates_management.CertificateAuthorityRule
		for i := range newCARules {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_authority_rules", i)
			converted, err := s.mapToCertificateAuthorityRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			// Check the rule type and skip the "CERTIFICATE_AUTHORITY_ISSUANCE_RULE" type
			if _, okIssuanceRule := converted.(oci_certificates_management.CertificateAuthorityIssuanceRule); okIssuanceRule {
				continue // Skip this rule if it's of type "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"
			}

			// If it wasn't skipped, append the rule to tmp
			tmpRules = append(tmpRules, converted)
		}

		oldRules, _ := s.D.GetChange("certificate_authority_rules")
		if oldRulesList, okOldRulesList := oldRules.([]interface{}); okOldRulesList {
			for _, rule := range oldRulesList {
				if ruleMap, ok := rule.(map[string]interface{}); ok {
					ruleType := ruleMap["rule_type"].(string)
					if strings.EqualFold("CERTIFICATE_AUTHORITY_ISSUANCE_RULE", ruleType) {
						caRule, err := caIssuanceRuleMapFromStateToCARule(ruleMap)
						if err != nil {
							return fmt.Errorf("unable to convert certificate_authority_rules from state, encountered error: %v", err)
						}
						tmpRules = append(tmpRules, caRule)
					}

				} else {
					return fmt.Errorf("not a map. unable to convert certificate_authority_rules from state, encountered error")
				}
			}
		} else {
			return fmt.Errorf("not a []interface. unable to convert certificate_authority_rules from state, encountered error")
		}

		if len(tmp) != 0 || s.D.HasChange("certificate_authority_rules") {
			request.CertificateAuthorityRules = tmpRules
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

	if s.Res.ExternalKeyDescription != nil {
		s.D.Set("external_key_description", *s.Res.ExternalKeyDescription)
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

func caIssuanceRuleMapFromStateToCARule(caIssuanceRuleMap map[string]interface{}) (oci_certificates_management.CertificateAuthorityRule, error) {
	var baseObject oci_certificates_management.CertificateAuthorityRule
	details := oci_certificates_management.CertificateAuthorityIssuanceRule{}

	if nameConstraint, ok := caIssuanceRuleMap["name_constraint"]; ok {
		if tmpList := nameConstraint.([]interface{}); ok {
			tmp, err := nameConstraintMapFromStateToNameConstraint(tmpList[0].(map[string]interface{}))
			if err != nil {
				return details, fmt.Errorf("unable to convert name_constraint, encountered error: %v", err)
			}
			details.NameConstraint = &tmp
		} else {
			return nil, fmt.Errorf("unable to convert name_constraint list, encountered error: %+v", nameConstraint)
		}
	}
	if pathLengthConstraint, ok := caIssuanceRuleMap["path_length_constraint"]; ok {
		tmp := pathLengthConstraint.(int)
		details.PathLengthConstraint = &tmp
	}

	baseObject = details
	return baseObject, nil
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
	case strings.ToLower("ROOT_CA_MANAGED_EXTERNALLY"):
		details := oci_certificates_management.UpdateRootCaManagedExternallyConfigDetails{}
		if certificatePem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_pem")); ok {
			tmp := certificatePem.(string)
			details.CertificatePem = &tmp
		}
		if versionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_name")); ok {
			tmp := versionName.(string)
			details.VersionName = &tmp
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
	case strings.ToLower("SUBORDINATE_CA_MANAGED_INTERNALLY_ISSUED_BY_EXTERNAL_CA"):
		details := oci_certificates_management.CreateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails{}
		if issuerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "issuer_certificate_authority_id")); ok {
			tmp := issuerId.(string)
			details.IssuerCertificateAuthorityId = &tmp
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
	case oci_certificates_management.UpdateRootCaManagedExternallyConfigDetails:
		result["config_type"] = "ROOT_CA_MANAGED_EXTERNALLY"

		if v.CertificatePem != nil {
			result["certificate_pem"] = string(*v.CertificatePem)
		}
	case oci_certificates_management.UpdateSubordinateCaIssuedByInternalCaConfigDetails:
		result["config_type"] = "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"

		if v.Validity != nil {
			result["validity"] = []interface{}{ValidityToMap(v.Validity)}
		}
	case oci_certificates_management.UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails:
		result["config_type"] = "SUBORDINATE_CA_MANAGED_INTERNALLY_ISSUED_BY_EXTERNAL_CA"

		if v.ActionDetails != nil {
			actionDetailsArray := []interface{}{}
			if actionDetailsMap := UpdateCertificateAuthorityActionDetailsToMap(&v.ActionDetails); actionDetailsMap != nil {
				actionDetailsArray = append(actionDetailsArray, actionDetailsMap)
			}
			result["action_details"] = actionDetailsArray
		}
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) mapToNameConstraint(fieldKeyFormat string) (oci_certificates_management.NameConstraint, error) {
	result := oci_certificates_management.NameConstraint{}

	if excludedSubtree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "excluded_subtree")); ok {
		interfaces := excludedSubtree.([]interface{})
		tmp := make([]oci_certificates_management.NameConstraintSubtreeNode, len(interfaces))
		for i := range interfaces {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "excluded_subtree"), i)
			converted, err := s.mapToNameConstraintSubtreeNode(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "excluded_subtree")) {
			result.ExcludedSubtree = tmp
		}
	}

	if permittedSubtree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "permitted_subtree")); ok {
		interfaces := permittedSubtree.([]interface{})
		tmp := make([]oci_certificates_management.NameConstraintSubtreeNode, len(interfaces))
		for i := range interfaces {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "permitted_subtree"), i)
			converted, err := s.mapToNameConstraintSubtreeNode(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "permitted_subtree")) {
			result.PermittedSubtree = tmp
		}
	}

	return result, nil
}

func nameConstraintMapFromStateToNameConstraint(nameConstraintMap map[string]interface{}) (oci_certificates_management.NameConstraint, error) {
	result := oci_certificates_management.NameConstraint{}

	if excludedSubtree, ok := nameConstraintMap["excluded_subtree"]; ok {
		excludedSubtrees := excludedSubtree.([]interface{})
		tmp := make([]oci_certificates_management.NameConstraintSubtreeNode, len(excludedSubtrees))
		for i, excludedTree := range excludedSubtrees {
			converted, err := nameConstraintSubtreeMapToNode(excludedTree)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 {
			result.ExcludedSubtree = tmp
		}
	}

	if permittedSubtree, ok := nameConstraintMap["permitted_subtree"]; ok {
		permittedSubtrees := permittedSubtree.([]interface{})
		tmp := make([]oci_certificates_management.NameConstraintSubtreeNode, len(permittedSubtrees))
		for i, permittedTree := range permittedSubtrees {
			converted, err := nameConstraintSubtreeMapToNode(permittedTree)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 {
			result.PermittedSubtree = tmp
		}
	}

	return result, nil
}

func nameConstraintSubtreesFromStateToNameConstraint(obj *oci_certificates_management.NameConstraint) map[string]interface{} {
	result := map[string]interface{}{}

	excludedSubtree := []interface{}{}
	for _, item := range obj.ExcludedSubtree {
		excludedSubtree = append(excludedSubtree, nameConstraintSubtreeNodeFromStateToNameConstraintSubtreeNode(item))
	}
	result["excluded_subtree"] = excludedSubtree

	permittedSubtree := []interface{}{}
	for _, item := range obj.PermittedSubtree {
		permittedSubtree = append(permittedSubtree, nameConstraintSubtreeNodeFromStateToNameConstraintSubtreeNode(item))
	}
	result["permitted_subtree"] = permittedSubtree

	return result
}

func (s *CertificatesManagementCertificateAuthorityResourceCrud) mapToNameConstraintSubtreeNode(fieldKeyFormat string) (oci_certificates_management.NameConstraintSubtreeNode, error) {
	result := oci_certificates_management.NameConstraintSubtreeNode{}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_certificates_management.NameConstraintTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func nameConstraintSubtreeMapToNode(subtree interface{}) (oci_certificates_management.NameConstraintSubtreeNode, error) {
	result := oci_certificates_management.NameConstraintSubtreeNode{}
	if subtreeMap, ok := subtree.(map[string]interface{}); ok {
		if type_, ok := subtreeMap["type"]; ok {
			result.Type = oci_certificates_management.NameConstraintTypeEnum(type_.(string))
		}

		if value, ok := subtreeMap["value"]; ok {
			tmp := value.(string)
			result.Value = &tmp
		}
		return result, nil
	}
	return result, fmt.Errorf("unable to convert subtree node,  %#v", subtree)
}

func nameConstraintSubtreeNodeFromStateToNameConstraintSubtreeNode(nameConstraintSubtreeNode oci_certificates_management.NameConstraintSubtreeNode) map[string]interface{} {
	result := map[string]interface{}{}

	result["type"] = string(nameConstraintSubtreeNode.Type)

	if nameConstraintSubtreeNode.Value != nil {
		result["value"] = string(*nameConstraintSubtreeNode.Value)
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

func (s *CertificatesManagementCertificateAuthorityResourceCrud) mapToUpdateCertificateAuthorityActionDetails(fieldKeyFormat string) (oci_certificates_management.UpdateCertificateAuthorityActionDetails, error) {
	var baseObject oci_certificates_management.UpdateCertificateAuthorityActionDetails
	//discriminator
	actionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_type"))
	var actionType string
	if ok {
		actionType = actionTypeRaw.(string)
	} else {
		actionType = "" // default value
	}
	switch strings.ToLower(actionType) {
	case strings.ToLower("GENERATE_CSR"):
		details := oci_certificates_management.UpdateCertificateAuthorityGenerateCsrDetails{}
		baseObject = details
	case strings.ToLower("UPDATE_CERTIFICATE"):
		details := oci_certificates_management.UpdateCertificateAuthorityCertificateDetails{}
		if certificatePem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_pem")); ok {
			tmp := certificatePem.(string)
			details.CertificatePem = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown action_type '%v' was specified", actionType)
	}
	return baseObject, nil
}

func UpdateCertificateAuthorityActionDetailsToMap(obj *oci_certificates_management.UpdateCertificateAuthorityActionDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_certificates_management.UpdateCertificateAuthorityGenerateCsrDetails:
		result["action_type"] = "GENERATE_CSR"
	case oci_certificates_management.UpdateCertificateAuthorityCertificateDetails:
		result["action_type"] = "UPDATE_CERTIFICATE"

		if v.CertificatePem != nil {
			result["certificate_pem"] = string(*v.CertificatePem)
		}
	default:
		log.Printf("[WARN] Received 'action_type' of unknown type %v", *obj)
		return nil
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

		baseObject = details
	case strings.ToLower("CERTIFICATE_AUTHORITY_ISSUANCE_RULE"):
		details := oci_certificates_management.CertificateAuthorityIssuanceRule{}
		if nameConstraint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name_constraint")); ok {
			if tmpList := nameConstraint.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "name_constraint"), 0)
				tmp, err := s.mapToNameConstraint(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert name_constraint, encountered error: %v", err)
				}
				details.NameConstraint = &tmp
			}
		}
		if pathLengthConstraint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path_length_constraint")); ok {
			tmp := pathLengthConstraint.(int)
			details.PathLengthConstraint = &tmp
		}

		baseObject = details
	default:
		return nil, fmt.Errorf("unknown rule_type '%v' was specified", ruleType)
	}

	return baseObject, nil
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
		result["path_length_constraint"] = 0
		result["name_constraint"] = nil
	case oci_certificates_management.CertificateAuthorityIssuanceRule:
		result["rule_type"] = "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"

		if v.NameConstraint != nil {
			result["name_constraint"] = []interface{}{nameConstraintSubtreesFromStateToNameConstraint(v.NameConstraint)}
		}

		if v.PathLengthConstraint != nil {
			result["path_length_constraint"] = int(*v.PathLengthConstraint)
		}
		result["certificate_authority_max_validity_duration"] = ""
		result["leaf_certificate_max_validity_duration"] = ""
	default:
		log.Printf("[WARN] Received 'rule_type' of unknown type %v", obj)
		return nil
	}

	return result
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
	case strings.ToLower("ROOT_CA_MANAGED_EXTERNALLY"):
		details := oci_certificates_management.UpdateRootCaManagedExternallyConfigDetails{}
		if certificatePem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_pem")); ok {
			tmp := certificatePem.(string)
			details.CertificatePem = &tmp
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
	case strings.ToLower("SUBORDINATE_CA_MANAGED_INTERNALLY_ISSUED_BY_EXTERNAL_CA"):
		details := oci_certificates_management.UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails{}
		if actionDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_details")); ok {
			if tmpList := actionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "action_details"), 0)
				tmp, err := s.mapToUpdateCertificateAuthorityActionDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert action_details, encountered error: %v", err)
				}
				details.ActionDetails = tmp
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

func nameOrPathConstraintDiffSuppress(k, old, new string, d *schema.ResourceData) bool {

	// Suppress only if the resource is already created
	if d.Id() == "" {
		return false
	}

	// Extract the rule index from the key path of format "certificate_authority_rules.0.name_constraint" or """certificate_authority_rules.0.path_length_constraint"
	keyParts := strings.Split(k, ".")
	if len(keyParts) < 3 {
		return false
	}
	keyStr := keyParts[2]

	//Doesn't matter what the actual change is if the key is name_constraint or path_constraint suppress. Diff will be taken care by ca_rules diff at the higher layer.
	return strings.EqualFold(keyStr, "name_constraint") || strings.EqualFold(keyStr, "path_length_constraint")
}

func certificateAuthorityRulesDiffSuppress(k, old, new string, d *schema.ResourceData) bool {

	// Check if the ID is not set (it's a create operation)
	if d.Id() == "" {
		return false
	}

	oldRules, newRules := d.GetChange("certificate_authority_rules")

	if oldList, okA := oldRules.([]interface{}); okA {
		if newList, okB := newRules.([]interface{}); okB {
			var srcList, destList []interface{}

			// Assign longer list to oldList, shorter to newList
			if len(oldList) >= len(newList) {
				srcList = oldList
				destList = newList
			} else {
				srcList = newList
				destList = oldList
			}

			// Build a map from destList using rule_type;
			destRulesMap := make(map[string]map[string]interface{})
			for _, rule := range destList {
				if ruleMap, ok := rule.(map[string]interface{}); ok {
					if rt, okRt := ruleMap["rule_type"].(string); okRt {
						_, found := destRulesMap[rt]
						if found {
							return false
						}
						destRulesMap[rt] = ruleMap
					}
				}
			}

			// Compare srcList entries with matching rule_type from destList
			for _, rule := range srcList {
				srcMap, ok := rule.(map[string]interface{})
				if !ok {
					continue
				}

				ruleType, ok := srcMap["rule_type"].(string)
				if !ok {
					continue
				}

				if strings.EqualFold("CERTIFICATE_AUTHORITY_ISSUANCE_RULE", ruleType) {
					continue
				}

				newMap, found := destRulesMap[ruleType]
				if !found {
					return false //mismatch
				}

				oldJSON, _ := json.MarshalIndent(srcMap, "", "  ")
				newJSON, _ := json.MarshalIndent(newMap, "", "  ")

				if string(oldJSON) != string(newJSON) {
					return false
				} else {
					delete(destRulesMap, ruleType)
				}
			}

			if len(destRulesMap) == 0 {
				return true
			}

		}
	}
	return false
}
