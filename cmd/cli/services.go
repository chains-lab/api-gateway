package cli

import (
	"context"
	"sync"

	"github.com/chains-lab/api-gateway/internal/api"
	"github.com/chains-lab/api-gateway/internal/config"
)

func runServices(ctx context.Context, cfg config.Config, wg *sync.WaitGroup) {
	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	API := api.NewAPI(cfg)
	run(func() { API.Run(ctx) })
}
