package packet

import (
	"sync"
	"sync/atomic"

	"github.com/google/gopacket"
	"github.com/google/gopacket/tcpassembly"
)

// HTTPStreamFactory implements StreamFactory interface for tcpassembly
type HTTPStreamFactory struct {
	runningStream *int32
	wg            *sync.WaitGroup
	seq           *uint
	uniStreams    *map[streamKey]*httpStreamPair
	eventChan     chan<- interface{}
}

// NewHTTPStreamFactory create a NewHTTPStreamFactory.
func NewHTTPStreamFactory(out chan<- interface{}) HTTPStreamFactory {
	var f HTTPStreamFactory
	f.seq = new(uint)
	*f.seq = 0
	f.wg = new(sync.WaitGroup)
	f.uniStreams = new(map[streamKey]*httpStreamPair)
	*f.uniStreams = make(map[streamKey]*httpStreamPair)
	f.eventChan = out
	f.runningStream = new(int32)
	return f
}

// Wait 等待所有流完成.
func (f HTTPStreamFactory) Wait() {
	f.wg.Wait()
}

// RunningStreamCount 当前流计数.
func (f *HTTPStreamFactory) RunningStreamCount() int32 {
	return atomic.LoadInt32(f.runningStream)
}

// runStreamPair 运行一个httpStreamPair.
func (f *HTTPStreamFactory) runStreamPair(streamPair *httpStreamPair) {
	atomic.AddInt32(f.runningStream, 1)

	defer f.wg.Done()
	defer func() { atomic.AddInt32(f.runningStream, -1) }()
	streamPair.run()
}

// New creates a HTTPStreamFactory.
func (f HTTPStreamFactory) New(netFlow, tcpFlow gopacket.Flow) (ret tcpassembly.Stream) {
	reverseKey := streamKey{netFlow.Reverse(), tcpFlow.Reverse()}
	streamPair, ok := (*f.uniStreams)[reverseKey]
	if ok {
		if streamPair.upStream == nil {
			panic("unbelievable!?")
		}
		delete(*f.uniStreams, reverseKey)
		key := streamKey{netFlow, tcpFlow}
		s := newHTTPStream(key)
		streamPair.downStream = &s
		ret = s
	} else {
		streamPair = newHTTPStreamPair(*f.seq, f.eventChan)
		key := streamKey{netFlow, tcpFlow}
		s := newHTTPStream(key)
		streamPair.upStream = &s
		(*f.uniStreams)[key] = streamPair
		*f.seq++
		f.wg.Add(1)
		go f.runStreamPair(streamPair)
		ret = s
	}
	return
}
