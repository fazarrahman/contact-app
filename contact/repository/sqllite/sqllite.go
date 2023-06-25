package sqllite

import (
	"context"
	"database/sql"
	"time"

	ent "github.com/fazarrahman/contact-app/contact/entity"
	"github.com/fazarrahman/contact-app/lib/errorHelper"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Mysqldb struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Mysqldb {
	return &Mysqldb{db: db}
}

func (m *Mysqldb) GetContacts(ctx context.Context, limit, offset int) ([]*ent.Contacts, *errorHelper.Error) {
	var contacts []*ent.Contacts
	err := m.db.SelectContext(ctx, &contacts, `SELECT id, name, gender, phone, email, created_at, updated_at
	FROM contacts LIMIT ? OFFSET ?`, limit, offset)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errorHelper.InternalServerError("Error get Contacts : " + err.Error())
	}
	return contacts, nil
}

func (m *Mysqldb) GetContactsById(ctx context.Context, id string) (*ent.Contacts, *errorHelper.Error) {
	var contacts ent.Contacts
	err := m.db.GetContext(ctx, &contacts, `SELECT id, name, gender, phone, email, created_at, updated_at
	FROM contacts 
	WHERE id = ?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errorHelper.InternalServerError("Error get contacts by id : " + err.Error())
	}
	return &contacts, nil
}

func (m *Mysqldb) UpsertContact(ctx context.Context, contact *ent.Contacts) *errorHelper.Error {
	now := time.Now()
	contact.CreatedAt = now
	contact.UpdatedAt = now
	if contact.Id == "" {
		contact.Id = uuid.New().String()
	}
	_, err := m.db.NamedExecContext(ctx, `INSERT OR REPLACE INTO contacts
	(id, name, gender, phone, email, created_at, updated_at)
	VALUES(:id, :name, :gender, :phone, :email, 
		COALESCE((SELECT created_at FROM contacts WHERE id = :id), :created_at), 
		:updated_at)`, contact)

	if err != nil {
		return errorHelper.InternalServerError("Error Upsert Contact : " + err.Error())
	}

	return nil
}

func (m *Mysqldb) DeleteContact(ctx context.Context, id string) *errorHelper.Error {
	_, err := m.db.ExecContext(ctx, `DELETE FROM contacts where id = ?`, id)
	if err != nil {
		return errorHelper.InternalServerError("Error Delete Contact : " + err.Error())
	}

	return nil
}
