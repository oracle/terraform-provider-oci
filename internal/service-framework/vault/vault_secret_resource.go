// Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vault

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_vault "github.com/oracle/oci-go-sdk/v65/vault"
)

var _ resource.Resource = &VaultSecretResource{}
var _ resource.ResourceWithConfigure = &VaultSecretResource{}

var (
	CreateTimeout = tfresource.TwentyMinutes
	ReadTimeout   = tfresource.TwentyMinutes
	UpdateTimeout = tfresource.TwentyMinutes
	DeleteTimeout = tfresource.TwentyMinutes
)

func NewVaultSecretResource() resource.Resource {
	return &VaultSecretResource{}
}

type VaultSecretResource struct {
	client *client.OracleClients
}

func (d *VaultSecretResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData != nil {
		d.client = request.ProviderData.(*client.OracleClients)
	}
}

func (d *VaultSecretResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "oci_vault_secret"
}

func (d *VaultSecretResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"compartment_id": schema.StringAttribute{
				Required: true,
			},
			"secret_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vault_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},

			// Optional
			"defined_tags": schema.MapAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				//PlanModifiers: []planmodifier.Map{
				//	mapplanmodifier.UseStateForUnknown(),
				//},
				//DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				//PlanModifiers: []planmodifier.Map{
				//	DefinedTagsPlanModifier,
				//},
			},
			"description": schema.StringAttribute{
				//CustomType: tfresource.CaseInsensitiveStringType{},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					StringEqualIgnoreCasePlanModifier(),
				},
			},
			"freeform_tags": schema.MapAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				//PlanModifiers: []planmodifier.Map{
				//	mapplanmodifier.UseStateForUnknown(),
				//},
			},
			"key_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"metadata": schema.MapAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
			},
			"current_version_number": schema.StringAttribute{
				Computed: true,
				//PlanModifiers: []planmodifier.String{
				//	stringplanmodifier.UseStateForUnknown(),
				//},
			},
			"last_rotation_time": schema.StringAttribute{
				Computed: true,
			},

			"lifecycle_details": schema.StringAttribute{
				Computed: true,
			},
			"next_rotation_time": schema.StringAttribute{
				Computed: true,
			},
			"rotation_status": schema.StringAttribute{
				Computed: true,
			},
			"state": schema.StringAttribute{
				Computed: true,
			},
			"time_created": schema.StringAttribute{
				Computed: true,
			},
			"time_of_current_version_expiry": schema.StringAttribute{
				Computed: true,
			},
			"time_of_deletion": schema.StringAttribute{
				Computed: true,
			},
		},
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
				Read:   true,
				Update: true,
				Delete: true,
			}),
			"rotation_config": schema.ListNestedBlock{
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"is_scheduled_rotation_enabled": schema.BoolAttribute{
							Optional: true,
							Computed: true,
						},
						"rotation_interval": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
					},
					Blocks: map[string]schema.Block{
						"target_system_details": schema.ListNestedBlock{
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
							},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"target_system_type": schema.StringAttribute{
										Required: true,
										// DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff, // Requires custom implementation
										Validators: []validator.String{
											stringvalidator.OneOf("ADB", "FUNCTION"),
										},
									},
									"adb_id": schema.StringAttribute{
										Optional: true,
										Computed: true,
									},
									"function_id": schema.StringAttribute{
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"secret_content": schema.ListNestedBlock{
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"content_type": schema.StringAttribute{
							Required: true,
							// DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff
							Validators: []validator.String{
								stringvalidator.OneOf("BASE64"),
							},
						},
						"content": schema.StringAttribute{
							Optional: true,
						},
						"name": schema.StringAttribute{
							Optional: true,
						},
						"stage": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"secret_rules": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"rule_type": schema.StringAttribute{
							Required: true,
							// DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff
							Validators: []validator.String{
								stringvalidator.OneOf("SECRET_EXPIRY_RULE", "SECRET_REUSE_RULE"),
							},
						},
						"is_enforced_on_deleted_secret_versions": schema.BoolAttribute{
							Optional: true,
							Computed: true,
						},
						"is_secret_content_retrieval_blocked_on_expiry": schema.BoolAttribute{
							Optional: true,
							Computed: true,
						},
						"secret_version_expiry_interval": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"time_of_absolute_expiry": schema.StringAttribute{
							Optional: true,
							Computed: true,
							// DiffSuppressFunc: tfresource.TimeDiffSuppressFunction, // Requires custom implementation
						},
					},
				},
			},
		},

		/*
			CustomizeDiff: customdiff.All(
					customdiff.ComputedIf("current_version_number", func(_ context.Context, diff *schema.ResourceDiff, meta interface{}) bool {
						return diff.HasChange("secret_content")
					}),
				),
		*/
	}
}

func StringEqualIgnoreCasePlanModifier() planmodifier.String {
	return &stringEqualIgnoreCasePlanModifier{}
}

type stringEqualIgnoreCasePlanModifier struct {
}

func (d *stringEqualIgnoreCasePlanModifier) Description(ctx context.Context) string {
	return "This attribute is case insensitive"
}

func (d *stringEqualIgnoreCasePlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *stringEqualIgnoreCasePlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.PlanValue.IsUnknown() && !req.PlanValue.IsNull() && !req.StateValue.IsNull() {
		if strings.EqualFold(req.PlanValue.ValueString(), req.StateValue.ValueString()) {
			resp.PlanValue = req.StateValue
		}
	}
}

/*
func (d *VaultSecretResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}
	var config, plan, state VaultSecretResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.Description.IsNull() && !state.Description.IsNull() {
		if strings.EqualFold(plan.Description.ValueString(), state.Description.ValueString()) {
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("description"), state.Description)...)
		}
	}
	//resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("last_rotation_time"), types.StringUnknown())...)
}
*/

type VaultSecretResourceCrud struct {
	tfresource.BaseCrudFW
	Client                 *oci_vault.VaultsClient
	Res                    *oci_vault.Secret
	DisableNotFoundRetries bool
}

type VaultSecretResourceModel struct {
	Timeouts      timeouts.Value `tfsdk:"timeouts"`
	ID            types.String   `tfsdk:"id"`
	CompartmentId types.String   `tfsdk:"compartment_id"`
	SecretName    types.String   `tfsdk:"secret_name"`
	VaultId       types.String   `tfsdk:"vault_id"`
	DefinedTags   types.Map      `tfsdk:"defined_tags"`
	//Description   tfresource.CaseInsensitiveStringValue `tfsdk:"description"`
	Description                types.String `tfsdk:"description"`
	FreeFormTags               types.Map    `tfsdk:"freeform_tags"`
	KeyId                      types.String `tfsdk:"key_id"`
	Metadata                   types.Map    `tfsdk:"metadata"`
	RotationConfig             types.List   `tfsdk:"rotation_config"`
	SecretContent              types.List   `tfsdk:"secret_content"`
	SecretRules                types.List   `tfsdk:"secret_rules"`
	CurrentVersionNumber       types.String `tfsdk:"current_version_number"`
	LastRotationTime           types.String `tfsdk:"last_rotation_time"`
	LifecycleDetails           types.String `tfsdk:"lifecycle_details"`
	NextRotationTime           types.String `tfsdk:"next_rotation_time"`
	RotationStatus             types.String `tfsdk:"rotation_status"`
	State                      types.String `tfsdk:"state"`
	TimeCreated                types.String `tfsdk:"time_created"`
	TimeOfCurrentVersionExpiry types.String `tfsdk:"time_of_current_version_expiry"`
	TimeOfDeletion             types.String `tfsdk:"time_of_deletion"`
}

type RotationConfig struct {
	IsScheduledRotationEnabled types.Bool   `tfsdk:"is_scheduled_rotation_enabled"`
	RotationInterval           types.String `tfsdk:"rotation_interval"`
	TargetSystemDetails        types.List   `tfsdk:"target_system_details"`
}

type TargetSystemDetails struct {
	AdbId            types.String `tfsdk:"adb_id"`
	FunctionId       types.String `tfsdk:"function_id"`
	TargetSystemType types.String `tfsdk:"target_system_type"`
}

type SecretContent struct {
	Content     types.String `tfsdk:"content"`
	ContentType types.String `tfsdk:"content_type"`
	Name        types.String `tfsdk:"name"`
	Stage       types.String `tfsdk:"stage"`
}

type SecretRules struct {
	IsEnforcedOnDeletedSecretVersions       types.Bool   `tfsdk:"is_enforced_on_deleted_secret_versions"`
	IsSecretContentRetrievalBlockedOnExpiry types.Bool   `tfsdk:"is_secret_content_retrieval_blocked_on_expiry"`
	RuleType                                types.String `tfsdk:"rule_type"`
	SecretVersionExpiryInterval             types.String `tfsdk:"secret_version_expiry_interval"`
	TimeOfAbsoluteExpiry                    types.String `tfsdk:"time_of_absolute_expiry"`
}

func (d *VaultSecretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	sync := &VaultSecretResourceCrud{}
	sync.Context = &ctx
	sync.Request = &req
	sync.Response = resp
	sync.RequestState = &resp.State
	sync.ResponseState = &resp.State
	sync.Client = d.client.VaultsClient()

	err := tfresource.CreateResourceFw(sync)
	if err != nil {
		resp.Diagnostics.AddError(err.Error(), "")
	}
}

func (d *VaultSecretResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	sync := &VaultSecretResourceCrud{}
	sync.Context = &ctx
	sync.Request = &req
	sync.Response = resp
	sync.RequestState = &req.State
	sync.ResponseState = &resp.State
	sync.Client = d.client.VaultsClient()

	err := tfresource.ReadResourceFw(sync)
	if err != nil {
		resp.Diagnostics.AddError(err.Error(), "")
	}
}

func (d *VaultSecretResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	sync := &VaultSecretResourceCrud{}
	sync.Context = &ctx
	sync.Request = &req
	sync.Response = resp
	sync.RequestState = &req.State
	sync.ResponseState = &resp.State
	sync.Client = d.client.VaultsClient()

	err := tfresource.UpdateResourceFw(sync)
	if err != nil {
		resp.Diagnostics.AddError(err.Error(), "")
	}
}

func (d *VaultSecretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	sync := &VaultSecretResourceCrud{}
	sync.Context = &ctx
	sync.Request = &req
	sync.Response = resp
	sync.RequestState = &req.State
	sync.ResponseState = &resp.State
	sync.Client = d.client.VaultsClient()

	err := tfresource.DeleteResourceFw(sync)
	if err != nil {
		resp.Diagnostics.AddError(err.Error(), "")
	}
}

func (d *VaultSecretResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (s *VaultSecretResourceCrud) VoidState() {
	s.ResponseState.RemoveResource(*s.Context)
}

func (s *VaultSecretResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VaultSecretResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_vault.SecretLifecycleStateCreating),
	}
}

func (s *VaultSecretResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_vault.SecretLifecycleStateActive),
	}
}

func (s *VaultSecretResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_vault.SecretLifecycleStateDeleting),
		string(oci_vault.SecretLifecycleStateSchedulingDeletion),
	}
}

func (s *VaultSecretResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_vault.SecretLifecycleStateDeleted),
		string(oci_vault.SecretLifecycleStatePendingDeletion),
	}
}

func (s *VaultSecretResourceCrud) Create() error {
	request := oci_vault.CreateSecretRequest{}
	var plan = VaultSecretResourceModel{}

	req := s.Request.(*resource.CreateRequest)
	resp := s.Response.(*resource.CreateResponse)
	diags := req.Config.Get(*s.Context, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return tfresource.DiagnosticsToError(resp.Diagnostics)
	}

	timeout, diags := plan.Timeouts.Create(*s.Context, CreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return tfresource.DiagnosticsToError(resp.Diagnostics)
	}

	s.OperationTimeout = timeout
	_, cancel := context.WithTimeout(*s.Context, timeout)
	defer cancel()

	if !plan.CompartmentId.IsNull() {
		tmp := plan.CompartmentId.ValueString()
		request.CompartmentId = &tmp
	}

	if !plan.DefinedTags.IsNull() {
		var definedTags map[string]interface{}
		diags = plan.DefinedTags.ElementsAs(*s.Context, &definedTags, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags)
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if !plan.Description.IsNull() {
		tmp := plan.Description.ValueString()
		request.Description = &tmp
	}

	if !plan.FreeFormTags.IsNull() {
		var freeformTags map[string]interface{}
		diags = plan.FreeFormTags.ElementsAs(*s.Context, &freeformTags, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags)
	}

	if !plan.KeyId.IsNull() {
		tmp := plan.KeyId.ValueString()
		request.KeyId = &tmp
	}

	if !plan.Metadata.IsNull() {
		var metadata map[string]interface{}
		diags = plan.DefinedTags.ElementsAs(*s.Context, &metadata, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		request.Metadata = metadata
	}

	if !plan.RotationConfig.IsNull() && len(plan.RotationConfig.Elements()) > 0 {
		rotationConfigCount := len(plan.RotationConfig.Elements())
		rotationConfigs := make([]RotationConfig, rotationConfigCount)
		diags = plan.RotationConfig.ElementsAs(*s.Context, &rotationConfigs, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		tmp, err := s.fwListToRotationConfig(rotationConfigs[0])
		if err != nil {
			return err
		}
		request.RotationConfig = &tmp
	}

	if !plan.SecretContent.IsNull() {
		secretContents := make([]SecretContent, 1)
		diags = plan.SecretContent.ElementsAs(*s.Context, &secretContents, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		if len(secretContents) > 0 {
			tmp, err := s.fwListToSecretContent(secretContents[0])
			if err != nil {
				return err
			}
			request.SecretContent = tmp
		}
	}

	if !plan.SecretName.IsNull() {
		tmp := plan.SecretName.ValueString()
		request.SecretName = &tmp
	}

	if !plan.SecretRules.IsNull() && len(plan.SecretRules.Elements()) > 0 {
		secretRulesCount := len(plan.SecretRules.Elements())
		secretRuless := make([]SecretRules, secretRulesCount)
		diags = plan.SecretRules.ElementsAs(*s.Context, &secretRuless, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		tmp := make([]oci_vault.SecretRule, secretRulesCount)
		for i := range secretRuless {
			converted, err := s.fwListToSecretRules(secretRuless[i])
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 {
			request.SecretRules = tmp
		}
	}

	if !plan.VaultId.IsNull() {
		tmp := plan.VaultId.ValueString()
		request.VaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vault")

	response, err := s.Client.CreateSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Secret

	plan.ID = types.StringValue(s.ID())
	diags = resp.State.Set(*s.Context, &plan)
	resp.Diagnostics.Append(diags...)
	/*diags = resp.State.SetAttribute(*s.Context, path.Root("id"), s.ID())
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(*s.Context, path.Root("secret_content"), plan.SecretContent)
	resp.Diagnostics.Append(diags...)*/
	if resp.Diagnostics.HasError() {
		return tfresource.DiagnosticsToError(resp.Diagnostics)
	}

	return nil
}

func (s *VaultSecretResourceCrud) Get() error {
	request := oci_vault.GetSecretRequest{}
	var state VaultSecretResourceModel

	diags := s.RequestState.Get(*s.Context, &state)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}

	timeout, diags := state.Timeouts.Read(*s.Context, ReadTimeout)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}

	s.OperationTimeout = timeout
	_, cancel := context.WithTimeout(*s.Context, timeout)
	defer cancel()

	tmp := state.ID.ValueString()
	request.SecretId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vault")

	response, err := s.Client.GetSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Secret
	return nil
}

func (s *VaultSecretResourceCrud) Update() error {
	var plan, state VaultSecretResourceModel

	req := s.Request.(*resource.UpdateRequest)
	resp := s.Response.(*resource.UpdateResponse)
	resp.Diagnostics.Append(req.Plan.Get(*s.Context, &plan)...)
	resp.Diagnostics.Append(req.State.Get(*s.Context, &state)...)

	if resp.Diagnostics.HasError() {
		return tfresource.DiagnosticsToError(resp.Diagnostics)
	}

	timeout, diags := plan.Timeouts.Update(*s.Context, UpdateTimeout)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}

	s.OperationTimeout = timeout
	_, cancel := context.WithTimeout(*s.Context, timeout)
	defer cancel()

	if !plan.CompartmentId.IsNull() {
		oldRaw := state.CompartmentId.ValueString()
		newRaw := plan.CompartmentId.ValueString()
		if oldRaw != newRaw && newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(newRaw)
			if err != nil {
				return err
			}
		}
	}
	request := oci_vault.UpdateSecretRequest{}

	/*if !plan.CurrentVersionNumber.IsNull() && !plan.CurrentVersionNumber.Equal(state.CurrentVersionNumber) {
		tmp := plan.CurrentVersionNumber.ValueString()
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert currentVersionNumber string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.CurrentVersionNumber = &tmpInt64
	}*/

	if !plan.DefinedTags.IsNull() && !plan.DefinedTags.IsUnknown() {
		var definedTags map[string]interface{}
		diags = plan.DefinedTags.ElementsAs(*s.Context, &definedTags, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags)
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		tmp := plan.Description.ValueString()
		request.Description = &tmp
	}

	if !plan.FreeFormTags.IsNull() && !plan.FreeFormTags.IsUnknown() {
		var freeformTags map[string]interface{}
		diags = plan.DefinedTags.ElementsAs(*s.Context, &freeformTags, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags)
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		var metadata map[string]interface{}
		diags = plan.DefinedTags.ElementsAs(*s.Context, &metadata, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		request.Metadata = metadata
	}

	if !plan.RotationConfig.IsNull() && !plan.RotationConfig.IsUnknown() && len(plan.RotationConfig.Elements()) > 0 {
		rotationConfigCount := len(plan.RotationConfig.Elements())
		rotationConfigs := make([]RotationConfig, rotationConfigCount)
		diags = plan.RotationConfig.ElementsAs(*s.Context, &rotationConfigs, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		tmp, err := s.fwListToRotationConfig(rotationConfigs[0])
		if err != nil {
			return err
		}
		request.RotationConfig = &tmp
	}

	if !plan.SecretContent.IsNull() && !plan.SecretContent.IsUnknown() && !plan.SecretContent.Equal(state.SecretContent) {
		secretContents := make([]SecretContent, 1)
		diags = plan.SecretContent.ElementsAs(*s.Context, &secretContents, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		if len(secretContents) > 0 {
			tmp, err := s.fwListToSecretContent(secretContents[0])
			if err != nil {
				return err
			}
			request.SecretContent = tmp
		}
	}

	var tmp string
	diags = req.State.GetAttribute(*s.Context, path.Root("id"), &tmp)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return tfresource.DiagnosticsToError(resp.Diagnostics)
	}
	request.SecretId = &tmp

	if !plan.SecretRules.IsNull() && !plan.SecretRules.IsUnknown() && !plan.SecretRules.Equal(state.SecretRules) && len(plan.SecretRules.Elements()) > 0 {
		secretRulesCount := len(plan.SecretRules.Elements())
		secretRuless := make([]SecretRules, secretRulesCount)
		diags := plan.SecretRules.ElementsAs(*s.Context, &secretRuless, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return tfresource.DiagnosticsToError(resp.Diagnostics)
		}
		tmp := make([]oci_vault.SecretRule, secretRulesCount)
		for i := range secretRuless {
			converted, err := s.fwListToSecretRules(secretRuless[i])
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 {
			request.SecretRules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vault")

	response, err := s.Client.UpdateSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Secret
	diags = resp.State.SetAttribute(*s.Context, path.Root("secret_content"), plan.SecretContent)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return tfresource.DiagnosticsToError(resp.Diagnostics)
	}

	return nil
}

func (s *VaultSecretResourceCrud) Delete() error {
	request := oci_vault.ScheduleSecretDeletionRequest{}
	var state VaultSecretResourceModel

	diags := s.ResponseState.Get(*s.Context, &state)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}

	timeout, diags := state.Timeouts.Delete(*s.Context, DeleteTimeout)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}

	s.OperationTimeout = timeout
	_, cancel := context.WithTimeout(*s.Context, timeout)
	defer cancel()

	tmp := state.ID.ValueString()
	request.SecretId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vault")

	_, err := s.Client.ScheduleSecretDeletion(context.Background(), request)
	return err
}

func (s *VaultSecretResourceCrud) SetData() error {
	var state = VaultSecretResourceModel{}
	var mapValue basetypes.MapValue

	var diags = s.ResponseState.Get(*s.Context, &state)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}

	state.ID = types.StringValue(s.ID())
	if s.Res.CompartmentId != nil {
		state.CompartmentId = types.StringValue(*s.Res.CompartmentId)
	}

	if s.Res.CurrentVersionNumber != nil {
		state.CurrentVersionNumber = types.StringValue(strconv.FormatInt(*s.Res.CurrentVersionNumber, 10))
	}

	if s.Res.DefinedTags != nil {
		mapValue, diags = types.MapValueFrom(*s.Context, types.StringType, tfresource.DefinedTagsToMap(s.Res.DefinedTags))
		if diags.HasError() {
			return tfresource.DiagnosticsToError(diags)
		}
		state.DefinedTags = mapValue
	}

	if s.Res.Description != nil {
		//state.Description = tfresource.CaseInsensitiveStringValue{StringValue: types.StringValue(*s.Res.Description)} // types.StringValue(*s.Res.Description)
		state.Description = types.StringValue(*s.Res.Description)
	}

	mapValue, diags = types.MapValueFrom(*s.Context, types.StringType, s.Res.FreeformTags)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}
	state.FreeFormTags = mapValue

	if s.Res.KeyId != nil {
		state.KeyId = types.StringValue(*s.Res.KeyId)
	}

	if s.Res.LastRotationTime != nil {
		state.LastRotationTime = types.StringValue(s.Res.LastRotationTime.String())
	}

	if s.Res.LifecycleDetails != nil {
		state.LifecycleDetails = types.StringValue(*s.Res.LifecycleDetails)
	}

	mapValue, diags = types.MapValueFrom(*s.Context, types.StringType, s.Res.Metadata)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}
	state.Metadata = mapValue

	if s.Res.NextRotationTime != nil {
		state.NextRotationTime = types.StringValue(s.Res.NextRotationTime.String())
	} else {
		state.NextRotationTime = types.StringValue("")
	}

	if s.Res.RotationConfig != nil {
		state.RotationConfig = RotationConfigToListFw(s.Res.RotationConfig)
	} else {
		state.RotationConfig = types.ListNull(types.ObjectType{AttrTypes: rotationConfigMap})
	}

	state.RotationStatus = types.StringValue(string(s.Res.RotationStatus))

	if s.Res.SecretName != nil {
		state.SecretName = types.StringValue(*s.Res.SecretName)
	}

	if len(s.Res.SecretRules) > 0 {
		state.SecretRules = SecretRulesToListFw(s.Res.SecretRules)
	}

	state.State = types.StringValue(string(s.Res.LifecycleState))

	if s.Res.TimeCreated != nil {
		state.TimeCreated = types.StringValue(s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfCurrentVersionExpiry != nil {
		state.TimeOfCurrentVersionExpiry = types.StringValue(s.Res.TimeOfCurrentVersionExpiry.String())
	}

	if s.Res.TimeOfDeletion != nil {
		state.TimeOfDeletion = types.StringValue(s.Res.TimeOfDeletion.String())
	}

	if s.Res.VaultId != nil {
		state.VaultId = types.StringValue(*s.Res.VaultId)
	}

	diags = s.ResponseState.Set(*s.Context, &state)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}

	return nil
}

var (
	rotationConfigMap = map[string]attr.Type{
		"is_scheduled_rotation_enabled": types.BoolType,
		"rotation_interval":             types.StringType,
		"target_system_details":         types.ListType{ElemType: types.ObjectType{AttrTypes: targetSystemDetailsMap}},
	}

	targetSystemDetailsMap = map[string]attr.Type{
		"adb_id":             types.StringType,
		"function_id":        types.StringType,
		"target_system_type": types.StringType,
	}

	secretContentMap = map[string]attr.Type{
		"content":      types.StringType,
		"content_type": types.StringType,
		"name":         types.StringType,
		"stage":        types.StringType,
	}

	secretRulesMap = map[string]attr.Type{
		"is_enforced_on_deleted_secret_versions":        types.BoolType,
		"is_secret_content_retrieval_blocked_on_expiry": types.BoolType,
		"rule_type":                      types.StringType,
		"secret_version_expiry_interval": types.StringType,
		"time_of_absolute_expiry":        types.StringType,
	}
)

func (s *VaultSecretResourceCrud) fwListToRotationConfig(rotationConfig RotationConfig) (oci_vault.RotationConfig, error) {
	result := oci_vault.RotationConfig{}

	if !rotationConfig.IsScheduledRotationEnabled.IsNull() {
		tmp := rotationConfig.IsScheduledRotationEnabled.ValueBool()
		result.IsScheduledRotationEnabled = &tmp
	}

	if !rotationConfig.RotationInterval.IsNull() {
		tmp := rotationConfig.RotationInterval.ValueString()
		result.RotationInterval = &tmp
	}

	if !rotationConfig.TargetSystemDetails.IsNull() && len(rotationConfig.TargetSystemDetails.Elements()) > 0 {
		tmpCount := len(rotationConfig.TargetSystemDetails.Elements())
		tmpList := make([]TargetSystemDetails, tmpCount)
		diags := rotationConfig.TargetSystemDetails.ElementsAs(*s.Context, &tmpList, false)
		if diags.HasError() {
			return result, tfresource.DiagnosticsToError(diags)
		}
		tmp, err := s.fwListToTargetSystemDetails(tmpList[0])
		if err != nil {
			return result, fmt.Errorf("unable to convert target_system_details, encountered error: %v", err)
		}
		result.TargetSystemDetails = tmp
	}

	return result, nil
}

func RotationConfigToListFw(obj *oci_vault.RotationConfig) types.List {
	elemType := types.ObjectType{AttrTypes: rotationConfigMap}

	valuesMap := map[string]attr.Value{}
	if obj.IsScheduledRotationEnabled != nil {
		valuesMap["is_scheduled_rotation_enabled"] = types.BoolValue(*obj.IsScheduledRotationEnabled)
	} else {
		valuesMap["is_scheduled_rotation_enabled"] = types.BoolValue(false)
	}

	if obj.RotationInterval != nil {
		valuesMap["rotation_interval"] = types.StringValue(*obj.RotationInterval)
	} else {
		valuesMap["rotation_interval"] = types.StringValue("")
	}

	if obj.TargetSystemDetails != nil {
		valuesMap["target_system_details"] = TargetSystemDetailsToListFw(&obj.TargetSystemDetails)
	}

	listValue := types.ObjectValueMust(rotationConfigMap, valuesMap)

	return types.ListValueMust(elemType, []attr.Value{listValue})
}

func (s *VaultSecretResourceCrud) fwListToSecretContent(secretContent SecretContent) (oci_vault.SecretContentDetails, error) {
	var baseObject oci_vault.SecretContentDetails
	//discriminator
	var contentType string
	if !secretContent.ContentType.IsNull() {
		contentType = secretContent.ContentType.ValueString()
	}

	switch strings.ToLower(contentType) {
	case strings.ToLower("BASE64"):
		details := oci_vault.Base64SecretContentDetails{}
		if !secretContent.Content.IsNull() && len(secretContent.Content.ValueString()) > 0 {
			tmp := secretContent.Content.ValueString()
			details.Content = &tmp
		} else {
			details.Content = nil
		}
		if !secretContent.Name.IsNull() && len(secretContent.Name.ValueString()) > 0 {
			tmp := secretContent.Name.ValueString()
			details.Name = &tmp
		} else {
			details.Content = nil
		}
		if !secretContent.Stage.IsNull() && len(secretContent.Stage.ValueString()) > 0 {
			details.Stage = oci_vault.SecretContentDetailsStageEnum(secretContent.Stage.ValueString())
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown content_type '%v' was specified", contentType)
	}
	return baseObject, nil
}

func (s *VaultSecretResourceCrud) fwListToSecretRules(secretRules SecretRules) (oci_vault.SecretRule, error) {
	var baseObject oci_vault.SecretRule
	//discriminator
	var ruleType string
	if !secretRules.RuleType.IsNull() {
		ruleType = secretRules.RuleType.ValueString()
	} else {
		ruleType = "" // default value
	}
	switch strings.ToLower(ruleType) {
	case strings.ToLower("SECRET_EXPIRY_RULE"):
		details := oci_vault.SecretExpiryRule{}
		if !secretRules.IsSecretContentRetrievalBlockedOnExpiry.IsNull() {
			tmp := secretRules.IsSecretContentRetrievalBlockedOnExpiry.ValueBool()
			details.IsSecretContentRetrievalBlockedOnExpiry = &tmp
		}
		if !secretRules.SecretVersionExpiryInterval.IsNull() {
			tmp := secretRules.SecretVersionExpiryInterval.ValueString()
			details.SecretVersionExpiryInterval = &tmp
		}
		if !secretRules.TimeOfAbsoluteExpiry.IsNull() {
			tmp, err := time.Parse(time.RFC3339, secretRules.TimeOfAbsoluteExpiry.ValueString())
			if err != nil {
				return details, err
			}
			details.TimeOfAbsoluteExpiry = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	case strings.ToLower("SECRET_REUSE_RULE"):
		details := oci_vault.SecretReuseRule{}
		if !secretRules.IsEnforcedOnDeletedSecretVersions.IsNull() {
			tmp := secretRules.IsEnforcedOnDeletedSecretVersions.ValueBool()
			details.IsEnforcedOnDeletedSecretVersions = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown rule_type '%v' was specified", ruleType)
	}
	return baseObject, nil
}

func SecretRulesToListFw(secretRules []oci_vault.SecretRule) types.List {
	elemType := types.ObjectType{AttrTypes: secretRulesMap}
	var listValues []attr.Value

	for _, item := range secretRules {
		values := map[string]attr.Value{}
		values["rule_type"] = types.StringValue("")
		values["is_secret_content_retrieval_blocked_on_expiry"] = types.BoolValue(false)
		values["secret_version_expiry_interval"] = types.StringValue("")
		values["time_of_absolute_expiry"] = types.StringValue("")
		values["is_enforced_on_deleted_secret_versions"] = types.BoolValue(false)
		switch v := (item).(type) {
		case oci_vault.SecretExpiryRule:
			values["rule_type"] = types.StringValue("SECRET_EXPIRY_RULE")
			if v.IsSecretContentRetrievalBlockedOnExpiry != nil {
				values["is_secret_content_retrieval_blocked_on_expiry"] = types.BoolValue(*v.IsSecretContentRetrievalBlockedOnExpiry)
			}

			if v.SecretVersionExpiryInterval != nil {
				values["secret_version_expiry_interval"] = types.StringValue(*v.SecretVersionExpiryInterval)
			}

			if v.TimeOfAbsoluteExpiry != nil {
				values["time_of_absolute_expiry"] = types.StringValue(v.TimeOfAbsoluteExpiry.Format(time.RFC3339Nano))
			}
		case oci_vault.SecretReuseRule:
			values["rule_type"] = types.StringValue("SECRET_REUSE_RULE")

			if v.IsEnforcedOnDeletedSecretVersions != nil {
				values["is_enforced_on_deleted_secret_versions"] = types.BoolValue(*v.IsEnforcedOnDeletedSecretVersions)
			}
		default:
			log.Printf("[WARN] Received 'rule_type' of unknown type %v", item)
		}
		listValue := types.ObjectValueMust(secretRulesMap, values)
		listValues = append(listValues, listValue)
	}
	return types.ListValueMust(elemType, listValues)
}

func (s *VaultSecretResourceCrud) fwListToTargetSystemDetails(targetSystemdetails TargetSystemDetails) (oci_vault.TargetSystemDetails, error) {
	var baseObject oci_vault.TargetSystemDetails
	//discriminator
	var targetSystemType string
	if !targetSystemdetails.TargetSystemType.IsNull() {
		targetSystemType = targetSystemdetails.TargetSystemType.ValueString()
	} else {
		targetSystemType = "" // default value
	}
	switch strings.ToLower(targetSystemType) {
	case strings.ToLower("ADB"):
		details := oci_vault.AdbTargetSystemDetails{}
		if !targetSystemdetails.AdbId.IsNull() {
			tmp := targetSystemdetails.AdbId.ValueString()
			details.AdbId = &tmp
		}
		baseObject = details
	case strings.ToLower("FUNCTION"):
		details := oci_vault.FunctionTargetSystemDetails{}
		if !targetSystemdetails.FunctionId.IsNull() {
			tmp := targetSystemdetails.FunctionId.ValueString()
			details.FunctionId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown target_system_type '%v' was specified", targetSystemType)
	}
	return baseObject, nil
}

func TargetSystemDetailsToListFw(obj *oci_vault.TargetSystemDetails) types.List {
	elemType := types.ObjectType{AttrTypes: targetSystemDetailsMap}

	valuesMap := map[string]attr.Value{}
	switch v := (*obj).(type) {
	case oci_vault.AdbTargetSystemDetails:
		valuesMap["target_system_type"] = types.StringValue("ADB")

		if v.AdbId != nil {
			valuesMap["adb_id"] = types.StringValue(*v.AdbId)
		}
		valuesMap["function_id"] = types.StringValue("")
	case oci_vault.FunctionTargetSystemDetails:
		valuesMap["target_system_type"] = types.StringValue("FUNCTION")

		if v.FunctionId != nil {
			valuesMap["function_id"] = types.StringValue(*v.FunctionId)
		}
		valuesMap["adb_id"] = types.StringValue("")
	default:
		log.Printf("[WARN] Received 'target_system_type' of unknown type %v", *obj)
		return types.ListNull(elemType)
	}

	listValue := types.ObjectValueMust(targetSystemDetailsMap, valuesMap)

	return types.ListValueMust(elemType, []attr.Value{listValue})
}

func (s *VaultSecretResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_vault.ChangeSecretCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	var state VaultSecretResourceModel

	diags := s.RequestState.Get(*s.Context, &state)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}

	idTmp := state.ID.ValueString()
	changeCompartmentRequest.SecretId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vault")

	_, err := s.Client.ChangeSecretCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateFw(s); waitErr != nil {
		return waitErr
	}

	return nil
}
