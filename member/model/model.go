package model

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/micro/micro/v3/service/logger"
	"gorm.io/gorm"
)

var JwtKey = []byte("my_secret_key")

type jwtClaim struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type Member struct {
	gorm.Model
	Avatar   string `json:"avatar" gorm:"type:varchar(100);"`
	Nickname string `json:"nickname" gorm:"type:varchar(100);"`
	Username string `json:"username" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password" gorm:"type:varchar(256)"`
}

// 加密密码
func (m Member) EncodePassword() string {
	password := md5.Sum([]byte(m.Password))
	logger.Debugf("加密密码 password=%s", m.Password, "encode=%x", password)
	return fmt.Sprintf("%x", password)
}

// 生成token
func (m Member) MakeToken() (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaim{
		UserId:   m.ID,
		Username: m.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}).SignedString(JwtKey)
}

func (m *Member) BeforeSave(tx *gorm.DB) (err error) {
	if m.Password != "" {
		m.Password = m.EncodePassword()
	}
	return nil
}

func (m *Member) BeforeCreate(tx *gorm.DB) (err error) {
	if m.Password != "" {
		m.Password = m.EncodePassword()
	}
	return nil
}

func (m *Member) BeforeUpdate(tx *gorm.DB) (err error) {
	if m.Password != "" {
		m.Password = m.EncodePassword()
	}
	return nil
}
