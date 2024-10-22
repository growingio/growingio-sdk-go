package growingio

import (
	"errors"
	"os"
	"strings"

	core "github.com/growingio/growingio-sdk-go/internal/core"
	logger "github.com/growingio/growingio-sdk-go/internal/logger"
	"gopkg.in/yaml.v3"
)

var lock = make(chan struct{}, 1)

func InitAnalytics(config *Config) error {
	lock <- struct{}{}
	defer func() { <-lock }()

	if core.InitializedSuccessfully {
		err := errors.New("initialization failed, already initialized")
		logger.Error(err, "Do not initialize GrowingAnalytics repeatedly")
		return err
	}

	logger.SetLogLevel(int(config.SdkConfig.LogLevel))

	if config.SdkConfig.AccountId == "" {
		err := errors.New("initialization failed, AccountId is empty")
		logger.Error(err, "Please enter your AccountId")
		return err
	}
	if config.SdkConfig.DataSourcesId == "" {
		err := errors.New("initialization failed, DataSourcesId is empty")
		logger.Error(err, "Please enter your DataSourcesId")
		return err
	}
	core.AccountId = config.SdkConfig.AccountId
	core.DataSourcesId = config.SdkConfig.DataSourcesId

	if config.HttpConfig.ServerHost == "" {
		err := errors.New("initialization failed, ServerHost is empty")
		logger.Error(err, "Please enter your ServerHost")
		return err
	}
	serverHost := strings.TrimSuffix(config.HttpConfig.ServerHost, "/")

	core.Urls = core.RequestUrl{
		Collect: serverHost + "/v3/projects/" + config.SdkConfig.AccountId + "/collect",
		Item:    serverHost + "/projects/" + config.SdkConfig.AccountId + "/collect/item",
	}
	core.RequestTimeout = int(config.HttpConfig.RequestTimeout)

	if config.BatchConfig.Enable {
		core.BatchEnable = config.BatchConfig.Enable
		core.MaxSize = config.BatchConfig.MaxSize
		core.FlushAfter = config.BatchConfig.FlushAfter
		core.RoutineCount = config.BatchConfig.RoutineCount
		core.MaxCacheSize = config.BatchConfig.MaxCacheSize
		core.InitBatch()
	}

	core.InitializedSuccessfully = true
	logger.Info("Thank you very much for using GrowingIO. We will do our best to provide you with the best service.", "sdkVersion", core.SdkVersion)
	logger.Debug(config.String())
	return nil
}

func InitAnalyticsByConfigFile(file string) error {
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		logger.Error(err, "Read File failed", "filePath", file)
		return err
	}
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		logger.Error(err, "yaml unmarshal failed", "filePath", file)
		return err
	}
	return InitAnalytics(config)
}

func TrackCustomEvent(builder *CustomEventBuilder) error {
	if !core.InitializedSuccessfully {
		err := errors.New("TrackCustomEvent failed, GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}

	if len(builder.EventName) == 0 {
		err := errors.New("TrackCustomEvent failed, EventName is empty")
		logger.Error(err, "Please enter eventName for customEvent")
		return err
	}

	if len(builder.AnonymousId)+len(builder.LoginUserId) == 0 {
		err := errors.New("TrackCustomEvent failed, AnonymousId and LoginUserId are empty")
		logger.Error(err, "Both AnonymousId and LoginUserId are empty. Please enter at least one of them")
		return err
	}

	b := core.EventBuilder{
		EventName:    builder.EventName,
		EventTime:    builder.EventTime,
		AnonymousId:  builder.AnonymousId,
		LoginUserId:  builder.LoginUserId,
		LoginUserKey: builder.LoginUserKey,
		Attributes:   builder.Attributes,
	}
	b.BuildCustomEvent()
	return nil
}

func TrackUser(builder *UserEventBuilder) error {
	if !core.InitializedSuccessfully {
		err := errors.New("TrackUser failed, GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}

	if len(builder.LoginUserId) == 0 {
		err := errors.New("TrackUser failed, LoginUserId is empty")
		logger.Error(err, "Please enter loginUserId for user")
		return err
	}

	b := core.EventBuilder{
		EventTime:    builder.EventTime,
		AnonymousId:  builder.AnonymousId,
		LoginUserId:  builder.LoginUserId,
		LoginUserKey: builder.LoginUserKey,
		Attributes:   builder.Attributes,
	}
	b.BuildUserLoginEvent()
	return nil
}

func SubmitItem(builder *ItemBuilder) error {
	if !core.InitializedSuccessfully {
		err := errors.New("SubmitItem failed, GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}

	if len(builder.ItemId) == 0 {
		err := errors.New("SubmitItem failed, ItemId is empty")
		logger.Error(err, "Please enter itemId for item")
		return err
	}

	if len(builder.ItemKey) == 0 {
		err := errors.New("SubmitItem failed, ItemKey is empty")
		logger.Error(err, "Please enter itemKey for item")
		return err
	}

	b := core.EventBuilder{
		ItemId:     builder.ItemId,
		ItemKey:    builder.ItemKey,
		Attributes: builder.Attributes,
	}
	b.BuildItemEvent()
	return nil
}
