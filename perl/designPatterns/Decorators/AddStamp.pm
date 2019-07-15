 package AddStamp;

 use strict;
 use warnings;

 sub TIEHANDLE {
     my ($class, $handle) = @_;
     return bless \$handle, $class;
 }

 sub PRINT {
     my $handle = shift;
     my $stamp = localtime();

     print $handle, "$stamp ", @_;
 }

 sub PRINTF { goto &PRINT }


 sub CLOSE {
     my $self = shift;
     close $self;
 }

1;
