go-libnecpp
===========

A Go wrapper around libnecpp, which is a C library for 
nec++, which is a C++ rewrite of nec2c, which is itself a C rewrite/translation
of the original FORTRAN code for NEC2, an antenna modelling package.

Installation
------------

This package requires nec++ to be installed. On Linux, it's often found in the distribution packages. Otherwise, and on other Unix-y OSes, it should build per the instructions in the necpp library.

Installing nec++ on Mac OS X is likely to be more difficult, because clang gets unhappy about some of the Fortranisms still creeping around in the code. Even when that's cleared up by applying a couple of the pull requests on the github page, you still need to configure nec++ with the `--without-lapack` option. A procudure for installing necpp smoothly is in the works.

This package should be able to build on Windows, assuming you've been able to build nec++ with the instructions for Windows on the github page.

Documentation
-------------

See the godocs for more detailed documentation for go-libnecpp.

Also very relevant:

* nec++'s github page can be found at https://github.com/tmolteno/necpp/.
* The nec++ documentation at http://tmolteno.github.io/necpp/
* Neoklis Kyriazis' website at http://www.qsl.net/5b4az/, which includes the homepage for nec2c.
* The user's manual for NEC2 at http://www.nec2.org/part_3/toc.html. This manual may make some of the more obscure portions of the documentation taken from nec++ more clear.

LICENSE
-------

This library is licensed under the GPL (since the original libnecpp library is). Details can be found in the LICENSE file in this repository.

AUTHOR
------

The golang libnecpp wrapper was written by Jeremy Bingham (<jeremy@goiardi.gl>).

Nec++ was written by Tim Molteno <tim@physics.otago.ac.nz>.

Nec2c (the C port which nec++ was based on) was done by by Neoklis Kyriazis.
