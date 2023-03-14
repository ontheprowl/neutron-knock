package endpoints

import (
	"log"

	"github.com/gofiber/fiber/v2"
	channels "neutron.money/knock/channels/sendinblue"
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

	var channel types.ChannelProvider = &channels.SendInBlueProvider{
		ApiKey:     "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",
		PartnerKey: "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",
	}

	channel.Init()

	scheduler.CronWithSeconds(job.Schedule).Tag(job.JobType.String()).Tag(job.Id).Do(
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
