HTMLINDENTHEADINGS(1)

# NAME

htmlindentheadings - indents all headings by a specified amount

# SYNOPSIS

*htmlindentheadings* INDENT_LEVELS [_FILE_]...

# DESCRIPTION

Indents the heading elements (h1 to h7) of all files by the specified level. For
example, if the specified indent level is 2, all h1 elements would become h3
elements, all h2 elements would become h4 elements and so on. htmlindentheadings
does not create invalid HTML, so it will not go beyond h7. INDENT_LEVELS can
also be a negative number, to decrease the indentation of heading elements.

# AUTHOR

Lukas Henkel <lh@entf.net>
