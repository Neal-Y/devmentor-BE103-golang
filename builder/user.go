package builder

import "shopping-cart/model/database"

type UserBuilder struct {
	user *database.User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: &database.User{}}
}

func (b *UserBuilder) WithLineID(lineID string) *UserBuilder {
	b.user.LineID = lineID
	return b
}

func (b *UserBuilder) WithDisplayName(displayName string) *UserBuilder {
	b.user.DisplayName = displayName
	return b
}

func (b *UserBuilder) WithPhone(phone string) *UserBuilder {
	b.user.Phone = phone
	return b
}

func (b *UserBuilder) WithIsMember(isMember bool) *UserBuilder {
	b.user.IsMember = isMember
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithLineToken(lineToken string) *UserBuilder {
	b.user.LineToken = lineToken
	return b
}

func (b *UserBuilder) Build() *database.User {
	return b.user
}
