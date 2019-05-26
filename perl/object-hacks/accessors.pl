use strict;
use warnings;
use lib '/home/malsoudani/Code/mumarasa/perl/object-hacks';
use Customer;
use BetterCustomer;

sub main {
    my $customer = BetterCustomer->new();
    $customer->first_name('Mohamed');
    $customer->last_name('Alsoudani');
    print $customer->full_name;
    print $customer->timestamp;
}

main();
