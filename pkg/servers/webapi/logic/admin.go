package logic

import (
	"context"
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	modelcode "github.com/CloudSilk/pkg/model"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"gorm.io/gorm"
)

func TryLogin(req *proto.LoginRequest, resp *proto.ServiceResponse) {
	var user *usercenter.UserProfile
	var token string
	if req.CardNo == "" {
		if req.Name == "" {
			resp.Code = modelcode.BadRequest
			resp.Message = "账号不能为空"
			return
		}
		if req.Password == "" {
			resp.Code = modelcode.BadRequest
			resp.Message = "密码不能为空"
			return
		}

		//根据帐号获取用户
		_user, _ := clients.UserClient.LoginByStaffNo(context.Background(), &usercenter.LoginByStaffNoRequest{UserName: req.Name, Password: req.Password})
		if _user.Code == modelcode.InternalServerError {
			resp.Code = modelcode.InternalServerError
			resp.Message = _user.Message
			return
		}
		if _user.Code == modelcode.UserIsNotExist || _user.Code == modelcode.UserNameOrPasswordIsWrong {
			resp.Code = modelcode.UserIsNotExist
			resp.Message = "用户或密码错误"
			return
		}
		if _user.Code == modelcode.UserDisabled {
			resp.Code = modelcode.UserDisabled
			resp.Message = "用户已停用"
			return
		}

		user = _user.User
		token = _user.Data
	} else {
		//根据卡号获取用户
		_user, _ := clients.UserClient.LoginByStaffNo(context.Background(), &usercenter.LoginByStaffNoRequest{StaffNo: req.CardNo})
		if _user.Code == modelcode.InternalServerError {
			resp.Code = modelcode.InternalServerError
			resp.Message = _user.Message
			return
		}
		if _user.Code == modelcode.UserIsNotExist {
			resp.Code = modelcode.UserIsNotExist
			resp.Message = "卡号不存在，请联系产线主管"
			return
		}
		if _user.Code == modelcode.UserDisabled {
			resp.Code = modelcode.UserDisabled
			resp.Message = "用户已停用"
			return
		}

		user = _user.User
		token = _user.Data
	}

	if req.ProductionStation == "" {
		resp.Code = modelcode.BadRequest
		resp.Message = "ProductionStation不能为空"
		return
	} else {
		if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
			productionStation := &model.ProductionStation{}
			if err := tx.Where(model.ProductionStation{Code: req.ProductionStation}).First(productionStation).Error; err == gorm.ErrRecordNotFound {
				resp.Message = "无效的工位编号"
				return err
			} else if err != nil {
				return err
			}
			productionStation.CurrentUserID = &user.Id

			productionStationSignup := &model.ProductionStationSignup{}
			if err := tx.Where(model.ProductionStationSignup{
				ProductionStationID: productionStation.ID,
				LoginUserID:         user.Id,
			}).Where("logout_time IS NULL").First(productionStationSignup).Error; err != nil && err != gorm.ErrRecordNotFound {
				return err
			}

			now := time.Now()
			if productionStationSignup.ID == "" {
				productionStationSignup = &model.ProductionStationSignup{
					LoginUserID:         user.Id,
					ProductionStationID: productionStation.ID,
					LoginTime:           now,
				}
				if err := tx.Create(productionStationSignup).Error; err != nil {
					return err
				}
			}

			productionStationSignup.LastHeartbeatTime = sql.NullTime{Time: now, Valid: true}
			productionStationSignup.LoginTime = now
			productionStationSignup.Duration = int32(now.Sub(productionStationSignup.LoginTime).Minutes())

			if err := tx.Save(productionStationSignup).Error; err != nil {
				return err
			}

			if err := tx.Save(productionStation).Error; err != nil {
				return err
			}

			return nil
		}); err != nil {
			if err != gorm.ErrRecordNotFound {
				resp.Message = err.Error()
			}
			return
		}
	}

	resp.Message = token
	resp.Data = &proto.Data{
		Id:          user.Id,
		Name:        user.UserName,
		ChineseName: user.ChineseName,
		EnglishName: user.EnglishName,
		Photo:       user.Mobile,
		StaffNo:     user.StaffNo,
	}
}

func TryLogout(req *proto.LogoutRequest, resp *proto.CommonResponse) {
	if req.Name == "" {
		resp.Code = modelcode.BadRequest
		resp.Message = "Name不能为空"
		return
	}
	if req.ProductionStation == "" {
		resp.Code = modelcode.BadRequest
		resp.Message = "ProductionStation不能为空"
		return
	}

	_productionStation, _ := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
	if _productionStation.Message == gorm.ErrRecordNotFound.Error() {
		resp.Code = modelcode.BadRequest
		resp.Message = "无效的工位编号"
		return
	}
	if _productionStation.Code != modelcode.Success {
		resp.Code = _productionStation.Code
		resp.Message = _productionStation.Message
		return
	}
	productionStation := _productionStation.Data

	//根据帐号注销用户
	_user, _ := clients.UserClient.LogoutByUserName(context.Background(), &usercenter.LogoutByUserNameRequest{UserName: req.Name})
	if _user.Code != modelcode.Success {
		resp.Code = proto.Code(_user.Code)
		resp.Message = _user.Message
		return
	}

	userID := _user.Message
	if userID == "" {
		resp.Code = modelcode.UserIsNotExist
		resp.Message = "无效的工位或账号"
		return
	}
	if productionStation.CurrentUserID != userID {
		resp.Code = modelcode.BadRequest
		resp.Message = "登录信息错误，用于账号与工位登录当前账号不符"
		return
	}

	_productionStationSignup, _ := clients.ProductionStationSignupClient.Get(context.Background(), &proto.GetProductionStationSignupRequest{
		ProductionStationID: productionStation.Id,
		LoginUserID:         productionStation.CurrentUserID,
		HasLogoutTime:       false,
	})
	if _productionStationSignup.Code == modelcode.InternalServerError && _productionStationSignup.Message != gorm.ErrRecordNotFound.Error() {
		resp.Code = _productionStationSignup.Code
		resp.Message = _productionStationSignup.Message
		return
	}

	productionStationSignup := _productionStationSignup.Data
	productionStationSignup.LastHeartbeatTime = time.Now().Format("2006-01-02 15:04:05")
	productionStationSignup.LogoutTime = time.Now().Format("2006-01-02 15:04:05")
	productionStationSignup.Duration = 0

	productionStation.CurrentUserID = ""
	if _resp, _ := clients.ProductionStationClient.Update(context.Background(), productionStation); _resp.Code != modelcode.Success {
		resp.Code = _resp.Code
		resp.Message = _resp.Message
		return
	}
	if _resp, _ := clients.ProductionStationSignupClient.Update(context.Background(), productionStationSignup); _resp.Code != modelcode.Success {
		resp.Code = _resp.Code
		resp.Message = _resp.Message
		return
	}
}
