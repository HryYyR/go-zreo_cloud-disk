// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Logger},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/login",
					Handler: UserLoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/detail/:identity",
					Handler: UserDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/register",
					Handler: UserRegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/mail/code/send/register",
					Handler: MailCodeSendRegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/share/basic/detail/:identity",
					Handler: ShareBasicDetailHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/file/upload",
					Handler: FileUploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/repository/save",
					Handler: UserRepositorySaveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/list",
					Handler: UserFileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/name/update",
					Handler: UserFileNameUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/folder/create/userfolder",
					Handler: UserFolderCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user/folder/delete/userfolder/:identity",
					Handler: UserFolderDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/folder/move/userfolder",
					Handler: UserFolderMoveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share/basic/create",
					Handler: ShareCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share/basic/save",
					Handler: ShareBasicSaveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/refresh/token",
					Handler: RefreshTokenHandler(serverCtx),
				},
			}...,
		),
	)
}
