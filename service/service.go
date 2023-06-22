package service

import (
	"context"

	contactRepo "github.com/fazarrahman/contact-app/contact/repository"
	"github.com/fazarrahman/contact-app/lib/errorHelper"
	"github.com/fazarrahman/contact-app/model"
)

type Svc struct {
	ContactRepository contactRepo.ContactRepository
}

// New ...
func New(_contactRepo contactRepo.ContactRepository) *Svc {
	return &Svc{ContactRepository: _contactRepo}
}

// Service ...
type ServiceInterface interface {
	GetContacts(ctx context.Context) ([]*model.Contacts, *errorHelper.Error)
	InsertContact(ctx context.Context, contact *model.Contacts) *errorHelper.Error
	UpdateContact(ctx context.Context, contact *model.Contacts) *errorHelper.Error
	DeleteContact(ctx context.Context, id string) *errorHelper.Error
}
