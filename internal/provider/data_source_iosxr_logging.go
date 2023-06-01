// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &LoggingDataSource{}
	_ datasource.DataSourceWithConfigure = &LoggingDataSource{}
)

func NewLoggingDataSource() datasource.DataSource {
	return &LoggingDataSource{}
}

type LoggingDataSource struct {
	client *client.Client
}

func (d *LoggingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_logging"
}

func (d *LoggingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Logging configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"ipv4_dscp": schema.StringAttribute{
				MarkdownDescription: "Set IP DSCP (DiffServ CodePoint)",
				Computed:            true,
			},
			"trap": schema.StringAttribute{
				MarkdownDescription: "Set trap logging",
				Computed:            true,
			},
			"events_display_location": schema.BoolAttribute{
				MarkdownDescription: "Include alarm source location in message text",
				Computed:            true,
			},
			"events_level": schema.StringAttribute{
				MarkdownDescription: "Log all events with equal or higher (lower level) severity",
				Computed:            true,
			},
			"console": schema.StringAttribute{
				MarkdownDescription: "Set console logging",
				Computed:            true,
			},
			"monitor": schema.StringAttribute{
				MarkdownDescription: "Set monitor logging",
				Computed:            true,
			},
			"buffered_logging_buffer_size": schema.Int64Attribute{
				MarkdownDescription: "Logging buffer size",
				Computed:            true,
			},
			"buffered_level": schema.StringAttribute{
				MarkdownDescription: "configure this node",
				Computed:            true,
			},
			"facility_level": schema.StringAttribute{
				MarkdownDescription: "configure this node",
				Computed:            true,
			},
			"hostnameprefix": schema.StringAttribute{
				MarkdownDescription: "Hostname prefix to add on msgs to servers",
				Computed:            true,
			},
			"suppress_duplicates": schema.BoolAttribute{
				MarkdownDescription: "Suppress consecutive duplicate messages",
				Computed:            true,
			},
		},
	}
}

func (d *LoggingDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *LoggingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Logging

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
