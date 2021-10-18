package mapreduce

import "fmt"

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

func (mr *MapReduce) RunMaster() []int {
	numMapJobs := mr.nMap
	numReduceJobs := mr.nReduce

	for mapJob := 0; mapJob < numMapJobs; mapJob++ {
		fmt.Println(<-mr.registerChannel)
	}

	for reduceJob := 0; reduceJob < numReduceJobs; reduceJob++ {

	}

	return mr.KillWorkers()
}
