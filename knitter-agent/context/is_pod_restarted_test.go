package context

import (
	"github.com/HyperNetworks/Knitter/knitter-agent/domain/cni"
	"github.com/HyperNetworks/Knitter/knitter-agent/domain/object/knitter-obj"
	"github.com/HyperNetworks/Knitter/knitter-agent/domain/object/pod-obj"
	"github.com/HyperNetworks/Knitter/knitter-agent/domain/object/port-obj"
	"github.com/HyperNetworks/Knitter/knitter-agent/mock/mock-pkg/mock-dbaccessor"
	"github.com/HyperNetworks/Knitter/knitter-manager/err-obj"
	"github.com/HyperNetworks/Knitter/pkg/trans-dsl"
	"github.com/coreos/etcd/client"
	"github.com/golang/gostub"
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsPodRestarted(t *testing.T) {
	action := &IsPodRestarted{}

	cniParam := &cni.CniParam{PodNs: "nw001"}
	knitterObj := &knitterobj.KnitterObj{CniParam: cniParam}

	var portObjs []*portobj.PortObj
	portObjs = append(portObjs,
		&portobj.PortObj{EagerAttr: portobj.PortEagerAttr{NetworkName: "lan"}},
		&portobj.PortObj{EagerAttr: portobj.PortEagerAttr{NetworkName: "net_api"}},
		&portobj.PortObj{EagerAttr: portobj.PortEagerAttr{NetworkName: "control"}})
	podObj := &podobj.PodObj{PortObjs: portObjs}
	knitterInfo := &KnitterInfo{KnitterObj: knitterObj, podObj: podObj}

	cfgMock := gomock.NewController(t)
	defer cfgMock.Finish()
	mockDB := mockdbaccessor.NewMockDbAccessor(cfgMock)
	var ports []*client.Node
	node := &client.Node{Key: "etcd_key"}
	ports = append(ports, node)

	mockDB.EXPECT().ReadDir(gomock.Any()).Return(ports, nil)

	stubs := gostub.StubFunc(&cni.GetGlobalContext, &cni.AgentContext{RemoteDB: mockDB})
	defer stubs.Reset()
	transInfo := &transdsl.TransInfo{AppInfo: knitterInfo}

	convey.Convey("TestIsPodRestarted for succ\n", t, func() {
		ok := action.Ok(transInfo)
		convey.So(ok, convey.ShouldEqual, true)
	})
}

func TestIsPodRestarted_false(t *testing.T) {
	action := &IsPodRestarted{}

	cniParam := &cni.CniParam{PodNs: "nw001"}
	knitterObj := &knitterobj.KnitterObj{CniParam: cniParam}

	var portObjs []*portobj.PortObj
	portObjs = append(portObjs,
		&portobj.PortObj{EagerAttr: portobj.PortEagerAttr{NetworkName: "lan"}},
		&portobj.PortObj{EagerAttr: portobj.PortEagerAttr{NetworkName: "net_api"}},
		&portobj.PortObj{EagerAttr: portobj.PortEagerAttr{NetworkName: "control"}})
	podObj := &podobj.PodObj{PortObjs: portObjs}
	knitterInfo := &KnitterInfo{KnitterObj: knitterObj, podObj: podObj}
	transInfo := &transdsl.TransInfo{AppInfo: knitterInfo}

	cfgMock := gomock.NewController(t)
	defer cfgMock.Finish()
	mockDB := mockdbaccessor.NewMockDbAccessor(cfgMock)
	var ports []*client.Node
	node := &client.Node{Key: "etcd_key"}
	ports = append(ports, node)

	mockDB.EXPECT().ReadDir(gomock.Any()).Return(ports, errobj.ErrAny)

	stubs := gostub.StubFunc(&cni.GetGlobalContext, &cni.AgentContext{RemoteDB: mockDB})
	defer stubs.Reset()

	action.RollBack(transInfo)
	convey.Convey("TestIsPodRestarted for return false\n", t, func() {
		ok := action.Ok(transInfo)
		convey.So(ok, convey.ShouldEqual, false)
	})
}
