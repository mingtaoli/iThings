package svc

import (
	"github.com/i-Things/things/shared/domain/schema"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/config"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceMsg"
	"github.com/i-Things/things/src/dmsvr/internal/domain/service/deviceData"
	"github.com/i-Things/things/src/dmsvr/internal/repo/event/publish/dataUpdate"
	"github.com/i-Things/things/src/dmsvr/internal/repo/event/publish/pubDev"
	"github.com/i-Things/things/src/dmsvr/internal/repo/mysql"
	"github.com/i-Things/things/src/dmsvr/internal/repo/tdengine/deviceDataRepo"
	"github.com/i-Things/things/src/dmsvr/internal/repo/tdengine/hubLogRepo"
	"github.com/i-Things/things/src/dmsvr/internal/repo/tdengine/sdkLogRepo"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"os"
)

type ServiceContext struct {
	Config         config.Config
	DeviceInfo     mysql.DeviceInfoModel
	ProductInfo    mysql.ProductInfoModel
	DmDB           mysql.DmModel
	DeviceID       *utils.SnowFlake
	ProductID      *utils.SnowFlake
	PubDev         pubDev.PubDev
	DataUpdate     dataUpdate.DataUpdate
	Store          kv.Store
	DeviceDataRepo deviceData.DeviceDataRepo
	HubLogRepo     deviceMsg.HubLogRepo
	SchemaRepo     schema.SchemaRepo
	SDKLogRepo     deviceMsg.SDKLogRepo
	FirmwareInfo   mysql.ProductFirmwareModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	hubLog := hubLogRepo.NewHubLogRepo(c.TDengine.DataSource)
	sdkLog := sdkLogRepo.NewSDKLogRepo(c.TDengine.DataSource)

	//TestTD(td)
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	di := mysql.NewDeviceInfoModel(conn)
	pi := mysql.NewProductInfoModel(conn)
	pt := mysql.NewProductSchemaModel(conn)
	tr := mysql.NewSchemaRepo(pt)
	deviceData := deviceDataRepo.NewDeviceDataRepo(c.TDengine.DataSource, tr.GetSchemaModel)
	fr := mysql.NewProductFirmwareModel(conn)
	DmDB := mysql.NewDmModel(conn)
	store := kv.NewStore(c.CacheRedis)
	nodeId := utils.GetNodeID(c.CacheRedis, c.Name)
	DeviceID := utils.NewSnowFlake(nodeId)
	ProductID := utils.NewSnowFlake(nodeId)
	pd, err := pubDev.NewPubDev(c.Event)
	if err != nil {
		logx.Error("NewPubDev err", err)
		os.Exit(-1)
	}
	du, err := dataUpdate.NewDataUpdate(c.Event)
	if err != nil {
		logx.Error("NewDataUpdate err", err)
		os.Exit(-1)
	}
	return &ServiceContext{
		Config:         c,
		DeviceInfo:     di,
		ProductInfo:    pi,
		FirmwareInfo:   fr,
		SchemaRepo:     tr,
		DmDB:           DmDB,
		DeviceID:       DeviceID,
		ProductID:      ProductID,
		PubDev:         pd,
		DataUpdate:     du,
		Store:          store,
		DeviceDataRepo: deviceData,
		HubLogRepo:     hubLog,
		SDKLogRepo:     sdkLog,
	}
}
