package models

type Room struct {
	Id           int    `json:"id"`
	Topic        string `json:"topic"`
	Participants []*Participant
}

func (r *Room) RemoveParticipant(i *Participant) {
	s := r.Participants

	for ind, p := range s {
		if p == i {
			s[ind] = s[len(s)-1]
			r.Participants = s[:len(s)-1]
		}
	}
}
