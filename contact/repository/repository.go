package repository

import (
	"context"

	ent "github.com/fazarrahman/contact-app/contact/entity"
	"github.com/fazarrahman/contact-app/lib/errorHelper"
)

type ContactRepository interface {
	GetContacts(ctx context.Context) ([]*ent.Contacts, *errorHelper.Error)
	GetContactsById(ctx context.Context, id string) (*ent.Contacts, *errorHelper.Error)
	UpsertContact(ctx context.Context, contact *ent.Contacts) *errorHelper.Error
	DeleteContact(ctx context.Context, id string) *errorHelper.Error
}
