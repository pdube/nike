package nike

import (
	"runtime"
)

//job holds the func to run on the JobWorker
type job struct {
	F func() error
}

//jobStatus contains an error if the job return an error
type jobStatus struct {
	E error
}

//jobWorker executes each job's func and returns the result on the results channel
func jobWorker(id int, jobs <-chan job, results chan<- jobStatus) {
	for j := range jobs {
		e := j.F()
		results <- jobStatus{
			E: e,
		}
	}
}

//JustDoIt parallelizes the execution of the funcs passed in based on the number of CPUs
func JustDoIt(funcs []func() error) []error {
	jobCount := len(funcs)
	jobs := make(chan job, jobCount)
	results := make(chan jobStatus, jobCount)
	cpuCount := runtime.NumCPU()

	for i := 0; i < cpuCount; i++ {
		go jobWorker(i, jobs, results)
	}

	for j := 0; j < jobCount; j++ {
		jobs <- job{
			F: funcs[j],
		}
	}
	close(jobs)
	errs := make([]error, 0)
	for a := 0; a < jobCount; a++ {
		r := <-results
		if r.E != nil {
			errs = append(errs, r.E)
		}
	}
	close(results)
	return errs
}
