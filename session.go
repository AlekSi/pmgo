package pmgo

import mgo "gopkg.in/mgo.v2"

type SessionManager interface {
	BuildInfo() (info mgo.BuildInfo, err error)
	Clone() SessionManager
	Close()
	Copy() SessionManager
	DB(name string) DatabaseManager
	DatabaseNames() (names []string, err error)
	EnsureSafe(safe *mgo.Safe)
	// FindRef(ref *DBRef) *Query
	// Fsync(async bool) error
	// FsyncLock() error
	// FsyncUnlock() error
	// LiveServers() (addrs []string)
	// Login(cred *Credential) error
	// LogoutAll()
	// Mode() Mode
	// New() *SessionManager
	Ping() error
	// Refresh()
	// ResetIndexCache()
	Run(cmd interface{}, result interface{}) error
	// Safe() (safe *Safe)
	// SelectServers(tags ...bson.D)
	// SetBatch(n int)
	// SetBypassValidation(bypass bool)
	// SetCursorTimeout(d time.Duration)
	SetMode(consistency mgo.Mode, refresh bool)
	// SetPoolLimit(limit int)
	// SetPrefetch(p float64)
	// SetSafe(safe *Safe)
	// SetSocketTimeout(d time.Duration)
	// SetSyncTimeout(d time.Duration)
}

type Session struct {
	session *mgo.Session
}

// This methos allows to use mgo's dbtest.DBServer in pmgo tests.
// Example:
// var Server dbtest.DBServer
// tempDir, _ := ioutil.TempDir("", "testing")
// Server.SetPath(tempDir)
// session := NewSessionManager(Server.Session())
func NewSessionManager(s *mgo.Session) SessionManager {
	return &Session{
		session: s,
	}
}

func (s *Session) BuildInfo() (info mgo.BuildInfo, err error) {
	return s.session.BuildInfo()
}

func (s *Session) Close() {
	s.session.Close()
}

func (s *Session) Clone() SessionManager {
	return &Session{
		session: s.session.Clone(),
	}
}

func (s *Session) Copy() SessionManager {
	return &Session{
		session: s.session.Copy(),
	}
}

func (s *Session) DB(name string) DatabaseManager {
	d := &Database{
		db: s.session.DB(name),
	}
	return d
}

func (s *Session) DatabaseNames() (names []string, err error) {
	return s.session.DatabaseNames()
}

func (s *Session) EnsureSafe(safe *mgo.Safe) {
	s.session.EnsureSafe(safe)
}

func (s *Session) Run(cmd interface{}, result interface{}) error {
	return s.session.Run(cmd, result)
}

func (s *Session) Ping() error {
	return s.session.Ping()
}

func (s *Session) SetMode(consistency mgo.Mode, refresh bool) {
	s.session.SetMode(consistency, refresh)
}
