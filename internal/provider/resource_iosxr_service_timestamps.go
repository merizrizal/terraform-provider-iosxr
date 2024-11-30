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
	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NewServiceTimestampsResource() resource.Resource {
	return &ServiceTimestampsResource{}
}

type ServiceTimestampsResource struct {
	client *client.Client
}

func (r *ServiceTimestampsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_timestamps"
}

func (r *ServiceTimestampsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Service Timestamps configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"debug_datetime_localtime_only": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp with date and time").String,
				Optional:            true,
			},
			"debug_datetime_localtime": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use local time zone for timestamps").String,
				Optional:            true,
			},
			"debug_datetime_msec": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Include milliseconds in timestamp").String,
				Optional:            true,
			},
			"debug_datetime_show_timezone": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Add time zone information to timestamp").String,
				Optional:            true,
			},
			"debug_datetime_year": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Add year to timestamp").String,
				Optional:            true,
			},
			"debug_uptime": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp with systime uptime").String,
				Optional:            true,
			},
			"debug_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable timestamp debug messages").String,
				Optional:            true,
			},
			"log_datetime_localtime_only": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp with date and time").String,
				Optional:            true,
			},
			"log_datetime_localtime": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use local time zone for timestamps").String,
				Optional:            true,
			},
			"log_datetime_msec": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Include milliseconds in timestamp").String,
				Optional:            true,
			},
			"log_datetime_show_timezone": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Add time zone information to timestamp").String,
				Optional:            true,
			},
			"log_datetime_year": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Add Year to timestamp").String,
				Optional:            true,
			},
			"log_uptime": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timestamp with systime uptime").String,
				Optional:            true,
			},
			"log_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable timestamp log messages").String,
				Optional:            true,
			},
		},
	}
}

func (r *ServiceTimestampsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *ServiceTimestampsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceTimestamps

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ServiceTimestampsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceTimestamps

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	import_ := false
	if state.Id.ValueString() == "" {
		import_ = true
		state.Id = types.StringValue(state.getPath())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	getResp, diags := r.client.Get(ctx, state.Device.ValueString(), state.Id.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	respBody := getResp.Notification[0].Update[0].Val.GetJsonIetfVal()
	if import_ {
		state.fromBody(ctx, respBody)
	} else {
		state.updateFromBody(ctx, respBody)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *ServiceTimestampsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ServiceTimestamps

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Update object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	deletedListItems := plan.getDeletedItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("Removed items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ServiceTimestampsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceTimestamps

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	var ops []client.SetOperation
	deleteMode := "all"

	if deleteMode == "all" {
		ops = append(ops, client.SetOperation{Path: state.Id.ValueString(), Body: "", Operation: client.Delete})
	} else {
		deletePaths := state.getDeletePaths(ctx)
		tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

		for _, i := range deletePaths {
			ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
		}
	}

	_, diags = r.client.Set(ctx, state.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *ServiceTimestampsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), "")...)
}
