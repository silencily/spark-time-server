package task

import (
	"github.com/kataras/golog"
	cron2 "github.com/robfig/cron/v3"
	"github.com/silencily/sparktime/core"
	"github.com/silencily/sparktime/services"
	"sync"
)

var once sync.Once

var tsm *taskScheduleManager

func getLogger() *golog.Logger {
	return core.GetLogger("TaskScheduleManager")
}

type TaskScheduleManager interface {
	Start()
	Stop()
}

type taskScheduleManager struct {
	cron         *cron2.Cron
	sparkService services.SparkService
}

func GetTaskScheduleManager() TaskScheduleManager {
	once.Do(func() {
		tsm = &taskScheduleManager{
			cron:         cron2.New(),
			sparkService: services.NewSparkService(),
		}
	})
	return tsm
}

func (tsm *taskScheduleManager) Start() {
	tsm.initTasks()

	tsm.cron.Start()
	getLogger().Info("TaskScheduleManager started...")
}

func (tsm *taskScheduleManager) Stop() {
	tsm.cron.Stop()
	getLogger().Info("TaskScheduleManager Stopped...")
}

func (tsm *taskScheduleManager) initTasks() {

	//每隔30秒清理过期sparks
	entryId, err := tsm.cron.AddFunc("@every 30s", func() {
		getLogger().Debug("Clean dying sparks start...")
		err := tsm.sparkService.Clean()
		if err != nil {
			getLogger().Error(err.Error())
		}
		getLogger().Debug("Clean dying sparks end...")
	})
	if err != nil {
		getLogger().Errorf("Add task failed:%s", err.Error())
		panic(err)
	}
	getLogger().Infof("Task added-[entryId:%d]", entryId)
}
