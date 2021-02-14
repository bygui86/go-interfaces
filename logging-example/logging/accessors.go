package logging

func (c *Config) Encoding() string {
	return c.encoding
}

func (c *Config) Level() string {
	return c.level
}

func (c *Config) OutputPath() string {
	return c.outputPath
}

func (c *Config) ErrOutputPath() string {
	return c.errOutputPath
}

// DEFAULTS

func EncodingDefault() string {
	return encodingDefault
}

func LevelDefault() string {
	return levelDefault
}

func OutputPathDefault() string {
	return outPathDefault
}

func ErrOutputPathDefault() string {
	return errOutPathDefault
}
