package context

import (
	"github.com/HyperNetworks/Knitter/knitter-agent/domain/cni"
	"github.com/HyperNetworks/Knitter/knitter-agent/domain/object/knitter-obj"
	"github.com/HyperNetworks/Knitter/knitter-agent/domain/object/port-obj"
	"github.com/HyperNetworks/Knitter/knitter-agent/infra"
	"github.com/HyperNetworks/Knitter/pkg/trans-dsl"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsTenantNetworkNeedDel(t *testing.T) {
	action := IsTenantNetworkNeedDel{}

	cniParam := &cni.CniParam{PodNs: "nw001"}
	knitterObj := &knitterobj.KnitterObj{CniParam: cniParam}

	portObj := &portobj.PortObj{LazyAttr: portobj.PortLazyAttr{NetAttr: portobj.NetworkAttrs{ID: "id1"}}}

	knitterInfo := &KnitterInfo{KnitterObj: knitterObj, portObj: portObj}
	transInfo := &transdsl.TransInfo{AppInfo: knitterInfo}

	infra.Init()

	convey.Convey("TestIsTenantNetworkNeedDel for succ!\n", t, func() {
		ok := action.Ok(transInfo)
		convey.So(ok, convey.ShouldEqual, false)
	})
}
