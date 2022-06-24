package udf

import (
	"abeam/libs/model"
	"abeam/model/avromodel"
	"fmt"

	libsModel "abeam/libs/model"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
)

func init() {
	beam.RegisterFunction(ToKVEvent)
	beam.RegisterFunction(CollectElements)

	beam.RegisterFunction(FormatKV)
}

var syncModeler = avromodel.SyncModeler[libsModel.UniPairSyncEvent, avromodel.SyncEvent]{}

func ToKVEvent(ele avromodel.SyncEvent, emit func(uint64, model.UniPairSyncEvent)) {
	syncModel, err := syncModeler.ToModel(ele)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v:%v\n", syncModel.BlockNumber, syncModel)
	emit(syncModel.BlockNumber, syncModel)
}

func CollectElements(blockNum uint64, iter func(*model.UniPairSyncEvent) bool) (uint64, []model.UniPairSyncEvent) {
	var groupedSyncEvents []model.UniPairSyncEvent
	var ele model.UniPairSyncEvent
	for iter(&ele) {
		groupedSyncEvents = append(groupedSyncEvents, ele)
	}
	return blockNum, groupedSyncEvents
}

func FormatKV(key uint64, value []model.UniPairSyncEvent) string {
	return fmt.Sprintf("%v:%v", key, value)
}
