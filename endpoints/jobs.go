package endpoints

import (
	"log"

	"github.com/gofiber/fiber/v2"

	channels "neutron.money/knock/channels"

	scheduler "neutron.money/knock/scheduler"
	"neutron.money/knock/types"
	utils "neutron.money/knock/utils"
)

func AffixJobsRoutes(router *fiber.Router) {
	jobsRouter := (*router).Group("/jobs")
	jobsRouter.Post("/create", createJob)
	// jobsRouter.Post("/update", updateJob)
	// jobsRouter.Delete("/delete", deleteJob)
	jobsRouter.Get("/clear", clearJobs)
	// jobsRouter.Get("/get", getJob)
	jobsRouter.Get("/list", listJobs)
}

func createJob(ctx *fiber.Ctx) error {
	job := new(types.Job)

	scheduler := scheduler.GetScheduler()

	if err := ctx.BodyParser(job); err != nil {
		return ctx.JSON(map[string]interface{}{
			"status":  -1,
			"message": err.Error(),
		})
	}

	if job.CallerID == "" {
		return ctx.JSON(map[string]interface{}{
			"status":  -1,
			"message": "No caller ID specified...",
		})
	}

	if job.JobType != types.Email && job.JobType != types.Whatsapp {
		return ctx.JSON(map[string]interface{}{
			"status":  -1,
			"message": "Could not create the specified job... JobType is invalid",
		})
	}

	channel, err := channels.GetChannel(job.JobType)

	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"status":  -1,
			"message": "No channels configured for the specified jobtype",
		})
	}

	jobRef, err := scheduler.Cron(job.Schedule).Tag(job.CallerID).Tag(job.JobType.String()).Tag(job.Id).Do(
		channel.SendMessage, job.Data, job.JobType)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return ctx.JSON(map[string]interface{}{
		"status": 0,
		"message": map[string]interface{}{
			"tags":    jobRef.Tags(),
			"nextRun": jobRef.NextRun(),
		},
	})
}

func clearJobs(ctx *fiber.Ctx) error {
	scheduler := scheduler.GetScheduler()
	scheduler.Clear()
	return ctx.JSON(map[string]interface{}{
		"status":  0,
		"message": "All jobs cleared...",
	})
}

func listJobs(ctx *fiber.Ctx) error {
	scheduler := scheduler.GetScheduler()
	jobs := scheduler.Jobs()

	if len(jobs) <= 0 {
		return ctx.JSON("No jobs scheduled currently...")
	}

	result := make([][]string, len(jobs))

	for _, v := range jobs {
		log.Print(v.Tags())
		result = append(result, v.Tags())
	}

	filteredJobs := utils.Filter(result, func(jobDetails []string) bool {
		return jobDetails != nil
	})

	return ctx.JSON(filteredJobs)
}
