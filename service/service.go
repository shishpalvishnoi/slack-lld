package service

import (
	"context"
	"slack-application/database"
	"slack-application/models"
	"time"
)

type Service struct {
}

func (s *Service) RegisterUser(ctx context.Context, registerUserReq models.RegUserReq) error {
	_, err := database.Client.Users.Create().
		SetName(registerUserReq.Name).
		SetEmail(registerUserReq.Email).
		SetPhoneNumber(registerUserReq.PhoneNumber).
		SetUserType(registerUserReq.UserType).
		SetVersionID(0).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	return err
}

func (s *Service) CreateWorkspace(ctx context.Context, createWorkspaceReq models.CreateWorkspaceReq) error {
	_, err := database.Client.Workspace.Create().
		SetName(createWorkspaceReq.Name).
		SetDescription(createWorkspaceReq.Description).
		SetCreatedAt(time.Now()).
		//SetCreatedBy(createWorkspaceReq.)
		Save(context.Background())
	return err
}

func (s *Service) CreateChannelInWorkspace(ctx context.Context, createChannelReq models.CreateChannelReq) error {
	// todo: check workspace exists and user creating channel has permission
	_, err := database.Client.Channel.Create().
		SetName(createChannelReq.Name).
		SetWorkspaceID(createChannelReq.WorkspaceId).
		SetCreatedBy(createChannelReq.CreatedBy).
		Save(context.Background())
	return err
}

func (s *Service) AddUserToWorkspace(ctx context.Context, addUserToWorkspaceReq models.AddUserToWorkspaceReq) error {
	_, err := database.Client.UserWorkspace.Create().
		SetWorkspaceID(addUserToWorkspaceReq.WorkSpaceId).
		SetUserID(addUserToWorkspaceReq.UserId).
		SetRoleType(addUserToWorkspaceReq.RoleType).
		Save(context.Background())
	return err
}

func (s *Service) AddUserToChannel(ctx context.Context, addUserToChannelReq models.AddUserToChannelReq) error {
	_, err := database.Client.UserChannel.Create().
		SetUserID(addUserToChannelReq.UserId).
		SetChannelID(addUserToChannelReq.ChannelId).
		SetRoleType(addUserToChannelReq.RoleType).
		Save(context.Background())
	return err
}

func (s *Service) SendMessageToChannel(ctx context.Context, sendMessageReq models.SendMessageToChReq) error {
	_, err := database.Client.Message.Create().
		SetUserID(sendMessageReq.UserId).
		SetWorkspaceID(sendMessageReq.WorkspaceId).
		SetAttachment(sendMessageReq.Attachment).
		SetData(sendMessageReq.Content).
		SetCreatedAt(time.Now()).
		Save(context.Background())

	// add messageId in channel message table

	return err
}

func (s *Service) SendMessageToUser(ctx context.Context, sendMessageReq models.SendMessageToUserReq) error {
	return nil
}
