package chat

import (
	"bjm/src/v1/chat/dto"
	"bjm/utils"

	db "bjm/db/benjamit"

	"github.com/gofiber/fiber/v2"

	auth "bjm/auth/jwt"

	"github.com/jsorb84/ssefiber"
)

// @Tags Chat
// @Description Send a message for guest
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body dto.SendRequestModel true "send request"
// @Success 201 {object} dto.SendResponseModel "created"
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 401 {object} utils.ErrorResponseModel "unauthorized"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/chat/user/send [post]
func send(c *fiber.Ctx, sse *ssefiber.FiberSSEApp) error {
	reqModel := &dto.SendRequestModel{}
	resModel := &dto.SendResponseModel{}
	getUuid := auth.DecodeToken(c)["uuid"].(string)
	err := c.BodyParser(reqModel)
	if err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}
	context, contextErr := db.Connect()
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}
	defer db.ConnectClose(context)

	service := &ChatService{context}
	serviceRes := service.Send(reqModel, resModel, getUuid, sse)

	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}

// @Tags Chat
// @Description Send a message (guest)
// @Accept json
// @Produce json
// @Param input body dto.SendForGuestRequestModel true "send request"
// @Success 201 {object} dto.SendForGuestResponseModel "created"
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/chat/send [post]
func sendForGuest(c *fiber.Ctx, sse *ssefiber.FiberSSEApp) error {
	reqModel := &dto.SendForGuestRequestModel{}
	resModel := &dto.SendForGuestResponseModel{}
	err := c.BodyParser(reqModel)
	if err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}
	context, contextErr := db.Connect()
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}
	defer db.ConnectClose(context)

	service := &ChatService{context}
	serviceRes := service.sendForGuest(reqModel, resModel, sse)

	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}

// @Tags Chat
// @Description Event chat message
// @Param channelName path string true "Channel Name"
// @Success 201 {object} dto.SendResponseModel "created"
// @Router /v1/chat/user/events/{channelName} [get]
func eventChat(c *fiber.Ctx, sse *ssefiber.FiberSSEApp) error {
	service := &ChatService{}
	serviceRes := service.EventChat(c, sse)
	return serviceRes
}

// @Tags Chat
// @Description Event chat message for guest
// @Param channelName path string true "Channel Name"
// @Success 201 {object} dto.SendForGuestResponseModel "created"
// @Router /v1/chat/events/guest [get]
func eventChatGuest(c *fiber.Ctx, sse *ssefiber.FiberSSEApp) error {
	service := &ChatService{}
	serviceRes := service.EventChatForGuest(c, sse)
	return serviceRes
}
