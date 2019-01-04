
use strict;
use warnings;
use LWP::Simple;

$|=1;

sub main {
    my @files = (
        '/Users/malsoudani/personal/mumarasa/perl/DownloadCrawler.pl',
        '/Users/malsoudani/personal/mumarasa/perl/DownloadCrawle.pl',
        '/Users/malsoudani/personal/mumarasa/perl/HelloWorld.pl',
    );

    foreach my $file(@files) {
        -f $file ? print "found file $file\n" : print "did not find $file\n";
    }
}

main();
