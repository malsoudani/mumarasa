package BetterCustomer;

use strict;
use warnings;
use Class::Tiny qw(first_name last_name), {
    timestamp => sub { time },
};

sub new { bless {}, shift}

sub full_name {
    my $self = shift;
    return join " ", $self->first_name, $self->last_name;
}

1;
