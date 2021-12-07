// This file is auto-generated by jwt/internal/cmd/gentoken/main.go. DO NOT EDIT

package openid

import (
	"time"

	"github.com/pkg/errors"
)

// Builder is a convenience wrapper around the New() constructor
// and the Set() methods to assign values to Token claims.
// Users can successively call Claim() on the Builder, and have it
// construct the Token when Build() is called. This alleviates the
// need for the user to check for the return value of every single
// Set() method call.
// Note that each call to Claim() overwrites the value set from the
// previous call.
type Builder struct {
	claims []*ClaimPair
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Claim(name string, value interface{}) *Builder {
	b.claims = append(b.claims, &ClaimPair{Key: name, Value: value})
	return b
}

func (b *Builder) Address(v *AddressClaim) *Builder {
	return b.Claim(AddressKey, v)
}

func (b *Builder) Audience(v []string) *Builder {
	return b.Claim(AudienceKey, v)
}

func (b *Builder) Birthdate(v *BirthdateClaim) *Builder {
	return b.Claim(BirthdateKey, v)
}

func (b *Builder) Email(v string) *Builder {
	return b.Claim(EmailKey, v)
}

func (b *Builder) EmailVerified(v bool) *Builder {
	return b.Claim(EmailVerifiedKey, v)
}

func (b *Builder) Expiration(v time.Time) *Builder {
	return b.Claim(ExpirationKey, v)
}

func (b *Builder) FamilyName(v string) *Builder {
	return b.Claim(FamilyNameKey, v)
}

func (b *Builder) Gender(v string) *Builder {
	return b.Claim(GenderKey, v)
}

func (b *Builder) GivenName(v string) *Builder {
	return b.Claim(GivenNameKey, v)
}

func (b *Builder) IssuedAt(v time.Time) *Builder {
	return b.Claim(IssuedAtKey, v)
}

func (b *Builder) Issuer(v string) *Builder {
	return b.Claim(IssuerKey, v)
}

func (b *Builder) JwtID(v string) *Builder {
	return b.Claim(JwtIDKey, v)
}

func (b *Builder) Locale(v string) *Builder {
	return b.Claim(LocaleKey, v)
}

func (b *Builder) MiddleName(v string) *Builder {
	return b.Claim(MiddleNameKey, v)
}

func (b *Builder) Name(v string) *Builder {
	return b.Claim(NameKey, v)
}

func (b *Builder) Nickname(v string) *Builder {
	return b.Claim(NicknameKey, v)
}

func (b *Builder) NotBefore(v time.Time) *Builder {
	return b.Claim(NotBeforeKey, v)
}

func (b *Builder) PhoneNumber(v string) *Builder {
	return b.Claim(PhoneNumberKey, v)
}

func (b *Builder) PhoneNumberVerified(v bool) *Builder {
	return b.Claim(PhoneNumberVerifiedKey, v)
}

func (b *Builder) Picture(v string) *Builder {
	return b.Claim(PictureKey, v)
}

func (b *Builder) PreferredUsername(v string) *Builder {
	return b.Claim(PreferredUsernameKey, v)
}

func (b *Builder) Profile(v string) *Builder {
	return b.Claim(ProfileKey, v)
}

func (b *Builder) Subject(v string) *Builder {
	return b.Claim(SubjectKey, v)
}

func (b *Builder) UpdatedAt(v time.Time) *Builder {
	return b.Claim(UpdatedAtKey, v)
}

func (b *Builder) Website(v string) *Builder {
	return b.Claim(WebsiteKey, v)
}

func (b *Builder) Zoneinfo(v string) *Builder {
	return b.Claim(ZoneinfoKey, v)
}

// Build creates a new token based on the claims that the builder has received
// so far. If a claim cannot be set, then the method returns a nil Token with
// a en error as a second return value
func (b *Builder) Build() (Token, error) {
	tok := New()
	for _, claim := range b.claims {
		if err := tok.Set(claim.Key.(string), claim.Value); err != nil {
			return nil, errors.Wrapf(err, `failed to set claim %q`, claim.Key.(string))
		}
	}
	return tok, nil
}
