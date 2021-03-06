package bean

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	UserGenderMale   = "m"
	UserGenderFemale = "f"
)

const (
	UserSexMale   = "m"
	UserSexFemale = "f"
)

type (
	UserBean struct {
		ID       bson.ObjectId `bson:"_id,omitempty"`
		Name     string        `json:"name" xml:"name" form:"name"`
		Password string        `json:"password" xml:"password" form:"password"`
		NickName string
		Gender   string
	}
	User struct {
		ID            int64      `xorm:"'t_user_id'" json:"id,omitempty" gorm:"column:t_user_id;primary_key"`
		UserName      string     `xorm:"'t_user_username'" json:"name,omitempty" form:"name" gorm:"column:t_user_username"`
		Password      string     `xorm:"'t_user_password'" json:"password,omitempty" form:"password" gorm:"column:t_user_password"`
		Birthday      *time.Time `xorm:"'t_user_birthday'" json:"birthday,omitempty" gorm:"column:t_user_birthday"` //生日
		Name          *string    `xorm:"'t_user_name'" json:"name,omitempty" gorm:"column:t_user_name"`             //姓名
		Icon          *string    `xorm:"'t_user_icon'" json:"icon,omitempty" gorm:"column:t_user_icon"`             //头像图片url
		Sex           *string    `xorm:"'t_user_sex'" json:"sex,omitempty" gorm:"column:t_user_sex"`                //性别
		NickName      *string    `xorm:"'t_user_nickname'" json:"nickname,omitempty" gorm:"column:t_user_nickname"`
		LastLoginDate *time.Time `xorm:"'t_user_last_login_date'" json:"last_login_date,omitempty" gorm:"column:t_user_last_login_date"`
		Email         *string    `xorm:"'t_user_email'" json:"email,omitempty" gorm:"column:t_user_email"`
		Mobile        *string    `xorm:"'t_user_mobile'" json:"mobile,omitempty" gorm:"column:t_user_mobile"`
		CreateTime    time.Time  `xorm:"'t_user_create_time'" json:"createTime" bson:"createTime" form:"createTime" gorm:"column:t_user_create_date"`
		UpdateTime    time.Time  `xorm:"'t_user_update_time'" json:"updateTime" bson:"updateTime" form:"updateTime" gorm:"column:t_user_update_time"`
	}

	UserExistBean struct {
		Name  string
		Exist bool
	}
)

func (u User) TableName() string {
	return "t_user"
}
