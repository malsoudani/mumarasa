use strict;
use warnings;
use LWP::Simple;

sub main {
    print "Downloading ... \n";
    # print get("http://www.google.com");

    my $httpCode = getstore("https://croatia.hr/sites/default/files/styles/image_full_width/public/2017-08/01_01_slide_nature.jpg", "nature.jpg");
    if ($httpCode == 200) {
        print "Success\n";
    } else {
        print $httpCode;
    }
    print "Finished";
    
}

main();
