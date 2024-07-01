package models

import (
	"database/sql"
	"time"

	"log"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"dateTime"`
	UserID      int64     `json:"userID"`
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
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

func GetAllEvents() ([]Event, error) {
	var events []Event
	query := "SELECT * FROM events"

	log.Println("Executing query:", query)

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("Error querying all events: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		var dateTimeStr string
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserID)
		if err != nil {
			log.Printf("Error scanning event row: %v", err)
			return nil, err
		}

		event.DateTime, err = time.Parse("2006-01-02 15:04:05-07:00", dateTimeStr)
		if err != nil {
			log.Printf("Error parsing dateTime: %v", err)
			return nil, err
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("Fetched %d events", len(events))
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	var dateTimeStr string
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Event not found with id: %d", id)
			return nil, err
		}
		log.Printf("Error scanning event row by ID: %v", err)
		return nil, err
	}

	event.DateTime, err = time.Parse("2006-01-02 15:04:05-07:00", dateTimeStr)
	if err != nil {
		log.Printf("Error parsing dateTime: %v", err)
		return nil, err
	}

	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}
