HTMLUNWRAP(1)

# NAME

htmlunwrap - unwrap elements in an HTML document

# SYNOPSIS

*htmlunwrap* SELECTOR [_FILE_]...

# DESCRIPTION

Reads each file in sequence and removes all elements matching SELECTOR. The
children of those elements will be added to their grandparents, thus unwrapping
them from their previous parent element. Prints resulting documents to standard
output. If no FILE is given or FILE is -, read standard input.

If any FILE cannot be processed, a message prefixed with the FILE name will be
written to standard error.

# AUTHOR

Lukas Henkel <lh@entf.net>
