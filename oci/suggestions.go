// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import "fmt"

func getSuggestionFromError(tfError customError) string {
	switch tfError.ErrorCode {
	case 400:
		return getSuggestionFor400(tfError)
	case 404:
		return getSuggestionFor404(tfError)
	case 409:
		return getSuggestionFor409(tfError)
	case 429:
		return getSuggestionFor429(tfError)
	case 500:
		return getSuggestionFor500(tfError)
	default:
		return getSuggestionForDefault(tfError)
	}
}

func getSuggestionFor400(tfError customError) string {
	if tfError.ErrorCodeName == "InvalidParameter" {
		return fmt.Sprintf("Please change Terraform config with valid parameter: %s", tfError.Message)
	}
	return fmt.Sprintf(tfError.Message)
}

func getSuggestionFor404(tfError customError) string {
	return fmt.Sprintf("The resource had been deleted or service %s need policy to access this resource.", tfError.Service)
}

func getSuggestionFor409(tfError customError) string {
	return fmt.Sprintf("The resource is in conflict state, please retry again or contact service %s", tfError.Service)
}

func getSuggestionFor429(tfError customError) string {
	return fmt.Sprintf("Please retry again or increase the retry timeout following this document: https://registry.terraform.io/providers/hashicorp/oci/latest/docs#retry_duration_seconds")
}

func getSuggestionFor500(tfError customError) string {
	return fmt.Sprintf("The service for this resource encountered an error. Please contact support for help with service %s", tfError.Service)
}

func getSuggestionForDefault(tfError customError) string {
	// return error message for default
	return tfError.Message
}
