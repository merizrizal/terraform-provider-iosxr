// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RouterVRRPInterfaceIPv4DataSource{}
	_ datasource.DataSourceWithConfigure = &RouterVRRPInterfaceIPv4DataSource{}
)

func NewRouterVRRPInterfaceIPv4DataSource() datasource.DataSource {
	return &RouterVRRPInterfaceIPv4DataSource{}
}

type RouterVRRPInterfaceIPv4DataSource struct {
	client *client.Client
}

func (d *RouterVRRPInterfaceIPv4DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_vrrp_interface_ipv4"
}

func (d *RouterVRRPInterfaceIPv4DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router VRRP Interface IPv4 configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "VRRP interface configuration subcommands",
				Required:            true,
			},
			"vrrp_id": schema.Int64Attribute{
				MarkdownDescription: "VRRP configuration",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "VRRP version",
				Required:            true,
			},
			"address": schema.StringAttribute{
				MarkdownDescription: "Enable VRRP and specify IP address(es)",
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Set priority level",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Configure VRRP Session name",
				Computed:            true,
			},
			"text_authentication": schema.StringAttribute{
				MarkdownDescription: "Set plain text authentication string",
				Computed:            true,
			},
			"secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "VRRP IPv4 address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "VRRP IPv4 address",
							Computed:            true,
						},
					},
				},
			},
			"timer_advertisement_seconds": schema.Int64Attribute{
				MarkdownDescription: "Advertisement time in seconds",
				Computed:            true,
			},
			"timer_advertisement_milliseconds": schema.Int64Attribute{
				MarkdownDescription: "Configure in milliseconds",
				Computed:            true,
			},
			"timer_force": schema.BoolAttribute{
				MarkdownDescription: "Force the configured values to be used",
				Computed:            true,
			},
			"preempt_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable preemption",
				Computed:            true,
			},
			"preempt_delay": schema.Int64Attribute{
				MarkdownDescription: "Wait before preempting",
				Computed:            true,
			},
			"accept_mode_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable accept mode",
				Computed:            true,
			},
			"track_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Track an interface",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Track an interface",
							Computed:            true,
						},
						"priority_decrement": schema.Int64Attribute{
							MarkdownDescription: "Priority decrement",
							Computed:            true,
						},
					},
				},
			},
			"track_objects": schema.ListNestedAttribute{
				MarkdownDescription: "Object Tracking",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"object_name": schema.StringAttribute{
							MarkdownDescription: "Object to be tracked",
							Computed:            true,
						},
						"priority_decrement": schema.Int64Attribute{
							MarkdownDescription: "Priority decrement",
							Computed:            true,
						},
					},
				},
			},
			"bfd_fast_detect_peer_ipv4": schema.StringAttribute{
				MarkdownDescription: "BFD peer interface IPv4 address",
				Computed:            true,
			},
		},
	}
}

func (d *RouterVRRPInterfaceIPv4DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterVRRPInterfaceIPv4DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterVRRPInterfaceIPv4Data

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, diags := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.fromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
