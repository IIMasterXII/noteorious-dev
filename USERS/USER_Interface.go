package USERS

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"time"
	"strconv"
	"mime/multipart"
	"errors"
	"google.golang.org/appengine"
	"github.com/Esseh/retrievable"
	"golang.org/x/net/context"
	"github.com/Esseh/notorious-dev/CORE"
	"github.com/Esseh/notorious-dev/COOKIE"
	"github.com/Esseh/notorious-dev/CLOUD"
)


// Uploads an avatar for a user into the cloud.
func UploadAvatar(ctx context.Context, userID int64, header *multipart.FileHeader, avatarReader io.ReadSeeker) error {
	m, _, err := image.Decode(avatarReader)
	if err != nil { return err }
	imageBounds := m.Bounds()
	if imageBounds.Dy() > CORE.MaxAvatarSize || imageBounds.Dx() > CORE.MaxAvatarSize {
		return errors.New("Filesize is Too Large")
	}
	avatarReader.Seek(0, 0)
	filename := CORE.GetAvatarPath(userID)
	return CLOUD.AddFile(ctx, filename, header.Header["Content-Type"][0], avatarReader)
}



// Retrieves an AUTH_User from the currently logged in user.
func GetUserFromSession(req *http.Request) (*User, error) {
	userID, err := GetUserIDFromRequest(req)
	if err != nil { return &User{}, err }
	ctx := appengine.NewContext(req)
	return GetUserFromID(ctx, userID)
}

// Retrieves an AUTH_User ID from the currently logged in user.
func GetUserIDFromRequest(req *http.Request) (int64, error) {
	s, err := GetSessionID(req)
	if err != nil { return 0, err }
	ctx := appengine.NewContext(req)
	userID, err := GetUserIDFromSession(ctx, s)
	if err != nil { return 0, err }
	return userID, nil
}

// Retireves an USERS_User from it's respective ID.
func GetUserFromID(ctx context.Context, userID int64) (*User, error) {
	u := &User{}
	getErr := retrievable.GetEntity(ctx, retrievable.IntID(userID), u)
	return u, getErr
}

// Retrieves an USERS_User ID from a AUTH_Session ID
func GetUserIDFromSession(ctx context.Context, sessionID int64) (userID int64, _ error) {
	sessionData, err := GetSession(ctx, sessionID)
	if err != nil { return 0, err }
	return sessionData.UserID, nil
}

// Retrieves an USERS_Session from its respective ID.
func GetSession(ctx context.Context, sessionID int64) (Session, error) {
	s := Session{}
	getErr := retrievable.GetEntity(ctx, sessionID, &s) // Get actual session from datastore
	if getErr != nil { return Session{}, errors.New("Not Logged In") }
	s.LastUsed = time.Now()
	if _, err := retrievable.PlaceEntity(ctx, sessionID, &s); err != nil { return Session{}, err }
	return s, nil
}

// Retrieves a Session ID from the currently logged in user.
func GetSessionID(req *http.Request) (int64, error) {
	sessionIDStr, err := COOKIE.GetValue(req, "session")
	if err != nil { return -1, errors.New("Not Logged In") }
	id, err := strconv.ParseInt(sessionIDStr, 10, 64) // Change cookie val into key
	if err != nil { return -1, errors.New("Invalid Logged In") }
	return id, nil
}