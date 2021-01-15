// This file is auto-generated by jwt/internal/cmd/gentoken/main.go. DO NOT EDIT

package openid

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/lestrrat-go/iter/mapiter"
	"github.com/lestrrat-go/jwx/internal/iter"
	"github.com/lestrrat-go/jwx/internal/json"
	"github.com/lestrrat-go/jwx/jwt/internal/types"
	"github.com/pkg/errors"
)

const (
	AudienceKey            = "aud"
	ExpirationKey          = "exp"
	IssuedAtKey            = "iat"
	IssuerKey              = "iss"
	JwtIDKey               = "jti"
	NotBeforeKey           = "nbf"
	SubjectKey             = "sub"
	NameKey                = "name"
	GivenNameKey           = "given_name"
	MiddleNameKey          = "middle_name"
	FamilyNameKey          = "family_name"
	NicknameKey            = "nickname"
	PreferredUsernameKey   = "preferred_username"
	ProfileKey             = "profile"
	PictureKey             = "picture"
	WebsiteKey             = "website"
	EmailKey               = "email"
	EmailVerifiedKey       = "email_verified"
	GenderKey              = "gender"
	BirthdateKey           = "birthdate"
	ZoneinfoKey            = "zoneinfo"
	LocaleKey              = "locale"
	PhoneNumberKey         = "phone_number"
	PhoneNumberVerifiedKey = "phone_number_verified"
	AddressKey             = "address"
	UpdatedAtKey           = "updated_at"
)

type Token interface {
	Audience() []string
	Expiration() time.Time
	IssuedAt() time.Time
	Issuer() string
	JwtID() string
	NotBefore() time.Time
	Subject() string
	Name() string
	GivenName() string
	MiddleName() string
	FamilyName() string
	Nickname() string
	PreferredUsername() string
	Profile() string
	Picture() string
	Website() string
	Email() string
	EmailVerified() bool
	Gender() string
	Birthdate() *BirthdateClaim
	Zoneinfo() string
	Locale() string
	PhoneNumber() string
	PhoneNumberVerified() bool
	Address() *AddressClaim
	UpdatedAt() time.Time
	PrivateClaims() map[string]interface{}
	Get(string) (interface{}, bool)
	Set(string, interface{}) error
	Iterate(context.Context) Iterator
	Walk(context.Context, Visitor) error
	AsMap(context.Context) (map[string]interface{}, error)
}
type stdToken struct {
	audience            types.StringList   // https://tools.ietf.org/html/rfc7519#section-4.1.3
	expiration          *types.NumericDate // https://tools.ietf.org/html/rfc7519#section-4.1.4
	issuedAt            *types.NumericDate // https://tools.ietf.org/html/rfc7519#section-4.1.6
	issuer              *string            // https://tools.ietf.org/html/rfc7519#section-4.1.1
	jwtID               *string            // https://tools.ietf.org/html/rfc7519#section-4.1.7
	notBefore           *types.NumericDate // https://tools.ietf.org/html/rfc7519#section-4.1.5
	subject             *string            // https://tools.ietf.org/html/rfc7519#section-4.1.2
	name                *string            //
	givenName           *string            //
	middleName          *string            //
	familyName          *string            //
	nickname            *string            //
	preferredUsername   *string            //
	profile             *string            //
	picture             *string            //
	website             *string            //
	email               *string            //
	emailVerified       *bool              //
	gender              *string            //
	birthdate           *BirthdateClaim    //
	zoneinfo            *string            //
	locale              *string            //
	phoneNumber         *string            //
	phoneNumberVerified *bool              //
	address             *AddressClaim      //
	updatedAt           *types.NumericDate //
	privateClaims       map[string]interface{}
}

type openidTokenMarshalProxy struct {
	Xaudience            types.StringList   `json:"aud,omitempty"`
	Xexpiration          *types.NumericDate `json:"exp,omitempty"`
	XissuedAt            *types.NumericDate `json:"iat,omitempty"`
	Xissuer              *string            `json:"iss,omitempty"`
	XjwtID               *string            `json:"jti,omitempty"`
	XnotBefore           *types.NumericDate `json:"nbf,omitempty"`
	Xsubject             *string            `json:"sub,omitempty"`
	Xname                *string            `json:"name,omitempty"`
	XgivenName           *string            `json:"given_name,omitempty"`
	XmiddleName          *string            `json:"middle_name,omitempty"`
	XfamilyName          *string            `json:"family_name,omitempty"`
	Xnickname            *string            `json:"nickname,omitempty"`
	XpreferredUsername   *string            `json:"preferred_username,omitempty"`
	Xprofile             *string            `json:"profile,omitempty"`
	Xpicture             *string            `json:"picture,omitempty"`
	Xwebsite             *string            `json:"website,omitempty"`
	Xemail               *string            `json:"email,omitempty"`
	XemailVerified       *bool              `json:"email_verified,omitempty"`
	Xgender              *string            `json:"gender,omitempty"`
	Xbirthdate           *BirthdateClaim    `json:"birthdate,omitempty"`
	Xzoneinfo            *string            `json:"zoneinfo,omitempty"`
	Xlocale              *string            `json:"locale,omitempty"`
	XphoneNumber         *string            `json:"phone_number,omitempty"`
	XphoneNumberVerified *bool              `json:"phone_number_verified,omitempty"`
	Xaddress             *AddressClaim      `json:"address,omitempty"`
	XupdatedAt           *types.NumericDate `json:"updated_at,omitempty"`
}

// New creates a standard token, with minimal knowledge of
// possible claims. Standard claims include"aud", "exp", "iat", "iss", "jti", "nbf", "sub", "name", "given_name", "middle_name", "family_name", "nickname", "preferred_username", "profile", "picture", "website", "email", "email_verified", "gender", "birthdate", "zoneinfo", "locale", "phone_number", "phone_number_verified", "address" and "updated_at".
// Convenience accessors are provided for these standard claims
func New() Token {
	return &stdToken{
		privateClaims: make(map[string]interface{}),
	}
}

// Size returns the number of valid claims stored in this token
func (t *stdToken) Size() int {
	var count int
	if len(t.audience) > 0 {
		count++
	}
	if t.birthdate != nil {
		count++
	}
	if t.address != nil {
		count++
	}
	count += len(t.privateClaims)
	return count
}

func (t *stdToken) Get(name string) (interface{}, bool) {
	switch name {
	case AudienceKey:
		if t.audience == nil {
			return nil, false
		}
		v := t.audience.Get()
		return v, true
	case ExpirationKey:
		if t.expiration == nil {
			return nil, false
		}
		v := t.expiration.Get()
		return v, true
	case IssuedAtKey:
		if t.issuedAt == nil {
			return nil, false
		}
		v := t.issuedAt.Get()
		return v, true
	case IssuerKey:
		if t.issuer == nil {
			return nil, false
		}
		v := *(t.issuer)
		return v, true
	case JwtIDKey:
		if t.jwtID == nil {
			return nil, false
		}
		v := *(t.jwtID)
		return v, true
	case NotBeforeKey:
		if t.notBefore == nil {
			return nil, false
		}
		v := t.notBefore.Get()
		return v, true
	case SubjectKey:
		if t.subject == nil {
			return nil, false
		}
		v := *(t.subject)
		return v, true
	case NameKey:
		if t.name == nil {
			return nil, false
		}
		v := *(t.name)
		return v, true
	case GivenNameKey:
		if t.givenName == nil {
			return nil, false
		}
		v := *(t.givenName)
		return v, true
	case MiddleNameKey:
		if t.middleName == nil {
			return nil, false
		}
		v := *(t.middleName)
		return v, true
	case FamilyNameKey:
		if t.familyName == nil {
			return nil, false
		}
		v := *(t.familyName)
		return v, true
	case NicknameKey:
		if t.nickname == nil {
			return nil, false
		}
		v := *(t.nickname)
		return v, true
	case PreferredUsernameKey:
		if t.preferredUsername == nil {
			return nil, false
		}
		v := *(t.preferredUsername)
		return v, true
	case ProfileKey:
		if t.profile == nil {
			return nil, false
		}
		v := *(t.profile)
		return v, true
	case PictureKey:
		if t.picture == nil {
			return nil, false
		}
		v := *(t.picture)
		return v, true
	case WebsiteKey:
		if t.website == nil {
			return nil, false
		}
		v := *(t.website)
		return v, true
	case EmailKey:
		if t.email == nil {
			return nil, false
		}
		v := *(t.email)
		return v, true
	case EmailVerifiedKey:
		if t.emailVerified == nil {
			return nil, false
		}
		v := *(t.emailVerified)
		return v, true
	case GenderKey:
		if t.gender == nil {
			return nil, false
		}
		v := *(t.gender)
		return v, true
	case BirthdateKey:
		if t.birthdate == nil {
			return nil, false
		}
		v := t.birthdate
		return v, true
	case ZoneinfoKey:
		if t.zoneinfo == nil {
			return nil, false
		}
		v := *(t.zoneinfo)
		return v, true
	case LocaleKey:
		if t.locale == nil {
			return nil, false
		}
		v := *(t.locale)
		return v, true
	case PhoneNumberKey:
		if t.phoneNumber == nil {
			return nil, false
		}
		v := *(t.phoneNumber)
		return v, true
	case PhoneNumberVerifiedKey:
		if t.phoneNumberVerified == nil {
			return nil, false
		}
		v := *(t.phoneNumberVerified)
		return v, true
	case AddressKey:
		if t.address == nil {
			return nil, false
		}
		v := t.address
		return v, true
	case UpdatedAtKey:
		if t.updatedAt == nil {
			return nil, false
		}
		v := t.updatedAt.Get()
		return v, true
	default:
		v, ok := t.privateClaims[name]
		return v, ok
	}
}

func (t *stdToken) Set(name string, value interface{}) error {
	switch name {
	case AudienceKey:
		var acceptor types.StringList
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, AudienceKey)
		}
		t.audience = acceptor
		return nil
	case ExpirationKey:
		var acceptor types.NumericDate
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, ExpirationKey)
		}
		t.expiration = &acceptor
		return nil
	case IssuedAtKey:
		var acceptor types.NumericDate
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, IssuedAtKey)
		}
		t.issuedAt = &acceptor
		return nil
	case IssuerKey:
		if v, ok := value.(string); ok {
			t.issuer = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, IssuerKey, value)
	case JwtIDKey:
		if v, ok := value.(string); ok {
			t.jwtID = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, JwtIDKey, value)
	case NotBeforeKey:
		var acceptor types.NumericDate
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, NotBeforeKey)
		}
		t.notBefore = &acceptor
		return nil
	case SubjectKey:
		if v, ok := value.(string); ok {
			t.subject = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, SubjectKey, value)
	case NameKey:
		if v, ok := value.(string); ok {
			t.name = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, NameKey, value)
	case GivenNameKey:
		if v, ok := value.(string); ok {
			t.givenName = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, GivenNameKey, value)
	case MiddleNameKey:
		if v, ok := value.(string); ok {
			t.middleName = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, MiddleNameKey, value)
	case FamilyNameKey:
		if v, ok := value.(string); ok {
			t.familyName = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, FamilyNameKey, value)
	case NicknameKey:
		if v, ok := value.(string); ok {
			t.nickname = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, NicknameKey, value)
	case PreferredUsernameKey:
		if v, ok := value.(string); ok {
			t.preferredUsername = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, PreferredUsernameKey, value)
	case ProfileKey:
		if v, ok := value.(string); ok {
			t.profile = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, ProfileKey, value)
	case PictureKey:
		if v, ok := value.(string); ok {
			t.picture = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, PictureKey, value)
	case WebsiteKey:
		if v, ok := value.(string); ok {
			t.website = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, WebsiteKey, value)
	case EmailKey:
		if v, ok := value.(string); ok {
			t.email = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, EmailKey, value)
	case EmailVerifiedKey:
		if v, ok := value.(bool); ok {
			t.emailVerified = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, EmailVerifiedKey, value)
	case GenderKey:
		if v, ok := value.(string); ok {
			t.gender = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, GenderKey, value)
	case BirthdateKey:
		var acceptor BirthdateClaim
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, BirthdateKey)
		}
		t.birthdate = &acceptor
		return nil
	case ZoneinfoKey:
		if v, ok := value.(string); ok {
			t.zoneinfo = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, ZoneinfoKey, value)
	case LocaleKey:
		if v, ok := value.(string); ok {
			t.locale = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, LocaleKey, value)
	case PhoneNumberKey:
		if v, ok := value.(string); ok {
			t.phoneNumber = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, PhoneNumberKey, value)
	case PhoneNumberVerifiedKey:
		if v, ok := value.(bool); ok {
			t.phoneNumberVerified = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, PhoneNumberVerifiedKey, value)
	case AddressKey:
		var acceptor AddressClaim
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, AddressKey)
		}
		t.address = &acceptor
		return nil
	case UpdatedAtKey:
		var acceptor types.NumericDate
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, UpdatedAtKey)
		}
		t.updatedAt = &acceptor
		return nil
	default:
		if t.privateClaims == nil {
			t.privateClaims = map[string]interface{}{}
		}
		t.privateClaims[name] = value
	}
	return nil
}

func (t *stdToken) Audience() []string {
	if t.audience != nil {
		return t.audience.Get()
	}
	return nil
}

func (t *stdToken) Expiration() time.Time {
	if t.expiration != nil {
		return t.expiration.Get()
	}
	return time.Time{}
}

func (t *stdToken) IssuedAt() time.Time {
	if t.issuedAt != nil {
		return t.issuedAt.Get()
	}
	return time.Time{}
}

func (t *stdToken) Issuer() string {
	if t.issuer != nil {
		return *(t.issuer)
	}
	return ""
}

func (t *stdToken) JwtID() string {
	if t.jwtID != nil {
		return *(t.jwtID)
	}
	return ""
}

func (t *stdToken) NotBefore() time.Time {
	if t.notBefore != nil {
		return t.notBefore.Get()
	}
	return time.Time{}
}

func (t *stdToken) Subject() string {
	if t.subject != nil {
		return *(t.subject)
	}
	return ""
}

func (t *stdToken) Name() string {
	if t.name != nil {
		return *(t.name)
	}
	return ""
}

func (t *stdToken) GivenName() string {
	if t.givenName != nil {
		return *(t.givenName)
	}
	return ""
}

func (t *stdToken) MiddleName() string {
	if t.middleName != nil {
		return *(t.middleName)
	}
	return ""
}

func (t *stdToken) FamilyName() string {
	if t.familyName != nil {
		return *(t.familyName)
	}
	return ""
}

func (t *stdToken) Nickname() string {
	if t.nickname != nil {
		return *(t.nickname)
	}
	return ""
}

func (t *stdToken) PreferredUsername() string {
	if t.preferredUsername != nil {
		return *(t.preferredUsername)
	}
	return ""
}

func (t *stdToken) Profile() string {
	if t.profile != nil {
		return *(t.profile)
	}
	return ""
}

func (t *stdToken) Picture() string {
	if t.picture != nil {
		return *(t.picture)
	}
	return ""
}

func (t *stdToken) Website() string {
	if t.website != nil {
		return *(t.website)
	}
	return ""
}

func (t *stdToken) Email() string {
	if t.email != nil {
		return *(t.email)
	}
	return ""
}

func (t *stdToken) EmailVerified() bool {
	if t.emailVerified != nil {
		return *(t.emailVerified)
	}
	return false
}

func (t *stdToken) Gender() string {
	if t.gender != nil {
		return *(t.gender)
	}
	return ""
}

func (t *stdToken) Birthdate() *BirthdateClaim {
	return t.birthdate
}

func (t *stdToken) Zoneinfo() string {
	if t.zoneinfo != nil {
		return *(t.zoneinfo)
	}
	return ""
}

func (t *stdToken) Locale() string {
	if t.locale != nil {
		return *(t.locale)
	}
	return ""
}

func (t *stdToken) PhoneNumber() string {
	if t.phoneNumber != nil {
		return *(t.phoneNumber)
	}
	return ""
}

func (t *stdToken) PhoneNumberVerified() bool {
	if t.phoneNumberVerified != nil {
		return *(t.phoneNumberVerified)
	}
	return false
}

func (t *stdToken) Address() *AddressClaim {
	return t.address
}

func (t *stdToken) UpdatedAt() time.Time {
	if t.updatedAt != nil {
		return t.updatedAt.Get()
	}
	return time.Time{}
}

func (t *stdToken) PrivateClaims() map[string]interface{} {
	return t.privateClaims
}

func (t *stdToken) iterate(ctx context.Context, ch chan *ClaimPair) {
	defer close(ch)

	var pairs []*ClaimPair
	if t.audience != nil {
		v := t.audience.Get()
		pairs = append(pairs, &ClaimPair{Key: AudienceKey, Value: v})
	}
	if t.expiration != nil {
		v := t.expiration.Get()
		pairs = append(pairs, &ClaimPair{Key: ExpirationKey, Value: v})
	}
	if t.issuedAt != nil {
		v := t.issuedAt.Get()
		pairs = append(pairs, &ClaimPair{Key: IssuedAtKey, Value: v})
	}
	if t.issuer != nil {
		v := *(t.issuer)
		pairs = append(pairs, &ClaimPair{Key: IssuerKey, Value: v})
	}
	if t.jwtID != nil {
		v := *(t.jwtID)
		pairs = append(pairs, &ClaimPair{Key: JwtIDKey, Value: v})
	}
	if t.notBefore != nil {
		v := t.notBefore.Get()
		pairs = append(pairs, &ClaimPair{Key: NotBeforeKey, Value: v})
	}
	if t.subject != nil {
		v := *(t.subject)
		pairs = append(pairs, &ClaimPair{Key: SubjectKey, Value: v})
	}
	if t.name != nil {
		v := *(t.name)
		pairs = append(pairs, &ClaimPair{Key: NameKey, Value: v})
	}
	if t.givenName != nil {
		v := *(t.givenName)
		pairs = append(pairs, &ClaimPair{Key: GivenNameKey, Value: v})
	}
	if t.middleName != nil {
		v := *(t.middleName)
		pairs = append(pairs, &ClaimPair{Key: MiddleNameKey, Value: v})
	}
	if t.familyName != nil {
		v := *(t.familyName)
		pairs = append(pairs, &ClaimPair{Key: FamilyNameKey, Value: v})
	}
	if t.nickname != nil {
		v := *(t.nickname)
		pairs = append(pairs, &ClaimPair{Key: NicknameKey, Value: v})
	}
	if t.preferredUsername != nil {
		v := *(t.preferredUsername)
		pairs = append(pairs, &ClaimPair{Key: PreferredUsernameKey, Value: v})
	}
	if t.profile != nil {
		v := *(t.profile)
		pairs = append(pairs, &ClaimPair{Key: ProfileKey, Value: v})
	}
	if t.picture != nil {
		v := *(t.picture)
		pairs = append(pairs, &ClaimPair{Key: PictureKey, Value: v})
	}
	if t.website != nil {
		v := *(t.website)
		pairs = append(pairs, &ClaimPair{Key: WebsiteKey, Value: v})
	}
	if t.email != nil {
		v := *(t.email)
		pairs = append(pairs, &ClaimPair{Key: EmailKey, Value: v})
	}
	if t.emailVerified != nil {
		v := *(t.emailVerified)
		pairs = append(pairs, &ClaimPair{Key: EmailVerifiedKey, Value: v})
	}
	if t.gender != nil {
		v := *(t.gender)
		pairs = append(pairs, &ClaimPair{Key: GenderKey, Value: v})
	}
	if t.birthdate != nil {
		v := t.birthdate
		pairs = append(pairs, &ClaimPair{Key: BirthdateKey, Value: v})
	}
	if t.zoneinfo != nil {
		v := *(t.zoneinfo)
		pairs = append(pairs, &ClaimPair{Key: ZoneinfoKey, Value: v})
	}
	if t.locale != nil {
		v := *(t.locale)
		pairs = append(pairs, &ClaimPair{Key: LocaleKey, Value: v})
	}
	if t.phoneNumber != nil {
		v := *(t.phoneNumber)
		pairs = append(pairs, &ClaimPair{Key: PhoneNumberKey, Value: v})
	}
	if t.phoneNumberVerified != nil {
		v := *(t.phoneNumberVerified)
		pairs = append(pairs, &ClaimPair{Key: PhoneNumberVerifiedKey, Value: v})
	}
	if t.address != nil {
		v := t.address
		pairs = append(pairs, &ClaimPair{Key: AddressKey, Value: v})
	}
	if t.updatedAt != nil {
		v := t.updatedAt.Get()
		pairs = append(pairs, &ClaimPair{Key: UpdatedAtKey, Value: v})
	}
	for k, v := range t.privateClaims {
		pairs = append(pairs, &ClaimPair{Key: k, Value: v})
	}
	for _, pair := range pairs {
		select {
		case <-ctx.Done():
			return
		case ch <- pair:
		}
	}
}

func (t *stdToken) UnmarshalJSON(buf []byte) error {
	var proxy openidTokenMarshalProxy
	if err := json.Unmarshal(buf, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal stdToken`)
	}
	t.audience = proxy.Xaudience
	t.expiration = proxy.Xexpiration
	t.issuedAt = proxy.XissuedAt
	t.issuer = proxy.Xissuer
	t.jwtID = proxy.XjwtID
	t.notBefore = proxy.XnotBefore
	t.subject = proxy.Xsubject
	t.name = proxy.Xname
	t.givenName = proxy.XgivenName
	t.middleName = proxy.XmiddleName
	t.familyName = proxy.XfamilyName
	t.nickname = proxy.Xnickname
	t.preferredUsername = proxy.XpreferredUsername
	t.profile = proxy.Xprofile
	t.picture = proxy.Xpicture
	t.website = proxy.Xwebsite
	t.email = proxy.Xemail
	t.emailVerified = proxy.XemailVerified
	t.gender = proxy.Xgender
	t.birthdate = proxy.Xbirthdate
	t.zoneinfo = proxy.Xzoneinfo
	t.locale = proxy.Xlocale
	t.phoneNumber = proxy.XphoneNumber
	t.phoneNumberVerified = proxy.XphoneNumberVerified
	t.address = proxy.Xaddress
	t.updatedAt = proxy.XupdatedAt
	var m map[string]interface{}
	if err := json.Unmarshal(buf, &m); err != nil {
		return errors.Wrap(err, `failed to parse private parameters`)
	}
	delete(m, AudienceKey)
	delete(m, ExpirationKey)
	delete(m, IssuedAtKey)
	delete(m, IssuerKey)
	delete(m, JwtIDKey)
	delete(m, NotBeforeKey)
	delete(m, SubjectKey)
	delete(m, NameKey)
	delete(m, GivenNameKey)
	delete(m, MiddleNameKey)
	delete(m, FamilyNameKey)
	delete(m, NicknameKey)
	delete(m, PreferredUsernameKey)
	delete(m, ProfileKey)
	delete(m, PictureKey)
	delete(m, WebsiteKey)
	delete(m, EmailKey)
	delete(m, EmailVerifiedKey)
	delete(m, GenderKey)
	delete(m, BirthdateKey)
	delete(m, ZoneinfoKey)
	delete(m, LocaleKey)
	delete(m, PhoneNumberKey)
	delete(m, PhoneNumberVerifiedKey)
	delete(m, AddressKey)
	delete(m, UpdatedAtKey)
	t.privateClaims = m
	return nil
}

func (t stdToken) MarshalJSON() ([]byte, error) {
	var proxy openidTokenMarshalProxy
	proxy.Xaudience = t.audience
	proxy.Xexpiration = t.expiration
	proxy.XissuedAt = t.issuedAt
	proxy.Xissuer = t.issuer
	proxy.XjwtID = t.jwtID
	proxy.XnotBefore = t.notBefore
	proxy.Xsubject = t.subject
	proxy.Xname = t.name
	proxy.XgivenName = t.givenName
	proxy.XmiddleName = t.middleName
	proxy.XfamilyName = t.familyName
	proxy.Xnickname = t.nickname
	proxy.XpreferredUsername = t.preferredUsername
	proxy.Xprofile = t.profile
	proxy.Xpicture = t.picture
	proxy.Xwebsite = t.website
	proxy.Xemail = t.email
	proxy.XemailVerified = t.emailVerified
	proxy.Xgender = t.gender
	proxy.Xbirthdate = t.birthdate
	proxy.Xzoneinfo = t.zoneinfo
	proxy.Xlocale = t.locale
	proxy.XphoneNumber = t.phoneNumber
	proxy.XphoneNumberVerified = t.phoneNumberVerified
	proxy.Xaddress = t.address
	proxy.XupdatedAt = t.updatedAt
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(proxy); err != nil {
		return nil, errors.Wrap(err, `failed to encode proxy to JSON`)
	}
	hasContent := buf.Len() > 3 // encoding/json always adds a newline, so "{}\n" is the empty hash
	if l := len(t.privateClaims); l > 0 {
		buf.Truncate(buf.Len() - 2)
		keys := make([]string, 0, l)
		for k := range t.privateClaims {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			if hasContent || i > 0 {
				fmt.Fprintf(&buf, `,`)
			}
			fmt.Fprintf(&buf, `%s:`, strconv.Quote(k))
			if err := enc.Encode(t.privateClaims[k]); err != nil {
				return nil, errors.Wrapf(err, `failed to encode private param %s`, k)
			}
		}
		fmt.Fprintf(&buf, `}`)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
		return nil, errors.Wrap(err, `failed to do second pass unmarshal during MarshalJSON`)
	}
	return json.Marshal(m)
}

func (t *stdToken) Iterate(ctx context.Context) Iterator {
	ch := make(chan *ClaimPair)
	go t.iterate(ctx, ch)
	return mapiter.New(ch)
}

func (t *stdToken) Walk(ctx context.Context, visitor Visitor) error {
	return iter.WalkMap(ctx, t, visitor)
}

func (t *stdToken) AsMap(ctx context.Context) (map[string]interface{}, error) {
	return iter.AsMap(ctx, t)
}
