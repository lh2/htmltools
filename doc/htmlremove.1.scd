HTMLREMOVE(1)

# NAME

htmlremove - remove elements from an HTML document

# SYNOPSIS

*htmlremove* SELECTOR [_FILE_]...

# DESCRIPTION

Reads each file in sequence and removes all elements matching SELECTOR. Prints
resulting documents to standard output. If no FILE is given or FILE is -, read
standard input.

If any FILE cannot be processed, a message prefixed with the FILE name will be
written to standard error.

# AUTHOR

Lukas Henkel <lh@entf.net>
