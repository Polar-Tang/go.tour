package main

import "errors"

func (a *analytics) handleEmailBounce(em email) error {
	em.recipient.updateStatus(em.status)
	if em.status == "email_failed" {
		return errors.New("error updating user status: invalid status: email_failed")
	}
	a.track(em.status)
	if em.status == "email_sent" {
		return errors.New("error updating user status: invalid status: email_sent")
	}
	return nil
}
