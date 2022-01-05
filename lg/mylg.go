package lg

import "fmt"

type lg interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type Profile struct {
	name        string
	rootLevel   LgLevel
	appenderRef []Appender
}

func NewProfile(name string, rootLevel LgLevel) *Profile {
	profile := &Profile{
		name:        name,
		rootLevel:   rootLevel,
		appenderRef: make([]Appender, 10),
	}
	return profile
}

func (p *Profile) AddAppender(appender Appender) {
	p.appenderRef = append(p.appenderRef, appender)
}

type Mylg struct {
	Oo map[string]Profile
}

func NewMylg() *Mylg {
	return &Mylg{
		Oo: make(map[string]Profile),
	}
}

func (m *Mylg) AddProfile(profile *Profile) {
	m.Oo[profile.name] = *profile
}

func (lg *Mylg) log(level LgLevel, format string, args ...interface{}) {
	for _, profile := range lg.Oo {
		if level >= profile.rootLevel {
			sprintf := fmt.Sprintf(format, args)
			fmt.Println(sprintf)
		}
	}
}

func (lg *Mylg) Debug(format string, args ...interface{}) {
	lg.log(Debug, format, args)
}

func (lg *Mylg) Info(format string, args ...interface{}) {
	lg.log(Info, format, args)
}

func (lg *Mylg) Error(format string, args ...interface{}) {
	lg.log(Error, format, args)
}
