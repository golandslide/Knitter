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
	"github.com/HyperNetworks/Knitter/knitter-agent/domain/object/os-obj"
	"github.com/HyperNetworks/Knitter/pkg/klog"
	"github.com/HyperNetworks/Knitter/pkg/trans-dsl"
)

type DestroyVethPairAction struct {
}

func (this *DestroyVethPairAction) Exec(transInfo *transdsl.TransInfo) (err error) {
	klog.Infof("***DestroyVethPairAction:Exec begin***")
	defer func() {
		if p := recover(); p != nil {
			RecoverErr(p, &err, "DestroyVethPairAction")
		}
		AppendActionName(&err, "DestroyVethPairAction")
	}()
	knitterInfo := transInfo.AppInfo.(*KnitterInfo)
	if !knitterInfo.vethNameOk {
		return nil
	}
	osObj := osobj.GetOsObjSingleton()
	osObj.VethPairRole.Destroy(knitterInfo.vethPair.VethNameOfBridge)
	klog.Infof("***DestroyVethPairAction:Exec end***")
	return nil
}

func (this *DestroyVethPairAction) RollBack(transInfo *transdsl.TransInfo) {
	klog.Infof("***DestroyVethPairAction:RollBack begin***")
	klog.Infof("***DestroyVethPairAction:RollBack end***")
}
