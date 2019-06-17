HTMLSELECT(1)

# NAME

htmlselect - print all elements matching a selector from an HTML document

# SYNOPSIS

*htmlselect* SELECTOR [_FILE_]...

# DESCRIPTION

Reads each file in sequence and prints all elements matching SELECTOR to
standard output. If no FILE is given or FILE is -, read standard input.

If any FILE cannot be processed, a message prefixed with the FILE name will be
written to standard error.

# AUTHOR

Lukas Henkel <lh@entf.net>
