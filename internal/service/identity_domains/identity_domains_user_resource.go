// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsUserResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsUser,
		Read:     readIdentityDomainsUser,
		Update:   updateIdentityDomainsUser,
		Delete:   deleteIdentityDomainsUser,
		Schema: map[string]*schema.Schema{
			// Required
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"family_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"formatted": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"given_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"honorific_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"honorific_suffix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"middle_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"country": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"formatted": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"locality": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"postal_code": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"primary": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"street_address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emails": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"primary": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"secondary": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"verified": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
						"pending_verification_data": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"entitlements": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"primary": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ims": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"primary": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"locale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nick_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"phone_numbers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"primary": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"verified": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"photos": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"primary": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"preferred_language": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"roles": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"primary": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"title": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urnietfparamsscimschemasextensionenterprise20user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"cost_center": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"department": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"division": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"employee_number": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"manager": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"organization": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextension_oci_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"defined_tags": {
							Type:             schema.TypeList,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"freeform_tags": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
						"tag_slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionadaptive_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"risk_level": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"risk_scores": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"last_update_timestamp": {
										Type:     schema.TypeString,
										Required: true,
									},
									"risk_level": {
										Type:     schema.TypeString,
										Required: true,
									},
									"score": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"source": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
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
			"urnietfparamsscimschemasoracleidcsextensioncapabilities_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"can_use_api_keys": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"can_use_auth_tokens": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"can_use_console": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"can_use_console_password": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"can_use_customer_secret_keys": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"can_use_db_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"can_use_oauth2client_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"can_use_smtp_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensiondb_credentials_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"db_user_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"db_login_attempts": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionkerberos_user_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"realm_users": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"principal_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"realm_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
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
			"urnietfparamsscimschemasoracleidcsextensionmfa_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"bypass_codes": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"devices": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"authentication_method": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"factor_status": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"factor_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"last_sync_time": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"third_party_vendor_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"login_attempts": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mfa_enabled_on": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mfa_ignored_apps": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"mfa_status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preferred_authentication_factor": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preferred_authentication_method": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preferred_device": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"preferred_third_party_vendor": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trusted_user_agents": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
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
			"urnietfparamsscimschemasoracleidcsextensionpasswordless_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"factor_identifier": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"factor_method": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"factor_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionposix_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"gecos": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gid_number": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"home_directory": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"login_shell": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uid_number": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"sec_questions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"answer": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"hint_text": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
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
			"urnietfparamsscimschemasoracleidcsextensionself_change_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"allow_self_change": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionself_registration_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"self_registration_profile": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						// Optional
						"consent_granted": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"user_token": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionsff_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"sff_auth_keys": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionsocial_account_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"social_accounts": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
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
			"urnietfparamsscimschemasoracleidcsextensionterms_of_use_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"terms_of_use_consents": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"ref": {
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
			"urnietfparamsscimschemasoracleidcsextensionuser_state_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"locked": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"expired": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"lock_date": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"on": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"reason": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"max_concurrent_sessions": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"recovery_locked": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"lock_date": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"on": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
						"last_failed_login_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_successful_login_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"login_attempts": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"previous_successful_login_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"recovery_attempts": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"recovery_enroll_attempts": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionuser_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"user_provider": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"account_recovery_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"bypass_notification": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"creation_mechanism": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"delegated_authentication_target_app": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"do_not_show_getting_started": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_authentication_delegated": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_federated_user": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_group_membership_normalized": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_group_membership_synced_to_users_groups": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"notification_email_template_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preferred_ui_landing_page": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"synced_from_app": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"user_flow_controlled_by_external_client": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
						"accounts": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"active": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"app_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
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
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"app_roles": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"admin_role": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"app_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"app_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"legacy_group_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"applicable_authentication_target_app": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"target_request_timeout": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"grants": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"app_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"grant_mechanism": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"grantor_id": {
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
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"group_membership_last_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_app_roles_limited_to_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"idcs_app_role_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ocid": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_account_recovery_enrolled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"support_accounts": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"user_provider": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ocid": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"user_id": {
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
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"user_token": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"user_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"x509certificates": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"primary": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"date_added": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"membership_ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"non_unique_display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
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
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
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
						"created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"urnietfparamsscimschemasoracleidcsextensiondb_user_user": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"db_global_roles": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"domain_level_schema": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"instance_level_schema": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_db_user": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"password_verifiers": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionpassword_state_user": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"applicable_password_policy": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"priority": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"cant_change": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"cant_expire": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"expired": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"last_failed_validation_date": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_successful_set_date": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_successful_validation_date": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"must_change": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionuser_credentials_user": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"api_keys": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ocid": {
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
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"auth_tokens": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"ocid": {
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
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"customer_secret_keys": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"ocid": {
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
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"db_credentials": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"ocid": {
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
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"o_auth2client_credentials": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"ocid": {
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
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"smtp_credentials": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"ocid": {
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
									"ref": {
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
		},
	}
}

func createIdentityDomainsUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsUserResourceCrud{}
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

	return tfresource.CreateResource(d, sync)
}

func readIdentityDomainsUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsUserResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "users")
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

func updateIdentityDomainsUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsUserResourceCrud{}
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

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityDomainsUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsUserResourceCrud{}
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
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityDomainsUserResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.User
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsUserResourceCrud) ID() string {
	return *s.Res.Id
	//return GetUserCompositeId(s.D.Get("id").(string))
}

func (s *IdentityDomainsUserResourceCrud) Create() error {
	request := oci_identity_domains.CreateUserRequest{}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
	}

	if addresses, ok := s.D.GetOkExists("addresses"); ok {
		interfaces := addresses.([]interface{})
		tmp := make([]oci_identity_domains.Addresses, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "addresses", stateDataIndex)
			converted, err := s.mapToaddresses(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("addresses") {
			request.Addresses = tmp
		}
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if emails, ok := s.D.GetOkExists("emails"); ok {
		interfaces := emails.([]interface{})
		tmp := make([]oci_identity_domains.UserEmails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "emails", stateDataIndex)
			converted, err := s.mapToUserEmails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("emails") {
			request.Emails = tmp
		}
	}

	if entitlements, ok := s.D.GetOkExists("entitlements"); ok {
		interfaces := entitlements.([]interface{})
		tmp := make([]oci_identity_domains.UserEntitlements, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "entitlements", stateDataIndex)
			converted, err := s.mapToUserEntitlements(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("entitlements") {
			request.Entitlements = tmp
		}
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if ims, ok := s.D.GetOkExists("ims"); ok {
		interfaces := ims.([]interface{})
		tmp := make([]oci_identity_domains.UserIms, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ims", stateDataIndex)
			converted, err := s.mapToUserIms(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("ims") {
			request.Ims = tmp
		}
	}

	if locale, ok := s.D.GetOkExists("locale"); ok {
		tmp := locale.(string)
		request.Locale = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		if tmpList := name.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "name", 0)
			tmp, err := s.mapToUserName(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Name = &tmp
		}
	}

	if nickName, ok := s.D.GetOkExists("nick_name"); ok {
		tmp := nickName.(string)
		request.NickName = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	if phoneNumbers, ok := s.D.GetOkExists("phone_numbers"); ok {
		interfaces := phoneNumbers.([]interface{})
		tmp := make([]oci_identity_domains.UserPhoneNumbers, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "phone_numbers", stateDataIndex)
			converted, err := s.mapToUserPhoneNumbers(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("phone_numbers") {
			request.PhoneNumbers = tmp
		}
	}

	if photos, ok := s.D.GetOkExists("photos"); ok {
		interfaces := photos.([]interface{})
		tmp := make([]oci_identity_domains.UserPhotos, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "photos", stateDataIndex)
			converted, err := s.mapToUserPhotos(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("photos") {
			request.Photos = tmp
		}
	}

	if preferredLanguage, ok := s.D.GetOkExists("preferred_language"); ok {
		tmp := preferredLanguage.(string)
		request.PreferredLanguage = &tmp
	}

	if profileUrl, ok := s.D.GetOkExists("profile_url"); ok {
		tmp := profileUrl.(string)
		request.ProfileUrl = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if roles, ok := s.D.GetOkExists("roles"); ok {
		interfaces := roles.([]interface{})
		tmp := make([]oci_identity_domains.UserRoles, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "roles", stateDataIndex)
			converted, err := s.mapToUserRoles(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("roles") {
			request.Roles = tmp
		}
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if timezone, ok := s.D.GetOkExists("timezone"); ok {
		tmp := timezone.(string)
		request.Timezone = &tmp
	}

	if title, ok := s.D.GetOkExists("title"); ok {
		tmp := title.(string)
		request.Title = &tmp
	}

	if urnietfparamsscimschemasextensionenterprise20User, ok := s.D.GetOkExists("urnietfparamsscimschemasextensionenterprise20user"); ok {
		if tmpList := urnietfparamsscimschemasextensionenterprise20User.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasextensionenterprise20user", 0)
			tmp, err := s.mapToExtensionEnterprise20User(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasExtensionEnterprise2_0User = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionOCITags, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextension_oci_tags"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionOCITags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextension_oci_tags", 0)
			tmp, err := s.mapToExtensionOCITags(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionadaptiveUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionadaptive_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionadaptiveUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionadaptive_user", 0)
			tmp, err := s.mapToExtensionAdaptiveUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionAdaptiveUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensioncapabilitiesUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensioncapabilities_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensioncapabilitiesUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensioncapabilities_user", 0)
			tmp, err := s.mapToExtensionCapabilitiesUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionCapabilitiesUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensiondbCredentialsUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensiondb_credentials_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensiondbCredentialsUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensiondb_credentials_user", 0)
			tmp, err := s.mapToExtensionDbCredentialsUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionDbCredentialsUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionkerberosUserUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionkerberos_user_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionkerberosUserUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionkerberos_user_user", 0)
			tmp, err := s.mapToExtensionKerberosUserUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosUserUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionmfaUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionmfa_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionmfaUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionmfa_user", 0)
			tmp, err := s.mapToExtensionMfaUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionMfaUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionpasswordlessUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionpasswordless_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionpasswordlessUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionpasswordless_user", 0)
			tmp, err := s.mapToExtensionPasswordlessUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordlessUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionposixUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionposix_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionposixUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionposix_user", 0)
			tmp, err := s.mapToExtensionPosixUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionsecurityQuestionsUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionsecurityQuestionsUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user", 0)
			tmp, err := s.mapToExtensionSecurityQuestionsUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSecurityQuestionsUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionselfChangeUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionself_change_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionselfChangeUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionself_change_user", 0)
			tmp, err := s.mapToExtensionSelfChangeUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionselfRegistrationUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionself_registration_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionselfRegistrationUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionself_registration_user", 0)
			tmp, err := s.mapToExtensionSelfRegistrationUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfRegistrationUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionsffUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionsff_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionsffUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionsff_user", 0)
			tmp, err := s.mapToExtensionSffUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSffUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionsocialAccountUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionsocial_account_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionsocialAccountUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionsocial_account_user", 0)
			tmp, err := s.mapToExtensionSocialAccountUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialAccountUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensiontermsOfUseUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionterms_of_use_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensiontermsOfUseUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionterms_of_use_user", 0)
			tmp, err := s.mapToExtensionTermsOfUseUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionTermsOfUseUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionuserStateUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionuser_state_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionuserStateUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionuser_state_user", 0)
			tmp, err := s.mapToExtensionUserStateUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionUserStateUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionuserUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionuser_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionuserUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionuser_user", 0)
			tmp, err := s.mapToExtensionUserUser(fieldKeyFormat, false)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionUserUser = &tmp
		}
	}

	if userName, ok := s.D.GetOkExists("user_name"); ok {
		tmp := userName.(string)
		request.UserName = &tmp
	}

	if userType, ok := s.D.GetOkExists("user_type"); ok {
		request.UserType = oci_identity_domains.UserUserTypeEnum(userType.(string))
	}

	if x509Certificates, ok := s.D.GetOkExists("x509certificates"); ok {
		interfaces := x509Certificates.([]interface{})
		tmp := make([]oci_identity_domains.UserX509Certificates, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "x509certificates", stateDataIndex)
			converted, err := s.mapToUserX509Certificates(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("x509certificates") {
			request.X509Certificates = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *IdentityDomainsUserResourceCrud) Get() error {
	request := oci_identity_domains.GetUserRequest{}

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

	tmp := s.D.Id()
	request.UserId = &tmp

	userId, err := parseUserCompositeId(s.D.Id())
	if err == nil {
		request.UserId = &userId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *IdentityDomainsUserResourceCrud) Update() error {
	request := oci_identity_domains.PutUserRequest{}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
	}

	if addresses, ok := s.D.GetOkExists("addresses"); ok {
		interfaces := addresses.([]interface{})
		tmp := make([]oci_identity_domains.Addresses, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "addresses", stateDataIndex)
			converted, err := s.mapToaddresses(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("addresses") {
			request.Addresses = tmp
		}
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if emails, ok := s.D.GetOkExists("emails"); ok {
		interfaces := emails.([]interface{})
		tmp := make([]oci_identity_domains.UserEmails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "emails", stateDataIndex)
			converted, err := s.mapToUserEmails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("emails") {
			request.Emails = tmp
		}
	}

	if entitlements, ok := s.D.GetOkExists("entitlements"); ok {
		interfaces := entitlements.([]interface{})
		tmp := make([]oci_identity_domains.UserEntitlements, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "entitlements", stateDataIndex)
			converted, err := s.mapToUserEntitlements(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("entitlements") {
			request.Entitlements = tmp
		}
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	tmp := s.D.Id()
	request.Id = &tmp

	if ims, ok := s.D.GetOkExists("ims"); ok {
		interfaces := ims.([]interface{})
		tmp := make([]oci_identity_domains.UserIms, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ims", stateDataIndex)
			converted, err := s.mapToUserIms(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("ims") {
			request.Ims = tmp
		}
	}

	if locale, ok := s.D.GetOkExists("locale"); ok {
		tmp := locale.(string)
		request.Locale = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		if tmpList := name.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "name", 0)
			tmp, err := s.mapToUserName(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Name = &tmp
		}
	}

	if nickName, ok := s.D.GetOkExists("nick_name"); ok {
		tmp := nickName.(string)
		request.NickName = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		if s.D.HasChange("password") {
			tmp := password.(string)
			request.Password = &tmp
		}
	}

	if phoneNumbers, ok := s.D.GetOkExists("phone_numbers"); ok {
		interfaces := phoneNumbers.([]interface{})
		tmp := make([]oci_identity_domains.UserPhoneNumbers, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "phone_numbers", stateDataIndex)
			converted, err := s.mapToUserPhoneNumbers(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("phone_numbers") {
			request.PhoneNumbers = tmp
		}
	}

	if photos, ok := s.D.GetOkExists("photos"); ok {
		interfaces := photos.([]interface{})
		tmp := make([]oci_identity_domains.UserPhotos, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "photos", stateDataIndex)
			converted, err := s.mapToUserPhotos(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("photos") {
			request.Photos = tmp
		}
	}

	if preferredLanguage, ok := s.D.GetOkExists("preferred_language"); ok {
		tmp := preferredLanguage.(string)
		request.PreferredLanguage = &tmp
	}

	if profileUrl, ok := s.D.GetOkExists("profile_url"); ok {
		tmp := profileUrl.(string)
		request.ProfileUrl = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if roles, ok := s.D.GetOkExists("roles"); ok {
		interfaces := roles.([]interface{})
		tmp := make([]oci_identity_domains.UserRoles, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "roles", stateDataIndex)
			converted, err := s.mapToUserRoles(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("roles") {
			request.Roles = tmp
		}
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if timezone, ok := s.D.GetOkExists("timezone"); ok {
		tmp := timezone.(string)
		request.Timezone = &tmp
	}

	if title, ok := s.D.GetOkExists("title"); ok {
		tmp := title.(string)
		request.Title = &tmp
	}

	if urnietfparamsscimschemasextensionenterprise20User, ok := s.D.GetOkExists("urnietfparamsscimschemasextensionenterprise20user"); ok {
		if tmpList := urnietfparamsscimschemasextensionenterprise20User.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasextensionenterprise20user", 0)
			tmp, err := s.mapToExtensionEnterprise20User(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasExtensionEnterprise2_0User = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionOCITags, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextension_oci_tags"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionOCITags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextension_oci_tags", 0)
			tmp, err := s.mapToExtensionOCITags(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionadaptiveUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionadaptive_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionadaptiveUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionadaptive_user", 0)
			tmp, err := s.mapToExtensionAdaptiveUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionAdaptiveUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensioncapabilitiesUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensioncapabilities_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensioncapabilitiesUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensioncapabilities_user", 0)
			tmp, err := s.mapToExtensionCapabilitiesUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionCapabilitiesUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensiondbCredentialsUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensiondb_credentials_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensiondbCredentialsUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensiondb_credentials_user", 0)
			tmp, err := s.mapToExtensionDbCredentialsUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionDbCredentialsUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionkerberosUserUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionkerberos_user_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionkerberosUserUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionkerberos_user_user", 0)
			tmp, err := s.mapToExtensionKerberosUserUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosUserUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionmfaUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionmfa_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionmfaUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionmfa_user", 0)
			tmp, err := s.mapToExtensionMfaUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionMfaUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionpasswordlessUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionpasswordless_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionpasswordlessUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionpasswordless_user", 0)
			tmp, err := s.mapToExtensionPasswordlessUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordlessUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionposixUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionposix_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionposixUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionposix_user", 0)
			tmp, err := s.mapToExtensionPosixUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionsecurityQuestionsUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionsecurityQuestionsUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user", 0)
			tmp, err := s.mapToExtensionSecurityQuestionsUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSecurityQuestionsUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionselfChangeUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionself_change_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionselfChangeUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionself_change_user", 0)
			tmp, err := s.mapToExtensionSelfChangeUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionselfRegistrationUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionself_registration_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionselfRegistrationUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionself_registration_user", 0)
			tmp, err := s.mapToExtensionSelfRegistrationUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfRegistrationUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionsffUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionsff_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionsffUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionsff_user", 0)
			tmp, err := s.mapToExtensionSffUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSffUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionsocialAccountUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionsocial_account_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionsocialAccountUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionsocial_account_user", 0)
			tmp, err := s.mapToExtensionSocialAccountUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialAccountUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensiontermsOfUseUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionterms_of_use_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensiontermsOfUseUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionterms_of_use_user", 0)
			tmp, err := s.mapToExtensionTermsOfUseUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionTermsOfUseUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionuserStateUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionuser_state_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionuserStateUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionuser_state_user", 0)
			tmp, err := s.mapToExtensionUserStateUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionUserStateUser = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionuserUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionuser_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionuserUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionuser_user", 0)
			tmp, err := s.mapToExtensionUserUser(fieldKeyFormat, true)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionUserUser = &tmp
		}
	}

	tmp = s.D.Id()
	request.UserId = &tmp

	if userName, ok := s.D.GetOkExists("user_name"); ok {
		tmp := userName.(string)
		request.UserName = &tmp
	}

	if userType, ok := s.D.GetOkExists("user_type"); ok {
		request.UserType = oci_identity_domains.UserUserTypeEnum(userType.(string))
	}

	if x509Certificates, ok := s.D.GetOkExists("x509certificates"); ok {
		interfaces := x509Certificates.([]interface{})
		tmp := make([]oci_identity_domains.UserX509Certificates, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "x509certificates", stateDataIndex)
			converted, err := s.mapToUserX509Certificates(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("x509certificates") {
			request.X509Certificates = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *IdentityDomainsUserResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteUserRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	tmp := s.D.Id()
	request.UserId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteUser(context.Background(), request)
	return err
}

func (s *IdentityDomainsUserResourceCrud) SetData() error {

	userId, err := parseUserCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(userId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
	}

	addresses := []interface{}{}
	for _, item := range s.Res.Addresses {
		addresses = append(addresses, addressesToMap(item))
	}
	s.D.Set("addresses", addresses)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	emails := []interface{}{}
	for _, item := range s.Res.Emails {
		emails = append(emails, UserEmailsToMap(item))
	}
	s.D.Set("emails", emails)

	entitlements := []interface{}{}
	for _, item := range s.Res.Entitlements {
		entitlements = append(entitlements, UserEntitlementsToMap(item))
	}
	s.D.Set("entitlements", entitlements)

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	groups := []interface{}{}
	for _, item := range s.Res.Groups {
		groups = append(groups, UserGroupsToMap(item))
	}
	s.D.Set("groups", groups)

	if s.Res.IdcsCreatedBy != nil {
		s.D.Set("idcs_created_by", []interface{}{idcsCreatedByToMap(s.Res.IdcsCreatedBy)})
	} else {
		s.D.Set("idcs_created_by", nil)
	}

	if s.Res.IdcsLastModifiedBy != nil {
		s.D.Set("idcs_last_modified_by", []interface{}{idcsLastModifiedByToMap(s.Res.IdcsLastModifiedBy)})
	} else {
		s.D.Set("idcs_last_modified_by", nil)
	}

	if s.Res.IdcsLastUpgradedInRelease != nil {
		s.D.Set("idcs_last_upgraded_in_release", *s.Res.IdcsLastUpgradedInRelease)
	}

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	ims := []interface{}{}
	for _, item := range s.Res.Ims {
		ims = append(ims, UserImsToMap(item))
	}
	s.D.Set("ims", ims)

	if s.Res.Locale != nil {
		s.D.Set("locale", *s.Res.Locale)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", []interface{}{UserNameToMap(s.Res.Name)})
	} else {
		s.D.Set("name", nil)
	}

	if s.Res.NickName != nil {
		s.D.Set("nick_name", *s.Res.NickName)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.Password != nil {
		s.D.Set("password", *s.Res.Password)
	}

	phoneNumbers := []interface{}{}
	for _, item := range s.Res.PhoneNumbers {
		phoneNumbers = append(phoneNumbers, UserPhoneNumbersToMap(item))
	}
	s.D.Set("phone_numbers", phoneNumbers)

	photos := []interface{}{}
	for _, item := range s.Res.Photos {
		photos = append(photos, UserPhotosToMap(item))
	}
	s.D.Set("photos", photos)

	if s.Res.PreferredLanguage != nil {
		s.D.Set("preferred_language", *s.Res.PreferredLanguage)
	}

	if s.Res.ProfileUrl != nil {
		s.D.Set("profile_url", *s.Res.ProfileUrl)
	}

	roles := []interface{}{}
	for _, item := range s.Res.Roles {
		roles = append(roles, UserRolesToMap(item))
	}
	s.D.Set("roles", roles)

	s.D.Set("schemas", s.Res.Schemas)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.Timezone != nil {
		s.D.Set("timezone", *s.Res.Timezone)
	}

	if s.Res.Title != nil {
		s.D.Set("title", *s.Res.Title)
	}

	if s.Res.UrnIetfParamsScimSchemasExtensionEnterprise2_0User != nil {
		s.D.Set("urnietfparamsscimschemasextensionenterprise20user", []interface{}{ExtensionEnterprise20UserToMap(s.Res.UrnIetfParamsScimSchemasExtensionEnterprise2_0User)})
	} else {
		s.D.Set("urnietfparamsscimschemasextensionenterprise20user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", []interface{}{ExtensionOCITagsToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionAdaptiveUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionadaptive_user", []interface{}{ExtensionAdaptiveUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionAdaptiveUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionadaptive_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionCapabilitiesUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensioncapabilities_user", []interface{}{ExtensionCapabilitiesUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionCapabilitiesUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensioncapabilities_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbCredentialsUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondb_credentials_user", []interface{}{ExtensionDbCredentialsUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbCredentialsUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondb_credentials_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbUserUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondb_user_user", []interface{}{ExtensionDbUserUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbUserUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondb_user_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosUserUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionkerberos_user_user", []interface{}{ExtensionKerberosUserUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosUserUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionkerberos_user_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionMfaUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmfa_user", []interface{}{ExtensionMfaUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionMfaUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmfa_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordStateUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionpassword_state_user", []interface{}{ExtensionPasswordStateUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordStateUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionpassword_state_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordlessUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionpasswordless_user", []interface{}{ExtensionPasswordlessUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordlessUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionpasswordless_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionposix_user", []interface{}{ExtensionPosixUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionposix_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSecurityQuestionsUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user", []interface{}{ExtensionSecurityQuestionsUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSecurityQuestionsUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionself_change_user", []interface{}{ExtensionSelfChangeUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionself_change_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfRegistrationUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionself_registration_user", []interface{}{ExtensionSelfRegistrationUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfRegistrationUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionself_registration_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSffUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsff_user", []interface{}{ExtensionSffUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSffUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsff_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialAccountUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsocial_account_user", []interface{}{ExtensionSocialAccountUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialAccountUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsocial_account_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionTermsOfUseUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionterms_of_use_user", []interface{}{ExtensionTermsOfUseUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionTermsOfUseUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionterms_of_use_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserCredentialsUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_credentials_user", []interface{}{ExtensionUserCredentialsUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserCredentialsUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_credentials_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserStateUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_state_user", []interface{}{ExtensionUserStateUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserStateUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_state_user", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_user", []interface{}{ExtensionUserUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionUserUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionuser_user", nil)
	}

	if s.Res.UserName != nil {
		s.D.Set("user_name", *s.Res.UserName)
	}

	s.D.Set("user_type", s.Res.UserType)

	x509Certificates := []interface{}{}
	for _, item := range s.Res.X509Certificates {
		x509Certificates = append(x509Certificates, UserX509CertificatesToMap(item))
	}
	s.D.Set("x509certificates", x509Certificates)

	return nil
}

func GetUserCompositeId(idcsEndpoint string, userId string) string {
	//id = url.PathEscape(id)
	//idcsEndpoint = url.PathEscape(idcsEndpoint)
	userId = url.PathEscape(userId)
	compositeId := "idcsEndpoint/" + idcsEndpoint + "/users/" + userId
	return compositeId
}

func parseUserCompositeId(compositeId string) (userId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/users/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	userId, _ = url.PathUnescape(parts[3])

	return
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionAdaptiveUser(fieldKeyFormat string) (oci_identity_domains.ExtensionAdaptiveUser, error) {
	result := oci_identity_domains.ExtensionAdaptiveUser{}

	if riskLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "risk_level")); ok {
		result.RiskLevel = oci_identity_domains.ExtensionAdaptiveUserRiskLevelEnum(riskLevel.(string))
	}

	if riskScores, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "risk_scores")); ok {
		interfaces := riskScores.([]interface{})
		tmp := make([]oci_identity_domains.UserExtRiskScores, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "risk_scores"), stateDataIndex)
			converted, err := s.mapToUserExtRiskScores(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "risk_scores")) {
			result.RiskScores = tmp
		}
	}

	return result, nil
}

func ExtensionAdaptiveUserToMap(obj *oci_identity_domains.ExtensionAdaptiveUser) map[string]interface{} {
	result := map[string]interface{}{}

	result["risk_level"] = string(obj.RiskLevel)

	riskScores := []interface{}{}
	for _, item := range obj.RiskScores {
		riskScores = append(riskScores, UserExtRiskScoresToMap(item))
	}
	result["risk_scores"] = riskScores

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionCapabilitiesUser(fieldKeyFormat string) (oci_identity_domains.ExtensionCapabilitiesUser, error) {
	result := oci_identity_domains.ExtensionCapabilitiesUser{}

	if canUseApiKeys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_use_api_keys")); ok {
		tmp := canUseApiKeys.(bool)
		result.CanUseApiKeys = &tmp
	}

	if canUseAuthTokens, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_use_auth_tokens")); ok {
		tmp := canUseAuthTokens.(bool)
		result.CanUseAuthTokens = &tmp
	}

	if canUseConsole, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_use_console")); ok {
		tmp := canUseConsole.(bool)
		result.CanUseConsole = &tmp
	}

	if canUseConsolePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_use_console_password")); ok {
		tmp := canUseConsolePassword.(bool)
		result.CanUseConsolePassword = &tmp
	}

	if canUseCustomerSecretKeys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_use_customer_secret_keys")); ok {
		tmp := canUseCustomerSecretKeys.(bool)
		result.CanUseCustomerSecretKeys = &tmp
	}

	if canUseDbCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_use_db_credentials")); ok {
		tmp := canUseDbCredentials.(bool)
		result.CanUseDbCredentials = &tmp
	}

	if canUseOAuth2ClientCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_use_oauth2client_credentials")); ok {
		tmp := canUseOAuth2ClientCredentials.(bool)
		result.CanUseOAuth2ClientCredentials = &tmp
	}

	if canUseSmtpCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_use_smtp_credentials")); ok {
		tmp := canUseSmtpCredentials.(bool)
		result.CanUseSmtpCredentials = &tmp
	}

	return result, nil
}

func ExtensionCapabilitiesUserToMap(obj *oci_identity_domains.ExtensionCapabilitiesUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CanUseApiKeys != nil {
		result["can_use_api_keys"] = bool(*obj.CanUseApiKeys)
	}

	if obj.CanUseAuthTokens != nil {
		result["can_use_auth_tokens"] = bool(*obj.CanUseAuthTokens)
	}

	if obj.CanUseConsole != nil {
		result["can_use_console"] = bool(*obj.CanUseConsole)
	}

	if obj.CanUseConsolePassword != nil {
		result["can_use_console_password"] = bool(*obj.CanUseConsolePassword)
	}

	if obj.CanUseCustomerSecretKeys != nil {
		result["can_use_customer_secret_keys"] = bool(*obj.CanUseCustomerSecretKeys)
	}

	if obj.CanUseDbCredentials != nil {
		result["can_use_db_credentials"] = bool(*obj.CanUseDbCredentials)
	}

	if obj.CanUseOAuth2ClientCredentials != nil {
		result["can_use_oauth2client_credentials"] = bool(*obj.CanUseOAuth2ClientCredentials)
	}

	if obj.CanUseSmtpCredentials != nil {
		result["can_use_smtp_credentials"] = bool(*obj.CanUseSmtpCredentials)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionDbCredentialsUser(fieldKeyFormat string) (oci_identity_domains.ExtensionDbCredentialsUser, error) {
	result := oci_identity_domains.ExtensionDbCredentialsUser{}

	if dbLoginAttempts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_login_attempts")); ok {
		tmp := dbLoginAttempts.(int)
		result.DbLoginAttempts = &tmp
	}

	if dbUserName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_user_name")); ok {
		tmp := dbUserName.(string)
		result.DbUserName = &tmp
	}

	return result, nil
}

func ExtensionDbCredentialsUserToMap(obj *oci_identity_domains.ExtensionDbCredentialsUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbLoginAttempts != nil {
		result["db_login_attempts"] = int(*obj.DbLoginAttempts)
	}

	if obj.DbUserName != nil {
		result["db_user_name"] = string(*obj.DbUserName)
	}

	return result
}

func ExtensionDbUserUserToMap(obj *oci_identity_domains.ExtensionDbUserUser) map[string]interface{} {
	result := map[string]interface{}{}

	result["db_global_roles"] = obj.DbGlobalRoles

	if obj.DomainLevelSchema != nil {
		result["domain_level_schema"] = string(*obj.DomainLevelSchema)
	}

	if obj.InstanceLevelSchema != nil {
		result["instance_level_schema"] = string(*obj.InstanceLevelSchema)
	}

	if obj.IsDbUser != nil {
		result["is_db_user"] = bool(*obj.IsDbUser)
	}

	passwordVerifiers := []interface{}{}
	for _, item := range obj.PasswordVerifiers {
		passwordVerifiers = append(passwordVerifiers, UserExtPasswordVerifiersToMap(item))
	}
	result["password_verifiers"] = passwordVerifiers

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionEnterprise20User(fieldKeyFormat string) (oci_identity_domains.ExtensionEnterprise20User, error) {
	result := oci_identity_domains.ExtensionEnterprise20User{}

	if costCenter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cost_center")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "cost_center")) {
			tmp := costCenter.(string)
			result.CostCenter = &tmp
		}
	}

	if department, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "department")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "department")) {
			tmp := department.(string)
			result.Department = &tmp
		}
	}

	if division, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "division")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "division")) {
			tmp := division.(string)
			result.Division = &tmp
		}
	}

	if employeeNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "employee_number")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "employee_number")) {
			tmp := employeeNumber.(string)
			result.EmployeeNumber = &tmp
		}
	}

	if manager, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "manager")); ok {
		if tmpList := manager.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "manager"), 0)
			tmp, err := s.mapToUserExtManager(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert manager, encountered error: %v", err)
			}
			result.Manager = &tmp
		}
	}

	if organization, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "organization")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "organization")) {
			tmp := organization.(string)
			result.Organization = &tmp
		}
	}

	return result, nil
}

func ExtensionEnterprise20UserToMap(obj *oci_identity_domains.ExtensionEnterprise20User) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CostCenter != nil {
		result["cost_center"] = string(*obj.CostCenter)
	}

	if obj.Department != nil {
		result["department"] = string(*obj.Department)
	}

	if obj.Division != nil {
		result["division"] = string(*obj.Division)
	}

	if obj.EmployeeNumber != nil {
		result["employee_number"] = string(*obj.EmployeeNumber)
	}

	if obj.Manager != nil {
		result["manager"] = []interface{}{UserExtManagerToMap(obj.Manager)}
	}

	if obj.Organization != nil {
		result["organization"] = string(*obj.Organization)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionKerberosUserUser(fieldKeyFormat string) (oci_identity_domains.ExtensionKerberosUserUser, error) {
	result := oci_identity_domains.ExtensionKerberosUserUser{}

	if realmUsers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "realm_users")); ok {
		interfaces := realmUsers.([]interface{})
		tmp := make([]oci_identity_domains.UserExtRealmUsers, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "realm_users"), stateDataIndex)
			converted, err := s.mapToUserExtRealmUsers(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "realm_users")) {
			result.RealmUsers = tmp
		}
	}

	return result, nil
}

func ExtensionKerberosUserUserToMap(obj *oci_identity_domains.ExtensionKerberosUserUser) map[string]interface{} {
	result := map[string]interface{}{}

	realmUsers := []interface{}{}
	for _, item := range obj.RealmUsers {
		realmUsers = append(realmUsers, UserExtRealmUsersToMap(item))
	}
	result["realm_users"] = realmUsers

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionMfaUser(fieldKeyFormat string) (oci_identity_domains.ExtensionMfaUser, error) {
	result := oci_identity_domains.ExtensionMfaUser{}

	if bypassCodes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bypass_codes")); ok {
		interfaces := bypassCodes.([]interface{})
		tmp := make([]oci_identity_domains.UserExtBypassCodes, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "bypass_codes"), stateDataIndex)
			converted, err := s.mapToUserExtBypassCodes(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "bypass_codes")) {
			result.BypassCodes = tmp
		}
	}

	if devices, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "devices")); ok {
		interfaces := devices.([]interface{})
		tmp := make([]oci_identity_domains.UserExtDevices, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "devices"), stateDataIndex)
			converted, err := s.mapToUserExtDevices(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "devices")) {
			result.Devices = tmp
		}
	}

	if loginAttempts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "login_attempts")); ok {
		tmp := loginAttempts.(int)
		result.LoginAttempts = &tmp
	}

	if mfaEnabledOn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mfa_enabled_on")); ok {
		tmp := mfaEnabledOn.(string)
		result.MfaEnabledOn = &tmp
	}

	if mfaIgnoredApps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mfa_ignored_apps")); ok {
		interfaces := mfaIgnoredApps.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "mfa_ignored_apps")) {
			result.MfaIgnoredApps = tmp
		}
	}

	if mfaStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mfa_status")); ok {
		result.MfaStatus = oci_identity_domains.ExtensionMfaUserMfaStatusEnum(mfaStatus.(string))
	}

	if preferredAuthenticationFactor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preferred_authentication_factor")); ok {
		result.PreferredAuthenticationFactor = oci_identity_domains.ExtensionMfaUserPreferredAuthenticationFactorEnum(preferredAuthenticationFactor.(string))
	}

	if preferredAuthenticationMethod, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preferred_authentication_method")); ok {
		tmp := preferredAuthenticationMethod.(string)
		result.PreferredAuthenticationMethod = &tmp
	}

	if preferredDevice, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preferred_device")); ok {
		if tmpList := preferredDevice.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "preferred_device"), 0)
			tmp, err := s.mapToUserExtPreferredDevice(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert preferred_device, encountered error: %v", err)
			}
			result.PreferredDevice = &tmp
		}
	}

	if preferredThirdPartyVendor, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preferred_third_party_vendor")); ok {
		tmp := preferredThirdPartyVendor.(string)
		result.PreferredThirdPartyVendor = &tmp
	}

	if trustedUserAgents, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trusted_user_agents")); ok {
		interfaces := trustedUserAgents.([]interface{})
		tmp := make([]oci_identity_domains.UserExtTrustedUserAgents, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "trusted_user_agents"), stateDataIndex)
			converted, err := s.mapToUserExtTrustedUserAgents(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "trusted_user_agents")) {
			result.TrustedUserAgents = tmp
		}
	}

	return result, nil
}

func ExtensionMfaUserToMap(obj *oci_identity_domains.ExtensionMfaUser) map[string]interface{} {
	result := map[string]interface{}{}

	bypassCodes := []interface{}{}
	for _, item := range obj.BypassCodes {
		bypassCodes = append(bypassCodes, UserExtBypassCodesToMap(item))
	}
	result["bypass_codes"] = bypassCodes

	devices := []interface{}{}
	for _, item := range obj.Devices {
		devices = append(devices, UserExtDevicesToMap(item))
	}
	result["devices"] = devices

	if obj.LoginAttempts != nil {
		result["login_attempts"] = int(*obj.LoginAttempts)
	}

	if obj.MfaEnabledOn != nil {
		result["mfa_enabled_on"] = string(*obj.MfaEnabledOn)
	}

	result["mfa_ignored_apps"] = obj.MfaIgnoredApps

	result["mfa_status"] = string(obj.MfaStatus)

	result["preferred_authentication_factor"] = string(obj.PreferredAuthenticationFactor)

	if obj.PreferredAuthenticationMethod != nil {
		result["preferred_authentication_method"] = string(*obj.PreferredAuthenticationMethod)
	}

	if obj.PreferredDevice != nil {
		result["preferred_device"] = []interface{}{UserExtPreferredDeviceToMap(obj.PreferredDevice)}
	}

	if obj.PreferredThirdPartyVendor != nil {
		result["preferred_third_party_vendor"] = string(*obj.PreferredThirdPartyVendor)
	}

	trustedUserAgents := []interface{}{}
	for _, item := range obj.TrustedUserAgents {
		trustedUserAgents = append(trustedUserAgents, UserExtTrustedUserAgentsToMap(item))
	}
	result["trusted_user_agents"] = trustedUserAgents

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionOCITags(fieldKeyFormat string) (oci_identity_domains.ExtensionOciTags, error) {
	result := oci_identity_domains.ExtensionOciTags{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		interfaces := definedTags.([]interface{})
		tmp := make([]oci_identity_domains.DefinedTags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "defined_tags"), stateDataIndex)
			converted, err := s.mapTodefinedTags(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "defined_tags")) {
			result.DefinedTags = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		interfaces := freeformTags.([]interface{})
		tmp := make([]oci_identity_domains.FreeformTags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "freeform_tags"), stateDataIndex)
			converted, err := s.mapTofreeformTags(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "freeform_tags")) {
			result.FreeformTags = tmp
		}
	}

	if tagSlug, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag_slug")); ok {
		result.TagSlug = &tagSlug
	}

	return result, nil
}

func ExtensionOCITagsToMap(obj *oci_identity_domains.ExtensionOciTags) map[string]interface{} {
	result := map[string]interface{}{}

	definedTags := []interface{}{}
	for _, item := range obj.DefinedTags {
		definedTags = append(definedTags, definedTagsToMap(item))
	}
	result["defined_tags"] = definedTags

	freeformTags := []interface{}{}
	for _, item := range obj.FreeformTags {
		freeformTags = append(freeformTags, freeformTagsToMap(item))
	}
	result["freeform_tags"] = freeformTags

	if obj.TagSlug != nil {
		result["tag_slug"] = fmt.Sprintf("%v", *obj.TagSlug)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapTodefinedTags(fieldKeyFormat string) (oci_identity_domains.DefinedTags, error) {
	result := oci_identity_domains.DefinedTags{}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func definedTagsToMap(obj oci_identity_domains.DefinedTags) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func ExtensionPasswordStateUserToMap(obj *oci_identity_domains.ExtensionPasswordStateUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicablePasswordPolicy != nil {
		result["applicable_password_policy"] = []interface{}{UserExtApplicablePasswordPolicyToMap(obj.ApplicablePasswordPolicy)}
	}

	if obj.CantChange != nil {
		result["cant_change"] = bool(*obj.CantChange)
	}

	if obj.CantExpire != nil {
		result["cant_expire"] = bool(*obj.CantExpire)
	}

	if obj.Expired != nil {
		result["expired"] = bool(*obj.Expired)
	}

	if obj.LastFailedValidationDate != nil {
		result["last_failed_validation_date"] = string(*obj.LastFailedValidationDate)
	}

	if obj.LastSuccessfulSetDate != nil {
		result["last_successful_set_date"] = string(*obj.LastSuccessfulSetDate)
	}

	if obj.LastSuccessfulValidationDate != nil {
		result["last_successful_validation_date"] = string(*obj.LastSuccessfulValidationDate)
	}

	if obj.MustChange != nil {
		result["must_change"] = bool(*obj.MustChange)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionPasswordlessUser(fieldKeyFormat string) (oci_identity_domains.ExtensionPasswordlessUser, error) {
	result := oci_identity_domains.ExtensionPasswordlessUser{}

	if factorIdentifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "factor_identifier")); ok {
		if tmpList := factorIdentifier.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "factor_identifier"), 0)
			tmp, err := s.mapToUserExtFactorIdentifier(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert factor_identifier, encountered error: %v", err)
			}
			result.FactorIdentifier = &tmp
		}
	}

	if factorMethod, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "factor_method")); ok {
		tmp := factorMethod.(string)
		result.FactorMethod = &tmp
	}

	if factorType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "factor_type")); ok {
		result.FactorType = oci_identity_domains.ExtensionPasswordlessUserFactorTypeEnum(factorType.(string))
	}

	return result, nil
}

func ExtensionPasswordlessUserToMap(obj *oci_identity_domains.ExtensionPasswordlessUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FactorIdentifier != nil {
		result["factor_identifier"] = []interface{}{UserExtFactorIdentifierToMap(obj.FactorIdentifier)}
	}

	if obj.FactorMethod != nil {
		result["factor_method"] = string(*obj.FactorMethod)
	}

	result["factor_type"] = string(obj.FactorType)

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionPosixUser(fieldKeyFormat string) (oci_identity_domains.ExtensionPosixUser, error) {
	result := oci_identity_domains.ExtensionPosixUser{}

	if gecos, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "gecos")); ok {
		tmp := gecos.(string)
		result.Gecos = &tmp
	}

	if gidNumber, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "gid_number")); ok {
		tmp := gidNumber.(int)
		result.GidNumber = &tmp
	}

	if homeDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "home_directory")); ok {
		tmp := homeDirectory.(string)
		result.HomeDirectory = &tmp
	}

	if loginShell, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "login_shell")); ok {
		tmp := loginShell.(string)
		result.LoginShell = &tmp
	}

	if uidNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "uid_number")); ok {
		tmp := uidNumber.(int)
		result.UidNumber = &tmp
	}

	return result, nil
}

func ExtensionPosixUserToMap(obj *oci_identity_domains.ExtensionPosixUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Gecos != nil {
		result["gecos"] = string(*obj.Gecos)
	}

	if obj.GidNumber != nil {
		result["gid_number"] = int(*obj.GidNumber)
	}

	if obj.HomeDirectory != nil {
		result["home_directory"] = string(*obj.HomeDirectory)
	}

	if obj.LoginShell != nil {
		result["login_shell"] = string(*obj.LoginShell)
	}

	if obj.UidNumber != nil {
		result["uid_number"] = int(*obj.UidNumber)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionSecurityQuestionsUser(fieldKeyFormat string) (oci_identity_domains.ExtensionSecurityQuestionsUser, error) {
	result := oci_identity_domains.ExtensionSecurityQuestionsUser{}

	if secQuestions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sec_questions")); ok {
		interfaces := secQuestions.([]interface{})
		tmp := make([]oci_identity_domains.UserExtSecQuestions, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "sec_questions"), stateDataIndex)
			converted, err := s.mapToUserExtSecQuestions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sec_questions")) {
			result.SecQuestions = tmp
		}
	}

	return result, nil
}

func ExtensionSecurityQuestionsUserToMap(obj *oci_identity_domains.ExtensionSecurityQuestionsUser) map[string]interface{} {
	result := map[string]interface{}{}

	secQuestions := []interface{}{}
	for _, item := range obj.SecQuestions {
		secQuestions = append(secQuestions, UserExtSecQuestionsToMap(item))
	}
	result["sec_questions"] = secQuestions

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionSelfChangeUser(fieldKeyFormat string) (oci_identity_domains.ExtensionSelfChangeUser, error) {
	result := oci_identity_domains.ExtensionSelfChangeUser{}

	if allowSelfChange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_self_change")); ok {
		tmp := allowSelfChange.(bool)
		result.AllowSelfChange = &tmp
	}

	return result, nil
}

func ExtensionSelfChangeUserToMap(obj *oci_identity_domains.ExtensionSelfChangeUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllowSelfChange != nil {
		result["allow_self_change"] = bool(*obj.AllowSelfChange)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionSelfRegistrationUser(fieldKeyFormat string) (oci_identity_domains.ExtensionSelfRegistrationUser, error) {
	result := oci_identity_domains.ExtensionSelfRegistrationUser{}

	if consentGranted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "consent_granted")); ok {
		tmp := consentGranted.(bool)
		result.ConsentGranted = &tmp
	}

	if selfRegistrationProfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "self_registration_profile")); ok {
		if tmpList := selfRegistrationProfile.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "self_registration_profile"), 0)
			tmp, err := s.mapToUserExtSelfRegistrationProfile(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert self_registration_profile, encountered error: %v", err)
			}
			result.SelfRegistrationProfile = &tmp
		}
	}

	if userToken, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_token")); ok {
		tmp := userToken.(string)
		result.UserToken = &tmp
	}

	return result, nil
}

func ExtensionSelfRegistrationUserToMap(obj *oci_identity_domains.ExtensionSelfRegistrationUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConsentGranted != nil {
		result["consent_granted"] = bool(*obj.ConsentGranted)
	}

	if obj.SelfRegistrationProfile != nil {
		result["self_registration_profile"] = []interface{}{UserExtSelfRegistrationProfileToMap(obj.SelfRegistrationProfile)}
	}

	if obj.UserToken != nil {
		result["user_token"] = string(*obj.UserToken)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionSffUser(fieldKeyFormat string) (oci_identity_domains.ExtensionSffUser, error) {
	result := oci_identity_domains.ExtensionSffUser{}

	if sffAuthKeys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sff_auth_keys")); ok {
		tmp := sffAuthKeys.(string)
		result.SffAuthKeys = &tmp
	}

	return result, nil
}

func ExtensionSffUserToMap(obj *oci_identity_domains.ExtensionSffUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SffAuthKeys != nil {
		result["sff_auth_keys"] = string(*obj.SffAuthKeys)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionSocialAccountUser(fieldKeyFormat string) (oci_identity_domains.ExtensionSocialAccountUser, error) {
	result := oci_identity_domains.ExtensionSocialAccountUser{}

	if socialAccounts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "social_accounts")); ok {
		interfaces := socialAccounts.([]interface{})
		tmp := make([]oci_identity_domains.UserExtSocialAccounts, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "social_accounts"), stateDataIndex)
			converted, err := s.mapToUserExtSocialAccounts(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "social_accounts")) {
			result.SocialAccounts = tmp
		}
	}

	return result, nil
}

func ExtensionSocialAccountUserToMap(obj *oci_identity_domains.ExtensionSocialAccountUser) map[string]interface{} {
	result := map[string]interface{}{}

	socialAccounts := []interface{}{}
	for _, item := range obj.SocialAccounts {
		socialAccounts = append(socialAccounts, UserExtSocialAccountsToMap(item))
	}
	result["social_accounts"] = socialAccounts

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionTermsOfUseUser(fieldKeyFormat string) (oci_identity_domains.ExtensionTermsOfUseUser, error) {
	result := oci_identity_domains.ExtensionTermsOfUseUser{}

	if termsOfUseConsents, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "terms_of_use_consents")); ok {
		interfaces := termsOfUseConsents.([]interface{})
		tmp := make([]oci_identity_domains.UserExtTermsOfUseConsents, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "terms_of_use_consents"), stateDataIndex)
			converted, err := s.mapToUserExtTermsOfUseConsents(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "terms_of_use_consents")) {
			result.TermsOfUseConsents = tmp
		}
	}

	return result, nil
}

func ExtensionTermsOfUseUserToMap(obj *oci_identity_domains.ExtensionTermsOfUseUser) map[string]interface{} {
	result := map[string]interface{}{}

	termsOfUseConsents := []interface{}{}
	for _, item := range obj.TermsOfUseConsents {
		termsOfUseConsents = append(termsOfUseConsents, UserExtTermsOfUseConsentsToMap(item))
	}
	result["terms_of_use_consents"] = termsOfUseConsents

	return result
}

func ExtensionUserCredentialsUserToMap(obj *oci_identity_domains.ExtensionUserCredentialsUser) map[string]interface{} {
	result := map[string]interface{}{}

	apiKeys := []interface{}{}
	for _, item := range obj.ApiKeys {
		apiKeys = append(apiKeys, UserExtApiKeysToMap(item))
	}
	result["api_keys"] = apiKeys

	authTokens := []interface{}{}
	for _, item := range obj.AuthTokens {
		authTokens = append(authTokens, UserExtAuthTokensToMap(item))
	}
	result["auth_tokens"] = authTokens

	customerSecretKeys := []interface{}{}
	for _, item := range obj.CustomerSecretKeys {
		customerSecretKeys = append(customerSecretKeys, UserExtCustomerSecretKeysToMap(item))
	}
	result["customer_secret_keys"] = customerSecretKeys

	dbCredentials := []interface{}{}
	for _, item := range obj.DbCredentials {
		dbCredentials = append(dbCredentials, UserExtDbCredentialsToMap(item))
	}
	result["db_credentials"] = dbCredentials

	oAuth2ClientCredentials := []interface{}{}
	for _, item := range obj.OAuth2ClientCredentials {
		oAuth2ClientCredentials = append(oAuth2ClientCredentials, UserExtOAuth2ClientCredentialsToMap(item))
	}
	result["o_auth2client_credentials"] = oAuth2ClientCredentials

	smtpCredentials := []interface{}{}
	for _, item := range obj.SmtpCredentials {
		smtpCredentials = append(smtpCredentials, UserExtSmtpCredentialsToMap(item))
	}
	result["smtp_credentials"] = smtpCredentials

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionUserStateUser(fieldKeyFormat string) (oci_identity_domains.ExtensionUserStateUser, error) {
	result := oci_identity_domains.ExtensionUserStateUser{}

	if lastFailedLoginDate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_failed_login_date")); ok {
		tmp := lastFailedLoginDate.(string)
		if tmp != "" {
			result.LastFailedLoginDate = &tmp
		}
	}

	if lastSuccessfulLoginDate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_successful_login_date")); ok {
		tmp := lastSuccessfulLoginDate.(string)
		if tmp != "" {
			result.LastSuccessfulLoginDate = &tmp
		}
	}

	if locked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "locked")); ok {
		if tmpList := locked.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "locked"), 0)
			tmp, err := s.mapToUserExtLocked(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert locked, encountered error: %v", err)
			}
			result.Locked = &tmp
		}
	}

	if loginAttempts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "login_attempts")); ok {
		tmp := loginAttempts.(int)
		result.LoginAttempts = &tmp
	}

	if maxConcurrentSessions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_concurrent_sessions")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "max_concurrent_sessions")) {
			tmp := maxConcurrentSessions.(int)
			result.MaxConcurrentSessions = &tmp
		}
	}

	if previousSuccessfulLoginDate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "previous_successful_login_date")); ok {
		tmp := previousSuccessfulLoginDate.(string)
		if tmp != "" {
			result.PreviousSuccessfulLoginDate = &tmp
		}
	}

	if recoveryAttempts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_attempts")); ok {
		tmp := recoveryAttempts.(int)
		result.RecoveryAttempts = &tmp
	}

	if recoveryEnrollAttempts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_enroll_attempts")); ok {
		tmp := recoveryEnrollAttempts.(int)
		result.RecoveryEnrollAttempts = &tmp
	}

	if recoveryLocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_locked")); ok {
		if tmpList := recoveryLocked.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "recovery_locked"), 0)
			tmp, err := s.mapToUserExtRecoveryLocked(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert recovery_locked, encountered error: %v", err)
			}
			result.RecoveryLocked = &tmp
		}
	}

	return result, nil
}

func ExtensionUserStateUserToMap(obj *oci_identity_domains.ExtensionUserStateUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LastFailedLoginDate != nil {
		result["last_failed_login_date"] = string(*obj.LastFailedLoginDate)
	}

	if obj.LastSuccessfulLoginDate != nil {
		result["last_successful_login_date"] = string(*obj.LastSuccessfulLoginDate)
	}

	if obj.Locked != nil {
		result["locked"] = []interface{}{UserExtLockedToMap(obj.Locked)}
	}

	if obj.LoginAttempts != nil {
		result["login_attempts"] = int(*obj.LoginAttempts)
	}

	if obj.MaxConcurrentSessions != nil {
		result["max_concurrent_sessions"] = int(*obj.MaxConcurrentSessions)
	}

	if obj.PreviousSuccessfulLoginDate != nil {
		result["previous_successful_login_date"] = string(*obj.PreviousSuccessfulLoginDate)
	}

	if obj.RecoveryAttempts != nil {
		result["recovery_attempts"] = int(*obj.RecoveryAttempts)
	}

	if obj.RecoveryEnrollAttempts != nil {
		result["recovery_enroll_attempts"] = int(*obj.RecoveryEnrollAttempts)
	}

	if obj.RecoveryLocked != nil {
		result["recovery_locked"] = []interface{}{UserExtRecoveryLockedToMap(obj.RecoveryLocked)}
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToExtensionUserUser(fieldKeyFormat string, forUpdate bool) (oci_identity_domains.ExtensionUserUser, error) {
	result := oci_identity_domains.ExtensionUserUser{}

	if userProvider, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_provider")); ok {
		result.Provider = oci_identity_domains.ExtensionUserUserProviderEnum(userProvider.(string))
	}

	if accountRecoveryRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "account_recovery_required")); ok {
		tmp := accountRecoveryRequired.(bool)
		result.AccountRecoveryRequired = &tmp
	}

	if bypassNotification, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bypass_notification")); ok {
		tmp := bypassNotification.(bool)
		result.BypassNotification = &tmp
	}

	if creationMechanism, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "creation_mechanism")); ok {
		result.CreationMechanism = oci_identity_domains.ExtensionUserUserCreationMechanismEnum(creationMechanism.(string))
	}

	if delegatedAuthenticationTargetApp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "delegated_authentication_target_app")); ok {
		if tmpList := delegatedAuthenticationTargetApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "delegated_authentication_target_app"), 0)
			tmp, err := s.mapToUserExtDelegatedAuthenticationTargetApp(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert delegated_authentication_target_app, encountered error: %v", err)
			}
			result.DelegatedAuthenticationTargetApp = &tmp
		}
	}

	if doNotShowGettingStarted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "do_not_show_getting_started")); ok {
		tmp := doNotShowGettingStarted.(bool)
		result.DoNotShowGettingStarted = &tmp
	}

	if isAccountRecoveryEnrolled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_account_recovery_enrolled")); ok {
		tmp := isAccountRecoveryEnrolled.(bool)
		result.IsAccountRecoveryEnrolled = &tmp
	}

	if isAuthenticationDelegated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_authentication_delegated")); ok {
		tmp := isAuthenticationDelegated.(bool)
		result.IsAuthenticationDelegated = &tmp
	}

	if isFederatedUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_federated_user")); ok {
		tmp := isFederatedUser.(bool)
		result.IsFederatedUser = &tmp
	}

	if isGroupMembershipNormalized, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_group_membership_normalized")); ok {
		tmp := isGroupMembershipNormalized.(bool)
		result.IsGroupMembershipNormalized = &tmp
	}

	if isGroupMembershipSyncedToUsersGroups, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_group_membership_synced_to_users_groups")); ok && !forUpdate {
		tmp := isGroupMembershipSyncedToUsersGroups.(bool)
		result.IsGroupMembershipSyncedToUsersGroups = &tmp
	}

	if notificationEmailTemplateId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "notification_email_template_id")); ok {
		tmp := notificationEmailTemplateId.(string)
		result.NotificationEmailTemplateId = &tmp
	}

	if preferredUiLandingPage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preferred_ui_landing_page")); ok {
		result.PreferredUiLandingPage = oci_identity_domains.ExtensionUserUserPreferredUiLandingPageEnum(preferredUiLandingPage.(string))
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_identity_domains.ExtensionUserUserStatusEnum(status.(string))
	}

	if syncedFromApp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "synced_from_app")); ok {
		if tmpList := syncedFromApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "synced_from_app"), 0)
			tmp, err := s.mapToUserExtSyncedFromApp(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert synced_from_app, encountered error: %v", err)
			}
			result.SyncedFromApp = &tmp
		}
	}

	if userFlowControlledByExternalClient, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_flow_controlled_by_external_client")); ok {
		tmp := userFlowControlledByExternalClient.(bool)
		result.UserFlowControlledByExternalClient = &tmp
	}

	return result, nil
}

func ExtensionUserUserToMap(obj *oci_identity_domains.ExtensionUserUser) map[string]interface{} {
	result := map[string]interface{}{}

	result["user_provider"] = string(obj.Provider)

	if obj.AccountRecoveryRequired != nil {
		result["account_recovery_required"] = bool(*obj.AccountRecoveryRequired)
	}

	accounts := []interface{}{}
	for _, item := range obj.Accounts {
		accounts = append(accounts, UserExtAccountsToMap(item))
	}
	result["accounts"] = accounts

	appRoles := []interface{}{}
	for _, item := range obj.AppRoles {
		appRoles = append(appRoles, UserExtAppRolesToMap(item))
	}
	result["app_roles"] = appRoles

	if obj.ApplicableAuthenticationTargetApp != nil {
		result["applicable_authentication_target_app"] = []interface{}{UserExtApplicableAuthenticationTargetAppToMap(obj.ApplicableAuthenticationTargetApp)}
	}

	if obj.BypassNotification != nil {
		result["bypass_notification"] = bool(*obj.BypassNotification)
	}

	result["creation_mechanism"] = string(obj.CreationMechanism)

	if obj.DelegatedAuthenticationTargetApp != nil {
		result["delegated_authentication_target_app"] = []interface{}{UserExtDelegatedAuthenticationTargetAppToMap(obj.DelegatedAuthenticationTargetApp)}
	}

	if obj.DoNotShowGettingStarted != nil {
		result["do_not_show_getting_started"] = bool(*obj.DoNotShowGettingStarted)
	}

	grants := []interface{}{}
	for _, item := range obj.Grants {
		grants = append(grants, UserExtGrantsToMap(item))
	}
	result["grants"] = grants

	if obj.GroupMembershipLastModified != nil {
		result["group_membership_last_modified"] = string(*obj.GroupMembershipLastModified)
	}

	idcsAppRolesLimitedToGroups := []interface{}{}
	for _, item := range obj.IdcsAppRolesLimitedToGroups {
		idcsAppRolesLimitedToGroups = append(idcsAppRolesLimitedToGroups, UserExtIdcsAppRolesLimitedToGroupsToMap(item))
	}
	result["idcs_app_roles_limited_to_groups"] = idcsAppRolesLimitedToGroups

	if obj.IsAccountRecoveryEnrolled != nil {
		result["is_account_recovery_enrolled"] = bool(*obj.IsAccountRecoveryEnrolled)
	}

	if obj.IsAuthenticationDelegated != nil {
		result["is_authentication_delegated"] = bool(*obj.IsAuthenticationDelegated)
	}

	if obj.IsFederatedUser != nil {
		result["is_federated_user"] = bool(*obj.IsFederatedUser)
	}

	if obj.IsGroupMembershipNormalized != nil {
		result["is_group_membership_normalized"] = bool(*obj.IsGroupMembershipNormalized)
	}

	if obj.IsGroupMembershipSyncedToUsersGroups != nil {
		result["is_group_membership_synced_to_users_groups"] = bool(*obj.IsGroupMembershipSyncedToUsersGroups)
	}

	if obj.NotificationEmailTemplateId != nil {
		result["notification_email_template_id"] = string(*obj.NotificationEmailTemplateId)
	}

	result["preferred_ui_landing_page"] = string(obj.PreferredUiLandingPage)

	result["status"] = string(obj.Status)

	supportAccounts := []interface{}{}
	for _, item := range obj.SupportAccounts {
		supportAccounts = append(supportAccounts, UserExtSupportAccountsToMap(item))
	}
	result["support_accounts"] = supportAccounts

	if obj.SyncedFromApp != nil {
		result["synced_from_app"] = []interface{}{UserExtSyncedFromAppToMap(obj.SyncedFromApp)}
	}

	if obj.UserFlowControlledByExternalClient != nil {
		result["user_flow_controlled_by_external_client"] = bool(*obj.UserFlowControlledByExternalClient)
	}

	if obj.UserToken != nil {
		result["user_token"] = []interface{}{UserExtUserTokenToMap(obj.UserToken)}
	}

	return result
}

func UserToMap(obj oci_identity_domains.User) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
	}

	addresses := []interface{}{}
	for _, item := range obj.Addresses {
		addresses = append(addresses, addressesToMap(item))
	}
	result["addresses"] = addresses

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	emails := []interface{}{}
	for _, item := range obj.Emails {
		emails = append(emails, UserEmailsToMap(item))
	}
	result["emails"] = emails

	entitlements := []interface{}{}
	for _, item := range obj.Entitlements {
		entitlements = append(entitlements, UserEntitlementsToMap(item))
	}
	result["entitlements"] = entitlements

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	groups := []interface{}{}
	for _, item := range obj.Groups {
		groups = append(groups, UserGroupsToMap(item))
	}
	result["groups"] = groups

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.IdcsLastUpgradedInRelease != nil {
		result["idcs_last_upgraded_in_release"] = string(*obj.IdcsLastUpgradedInRelease)
	}

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	ims := []interface{}{}
	for _, item := range obj.Ims {
		ims = append(ims, UserImsToMap(item))
	}
	result["ims"] = ims

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Name != nil {
		result["name"] = []interface{}{UserNameToMap(obj.Name)}
	}

	if obj.NickName != nil {
		result["nick_name"] = string(*obj.NickName)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Password != nil {
		result["password"] = string(*obj.Password)
	}

	phoneNumbers := []interface{}{}
	for _, item := range obj.PhoneNumbers {
		phoneNumbers = append(phoneNumbers, UserPhoneNumbersToMap(item))
	}
	result["phone_numbers"] = phoneNumbers

	photos := []interface{}{}
	for _, item := range obj.Photos {
		photos = append(photos, UserPhotosToMap(item))
	}
	result["photos"] = photos

	if obj.PreferredLanguage != nil {
		result["preferred_language"] = string(*obj.PreferredLanguage)
	}

	if obj.ProfileUrl != nil {
		result["profile_url"] = string(*obj.ProfileUrl)
	}

	roles := []interface{}{}
	for _, item := range obj.Roles {
		roles = append(roles, UserRolesToMap(item))
	}
	result["roles"] = roles

	result["schemas"] = obj.Schemas

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.Timezone != nil {
		result["timezone"] = string(*obj.Timezone)
	}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	if obj.UrnIetfParamsScimSchemasExtensionEnterprise2_0User != nil {
		result["urnietfparamsscimschemasextensionenterprise20user"] = []interface{}{ExtensionEnterprise20UserToMap(obj.UrnIetfParamsScimSchemasExtensionEnterprise2_0User)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags != nil {
		result["urnietfparamsscimschemasoracleidcsextension_oci_tags"] = []interface{}{ExtensionOCITagsToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionAdaptiveUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionadaptive_user"] = []interface{}{ExtensionAdaptiveUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionAdaptiveUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionCapabilitiesUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensioncapabilities_user"] = []interface{}{ExtensionCapabilitiesUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionCapabilitiesUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionDbCredentialsUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensiondb_credentials_user"] = []interface{}{ExtensionDbCredentialsUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionDbCredentialsUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionDbUserUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensiondb_user_user"] = []interface{}{ExtensionDbUserUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionDbUserUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosUserUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionkerberos_user_user"] = []interface{}{ExtensionKerberosUserUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosUserUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionMfaUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionmfa_user"] = []interface{}{ExtensionMfaUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionMfaUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordStateUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionpassword_state_user"] = []interface{}{ExtensionPasswordStateUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordStateUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordlessUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionpasswordless_user"] = []interface{}{ExtensionPasswordlessUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordlessUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionposix_user"] = []interface{}{ExtensionPosixUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSecurityQuestionsUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user"] = []interface{}{ExtensionSecurityQuestionsUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSecurityQuestionsUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionself_change_user"] = []interface{}{ExtensionSelfChangeUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfRegistrationUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionself_registration_user"] = []interface{}{ExtensionSelfRegistrationUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfRegistrationUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSffUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionsff_user"] = []interface{}{ExtensionSffUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSffUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialAccountUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionsocial_account_user"] = []interface{}{ExtensionSocialAccountUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialAccountUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionTermsOfUseUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionterms_of_use_user"] = []interface{}{ExtensionTermsOfUseUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionTermsOfUseUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionUserCredentialsUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionuser_credentials_user"] = []interface{}{ExtensionUserCredentialsUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionUserCredentialsUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionUserStateUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionuser_state_user"] = []interface{}{ExtensionUserStateUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionUserStateUser)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionUserUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionuser_user"] = []interface{}{ExtensionUserUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionUserUser)}
	}

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	result["user_type"] = string(obj.UserType)

	x509Certificates := []interface{}{}
	for _, item := range obj.X509Certificates {
		x509Certificates = append(x509Certificates, UserX509CertificatesToMap(item))
	}
	result["x509certificates"] = x509Certificates

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserEmails(fieldKeyFormat string) (oci_identity_domains.UserEmails, error) {
	result := oci_identity_domains.UserEmails{}

	if pendingVerificationData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pending_verification_data")); ok {
		tmp := pendingVerificationData.(string)
		result.PendingVerificationData = &tmp
	}

	if primary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary")); ok {
		tmp := primary.(bool)
		result.Primary = &tmp
	}

	if secondary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secondary")); ok {
		tmp := secondary.(bool)
		result.Secondary = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.UserEmailsTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	if verified, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verified")); ok {
		tmp := verified.(bool)
		result.Verified = &tmp
	}

	return result, nil
}

func UserEmailsToMap(obj oci_identity_domains.UserEmails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PendingVerificationData != nil {
		result["pending_verification_data"] = string(*obj.PendingVerificationData)
	}

	if obj.Primary != nil {
		result["primary"] = bool(*obj.Primary)
	}

	if obj.Secondary != nil {
		result["secondary"] = bool(*obj.Secondary)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	if obj.Verified != nil {
		result["verified"] = bool(*obj.Verified)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserEntitlements(fieldKeyFormat string) (oci_identity_domains.UserEntitlements, error) {
	result := oci_identity_domains.UserEntitlements{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if primary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary")); ok {
		tmp := primary.(bool)
		result.Primary = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserEntitlementsToMap(obj oci_identity_domains.UserEntitlements) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Primary != nil {
		result["primary"] = bool(*obj.Primary)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UserExtAccountsToMap(obj oci_identity_domains.UserExtAccounts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
	}

	if obj.AppId != nil {
		result["app_id"] = string(*obj.AppId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtApiKeys(fieldKeyFormat string) (oci_identity_domains.UserExtApiKeys, error) {
	result := oci_identity_domains.UserExtApiKeys{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if ocid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocid")); ok {
		tmp := ocid.(string)
		result.Ocid = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtApiKeysToMap(obj oci_identity_domains.UserExtApiKeys) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UserExtAppRolesToMap(obj oci_identity_domains.UserExtAppRoles) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminRole != nil {
		result["admin_role"] = bool(*obj.AdminRole)
	}

	if obj.AppId != nil {
		result["app_id"] = string(*obj.AppId)
	}

	if obj.AppName != nil {
		result["app_name"] = string(*obj.AppName)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.LegacyGroupName != nil {
		result["legacy_group_name"] = string(*obj.LegacyGroupName)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UserExtApplicableAuthenticationTargetAppToMap(obj *oci_identity_domains.UserExtApplicableAuthenticationTargetApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.TargetRequestTimeout != nil {
		result["target_request_timeout"] = int(*obj.TargetRequestTimeout)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtApplicablePasswordPolicy(fieldKeyFormat string) (oci_identity_domains.UserExtApplicablePasswordPolicy, error) {
	result := oci_identity_domains.UserExtApplicablePasswordPolicy{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if priority, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "priority")); ok {
		tmp := priority.(int)
		result.Priority = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtApplicablePasswordPolicyToMap(obj *oci_identity_domains.UserExtApplicablePasswordPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Priority != nil {
		result["priority"] = int(*obj.Priority)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtAuthTokens(fieldKeyFormat string) (oci_identity_domains.UserExtAuthTokens, error) {
	result := oci_identity_domains.UserExtAuthTokens{}

	if ocid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocid")); ok {
		tmp := ocid.(string)
		result.Ocid = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtAuthTokensToMap(obj oci_identity_domains.UserExtAuthTokens) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtBypassCodes(fieldKeyFormat string) (oci_identity_domains.UserExtBypassCodes, error) {
	result := oci_identity_domains.UserExtBypassCodes{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtBypassCodesToMap(obj oci_identity_domains.UserExtBypassCodes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtCustomerSecretKeys(fieldKeyFormat string) (oci_identity_domains.UserExtCustomerSecretKeys, error) {
	result := oci_identity_domains.UserExtCustomerSecretKeys{}

	if ocid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocid")); ok {
		tmp := ocid.(string)
		result.Ocid = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtCustomerSecretKeysToMap(obj oci_identity_domains.UserExtCustomerSecretKeys) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtDbCredentials(fieldKeyFormat string) (oci_identity_domains.UserExtDbCredentials, error) {
	result := oci_identity_domains.UserExtDbCredentials{}

	if ocid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocid")); ok {
		tmp := ocid.(string)
		result.Ocid = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtDbCredentialsToMap(obj oci_identity_domains.UserExtDbCredentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtDelegatedAuthenticationTargetApp(fieldKeyFormat string) (oci_identity_domains.UserExtDelegatedAuthenticationTargetApp, error) {
	result := oci_identity_domains.UserExtDelegatedAuthenticationTargetApp{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.UserExtDelegatedAuthenticationTargetAppTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtDelegatedAuthenticationTargetAppToMap(obj *oci_identity_domains.UserExtDelegatedAuthenticationTargetApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtDevices(fieldKeyFormat string) (oci_identity_domains.UserExtDevices, error) {
	result := oci_identity_domains.UserExtDevices{}

	if authenticationMethod, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication_method")); ok {
		tmp := authenticationMethod.(string)
		result.AuthenticationMethod = &tmp
	}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if factorStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "factor_status")); ok {
		tmp := factorStatus.(string)
		result.FactorStatus = &tmp
	}

	if factorType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "factor_type")); ok {
		tmp := factorType.(string)
		result.FactorType = &tmp
	}

	if lastSyncTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_sync_time")); ok {
		tmp := lastSyncTime.(string)
		result.LastSyncTime = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		tmp := status.(string)
		result.Status = &tmp
	}

	if thirdPartyVendorName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "third_party_vendor_name")); ok {
		tmp := thirdPartyVendorName.(string)
		result.ThirdPartyVendorName = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtDevicesToMap(obj oci_identity_domains.UserExtDevices) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuthenticationMethod != nil {
		result["authentication_method"] = string(*obj.AuthenticationMethod)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.FactorStatus != nil {
		result["factor_status"] = string(*obj.FactorStatus)
	}

	if obj.FactorType != nil {
		result["factor_type"] = string(*obj.FactorType)
	}

	if obj.LastSyncTime != nil {
		result["last_sync_time"] = string(*obj.LastSyncTime)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.ThirdPartyVendorName != nil {
		result["third_party_vendor_name"] = string(*obj.ThirdPartyVendorName)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtFactorIdentifier(fieldKeyFormat string) (oci_identity_domains.UserExtFactorIdentifier, error) {
	result := oci_identity_domains.UserExtFactorIdentifier{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtFactorIdentifierToMap(obj *oci_identity_domains.UserExtFactorIdentifier) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UserExtGrantsToMap(obj oci_identity_domains.UserExtGrants) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AppId != nil {
		result["app_id"] = string(*obj.AppId)
	}

	result["grant_mechanism"] = string(obj.GrantMechanism)

	if obj.GrantorId != nil {
		result["grantor_id"] = string(*obj.GrantorId)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UserExtIdcsAppRolesLimitedToGroupsToMap(obj oci_identity_domains.UserExtIdcsAppRolesLimitedToGroups) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.IdcsAppRoleId != nil {
		result["idcs_app_role_id"] = string(*obj.IdcsAppRoleId)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtLocked(fieldKeyFormat string) (oci_identity_domains.UserExtLocked, error) {
	result := oci_identity_domains.UserExtLocked{}

	if expired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "expired")); ok {
		tmp := expired.(bool)
		result.Expired = &tmp
	}

	if lockDate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lock_date")); ok {
		tmp := lockDate.(string)
		if tmp != "" {
			result.LockDate = &tmp
		}
	}

	if on, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on")); ok {
		tmp := on.(bool)
		result.On = &tmp
	}

	if reason, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reason")); ok {
		if *result.On == true {
			tmp := reason.(int)
			result.Reason = &tmp
		}
	}

	return result, nil
}

func UserExtLockedToMap(obj *oci_identity_domains.UserExtLocked) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Expired != nil {
		result["expired"] = bool(*obj.Expired)
	}

	if obj.LockDate != nil {
		result["lock_date"] = string(*obj.LockDate)
	}

	if obj.On != nil {
		result["on"] = bool(*obj.On)
	}

	if obj.Reason != nil {
		result["reason"] = int(*obj.Reason)
	}
	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtManager(fieldKeyFormat string) (oci_identity_domains.UserExtManager, error) {
	result := oci_identity_domains.UserExtManager{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtManagerToMap(obj *oci_identity_domains.UserExtManager) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtOAuth2ClientCredentials(fieldKeyFormat string) (oci_identity_domains.UserExtOAuth2ClientCredentials, error) {
	result := oci_identity_domains.UserExtOAuth2ClientCredentials{}

	if ocid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocid")); ok {
		tmp := ocid.(string)
		result.Ocid = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtOAuth2ClientCredentialsToMap(obj oci_identity_domains.UserExtOAuth2ClientCredentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtPasswordVerifiers(fieldKeyFormat string) (oci_identity_domains.UserExtPasswordVerifiers, error) {
	result := oci_identity_domains.UserExtPasswordVerifiers{}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtPasswordVerifiersToMap(obj oci_identity_domains.UserExtPasswordVerifiers) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtPreferredDevice(fieldKeyFormat string) (oci_identity_domains.UserExtPreferredDevice, error) {
	result := oci_identity_domains.UserExtPreferredDevice{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtPreferredDeviceToMap(obj *oci_identity_domains.UserExtPreferredDevice) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtRealmUsers(fieldKeyFormat string) (oci_identity_domains.UserExtRealmUsers, error) {
	result := oci_identity_domains.UserExtRealmUsers{}

	if principalName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "principal_name")); ok {
		tmp := principalName.(string)
		result.PrincipalName = &tmp
	}

	if realmName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "realm_name")); ok {
		tmp := realmName.(string)
		result.RealmName = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtRealmUsersToMap(obj oci_identity_domains.UserExtRealmUsers) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PrincipalName != nil {
		result["principal_name"] = string(*obj.PrincipalName)
	}

	if obj.RealmName != nil {
		result["realm_name"] = string(*obj.RealmName)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtRecoveryLocked(fieldKeyFormat string) (oci_identity_domains.UserExtRecoveryLocked, error) {
	result := oci_identity_domains.UserExtRecoveryLocked{}

	if lockDate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lock_date")); ok {
		tmp := lockDate.(string)
		if tmp != "" {
			result.LockDate = &tmp
		}
	}

	if on, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on")); ok {
		tmp := on.(bool)
		result.On = &tmp
	}

	return result, nil
}

func UserExtRecoveryLockedToMap(obj *oci_identity_domains.UserExtRecoveryLocked) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LockDate != nil {
		result["lock_date"] = string(*obj.LockDate)
	}

	if obj.On != nil {
		result["on"] = bool(*obj.On)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtRiskScores(fieldKeyFormat string) (oci_identity_domains.UserExtRiskScores, error) {
	result := oci_identity_domains.UserExtRiskScores{}

	if lastUpdateTimestamp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_update_timestamp")); ok {
		tmp := lastUpdateTimestamp.(string)
		result.LastUpdateTimestamp = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if riskLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "risk_level")); ok {
		result.RiskLevel = oci_identity_domains.UserExtRiskScoresRiskLevelEnum(riskLevel.(string))
	}

	if score, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "score")); ok {
		tmp := score.(int)
		result.Score = &tmp
	}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		tmp := source.(string)
		result.Source = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		tmp := status.(string)
		result.Status = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtRiskScoresToMap(obj oci_identity_domains.UserExtRiskScores) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LastUpdateTimestamp != nil {
		result["last_update_timestamp"] = string(*obj.LastUpdateTimestamp)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["risk_level"] = string(obj.RiskLevel)

	if obj.Score != nil {
		result["score"] = int(*obj.Score)
	}

	if obj.Source != nil {
		result["source"] = string(*obj.Source)
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtSecQuestions(fieldKeyFormat string) (oci_identity_domains.UserExtSecQuestions, error) {
	result := oci_identity_domains.UserExtSecQuestions{}

	if answer, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "answer")); ok {
		tmp := answer.(string)
		result.Answer = &tmp
	}

	if hintText, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hint_text")); ok {
		tmp := hintText.(string)
		result.HintText = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtSecQuestionsToMap(obj oci_identity_domains.UserExtSecQuestions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Answer != nil {
		result["answer"] = string(*obj.Answer)
	}

	if obj.HintText != nil {
		result["hint_text"] = string(*obj.HintText)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtSelfRegistrationProfile(fieldKeyFormat string) (oci_identity_domains.UserExtSelfRegistrationProfile, error) {
	result := oci_identity_domains.UserExtSelfRegistrationProfile{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtSelfRegistrationProfileToMap(obj *oci_identity_domains.UserExtSelfRegistrationProfile) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtSmtpCredentials(fieldKeyFormat string) (oci_identity_domains.UserExtSmtpCredentials, error) {
	result := oci_identity_domains.UserExtSmtpCredentials{}

	if ocid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocid")); ok {
		tmp := ocid.(string)
		result.Ocid = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtSmtpCredentialsToMap(obj oci_identity_domains.UserExtSmtpCredentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtSocialAccounts(fieldKeyFormat string) (oci_identity_domains.UserExtSocialAccounts, error) {
	result := oci_identity_domains.UserExtSocialAccounts{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtSocialAccountsToMap(obj oci_identity_domains.UserExtSocialAccounts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UserExtSupportAccountsToMap(obj oci_identity_domains.UserExtSupportAccounts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Provider != nil {
		result["user_provider"] = string(*obj.Provider)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.UserId != nil {
		result["user_id"] = string(*obj.UserId)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtSyncedFromApp(fieldKeyFormat string) (oci_identity_domains.UserExtSyncedFromApp, error) {
	result := oci_identity_domains.UserExtSyncedFromApp{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.UserExtSyncedFromAppTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtSyncedFromAppToMap(obj *oci_identity_domains.UserExtSyncedFromApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtTermsOfUseConsents(fieldKeyFormat string) (oci_identity_domains.UserExtTermsOfUseConsents, error) {
	result := oci_identity_domains.UserExtTermsOfUseConsents{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtTermsOfUseConsentsToMap(obj oci_identity_domains.UserExtTermsOfUseConsents) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserExtTrustedUserAgents(fieldKeyFormat string) (oci_identity_domains.UserExtTrustedUserAgents, error) {
	result := oci_identity_domains.UserExtTrustedUserAgents{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserExtTrustedUserAgentsToMap(obj oci_identity_domains.UserExtTrustedUserAgents) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UserExtUserTokenToMap(obj *oci_identity_domains.UserExtUserToken) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UserGroupsToMap(obj oci_identity_domains.UserGroups) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DateAdded != nil {
		result["date_added"] = string(*obj.DateAdded)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	if obj.MembershipOcid != nil {
		result["membership_ocid"] = string(*obj.MembershipOcid)
	}

	if obj.NonUniqueDisplay != nil {
		result["non_unique_display"] = string(*obj.NonUniqueDisplay)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserIms(fieldKeyFormat string) (oci_identity_domains.UserIms, error) {
	result := oci_identity_domains.UserIms{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "display")) {
			tmp := display.(string)
			result.Display = &tmp
		}
	}

	if primary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "primary")) {
			tmp := primary.(bool)
			result.Primary = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.UserImsTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "value")) {
			tmp := value.(string)
			result.Value = &tmp
		}
	}

	return result, nil
}

func UserImsToMap(obj oci_identity_domains.UserIms) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Primary != nil {
		result["primary"] = bool(*obj.Primary)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserName(fieldKeyFormat string) (oci_identity_domains.UserName, error) {
	result := oci_identity_domains.UserName{}

	if familyName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "family_name")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "family_name")) {
			tmp := familyName.(string)
			result.FamilyName = &tmp
		}
	}

	if formatted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "formatted")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "formatted")) {
			tmp := formatted.(string)
			result.Formatted = &tmp
		}
	}

	if givenName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "given_name")); ok {
		// if it's optional with an empty value ("", 0), don't set it to result
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "given_name")) {
			tmp := givenName.(string)
			result.GivenName = &tmp
		}
	}

	if honorificPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "honorific_prefix")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "honorific_prefix")) {
			tmp := honorificPrefix.(string)
			result.HonorificPrefix = &tmp
		}
	}

	if honorificSuffix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "honorific_suffix")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "honorific_suffix")) {
			tmp := honorificSuffix.(string)
			result.HonorificSuffix = &tmp
		}
	}

	if middleName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "middle_name")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "middle_name")) {
			tmp := middleName.(string)
			result.MiddleName = &tmp
		}
	}

	return result, nil
}

func UserNameToMap(obj *oci_identity_domains.UserName) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FamilyName != nil {
		result["family_name"] = string(*obj.FamilyName)
	}

	if obj.Formatted != nil {
		result["formatted"] = string(*obj.Formatted)
	}

	if obj.GivenName != nil {
		result["given_name"] = string(*obj.GivenName)
	}

	if obj.HonorificPrefix != nil {
		result["honorific_prefix"] = string(*obj.HonorificPrefix)
	}

	if obj.HonorificSuffix != nil {
		result["honorific_suffix"] = string(*obj.HonorificSuffix)
	}

	if obj.MiddleName != nil {
		result["middle_name"] = string(*obj.MiddleName)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserPhoneNumbers(fieldKeyFormat string) (oci_identity_domains.UserPhoneNumbers, error) {
	result := oci_identity_domains.UserPhoneNumbers{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if primary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary")); ok {
		tmp := primary.(bool)
		result.Primary = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.UserPhoneNumbersTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	if verified, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verified")); ok {
		tmp := verified.(bool)
		result.Verified = &tmp
	}

	return result, nil
}

func UserPhoneNumbersToMap(obj oci_identity_domains.UserPhoneNumbers) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Primary != nil {
		result["primary"] = bool(*obj.Primary)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	if obj.Verified != nil {
		result["verified"] = bool(*obj.Verified)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserPhotos(fieldKeyFormat string) (oci_identity_domains.UserPhotos, error) {
	result := oci_identity_domains.UserPhotos{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if primary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary")); ok {
		tmp := primary.(bool)
		result.Primary = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.UserPhotosTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserPhotosToMap(obj oci_identity_domains.UserPhotos) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Primary != nil {
		result["primary"] = bool(*obj.Primary)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserRoles(fieldKeyFormat string) (oci_identity_domains.UserRoles, error) {
	result := oci_identity_domains.UserRoles{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if primary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary")); ok {
		tmp := primary.(bool)
		result.Primary = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func UserRolesToMap(obj oci_identity_domains.UserRoles) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Primary != nil {
		result["primary"] = bool(*obj.Primary)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToUserX509Certificates(fieldKeyFormat string) (oci_identity_domains.UserX509Certificates, error) {
	result := oci_identity_domains.UserX509Certificates{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if primary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary")); ok {
		tmp := primary.(bool)
		result.Primary = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		result.Value = &value
	}

	return result, nil
}

func UserX509CertificatesToMap(obj oci_identity_domains.UserX509Certificates) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Primary != nil {
		result["primary"] = bool(*obj.Primary)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = fmt.Sprintf("%v", *obj.Value)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapToaddresses(fieldKeyFormat string) (oci_identity_domains.Addresses, error) {
	result := oci_identity_domains.Addresses{}

	if country, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "country")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "country")) {
			tmp := country.(string)
			result.Country = &tmp
		}
	}

	if formatted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "formatted")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "formatted")) {
			tmp := formatted.(string)
			result.Formatted = &tmp
		}
	}

	if locality, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "locality")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "locality")) {
			tmp := locality.(string)
			result.Locality = &tmp
		}
	}

	if postalCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "postal_code")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "postal_code")) {
			tmp := postalCode.(string)
			result.PostalCode = &tmp
		}
	}

	if primary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "primary")) {
			tmp := primary.(bool)
			result.Primary = &tmp
		}
	}

	if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "region")) {
			tmp := region.(string)
			result.Region = &tmp
		}
	}

	if streetAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "street_address")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsUserResource(), s.D, fmt.Sprintf(fieldKeyFormat, "street_address")) {
			tmp := streetAddress.(string)
			result.StreetAddress = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.AddressesTypeEnum(type_.(string))
	}

	return result, nil
}

func addressesToMap(obj oci_identity_domains.Addresses) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Country != nil {
		result["country"] = string(*obj.Country)
	}

	if obj.Formatted != nil {
		result["formatted"] = string(*obj.Formatted)
	}

	if obj.Locality != nil {
		result["locality"] = string(*obj.Locality)
	}

	if obj.PostalCode != nil {
		result["postal_code"] = string(*obj.PostalCode)
	}

	if obj.Primary != nil {
		result["primary"] = bool(*obj.Primary)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.StreetAddress != nil {
		result["street_address"] = string(*obj.StreetAddress)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapTofreeformTags(fieldKeyFormat string) (oci_identity_domains.FreeformTags, error) {
	result := oci_identity_domains.FreeformTags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func freeformTagsToMap(obj oci_identity_domains.FreeformTags) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func idcsCreatedByToMap(obj *oci_identity_domains.IdcsCreatedBy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func idcsLastModifiedByToMap(obj *oci_identity_domains.IdcsLastModifiedBy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func metaToMap(obj *oci_identity_domains.Meta) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Created != nil {
		result["created"] = string(*obj.Created)
	}

	if obj.LastModified != nil {
		result["last_modified"] = string(*obj.LastModified)
	}

	if obj.Location != nil {
		result["location"] = string(*obj.Location)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func (s *IdentityDomainsUserResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
	result := oci_identity_domains.Tags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func tagsToMap(obj oci_identity_domains.Tags) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
