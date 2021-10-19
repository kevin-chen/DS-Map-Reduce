package mapreduce

import "fmt"
import "sync"

type WorkerInfo struct {
	address string
}

// Clean up all workers by sending a Shutdown RPC to each one of them Collect
// the number of jobs each work has performed.
func (mr *MapReduce) KillWorkers() []int {
	l := make([]int, 0)
	for _, w := range mr.Workers {
		DPrintf("DoWork: shutdown %s\n", w.address)
		args := &ShutdownArgs{}
		var reply ShutdownReply
		ok := call(w.address, "Worker.Shutdown", args, &reply)
		if ok == false || reply.OK == false {
			fmt.Printf("DoWork: RPC %s shutdown error\n", w.address)
		} else {
			l = append(l, reply.Njobs)
		}
	}
	return l
}

/*
Input: MapReduce object
Output: ints of shutdown workers
Purpose: iterate over all map jobs, run map/reduce job threads for each map/reduce job (when worker is available)
Behavior: first run map job: when a worker is available, run the map job on that worker,
when worker is done, write to registerChannel to notify free worker; use wait group to wait
for all threads to finish
Similarly for Reduce Jobs
*/
func (mr *MapReduce) RunMaster() []int {
	numMapJobs := mr.nMap
	numReduceJobs := mr.nReduce
	var w sync.WaitGroup

	for mapJob := 0; mapJob < numMapJobs; mapJob++ {
		availableWorker := <-mr.registerChannel
		fmt.Println("USING WORKER", availableWorker, "for Map Job")
		w.Add(1)
		go func(worker string, i int) {
			defer w.Done()
			var reply DoJobReply
			args := &DoJobArgs{mr.file, Map, i, mr.nReduce}
			ok := call(worker, "Worker.DoJob", args, &reply)
			if !ok {
				fmt.Println("Map Job", i, "has FAILED")
			} else {
				fmt.Println("Map Job", i, "is SUCCESS")
			}
			mr.registerChannel <- worker
		}(availableWorker, mapJob)
	}

	w.Wait()
	fmt.Println("DONE WITH ALL MAP JOBS")

	for reduceJob := 0; reduceJob < numReduceJobs; reduceJob++ {
		availableWorker := <-mr.registerChannel
		fmt.Println("USING WORKER", availableWorker, "for Reduce Job")
		w.Add(1)
		go func(worker string, i int) {
			defer w.Done()
			var reply DoJobReply
			args := &DoJobArgs{mr.file, Reduce, i, mr.nMap}
			ok := call(worker, "Worker.DoJob", args, &reply)
			if !ok {
				fmt.Println("Reduce Job", i, "has FAILED")
			} else {
				fmt.Println("Reduce Job", i, "is SUCCESS")
			}
			mr.registerChannel <- worker
		}(availableWorker, reduceJob)
	}

	w.Wait()
	fmt.Println("DONE WITH ALL REDUCE JOBS")

	return mr.KillWorkers()
}
