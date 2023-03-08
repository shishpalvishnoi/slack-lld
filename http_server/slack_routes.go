package http_server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slack-application/models"
	"slack-application/service"
)

func registerUser(context *gin.Context) {
	var regUserReq models.RegUserReq
	err := context.BindJSON(&regUserReq)
	if err != nil {
		logger.Error("registerUser error: ", zap.Error(err))
	}
	logger.Info("register user request received: ", zap.Any("body", regUserReq))

	// register user
	s := service.Service{}
	err = s.RegisterUser(context, regUserReq)
	if err != nil {
		logger.Error("registerUser error: ", zap.Error(err))
		// return error here
		res := models.ResponseFailure(nil, "failed to register user")
		context.JSON(400, res)
		return
	}
	res := models.ResponseSuccess(nil)
	context.JSON(200, res)
}

func createWorkspace(context *gin.Context) {
	var createWorkspaceReq models.CreateWorkspaceReq
	err := context.BindJSON(&createWorkspaceReq)
	if err != nil {
		logger.Error("createWorkspace error: ", zap.Error(err))
	}
	logger.Info("create workspace request received: ", zap.Any("body", createWorkspaceReq))

	// create workspace
	s := service.Service{}
	err = s.CreateWorkspace(context, createWorkspaceReq)
	if err != nil {
		logger.Error("createWorkspace error: ", zap.Error(err))
		// return error here
		res := models.ResponseFailure(nil, "failed to create workspace")
		context.JSON(400, res)
		return
	}
	res := models.ResponseSuccess(nil)
	context.JSON(200, res)
}

func createChannelInWorkspace(context *gin.Context) {
	var createChannelReq models.CreateChannelReq
	err := context.BindJSON(&createChannelReq)
	if err != nil {
		logger.Error("createChannel error: ", zap.Error(err))
	}
	logger.Info("create channel in workspace request received: ", zap.Any("body", createChannelReq))
	// register user
	s := service.Service{}
	err = s.CreateChannelInWorkspace(context, createChannelReq)
	if err != nil {
		logger.Error("createChannelInWorkspace error: ", zap.Error(err))
		// return error here
		res := models.ResponseFailure(nil, "failed to create channel")
		context.JSON(400, res)
		return
	}
	res := models.ResponseSuccess(nil)
	context.JSON(200, res)
}

func addUserToWorkspace(context *gin.Context) {
	var addUserToWorkspaceReq models.AddUserToWorkspaceReq
	err := context.BindJSON(&addUserToWorkspaceReq)
	if err != nil {
		logger.Error("addUserToWorkspace error: ", zap.Error(err))
	}
	logger.Info("add user to workspace request received: ", zap.Any("body", addUserToWorkspaceReq))

	// register user
	s := service.Service{}
	err = s.AddUserToWorkspace(context, addUserToWorkspaceReq)
	if err != nil {
		logger.Error("addUserToWorkspace error: ", zap.Error(err))
		// return error here
		res := models.ResponseFailure(nil, "failed to add user to workspace")
		context.JSON(400, res)
		return
	}

	res := models.ResponseSuccess(nil)
	context.JSON(200, res)
}

func addUserToChannel(context *gin.Context) {
	var addUserToChannelReq models.AddUserToChannelReq
	err := context.BindJSON(&addUserToChannelReq)
	if err != nil {
		logger.Error("addUserToChannelReq error: ", zap.Error(err))
	}
	logger.Info("add user to channel request received: ", zap.Any("body", addUserToChannelReq))

	// add user to channel
	s := service.Service{}
	err = s.AddUserToChannel(context, addUserToChannelReq)
	if err != nil {
		logger.Error("addUserToChannel error: ", zap.Error(err))
		// return error here
		res := models.ResponseFailure(nil, "failed to add user to channel")
		context.JSON(400, res)
		return
	}

	res := models.ResponseSuccess(nil)
	context.JSON(200, res)
}

func sendMessageToChannel(context *gin.Context) {
	var sendMessageToChannelReq models.SendMessageToChReq
	err := context.BindJSON(&sendMessageToChannelReq)
	if err != nil {
		logger.Error("sendMessageToChannelReq error: ", zap.Error(err))
	}
	logger.Info("send message to channel request received: ", zap.Any("body", sendMessageToChannelReq))

	// send message to channel
	s := service.Service{}
	err = s.SendMessageToChannel(context, sendMessageToChannelReq)
	if err != nil {
		logger.Error("sendMessageToChannel error: ", zap.Error(err))
		// return error here
		res := models.ResponseFailure(nil, "failed to send message to channel")
		context.JSON(400, res)
		return
	}

	res := models.ResponseSuccess(nil)
	context.JSON(200, res)
}

func sendMessageToUser(context *gin.Context) {
	var sendMessageToUserReq models.SendMessageToUserReq
	err := context.BindJSON(&sendMessageToUserReq)
	if err != nil {
		logger.Error("sendMessageToUserReq error: ", zap.Error(err))
	}

	logger.Info("send message to user request received: ", zap.Any("body", sendMessageToUserReq))

	// send message to user
	s := service.Service{}
	err = s.SendMessageToUser(context, sendMessageToUserReq)
	if err != nil {
		logger.Error("sendMessageToUser error: ", zap.Error(err))
		// return error here
		res := models.ResponseFailure(nil, "failed to send message to user")
		context.JSON(400, res)
		return
	}

	res := models.ResponseSuccess(nil)
	context.JSON(200, res)
}
