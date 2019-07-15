#!/usr/bin/perl
use strict;
use warnings;
use lib ".";
use AddStamp;

open LOG, ">output.tmp" or die "Couldn't write output.tmp: $!\n";
tie *STAMPED_LOG, "AddStamp", *LOG;

while (<>) {
    print STAMPED_LOG;
}

close STAMPED_LOG;
