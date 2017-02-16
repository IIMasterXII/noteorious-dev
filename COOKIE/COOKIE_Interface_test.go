package COOKIE
import(
	"testing"
	"net/http/httptest"
)
func TestMake(t *testing.T){
	res := httptest.NewRecorder()
	
	testing.Coverage()
}
func TestGetValue(t *testing.T){
	res := httptest.NewRecorder()
	testing.Coverage()
}
func TestDelete(t *testing.T){
	res := httptest.NewRecorder()
	testing.Coverage()
}

/*
// Initializes a cookie into the current session.
func Make(res http.ResponseWriter, name, value string) error {
	mac, err := CORE.CreateHmac(value)
	if err != nil { return err }
	c := &http.Cookie{
		Name:     name,
		Value:    value + "." + base64.RawURLEncoding.EncodeToString(mac),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   CORE.SessionTime,
	}
	http.SetCookie(res, c)
	return nil
}

// Retrieves the value located inside of a cookie.
func GetValue(req *http.Request, name string) (string, error) {
	cookie, err := req.Cookie(name)
	if err != nil {
		return "", err
	}
	val, mac := CORE.SplitMac(cookie.Value)
	if good := CORE.CheckMac(val, mac); !good {
		return "", errors.New("Hmac checking failed")
	}
	return val, nil
}
// Deletes a cookie held in the current session by name.
func Delete(res http.ResponseWriter, name string) {
	http.SetCookie(res, &http.Cookie{
		Name:   name,
		MaxAge: -1,
		Path:   "/",
	})
}

*/