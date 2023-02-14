package endpoints

import (
	"log"

	"github.com/gofiber/fiber/v2"
	scheduler "neutron.money/knock/scheduler"
)

type JobType int

const (
	Email    = 0
	Whatsapp = 1
)

type Job struct {
	Id      string  `json:"id" xml:"id" form:"id"`
	JobType JobType `json:"jobType" xml:"jobType" form:"jobType"`
}

func AffixJobsRoutes(router *fiber.Router) {
	jobsRouter := (*router).Group("/jobs")
	jobsRouter.Get("/create", createJob)
	// jobsRouter.Post("/update", updateJob)
	// jobsRouter.Delete("/delete", deleteJob)
	// jobsRouter.Get("/get", getJob)
	jobsRouter.Get("/list", listJobs)
}

func createJob(ctx *fiber.Ctx) error {
	// job := new(Job)
	scheduler := scheduler.GetScheduler()

	// if err := ctx.BodyParser(job); err != nil {
	// 	return err
	// }

	// if job.JobType != Email && job.JobType != Whatsapp {
	// 	return ctx.JSON("Could not create the specified job... JobType is invalid")
	// }

	scheduler.Every(5).Seconds().Tag("testJob").Do(func() {
		log.Println("New Task Scheduled...")
	})

	// log.Println(job.Id)      // john
	// log.Println(job.JobType) // doe
	return ctx.JSON("success!")
}

func listJobs(ctx *fiber.Ctx) error {
	scheduler := scheduler.GetScheduler()
	jobs := scheduler.Jobs()

	log.Println(jobs)
	result := make([][]string, len(jobs))

	for _, v := range jobs {
		log.Print(v.Tags())
		result = append(result, v.Tags())
	}

	if len(jobs) <= 0 {
		return ctx.JSON("No jobs scheduled currently...")
	}
	return ctx.JSON(result)
}
