package service

import (
	"context"

	"github.com/fazarrahman/contact-app/contact/entity"
	"github.com/fazarrahman/contact-app/lib/errorHelper"
	"github.com/fazarrahman/contact-app/model"
	"github.com/gin-gonic/gin"
)

func validate(r *model.Contacts, isUpdate bool) *errorHelper.Error {
	if isUpdate && r.Id == "" {
		return errorHelper.BadRequest("Id is required")
	}

	if r.Name == "" {
		return errorHelper.BadRequest("Name is required")
	} else if r.Gender == "" {
		return errorHelper.BadRequest("Gender is required")
	} else if r.Email == "" {
		return errorHelper.BadRequest("Email is required")
	}
	return nil
}

func (s *Svc) GetContacts(ctx *gin.Context) ([]*model.Contacts, *errorHelper.Error) {
	contacts, err := s.ContactRepository.GetContacts(ctx)
	if err != nil {
		return nil, err
	}
	var ctModel []*model.Contacts
	for _, c := range contacts {
		ctModel = append(ctModel, &model.Contacts{
			Id:        c.Id,
			Name:      c.Name,
			Gender:    c.Gender,
			Email:     c.Email,
			Phone:     c.Phone,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
	return ctModel, nil
}

func (s *Svc) InsertContact(ctx context.Context, contact *model.Contacts) *errorHelper.Error {
	err := validate(contact, false)
	if err != nil {
		return err
	}

	err = s.ContactRepository.UpsertContact(ctx, &entity.Contacts{
		Id:     contact.Id,
		Name:   contact.Name,
		Gender: contact.Gender,
		Phone:  contact.Phone,
		Email:  contact.Email,
	})
	if err != nil {
		return err
	}

	return nil
}
