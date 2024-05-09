package email

import (
	oci_email "github.com/oracle/oci-go-sdk/v65/email"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("email", emailResourceGraph)
	tf_export.RegisterTenancyGraphs("email_tenancy", emailTenancyResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportEmailSuppressionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_email_suppression",
	DatasourceClass:      "oci_email_suppressions",
	DatasourceItemsAttr:  "suppressions",
	ResourceAbbreviation: "suppression",
}

var exportEmailSenderHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_email_sender",
	DatasourceClass:      "oci_email_senders",
	DatasourceItemsAttr:  "senders",
	ResourceAbbreviation: "sender",
	DiscoverableLifecycleStates: []string{
		string(oci_email.SenderLifecycleStateActive),
		string(oci_email.SenderLifecycleStateNeedsAttention),
	},
}

var exportEmailEmailDomainHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_email_email_domain",
	DatasourceClass:        "oci_email_email_domains",
	DatasourceItemsAttr:    "email_domain_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "email_domain",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_email.EmailDomainLifecycleStateActive),
	},
}

var exportEmailDkimHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_email_dkim",
	DatasourceClass:        "oci_email_dkims",
	DatasourceItemsAttr:    "dkim_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "dkim",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_email.DkimLifecycleStateActive),
		string(oci_email.DkimLifecycleStateNeedsAttention),
	},
}

var exportEmailEmailReturnPathHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_email_email_return_path",
	DatasourceClass:        "oci_email_email_return_paths",
	DatasourceItemsAttr:    "email_return_path_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "email_return_path",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_email.EmailReturnPathLifecycleStateActive),
		string(oci_email.EmailReturnPathLifecycleStateNeedsAttention),
	},
}

var emailResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportEmailSenderHints},
		{TerraformResourceHints: exportEmailEmailDomainHints},
		{TerraformResourceHints: exportEmailEmailReturnPathHints},
	},
	"oci_email_email_domain": {
		{
			TerraformResourceHints: exportEmailDkimHints,
			DatasourceQueryParams: map[string]string{
				"email_domain_id": "id",
			},
		},
	},
}

var emailTenancyResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportEmailSuppressionHints},
	},
}
