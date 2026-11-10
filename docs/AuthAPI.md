# \AuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiAuth2fa**](AuthAPI.md#GetApiAuth2fa) | **Get** /api/auth/2fa | Two-factor status of the signed-in user
[**GetApiAuthInviteByToken**](AuthAPI.md#GetApiAuthInviteByToken) | **Get** /api/auth/invite/{token} | Preview invitation metadata
[**GetApiAuthMe**](AuthAPI.md#GetApiAuthMe) | **Get** /api/auth/me | Current dashboard user
[**PostApiAuth2faDisable**](AuthAPI.md#PostApiAuth2faDisable) | **Post** /api/auth/2fa/disable | Disable two-factor authentication
[**PostApiAuth2faEnable**](AuthAPI.md#PostApiAuth2faEnable) | **Post** /api/auth/2fa/enable | Confirm enrolment and enable two-factor authentication
[**PostApiAuth2faEnrollConfirm**](AuthAPI.md#PostApiAuth2faEnrollConfirm) | **Post** /api/auth/2fa/enroll/confirm | Confirm the required enrolment and finish signing in
[**PostApiAuth2faEnrollStart**](AuthAPI.md#PostApiAuth2faEnrollStart) | **Post** /api/auth/2fa/enroll/start | Start the enrolment required by the workspace (during sign-in)
[**PostApiAuth2faRecoveryCodes**](AuthAPI.md#PostApiAuth2faRecoveryCodes) | **Post** /api/auth/2fa/recovery-codes | Replace the recovery codes
[**PostApiAuth2faSetup**](AuthAPI.md#PostApiAuth2faSetup) | **Post** /api/auth/2fa/setup | Start two-factor enrolment
[**PostApiAuth2faVerify**](AuthAPI.md#PostApiAuth2faVerify) | **Post** /api/auth/2fa/verify | Complete sign-in with a two-factor code
[**PostApiAuthAcceptInvite**](AuthAPI.md#PostApiAuthAcceptInvite) | **Post** /api/auth/accept-invite | Accept a workspace invitation
[**PostApiAuthLogin**](AuthAPI.md#PostApiAuthLogin) | **Post** /api/auth/login | Sign in with email and password
[**PostApiAuthLogout**](AuthAPI.md#PostApiAuthLogout) | **Post** /api/auth/logout | Revoke refresh token
[**PostApiAuthRefresh**](AuthAPI.md#PostApiAuthRefresh) | **Post** /api/auth/refresh | Rotate refresh token and issue new access JWT
[**PostApiAuthSignup**](AuthAPI.md#PostApiAuthSignup) | **Post** /api/auth/signup | Create account and workspace
[**PostApiAuthSwitchWorkspace**](AuthAPI.md#PostApiAuthSwitchWorkspace) | **Post** /api/auth/switch-workspace | Switch active workspace



## GetApiAuth2fa

> GetApiAuth2fa(ctx).Execute()

Two-factor status of the signed-in user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.GetApiAuth2fa(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetApiAuth2fa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiAuth2faRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiAuthInviteByToken

> GetApiAuthInviteByToken(ctx, token).Execute()

Preview invitation metadata

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.GetApiAuthInviteByToken(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetApiAuthInviteByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiAuthInviteByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiAuthMe

> GetApiAuthMe(ctx).Execute()

Current dashboard user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.GetApiAuthMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetApiAuthMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiAuthMeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuth2faDisable

> PostApiAuth2faDisable(ctx).PostApiAuth2faDisableRequest(postApiAuth2faDisableRequest).IdempotencyKey(idempotencyKey).Execute()

Disable two-factor authentication



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuth2faDisableRequest := *openapiclient.NewPostApiAuth2faDisableRequest() // PostApiAuth2faDisableRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuth2faDisable(context.Background()).PostApiAuth2faDisableRequest(postApiAuth2faDisableRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuth2faDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuth2faDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuth2faDisableRequest** | [**PostApiAuth2faDisableRequest**](PostApiAuth2faDisableRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuth2faEnable

> PostApiAuth2faEnable(ctx).PostApiAuth2faEnableRequest(postApiAuth2faEnableRequest).IdempotencyKey(idempotencyKey).Execute()

Confirm enrolment and enable two-factor authentication



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuth2faEnableRequest := *openapiclient.NewPostApiAuth2faEnableRequest("Code_example") // PostApiAuth2faEnableRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuth2faEnable(context.Background()).PostApiAuth2faEnableRequest(postApiAuth2faEnableRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuth2faEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuth2faEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuth2faEnableRequest** | [**PostApiAuth2faEnableRequest**](PostApiAuth2faEnableRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuth2faEnrollConfirm

> PostApiAuth2faEnrollConfirm(ctx).PostApiAuth2faEnrollConfirmRequest(postApiAuth2faEnrollConfirmRequest).IdempotencyKey(idempotencyKey).Execute()

Confirm the required enrolment and finish signing in



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuth2faEnrollConfirmRequest := *openapiclient.NewPostApiAuth2faEnrollConfirmRequest("PendingToken_example", "Code_example") // PostApiAuth2faEnrollConfirmRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuth2faEnrollConfirm(context.Background()).PostApiAuth2faEnrollConfirmRequest(postApiAuth2faEnrollConfirmRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuth2faEnrollConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuth2faEnrollConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuth2faEnrollConfirmRequest** | [**PostApiAuth2faEnrollConfirmRequest**](PostApiAuth2faEnrollConfirmRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuth2faEnrollStart

> PostApiAuth2faEnrollStart(ctx).PostApiAuth2faEnrollStartRequest(postApiAuth2faEnrollStartRequest).IdempotencyKey(idempotencyKey).Execute()

Start the enrolment required by the workspace (during sign-in)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuth2faEnrollStartRequest := *openapiclient.NewPostApiAuth2faEnrollStartRequest("PendingToken_example") // PostApiAuth2faEnrollStartRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuth2faEnrollStart(context.Background()).PostApiAuth2faEnrollStartRequest(postApiAuth2faEnrollStartRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuth2faEnrollStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuth2faEnrollStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuth2faEnrollStartRequest** | [**PostApiAuth2faEnrollStartRequest**](PostApiAuth2faEnrollStartRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuth2faRecoveryCodes

> PostApiAuth2faRecoveryCodes(ctx).PostApiAuth2faDisableRequest(postApiAuth2faDisableRequest).IdempotencyKey(idempotencyKey).Execute()

Replace the recovery codes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuth2faDisableRequest := *openapiclient.NewPostApiAuth2faDisableRequest() // PostApiAuth2faDisableRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuth2faRecoveryCodes(context.Background()).PostApiAuth2faDisableRequest(postApiAuth2faDisableRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuth2faRecoveryCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuth2faRecoveryCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuth2faDisableRequest** | [**PostApiAuth2faDisableRequest**](PostApiAuth2faDisableRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuth2faSetup

> PostApiAuth2faSetup(ctx).IdempotencyKey(idempotencyKey).Execute()

Start two-factor enrolment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuth2faSetup(context.Background()).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuth2faSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuth2faSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuth2faVerify

> PostApiAuth2faVerify(ctx).PostApiAuth2faVerifyRequest(postApiAuth2faVerifyRequest).IdempotencyKey(idempotencyKey).Execute()

Complete sign-in with a two-factor code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuth2faVerifyRequest := *openapiclient.NewPostApiAuth2faVerifyRequest("PendingToken_example") // PostApiAuth2faVerifyRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuth2faVerify(context.Background()).PostApiAuth2faVerifyRequest(postApiAuth2faVerifyRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuth2faVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuth2faVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuth2faVerifyRequest** | [**PostApiAuth2faVerifyRequest**](PostApiAuth2faVerifyRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuthAcceptInvite

> PostApiAuthAcceptInvite(ctx).PostApiAuthAcceptInviteRequest(postApiAuthAcceptInviteRequest).IdempotencyKey(idempotencyKey).Execute()

Accept a workspace invitation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuthAcceptInviteRequest := *openapiclient.NewPostApiAuthAcceptInviteRequest("Token_example", "Password_example") // PostApiAuthAcceptInviteRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuthAcceptInvite(context.Background()).PostApiAuthAcceptInviteRequest(postApiAuthAcceptInviteRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuthAcceptInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuthAcceptInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuthAcceptInviteRequest** | [**PostApiAuthAcceptInviteRequest**](PostApiAuthAcceptInviteRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuthLogin

> PostApiAuthLogin(ctx).PostApiAuthLoginRequest(postApiAuthLoginRequest).IdempotencyKey(idempotencyKey).Execute()

Sign in with email and password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuthLoginRequest := *openapiclient.NewPostApiAuthLoginRequest("Email_example", "Password_example") // PostApiAuthLoginRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuthLogin(context.Background()).PostApiAuthLoginRequest(postApiAuthLoginRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuthLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuthLoginRequest** | [**PostApiAuthLoginRequest**](PostApiAuthLoginRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuthLogout

> PostApiAuthLogout(ctx).PostApiAuthLogoutRequest(postApiAuthLogoutRequest).IdempotencyKey(idempotencyKey).Execute()

Revoke refresh token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuthLogoutRequest := *openapiclient.NewPostApiAuthLogoutRequest() // PostApiAuthLogoutRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuthLogout(context.Background()).PostApiAuthLogoutRequest(postApiAuthLogoutRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuthLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuthLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuthLogoutRequest** | [**PostApiAuthLogoutRequest**](PostApiAuthLogoutRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuthRefresh

> PostApiAuthRefresh(ctx).PostApiAuthRefreshRequest(postApiAuthRefreshRequest).IdempotencyKey(idempotencyKey).Execute()

Rotate refresh token and issue new access JWT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuthRefreshRequest := *openapiclient.NewPostApiAuthRefreshRequest("RefreshToken_example") // PostApiAuthRefreshRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuthRefresh(context.Background()).PostApiAuthRefreshRequest(postApiAuthRefreshRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuthRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuthRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuthRefreshRequest** | [**PostApiAuthRefreshRequest**](PostApiAuthRefreshRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuthSignup

> PostApiAuthSignup(ctx).PostApiAuthSignupRequest(postApiAuthSignupRequest).IdempotencyKey(idempotencyKey).Execute()

Create account and workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuthSignupRequest := *openapiclient.NewPostApiAuthSignupRequest("Email_example", "Password_example") // PostApiAuthSignupRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuthSignup(context.Background()).PostApiAuthSignupRequest(postApiAuthSignupRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuthSignup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuthSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuthSignupRequest** | [**PostApiAuthSignupRequest**](PostApiAuthSignupRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiAuthSwitchWorkspace

> PostApiAuthSwitchWorkspace(ctx).PostApiAuthSwitchWorkspaceRequest(postApiAuthSwitchWorkspaceRequest).IdempotencyKey(idempotencyKey).Execute()

Switch active workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiAuthSwitchWorkspaceRequest := *openapiclient.NewPostApiAuthSwitchWorkspaceRequest("TenantId_example") // PostApiAuthSwitchWorkspaceRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.PostApiAuthSwitchWorkspace(context.Background()).PostApiAuthSwitchWorkspaceRequest(postApiAuthSwitchWorkspaceRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiAuthSwitchWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiAuthSwitchWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiAuthSwitchWorkspaceRequest** | [**PostApiAuthSwitchWorkspaceRequest**](PostApiAuthSwitchWorkspaceRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

