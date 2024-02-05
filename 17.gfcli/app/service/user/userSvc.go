package user

import (
	"errors"
	"gfcli/app/model/user"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

// 请求参数
type Request struct {
	user.Entity
}

// 通过id获取实体
func GetById(id int64) (*user.Entity, error) {
	if id <= 0 {
		glog.Error(" get id error")
		return nil, errors.New("参数不合法")
	}

	return user.Model.One(" id = ?", id)
}

// 删除实体
func Delete(id int64) (int64, error) {
	if id <= 0 {
		glog.Error("delete id error")
		return 0, errors.New("参数不合法")
	}

	// 获取删除对象
	r, err1 := user.Model.Delete(" id = ?", id)
	if err1 != nil {
		return 0, err1
	}

	return r.RowsAffected()
}

// 保存实体
func Save(request *Request) (int64, error) {
	entity := (*user.Entity)(nil)
	err := gconv.StructDeep(request.Entity, &entity)
	if err != nil {
		return 0, errors.New("数据错误")
	}

	// 判断新增还是修改
	if entity.Id <= 0 {

		r, err := user.Model.Insert(entity)
		if err != nil {
			return 0, err
		}

		return r.RowsAffected()
	} else {
		r, err := user.Model.OmitEmpty().Where(" id = ?", entity.Id).Update(entity)
		if err != nil {
			return 0, err
		}

		return r.RowsAffected()
	}
}

// 分页查询
func Page(page, limit int) ([]*user.Entity, int, error) {
	if page <= 0 || limit <= 0 {
		glog.Error("page param error", form.Page, form.Rows)
		return nil, 0, nil
	}

	num, err := user.Model.As("t").FindCount()

	if err != nil {
		glog.Error("page count error", err)
		return nil, 0, err
	}

	dbModel, err := user.Model.As("t").Page(page, limit).All()
	if err != nil {
		glog.Error("page list error", err)
		return nil, 0, err
	}

	return dbModel, num, nil
}
