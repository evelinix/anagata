package splash

type NativeSplash struct {
	platformSplash platformSplash
}

type platformSplash interface {
	Start() error
	SetStatus(message string)
	Close()
}

func NewNativeSplash() *NativeSplash {
	return &NativeSplash{
		platformSplash: newPlatformSplash(),
	}
}

func (s *NativeSplash) Start() error {
	return s.platformSplash.Start()
}

func (s *NativeSplash) SetStatus(message string) {
	s.platformSplash.SetStatus(message)
}

func (s *NativeSplash) Close() {
	s.platformSplash.Close()
}
