package infra

import (
	"time"

	"github.com/HyperNetworks/Knitter/pkg/db-accessor"
	"github.com/HyperNetworks/Knitter/pkg/klog"
	"github.com/HyperNetworks/Knitter/pkg/uuid"
)

var etcdDataBase dbaccessor.DbAccessor = nil

//todo how to use managerUUID
var managerUUID = ""

func SetDataBase(i dbaccessor.DbAccessor) error {
	etcdDataBase = i
	return nil
}

var GetDataBase = func() dbaccessor.DbAccessor {
	return etcdDataBase
}

func CheckDB() error {
	for {
		/*wait for etcd database OK*/
		e := dbaccessor.CheckDataBase(etcdDataBase)
		if e != nil {
			klog.Error("DataBase Config is ERROR", e)
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	for {
		SetClusterID()
		if GetClusterID() != uuid.NIL.String() {
			break
		}
	}

	return nil
}
