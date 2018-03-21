/*
Copyright 2018 ZTE Corporation. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package implement

import (
	"github.com/HyperNetworks/Knitter/knitter-agent/infra/util/base"
	"github.com/vishvananda/netns"
)

type NameSpace struct {
	base.NameSpace
}

func (self *NameSpace) Set(ns netns.NsHandle) (err error) {
	return netns.Set(ns)
}

func (self *NameSpace) Get() (netns.NsHandle, error) {
	nsHandle, err := netns.Get()
	return nsHandle, err
}
