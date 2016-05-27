package mapreduce

import "container/list"
import "fmt"
import "log"
import "sync"

/*
TODO :

Close workerAgents, and all goroutines cleanly on exit
Reuse DoJob objects to avoid creating too many DoJob objects


*/

type WorkerInfo struct {
	address string
	// You can add definitions here.
}

// Clean up all workers by sending a Shutdown RPC to each one of them Collect
// the number of jobs each work has performed.
func (mr *MapReduce) KillWorkers() *list.List {
	l := list.New()
	for _, w := range mr.Workers {
		DPrintf("DoWork: shutdown %s\n", w.address)
		args := &ShutdownArgs{}
		var reply ShutdownReply
		ok := call(w.address, "Worker.Shutdown", args, &reply)
		if ok == false {
			fmt.Printf("DoWork: RPC %s shutdown error\n", w.address)
		} else {
			l.PushBack(reply.Njobs)
		}
	}
	return l
}

func workerAgent(done <-chan struct{},
	worker string,
	jobChannel <-chan DoJobArgs,
	jobCompletedChannel chan DoJobArgs) {

	for {
		select {
		case jobArgs := <-jobChannel:
			fmt.Printf("%s Job Received %d\n", worker, jobArgs.JobNumber)
			reply := DoJobReply{OK: false}
			call(worker, "Worker.DoJob", jobArgs, &reply)

			if reply.OK {
				fmt.Printf("%s Job Completed %d\n", worker, jobArgs.JobNumber)
				jobCompletedChannel <- jobArgs
			} else {
				log.Fatalf("WorkerAgent Job failed %i\n", jobArgs.JobNumber)
			}
		case <-done:
			fmt.Printf("WorkerAgent Exiting %s\n", worker)
			return
		}
	}
}

func (mr *MapReduce) startWorkerAgents(done <-chan struct{}, jobChannel <-chan DoJobArgs, jobCompletedChannel chan DoJobArgs) {
	var wg sync.WaitGroup
	for {
		select {
		case worker := <-mr.registerChannel:
			wg.Add(1)
			go func() {
				workerAgent(done, worker, jobChannel, jobCompletedChannel)
				wg.Done()
			}()
		case <-done:
			break
		}
	}

	wg.Wait()
}

func (mr *MapReduce) generateMapJobs(jobChannel chan DoJobArgs) {
	for i := 0; i < mr.nMap; i++ {
		jobChannel <- DoJobArgs{File: mr.file,
			Operation:     Map,
			JobNumber:     i,
			NumOtherPhase: mr.nReduce}
	}
}

func (mr *MapReduce) generateReduceJobs(jobChannel chan DoJobArgs) {
	for j := 0; j < mr.nReduce; j++ {
		jobChannel <- DoJobArgs{File: mr.file,
			Operation:     Reduce,
			JobNumber:     j,
			NumOtherPhase: mr.nMap}
	}
}

func (mr *MapReduce) RunMaster() *list.List {
	var wg sync.WaitGroup

	done := make(chan struct{})

	jobs := make(chan DoJobArgs)
	jobCompletedChannel := make(chan DoJobArgs, 5)

	wg.Add(1)
	go func() {
		mr.startWorkerAgents(done, jobs, jobCompletedChannel)
		wg.Done()
	}()

	go mr.generateMapJobs(jobs)

	completedMaps := 0
	for jobArgs := range jobCompletedChannel {
		fmt.Printf("Map Job Completed %d\n", jobArgs.JobNumber)
		completedMaps += 1
		if completedMaps == mr.nMap {
			break
		}
	}

	fmt.Printf("All Maps Completed\n")

	go mr.generateReduceJobs(jobs)

	completedReduces := 0
	for jobArgs := range jobCompletedChannel {
		fmt.Printf("Reduce Job Completed %d\n", jobArgs.JobNumber)
		completedReduces += 1
		if completedReduces == mr.nReduce {
			break
		}
	}
	fmt.Printf("All Reduces Completed\n")
	close(done)
	wg.Wait()

	return mr.KillWorkers()
}
