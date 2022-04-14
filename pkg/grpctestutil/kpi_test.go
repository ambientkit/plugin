package grpctestutil_test

import (
	"context"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ambientkit/ambient/pkg/ambientapp"
)

// Total number of tests to run.
const numberTests = 5000
const numberTestsConcurrent = 100

// Number of requests to send out at the same time.
const concurrentJobs = 50

func TestKPIStandard(t *testing.T) {
	if testing.Short() {
		t.Skip("skip kpi in short mode")
	}

	os.Setenv("AMB_LOGLEVEL", "FATAL")
	app := standardSetup(t)
	responseTime(t, app, numberTests)
}

func TestKPIGRPC(t *testing.T) {
	if testing.Short() {
		t.Skip("skip kpi in short mode")
	}

	os.Setenv("AMB_LOGLEVEL", "FATAL")
	app := grpcSetup(t)
	responseTime(t, app, numberTests)
}

func TestKPIStandardConcurrent(t *testing.T) {
	if testing.Short() {
		t.Skip("skip kpi in short mode")
	}

	os.Setenv("AMB_LOGLEVEL", "FATAL")
	app := standardSetup(t)
	responseTimeConcurrent(t, app, numberTestsConcurrent, concurrentJobs)
}

func TestKPIGRPCConcurrent(t *testing.T) {
	if testing.Short() {
		t.Skip("skip kpi in short mode")
	}

	os.Setenv("AMB_LOGLEVEL", "FATAL")
	app := grpcSetup(t)
	responseTimeConcurrent(t, app, numberTestsConcurrent, concurrentJobs)
}

func responseTime(t *testing.T, app *ambientapp.App, totalTests int) {
	mux := setGrants(t, context.Background(), app)

	arrTimes := make([]float64, 0)
	var max int64 = 0
	var min int64 = 1000

	for n := 0; n < numberTests; n++ {
		start := time.Now()
		// Use endpoint since it uses: middleware, router, funcmaps, assets.
		doRequest(t, mux, httptest.NewRequest("GET", "/assetsHello", nil))
		actualMS := time.Since(start).Milliseconds()
		arrTimes = append(arrTimes, float64(actualMS))
		if actualMS > max {
			max = actualMS
		}
		if actualMS < min {
			min = actualMS
		}
	}

	var avg float64 = 0
	for _, v := range arrTimes {
		avg += v

	}
	avg = avg / float64(len(arrTimes))

	t.Logf("Response time (tests: %v): average: %vms | max: %vms | min: %vms\n", len(arrTimes), avg, max, min)
}

func responseTimeConcurrent(t *testing.T, app *ambientapp.App, totalTests int, concurrent int) {
	mux := setGrants(t, context.Background(), app)

	arrTimes := make([]float64, 0)
	var max int64 = 0
	var min int64 = 1000
	var m sync.Mutex

	for n := 0; n < totalTests; n++ {
		var wg sync.WaitGroup
		for i := 0; i < concurrent; i++ {
			wg.Add(1)
			go func() {
				start := time.Now()
				// Use endpoint since it uses: middleware, router, funcmaps, assets.
				doRequest(t, mux, httptest.NewRequest("GET", "/assetsHello", nil))
				actualMS := time.Since(start).Milliseconds()
				m.Lock()
				arrTimes = append(arrTimes, float64(actualMS))
				if actualMS > max {
					max = actualMS
				}
				if actualMS < min {
					min = actualMS
				}
				m.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()
	}

	var avg float64 = 0
	for _, v := range arrTimes {
		avg += v

	}
	avg = avg / float64(len(arrTimes))

	t.Logf("Response time (tests: %v): average: %vms | max: %vms | min: %vms\n", len(arrTimes), avg, max, min)
}
