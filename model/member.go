package model

import (
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
)

type Member struct {
	MemberId int64  `db:"member_id" json:"member_id"` // チャットサービス内のID
	Service  string `db:"service" json:"service"`     // チャットサービス名
	AppName  string `db:"app" json:"app"`             // 選択中のチャットアプリケーション
}

const DefaultApp = "WordChain"

func NewMember(member_id int64, service string) *Member {
	return &Member{
		MemberId: member_id,
		Service:  service,
		AppName:  DefaultApp,
	}
}

func (m *Member) SaveMember(tx *dbr.Tx) error {

	_, err := tx.InsertInto("member").
		Columns("member_id", "service", "app").
		Record(m).
		Exec()

	if err != nil {
		logrus.Error("Error")
	}

	return err
}

func (m *Member) LoadMember(tx *dbr.Tx, number int64) error {

	_, err := tx.Select("*").
		From("member").
		Where("id = ?", number).
		Load(m)

	if err != nil {
		logrus.Error("Error")
	}
	return err
}

type Members []Member

func (m *Members) LoadMembers(tx *dbr.Tx) error {

	_, err := tx.Select("*").
		From("member").
		Load(m)

	if err != nil {
		logrus.Error("Error")
	}
	return err
}
