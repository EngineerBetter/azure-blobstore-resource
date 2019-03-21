package operationalinsights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// QueryClient is the log Analytics Data Plane Client
type QueryClient struct {
	BaseClient
}

// NewQueryClient creates an instance of the QueryClient client.
func NewQueryClient() QueryClient {
	return NewQueryClientWithBaseURI(DefaultBaseURI)
}

// NewQueryClientWithBaseURI creates an instance of the QueryClient client.
func NewQueryClientWithBaseURI(baseURI string) QueryClient {
	return QueryClient{NewWithBaseURI(baseURI)}
}

// Execute executes an Analytics query for data. [Here](https://dev.loganalytics.io/documentation/Using-the-API) is an
// example for using POST with an Analytics query.
// Parameters:
// workspaceID - ID of the workspace. This is Workspace ID from the Properties blade in the Azure portal.
// body - the Analytics query. Learn more about the [Analytics query
// syntax](https://azure.microsoft.com/documentation/articles/app-insights-analytics-reference/)
func (client QueryClient) Execute(ctx context.Context, workspaceID string, body QueryBody) (result QueryResults, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryClient.Execute")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Query", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.QueryClient", "Execute", err.Error())
	}

	req, err := client.ExecutePreparer(ctx, workspaceID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueryClient", "Execute", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecuteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.QueryClient", "Execute", resp, "Failure sending request")
		return
	}

	result, err = client.ExecuteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.QueryClient", "Execute", resp, "Failure responding to request")
	}

	return
}

// ExecutePreparer prepares the Execute request.
func (client QueryClient) ExecutePreparer(ctx context.Context, workspaceID string, body QueryBody) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"workspaceId": autorest.Encode("path", workspaceID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/workspaces/{workspaceId}/query", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteSender sends the Execute request. The method will close the
// http.Response Body if it receives an error.
func (client QueryClient) ExecuteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ExecuteResponder handles the response to the Execute request. The method always
// closes the http.Response Body.
func (client QueryClient) ExecuteResponder(resp *http.Response) (result QueryResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
