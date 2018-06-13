What?
=====

(eventually) a tar-like tool for extracting stuff from fat filesystems.

I'm working on this for my own personal learning and interest - please do not
use this as a source of truth for how to work with FAT, and please do not use
this in any kind of production environment.

Can't be bothered to put a LICENSE file in, so: It's MIT licensed. Please heed
especially the part in the MIT license about there being no warranty for this
software.

Progress
========

FAT32
-----

* [ ] Correctly read 8.3 file entries from a directory
* [ ] Correctly read long name file entries from a directory
* [ ] Given a mock cluster map & fat header, correctly read a file
* [ ] Correctly read the root directory of a filesystem
* [ ] Correctly read a file from an actual filesystem

Stretch goals
-------------

* [ ] Extend support to exFAT
* [ ] Extend support to FAT16
* [ ] Extend support to FAT12
* [ ] Add a file to a FAT filesystem
