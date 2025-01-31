package interop

// Interop represents the interoperability interface between Go and other languages
type Interop struct {
	language string
	filePath string
}

// NewInteropRunner creates a new InteropRunner instance
func NewInteropRunner(interop Interop) InteropRunner {
	switch interop.language {
	case "javascript", "js":
		return &JavascriptRunner{I: interop}
	case "java":
		return &JavaRunner{I: interop}
	case "python":
		return &PythonRunner{I: interop}
	default:
		return &UnknownRunner{}
	}
}
