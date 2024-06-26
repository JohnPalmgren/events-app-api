package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

// Save inserts a new event record into the database with the event details.
// It sets the ID of the Event object upon successful insertion.
// It returns any error encountered during the process.
func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	e.ID = id

	return nil
}

// GetAllEvents retrieves all events from the database.
// It returns a slice of Event objects and any error encountered during the process.
func GetAllEvents() ([]Event, error) {

	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

// GetEventByID retrieves a specific event by its ID from the database.
// It returns the Event object and any error encountered during the process.
func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Update updates the details of an existing event in the database.
// It returns any error encountered during the process.
func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

// Delete removes an existing event from the database by its ID.
// It returns any error encountered during the process.
func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(event.ID)

	return err
}

// Register registers a user for an event by inserting a record into the registrations table.
// It returns any error encountered during the process.
func (event Event) Register(userId int64) error {
	query := ("INSERT INTO registrations (user_id, event_id) VALUES (?, ?)")

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(userId, event.ID)

	return err
}

// DeleteRegistration cancels a user's registration for an event by deleting the corresponding record
// from the registrations table. It returns any error encountered during the process.
func (event Event) DeleteRegistration(userId int64) error {
	query := ("DELETE FROM registrations WHERE user_id = ? AND event_id = ?")

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(userId, event.ID)

	return err
}
