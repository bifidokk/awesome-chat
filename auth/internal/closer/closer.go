package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = New()

// Closer is a struct that manages the graceful shutdown of multiple resources.
type Closer struct {
	mu    sync.Mutex
	once  sync.Once
	done  chan struct{}
	funcs []func() error
}

// Add adds one or more functions to the globalCloser that will be called during shutdown.
func Add(f ...func() error) {
	globalCloser.Add(f...)
}

// Wait waits for the globalCloser to complete all shutdown functions.
func Wait() {
	globalCloser.Wait()
}

// CloseAll initiates the shutdown process for the globalCloser, ensuring each function is called.
func CloseAll() {
	log.Println("dd")
	globalCloser.CloseAll()
}

// New creates a new Closer instance.
func New(sig ...os.Signal) *Closer {
	c := &Closer{
		done: make(chan struct{}),
	}

	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.CloseAll()
		}()
	}

	return c
}

// Add adds one or more functions to the Closer that will be called during shutdown.
func (c *Closer) Add(f ...func() error) {
	c.mu.Lock()
	c.funcs = append(c.funcs, f...)
	c.mu.Unlock()
}

// Wait waits for the Closer to complete all shutdown functions.
func (c *Closer) Wait() {
	<-c.done
}

// CloseAll initiates the shutdown process for the Closer, ensuring each function is called only once.
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.mu.Unlock()

		// call all Closer funcs async
		errs := make(chan error, len(funcs))
		for _, f := range funcs {
			go func(f func() error) {
				errs <- f()
			}(f)
		}

		for i := 0; i < cap(errs); i++ {
			if err := <-errs; err != nil {
				log.Println("error returned from Closer")
			}
		}
	})
}
