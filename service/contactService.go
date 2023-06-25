package service

import (
	"github.com/fazarrahman/contact-app/contact/entity"
	"github.com/fazarrahman/contact-app/lib/errorHelper"
	"github.com/fazarrahman/contact-app/model"
	"github.com/gin-gonic/gin"
)

func (s *Svc) validate(c *gin.Context, r *model.Contacts, isUpdate bool) *errorHelper.Error {
	if isUpdate {
		if r.Id == "" {
			return errorHelper.BadRequest("Id is required")
		}
		ct, err := s.ContactRepository.GetContactsById(c, r.Id)
		if err != nil {
			return errorHelper.InternalServerError(err.Message)
		}
		if ct == nil {
			return errorHelper.BadRequest("Invalid Id")
		}
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

func (s *Svc) GetContacts(ctx *gin.Context, limit, offset int) ([]*model.Contacts, *errorHelper.Error) {
	contacts, err := s.ContactRepository.GetContacts(ctx, limit, offset)
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

func (s *Svc) InsertContact(ctx *gin.Context, contact *model.Contacts) *errorHelper.Error {
	err := s.validate(ctx, contact, false)
	if err != nil {
		return err
	}

	err = s.ContactRepository.UpsertContact(ctx, &entity.Contacts{
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

func (s *Svc) UpdateContact(ctx *gin.Context, contact *model.Contacts) *errorHelper.Error {
	err := s.validate(ctx, contact, true)
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

func (s *Svc) DeleteContact(ctx *gin.Context, id string) *errorHelper.Error {
	err := s.ContactRepository.DeleteContact(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
