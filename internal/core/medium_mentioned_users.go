package core

func (mu MentionedUsers) getUserData(id string) MentionedUser {
	for _, u := range mu {
		if u.UserID == id {
			return u
		}
	}
	return MentionedUser{}
}
