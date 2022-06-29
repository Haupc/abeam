package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
)

func main() {
	flag.Parse()
	beam.Init()
	// In order to start creating the pipeline for execution, a Pipeline object is needed.
	p := beam.NewPipeline()
	s := p.Root()

	firstPCol := beam.Create(s, "first", "mid1", "mid2", "mid3", "last")

	beam.ParDo(s, func(ele string, emit func(string, string)) {
		emit(ele, ele)
	}, firstPCol)
	// textio.Write(s, "/Users/haupc/project/abeam/output.txt", firstPCol)

	if err := beamx.Run(context.Background(), p); err != nil {
		fmt.Printf("Pipeline failed: %v", err)
	}
}
