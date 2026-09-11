//go:build !windows

package splash

type otherPlatformSplash struct{}

func newPlatformSplash() platformSplash {
	return &otherPlatformSplash{}
}

func (s *otherPlatformSplash) Start() error {
	return nil
}

func (s *otherPlatformSplash) SetStatus(message string) {}

func (s *otherPlatformSplash) Close() {}
