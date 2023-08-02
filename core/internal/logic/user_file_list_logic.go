package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/define"
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Id   int `json:"id,optional"`
// Page int `json:"page,optional"`
// Size int `json:"size,optional"`
func (l *UserFileListLogic) UserFileList(req *types.UserFileListreq, userIdentity string) (resp *types.UserFileListres, err error) {
	uf := make([]*types.UserFile, 0)
	resp = new(types.UserFileListres)
	size := req.Size //每页的数量
	if size == 0 {
		size = define.PageSize
	}

	page := req.Page // 第几页
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * size
	err = l.svcCtx.Xorm.Table("user_repository").
		Where("parent_id=? and user_identity=?", req.Id, userIdentity).
		Select(
			"user_repository.id , user_repository.identity , user_repository.repository_identity , user_repository.ext , user_repository.name,"+
				"repository_pool.path , repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? or user_repository.deleted_at is NULL ", time.Time{}.Format(define.DateTime)).
		Limit(size, offset).Find(&uf)
	if err != nil {
		return nil, err
	}

	// 查询总数
	cnt, err := l.svcCtx.Xorm.Where("parent_id=? and user_identity=?", req.Id, userIdentity).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	resp.List = uf
	resp.Count = int(cnt)
	return resp, nil
}
