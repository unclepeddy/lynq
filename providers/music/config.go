package main

type Config struct {
	music	*Service
}

// loadAndValidate ensures we have proper credentials for communicating
// with all backends required for this provider
func (c *Config) loadAndValidate(userId string) error {
	c.music = NewService(userId)
	return nil
}
