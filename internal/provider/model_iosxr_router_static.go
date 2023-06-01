// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RouterStatic struct {
	Device                    types.String                            `tfsdk:"device"`
	Id                        types.String                            `tfsdk:"id"`
	PrefixAddress             types.String                            `tfsdk:"prefix_address"`
	PrefixLength              types.Int64                             `tfsdk:"prefix_length"`
	NexthopInterfaces         []RouterStaticNexthopInterfaces         `tfsdk:"nexthop_interfaces"`
	NexthopInterfaceAddresses []RouterStaticNexthopInterfaceAddresses `tfsdk:"nexthop_interface_addresses"`
	NexthopAddresses          []RouterStaticNexthopAddresses          `tfsdk:"nexthop_addresses"`
}
type RouterStaticNexthopInterfaces struct {
	InterfaceName  types.String `tfsdk:"interface_name"`
	Description    types.String `tfsdk:"description"`
	Tag            types.Int64  `tfsdk:"tag"`
	DistanceMetric types.Int64  `tfsdk:"distance_metric"`
}
type RouterStaticNexthopInterfaceAddresses struct {
	InterfaceName  types.String `tfsdk:"interface_name"`
	Address        types.String `tfsdk:"address"`
	Description    types.String `tfsdk:"description"`
	Tag            types.Int64  `tfsdk:"tag"`
	DistanceMetric types.Int64  `tfsdk:"distance_metric"`
}
type RouterStaticNexthopAddresses struct {
	Address        types.String `tfsdk:"address"`
	Description    types.String `tfsdk:"description"`
	Tag            types.Int64  `tfsdk:"tag"`
	DistanceMetric types.Int64  `tfsdk:"distance_metric"`
}

func (data RouterStatic) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-static-cfg:/router/static/address-family/ipv4/unicast/prefixes/prefix[prefix-address=%s][prefix-length=%d]", data.PrefixAddress.ValueString(), data.PrefixLength.ValueInt64())
}

func (data RouterStatic) toBody(ctx context.Context) string {
	body := "{}"
	if !data.PrefixAddress.IsNull() && !data.PrefixAddress.IsUnknown() {
		body, _ = sjson.Set(body, "prefix-address", data.PrefixAddress.ValueString())
	}
	if !data.PrefixLength.IsNull() && !data.PrefixLength.IsUnknown() {
		body, _ = sjson.Set(body, "prefix-length", strconv.FormatInt(data.PrefixLength.ValueInt64(), 10))
	}
	if len(data.NexthopInterfaces) > 0 {
		body, _ = sjson.Set(body, "nexthop-interfaces.nexthop-interface", []interface{}{})
		for index, item := range data.NexthopInterfaces {
			if !item.InterfaceName.IsNull() && !item.InterfaceName.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-interfaces.nexthop-interface"+"."+strconv.Itoa(index)+"."+"interface-name", item.InterfaceName.ValueString())
			}
			if !item.Description.IsNull() && !item.Description.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-interfaces.nexthop-interface"+"."+strconv.Itoa(index)+"."+"description", item.Description.ValueString())
			}
			if !item.Tag.IsNull() && !item.Tag.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-interfaces.nexthop-interface"+"."+strconv.Itoa(index)+"."+"tag", strconv.FormatInt(item.Tag.ValueInt64(), 10))
			}
			if !item.DistanceMetric.IsNull() && !item.DistanceMetric.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-interfaces.nexthop-interface"+"."+strconv.Itoa(index)+"."+"distance-metric", strconv.FormatInt(item.DistanceMetric.ValueInt64(), 10))
			}
		}
	}
	if len(data.NexthopInterfaceAddresses) > 0 {
		body, _ = sjson.Set(body, "nexthop-interface-addresses.nexthop-interface-address", []interface{}{})
		for index, item := range data.NexthopInterfaceAddresses {
			if !item.InterfaceName.IsNull() && !item.InterfaceName.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-interface-addresses.nexthop-interface-address"+"."+strconv.Itoa(index)+"."+"interface-name", item.InterfaceName.ValueString())
			}
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-interface-addresses.nexthop-interface-address"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.Description.IsNull() && !item.Description.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-interface-addresses.nexthop-interface-address"+"."+strconv.Itoa(index)+"."+"description", item.Description.ValueString())
			}
			if !item.Tag.IsNull() && !item.Tag.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-interface-addresses.nexthop-interface-address"+"."+strconv.Itoa(index)+"."+"tag", strconv.FormatInt(item.Tag.ValueInt64(), 10))
			}
			if !item.DistanceMetric.IsNull() && !item.DistanceMetric.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-interface-addresses.nexthop-interface-address"+"."+strconv.Itoa(index)+"."+"distance-metric", strconv.FormatInt(item.DistanceMetric.ValueInt64(), 10))
			}
		}
	}
	if len(data.NexthopAddresses) > 0 {
		body, _ = sjson.Set(body, "nexthop-addresses.nexthop-address", []interface{}{})
		for index, item := range data.NexthopAddresses {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-addresses.nexthop-address"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.Description.IsNull() && !item.Description.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-addresses.nexthop-address"+"."+strconv.Itoa(index)+"."+"description", item.Description.ValueString())
			}
			if !item.Tag.IsNull() && !item.Tag.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-addresses.nexthop-address"+"."+strconv.Itoa(index)+"."+"tag", strconv.FormatInt(item.Tag.ValueInt64(), 10))
			}
			if !item.DistanceMetric.IsNull() && !item.DistanceMetric.IsUnknown() {
				body, _ = sjson.Set(body, "nexthop-addresses.nexthop-address"+"."+strconv.Itoa(index)+"."+"distance-metric", strconv.FormatInt(item.DistanceMetric.ValueInt64(), 10))
			}
		}
	}
	return body
}

func (data *RouterStatic) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.NexthopInterfaces {
		keys := [...]string{"interface-name"}
		keyValues := [...]string{data.NexthopInterfaces[i].InterfaceName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "nexthop-interfaces.nexthop-interface").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("interface-name"); value.Exists() && !data.NexthopInterfaces[i].InterfaceName.IsNull() {
			data.NexthopInterfaces[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.NexthopInterfaces[i].InterfaceName = types.StringNull()
		}
		if value := r.Get("description"); value.Exists() && !data.NexthopInterfaces[i].Description.IsNull() {
			data.NexthopInterfaces[i].Description = types.StringValue(value.String())
		} else {
			data.NexthopInterfaces[i].Description = types.StringNull()
		}
		if value := r.Get("tag"); value.Exists() && !data.NexthopInterfaces[i].Tag.IsNull() {
			data.NexthopInterfaces[i].Tag = types.Int64Value(value.Int())
		} else {
			data.NexthopInterfaces[i].Tag = types.Int64Null()
		}
		if value := r.Get("distance-metric"); value.Exists() && !data.NexthopInterfaces[i].DistanceMetric.IsNull() {
			data.NexthopInterfaces[i].DistanceMetric = types.Int64Value(value.Int())
		} else {
			data.NexthopInterfaces[i].DistanceMetric = types.Int64Null()
		}
	}
	for i := range data.NexthopInterfaceAddresses {
		keys := [...]string{"interface-name", "address"}
		keyValues := [...]string{data.NexthopInterfaceAddresses[i].InterfaceName.ValueString(), data.NexthopInterfaceAddresses[i].Address.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "nexthop-interface-addresses.nexthop-interface-address").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("interface-name"); value.Exists() && !data.NexthopInterfaceAddresses[i].InterfaceName.IsNull() {
			data.NexthopInterfaceAddresses[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.NexthopInterfaceAddresses[i].InterfaceName = types.StringNull()
		}
		if value := r.Get("address"); value.Exists() && !data.NexthopInterfaceAddresses[i].Address.IsNull() {
			data.NexthopInterfaceAddresses[i].Address = types.StringValue(value.String())
		} else {
			data.NexthopInterfaceAddresses[i].Address = types.StringNull()
		}
		if value := r.Get("description"); value.Exists() && !data.NexthopInterfaceAddresses[i].Description.IsNull() {
			data.NexthopInterfaceAddresses[i].Description = types.StringValue(value.String())
		} else {
			data.NexthopInterfaceAddresses[i].Description = types.StringNull()
		}
		if value := r.Get("tag"); value.Exists() && !data.NexthopInterfaceAddresses[i].Tag.IsNull() {
			data.NexthopInterfaceAddresses[i].Tag = types.Int64Value(value.Int())
		} else {
			data.NexthopInterfaceAddresses[i].Tag = types.Int64Null()
		}
		if value := r.Get("distance-metric"); value.Exists() && !data.NexthopInterfaceAddresses[i].DistanceMetric.IsNull() {
			data.NexthopInterfaceAddresses[i].DistanceMetric = types.Int64Value(value.Int())
		} else {
			data.NexthopInterfaceAddresses[i].DistanceMetric = types.Int64Null()
		}
	}
	for i := range data.NexthopAddresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.NexthopAddresses[i].Address.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "nexthop-addresses.nexthop-address").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("address"); value.Exists() && !data.NexthopAddresses[i].Address.IsNull() {
			data.NexthopAddresses[i].Address = types.StringValue(value.String())
		} else {
			data.NexthopAddresses[i].Address = types.StringNull()
		}
		if value := r.Get("description"); value.Exists() && !data.NexthopAddresses[i].Description.IsNull() {
			data.NexthopAddresses[i].Description = types.StringValue(value.String())
		} else {
			data.NexthopAddresses[i].Description = types.StringNull()
		}
		if value := r.Get("tag"); value.Exists() && !data.NexthopAddresses[i].Tag.IsNull() {
			data.NexthopAddresses[i].Tag = types.Int64Value(value.Int())
		} else {
			data.NexthopAddresses[i].Tag = types.Int64Null()
		}
		if value := r.Get("distance-metric"); value.Exists() && !data.NexthopAddresses[i].DistanceMetric.IsNull() {
			data.NexthopAddresses[i].DistanceMetric = types.Int64Value(value.Int())
		} else {
			data.NexthopAddresses[i].DistanceMetric = types.Int64Null()
		}
	}
}

func (data *RouterStatic) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "nexthop-interfaces.nexthop-interface"); value.Exists() {
		data.NexthopInterfaces = make([]RouterStaticNexthopInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterStaticNexthopInterfaces{}
			if cValue := v.Get("interface-name"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			}
			if cValue := v.Get("tag"); cValue.Exists() {
				item.Tag = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("distance-metric"); cValue.Exists() {
				item.DistanceMetric = types.Int64Value(cValue.Int())
			}
			data.NexthopInterfaces = append(data.NexthopInterfaces, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "nexthop-interface-addresses.nexthop-interface-address"); value.Exists() {
		data.NexthopInterfaceAddresses = make([]RouterStaticNexthopInterfaceAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterStaticNexthopInterfaceAddresses{}
			if cValue := v.Get("interface-name"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			}
			if cValue := v.Get("tag"); cValue.Exists() {
				item.Tag = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("distance-metric"); cValue.Exists() {
				item.DistanceMetric = types.Int64Value(cValue.Int())
			}
			data.NexthopInterfaceAddresses = append(data.NexthopInterfaceAddresses, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "nexthop-addresses.nexthop-address"); value.Exists() {
		data.NexthopAddresses = make([]RouterStaticNexthopAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterStaticNexthopAddresses{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			}
			if cValue := v.Get("tag"); cValue.Exists() {
				item.Tag = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("distance-metric"); cValue.Exists() {
				item.DistanceMetric = types.Int64Value(cValue.Int())
			}
			data.NexthopAddresses = append(data.NexthopAddresses, item)
			return true
		})
	}
}

func (data *RouterStatic) fromPlan(ctx context.Context, plan RouterStatic) {
	data.Device = plan.Device
	data.PrefixAddress = types.StringValue(plan.PrefixAddress.ValueString())
	data.PrefixLength = types.Int64Value(plan.PrefixLength.ValueInt64())
}

func (data *RouterStatic) getDeletedListItems(ctx context.Context, state RouterStatic) []string {
	deletedListItems := make([]string, 0)
	for i := range state.NexthopInterfaces {
		keys := [...]string{"interface-name"}
		stateKeyValues := [...]string{state.NexthopInterfaces[i].InterfaceName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.NexthopInterfaces[i].InterfaceName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.NexthopInterfaces {
			found = true
			if state.NexthopInterfaces[i].InterfaceName.ValueString() != data.NexthopInterfaces[j].InterfaceName.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/nexthop-interfaces/nexthop-interface%v", state.getPath(), keyString))
		}
	}
	for i := range state.NexthopInterfaceAddresses {
		keys := [...]string{"interface-name", "address"}
		stateKeyValues := [...]string{state.NexthopInterfaceAddresses[i].InterfaceName.ValueString(), state.NexthopInterfaceAddresses[i].Address.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.NexthopInterfaceAddresses[i].InterfaceName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.NexthopInterfaceAddresses[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.NexthopInterfaceAddresses {
			found = true
			if state.NexthopInterfaceAddresses[i].InterfaceName.ValueString() != data.NexthopInterfaceAddresses[j].InterfaceName.ValueString() {
				found = false
			}
			if state.NexthopInterfaceAddresses[i].Address.ValueString() != data.NexthopInterfaceAddresses[j].Address.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/nexthop-interface-addresses/nexthop-interface-address%v", state.getPath(), keyString))
		}
	}
	for i := range state.NexthopAddresses {
		keys := [...]string{"address"}
		stateKeyValues := [...]string{state.NexthopAddresses[i].Address.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.NexthopAddresses[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.NexthopAddresses {
			found = true
			if state.NexthopAddresses[i].Address.ValueString() != data.NexthopAddresses[j].Address.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/nexthop-addresses/nexthop-address%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *RouterStatic) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}
