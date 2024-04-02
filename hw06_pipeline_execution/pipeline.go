package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	inCh := in
	for _, stage := range stages {
		inCh = stage(execStage(inCh, done))
	}

	return execStage(inCh, done)
}

func execStage(in In, done In) Out {
	outCh := make(Bi)

	go func() {
		defer close(outCh)
		for {
			select {
			case <-done:
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				outCh <- value
			}
		}
	}()

	return outCh
}
