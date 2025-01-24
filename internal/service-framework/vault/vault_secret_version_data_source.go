// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vault

import (
	"context"
	"fmt"
	"log"
	"strconv"

	oci_vault "github.com/oracle/oci-go-sdk/v65/vault"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ datasource.DataSource = &VaultSecretVersionDataSource{}
var _ datasource.DataSourceWithConfigure = &VaultSecretVersionDataSource{}

func NewVaultSecretVersionDataSource() datasource.DataSource {
	return &VaultSecretVersionDataSource{}
}

type VaultSecretVersionDataSource struct {
	client *client.OracleClients
}

func (d *VaultSecretVersionDataSource) Configure(ctx context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	log.Println("VaultSecretVersionDataSource providerData", request.ProviderData)
	if request.ProviderData != nil {
		d.client = request.ProviderData.(*client.OracleClients)
	}
}

func (d *VaultSecretVersionDataSource) Metadata(ctx context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = "oci_vault_secret_version"
}

func (d *VaultSecretVersionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"secret_id": schema.StringAttribute{
				Required: true,
			},
			"secret_version_number": schema.StringAttribute{
				Required: true,
			},
			// Computed
			"id": schema.StringAttribute{
				Computed: true,
			},
			"content_type": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"stages": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
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
			"version_number": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type VaultSecretVersionDataSourceCrud struct {
	Context  *context.Context
	Request  *datasource.ReadRequest
	Response *datasource.ReadResponse
	Client   *oci_vault.VaultsClient
	Res      *oci_vault.GetSecretVersionResponse
}

func (d *VaultSecretVersionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	sync := &VaultSecretVersionDataSourceCrud{}
	sync.Context = &ctx
	sync.Request = &req
	sync.Response = resp
	sync.Client = d.client.VaultsClient()

	err := tfresource.ReadResource(sync)
	if err != nil {
		resp.Diagnostics.AddError(err.Error(), "")
	}
}

type VaultSecretVersionDataSourceModel struct {
	SecretId                   types.String `tfsdk:"secret_id"`
	SecretVersionNumber        types.String `tfsdk:"secret_version_number"`
	ID                         types.String `tfsdk:"id"`
	ContentType                types.String `tfsdk:"content_type"`
	Name                       types.String `tfsdk:"name"`
	Stages                     types.List   `tfsdk:"stages"`
	TimeCreated                types.String `tfsdk:"time_created"`
	TimeOfCurrentVersionExpiry types.String `tfsdk:"time_of_current_version_expiry"`
	TimeOfDeletion             types.String `tfsdk:"time_of_deletion"`
	VersionNumber              types.String `tfsdk:"version_number"`
}

func (s *VaultSecretVersionDataSourceCrud) VoidState() {
	var state = VaultSecretVersionDataSourceModel{}
	resp := s.Response
	diags := resp.State.Get(*s.Context, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.ID = types.StringValue("")

	diags = resp.State.Set(*s.Context, &state)
	resp.Diagnostics.Append(diags...)
}

func (s *VaultSecretVersionDataSourceCrud) Get() error {
	request := oci_vault.GetSecretVersionRequest{}

	var config = VaultSecretVersionDataSourceModel{}
	req := s.Request
	resp := s.Response
	diags := req.Config.Get(*s.Context, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return tfresource.DiagnosticsToError(resp.Diagnostics)
	}

	if !(config.SecretId.IsNull() || config.SecretId.IsUnknown()) {
		tmp := config.SecretId.ValueString()
		request.SecretId = &tmp
	}

	if !(config.SecretVersionNumber.IsNull() || config.SecretVersionNumber.IsUnknown()) {
		tmp := config.SecretVersionNumber.ValueString()
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert secretVersionNumber string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SecretVersionNumber = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "vault")

	response, err := s.Client.GetSecretVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *VaultSecretVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	var state = VaultSecretVersionDataSourceModel{}
	var listValue basetypes.ListValue

	resp := s.Response
	diags := resp.State.Get(*s.Context, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return tfresource.DiagnosticsToError(resp.Diagnostics)
	}

	state.ID = types.StringValue(tfresource.GenerateFrameworkDataSourceHashID("VaultSecretVersionDataSource-", *s.Context, resp.State))

	state.ContentType = types.StringValue(string(s.Res.ContentType))

	if s.Res.Name != nil {
		state.Name = types.StringValue(*s.Res.Name)
	}

	/*stagesArray := make([]string, 0)
	for _, stage := range s.Res.Stages {
		stagesArray = append(stagesArray, string(stage))
	}*/
	listValue, diags = types.ListValueFrom(*s.Context, types.StringType, s.Res.Stages)
	if diags.HasError() {
		return tfresource.DiagnosticsToError(diags)
	}
	state.Stages = listValue

	if s.Res.TimeCreated != nil {
		state.TimeCreated = types.StringValue(s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfCurrentVersionExpiry != nil {
		state.TimeOfCurrentVersionExpiry = types.StringValue(s.Res.TimeOfCurrentVersionExpiry.String())
	}

	if s.Res.TimeOfDeletion != nil {
		state.TimeOfDeletion = types.StringValue(s.Res.TimeOfDeletion.String())
	}

	if s.Res.VersionNumber != nil {
		state.VersionNumber = types.StringValue(strconv.FormatInt(*s.Res.VersionNumber, 10))
	}

	diags = resp.State.Set(*s.Context, &state)
	resp.Diagnostics.Append(diags...)

	return nil
}
