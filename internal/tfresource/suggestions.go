// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import "fmt"

func getSuggestionFromError(tfError customError) string {
	switch tfError.TypeOfError {
	case ServiceError:
		return getSuggestionForServiceError(tfError)
	case TimeoutError:
		return getSuggestionForTimeoutError()
	case UnexpectedStateError:
		return getSuggestionForUnexpectedState(tfError)
	default:
		return getSuggestionForDefault(tfError)
	}
}

func getSuggestionForServiceError(tfError customError) string {
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
	switch tfError.ErrorCodeName {
	case "InvalidParameter":
		return fmt.Sprintf("Please update the parameter(s) in the Terraform config as per error message %s", tfError.Message)
	case "LimitExceeded":
		return fmt.Sprintf("Request a service limit increase for this resource %s", tfError.Service)
	case "QuotaExceeded":
		return fmt.Sprintf("Contact your administrator to increase limit for your account or compartment for this service: %s", tfError.Service)
	default:
		return getSuggestionForDefault(tfError)
	}
}

func getSuggestionFor404(tfError customError) string {
	return fmt.Sprintf("Either the resource has been deleted or service %s need policy to access this resource. Policy reference: https://docs.oracle.com/en-us/iaas/Content/Identity/Reference/policyreference.htm", tfError.Service)
}

func getSuggestionFor409(tfError customError) string {
	return fmt.Sprintf("The resource is in a conflicted state. Please retry again or contact support for help with service: %s", tfError.Service)
}

func getSuggestionFor429(tfError customError) string {
	return fmt.Sprintf("Please re-apply your Terraform config and/or increase the retry timeout using this document: https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformtroubleshooting.htm#common_issues__automaticretries")
}

func getSuggestionFor500(tfError customError) string {
	return fmt.Sprintf("The service for this resource encountered an error. Please contact support for help with service: %s", tfError.Service)
}

func getSuggestionForDefault(tfError customError) string {
	// return error message for default
	return fmt.Sprintf("Please retry or contact support for help with service: %s", tfError.Service)
}

func getSuggestionForTimeoutError() string {
	return fmt.Sprintf("Try increasing the timeout by referring to this document: https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformtroubleshooting.htm#common_issues__timeoutwhilewaiting")
}

func getSuggestionForUnexpectedState(tfError customError) string {
	return fmt.Sprintf("Please retry or contact support for help with service: %s", tfError.Service)
}
