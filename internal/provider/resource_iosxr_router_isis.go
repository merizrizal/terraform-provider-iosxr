// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/helpers"
)

type resourceRouterISISType struct{}

func (t resourceRouterISISType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router ISIS configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			"process_id": {
				MarkdownDescription: helpers.NewAttributeDescription("Process ID").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"is_type": {
				MarkdownDescription: helpers.NewAttributeDescription("Area type (level)").AddStringEnumDescription("level-1", "level-1-2", "level-2-only").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("level-1", "level-1-2", "level-2-only"),
				},
			},
			"nets": {
				MarkdownDescription: helpers.NewAttributeDescription("A Network Entity Title (NET) for this process").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"net_id": {
						MarkdownDescription: helpers.NewAttributeDescription("A Network Entity Title (NET) for this process").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
			"address_families": {
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS address family").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"af_name": {
						MarkdownDescription: helpers.NewAttributeDescription("Address family name").AddStringEnumDescription("ipv4", "ipv6").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringEnumValidator("ipv4", "ipv6"),
						},
					},
					"saf_name": {
						MarkdownDescription: helpers.NewAttributeDescription("Sub address family name").AddStringEnumDescription("multicast", "unicast").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringEnumValidator("multicast", "unicast"),
						},
					},
					"mpls_ldp_auto_config": {
						MarkdownDescription: helpers.NewAttributeDescription("Enable LDP IGP interface auto-configuration").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"metric_style_narrow": {
						MarkdownDescription: helpers.NewAttributeDescription("Use old style of TLVs with narrow metric").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"metric_style_wide": {
						MarkdownDescription: helpers.NewAttributeDescription("Use new style of TLVs to carry wider metric").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"metric_style_transition": {
						MarkdownDescription: helpers.NewAttributeDescription("Send and accept both styles of TLVs during transition").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"router_id_interface_name": {
						MarkdownDescription: helpers.NewAttributeDescription("Router ID Interface").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringPatternValidator(0, 0, `[a-zA-Z0-9.:_/-]+`),
						},
					},
					"router_id_ip_address": {
						MarkdownDescription: helpers.NewAttributeDescription("Router ID address").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
					"default_information_originate": {
						MarkdownDescription: helpers.NewAttributeDescription("Distribute a default route").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
			"interfaces": {
				MarkdownDescription: helpers.NewAttributeDescription("Enter the IS-IS interface configuration submode").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"interface_name": {
						MarkdownDescription: helpers.NewAttributeDescription("Enter the IS-IS interface configuration submode").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringPatternValidator(0, 0, `[a-zA-Z0-9.:_/-]+`),
						},
					},
					"circuit_type": {
						MarkdownDescription: helpers.NewAttributeDescription("Configure circuit type for interface").AddStringEnumDescription("level-1", "level-1-2", "level-2-only").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringEnumValidator("level-1", "level-1-2", "level-2-only"),
						},
					},
					"hello_padding_disable": {
						MarkdownDescription: helpers.NewAttributeDescription("Disable hello-padding").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"hello_padding_sometimes": {
						MarkdownDescription: helpers.NewAttributeDescription("Enable hello-padding during adjacency formation only").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"priority": {
						MarkdownDescription: helpers.NewAttributeDescription("Set priority for Designated Router election").AddIntegerRangeDescription(0, 127).String,
						Type:                types.Int64Type,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.IntegerRangeValidator(0, 127),
						},
					},
					"point_to_point": {
						MarkdownDescription: helpers.NewAttributeDescription("Treat active LAN interface as point-to-point").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"passive": {
						MarkdownDescription: helpers.NewAttributeDescription("Do not establish adjacencies over this interface").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"suppressed": {
						MarkdownDescription: helpers.NewAttributeDescription("Do not advertise connected prefixes of this interface").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
					"shutdown": {
						MarkdownDescription: helpers.NewAttributeDescription("Shutdown IS-IS on this interface").String,
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t resourceRouterISISType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceRouterISIS{
		provider: provider,
	}, diags
}

type resourceRouterISIS struct {
	provider provider
}

func (r resourceRouterISIS) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan RouterISIS

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody()

	_, diags = r.provider.client.Set(ctx, plan.Device.Value, plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete()
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		_, diags = r.provider.client.Set(ctx, plan.Device.Value, i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	plan.setUnknownValues()

	plan.Id = types.String{Value: plan.getPath()}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourceRouterISIS) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state RouterISIS

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	getResp, diags := r.provider.client.Get(ctx, state.Device.Value, state.Id.Value)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.updateFromBody(getResp.Notification[0].Update[0].Val.GetJsonIetfVal())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceRouterISIS) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state RouterISIS

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.Value))

	// Update object
	body := plan.toBody()

	_, diags = r.provider.client.Set(ctx, plan.Device.Value, plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.setUnknownValues()

	deletedListItems := plan.getDeletedListItems(state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		_, diags = r.provider.client.Set(ctx, plan.Device.Value, i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete()
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		_, diags = r.provider.client.Set(ctx, plan.Device.Value, i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.Value))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourceRouterISIS) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state RouterISIS

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.Value))

	_, diags = r.provider.client.Set(ctx, state.Device.Value, state.getPath(), "", client.Delete)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.Value))

	resp.State.RemoveResource(ctx)
}

func (r resourceRouterISIS) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
