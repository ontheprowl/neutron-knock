package endpoints

import (
	"log"

	"github.com/gofiber/fiber/v2"

	channels "neutron.money/knock/channels"

	scheduler "neutron.money/knock/scheduler"
	"neutron.money/knock/types"
)

func AffixJobsRoutes(router *fiber.Router) {
	jobsRouter := (*router).Group("/jobs")
	jobsRouter.Post("/create", createJob)
	// jobsRouter.Post("/update", updateJob)
	// jobsRouter.Delete("/delete", deleteJob)
	// jobsRouter.Get("/get", getJob)
	jobsRouter.Get("/list", listJobs)
}

func createJob(ctx *fiber.Ctx) error {
	job := new(types.Job)
	scheduler := scheduler.GetScheduler()

	if err := ctx.BodyParser(job); err != nil {
		return err
	}

	if job.JobType != types.Email && job.JobType != types.Whatsapp {
		return ctx.JSON("Could not create the specified job... JobType is invalid")
	}

	channel, err := channels.GetChannel(job.JobType)

	if err != nil {
		return ctx.JSON("No channels configured for the specified jobtype")
	}

	scheduler.Cron(job.Schedule).Tag(job.JobType.String()).Tag(job.Id).Do(
		channel.SendMessage, job.Data, job.JobType)

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
