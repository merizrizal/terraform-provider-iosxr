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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterISIS(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "process_id", "P1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "is_type", "level-1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "set_overload_bit_levels.0.level_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "set_overload_bit_levels.0.on_startup_advertise_as_overloaded", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "set_overload_bit_levels.0.on_startup_advertise_as_overloaded_time_to_advertise", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "set_overload_bit_levels.0.on_startup_wait_for_bgp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "set_overload_bit_levels.0.advertise_external", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "set_overload_bit_levels.0.advertise_interlevel", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "nsr", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "nsf_cisco", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "nsf_ietf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "nsf_lifetime", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "nsf_interface_timer", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "nsf_interface_expires", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "log_adjacency_changes", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "lsp_gen_interval_maximum_wait", "5000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "lsp_gen_interval_initial_wait", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "lsp_gen_interval_secondary_wait", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "lsp_refresh_interval", "65000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "max_lsp_lifetime", "65535"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "lsp_password_keychain", "ISIS-KEY"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "distribute_link_state_instance_id", "32"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "affinity_maps.0.name", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "affinity_maps.0.bit_position", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "flex_algos.0.algorithm_number", "128"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "flex_algos.0.advertise_definition", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "flex_algos.0.metric_type_delay", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "nets.0.net_id", "49.0001.2222.2222.2222.00"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.circuit_type", "level-1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.hello_padding_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.hello_padding_sometimes", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.point_to_point", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.passive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.suppressed", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.shutdown", "false"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrRouterISISConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrRouterISISConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_router_isis.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=P1]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrRouterISISConfig_minimum() string {
	config := `resource "iosxr_router_isis" "test" {` + "\n"
	config += `	process_id = "P1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterISISConfig_all() string {
	config := `resource "iosxr_router_isis" "test" {` + "\n"
	config += `	process_id = "P1"` + "\n"
	config += `	is_type = "level-1"` + "\n"
	config += `	set_overload_bit_levels = [{` + "\n"
	config += `		level_id = 1` + "\n"
	config += `		on_startup_advertise_as_overloaded = true` + "\n"
	config += `		on_startup_advertise_as_overloaded_time_to_advertise = 10` + "\n"
	config += `		on_startup_wait_for_bgp = false` + "\n"
	config += `		advertise_external = true` + "\n"
	config += `		advertise_interlevel = true` + "\n"
	config += `		}]` + "\n"
	config += `	nsr = true` + "\n"
	config += `	nsf_cisco = true` + "\n"
	config += `	nsf_ietf = false` + "\n"
	config += `	nsf_lifetime = 10` + "\n"
	config += `	nsf_interface_timer = 5` + "\n"
	config += `	nsf_interface_expires = 2` + "\n"
	config += `	log_adjacency_changes = true` + "\n"
	config += `	lsp_gen_interval_maximum_wait = 5000` + "\n"
	config += `	lsp_gen_interval_initial_wait = 50` + "\n"
	config += `	lsp_gen_interval_secondary_wait = 200` + "\n"
	config += `	lsp_refresh_interval = 65000` + "\n"
	config += `	max_lsp_lifetime = 65535` + "\n"
	config += `	lsp_password_keychain = "ISIS-KEY"` + "\n"
	config += `	distribute_link_state_instance_id = 32` + "\n"
	config += `	affinity_maps = [{` + "\n"
	config += `		name = "22"` + "\n"
	config += `		bit_position = 4` + "\n"
	config += `		}]` + "\n"
	config += `	flex_algos = [{` + "\n"
	config += `		algorithm_number = 128` + "\n"
	config += `		advertise_definition = true` + "\n"
	config += `		metric_type_delay = true` + "\n"
	config += `		}]` + "\n"
	config += `	nets = [{` + "\n"
	config += `		net_id = "49.0001.2222.2222.2222.00"` + "\n"
	config += `		}]` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `		circuit_type = "level-1"` + "\n"
	config += `		hello_padding_disable = true` + "\n"
	config += `		hello_padding_sometimes = false` + "\n"
	config += `		priority = 10` + "\n"
	config += `		point_to_point = false` + "\n"
	config += `		passive = false` + "\n"
	config += `		suppressed = false` + "\n"
	config += `		shutdown = false` + "\n"
	config += `		}]` + "\n"
	config += `}` + "\n"
	return config
}
