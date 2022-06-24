package main

import (
	"abeam/model/avromodel"
	"abeam/udf"
	"context"
	"flag"
	"fmt"
	"reflect"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/avroio"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
)

func main() {
	flag.Parse()
	beam.Init()
	// In order to start creating the pipeline for execution, a Pipeline object is needed.
	p := beam.NewPipeline()
	s := p.Root()

	syncPCol := avroio.Read(s, "/Users/haupc/project/abeam/bsc-*.avro", reflect.TypeOf(avromodel.SyncEvent{}))

	mappedPCol := beam.ParDo(s, udf.ToKVEvent, syncPCol)
	groupedPCol := beam.GroupByKey(s, mappedPCol)
	collectedPCol := beam.ParDo(s, udf.CollectElements, groupedPCol)

	// Write the output to a file.
	textio.Write(s, "/Users/haupc/project/abeam/output.txt", formated)

	if err := beamx.Run(context.Background(), p); err != nil {
		fmt.Printf("Pipeline failed: %v", err)
	}
}
