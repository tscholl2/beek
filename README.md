# beek
Peek into binary files

A command line utility that can peek into binary files and display values in multiple formats. Useful to see if a file containing a bit stream is all 0's or something.

    NAME:
       beek - Peek at binary files.

    EXAMPLES:
        # read first 16 bytes of test.bin in hex format
        beek test.bin --format x -l 16
    	# read last 8 bytes of test.bin in binary format
    	beek --format b test.bin

    USAGE:
       beek [global options] command [command options] [arguments...]

    VERSION:
       0.0.0

    COMMANDS:
       help, h	Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --length, -l "8"	Number of bytes to read from file. Default: 8.
       --format, -s "d"	Format of output:
                'x' : hex
                'X' : hex with capitol letters
                'd' : base10
                'a' : ascii
       --delimiter, -d "\\n"    delimiter. Default is newline character
       --offset, -o "0"	offset of reading in file. 0 for head, 10 for starting with 10th byte. offset of -10 reads last 10 bytes starting at -10,-9,-8,...
       --help, -h		show help
       --version, -v	print the version
