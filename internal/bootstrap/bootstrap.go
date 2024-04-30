package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"botperational/config"
	"botperational/internal/adapter/rest"
	"botperational/internal/application/service"
	"botperational/internal/pkg/validator"

	"github.com/jasonlvhit/gocron"
	"github.com/runsystemid/golog"
	"github.com/runsystemid/gontainer"
)

var appContainer = gontainer.New()

func Run(conf *config.Config) {
	appContainer.RegisterService("config", conf)

	// Initialize struct validator
	appContainer.RegisterService("validator", validator.NewGoValidator())

	bootstrapContext := context.Background()
	golog.Info(bootstrapContext, "Serving...")

	// Register adapter
	RegisterDatabase()
	RegisterRest()
	RegisterRepository()

	// Register application
	RegisterService()

	// Startup the container
	if err := appContainer.Ready(); err != nil {
		golog.Panic(bootstrapContext, "Failed to populate service", err)
	}

	// Start server
	fiberApp := appContainer.GetServiceOrNil("fiber").(*rest.Fiber)
	errs := make(chan error, 2)
	go func() {
		golog.Info(bootstrapContext, fmt.Sprintf("Listening on port :%d", conf.Http.Port))
		errs <- fiberApp.Listen(fmt.Sprintf(":%d", conf.Http.Port))
	}()

	golog.Info(bootstrapContext, "Your app started")

	//cron
	startScheduler(bootstrapContext, conf)

	gracefulShutdown(bootstrapContext)
}

func gracefulShutdown(ctx context.Context) {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	delay := 5 * time.Second

	golog.Info(ctx, fmt.Sprintf("Signal termination received. Waiting %v to shutdown.", delay))

	time.Sleep(delay)

	golog.Info(ctx, "Cleaning up resources...")

	// This will shuting down all the resources
	appContainer.Shutdown()

	golog.Info(ctx, "Bye")
}

func startScheduler(ctx context.Context, conf *config.Config) {

	IsSchedulerSettingExisted := false
	g := gocron.NewScheduler()

	if len(conf.Scheduler.MultiTime) > 0 {
		s := strings.Split(conf.Scheduler.MultiTime, ";")
		for _, x := range s {
			if len(x) > 0 {
				today := time.Now().Weekday()

				if today != time.Saturday && today != time.Sunday {
					g.Every(1).Weekday(time.Monday).At(x).Do(processJob)
					IsSchedulerSettingExisted = true
				}
			}
		}
	}

	if len(conf.Scheduler.IntervalInSecond) > 0 {
		i, err := strconv.ParseUint(conf.Scheduler.IntervalInSecond, 10, 64)
		if err != nil {
			golog.Error(ctx, "Error convert scheduler interval", err)
		}
		g.Every(i).Seconds().Do(processJob)
		IsSchedulerSettingExisted = true
	}
	if IsSchedulerSettingExisted {
		<-g.Start()
	}

}

func processJob() {
	bootstrapContext := context.Background()

	//on leave
	onLeaveService := appContainer.GetServiceOrNil("onLeaveService").(*service.OnLeaveService)
	err := onLeaveService.ProcessOnLeaveData(bootstrapContext)
	if err != nil {
		golog.Error(
			bootstrapContext,
			fmt.Sprintf("Error on leave ProcessJob: %s", err.Error()),
			err,
		)
	}

	//on birthday
	onBirthdayService := appContainer.GetServiceOrNil("onBirthdayService").(*service.OnBirthdayService)
	err = onBirthdayService.ProcessOnBirthdayData(bootstrapContext)
	if err != nil {
		golog.Error(
			bootstrapContext,
			fmt.Sprintf("Error on birthday ProcessJob: %s", err.Error()),
			err,
		)
	}
}
