// Copyright (c) 2015 Spinpunch, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

const (
	TEAM_OPEN   = "O"
	TEAM_INVITE = "I"
)

type Team struct {
	Id             string `json:"id"`
	CreateAt       int64  `json:"create_at"`
	UpdateAt       int64  `json:"update_at"`
	DeleteAt       int64  `json:"delete_at"`
	Name           string `json:"name"`
	URLId          string `json:"urlid"`
	Email          string `json:"email"`
	Type           string `json:"type"`
	CompanyName    string `json:"company_name"`
	AllowedDomains string `json:"allowed_domains"`
	AllowValet     bool   `json:"allow_valet"`
}

type Invites struct {
	Invites []map[string]string `json:"invites"`
}

func InvitesFromJson(data io.Reader) *Invites {
	decoder := json.NewDecoder(data)
	var o Invites
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *Invites) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (o *Team) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func TeamFromJson(data io.Reader) *Team {
	decoder := json.NewDecoder(data)
	var o Team
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *Team) Etag() string {
	return Etag(o.Id, o.UpdateAt)
}

func (o *Team) IsValid() *AppError {

	if len(o.Id) != 26 {
		return NewAppError("Team.IsValid", "Invalid Id", "")
	}

	if o.CreateAt == 0 {
		return NewAppError("Team.IsValid", "Create at must be a valid time", "id="+o.Id)
	}

	if o.UpdateAt == 0 {
		return NewAppError("Team.IsValid", "Update at must be a valid time", "id="+o.Id)
	}

	if len(o.Email) > 128 {
		return NewAppError("Team.IsValid", "Invalid email", "id="+o.Id)
	}

	if !IsValidEmail(o.Email) {
		return NewAppError("Team.IsValid", "Invalid email", "id="+o.Id)
	}

	if len(o.Name) > 64 {
		return NewAppError("Team.IsValid", "Invalid name", "id="+o.Id)
	}

	if len(o.URLId) > 64 {
		return NewAppError("Team.IsValid", "Invalid URL Identifier", "id="+o.Id)
	}

	if IsReservedURLId(o.URLId) {
		return NewAppError("Team.IsValid", "This URL is unavailable. Please try another.", "id="+o.Id)
	}

	if !IsValidURLId(o.URLId) {
		return NewAppError("Team.IsValid", "URLIdentifier must be 4 or more lowercase alphanumeric characters", "id="+o.Id)
	}

	if !(o.Type == TEAM_OPEN || o.Type == TEAM_INVITE) {
		return NewAppError("Team.IsValid", "Invalid type", "id="+o.Id)
	}

	if len(o.CompanyName) > 64 {
		return NewAppError("Team.IsValid", "Invalid company name", "id="+o.Id)
	}

	if len(o.AllowedDomains) > 500 {
		return NewAppError("Team.IsValid", "Invalid allowed domains", "id="+o.Id)
	}

	return nil
}

func (o *Team) PreSave() {
	if o.Id == "" {
		o.Id = NewId()
	}

	o.CreateAt = GetMillis()
	o.UpdateAt = o.CreateAt
}

func (o *Team) PreUpdate() {
	o.UpdateAt = GetMillis()
}
