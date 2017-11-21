package sourcemap

func Schema() string {
	return sourcemapSchema
}

var sourcemapSchema = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "docs/spec/sourcemaps/sourcemap-wrapper.json",
    "title": "Sourcemap Wrapper",
    "description": "Sourcemap + Metadata",
    "type": "object",
    "properties": {
        "bundle_filepath": {
            "description": "relative path of the minified bundle file",
            "type": "string",
            "maxLength": 1024,
            "minLength": 1
        },
        "app": {
            "allOf": [
                {     "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "doc/spec/app_name.json",
    "title": "App name property",
    "type": "object",
    "properties": {
        "name": {
            "description": "Immutable name of the app emitting this event",
            "type": "string",
            "pattern": "^[a-zA-Z0-9 _-]+$",
            "maxLength": 1024
        }
    },
    "required": ["name"] },
                { "properties": {
                    "version": {
                        "description": "Version of the app emitting this event",
                        "type": "string",
                        "maxLength": 1024,
                        "minLength": 1
                    }
                },
                "required": ["version"]
                }
            ]
        }
    },
    "required": ["bundle_filepath"]
}
`