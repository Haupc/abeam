package udf

import (
	"abeam/libs/model"
	"abeam/model/avromodel"
	"fmt"

	libsModel "abeam/libs/model"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/slices"
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

func CollectElements(blockNum uint64, iter func(*model.UniPairSyncEvent) bool) []model.UniPairSyncEvent {
	var groupedSyncEvents []model.UniPairSyncEvent
	var ele model.UniPairSyncEvent
	for iter(&ele) {
		groupedSyncEvents = append(groupedSyncEvents, ele)
	}
	return groupedSyncEvents
}

func FormatKV(key uint64, value []model.UniPairSyncEvent) string {
	return fmt.Sprintf("%v:%v", key, value)
}

func BlockToCandles(values []model.UniPairSyncEvent, emit func(Candle)) {
	syncEventMap := make(map[common.Address][]model.UniPairSyncEvent) // map pool address, sync event in same block
	for _, syncEvent := range values {
		syncEventMap[syncEvent.PoolAddress] = append(syncEventMap[syncEvent.PoolAddress], syncEvent)
	}

	// sort by log index
	for poolAddress := range syncEventMap {
		slices.SortFunc(syncEventMap[poolAddress], func(a, b model.UniPairSyncEvent) bool {
			return a.BaseEvent.LogIndex < b.BaseEvent.LogIndex
		})
		var o, h, c, l float64
	}
}
