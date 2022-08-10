// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package settings

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	settingspb "google.golang.org/genproto/googleapis/cloud/securitycenter/settings/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newSecurityCenterSettingsClientHook clientHook

// SecurityCenterSettingsCallOptions contains the retry settings for each method of SecurityCenterSettingsClient.
type SecurityCenterSettingsCallOptions struct {
	GetServiceAccount                   []gax.CallOption
	GetSettings                         []gax.CallOption
	UpdateSettings                      []gax.CallOption
	ResetSettings                       []gax.CallOption
	BatchGetSettings                    []gax.CallOption
	CalculateEffectiveSettings          []gax.CallOption
	BatchCalculateEffectiveSettings     []gax.CallOption
	GetComponentSettings                []gax.CallOption
	UpdateComponentSettings             []gax.CallOption
	ResetComponentSettings              []gax.CallOption
	CalculateEffectiveComponentSettings []gax.CallOption
	ListDetectors                       []gax.CallOption
	ListComponents                      []gax.CallOption
}

func defaultSecurityCenterSettingsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("securitycenter.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("securitycenter.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://securitycenter.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSecurityCenterSettingsCallOptions() *SecurityCenterSettingsCallOptions {
	return &SecurityCenterSettingsCallOptions{
		GetServiceAccount: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ResetSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchGetSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CalculateEffectiveSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchCalculateEffectiveSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetComponentSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateComponentSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ResetComponentSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CalculateEffectiveComponentSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListDetectors: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListComponents: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalSecurityCenterSettingsClient is an interface that defines the methods available from Cloud Security Command Center API.
type internalSecurityCenterSettingsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetServiceAccount(context.Context, *settingspb.GetServiceAccountRequest, ...gax.CallOption) (*settingspb.ServiceAccount, error)
	GetSettings(context.Context, *settingspb.GetSettingsRequest, ...gax.CallOption) (*settingspb.Settings, error)
	UpdateSettings(context.Context, *settingspb.UpdateSettingsRequest, ...gax.CallOption) (*settingspb.Settings, error)
	ResetSettings(context.Context, *settingspb.ResetSettingsRequest, ...gax.CallOption) error
	BatchGetSettings(context.Context, *settingspb.BatchGetSettingsRequest, ...gax.CallOption) (*settingspb.BatchGetSettingsResponse, error)
	CalculateEffectiveSettings(context.Context, *settingspb.CalculateEffectiveSettingsRequest, ...gax.CallOption) (*settingspb.Settings, error)
	BatchCalculateEffectiveSettings(context.Context, *settingspb.BatchCalculateEffectiveSettingsRequest, ...gax.CallOption) (*settingspb.BatchCalculateEffectiveSettingsResponse, error)
	GetComponentSettings(context.Context, *settingspb.GetComponentSettingsRequest, ...gax.CallOption) (*settingspb.ComponentSettings, error)
	UpdateComponentSettings(context.Context, *settingspb.UpdateComponentSettingsRequest, ...gax.CallOption) (*settingspb.ComponentSettings, error)
	ResetComponentSettings(context.Context, *settingspb.ResetComponentSettingsRequest, ...gax.CallOption) error
	CalculateEffectiveComponentSettings(context.Context, *settingspb.CalculateEffectiveComponentSettingsRequest, ...gax.CallOption) (*settingspb.ComponentSettings, error)
	ListDetectors(context.Context, *settingspb.ListDetectorsRequest, ...gax.CallOption) *DetectorIterator
	ListComponents(context.Context, *settingspb.ListComponentsRequest, ...gax.CallOption) *StringIterator
}

// SecurityCenterSettingsClient is a client for interacting with Cloud Security Command Center API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// API OverviewThe SecurityCenterSettingsService is a sub-api of
// securitycenter.googleapis.com. The service provides methods to manage
// Security Center Settings, and Component Settings for GCP organizations,
// folders, projects, and clusters.
type SecurityCenterSettingsClient struct {
	// The internal transport-dependent client.
	internalClient internalSecurityCenterSettingsClient

	// The call options for this service.
	CallOptions *SecurityCenterSettingsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SecurityCenterSettingsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SecurityCenterSettingsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SecurityCenterSettingsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetServiceAccount retrieves the organizations service account, if it exists, otherwise it
// creates the organization service account. This API is idempotent and
// will only create a service account once. On subsequent calls it will
// return the previously created service account.  SHA, SCC and CTD Infra
// Automation will use this SA.  This SA will not have any permissions when
// created.  The UI will provision this via IAM or the user will using
// their own internal process. This API only creates SAs on the organization.
// Folders are not supported and projects will use per-project SAs associated
// with APIs enabled on a project. This API will be called by the UX
// onboarding workflow.
func (c *SecurityCenterSettingsClient) GetServiceAccount(ctx context.Context, req *settingspb.GetServiceAccountRequest, opts ...gax.CallOption) (*settingspb.ServiceAccount, error) {
	return c.internalClient.GetServiceAccount(ctx, req, opts...)
}

// GetSettings gets the Settings.
func (c *SecurityCenterSettingsClient) GetSettings(ctx context.Context, req *settingspb.GetSettingsRequest, opts ...gax.CallOption) (*settingspb.Settings, error) {
	return c.internalClient.GetSettings(ctx, req, opts...)
}

// UpdateSettings updates the Settings.
func (c *SecurityCenterSettingsClient) UpdateSettings(ctx context.Context, req *settingspb.UpdateSettingsRequest, opts ...gax.CallOption) (*settingspb.Settings, error) {
	return c.internalClient.UpdateSettings(ctx, req, opts...)
}

// ResetSettings reset the organization, folder or project’s settings and return
// the settings of just that resource to the default.
//
// Settings are present at the organization, folder, project, and cluster
// levels. Using Reset on a sub-organization level will remove that resource’s
// override and result in the parent’s settings being used (eg: if Reset on a
// cluster, project settings will be used).
//
// Using Reset on organization will remove the override that was set and
// result in default settings being used.
func (c *SecurityCenterSettingsClient) ResetSettings(ctx context.Context, req *settingspb.ResetSettingsRequest, opts ...gax.CallOption) error {
	return c.internalClient.ResetSettings(ctx, req, opts...)
}

// BatchGetSettings gets a list of settings.
func (c *SecurityCenterSettingsClient) BatchGetSettings(ctx context.Context, req *settingspb.BatchGetSettingsRequest, opts ...gax.CallOption) (*settingspb.BatchGetSettingsResponse, error) {
	return c.internalClient.BatchGetSettings(ctx, req, opts...)
}

// CalculateEffectiveSettings calculateEffectiveSettings looks up all of the Security Center
// Settings resources in the GCP resource hierarchy, and calculates the
// effective settings on that resource by applying the following rules:
//
//	Settings provided closer to the target resource take precedence over
//	those further away (e.g. folder will override organization level
//	settings).
//
//	Product defaults can be overridden at org, folder, project, and cluster
//	levels.
//
//	Detectors will be filtered out if they belong to a billing tier the
//	customer
//	has not configured.
func (c *SecurityCenterSettingsClient) CalculateEffectiveSettings(ctx context.Context, req *settingspb.CalculateEffectiveSettingsRequest, opts ...gax.CallOption) (*settingspb.Settings, error) {
	return c.internalClient.CalculateEffectiveSettings(ctx, req, opts...)
}

// BatchCalculateEffectiveSettings gets a list of effective settings.
func (c *SecurityCenterSettingsClient) BatchCalculateEffectiveSettings(ctx context.Context, req *settingspb.BatchCalculateEffectiveSettingsRequest, opts ...gax.CallOption) (*settingspb.BatchCalculateEffectiveSettingsResponse, error) {
	return c.internalClient.BatchCalculateEffectiveSettings(ctx, req, opts...)
}

// GetComponentSettings gets the Component Settings.
func (c *SecurityCenterSettingsClient) GetComponentSettings(ctx context.Context, req *settingspb.GetComponentSettingsRequest, opts ...gax.CallOption) (*settingspb.ComponentSettings, error) {
	return c.internalClient.GetComponentSettings(ctx, req, opts...)
}

// UpdateComponentSettings updates the Component Settings.
func (c *SecurityCenterSettingsClient) UpdateComponentSettings(ctx context.Context, req *settingspb.UpdateComponentSettingsRequest, opts ...gax.CallOption) (*settingspb.ComponentSettings, error) {
	return c.internalClient.UpdateComponentSettings(ctx, req, opts...)
}

// ResetComponentSettings reset the organization, folder or project’s component settings and return
// the settings to the default. Settings are present at the
// organization, folder and project levels. Using Reset for a folder or
// project will remove the override that was set and result in the
// organization-level settings being used.
func (c *SecurityCenterSettingsClient) ResetComponentSettings(ctx context.Context, req *settingspb.ResetComponentSettingsRequest, opts ...gax.CallOption) error {
	return c.internalClient.ResetComponentSettings(ctx, req, opts...)
}

// CalculateEffectiveComponentSettings gets the Effective Component Settings.
func (c *SecurityCenterSettingsClient) CalculateEffectiveComponentSettings(ctx context.Context, req *settingspb.CalculateEffectiveComponentSettingsRequest, opts ...gax.CallOption) (*settingspb.ComponentSettings, error) {
	return c.internalClient.CalculateEffectiveComponentSettings(ctx, req, opts...)
}

// ListDetectors retrieves an unordered list of available detectors.
func (c *SecurityCenterSettingsClient) ListDetectors(ctx context.Context, req *settingspb.ListDetectorsRequest, opts ...gax.CallOption) *DetectorIterator {
	return c.internalClient.ListDetectors(ctx, req, opts...)
}

// ListComponents retrieves an unordered list of available SCC components.
func (c *SecurityCenterSettingsClient) ListComponents(ctx context.Context, req *settingspb.ListComponentsRequest, opts ...gax.CallOption) *StringIterator {
	return c.internalClient.ListComponents(ctx, req, opts...)
}

// securityCenterSettingsGRPCClient is a client for interacting with Cloud Security Command Center API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type securityCenterSettingsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing SecurityCenterSettingsClient
	CallOptions **SecurityCenterSettingsCallOptions

	// The gRPC API client.
	securityCenterSettingsClient settingspb.SecurityCenterSettingsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSecurityCenterSettingsClient creates a new security center settings service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// API OverviewThe SecurityCenterSettingsService is a sub-api of
// securitycenter.googleapis.com. The service provides methods to manage
// Security Center Settings, and Component Settings for GCP organizations,
// folders, projects, and clusters.
func NewSecurityCenterSettingsClient(ctx context.Context, opts ...option.ClientOption) (*SecurityCenterSettingsClient, error) {
	clientOpts := defaultSecurityCenterSettingsGRPCClientOptions()
	if newSecurityCenterSettingsClientHook != nil {
		hookOpts, err := newSecurityCenterSettingsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := SecurityCenterSettingsClient{CallOptions: defaultSecurityCenterSettingsCallOptions()}

	c := &securityCenterSettingsGRPCClient{
		connPool:                     connPool,
		disableDeadlines:             disableDeadlines,
		securityCenterSettingsClient: settingspb.NewSecurityCenterSettingsServiceClient(connPool),
		CallOptions:                  &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *securityCenterSettingsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *securityCenterSettingsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *securityCenterSettingsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *securityCenterSettingsGRPCClient) GetServiceAccount(ctx context.Context, req *settingspb.GetServiceAccountRequest, opts ...gax.CallOption) (*settingspb.ServiceAccount, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetServiceAccount[0:len((*c.CallOptions).GetServiceAccount):len((*c.CallOptions).GetServiceAccount)], opts...)
	var resp *settingspb.ServiceAccount
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.securityCenterSettingsClient.GetServiceAccount(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *securityCenterSettingsGRPCClient) GetSettings(ctx context.Context, req *settingspb.GetSettingsRequest, opts ...gax.CallOption) (*settingspb.Settings, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetSettings[0:len((*c.CallOptions).GetSettings):len((*c.CallOptions).GetSettings)], opts...)
	var resp *settingspb.Settings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.securityCenterSettingsClient.GetSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *securityCenterSettingsGRPCClient) UpdateSettings(ctx context.Context, req *settingspb.UpdateSettingsRequest, opts ...gax.CallOption) (*settingspb.Settings, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "settings.name", url.QueryEscape(req.GetSettings().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateSettings[0:len((*c.CallOptions).UpdateSettings):len((*c.CallOptions).UpdateSettings)], opts...)
	var resp *settingspb.Settings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.securityCenterSettingsClient.UpdateSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *securityCenterSettingsGRPCClient) ResetSettings(ctx context.Context, req *settingspb.ResetSettingsRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ResetSettings[0:len((*c.CallOptions).ResetSettings):len((*c.CallOptions).ResetSettings)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.securityCenterSettingsClient.ResetSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *securityCenterSettingsGRPCClient) BatchGetSettings(ctx context.Context, req *settingspb.BatchGetSettingsRequest, opts ...gax.CallOption) (*settingspb.BatchGetSettingsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchGetSettings[0:len((*c.CallOptions).BatchGetSettings):len((*c.CallOptions).BatchGetSettings)], opts...)
	var resp *settingspb.BatchGetSettingsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.securityCenterSettingsClient.BatchGetSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *securityCenterSettingsGRPCClient) CalculateEffectiveSettings(ctx context.Context, req *settingspb.CalculateEffectiveSettingsRequest, opts ...gax.CallOption) (*settingspb.Settings, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CalculateEffectiveSettings[0:len((*c.CallOptions).CalculateEffectiveSettings):len((*c.CallOptions).CalculateEffectiveSettings)], opts...)
	var resp *settingspb.Settings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.securityCenterSettingsClient.CalculateEffectiveSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *securityCenterSettingsGRPCClient) BatchCalculateEffectiveSettings(ctx context.Context, req *settingspb.BatchCalculateEffectiveSettingsRequest, opts ...gax.CallOption) (*settingspb.BatchCalculateEffectiveSettingsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchCalculateEffectiveSettings[0:len((*c.CallOptions).BatchCalculateEffectiveSettings):len((*c.CallOptions).BatchCalculateEffectiveSettings)], opts...)
	var resp *settingspb.BatchCalculateEffectiveSettingsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.securityCenterSettingsClient.BatchCalculateEffectiveSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *securityCenterSettingsGRPCClient) GetComponentSettings(ctx context.Context, req *settingspb.GetComponentSettingsRequest, opts ...gax.CallOption) (*settingspb.ComponentSettings, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetComponentSettings[0:len((*c.CallOptions).GetComponentSettings):len((*c.CallOptions).GetComponentSettings)], opts...)
	var resp *settingspb.ComponentSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.securityCenterSettingsClient.GetComponentSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *securityCenterSettingsGRPCClient) UpdateComponentSettings(ctx context.Context, req *settingspb.UpdateComponentSettingsRequest, opts ...gax.CallOption) (*settingspb.ComponentSettings, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "component_settings.name", url.QueryEscape(req.GetComponentSettings().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateComponentSettings[0:len((*c.CallOptions).UpdateComponentSettings):len((*c.CallOptions).UpdateComponentSettings)], opts...)
	var resp *settingspb.ComponentSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.securityCenterSettingsClient.UpdateComponentSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *securityCenterSettingsGRPCClient) ResetComponentSettings(ctx context.Context, req *settingspb.ResetComponentSettingsRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ResetComponentSettings[0:len((*c.CallOptions).ResetComponentSettings):len((*c.CallOptions).ResetComponentSettings)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.securityCenterSettingsClient.ResetComponentSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *securityCenterSettingsGRPCClient) CalculateEffectiveComponentSettings(ctx context.Context, req *settingspb.CalculateEffectiveComponentSettingsRequest, opts ...gax.CallOption) (*settingspb.ComponentSettings, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CalculateEffectiveComponentSettings[0:len((*c.CallOptions).CalculateEffectiveComponentSettings):len((*c.CallOptions).CalculateEffectiveComponentSettings)], opts...)
	var resp *settingspb.ComponentSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.securityCenterSettingsClient.CalculateEffectiveComponentSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *securityCenterSettingsGRPCClient) ListDetectors(ctx context.Context, req *settingspb.ListDetectorsRequest, opts ...gax.CallOption) *DetectorIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListDetectors[0:len((*c.CallOptions).ListDetectors):len((*c.CallOptions).ListDetectors)], opts...)
	it := &DetectorIterator{}
	req = proto.Clone(req).(*settingspb.ListDetectorsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*settingspb.Detector, string, error) {
		resp := &settingspb.ListDetectorsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.securityCenterSettingsClient.ListDetectors(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetDetectors(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *securityCenterSettingsGRPCClient) ListComponents(ctx context.Context, req *settingspb.ListComponentsRequest, opts ...gax.CallOption) *StringIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListComponents[0:len((*c.CallOptions).ListComponents):len((*c.CallOptions).ListComponents)], opts...)
	it := &StringIterator{}
	req = proto.Clone(req).(*settingspb.ListComponentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]string, string, error) {
		resp := &settingspb.ListComponentsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.securityCenterSettingsClient.ListComponents(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetComponents(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// DetectorIterator manages a stream of *settingspb.Detector.
type DetectorIterator struct {
	items    []*settingspb.Detector
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*settingspb.Detector, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *DetectorIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *DetectorIterator) Next() (*settingspb.Detector, error) {
	var item *settingspb.Detector
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *DetectorIterator) bufLen() int {
	return len(it.items)
}

func (it *DetectorIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// StringIterator manages a stream of string.
type StringIterator struct {
	items    []string
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []string, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *StringIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *StringIterator) Next() (string, error) {
	var item string
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *StringIterator) bufLen() int {
	return len(it.items)
}

func (it *StringIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
