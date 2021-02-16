HTMLATTR(1)

# NAME

htmlattr - print specific attribute values from top level nodes

# SYNOPSIS

*htmlattr* [-fs FIELD_SEPERATOR] ATTRIBUTES [_FILE_]...
# OPTIONS

*-fs*
	character that will be used to both split the ATTRIBUTES argument and
	seperate the output columns. "," by default.


# DESCRIPTION

Reads each file in sequence, iterates over all top level nodes and prints
requested attributes for each element line by line. If an element does not
contain any attributes requested, no line is printed. If some but not all
attributes exist on a node, the respective columns will be empty. htmlattr
prints the same number of columns for every node. If no FILE is given or FILE
is -, read standard input.

If any FILE cannot be processed, a message prefixed with the FILE name will be
written to standard error.

# AUTHOR

Lukas Henkel <lh@entf.net>
