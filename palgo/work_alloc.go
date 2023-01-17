package palgo

import (
	"runtime"
	"sync"

	"github.com/EduardGomezEscandell/algo/utils"
)

// minWork is amount of work per worker. Some algorithms require
// at least two items per worker, so this cannot be pushed below 2.
const minWork = 3

// workAlloc represents the work to be done
// by a single worker.
type workAlloc struct {
	worker, begin, end int
}

func (r workAlloc) len() int {
	return r.end - r.begin
}

func roundUpDiv(x, n int) int {
	return (x + n - 1) / n
}

func workAllocation(workload int) []workAlloc {
	if workload < minWork {
		return []workAlloc{{
			begin: 0,
			end:   workload,
		}}
	}

	chunkSize := utils.Max(minWork, roundUpDiv(workload, runtime.NumCPU()))
	nWorkers := roundUpDiv(workload, chunkSize)

	r := make([]workAlloc, nWorkers)
	for i := 0; i < nWorkers; i++ {
		r[i] = workAlloc{
			worker: i,
			begin:  i * chunkSize,
			end:    (i + 1) * chunkSize,
		}
	}

	r[len(r)-1].end = utils.Min(r[len(r)-1].end, workload)
	lastChunkSize := r[len(r)-1].len()
	if lastChunkSize < minWork {
		r[len(r)-2].end += lastChunkSize
		r = r[:len(r)-1]
	}

	return r
}

func distribute(work []workAlloc, f func(workAlloc)) {
	var wg sync.WaitGroup
	for _, chunk := range work {
		wg.Add(1)
		chunk := chunk
		go func() {
			defer wg.Done()
			f(chunk)
		}()
	}

	wg.Wait()
}
