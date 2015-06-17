package main

type cacheModeSetting int

const (
	// Possible cache modes
	cacheModeUnset       = iota // cache mode has not been set
	cacheModeOn                 // cache everything
	cacheModeDevelopment        // cache everything, except Amber, Lua, GCSS and Markdown
	cacheModeProduction         // cache everything, except Amber and Lua
	cacheModeImages             // cache images (png, jpg, gif, svg)
	cacheModeBinary             // cache all binary files (zip, png etc. but not svg)
	cacheModeOff                // cache nothing
)

const cacheModeDefault = cacheModeOn

var (
	// Table of cache mode setting names
	cacheModeNames = map[cacheModeSetting]string{
		cacheModeUnset:       "unset",
		cacheModeOn:          "On",
		cacheModeDevelopment: "Development",
		cacheModeProduction:  "Production",
		cacheModeImages:      "Images",
		cacheModeBinary:      "Binary",
		cacheModeOff:         "Off",
	}
)

// Create a new cache mode setting
func NewCacheModeSetting(cacheModeString string) cacheModeSetting {
	switch cacheModeString {
	case "everything", "all", "on", "1", "enabled", "yes", "enable": //- Cache everything.
		return cacheModeOn
	case "production", "prod": // - Cache everything, except: Amber and Lua.
		return cacheModeProduction
	case "images", "image": // - Cache images (png, jpg, gif, svg).
		return cacheModeImages
	case "binary", "bin": // - Cache all binary files.
		return cacheModeBinary
	case "off", "disabled", "0", "no", "disable": // - Disable caching.
		return cacheModeOff
	case "dev", "default", "unset": //- Cache everything, except: Amber, Lua, GCSS and Markdown.
		fallthrough
	default:
		return cacheModeDefault
	}
}

// Return the name of the cache mode setting, if set
func (cms cacheModeSetting) String() string {
	for k, v := range cacheModeNames {
		if k == cms {
			return v
		}
	}
	// Could not find the name
	return cacheModeNames[cacheModeUnset]
}

// Return true of the given file type (extension) should be cached
func (cmd cacheModeSetting) ShouldCache(ext string) bool {
	switch cacheMode {
	case cacheModeOn:
		return true
	case cacheModeProduction:
		switch ext {
		case ".amber", ".lua":
			return false
		default:
			return true
		}
	case cacheModeImages:
		switch ext {
		case ".png", ".jpg", ".gif", ".svg", ".jpeg", ".ico", ".bmp", ".apng":
			return true
		default:
			return false
		}
	case cacheModeBinary:
		// TODO: This test should be rewritten to take the file data into consideration instead.
		switch ext {
		case ".txt", ".html", ".css", ".js", ".xml", ".htm", ".gcss", ".amber", ".lua", ".md", ".text", ".nfo", ".jsx":
			return false
		default:
			return true
		}
	case cacheModeOff:
		return false
	case cacheModeDevelopment, cacheModeUnset:
		fallthrough
	default:
		switch ext {
		case ".amber", ".lua", ".md", ".gcss", ".jsx":
			return false
		default:
			return true
		}
	}
}