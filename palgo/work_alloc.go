package palgo

import (
	"runtime"
	"sync"

	"github.com/EduardGomezEscandell/algo/utils"
)

// WorkDistribution is an struct containing information regarding workers.
type WorkDistribution struct {
	Work []WorkAlloc
}

// NewWorkDistribution takes a total workload and a minimum workload per worker
// and distributes it along as many workers as possible, up to the number of cores.
func NewWorkDistribution(workload, minWorkload int) WorkDistribution {
	if workload < minWorkload {
		return WorkDistribution{Work: []WorkAlloc{{
			begin: 0,
			end:   workload,
		}}}
	}

	chunkSize := utils.Max(minWorkload, roundUpDiv(workload, runtime.NumCPU()))
	nWorkers := roundUpDiv(workload, chunkSize)

	w := make([]WorkAlloc, nWorkers)
	for i := 0; i < nWorkers; i++ {
		w[i] = WorkAlloc{
			worker: i,
			begin:  i * chunkSize,
			end:    (i + 1) * chunkSize,
		}
	}

	w[len(w)-1].end = utils.Min(w[len(w)-1].end, workload)
	lastChunkSize := w[len(w)-1].len()
	if lastChunkSize < minWorkload {
		w[len(w)-2].end += lastChunkSize
		w = w[:len(w)-1]
	}

	return WorkDistribution{Work: w}
}

// NWorkers is the count of goroutines this distribution will launch.
func (dist WorkDistribution) NWorkers() int {
	return len(dist.Work)
}

// Run executes one the function in each of its goroutines.
func (dist WorkDistribution) Run(f func(WorkAlloc)) {
	var wg sync.WaitGroup
	for _, chunk := range dist.Work {
		wg.Add(1)
		chunk := chunk
		go func() {
			defer wg.Done()
			f(chunk)
		}()
	}

	wg.Wait()
}

// WorkAlloc represents the work to be done by a single worker.
type WorkAlloc struct {
	worker, begin, end int
}

func (a WorkAlloc) len() int {
	return a.end - a.begin
}

func roundUpDiv(x, n int) int {
	return (x + n - 1) / n
}
