package security_attribute

import (
	"fmt"

	oci_security_attribute "github.com/oracle/oci-go-sdk/v65/securityattribute"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportSecurityAttributeSecurityAttributeHints.GetIdFn = getSecurityAttributeSecurityAttributeId
	tf_export.RegisterTenancyGraphs("security_attribute", securityAttributeResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getSecurityAttributeSecurityAttributeId(resource *tf_export.OCIResource) (string, error) {

	securityAttributeName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find securityAttributeName for SecurityAttribute SecurityAttribute")
	}
	securityAttributeNamespaceId := resource.Parent.Id
	return GetSecurityAttributeCompositeId(securityAttributeName, securityAttributeNamespaceId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportSecurityAttributeSecurityAttributeNamespaceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_security_attribute_security_attribute_namespace",
	DatasourceClass:      "oci_security_attribute_security_attribute_namespaces",
	DatasourceItemsAttr:  "security_attribute_namespaces",
	ResourceAbbreviation: "security_attribute_namespace",
	DiscoverableLifecycleStates: []string{
		string(oci_security_attribute.SecurityAttributeNamespaceLifecycleStateActive),
		string(oci_security_attribute.SecurityAttributeNamespaceLifecycleStateInactive),
	},
}

var exportSecurityAttributeSecurityAttributeHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_security_attribute_security_attribute",
	DatasourceClass:        "oci_security_attribute_security_attributes",
	DatasourceItemsAttr:    "security_attributes",
	ResourceAbbreviation:   "security_attribute",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_security_attribute.SecurityAttributeLifecycleStateActive),
		string(oci_security_attribute.SecurityAttributeLifecycleStateInactive),
	},
}

var securityAttributeResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportSecurityAttributeSecurityAttributeNamespaceHints},
	},
	"oci_security_attribute_security_attribute_namespace": {
		{
			TerraformResourceHints: exportSecurityAttributeSecurityAttributeHints,
			DatasourceQueryParams: map[string]string{
				"security_attribute_namespace_id": "id",
			},
		},
	},
}
