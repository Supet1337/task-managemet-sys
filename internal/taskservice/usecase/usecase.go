package eventLogic

import models "golang-pkg/internal/taskservice"

// создание ивента
func CreateEvent(task models.Task) error {
	// check, user can't be organizer and member
	err := event.ValidValue()
	if err.Message != nil {
		return err
	} else if err = checkNoteAtTime(event.OrganizersId, event.TimeStart, event.TimeEnd); err.Message != nil {
		return err
	} else {
		err = repository.CreateNode(event)
	}
	return err
}
