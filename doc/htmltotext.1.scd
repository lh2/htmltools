HTMLTOTEXT(1)

# NAME

htmltotext - extract all text from an HTML document

# SYNOPSIS

*htmltotext* [_FILE_]...

# DESCRIPTION

Reads each file in sequence and prints all text without the HTML tags to
standard output. If no FILE is given or FILE is -, read standard input.

If any FILE cannot be processed, a message prefixed with the FILE name will be
written to standard error.

# AUTHOR

Lukas Henkel <lh@entf.net>
