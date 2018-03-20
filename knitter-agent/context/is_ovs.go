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

package context

import (
	"github.com/HyperNetworks/Knitter/knitter-agent/domain/const-value"
	"github.com/HyperNetworks/Knitter/pkg/klog"
	"github.com/HyperNetworks/Knitter/pkg/trans-dsl"
)

type IsOvs struct {
}

func (this *IsOvs) Ok(transInfo *transdsl.TransInfo) bool {
	portObj := transInfo.AppInfo.(*KnitterInfo).podObj.PortObjs[transInfo.RepeatIdx]
	if portObj.EagerAttr.VnicType == constvalue.MechDriverOvs {
		klog.Infof("***IsOvs: true***")
		return true
	}
	klog.Infof("***IsOvs: false***")
	return false
}